// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/gagarinchain/common/eth/common"
	mock "github.com/stretchr/testify/mock"

	pb "github.com/gagarinchain/common/protobuff"
)

// Receipt is an autogenerated mock type for the Receipt type
type Receipt struct {
	mock.Mock
}

// From provides a mock function with given fields:
func (_m *Receipt) From() common.Address {
	ret := _m.Called()

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// FromValue provides a mock function with given fields:
func (_m *Receipt) FromValue() *big.Int {
	ret := _m.Called()

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// To provides a mock function with given fields:
func (_m *Receipt) To() common.Address {
	ret := _m.Called()

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// ToStorageProto provides a mock function with given fields:
func (_m *Receipt) ToStorageProto() *pb.Receipt {
	ret := _m.Called()

	var r0 *pb.Receipt
	if rf, ok := ret.Get(0).(func() *pb.Receipt); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Receipt)
		}
	}

	return r0
}

// ToValue provides a mock function with given fields:
func (_m *Receipt) ToValue() *big.Int {
	ret := _m.Called()

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// TxHash provides a mock function with given fields:
func (_m *Receipt) TxHash() common.Hash {
	ret := _m.Called()

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func() common.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	return r0
}

// TxIndex provides a mock function with given fields:
func (_m *Receipt) TxIndex() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// Value provides a mock function with given fields:
func (_m *Receipt) Value() *big.Int {
	ret := _m.Called()

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}
