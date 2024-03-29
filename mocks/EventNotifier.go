// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	api "github.com/gagarinchain/common/api"
	mock "github.com/stretchr/testify/mock"
)

// EventNotifier is an autogenerated mock type for the EventNotifier type
type EventNotifier struct {
	mock.Mock
}

// FireEvent provides a mock function with given fields: event
func (_m *EventNotifier) FireEvent(event api.Event) {
	_m.Called(event)
}

// NotifyEvent provides a mock function with given fields: event
func (_m *EventNotifier) NotifyEvent(event api.Event) {
	_m.Called(event)
}

// SubscribeProtocolEvents provides a mock function with given fields: sub
func (_m *EventNotifier) SubscribeProtocolEvents(sub chan api.Event) {
	_m.Called(sub)
}
