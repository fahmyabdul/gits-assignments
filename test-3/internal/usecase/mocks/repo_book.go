// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/fahmyabdul/gits-assignments/test-3/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// RepoBook is an autogenerated mock type for the RepoBook type
type RepoBook struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, request, author_id
func (_m *RepoBook) Create(ctx context.Context, request *entity.Book, author_id int) error {
	ret := _m.Called(ctx, request, author_id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Book, int) error); ok {
		r0 = rf(ctx, request, author_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *RepoBook) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchAll provides a mock function with given fields: ctx
func (_m *RepoBook) FetchAll(ctx context.Context) ([]*entity.Book, error) {
	ret := _m.Called(ctx)

	var r0 []*entity.Book
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.Book); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchByAuthorId provides a mock function with given fields: ctx, author_id
func (_m *RepoBook) FetchByAuthorId(ctx context.Context, author_id int) (*entity.BookByAuthor, error) {
	ret := _m.Called(ctx, author_id)

	var r0 *entity.BookByAuthor
	if rf, ok := ret.Get(0).(func(context.Context, int) *entity.BookByAuthor); ok {
		r0 = rf(ctx, author_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.BookByAuthor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, author_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchById provides a mock function with given fields: ctx, id
func (_m *RepoBook) FetchById(ctx context.Context, id int) (*entity.Book, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Book
	if rf, ok := ret.Get(0).(func(context.Context, int) *entity.Book); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchByName provides a mock function with given fields: ctx, name
func (_m *RepoBook) FetchByName(ctx context.Context, name string) ([]*entity.Book, error) {
	ret := _m.Called(ctx, name)

	var r0 []*entity.Book
	if rf, ok := ret.Get(0).(func(context.Context, string) []*entity.Book); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, author_id, request
func (_m *RepoBook) Update(ctx context.Context, id int, author_id int, request *entity.Book) error {
	ret := _m.Called(ctx, id, author_id, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, *entity.Book) error); ok {
		r0 = rf(ctx, id, author_id, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepoBook interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepoBook creates a new instance of RepoBook. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepoBook(t mockConstructorTestingTNewRepoBook) *RepoBook {
	mock := &RepoBook{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
