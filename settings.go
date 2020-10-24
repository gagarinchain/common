package common

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gagarinchain/common/eth/crypto"
	p2pcrypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/op/go-logging"
	"io/ioutil"
	"strconv"
)

var log = logging.MustGetLogger("hotstuff")

type Settings struct {
	Hotstuff struct {
		CommitteeSize int    `yaml:"CommitteeSize"`
		Me            int    `yaml:"Me"`
		Delta         int    `yaml:"Delta"`
		BlockDelta    int    `yaml:"BlockDelta"`
		SeedPath      string `yaml:"SeedPath"`
	} `yaml:"Hotstuff"`
	Network struct {
		ExtAddr           string `yaml:"ExtAddr"`
		MinPeerThreshold  int    `yaml:"MinPeerThreshold"`
		ReconnectPeriod   int    `yaml:"ReconnectPeriod"`
		ConnectionTimeout int    `yaml:"ConnectionTimeout"`
	} `yaml:"Network"`
	Storage struct {
		Dir string `yaml:"Dir"`
	} `yaml:"Storage"`
	Static struct {
		Dir string `yaml:"Dir"`
	} `yaml:"Static"`
	Rpc struct {
		Address              string `yaml:"Address"`
		MaxConcurrentStreams uint32 `yaml:"MaxConcurrentStreams"`
	} `yaml:"Rpc"`
	Log struct {
		Level string `yaml:"Level"`
	} `yaml:"Log"`
	Plugins []Plugin `yaml:"Plugins"`
}

type Plugin struct {
	Address    string   `yaml:"Address"`
	Interfaces []string `yaml:"Interfaces"`
}

func GenerateIdentities() {
	committee := CommitteeData{}

	for i := 0; i < 10; i++ {
		pk, _ := crypto.GenerateKey()
		pkbytes := pk.V().Serialize()
		pubbytes := pk.PublicKey().Bytes()
		pkstring := hex.EncodeToString(pkbytes[:])
		pubstring := hex.EncodeToString(pubbytes[:])
		address := crypto.PubkeyToAddress(pk.PublicKey())

		privKey, _, _ := p2pcrypto.GenerateSecp256k1Key(rand.Reader)
		id, _ := peer.IDFromPrivateKey(privKey)
		b, _ := p2pcrypto.MarshalPrivateKey(privKey)

		v := map[string]interface{}{"addr": address, "pk": pkstring, "pub": pubstring, "id": id.Pretty(), "pkpeer": hex.EncodeToString(b)}

		marshal, _ := json.Marshal(v)
		var out bytes.Buffer
		if err := json.Indent(&out, marshal, "", "\t"); err != nil {
			panic(err)
		}

		log.Info(out.String())
		err := ioutil.WriteFile("static/peer"+strconv.Itoa(i)+".json", out.Bytes(), 0644)

		if err != nil {
			panic(err)
		}

		multiaddr := fmt.Sprintf("/ip4/127.0.0.1/tcp/908%d/p2p/%s", i, id.Pretty())

		p := PeerData{
			Address:      address.Hex(),
			Pub:          pubstring,
			MultiAddress: multiaddr,
		}

		committee.Peers = append(committee.Peers, p)
	}

	marshal, e := json.Marshal(committee)
	if e != nil {
		panic(e)
	}

	var out bytes.Buffer
	if err := json.Indent(&out, marshal, "", "\t"); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile("static/peers.json", out.Bytes(), 0644); err != nil {
		panic(err)
	}

}
