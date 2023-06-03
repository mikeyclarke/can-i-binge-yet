// Code generated by mockery v2.28.0. DO NOT EDIT.

package url

import mock "github.com/stretchr/testify/mock"

// MockSlugGeneratorInterface is an autogenerated mock type for the SlugGeneratorInterface type
type MockSlugGeneratorInterface struct {
	mock.Mock
}

type MockSlugGeneratorInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSlugGeneratorInterface) EXPECT() *MockSlugGeneratorInterface_Expecter {
	return &MockSlugGeneratorInterface_Expecter{mock: &_m.Mock}
}

// Generate provides a mock function with given fields: textToSlugify, maxLength
func (_m *MockSlugGeneratorInterface) Generate(textToSlugify string, maxLength int) string {
	ret := _m.Called(textToSlugify, maxLength)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, int) string); ok {
		r0 = rf(textToSlugify, maxLength)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockSlugGeneratorInterface_Generate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Generate'
type MockSlugGeneratorInterface_Generate_Call struct {
	*mock.Call
}

// Generate is a helper method to define mock.On call
//   - textToSlugify string
//   - maxLength int
func (_e *MockSlugGeneratorInterface_Expecter) Generate(textToSlugify interface{}, maxLength interface{}) *MockSlugGeneratorInterface_Generate_Call {
	return &MockSlugGeneratorInterface_Generate_Call{Call: _e.mock.On("Generate", textToSlugify, maxLength)}
}

func (_c *MockSlugGeneratorInterface_Generate_Call) Run(run func(textToSlugify string, maxLength int)) *MockSlugGeneratorInterface_Generate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int))
	})
	return _c
}

func (_c *MockSlugGeneratorInterface_Generate_Call) Return(_a0 string) *MockSlugGeneratorInterface_Generate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSlugGeneratorInterface_Generate_Call) RunAndReturn(run func(string, int) string) *MockSlugGeneratorInterface_Generate_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockSlugGeneratorInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockSlugGeneratorInterface creates a new instance of MockSlugGeneratorInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockSlugGeneratorInterface(t mockConstructorTestingTNewMockSlugGeneratorInterface) *MockSlugGeneratorInterface {
	mock := &MockSlugGeneratorInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}