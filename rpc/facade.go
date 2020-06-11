package rpc

import (
	"github.com/emirpasic/gods/maps/treemap"
	"github.com/gagarinchain/common/api"
)

type PluginsFacade struct {
	onReceiveProposal api.OnReceiveProposal
	onProposal        api.OnProposal
	onVoteReceived    api.OnVoteReceived
	onNewBlockCreated api.OnNewBlockCreated
	onNextView        api.OnNextView
	onNextEpoch       api.OnNextEpoch
	blockCommit       api.OnBlockCommit
}

func NewPluginsFacade(onReceiveProposal api.OnReceiveProposal, onProposal api.OnProposal, onVoteReceived api.OnVoteReceived,
	onNewBlockCreated api.OnNewBlockCreated, onNextView api.OnNextView, onNextEpoch api.OnNextEpoch, blockCommit api.OnBlockCommit) *PluginsFacade {
	return &PluginsFacade{
		onReceiveProposal: onReceiveProposal,
		onProposal:        onProposal,
		onVoteReceived:    onVoteReceived,
		onNewBlockCreated: onNewBlockCreated,
		onNextView:        onNextView,
		onNextEpoch:       onNextEpoch,
		blockCommit:       blockCommit}
}

func (f *PluginsFacade) OnBlockCommit(bc api.Blockchain, block api.Block, orphans *treemap.Map) error {
	panic("implement me")
}

func (f *PluginsFacade) OnNewEpoch(pacer api.Pacer, bc api.Blockchain, newEpoch int32) error {
	panic("implement me")
}

func (f *PluginsFacade) OnNewView(pacer api.Pacer, bc api.Blockchain, newView int32) error {
	panic("implement me")
}

func (f *PluginsFacade) OnBlockCreated(bc api.Blockchain, builder api.BlockBuilder, receipts []api.Receipt) error {
	panic("implement me")
}

func (f *PluginsFacade) OnVoteReceived(pacer api.Pacer, bc api.Blockchain, vote api.Vote) error {
	panic("implement me")
}

func (f *PluginsFacade) OnQCFinished(pacer api.Pacer, bc api.Blockchain, qc api.QuorumCertificate) error {
	panic("implement me")
}

func (f *PluginsFacade) OnProposal(pacer api.Pacer, bc api.Blockchain, proposal api.Proposal) error {
	panic("implement me")
}

func (f *PluginsFacade) BeforeProposedBlockAdded(pacer api.Pacer, bc api.Blockchain, proposal api.Proposal) error {
	panic("implement me")
}

func (f *PluginsFacade) AfterProposedBlockAdded(pacer api.Pacer, bc api.Blockchain, proposal api.Proposal, receipts []api.Receipt) error {
	panic("implement me")
}

func (f *PluginsFacade) BeforeVoted(pacer api.Pacer, bc api.Blockchain, vote api.Vote) error {
	panic("implement me")
}

func (f *PluginsFacade) AfterVoted(pacer api.Pacer, bc api.Blockchain, vote api.Vote) error {
	panic("implement me")
}
