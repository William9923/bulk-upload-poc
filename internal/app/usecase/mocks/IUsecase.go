// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	bytes "bytes"
	context "context"

	io "io"

	mock "github.com/stretchr/testify/mock"

	usecase "github.com/William9923/bulk-upload-poc/internal/app/usecase"
)

// IUsecase is an autogenerated mock type for the IUsecase type
type IUsecase struct {
	mock.Mock
}

// ExportResult provides a mock function with given fields: ctx, id
func (_m *IUsecase) ExportResult(ctx context.Context, id int64) (bytes.Buffer, error) {
	ret := _m.Called(ctx, id)

	var r0 bytes.Buffer
	if rf, ok := ret.Get(0).(func(context.Context, int64) bytes.Buffer); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bytes.Buffer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExportUsers provides a mock function with given fields: ctx
func (_m *IUsecase) ExportUsers(ctx context.Context) (bytes.Buffer, error) {
	ret := _m.Called(ctx)

	var r0 bytes.Buffer
	if rf, ok := ret.Get(0).(func(context.Context) bytes.Buffer); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bytes.Buffer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowResults provides a mock function with given fields: ctx
func (_m *IUsecase) ShowResults(ctx context.Context) (usecase.ResultsDTO, error) {
	ret := _m.Called(ctx)

	var r0 usecase.ResultsDTO
	if rf, ok := ret.Get(0).(func(context.Context) usecase.ResultsDTO); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(usecase.ResultsDTO)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShowUsers provides a mock function with given fields: ctx
func (_m *IUsecase) ShowUsers(ctx context.Context) (usecase.UsersDTO, error) {
	ret := _m.Called(ctx)

	var r0 usecase.UsersDTO
	if rf, ok := ret.Get(0).(func(context.Context) usecase.UsersDTO); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(usecase.UsersDTO)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadWhitelists provides a mock function with given fields: ctx, file
func (_m *IUsecase) UploadWhitelists(ctx context.Context, file io.Reader) (usecase.ResultDTO, error) {
	ret := _m.Called(ctx, file)

	var r0 usecase.ResultDTO
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader) usecase.ResultDTO); ok {
		r0 = rf(ctx, file)
	} else {
		r0 = ret.Get(0).(usecase.ResultDTO)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, io.Reader) error); ok {
		r1 = rf(ctx, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
