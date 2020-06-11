package api

import (
	"github.com/gagarinchain/common/eth/common"
	pb "github.com/gagarinchain/common/protobuff"
	"github.com/gagarinchain/common/trie/sparse"
	"math/big"
)

type Account interface {
	Balance() *big.Int
	Nonce() uint64
	Copy() Account
	Voters() []common.Address
	Origin() common.Address
	SetOrigin(origin common.Address)
	AddVoters(from common.Address)
	IncrementNonce()
	Serialize() []byte
	ToStorageProto() *pb.Account
}

//record stores information about version tree structure and provides base account processing logic
type Record interface {
	Pending() Record
	SetPending(pending Record)
	NewPendingRecord(proposer common.Address) Record
	Get(address common.Address) (acc Account, found bool)
	GetForUpdate(address common.Address) (acc Account, found bool)
	Proof(address common.Address) (proof *sparse.Proof)
	RootProof() common.Hash
	SetParent(parent Record)
	Parent() Record
	Serialize() []byte
	Hash() common.Hash
	Trie() *sparse.SMT
	AddSibling(record Record)
	SetHash(pending common.Hash)
	Siblings() []Record
	ApplyTransaction(t Transaction) (r []Receipt, err error)
	Update(address common.Address, account Account) error
	Put(address common.Address, account Account)
}

type Receipt interface {
	Value() *big.Int
	To() common.Address
	From() common.Address
	TxIndex() int32
	TxHash() common.Hash
	FromValue() *big.Int
	ToValue() *big.Int
	ToStorageProto() *pb.Receipt
}
