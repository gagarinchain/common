// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	api "github.com/gagarinchain/common/api"
	crypto "github.com/gagarinchain/common/eth/crypto"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/gagarinchain/common/protobuff"
)

// Block is an autogenerated mock type for the Block type
type Block struct {
	mock.Mock
}

// AddTransaction provides a mock function with given fields: t
func (_m *Block) AddTransaction(t api.Transaction) {
	_m.Called(t)
}

// Data provides a mock function with given fields:
func (_m *Block) Data() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// GetMessage provides a mock function with given fields:
func (_m *Block) GetMessage() *pb.Block {
	ret := _m.Called()

	var r0 *pb.Block
	if rf, ok := ret.Get(0).(func() *pb.Block); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Block)
		}
	}

	return r0
}

// Header provides a mock function with given fields:
func (_m *Block) Header() api.Header {
	ret := _m.Called()

	var r0 api.Header
	if rf, ok := ret.Get(0).(func() api.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.Header)
		}
	}

	return r0
}

// Height provides a mock function with given fields:
func (_m *Block) Height() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// QC provides a mock function with given fields:
func (_m *Block) QC() api.QuorumCertificate {
	ret := _m.Called()

	var r0 api.QuorumCertificate
	if rf, ok := ret.Get(0).(func() api.QuorumCertificate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.QuorumCertificate)
		}
	}

	return r0
}

// QRef provides a mock function with given fields:
func (_m *Block) QRef() api.Header {
	ret := _m.Called()

	var r0 api.Header
	if rf, ok := ret.Get(0).(func() api.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.Header)
		}
	}

	return r0
}

// Receipts provides a mock function with given fields:
func (_m *Block) Receipts() []api.Receipt {
	ret := _m.Called()

	var r0 []api.Receipt
	if rf, ok := ret.Get(0).(func() []api.Receipt); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.Receipt)
		}
	}

	return r0
}

// Serialize provides a mock function with given fields:
func (_m *Block) Serialize() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetQC provides a mock function with given fields: qc
func (_m *Block) SetQC(qc api.QuorumCertificate) {
	_m.Called(qc)
}

// SetReceipts provides a mock function with given fields: receipts
func (_m *Block) SetReceipts(receipts []api.Receipt) {
	_m.Called(receipts)
}

// SetSignature provides a mock function with given fields: s
func (_m *Block) SetSignature(s *crypto.SignatureAggregate) {
	_m.Called(s)
}

// Signature provides a mock function with given fields:
func (_m *Block) Signature() *crypto.SignatureAggregate {
	ret := _m.Called()

	var r0 *crypto.SignatureAggregate
	if rf, ok := ret.Get(0).(func() *crypto.SignatureAggregate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*crypto.SignatureAggregate)
		}
	}

	return r0
}

// ToStorageProto provides a mock function with given fields:
func (_m *Block) ToStorageProto() *pb.BlockS {
	ret := _m.Called()

	var r0 *pb.BlockS
	if rf, ok := ret.Get(0).(func() *pb.BlockS); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.BlockS)
		}
	}

	return r0
}

// Txs provides a mock function with given fields:
func (_m *Block) Txs() api.Iterator {
	ret := _m.Called()

	var r0 api.Iterator
	if rf, ok := ret.Get(0).(func() api.Iterator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.Iterator)
		}
	}

	return r0
}

// TxsCount provides a mock function with given fields:
func (_m *Block) TxsCount() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
