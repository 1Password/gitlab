// Code generated by MockGen. DO NOT EDIT.
// Source: issuesstatisticsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=issuesstatisticsservice_inf.go -destination=mocks/issuesstatisticsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockIssuesStatisticsService is a mock of IssuesStatisticsService interface.
type MockIssuesStatisticsService struct {
	ctrl     *gomock.Controller
	recorder *MockIssuesStatisticsServiceMockRecorder
	isgomock struct{}
}

// MockIssuesStatisticsServiceMockRecorder is the mock recorder for MockIssuesStatisticsService.
type MockIssuesStatisticsServiceMockRecorder struct {
	mock *MockIssuesStatisticsService
}

// NewMockIssuesStatisticsService creates a new mock instance.
func NewMockIssuesStatisticsService(ctrl *gomock.Controller) *MockIssuesStatisticsService {
	mock := &MockIssuesStatisticsService{ctrl: ctrl}
	mock.recorder = &MockIssuesStatisticsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIssuesStatisticsService) EXPECT() *MockIssuesStatisticsServiceMockRecorder {
	return m.recorder
}

// GetGroupIssuesStatistics mocks base method.
func (m *MockIssuesStatisticsService) GetGroupIssuesStatistics(gid any, opt *gitlab.GetGroupIssuesStatisticsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.IssuesStatistics, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroupIssuesStatistics", varargs...)
	ret0, _ := ret[0].(*gitlab.IssuesStatistics)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGroupIssuesStatistics indicates an expected call of GetGroupIssuesStatistics.
func (mr *MockIssuesStatisticsServiceMockRecorder) GetGroupIssuesStatistics(gid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupIssuesStatistics", reflect.TypeOf((*MockIssuesStatisticsService)(nil).GetGroupIssuesStatistics), varargs...)
}

// GetIssuesStatistics mocks base method.
func (m *MockIssuesStatisticsService) GetIssuesStatistics(opt *gitlab.GetIssuesStatisticsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.IssuesStatistics, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIssuesStatistics", varargs...)
	ret0, _ := ret[0].(*gitlab.IssuesStatistics)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIssuesStatistics indicates an expected call of GetIssuesStatistics.
func (mr *MockIssuesStatisticsServiceMockRecorder) GetIssuesStatistics(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssuesStatistics", reflect.TypeOf((*MockIssuesStatisticsService)(nil).GetIssuesStatistics), varargs...)
}

// GetProjectIssuesStatistics mocks base method.
func (m *MockIssuesStatisticsService) GetProjectIssuesStatistics(pid any, opt *gitlab.GetProjectIssuesStatisticsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.IssuesStatistics, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProjectIssuesStatistics", varargs...)
	ret0, _ := ret[0].(*gitlab.IssuesStatistics)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProjectIssuesStatistics indicates an expected call of GetProjectIssuesStatistics.
func (mr *MockIssuesStatisticsServiceMockRecorder) GetProjectIssuesStatistics(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectIssuesStatistics", reflect.TypeOf((*MockIssuesStatisticsService)(nil).GetProjectIssuesStatistics), varargs...)
}
