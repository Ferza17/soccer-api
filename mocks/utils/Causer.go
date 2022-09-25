// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Causer is an autogenerated mock type for the Causer type
type Causer struct {
	mock.Mock
}

// Cause provides a mock function with given fields:
func (_m *Causer) Cause() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCauser interface {
	mock.TestingT
	Cleanup(func())
}

// NewCauser creates a new instance of Causer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCauser(t mockConstructorTestingTNewCauser) *Causer {
	mock := &Causer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}