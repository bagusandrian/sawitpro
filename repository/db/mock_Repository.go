// Code generated by mockery v2.23.2. DO NOT EDIT.

package db

import (
	context "context"

	model "github.com/bagusandrian/sawitpro/model"
	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// GetUserDataByID provides a mock function with given fields: ctx, userID
func (_m *MockRepository) GetUserDataByID(ctx context.Context, userID int64) (model.ResponseGetProfile, error) {
	ret := _m.Called(ctx, userID)

	var r0 model.ResponseGetProfile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (model.ResponseGetProfile, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.ResponseGetProfile); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(model.ResponseGetProfile)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, req
func (_m *MockRepository) Login(ctx context.Context, req model.RequestLogin) (model.ResponseLogin, error) {
	ret := _m.Called(ctx, req)

	var r0 model.ResponseLogin
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.RequestLogin) (model.ResponseLogin, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.RequestLogin) model.ResponseLogin); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.ResponseLogin)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.RequestLogin) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Registration provides a mock function with given fields: ctx, req
func (_m *MockRepository) Registration(ctx context.Context, req model.RequestRegistration) (model.ResponseRegristration, error) {
	ret := _m.Called(ctx, req)

	var r0 model.ResponseRegristration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.RequestRegistration) (model.ResponseRegristration, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.RequestRegistration) model.ResponseRegristration); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.ResponseRegristration)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.RequestRegistration) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProfile provides a mock function with given fields: ctx, req
func (_m *MockRepository) UpdateProfile(ctx context.Context, req model.RequestUpdateProfile) (model.ResponseUpdateProfile, error) {
	ret := _m.Called(ctx, req)

	var r0 model.ResponseUpdateProfile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.RequestUpdateProfile) (model.ResponseUpdateProfile, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.RequestUpdateProfile) model.ResponseUpdateProfile); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.ResponseUpdateProfile)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.RequestUpdateProfile) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRepository(t mockConstructorTestingTNewMockRepository) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
