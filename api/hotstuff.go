package api

import (
	"github.com/gagarinchain/common"
	cmn "github.com/gagarinchain/common/eth/common"
	"github.com/gagarinchain/common/eth/crypto"
	pb "github.com/gagarinchain/common/protobuff"
	"math/big"
)

type ProposerForHeight interface {
	ProposerForHeight(blockHeight int32) *common.Peer
	GetBitmap(src map[cmn.Address]*crypto.Signature) (bitmap *big.Int)
	GetPeers() []*common.Peer
}

type EventNotifier interface {
	SubscribeProtocolEvents(sub chan Event)
	FireEvent(event Event)
	NotifyEvent(event Event)
}

type Pacer interface {
	EventNotifier
	ProposerForHeight
	GetCurrentView() int32
	GetCurrentEpoch() int32
	GetCurrent() *common.Peer
	GetNext() *common.Peer
}

type Proposal interface {
	Cert() Certificate
	Sender() *common.Peer
	NewBlock() Block
	Signature() *crypto.Signature
	GetMessage() *pb.ProposalPayload
	Sign(key *crypto.PrivateKey)
}

type Sync interface {
	Cert() Certificate
	Height() int32
	Voting() bool
	Sign(key *crypto.PrivateKey)
	GetMessage() *pb.SynchronizePayload
	Sender() *common.Peer
	Signature() *crypto.Signature
	VotingSignature() *crypto.Signature
}

type Vote interface {
	Cert() Certificate
	Sign(key *crypto.PrivateKey)
	GetMessage() *pb.VotePayload
	Sender() *common.Peer
	Header() Header
	Signature() *crypto.Signature
}

type EventPayload interface{}
type EventType int

const (
	TimedOut            EventType = iota
	EpochStarted        EventType = iota
	EpochStartTriggered EventType = iota
	Voted               EventType = iota
	VotesCollected      EventType = iota
	Proposed            EventType = iota
	ChangedView         EventType = iota
	HCUpdated           EventType = iota
	SCCreated           EventType = iota
)

// We intentionally duplicated subset of common.Events here, simply to isolate protocol logic
type Event struct {
	Payload EventPayload
	T       EventType
}

type EventHandler = func(event Event)
