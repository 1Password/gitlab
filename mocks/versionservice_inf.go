// Code generated by MockGen. DO NOT EDIT.
// Source: versionservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=versionservice_inf.go -destination=mocks/versionservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockVersionService is a mock of VersionService interface.
type MockVersionService struct {
	ctrl     *gomock.Controller
	recorder *MockVersionServiceMockRecorder
}

// MockVersionServiceMockRecorder is the mock recorder for MockVersionService.
type MockVersionServiceMockRecorder struct {
	mock *MockVersionService
}

// NewMockVersionService creates a new mock instance.
func NewMockVersionService(ctrl *gomock.Controller) *MockVersionService {
	mock := &MockVersionService{ctrl: ctrl}
	mock.recorder = &MockVersionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVersionService) EXPECT() *MockVersionServiceMockRecorder {
	return m.recorder
}

// GetVersion mocks base method.
func (m *MockVersionService) GetVersion(options ...gitlab.RequestOptionFunc) (*gitlab.Version, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVersion", varargs...)
	ret0, _ := ret[0].(*gitlab.Version)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetVersion indicates an expected call of GetVersion.
func (mr *MockVersionServiceMockRecorder) GetVersion(options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockVersionService)(nil).GetVersion), options...)
}
