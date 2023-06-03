// Code generated by mockery v2.28.0. DO NOT EDIT.

package themoviedb

import mock "github.com/stretchr/testify/mock"

// MockTheMovieDbClientInterface is an autogenerated mock type for the TheMovieDbClientInterface type
type MockTheMovieDbClientInterface struct {
	mock.Mock
}

type MockTheMovieDbClientInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTheMovieDbClientInterface) EXPECT() *MockTheMovieDbClientInterface_Expecter {
	return &MockTheMovieDbClientInterface_Expecter{mock: &_m.Mock}
}

// GetConfiguration provides a mock function with given fields:
func (_m *MockTheMovieDbClientInterface) GetConfiguration() (ApiConfigurationResult, error) {
	ret := _m.Called()

	var r0 ApiConfigurationResult
	var r1 error
	if rf, ok := ret.Get(0).(func() (ApiConfigurationResult, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() ApiConfigurationResult); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ApiConfigurationResult)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTheMovieDbClientInterface_GetConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConfiguration'
type MockTheMovieDbClientInterface_GetConfiguration_Call struct {
	*mock.Call
}

// GetConfiguration is a helper method to define mock.On call
func (_e *MockTheMovieDbClientInterface_Expecter) GetConfiguration() *MockTheMovieDbClientInterface_GetConfiguration_Call {
	return &MockTheMovieDbClientInterface_GetConfiguration_Call{Call: _e.mock.On("GetConfiguration")}
}

func (_c *MockTheMovieDbClientInterface_GetConfiguration_Call) Run(run func()) *MockTheMovieDbClientInterface_GetConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTheMovieDbClientInterface_GetConfiguration_Call) Return(_a0 ApiConfigurationResult, _a1 error) *MockTheMovieDbClientInterface_GetConfiguration_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTheMovieDbClientInterface_GetConfiguration_Call) RunAndReturn(run func() (ApiConfigurationResult, error)) *MockTheMovieDbClientInterface_GetConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// GetShow provides a mock function with given fields: id
func (_m *MockTheMovieDbClientInterface) GetShow(id int) (ApiShowResult, error) {
	ret := _m.Called(id)

	var r0 ApiShowResult
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (ApiShowResult, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) ApiShowResult); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(ApiShowResult)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTheMovieDbClientInterface_GetShow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetShow'
type MockTheMovieDbClientInterface_GetShow_Call struct {
	*mock.Call
}

// GetShow is a helper method to define mock.On call
//   - id int
func (_e *MockTheMovieDbClientInterface_Expecter) GetShow(id interface{}) *MockTheMovieDbClientInterface_GetShow_Call {
	return &MockTheMovieDbClientInterface_GetShow_Call{Call: _e.mock.On("GetShow", id)}
}

func (_c *MockTheMovieDbClientInterface_GetShow_Call) Run(run func(id int)) *MockTheMovieDbClientInterface_GetShow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockTheMovieDbClientInterface_GetShow_Call) Return(_a0 ApiShowResult, _a1 error) *MockTheMovieDbClientInterface_GetShow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTheMovieDbClientInterface_GetShow_Call) RunAndReturn(run func(int) (ApiShowResult, error)) *MockTheMovieDbClientInterface_GetShow_Call {
	_c.Call.Return(run)
	return _c
}

// GetShowSeason provides a mock function with given fields: id, seasonNumber
func (_m *MockTheMovieDbClientInterface) GetShowSeason(id int, seasonNumber int) (ApiSeasonResult, error) {
	ret := _m.Called(id, seasonNumber)

	var r0 ApiSeasonResult
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (ApiSeasonResult, error)); ok {
		return rf(id, seasonNumber)
	}
	if rf, ok := ret.Get(0).(func(int, int) ApiSeasonResult); ok {
		r0 = rf(id, seasonNumber)
	} else {
		r0 = ret.Get(0).(ApiSeasonResult)
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(id, seasonNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTheMovieDbClientInterface_GetShowSeason_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetShowSeason'
type MockTheMovieDbClientInterface_GetShowSeason_Call struct {
	*mock.Call
}

// GetShowSeason is a helper method to define mock.On call
//   - id int
//   - seasonNumber int
func (_e *MockTheMovieDbClientInterface_Expecter) GetShowSeason(id interface{}, seasonNumber interface{}) *MockTheMovieDbClientInterface_GetShowSeason_Call {
	return &MockTheMovieDbClientInterface_GetShowSeason_Call{Call: _e.mock.On("GetShowSeason", id, seasonNumber)}
}

func (_c *MockTheMovieDbClientInterface_GetShowSeason_Call) Run(run func(id int, seasonNumber int)) *MockTheMovieDbClientInterface_GetShowSeason_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int))
	})
	return _c
}

