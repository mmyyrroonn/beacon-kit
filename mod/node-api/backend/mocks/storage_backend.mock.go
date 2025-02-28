// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// StorageBackend is an autogenerated mock type for the StorageBackend type
type StorageBackend[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}] struct {
	mock.Mock
}

type StorageBackend_Expecter[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}] struct {
	mock *mock.Mock
}

func (_m *StorageBackend[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) EXPECT() *StorageBackend_Expecter[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	return &StorageBackend_Expecter[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]{mock: &_m.Mock}
}

// AvailabilityStore provides a mock function with given fields:
func (_m *StorageBackend[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) AvailabilityStore() AvailabilityStoreT {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AvailabilityStore")
	}

	var r0 AvailabilityStoreT
	if rf, ok := ret.Get(0).(func() AvailabilityStoreT); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(AvailabilityStoreT)
	}

	return r0
}

// StorageBackend_AvailabilityStore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AvailabilityStore'
type StorageBackend_AvailabilityStore_Call[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}] struct {
	*mock.Call
}

// AvailabilityStore is a helper method to define mock.On call
func (_e *StorageBackend_Expecter[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) AvailabilityStore() *StorageBackend_AvailabilityStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	return &StorageBackend_AvailabilityStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]{Call: _e.mock.On("AvailabilityStore")}
}

func (_c *StorageBackend_AvailabilityStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) Run(run func()) *StorageBackend_AvailabilityStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *StorageBackend_AvailabilityStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) Return(_a0 AvailabilityStoreT) *StorageBackend_AvailabilityStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageBackend_AvailabilityStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) RunAndReturn(run func() AvailabilityStoreT) *StorageBackend_AvailabilityStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	_c.Call.Return(run)
	return _c
}

// BlockStore provides a mock function with given fields:
func (_m *StorageBackend[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) BlockStore() BlockStoreT {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BlockStore")
	}

	var r0 BlockStoreT
	if rf, ok := ret.Get(0).(func() BlockStoreT); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(BlockStoreT)
	}

	return r0
}

// StorageBackend_BlockStore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockStore'
type StorageBackend_BlockStore_Call[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}] struct {
	*mock.Call
}

// BlockStore is a helper method to define mock.On call
func (_e *StorageBackend_Expecter[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) BlockStore() *StorageBackend_BlockStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	return &StorageBackend_BlockStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]{Call: _e.mock.On("BlockStore")}
}

func (_c *StorageBackend_BlockStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) Run(run func()) *StorageBackend_BlockStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *StorageBackend_BlockStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) Return(_a0 BlockStoreT) *StorageBackend_BlockStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageBackend_BlockStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) RunAndReturn(run func() BlockStoreT) *StorageBackend_BlockStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	_c.Call.Return(run)
	return _c
}

// DepositStore provides a mock function with given fields:
func (_m *StorageBackend[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) DepositStore() DepositStoreT {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DepositStore")
	}

	var r0 DepositStoreT
	if rf, ok := ret.Get(0).(func() DepositStoreT); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(DepositStoreT)
	}

	return r0
}

// StorageBackend_DepositStore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DepositStore'
type StorageBackend_DepositStore_Call[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}] struct {
	*mock.Call
}

// DepositStore is a helper method to define mock.On call
func (_e *StorageBackend_Expecter[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) DepositStore() *StorageBackend_DepositStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	return &StorageBackend_DepositStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]{Call: _e.mock.On("DepositStore")}
}

func (_c *StorageBackend_DepositStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) Run(run func()) *StorageBackend_DepositStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *StorageBackend_DepositStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) Return(_a0 DepositStoreT) *StorageBackend_DepositStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageBackend_DepositStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) RunAndReturn(run func() DepositStoreT) *StorageBackend_DepositStore_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	_c.Call.Return(run)
	return _c
}

// StateFromContext provides a mock function with given fields: _a0
func (_m *StorageBackend[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) StateFromContext(_a0 context.Context) BeaconStateT {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for StateFromContext")
	}

	var r0 BeaconStateT
	if rf, ok := ret.Get(0).(func(context.Context) BeaconStateT); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(BeaconStateT)
	}

	return r0
}

// StorageBackend_StateFromContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StateFromContext'
type StorageBackend_StateFromContext_Call[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}] struct {
	*mock.Call
}

// StateFromContext is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *StorageBackend_Expecter[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) StateFromContext(_a0 interface{}) *StorageBackend_StateFromContext_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	return &StorageBackend_StateFromContext_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]{Call: _e.mock.On("StateFromContext", _a0)}
}

func (_c *StorageBackend_StateFromContext_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) Run(run func(_a0 context.Context)) *StorageBackend_StateFromContext_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *StorageBackend_StateFromContext_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) Return(_a0 BeaconStateT) *StorageBackend_StateFromContext_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StorageBackend_StateFromContext_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]) RunAndReturn(run func(context.Context) BeaconStateT) *StorageBackend_StateFromContext_Call[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	_c.Call.Return(run)
	return _c
}

// NewStorageBackend creates a new instance of StorageBackend. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorageBackend[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *StorageBackend[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT] {
	mock := &StorageBackend[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
