package api

import (
	"github.com/emirpasic/gods/maps/treemap"
	"github.com/gagarinchain/common/eth/common"
	"github.com/gagarinchain/common/eth/crypto"
	pb "github.com/gagarinchain/common/protobuff"
	"github.com/gagarinchain/common/trie"
	"math/big"
	"time"
)

type Blockchain interface {
	GetBlockByHash(hash common.Hash) (block Block)
	GetBlockByHeight(height int32) (res []Block)
	GetFork(height int32, headHash common.Hash) (res []Block)
	Contains(hash common.Hash) bool
	GetThreeChain(twoHash common.Hash) (zero Block, one Block, two Block)
	OnCommit(b Block) (toCommit []Block, orphans *treemap.Map, err error)
	GetHead() Block
	GetHeadRecord() Record
	GetTopHeight() int32
	GetTopHeightBlocks() []Block
	AddBlock(block Block) ([]Receipt, error)
	RemoveBlock(block Block) error
	GetGenesisBlock() Block
	GetGenesisCert() QuorumCertificate
	IsSibling(sibling Header, ancestor Header) bool
	NewBlock(parent Block, qc QuorumCertificate, data []byte) Block
	PadEmptyBlock(head Block, qc QuorumCertificate) Block
	GetGenesisBlockSignedHash(key *crypto.PrivateKey) *crypto.Signature
	ValidateGenesisBlockSignature(signature *crypto.Signature, address common.Address) bool
	GetTopCommittedBlock() Block
	UpdateGenesisBlockQC(certificate QuorumCertificate)
	SetProposerGetter(proposerGetter ProposerForHeight)
}

//Caution: every Set* method MUST be called with great aware since it can potentially change hash of the block.
type Block interface {
	TxsCount() int
	Txs() Iterator
	AddTransaction(t Transaction)
	Header() Header
	Data() []byte
	Height() int32
	QC() QuorumCertificate
	//Sets block QC to concrete value. Use this method with big care, since it mutates block's field that is used in hash calculation.
	//Obviously such a mutation must lead to blocks rehashing and rehashing would lead to update in all maps and this can be very dangerous.
	SetQC(qc QuorumCertificate)
	Signature() *crypto.SignatureAggregate
	//Signature is not used in hash calculation, this set can be called freely
	SetSignature(s *crypto.SignatureAggregate)
	QRef() Header
	GetMessage() *pb.Block
	ToStorageProto() *pb.BlockS
	Serialize() ([]byte, error)
}

type BlockBuilder interface {
	SetHeader(header Header) BlockBuilder
	SetQC(qc QuorumCertificate) BlockBuilder
	SetTxs(txs *trie.FixedLengthHexKeyMerkleTrie) BlockBuilder
	SetData(data []byte) BlockBuilder
	Header() Header
	QC() QuorumCertificate
	Txs() *trie.FixedLengthHexKeyMerkleTrie
	Data() []byte
	Build() Block
}

type Header interface {
	DataHash() common.Hash
	StateHash() common.Hash
	Height() int32
	Hash() common.Hash
	TxHash() common.Hash
	QCHash() common.Hash
	SetQCHash(hash common.Hash)
	Parent() common.Hash
	Timestamp() time.Time
	IsGenesisBlock() bool
	GetMessage() *pb.BlockHeader
	ToStorageProto() *pb.BlockHeaderS
	SetHash()
	Sign(key *crypto.PrivateKey) *crypto.Signature
}

type QuorumCertificate interface {
	SignatureAggregate() *crypto.SignatureAggregate
	QrefBlock() Header
	GetMessage() *pb.QuorumCertificate
	GetHash() common.Hash
	IsValid(qcHash common.Hash, committee []*crypto.PublicKey) (bool, error)
	ToStorageProto() *pb.QuorumCertificateS
}

type Transaction interface {
	Serialized() []byte
	Data() []byte
	Signature() *crypto.Signature
	Value() *big.Int
	Nonce() uint64
	From() common.Address
	To() common.Address
	SetTo(to common.Address)
	TxType() Type
	Fee() *big.Int
	SetFrom(from common.Address)
	CreateProof(pk *crypto.PrivateKey) (e error)
	RecoverProver() (aggregate *crypto.SignatureAggregate, e error)
	GetMessage() *pb.Transaction
	ToStorageProto() *pb.TransactionS
	Sign(key *crypto.PrivateKey)
	Hash() common.Hash
	DropSignature()
}

type Receipt interface {
	Value() *big.Int
	To() common.Address
	From() common.Address
	TxIndex() int32
	TxHash() common.Hash
	FromValue() *big.Int
	ToValue() *big.Int
}

type Iterator interface {
	Next() Transaction
	HasNext() bool
}

type Type int

const (
	SettlementAddressHex         = "0x6522b1ac0c0c078f1fcc696b9cf72c59bb3624b7d2a9d82059b2f3832fd9973d"
	DefaultSettlementReward      = 10 //probably this value should be set from config or via consensus or moved to different TX field
	DefaultAgreementFee          = 2
	Payment                 Type = iota
	Slashing                Type = iota
	Settlement              Type = iota
	Agreement               Type = iota
	Proof                   Type = iota
	Redeem                  Type = iota
)
