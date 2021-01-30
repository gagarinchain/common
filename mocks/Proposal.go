// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	common "github.com/gagarinchain/common"
	api "github.com/gagarinchain/common/api"

	crypto "github.com/gagarinchain/common/eth/crypto"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/gagarinchain/common/protobuff"
)

// Proposal is an autogenerated mock type for the Proposal type
type Proposal struct {
	mock.Mock
}

// GetMessage provides a mock function with given fields:
func (_m *Proposal) GetMessage() *pb.ProposalPayload {
	ret := _m.Called()

	var r0 *pb.ProposalPayload
	if rf, ok := ret.Get(0).(func() *pb.ProposalPayload); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.ProposalPayload)
		}
	}

	return r0
}

// NewBlock provides a mock function with given fields:
func (_m *Proposal) NewBlock() api.Block {
	ret := _m.Called()

	var r0 api.Block
	if rf, ok := ret.Get(0).(func() api.Block); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.Block)
		}
	}

	return r0
}

// Sender provides a mock function with given fields:
func (_m *Proposal) Sender() *common.Peer {
	ret := _m.Called()

	var r0 *common.Peer
	if rf, ok := ret.Get(0).(func() *common.Peer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Peer)
		}
	}

	return r0
}

// Sign provides a mock function with given fields: key
func (_m *Proposal) Sign(key *crypto.PrivateKey) {
	_m.Called(key)
}

// Signature provides a mock function with given fields:
func (_m *Proposal) Signature() *crypto.Signature {
	ret := _m.Called()

	var r0 *crypto.Signature
	if rf, ok := ret.Get(0).(func() *crypto.Signature); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*crypto.Signature)
		}
	}

	return r0
}
