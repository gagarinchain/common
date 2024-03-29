// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	context "context"

	pb "github.com/gagarinchain/common/protobuff"
	mock "github.com/stretchr/testify/mock"
)

// OnReceiveProposalServer is an autogenerated mock type for the OnReceiveProposalServer type
type OnReceiveProposalServer struct {
	mock.Mock
}

// AfterProposedBlockAdded provides a mock function with given fields: _a0, _a1
func (_m *OnReceiveProposalServer) AfterProposedBlockAdded(_a0 context.Context, _a1 *pb.AfterProposedBlockAddedRequest) (*pb.AfterProposedBlockAddedResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.AfterProposedBlockAddedResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.AfterProposedBlockAddedRequest) *pb.AfterProposedBlockAddedResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.AfterProposedBlockAddedResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.AfterProposedBlockAddedRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AfterVoted provides a mock function with given fields: _a0, _a1
func (_m *OnReceiveProposalServer) AfterVoted(_a0 context.Context, _a1 *pb.AfterVotedRequest) (*pb.AfterVotedResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.AfterVotedResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.AfterVotedRequest) *pb.AfterVotedResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.AfterVotedResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.AfterVotedRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeforeProposedBlockAdded provides a mock function with given fields: _a0, _a1
func (_m *OnReceiveProposalServer) BeforeProposedBlockAdded(_a0 context.Context, _a1 *pb.BeforeProposedBlockAddedRequest) (*pb.BeforeProposedBlockAddedResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.BeforeProposedBlockAddedResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.BeforeProposedBlockAddedRequest) *pb.BeforeProposedBlockAddedResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.BeforeProposedBlockAddedResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.BeforeProposedBlockAddedRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeforeVoted provides a mock function with given fields: _a0, _a1
func (_m *OnReceiveProposalServer) BeforeVoted(_a0 context.Context, _a1 *pb.BeforeVotedRequest) (*pb.BeforeVotedResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.BeforeVotedResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.BeforeVotedRequest) *pb.BeforeVotedResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.BeforeVotedResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.BeforeVotedRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
