package common

import (
	"encoding/hex"
	"encoding/json"
	"github.com/gagarinchain/common/eth/common"
	"github.com/gagarinchain/common/eth/crypto"
	"github.com/gagarinchain/common/eth/crypto/bls12_381"
	p2pcrypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/phoreproject/bls/g1pubs"
	"io/ioutil"
)

type CommitteeLoader interface {
	LoadPeerListFromFile(filePath string) []*Peer
	LoadPeerFromFile(fileName string, peer *Peer) (peerKey p2pcrypto.PrivKey, err error)
}

type CommitteeData struct {
	Peers []PeerData `json:"peers"`
}

type PeerData struct {
	Address      string `json:"addr"`
	Pub          string `json:"pub"`
	MultiAddress string `json: ma"`
}

type CommitteeLoaderImpl struct {
}

func (c *CommitteeLoaderImpl) LoadPeerListFromFile(filePath string) (res []*Peer) {
	byteValue, e := ioutil.ReadFile(filePath)
	if e != nil {
		log.Fatal("Can't load committee list", e)
		return nil
	}

	var data CommitteeData
	if err := json.Unmarshal(byteValue, &data); err != nil {
		log.Fatal("Can't unmarshal committee file", err)
		return nil
	}

	for _, v := range data.Peers {
		address := common.HexToAddress(v.Address)
		pub := common.Hex2Bytes(v.Pub)
		addr, e := ma.NewMultiaddr(v.MultiAddress)
		if e != nil {
			log.Fatal("Can't create address", e)
			return nil
		}

		info, e := peer.AddrInfoFromP2pAddr(addr)
		if e != nil {
			log.Fatal("Can't create address", e)
			return nil
		}
		key, _ := g1pubs.DeserializePublicKey(bls12_381.ToBytes48(pub))
		res = append(res, &Peer{
			address:   address,
			publicKey: crypto.NewPublicKey(key),
			peerInfo:  info,
		})

	}

	return res
}

func (c *CommitteeLoaderImpl) LoadPeerFromFile(fileName string, peer *Peer) (peerKey p2pcrypto.PrivKey, err error) {
	var v map[string]interface{}

	bytes, e := ioutil.ReadFile(fileName)
	if e != nil {
		return nil, e
	}
	e = json.Unmarshal(bytes, &v)
	if e != nil {
		return nil, e
	}

	// First let's create a new identity key pair for our node. If this was your
	// application you would likely save this private key to a database and load
	// it from the db on subsequent start ups.
	pkpeer := v["pkpeer"].(string)
	addr := v["addr"].(string)
	pub := v["pub"].(string)
	pk := v["pk"].(string)

	decodeString, e := hex.DecodeString(pkpeer)
	privKey, err := p2pcrypto.UnmarshalPrivateKey(decodeString)
	if err != nil {
		return nil, err
	}

	pkbytes, e := hex.DecodeString(pk)
	if e != nil {
		return nil, e
	}
	pubbytes, e := hex.DecodeString(pub)
	if e != nil {
		return nil, e
	}

	key := g1pubs.DeserializeSecretKey(bls12_381.ToBytes32(pkbytes))
	publicKey, _ := g1pubs.DeserializePublicKey(bls12_381.ToBytes48(pubbytes))

	peer.SetAddress(common.HexToAddress(addr))
	peer.SetPublicKey(crypto.NewPublicKey(publicKey))
	peer.SetPrivateKey(crypto.NewPrivateKey(key))

	return privKey, nil
}
