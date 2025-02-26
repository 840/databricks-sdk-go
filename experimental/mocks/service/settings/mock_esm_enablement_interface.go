// Code generated by mockery v2.39.1. DO NOT EDIT.

package settings

import (
	context "context"

	settings "github.com/databricks/databricks-sdk-go/service/settings"
	mock "github.com/stretchr/testify/mock"
)

// MockEsmEnablementInterface is an autogenerated mock type for the EsmEnablementInterface type
type MockEsmEnablementInterface struct {
	mock.Mock
}

type MockEsmEnablementInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEsmEnablementInterface) EXPECT() *MockEsmEnablementInterface_Expecter {
	return &MockEsmEnablementInterface_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockEsmEnablementInterface) Get(ctx context.Context, request settings.GetEsmEnablementRequest) (*settings.EsmEnablementSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *settings.EsmEnablementSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetEsmEnablementRequest) (*settings.EsmEnablementSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetEsmEnablementRequest) *settings.EsmEnablementSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.EsmEnablementSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.GetEsmEnablementRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEsmEnablementInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockEsmEnablementInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.GetEsmEnablementRequest
func (_e *MockEsmEnablementInterface_Expecter) Get(ctx interface{}, request interface{}) *MockEsmEnablementInterface_Get_Call {
	return &MockEsmEnablementInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockEsmEnablementInterface_Get_Call) Run(run func(ctx context.Context, request settings.GetEsmEnablementRequest)) *MockEsmEnablementInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.GetEsmEnablementRequest))
	})
	return _c
}

func (_c *MockEsmEnablementInterface_Get_Call) Return(_a0 *settings.EsmEnablementSetting, _a1 error) *MockEsmEnablementInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEsmEnablementInterface_Get_Call) RunAndReturn(run func(context.Context, settings.GetEsmEnablementRequest) (*settings.EsmEnablementSetting, error)) *MockEsmEnablementInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockEsmEnablementInterface) Impl() settings.EsmEnablementService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 settings.EsmEnablementService
	if rf, ok := ret.Get(0).(func() settings.EsmEnablementService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.EsmEnablementService)
		}
	}

	return r0
}

// MockEsmEnablementInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockEsmEnablementInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockEsmEnablementInterface_Expecter) Impl() *MockEsmEnablementInterface_Impl_Call {
	return &MockEsmEnablementInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockEsmEnablementInterface_Impl_Call) Run(run func()) *MockEsmEnablementInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEsmEnablementInterface_Impl_Call) Return(_a0 settings.EsmEnablementService) *MockEsmEnablementInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEsmEnablementInterface_Impl_Call) RunAndReturn(run func() settings.EsmEnablementService) *MockEsmEnablementInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockEsmEnablementInterface) Update(ctx context.Context, request settings.UpdateEsmEnablementSettingRequest) (*settings.EsmEnablementSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *settings.EsmEnablementSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateEsmEnablementSettingRequest) (*settings.EsmEnablementSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateEsmEnablementSettingRequest) *settings.EsmEnablementSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.EsmEnablementSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.UpdateEsmEnablementSettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEsmEnablementInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockEsmEnablementInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.UpdateEsmEnablementSettingRequest
func (_e *MockEsmEnablementInterface_Expecter) Update(ctx interface{}, request interface{}) *MockEsmEnablementInterface_Update_Call {
	return &MockEsmEnablementInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockEsmEnablementInterface_Update_Call) Run(run func(ctx context.Context, request settings.UpdateEsmEnablementSettingRequest)) *MockEsmEnablementInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.UpdateEsmEnablementSettingRequest))
	})
	return _c
}

func (_c *MockEsmEnablementInterface_Update_Call) Return(_a0 *settings.EsmEnablementSetting, _a1 error) *MockEsmEnablementInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEsmEnablementInterface_Update_Call) RunAndReturn(run func(context.Context, settings.UpdateEsmEnablementSettingRequest) (*settings.EsmEnablementSetting, error)) *MockEsmEnablementInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockEsmEnablementInterface) WithImpl(impl settings.EsmEnablementService) settings.EsmEnablementInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 settings.EsmEnablementInterface
	if rf, ok := ret.Get(0).(func(settings.EsmEnablementService) settings.EsmEnablementInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.EsmEnablementInterface)
		}
	}

	return r0
}

// MockEsmEnablementInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockEsmEnablementInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl settings.EsmEnablementService
func (_e *MockEsmEnablementInterface_Expecter) WithImpl(impl interface{}) *MockEsmEnablementInterface_WithImpl_Call {
	return &MockEsmEnablementInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockEsmEnablementInterface_WithImpl_Call) Run(run func(impl settings.EsmEnablementService)) *MockEsmEnablementInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(settings.EsmEnablementService))
	})
	return _c
}

func (_c *MockEsmEnablementInterface_WithImpl_Call) Return(_a0 settings.EsmEnablementInterface) *MockEsmEnablementInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEsmEnablementInterface_WithImpl_Call) RunAndReturn(run func(settings.EsmEnablementService) settings.EsmEnablementInterface) *MockEsmEnablementInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEsmEnablementInterface creates a new instance of MockEsmEnablementInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEsmEnablementInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEsmEnablementInterface {
	mock := &MockEsmEnablementInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
