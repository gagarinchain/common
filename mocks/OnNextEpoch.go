// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// OnNextEpoch is an autogenerated mock type for the OnNextEpoch type
type OnNextEpoch struct {
	mock.Mock
}

// OnNewEpoch provides a mock function with given fields: ctx, newEpoch
func (_m *OnNextEpoch) OnNewEpoch(ctx context.Context, newEpoch int32) error {
	ret := _m.Called(ctx, newEpoch)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, newEpoch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
