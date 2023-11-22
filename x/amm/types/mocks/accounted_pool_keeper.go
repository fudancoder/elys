// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	math "cosmossdk.io/math"
	types "github.com/cosmos/cosmos-sdk/types"
	mock "github.com/stretchr/testify/mock"
)

// AccountedPoolKeeper is an autogenerated mock type for the AccountedPoolKeeper type
type AccountedPoolKeeper struct {
	mock.Mock
}

type AccountedPoolKeeper_Expecter struct {
	mock *mock.Mock
}

func (_m *AccountedPoolKeeper) EXPECT() *AccountedPoolKeeper_Expecter {
	return &AccountedPoolKeeper_Expecter{mock: &_m.Mock}
}

// GetAccountedBalance provides a mock function with given fields: _a0, _a1, _a2
func (_m *AccountedPoolKeeper) GetAccountedBalance(_a0 types.Context, _a1 uint64, _a2 string) math.Int {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 math.Int
	if rf, ok := ret.Get(0).(func(types.Context, uint64, string) math.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	return r0
}

// AccountedPoolKeeper_GetAccountedBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccountedBalance'
type AccountedPoolKeeper_GetAccountedBalance_Call struct {
	*mock.Call
}

// GetAccountedBalance is a helper method to define mock.On call
//   - _a0 types.Context
//   - _a1 uint64
//   - _a2 string
func (_e *AccountedPoolKeeper_Expecter) GetAccountedBalance(_a0 interface{}, _a1 interface{}, _a2 interface{}) *AccountedPoolKeeper_GetAccountedBalance_Call {
	return &AccountedPoolKeeper_GetAccountedBalance_Call{Call: _e.mock.On("GetAccountedBalance", _a0, _a1, _a2)}
}

func (_c *AccountedPoolKeeper_GetAccountedBalance_Call) Run(run func(_a0 types.Context, _a1 uint64, _a2 string)) *AccountedPoolKeeper_GetAccountedBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(uint64), args[2].(string))
	})
	return _c
}

func (_c *AccountedPoolKeeper_GetAccountedBalance_Call) Return(_a0 math.Int) *AccountedPoolKeeper_GetAccountedBalance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AccountedPoolKeeper_GetAccountedBalance_Call) RunAndReturn(run func(types.Context, uint64, string) math.Int) *AccountedPoolKeeper_GetAccountedBalance_Call {
	_c.Call.Return(run)
	return _c
}

// NewAccountedPoolKeeper creates a new instance of AccountedPoolKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountedPoolKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountedPoolKeeper {
	mock := &AccountedPoolKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}