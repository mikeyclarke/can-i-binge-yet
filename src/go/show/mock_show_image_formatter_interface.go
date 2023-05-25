// Code generated by mockery v2.28.0. DO NOT EDIT.

package show

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockShowImageFormatterInterface is an autogenerated mock type for the ShowImageFormatterInterface type
type MockShowImageFormatterInterface struct {
	mock.Mock
}

type MockShowImageFormatterInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockShowImageFormatterInterface) EXPECT() *MockShowImageFormatterInterface_Expecter {
	return &MockShowImageFormatterInterface_Expecter{mock: &_m.Mock}
}

// Format provides a mock function with given fields: ctx, imageType, imagePath, defaultSize
func (_m *MockShowImageFormatterInterface) Format(ctx context.Context, imageType string, imagePath string, defaultSize string) ShowImageResult {
	ret := _m.Called(ctx, imageType, imagePath, defaultSize)

	var r0 ShowImageResult
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) ShowImageResult); ok {
		r0 = rf(ctx, imageType, imagePath, defaultSize)
	} else {
		r0 = ret.Get(0).(ShowImageResult)
	}

	return r0
}

// MockShowImageFormatterInterface_Format_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Format'
type MockShowImageFormatterInterface_Format_Call struct {
	*mock.Call
}

// Format is a helper method to define mock.On call
//   - ctx context.Context
//   - imageType string
//   - imagePath string
//   - defaultSize string
func (_e *MockShowImageFormatterInterface_Expecter) Format(ctx interface{}, imageType interface{}, imagePath interface{}, defaultSize interface{}) *MockShowImageFormatterInterface_Format_Call {
	return &MockShowImageFormatterInterface_Format_Call{Call: _e.mock.On("Format", ctx, imageType, imagePath, defaultSize)}
}

func (_c *MockShowImageFormatterInterface_Format_Call) Run(run func(ctx context.Context, imageType string, imagePath string, defaultSize string)) *MockShowImageFormatterInterface_Format_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockShowImageFormatterInterface_Format_Call) Return(_a0 ShowImageResult) *MockShowImageFormatterInterface_Format_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockShowImageFormatterInterface_Format_Call) RunAndReturn(run func(context.Context, string, string, string) ShowImageResult) *MockShowImageFormatterInterface_Format_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockShowImageFormatterInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockShowImageFormatterInterface creates a new instance of MockShowImageFormatterInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockShowImageFormatterInterface(t mockConstructorTestingTNewMockShowImageFormatterInterface) *MockShowImageFormatterInterface {
	mock := &MockShowImageFormatterInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
