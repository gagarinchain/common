// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	api "github.com/gagarinchain/common/api"
	common "github.com/gagarinchain/common/eth/common"

	crypto "github.com/gagarinchain/common/eth/crypto"

	mock "github.com/stretchr/testify/mock"

	treemap "github.com/emirpasic/gods/maps/treemap"
)

// Blockchain is an autogenerated mock type for the Blockchain type
type Blockchain struct {
	mock.Mock
}

// AddBlock provides a mock function with given fields: block
func (_m *Blockchain) AddBlock(block api.Block) ([]api.Receipt, error) {
	ret := _m.Called(block)

	var r0 []api.Receipt
	if rf, ok := ret.Get(0).(func(api.Block) []api.Receipt); ok {
		r0 = rf(block)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(api.Block) error); ok {
		r1 = rf(block)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Contains provides a mock function with given fields: hash
func (_m *Blockchain) Contains(hash common.Hash) bool {
	ret := _m.Called(hash)

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.Hash) bool); ok {
		r0 = rf(hash)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetBlockByHash provides a mock function with given fields: hash
func (_m *Blockchain) GetBlockByHash(hash common.Hash) api.Block {
	ret := _m.Called(hash)

	var r0 api.Block
	if rf, ok := ret.Get(0).(func(common.Hash) api.Block); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.Block)
		}
	}

	return r0
}

// GetBlockByHeight provides a mock function with given fields: height
func (_m *Blockchain) GetBlockByHeight(height int32) []api.Block {
	ret := _m.Called(height)

	var r0 []api.Block
	if rf, ok := ret.Get(0).(func(int32) []api.Block); ok {
		r0 = rf(height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.Block)
		}
	}

	return r0
}

// GetFork provides a mock function with given fields: height, headHash
func (_m *Blockchain) GetFork(height int32, headHash common.Hash) []api.Block {
	ret := _m.Called(height, headHash)

	var r0 []api.Block
	if rf, ok := ret.Get(0).(func(int32, common.Hash) []api.Block); ok {
		r0 = rf(height, headHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.Block)
		}
	}

	return r0
}

// GetGenesisBlock provides a mock function with given fields:
func (_m *Blockchain) GetGenesisBlock() api.Block {
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

// GetGenesisBlockSignedHash provides a mock function with given fields: key
func (_m *Blockchain) GetGenesisBlockSignedHash(key *crypto.PrivateKey) *crypto.Signature {
	ret := _m.Called(key)

	var r0 *crypto.Signature
	if rf, ok := ret.Get(0).(func(*crypto.PrivateKey) *crypto.Signature); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*crypto.Signature)
		}
	}

	return r0
}

// GetGenesisCert provides a mock function with given fields:
func (_m *Blockchain) GetGenesisCert() api.QuorumCertificate {
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

// GetHead provides a mock function with given fields:
func (_m *Blockchain) GetHead() api.Block {
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

// GetHeadRecord provides a mock function with given fields:
func (_m *Blockchain) GetHeadRecord() api.Record {
	ret := _m.Called()

	var r0 api.Record
	if rf, ok := ret.Get(0).(func() api.Record); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.Record)
		}
	}

	return r0
}

// GetThreeChain provides a mock function with given fields: twoHash
func (_m *Blockchain) GetThreeChain(twoHash common.Hash) (api.Block, api.Block, api.Block) {
	ret := _m.Called(twoHash)

	var r0 api.Block
	if rf, ok := ret.Get(0).(func(common.Hash) api.Block); ok {
		r0 = rf(twoHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.Block)
		}
	}

	var r1 api.Block
	if rf, ok := ret.Get(1).(func(common.Hash) api.Block); ok {
		r1 = rf(twoHash)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(api.Block)
		}
	}

	var r2 api.Block
	if rf, ok := ret.Get(2).(func(common.Hash) api.Block); ok {
		r2 = rf(twoHash)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(api.Block)
		}
	}

	return r0, r1, r2
}

// GetTopCommittedBlock provides a mock function with given fields:
func (_m *Blockchain) GetTopCommittedBlock() api.Block {
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

// GetTopHeight provides a mock function with given fields:
func (_m *Blockchain) GetTopHeight() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// GetTopHeightBlocks provides a mock function with given fields:
func (_m *Blockchain) GetTopHeightBlocks() []api.Block {
	ret := _m.Called()

	var r0 []api.Block
	if rf, ok := ret.Get(0).(func() []api.Block); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.Block)
		}
	}

	return r0
}

// IsSibling provides a mock function with given fields: sibling, ancestor
func (_m *Blockchain) IsSibling(sibling api.Header, ancestor api.Header) bool {
	ret := _m.Called(sibling, ancestor)

	var r0 bool
	if rf, ok := ret.Get(0).(func(api.Header, api.Header) bool); ok {
		r0 = rf(sibling, ancestor)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewBlock provides a mock function with given fields: parent, qc, data
func (_m *Blockchain) NewBlock(parent api.Block, qc api.QuorumCertificate, data []byte) (api.Block, error) {
	ret := _m.Called(parent, qc, data)

	var r0 api.Block
	if rf, ok := ret.Get(0).(func(api.Block, api.QuorumCertificate, []byte) api.Block); ok {
		r0 = rf(parent, qc, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(api.Block, api.QuorumCertificate, []byte) error); ok {
		r1 = rf(parent, qc, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OnCommit provides a mock function with given fields: b
func (_m *Blockchain) OnCommit(b api.Block) ([]api.Block, *treemap.Map, error) {
	ret := _m.Called(b)

	var r0 []api.Block
	if rf, ok := ret.Get(0).(func(api.Block) []api.Block); ok {
		r0 = rf(b)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.Block)
		}
	}

	var r1 *treemap.Map
	if rf, ok := ret.Get(1).(func(api.Block) *treemap.Map); ok {
		r1 = rf(b)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*treemap.Map)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(api.Block) error); ok {
		r2 = rf(b)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PadEmptyBlock provides a mock function with given fields: head, qc
func (_m *Blockchain) PadEmptyBlock(head api.Block, qc api.QuorumCertificate) (api.Block, error) {
	ret := _m.Called(head, qc)

	var r0 api.Block
	if rf, ok := ret.Get(0).(func(api.Block, api.QuorumCertificate) api.Block); ok {
		r0 = rf(head, qc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(api.Block, api.QuorumCertificate) error); ok {
		r1 = rf(head, qc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveBlock provides a mock function with given fields: block
func (_m *Blockchain) RemoveBlock(block api.Block) error {
	ret := _m.Called(block)

	var r0 error
	if rf, ok := ret.Get(0).(func(api.Block) error); ok {
		r0 = rf(block)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetProposerGetter provides a mock function with given fields: proposerGetter
func (_m *Blockchain) SetProposerGetter(proposerGetter api.ProposerForHeight) {
	_m.Called(proposerGetter)
}

// UpdateGenesisBlockQC provides a mock function with given fields: certificate
func (_m *Blockchain) UpdateGenesisBlockQC(certificate api.QuorumCertificate) {
	_m.Called(certificate)
}

// ValidateGenesisBlockSignature provides a mock function with given fields: signature, address
func (_m *Blockchain) ValidateGenesisBlockSignature(signature *crypto.Signature, address common.Address) bool {
	ret := _m.Called(signature, address)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*crypto.Signature, common.Address) bool); ok {
		r0 = rf(signature, address)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
