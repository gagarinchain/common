// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	context "context"

	pb "github.com/gagarinchain/common/protobuff"
	mock "github.com/stretchr/testify/mock"
)

// OnProposalServer is an autogenerated mock type for the OnProposalServer type
type OnProposalServer struct {
	mock.Mock
}

// OnProposal provides a mock function with given fields: _a0, _a1
func (_m *OnProposalServer) OnProposal(_a0 context.Context, _a1 *pb.OnProposalRequest) (*pb.OnProposalResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.OnProposalResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.OnProposalRequest) *pb.OnProposalResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.OnProposalResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.OnProposalRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