func (_c *MockTheMovieDbClientInterface_GetShowSeason_Call) Return(_a0 ApiSeasonResult, _a1 error) *MockTheMovieDbClientInterface_GetShowSeason_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTheMovieDbClientInterface_GetShowSeason_Call) RunAndReturn(run func(int, int) (ApiSeasonResult, error)) *MockTheMovieDbClientInterface_GetShowSeason_Call {
	_c.Call.Return(run)
	return _c
}

// GetTrendingShows provides a mock function with given fields: timeWindow
func (_m *MockTheMovieDbClientInterface) GetTrendingShows(timeWindow string) (ApiTrendingShowsResult, error) {
	ret := _m.Called(timeWindow)

	var r0 ApiTrendingShowsResult
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (ApiTrendingShowsResult, error)); ok {
		return rf(timeWindow)
	}
	if rf, ok := ret.Get(0).(func(string) ApiTrendingShowsResult); ok {
		r0 = rf(timeWindow)
	} else {
		r0 = ret.Get(0).(ApiTrendingShowsResult)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(timeWindow)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTheMovieDbClientInterface_GetTrendingShows_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTrendingShows'
type MockTheMovieDbClientInterface_GetTrendingShows_Call struct {
	*mock.Call
}

// GetTrendingShows is a helper method to define mock.On call
//   - timeWindow string
func (_e *MockTheMovieDbClientInterface_Expecter) GetTrendingShows(timeWindow interface{}) *MockTheMovieDbClientInterface_GetTrendingShows_Call {
	return &MockTheMovieDbClientInterface_GetTrendingShows_Call{Call: _e.mock.On("GetTrendingShows", timeWindow)}
}

func (_c *MockTheMovieDbClientInterface_GetTrendingShows_Call) Run(run func(timeWindow string)) *MockTheMovieDbClientInterface_GetTrendingShows_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockTheMovieDbClientInterface_GetTrendingShows_Call) Return(_a0 ApiTrendingShowsResult, _a1 error) *MockTheMovieDbClientInterface_GetTrendingShows_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTheMovieDbClientInterface_GetTrendingShows_Call) RunAndReturn(run func(string) (ApiTrendingShowsResult, error)) *MockTheMovieDbClientInterface_GetTrendingShows_Call {
	_c.Call.Return(run)
	return _c
}

// SearchShows provides a mock function with given fields: searchToken, page
func (_m *MockTheMovieDbClientInterface) SearchShows(searchToken string, page int) (ApiSearchShowsResult, error) {
	ret := _m.Called(searchToken, page)

	var r0 ApiSearchShowsResult
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int) (ApiSearchShowsResult, error)); ok {
		return rf(searchToken, page)
	}
	if rf, ok := ret.Get(0).(func(string, int) ApiSearchShowsResult); ok {
		r0 = rf(searchToken, page)
	} else {
		r0 = ret.Get(0).(ApiSearchShowsResult)
	}

	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(searchToken, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTheMovieDbClientInterface_SearchShows_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchShows'
type MockTheMovieDbClientInterface_SearchShows_Call struct {
	*mock.Call
}

// SearchShows is a helper method to define mock.On call
//   - searchToken string
//   - page int
func (_e *MockTheMovieDbClientInterface_Expecter) SearchShows(searchToken interface{}, page interface{}) *MockTheMovieDbClientInterface_SearchShows_Call {
	return &MockTheMovieDbClientInterface_SearchShows_Call{Call: _e.mock.On("SearchShows", searchToken, page)}
}

func (_c *MockTheMovieDbClientInterface_SearchShows_Call) Run(run func(searchToken string, page int)) *MockTheMovieDbClientInterface_SearchShows_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int))
	})
	return _c
}

func (_c *MockTheMovieDbClientInterface_SearchShows_Call) Return(_a0 ApiSearchShowsResult, _a1 error) *MockTheMovieDbClientInterface_SearchShows_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTheMovieDbClientInterface_SearchShows_Call) RunAndReturn(run func(string, int) (ApiSearchShowsResult, error)) *MockTheMovieDbClientInterface_SearchShows_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockTheMovieDbClientInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockTheMovieDbClientInterface creates a new instance of MockTheMovieDbClientInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTheMovieDbClientInterface(t mockConstructorTestingTNewMockTheMovieDbClientInterface) *MockTheMovieDbClientInterface {
	mock := &MockTheMovieDbClientInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
