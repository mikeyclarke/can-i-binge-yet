// Code generated by mockery v2.28.0. DO NOT EDIT.

package asset

import mock "github.com/stretchr/testify/mock"

// MockAssetManifestInterface is an autogenerated mock type for the AssetManifestInterface type
type MockAssetManifestInterface struct {
	mock.Mock
}

type MockAssetManifestInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAssetManifestInterface) EXPECT() *MockAssetManifestInterface_Expecter {
	return &MockAssetManifestInterface_Expecter{mock: &_m.Mock}
}

// GetAssetUrl provides a mock function with given fields: filename
func (_m *MockAssetManifestInterface) GetAssetUrl(filename string) string {
	ret := _m.Called(filename)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(filename)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockAssetManifestInterface_GetAssetUrl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAssetUrl'
type MockAssetManifestInterface_GetAssetUrl_Call struct {
	*mock.Call
}

// GetAssetUrl is a helper method to define mock.On call
//   - filename string
func (_e *MockAssetManifestInterface_Expecter) GetAssetUrl(filename interface{}) *MockAssetManifestInterface_GetAssetUrl_Call {
	return &MockAssetManifestInterface_GetAssetUrl_Call{Call: _e.mock.On("GetAssetUrl", filename)}
}

func (_c *MockAssetManifestInterface_GetAssetUrl_Call) Run(run func(filename string)) *MockAssetManifestInterface_GetAssetUrl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAssetManifestInterface_GetAssetUrl_Call) Return(_a0 string) *MockAssetManifestInterface_GetAssetUrl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAssetManifestInterface_GetAssetUrl_Call) RunAndReturn(run func(string) string) *MockAssetManifestInterface_GetAssetUrl_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockAssetManifestInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAssetManifestInterface creates a new instance of MockAssetManifestInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAssetManifestInterface(t mockConstructorTestingTNewMockAssetManifestInterface) *MockAssetManifestInterface {
	mock := &MockAssetManifestInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
