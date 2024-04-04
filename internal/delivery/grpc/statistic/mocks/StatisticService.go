// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/shamank/ai-marketplace-stats-service/internal/domain/models"
	mock "github.com/stretchr/testify/mock"
)

// StatisticService is an autogenerated mock type for the StatisticService type
type StatisticService struct {
	mock.Mock
}

// Call provides a mock function with given fields: ctx, AIServiceUID, userUID
func (_m *StatisticService) Call(ctx context.Context, AIServiceUID string, userUID string) error {
	ret := _m.Called(ctx, AIServiceUID, userUID)

	if len(ret) == 0 {
		panic("no return value specified for Call")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, AIServiceUID, userUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateService provides a mock function with given fields: ctx, create
func (_m *StatisticService) CreateService(ctx context.Context, create models.AIServiceCreate) (string, error) {
	ret := _m.Called(ctx, create)

	if len(ret) == 0 {
		panic("no return value specified for CreateService")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.AIServiceCreate) (string, error)); ok {
		return rf(ctx, create)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.AIServiceCreate) string); ok {
		r0 = rf(ctx, create)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.AIServiceCreate) error); ok {
		r1 = rf(ctx, create)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCalls provides a mock function with given fields: ctx, filter
func (_m *StatisticService) GetCalls(ctx context.Context, filter models.StatisticFilter) ([]models.StatisticRead, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetCalls")
	}

	var r0 []models.StatisticRead
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.StatisticFilter) ([]models.StatisticRead, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.StatisticFilter) []models.StatisticRead); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.StatisticRead)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.StatisticFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStatisticService creates a new instance of StatisticService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStatisticService(t interface {
	mock.TestingT
	Cleanup(func())
}) *StatisticService {
	mock := &StatisticService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
