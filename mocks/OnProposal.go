// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/gagarinchain/common/api"

	mock "github.com/stretchr/testify/mock"
)

// OnProposal is an autogenerated mock type for the OnProposal type
type OnProposal struct {
	mock.Mock
}

// OnProposal provides a mock function with given fields: ctx, proposal
func (_m *OnProposal) OnProposal(ctx context.Context, proposal api.Proposal) error {
	ret := _m.Called(ctx, proposal)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, api.Proposal) error); ok {
		r0 = rf(ctx, proposal)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
