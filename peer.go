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
	id, err := p.peerInfo.ID.Marshal()
	if err != nil {
		log.Error(err)
	}
	var addrs []string
	for _, addr := range p.peerInfo.Addrs {
		addrs = append(addrs, addr.String())
	}
	return &pb.Peer{
		Address:   p.address.Bytes(),
		PublicKey: p.publicKey.Bytes(),
		Id:        id,
		Addresses: addrs,
	}
}

func PeersToPubs(peers []*Peer) (pubs []*crypto.PublicKey) {
	for _, p := range peers {
		pubs = append(pubs, p.PublicKey())
	}

	return pubs
}
