// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/Ferza17/soccer-api/model/graph/model"
)

// QueryResolver is an autogenerated mock type for the QueryResolver type
type QueryResolver struct {
	mock.Mock
}

// FindPlayer provides a mock function with given fields: ctx, input
func (_m *QueryResolver) FindPlayer(ctx context.Context, input model.FindPlayerInputArgs) (*model.Player, error) {
	ret := _m.Called(ctx, input)

	var r0 *model.Player
	if rf, ok := ret.Get(0).(func(context.Context, model.FindPlayerInputArgs) *model.Player); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Player)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.FindPlayerInputArgs) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPlayers provides a mock function with given fields: ctx, input
func (_m *QueryResolver) FindPlayers(ctx context.Context, input model.FindPlayersInputArgs) ([]*model.Player, error) {
	ret := _m.Called(ctx, input)

	var r0 []*model.Player
	if rf, ok := ret.Get(0).(func(context.Context, model.FindPlayersInputArgs) []*model.Player); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Player)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.FindPlayersInputArgs) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTeam provides a mock function with given fields: ctx, input
func (_m *QueryResolver) FindTeam(ctx context.Context, input model.FindTeamInputArgs) (*model.Team, error) {
	ret := _m.Called(ctx, input)

	var r0 *model.Team
	if rf, ok := ret.Get(0).(func(context.Context, model.FindTeamInputArgs) *model.Team); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.FindTeamInputArgs) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTeams provides a mock function with given fields: ctx, input
func (_m *QueryResolver) FindTeams(ctx context.Context, input model.FindTeamsInputArgs) ([]*model.Team, error) {
	ret := _m.Called(ctx, input)

	var r0 []*model.Team
	if rf, ok := ret.Get(0).(func(context.Context, model.FindTeamsInputArgs) []*model.Team); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.FindTeamsInputArgs) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewQueryResolver interface {
	mock.TestingT
	Cleanup(func())
}

// NewQueryResolver creates a new instance of QueryResolver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewQueryResolver(t mockConstructorTestingTNewQueryResolver) *QueryResolver {
	mock := &QueryResolver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
