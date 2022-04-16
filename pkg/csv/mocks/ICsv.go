// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	bytes "bytes"

	mock "github.com/stretchr/testify/mock"
)

// ICsv is an autogenerated mock type for the ICsv type
type ICsv struct {
	mock.Mock
}

// Export provides a mock function with given fields:
func (_m *ICsv) Export() (bytes.Buffer, error) {
	ret := _m.Called()

	var r0 bytes.Buffer
	if rf, ok := ret.Get(0).(func() bytes.Buffer); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bytes.Buffer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContents provides a mock function with given fields:
func (_m *ICsv) GetContents() [][]string {
	ret := _m.Called()

	var r0 [][]string
	if rf, ok := ret.Get(0).(func() [][]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]string)
		}
	}

	return r0
}

// GetHeader provides a mock function with given fields:
func (_m *ICsv) GetHeader() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Save provides a mock function with given fields: path
func (_m *ICsv) Save(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}