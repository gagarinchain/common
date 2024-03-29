// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/gagarinchain/common/api"

	mock "github.com/stretchr/testify/mock"

	treemap "github.com/emirpasic/gods/maps/treemap"
)

// OnBlockCommit is an autogenerated mock type for the OnBlockCommit type
type OnBlockCommit struct {
	mock.Mock
}

// OnBlockCommit provides a mock function with given fields: ctx, block, orphans
func (_m *OnBlockCommit) OnBlockCommit(ctx context.Context, block api.Block, orphans *treemap.Map) error {
	ret := _m.Called(ctx, block, orphans)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, api.Block, *treemap.Map) error); ok {
		r0 = rf(ctx, block, orphans)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
