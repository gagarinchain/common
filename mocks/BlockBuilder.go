// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	api "github.com/gagarinchain/common/api"
	mock "github.com/stretchr/testify/mock"
)

// BlockBuilder is an autogenerated mock type for the BlockBuilder type
type BlockBuilder struct {
	mock.Mock
}

// AddTx provides a mock function with given fields: tx
func (_m *BlockBuilder) AddTx(tx api.Transaction) api.BlockBuilder {
	ret := _m.Called(tx)

	var r0 api.BlockBuilder
	if rf, ok := ret.Get(0).(func(api.Transaction) api.BlockBuilder); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.BlockBuilder)
		}
	}

	return r0
}

// Build provides a mock function with given fields:
func (_m *BlockBuilder) Build() api.Block {
	ret := _m.Called()

	var r0 api.Block
	if rf, ok := ret.Get(0).(func() api.Block); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.Block)
		}
	}

	return r0
}

// Data provides a mock function with given fields:
func (_m *BlockBuilder) Data() []byte {
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

// Header provides a mock function with given fields:
func (_m *BlockBuilder) Header() api.Header {
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

// QC provides a mock function with given fields:
func (_m *BlockBuilder) QC() api.QuorumCertificate {
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

// SetData provides a mock function with given fields: data
func (_m *BlockBuilder) SetData(data []byte) api.BlockBuilder {
	ret := _m.Called(data)

	var r0 api.BlockBuilder
	if rf, ok := ret.Get(0).(func([]byte) api.BlockBuilder); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.BlockBuilder)
		}
	}

	return r0
}

// SetHeader provides a mock function with given fields: header
func (_m *BlockBuilder) SetHeader(header api.Header) api.BlockBuilder {
	ret := _m.Called(header)

	var r0 api.BlockBuilder
	if rf, ok := ret.Get(0).(func(api.Header) api.BlockBuilder); ok {
		r0 = rf(header)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.BlockBuilder)
		}
	}

	return r0
}

// SetQC provides a mock function with given fields: qc
func (_m *BlockBuilder) SetQC(qc api.QuorumCertificate) api.BlockBuilder {
	ret := _m.Called(qc)

	var r0 api.BlockBuilder
	if rf, ok := ret.Get(0).(func(api.QuorumCertificate) api.BlockBuilder); ok {
		r0 = rf(qc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.BlockBuilder)
		}
	}

	return r0
}

// SetTxs provides a mock function with given fields: txs
func (_m *BlockBuilder) SetTxs(txs []api.Transaction) api.BlockBuilder {
	ret := _m.Called(txs)

	var r0 api.BlockBuilder
	if rf, ok := ret.Get(0).(func([]api.Transaction) api.BlockBuilder); ok {
		r0 = rf(txs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.BlockBuilder)
		}
	}

	return r0
}

// Txs provides a mock function with given fields:
func (_m *BlockBuilder) Txs() []api.Transaction {
	ret := _m.Called()

	var r0 []api.Transaction
	if rf, ok := ret.Get(0).(func() []api.Transaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.Transaction)
		}
	}

	return r0
}
