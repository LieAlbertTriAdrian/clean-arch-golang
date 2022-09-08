// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/LieAlbertTriAdrian/clean-arch-golang/domain"
	mock "github.com/stretchr/testify/mock"
)

// ITodoRepository is an autogenerated mock type for the ITodoRepository type
type ITodoRepository struct {
	mock.Mock
}

// AddTodo provides a mock function with given fields: ctx, todo
func (_m *ITodoRepository) AddTodo(ctx context.Context, todo *domain.Todo) error {
	ret := _m.Called(ctx, todo)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Todo) error); ok {
		r0 = rf(ctx, todo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, param
func (_m *ITodoRepository) Fetch(ctx context.Context, param domain.FetchTodoParam) ([]domain.Todo, string, error) {
	ret := _m.Called(ctx, param)

	var r0 []domain.Todo
	if rf, ok := ret.Get(0).(func(context.Context, domain.FetchTodoParam) []domain.Todo); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Todo)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, domain.FetchTodoParam) string); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, domain.FetchTodoParam) error); ok {
		r2 = rf(ctx, param)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewITodoRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewITodoRepository creates a new instance of ITodoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewITodoRepository(t mockConstructorTestingTNewITodoRepository) *ITodoRepository {
	mock := &ITodoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
