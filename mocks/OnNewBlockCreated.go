// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	api "github.com/gagarinchain/common/api"
	mock "github.com/stretchr/testify/mock"
)

// OnNewBlockCreated is an autogenerated mock type for the OnNewBlockCreated type
type OnNewBlockCreated struct {
	mock.Mock
}

// OnBlockCreated provides a mock function with given fields: bc, builder, receipts
func (_m *OnNewBlockCreated) OnBlockCreated(bc api.Blockchain, builder api.BlockBuilder, receipts []api.Receipt) error {
	ret := _m.Called(bc, builder, receipts)

	var r0 error
	if rf, ok := ret.Get(0).(func(api.Blockchain, api.BlockBuilder, []api.Receipt) error); ok {
		r0 = rf(bc, builder, receipts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}