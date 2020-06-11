package api

import (
	"context"
	"github.com/emirpasic/gods/maps/treemap"
)

type OnReceiveProposal interface {
	BeforeProposedBlockAdded(ctx context.Context, proposal Proposal) error
	AfterProposedBlockAdded(ctx context.Context, proposal Proposal, receipts []Receipt) error
	BeforeVoted(ctx context.Context, vote Vote) error
	AfterVoted(ctx context.Context, vote Vote) error
}

type OnProposal interface {
	OnProposal(ctx context.Context, proposal Proposal) error
}

type OnVoteReceived interface {
	OnVoteReceived(ctx context.Context, vote Vote) error
	OnQCFinished(ctx context.Context, qc QuorumCertificate) error
}

type OnNewBlockCreated interface {
	OnNewBlockCreated(ctx context.Context, builder BlockBuilder, receipts []Receipt) (Block, error)
}

type OnNextView interface {
	OnNewView(ctx context.Context, newView int32) error
}

type OnNextEpoch interface {
	OnNewEpoch(ctx context.Context, newEpoch int32) error
}

type OnBlockCommit interface {
	OnBlockCommit(ctx context.Context, block Block, orphans *treemap.Map) error
}

type NullOnReceiveProposal struct{}

func (NullOnReceiveProposal) BeforeProposedBlockAdded(ctx context.Context, proposal Proposal) error {
	return nil
}

func (NullOnReceiveProposal) AfterProposedBlockAdded(ctx context.Context, proposal Proposal, receipts []Receipt) error {
	return nil
}

func (NullOnReceiveProposal) BeforeVoted(ctx context.Context, vote Vote) error {
	return nil
}

func (NullOnReceiveProposal) AfterVoted(ctx context.Context, vote Vote) error {
	return nil
}

type NullOnVoteReceived struct{}

func (NullOnVoteReceived) OnVoteReceived(ctx context.Context, vote Vote) error {
	return nil
}

func (NullOnVoteReceived) OnQCFinished(ctx context.Context, qc QuorumCertificate) error {
	return nil
}

type NullOnNextView struct{}

func (NullOnNextView) OnNewView(ctx context.Context, newView int32) error {
	return nil
}

type NullOnNextEpoch struct{}

func (NullOnNextEpoch) OnNewEpoch(ctx context.Context, newEpoch int32) error {
	return nil
}

type NullOnBlockCommit struct{}

func (NullOnBlockCommit) OnBlockCommit(ctx context.Context, block Block, orphans *treemap.Map) error {
	return nil
}

type NullOnProposal struct{}

func (NullOnProposal) OnProposal(ctx context.Context, proposal Proposal) error {
	return nil
}

type NullOnNewBlockCreated struct{}

func (NullOnNewBlockCreated) OnNewBlockCreated(ctx context.Context, builder BlockBuilder, receipts []Receipt) (Block, error) {
	return nil, nil
}
