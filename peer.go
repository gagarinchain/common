package common

import (
	"bytes"
	"github.com/gagarinchain/common/eth/common"
	"github.com/gagarinchain/common/eth/crypto"
	pb "github.com/gagarinchain/common/protobuff"
	"github.com/libp2p/go-libp2p-core/peer"
)

type Peer struct {
	address    common.Address
	publicKey  *crypto.PublicKey
	privateKey *crypto.PrivateKey
	peerInfo   *peer.AddrInfo
}

func (p *Peer) SetPublicKey(publicKey *crypto.PublicKey) {
	p.publicKey = publicKey
}

func (p *Peer) SetPrivateKey(privateKey *crypto.PrivateKey) {
	p.privateKey = privateKey
}

func (p *Peer) GetAddress() common.Address {
	return p.address
}

func (p *Peer) SetAddress(address common.Address) {
	p.address = address
}

func (p *Peer) GetPrivateKey() *crypto.PrivateKey {
	return p.privateKey
}

func CreatePeer(publicKey *crypto.PublicKey, privateKey *crypto.PrivateKey, peerInfo *peer.AddrInfo) *Peer {
	peer := &Peer{
		publicKey:  publicKey,
		privateKey: privateKey,
		peerInfo:   peerInfo,
	}

	if publicKey != nil {
		peer.address = common.BytesToAddress(publicKey.Bytes())
	}
	return peer
}

func (p *Peer) Equals(toCompare *Peer) bool {
	return bytes.Compare(p.address.Bytes(), toCompare.address.Bytes()) == 0
}

func (p *Peer) GetPeerInfo() *peer.AddrInfo {
	return p.peerInfo
}

func (p *Peer) PublicKey() *crypto.PublicKey {
	return p.publicKey
}

func (p *Peer) ToStorageProto() *pb.Peer {
	json, err := p.peerInfo.MarshalJSON()
	if err != nil {
		log.Error("Error while serializing peerInfo", err)
		return nil
	}
	return &pb.Peer{
		Address:   p.address.Bytes(),
		PublicKey: p.publicKey.Bytes(),
		AddrInfo:  json,
	}
}

func CreatePeerFromStorage(p *pb.Peer) *Peer {
	info := &peer.AddrInfo{}
	if err := info.UnmarshalJSON(p.AddrInfo); err != nil {
		log.Error("Can't unmarshal addrinfo")
		return nil
	}
	return &Peer{
		address:    common.BytesToAddress(p.Address),
		publicKey:  crypto.FromBytes(p.PublicKey),
		privateKey: nil,
		peerInfo:   info,
	}
}

func PeersToPubs(peers []*Peer) (pubs []*crypto.PublicKey) {
	for _, p := range peers {
		pubs = append(pubs, p.PublicKey())
	}

	return pubs
}
