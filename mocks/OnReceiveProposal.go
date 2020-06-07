// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	api "github.com/gagarinchain/common/api"
	mock "github.com/stretchr/testify/mock"
)

// OnReceiveProposal is an autogenerated mock type for the OnReceiveProposal type
type OnReceiveProposal struct {
	mock.Mock
}

// AfterProposedBlockAdded provides a mock function with given fields: pacer, bc, proposal, receipts
func (_m *OnReceiveProposal) AfterProposedBlockAdded(pacer api.Pacer, bc api.Blockchain, proposal api.Proposal, receipts []api.Receipt) error {
	ret := _m.Called(pacer, bc, proposal, receipts)

	var r0 error
	if rf, ok := ret.Get(0).(func(api.Pacer, api.Blockchain, api.Proposal, []api.Receipt) error); ok {
		r0 = rf(pacer, bc, proposal, receipts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AfterVoted provides a mock function with given fields: pacer, bc, vote
func (_m *OnReceiveProposal) AfterVoted(pacer api.Pacer, bc api.Blockchain, vote api.Vote) error {
	ret := _m.Called(pacer, bc, vote)

	var r0 error
	if rf, ok := ret.Get(0).(func(api.Pacer, api.Blockchain, api.Vote) error); ok {
		r0 = rf(pacer, bc, vote)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BeforeProposedBlockAdded provides a mock function with given fields: pacer, bc, proposal
func (_m *OnReceiveProposal) BeforeProposedBlockAdded(pacer api.Pacer, bc api.Blockchain, proposal api.Proposal) error {
	ret := _m.Called(pacer, bc, proposal)

	var r0 error
	if rf, ok := ret.Get(0).(func(api.Pacer, api.Blockchain, api.Proposal) error); ok {
		r0 = rf(pacer, bc, proposal)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BeforeVoted provides a mock function with given fields: pacer, bc, vote
func (_m *OnReceiveProposal) BeforeVoted(pacer api.Pacer, bc api.Blockchain, vote api.Vote) error {
	ret := _m.Called(pacer, bc, vote)

	var r0 error
	if rf, ok := ret.Get(0).(func(api.Pacer, api.Blockchain, api.Vote) error); ok {
		r0 = rf(pacer, bc, vote)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}