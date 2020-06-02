package api

import (
	"github.com/emirpasic/gods/maps/treemap"
)

type OnReceiveProposal interface {
	BeforeProposedBlockAdded(pacer Pacer, bc Blockchain, proposal Proposal) error
	AfterProposedBlockAdded(pacer Pacer, bc Blockchain, proposal Proposal, receipts []Receipt) error
	BeforeVoted(pacer Pacer, bc Blockchain, vote Vote) error
	AfterVoted(pacer Pacer, bc Blockchain, vote Vote) error
}

type OnProposal interface {
	OnProposal(pacer Pacer, bc Blockchain, proposal Proposal) error
}

type OnVoteReceived interface {
	OnVoteReceived(pacer Pacer, bc Blockchain, vote Vote) error
	OnQCFinished(pacer Pacer, bc Blockchain, qc QuorumCertificate) error
}

type OnNewBlockCreated interface {
	OnBlockCreated(bc Blockchain, builder BlockBuilder, receipts []Receipt) error
}

type OnNextView interface {
	OnNewView(pacer Pacer, bc Blockchain, newView int32) error
}

type OnNextEpoch interface {
	OnNewEpoch(pacer Pacer, bc Blockchain, newEpoch int32) error
}

type OnBlockCommit interface {
	OnBlockCommit(bc Blockchain, block Block, orphans *treemap.Map) error
}

type NullOnReceiveProposal struct{}

func (NullOnReceiveProposal) BeforeProposedBlockAdded(pacer Pacer, bc Blockchain, proposal Proposal) error {
	return nil
}

func (NullOnReceiveProposal) AfterProposedBlockAdded(pacer Pacer, bc Blockchain, proposal Proposal, receipts []Receipt) error {
	return nil
}

func (NullOnReceiveProposal) BeforeVoted(pacer Pacer, bc Blockchain, vote Vote) error {
	return nil
}

func (NullOnReceiveProposal) AfterVoted(pacer Pacer, bc Blockchain, vote Vote) error {
	return nil
}

type NullOnVoteReceived struct{}

func (NullOnVoteReceived) OnVoteReceived(pacer Pacer, bc Blockchain, vote Vote) error {
	return nil
}

func (NullOnVoteReceived) OnQCFinished(pacer Pacer, bc Blockchain, qc QuorumCertificate) error {
	return nil
}

type NullOnNextView struct{}

func (NullOnNextView) OnNewView(pacer Pacer, bc Blockchain, newView int32) error {
	return nil
}

type NullOnNextEpoch struct{}

func (NullOnNextEpoch) OnNewEpoch(pacer Pacer, bc Blockchain, newEpoch int32) error {
	return nil
}

type NullOnBlockCommit struct{}

func (NullOnBlockCommit) OnBlockCommit(bc Blockchain, block Block, orphans *treemap.Map) error {
	return nil
}

type NullOnProposal struct{}

func (NullOnProposal) OnProposal(pacer Pacer, bc Blockchain, proposal Proposal) error {
	return nil
}

type NullOnNewBlockCreated struct{}

func (NullOnNewBlockCreated) OnBlockCreated(bc Blockchain, builder BlockBuilder, receipts []Receipt) error {
	return nil
}
