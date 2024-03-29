package network

import (
	"context"
	"fmt"
	"github.com/gagarinchain/common"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	disc "github.com/libp2p/go-libp2p-discovery"
	"github.com/libp2p/go-libp2p-kad-dht"
	"github.com/op/go-logging"
	"sync"
	"time"
)

var log = logging.MustGetLogger("bootstrap")

type BootstrapConfig struct {
	//Minimum amount of peers needed to start correctly
	MinPeerThreshold int

	// Period duration for bootstrap repetition.
	// Should set up it carefully since bootstrap process could be resource heavy
	Period time.Duration

	// Simply connection timeout
	ConnectionTimeout time.Duration

	//discover cid
	RendezvousNs string

	//Peers we started with, base peers used to discover additional ones
	InitialPeers []peer.AddrInfo
}

//Default parameters for bootstrapping
var DefaultBootstrapConfig = BootstrapConfig{
	MinPeerThreshold:  3,
	Period:            10 * time.Second,
	ConnectionTimeout: (10 * time.Second) / 3,
}

func bootstrapWithPeers(committee []*common.Peer) BootstrapConfig {
	peers := make([]peer.AddrInfo, len(committee))
	for i, c := range committee {
		peers[i] = *c.GetPeerInfo()
	}

	cfg := DefaultBootstrapConfig
	cfg.InitialPeers = peers
	return cfg
}

//Start bus services here
//Note that now we simply connect to every node in committee, and actually don't use peer discovery via DHT.
//It is considered that sooner we will have small amount of bootstrapped nodes and bigger amount of committee which can be found using DHT routing tables and connected to
func Bootstrap(ctx context.Context, peerHost host.Host, routing *dht.IpfsDHT, cfg BootstrapConfig) (statusChan chan int, errChan chan error) {
	statusChan = make(chan int)
	errChan = make(chan error)

	err := routing.Bootstrap(ctx)

	if err != nil {
		go func() {
			errChan <- err
		}()
	}

	d := disc.NewRoutingDiscovery(routing)
	//Start DHT
	disc.Advertise(ctx, d, cfg.RendezvousNs)

	watchdog := &Watchdog{
		host:      peerHost,
		cfg:       cfg,
		discovery: d,
		res:       statusChan,
		err:       errChan,
	}
	go watchdog.watch(ctx)

	return statusChan, errChan
}

type Watchdog struct {
	host      host.Host
	discovery *disc.RoutingDiscovery
	cfg       BootstrapConfig
	res       chan int
	err       chan error
}

func (w *Watchdog) watch(ctx context.Context) {
	t := time.NewTicker(w.cfg.Period)
	go func() {
		for {
			select {
			case <-t.C:
				if err := bootstrapTick(ctx, w.host, w.discovery, w.cfg); err != nil {
					w.err <- err
				} else {
					w.res <- 0
				}
			case <-ctx.Done():
				t.Stop()
				return
			}
		}
	}()
	if err := bootstrapTick(ctx, w.host, w.discovery, w.cfg); err != nil {
		w.err <- err
	}
}

func bootstrapTick(ctx context.Context, host host.Host, discovery *disc.RoutingDiscovery, cfg BootstrapConfig) error {
	ctx, cancel := context.WithTimeout(ctx, cfg.ConnectionTimeout)
	defer cancel()
	id := host.ID()
	peers := cfg.InitialPeers

	//manage bootstrap connections
	connected := host.Network().Peers()
	toConnect := cfg.MinPeerThreshold - len(connected)
	if toConnect <= 0 {
		log.Debugf("Connected to %d (> %d) nodes", len(connected), cfg.MinPeerThreshold)
		return nil
	}

	// filter out connected nodes
	var notConnected []peer.AddrInfo
	for _, p := range peers {
		if host.Network().Connectedness(p.ID) != network.Connected {
			notConnected = append(notConnected, p)
		}
	}

	if len(notConnected) > 0 {
		log.Debugf("%s connecting to %d not connected bootstrap nodes: %v", id, len(notConnected), notConnected)
		connected := bootstrapConnect(ctx, host, notConnected)
		toConnect = toConnect - connected
	}

	if toConnect > 0 {
		log.Debugf("try to discover and connect to %d nodes", toConnect)

		toConnect = toConnect - discoverAndConnect(ctx, host, discovery, cfg, toConnect)
	}

	if toConnect > 0 {
		return fmt.Errorf("can't connect to all peers, need %v more", toConnect)
	}

	return nil
}

func bootstrapConnect(ctx context.Context, ph host.Host, peers []peer.AddrInfo) int {
	if len(peers) < 1 {
		return 0
	}

	errs := make(chan error, len(peers))
	var wg sync.WaitGroup
	for _, p := range peers {
		wg.Add(1)

		//running startup routines
		go func(p peer.AddrInfo) {
			defer wg.Done()
			log.Debugf("%s connecting %s", ph.ID(), p.ID)
			log.Debugf("Got addresses [%v] for peer [%s]", p.Addrs, p.ID)

			ph.Peerstore().AddAddrs(p.ID, p.Addrs, peerstore.PermanentAddrTTL)
			if err := ph.Connect(ctx, p); err != nil {
				log.Debugf("failed to connect with %v: %s", p.ID, err)
				errs <- err
				return
			}
			//TODO refresh committee peer addresses here or later, to have their actual addresses
			log.Infof("connected %v", p.ID)
		}(p)
	}
	wg.Wait()

	// cleaning error channel
	close(errs)
	count := 0
	var err error
	for err = range errs {
		if err != nil {
			count++
		}
	}

	return len(peers) - count
}

// Subscribe returns a new Subscription for the given topic.
// While subscribing we calculate topics cid and provide this value to underlying DHT-router
func discoverAndConnect(ctx context.Context, ph host.Host, d *disc.RoutingDiscovery, cfg BootstrapConfig, toConnect int) int {
	timeout, _ := context.WithTimeout(ctx, cfg.ConnectionTimeout)

	//take 10 providers to connect
	provs, err := disc.FindPeers(timeout, d, cfg.RendezvousNs)
	if err != nil {
		return 0
	}

	// filter out connected nodes
	var notConnected []peer.AddrInfo
	for _, p := range provs {
		if ph.Network().Connectedness(p.ID) != network.Connected {
			notConnected = append(notConnected, p)
		}
		if len(notConnected) > toConnect {
			break
		}
	}

	if len(notConnected) < 1 {
		return 0
	}

	errs := make(chan error, len(notConnected))
	var wg sync.WaitGroup
	for _, p := range notConnected {
		wg.Add(1)

		//running startup routines
		go func(p peer.AddrInfo) {
			defer wg.Done()
			log.Debugf("%s connecting %s", ph.ID(), p.ID)
			log.Debugf("Got addresses [%v] for peer [%s]", p.Addrs, p.ID)

			ph.Peerstore().AddAddrs(p.ID, p.Addrs, peerstore.PermanentAddrTTL)
			if err := ph.Connect(ctx, p); err != nil {
				log.Debugf("failed to connect with %v: %s", p.ID, err)
				errs <- err
				return
			}
			log.Infof("connected %v", p.ID)
		}(p)
	}
	wg.Wait()

	// cleaning error channel
	close(errs)
	count := 0
	for err = range errs {
		if err != nil {
			count++
		}
	}

	return len(notConnected) - count
}

func IntMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}
