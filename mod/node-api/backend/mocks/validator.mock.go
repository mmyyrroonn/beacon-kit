// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	backend "github.com/berachain/beacon-kit/mod/node-api/backend"
	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"

	mock "github.com/stretchr/testify/mock"
)

// Validator is an autogenerated mock type for the Validator type
type Validator[WithdrawalCredentialsT backend.WithdrawalCredentials] struct {
	mock.Mock
}

type Validator_Expecter[WithdrawalCredentialsT backend.WithdrawalCredentials] struct {
	mock *mock.Mock
}

func (_m *Validator[WithdrawalCredentialsT]) EXPECT() *Validator_Expecter[WithdrawalCredentialsT] {
	return &Validator_Expecter[WithdrawalCredentialsT]{mock: &_m.Mock}
}

// GetWithdrawalCredentials provides a mock function with given fields:
func (_m *Validator[WithdrawalCredentialsT]) GetWithdrawalCredentials() WithdrawalCredentialsT {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetWithdrawalCredentials")
	}

	var r0 WithdrawalCredentialsT
	if rf, ok := ret.Get(0).(func() WithdrawalCredentialsT); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(WithdrawalCredentialsT)
	}

	return r0
}

// Validator_GetWithdrawalCredentials_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWithdrawalCredentials'
type Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT backend.WithdrawalCredentials] struct {
	*mock.Call
}

// GetWithdrawalCredentials is a helper method to define mock.On call
func (_e *Validator_Expecter[WithdrawalCredentialsT]) GetWithdrawalCredentials() *Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT] {
	return &Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT]{Call: _e.mock.On("GetWithdrawalCredentials")}
}

func (_c *Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT]) Run(run func()) *Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT]) Return(_a0 WithdrawalCredentialsT) *Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT]) RunAndReturn(run func() WithdrawalCredentialsT) *Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT] {
	_c.Call.Return(run)
	return _c
}

// IsFullyWithdrawable provides a mock function with given fields: amount, epoch
func (_m *Validator[WithdrawalCredentialsT]) IsFullyWithdrawable(amount math.U64, epoch math.U64) bool {
	ret := _m.Called(amount, epoch)

	if len(ret) == 0 {
		panic("no return value specified for IsFullyWithdrawable")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(math.U64, math.U64) bool); ok {
		r0 = rf(amount, epoch)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Validator_IsFullyWithdrawable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsFullyWithdrawable'
type Validator_IsFullyWithdrawable_Call[WithdrawalCredentialsT backend.WithdrawalCredentials] struct {
	*mock.Call
}

// IsFullyWithdrawable is a helper method to define mock.On call
//   - amount math.U64
//   - epoch math.U64
func (_e *Validator_Expecter[WithdrawalCredentialsT]) IsFullyWithdrawable(amount interface{}, epoch interface{}) *Validator_IsFullyWithdrawable_Call[WithdrawalCredentialsT] {
	return &Validator_IsFullyWithdrawable_Call[WithdrawalCredentialsT]{Call: _e.mock.On("IsFullyWithdrawable", amount, epoch)}
}

func (_c *Validator_IsFullyWithdrawable_Call[WithdrawalCredentialsT]) Run(run func(amount math.U64, epoch math.U64)) *Validator_IsFullyWithdrawable_Call[WithdrawalCredentialsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(math.U64), args[1].(math.U64))
	})
	return _c
}

func (_c *Validator_IsFullyWithdrawable_Call[WithdrawalCredentialsT]) Return(_a0 bool) *Validator_IsFullyWithdrawable_Call[WithdrawalCredentialsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Validator_IsFullyWithdrawable_Call[WithdrawalCredentialsT]) RunAndReturn(run func(math.U64, math.U64) bool) *Validator_IsFullyWithdrawable_Call[WithdrawalCredentialsT] {
	_c.Call.Return(run)
	return _c
}

// IsPartiallyWithdrawable provides a mock function with given fields: amount1, amount2
func (_m *Validator[WithdrawalCredentialsT]) IsPartiallyWithdrawable(amount1 math.U64, amount2 math.U64) bool {
	ret := _m.Called(amount1, amount2)

	if len(ret) == 0 {
		panic("no return value specified for IsPartiallyWithdrawable")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(math.U64, math.U64) bool); ok {
		r0 = rf(amount1, amount2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Validator_IsPartiallyWithdrawable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsPartiallyWithdrawable'
type Validator_IsPartiallyWithdrawable_Call[WithdrawalCredentialsT backend.WithdrawalCredentials] struct {
	*mock.Call
}

// IsPartiallyWithdrawable is a helper method to define mock.On call
//   - amount1 math.U64
//   - amount2 math.U64
func (_e *Validator_Expecter[WithdrawalCredentialsT]) IsPartiallyWithdrawable(amount1 interface{}, amount2 interface{}) *Validator_IsPartiallyWithdrawable_Call[WithdrawalCredentialsT] {
	return &Validator_IsPartiallyWithdrawable_Call[WithdrawalCredentialsT]{Call: _e.mock.On("IsPartiallyWithdrawable", amount1, amount2)}
}

func (_c *Validator_IsPartiallyWithdrawable_Call[WithdrawalCredentialsT]) Run(run func(amount1 math.U64, amount2 math.U64)) *Validator_IsPartiallyWithdrawable_Call[WithdrawalCredentialsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(math.U64), args[1].(math.U64))
	})
	return _c
}

func (_c *Validator_IsPartiallyWithdrawable_Call[WithdrawalCredentialsT]) Return(_a0 bool) *Validator_IsPartiallyWithdrawable_Call[WithdrawalCredentialsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Validator_IsPartiallyWithdrawable_Call[WithdrawalCredentialsT]) RunAndReturn(run func(math.U64, math.U64) bool) *Validator_IsPartiallyWithdrawable_Call[WithdrawalCredentialsT] {
	_c.Call.Return(run)
	return _c
}

// NewValidator creates a new instance of Validator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewValidator[WithdrawalCredentialsT backend.WithdrawalCredentials](t interface {
	mock.TestingT
	Cleanup(func())
}) *Validator[WithdrawalCredentialsT] {
	mock := &Validator[WithdrawalCredentialsT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
