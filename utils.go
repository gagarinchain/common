package common

import (
	"github.com/gagarinchain/common/eth/common"
	"github.com/gagarinchain/common/eth/crypto"
)

func GenerateAddress() common.Address {
	p1, err := crypto.GenerateKey()
	if err != nil {
		return common.Address{}
	}
	return crypto.PubkeyToAddress(p1.PublicKey())
}

func GenerateAddresses(n int) (a []common.Address) {
	for i := 0; i < n; i++ {
		a = append(a, GenerateAddress())
	}
	return a
}

func GeneratePeer() *Peer {
	p1, _ := crypto.GenerateKey()
	return CreatePeer(p1.PublicKey(), p1, nil)
}

func GeneratePeers(n int) (peers []*Peer) {
	for i := 0; i < n; i++ {
		peers = append(peers, GeneratePeer())
	}
	return peers
}
