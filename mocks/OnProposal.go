// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	api "github.com/gagarinchain/common/api"
	mock "github.com/stretchr/testify/mock"
)

// OnProposal is an autogenerated mock type for the OnProposal type
type OnProposal struct {
	mock.Mock
}

// OnProposal provides a mock function with given fields: pacer, bc, proposal
func (_m *OnProposal) OnProposal(pacer api.Pacer, bc api.Blockchain, proposal api.Proposal) error {
	ret := _m.Called(pacer, bc, proposal)

	var r0 error
	if rf, ok := ret.Get(0).(func(api.Pacer, api.Blockchain, api.Proposal) error); ok {
		r0 = rf(pacer, bc, proposal)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
