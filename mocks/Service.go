// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/gagarinchain/common"

	message "github.com/gagarinchain/common/message"

	mock "github.com/stretchr/testify/mock"

	network "github.com/gagarinchain/common/network"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Bootstrap provides a mock function with given fields: ctx, cfg
func (_m *Service) Bootstrap(ctx context.Context, cfg *network.BootstrapConfig) (chan int, chan error) {
	ret := _m.Called(ctx, cfg)

	var r0 chan int
	if rf, ok := ret.Get(0).(func(context.Context, *network.BootstrapConfig) chan int); ok {
		r0 = rf(ctx, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan int)
		}
	}

	var r1 chan error
	if rf, ok := ret.Get(1).(func(context.Context, *network.BootstrapConfig) chan error); ok {
		r1 = rf(ctx, cfg)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(chan error)
		}
	}

	return r0, r1
}

// Broadcast provides a mock function with given fields: ctx, msg
func (_m *Service) Broadcast(ctx context.Context, msg *message.Message) {
	_m.Called(ctx, msg)
}

// BroadcastTransaction provides a mock function with given fields: ctx, msg
func (_m *Service) BroadcastTransaction(ctx context.Context, msg *message.Message) {
	_m.Called(ctx, msg)
}

// SendMessage provides a mock function with given fields: ctx, peer, msg
func (_m *Service) SendMessage(ctx context.Context, peer *common.Peer, msg *message.Message) {
	_m.Called(ctx, peer, msg)
}

// SendMessageTriggered provides a mock function with given fields: ctx, peer, msg, trigger
func (_m *Service) SendMessageTriggered(ctx context.Context, peer *common.Peer, msg *message.Message, trigger chan interface{}) {
	_m.Called(ctx, peer, msg, trigger)
}

// SendRequest provides a mock function with given fields: ctx, peer, msg
func (_m *Service) SendRequest(ctx context.Context, peer *common.Peer, msg *message.Message) (chan *message.Message, chan error) {
	ret := _m.Called(ctx, peer, msg)

	var r0 chan *message.Message
	if rf, ok := ret.Get(0).(func(context.Context, *common.Peer, *message.Message) chan *message.Message); ok {
		r0 = rf(ctx, peer, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *message.Message)
		}
	}

	var r1 chan error
	if rf, ok := ret.Get(1).(func(context.Context, *common.Peer, *message.Message) chan error); ok {
		r1 = rf(ctx, peer, msg)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(chan error)
		}
	}

	return r0, r1
}

// SendRequestToRandomPeer provides a mock function with given fields: ctx, req
func (_m *Service) SendRequestToRandomPeer(ctx context.Context, req *message.Message) (chan *message.Message, chan error) {
	ret := _m.Called(ctx, req)

	var r0 chan *message.Message
	if rf, ok := ret.Get(0).(func(context.Context, *message.Message) chan *message.Message); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *message.Message)
		}
	}

	var r1 chan error
	if rf, ok := ret.Get(1).(func(context.Context, *message.Message) chan error); ok {
		r1 = rf(ctx, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(chan error)
		}
	}

	return r0, r1
}

// SendResponse provides a mock function with given fields: ctx, msg
func (_m *Service) SendResponse(ctx context.Context, msg *message.Message) {
	_m.Called(ctx, msg)
}
