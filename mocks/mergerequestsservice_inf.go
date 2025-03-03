// Code generated by MockGen. DO NOT EDIT.
// Source: mergerequestsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=mergerequestsservice_inf.go -destination=mocks/mergerequestsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockMergeRequestsService is a mock of MergeRequestsService interface.
type MockMergeRequestsService struct {
	ctrl     *gomock.Controller
	recorder *MockMergeRequestsServiceMockRecorder
	isgomock struct{}
}

// MockMergeRequestsServiceMockRecorder is the mock recorder for MockMergeRequestsService.
type MockMergeRequestsServiceMockRecorder struct {
	mock *MockMergeRequestsService
}

// NewMockMergeRequestsService creates a new mock instance.
func NewMockMergeRequestsService(ctrl *gomock.Controller) *MockMergeRequestsService {
	mock := &MockMergeRequestsService{ctrl: ctrl}
	mock.recorder = &MockMergeRequestsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMergeRequestsService) EXPECT() *MockMergeRequestsServiceMockRecorder {
	return m.recorder
}

// AcceptMergeRequest mocks base method.
func (m *MockMergeRequestsService) AcceptMergeRequest(pid any, mergeRequest int, opt *gitlab.AcceptMergeRequestOptions, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AcceptMergeRequest", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AcceptMergeRequest indicates an expected call of AcceptMergeRequest.
func (mr *MockMergeRequestsServiceMockRecorder) AcceptMergeRequest(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptMergeRequest", reflect.TypeOf((*MockMergeRequestsService)(nil).AcceptMergeRequest), varargs...)
}

// AddSpentTime mocks base method.
func (m *MockMergeRequestsService) AddSpentTime(pid any, mergeRequest int, opt *gitlab.AddSpentTimeOptions, options ...gitlab.RequestOptionFunc) (*gitlab.TimeStats, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddSpentTime", varargs...)
	ret0, _ := ret[0].(*gitlab.TimeStats)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddSpentTime indicates an expected call of AddSpentTime.
func (mr *MockMergeRequestsServiceMockRecorder) AddSpentTime(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSpentTime", reflect.TypeOf((*MockMergeRequestsService)(nil).AddSpentTime), varargs...)
}

// CancelMergeWhenPipelineSucceeds mocks base method.
func (m *MockMergeRequestsService) CancelMergeWhenPipelineSucceeds(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelMergeWhenPipelineSucceeds", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CancelMergeWhenPipelineSucceeds indicates an expected call of CancelMergeWhenPipelineSucceeds.
func (mr *MockMergeRequestsServiceMockRecorder) CancelMergeWhenPipelineSucceeds(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelMergeWhenPipelineSucceeds", reflect.TypeOf((*MockMergeRequestsService)(nil).CancelMergeWhenPipelineSucceeds), varargs...)
}

// CreateMergeRequest mocks base method.
func (m *MockMergeRequestsService) CreateMergeRequest(pid any, opt *gitlab.CreateMergeRequestOptions, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMergeRequest", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateMergeRequest indicates an expected call of CreateMergeRequest.
func (mr *MockMergeRequestsServiceMockRecorder) CreateMergeRequest(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMergeRequest", reflect.TypeOf((*MockMergeRequestsService)(nil).CreateMergeRequest), varargs...)
}

// CreateMergeRequestDependency mocks base method.
func (m *MockMergeRequestsService) CreateMergeRequestDependency(pid any, mergeRequest int, opts gitlab.CreateMergeRequestDependencyOptions, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequestDependency, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMergeRequestDependency", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequestDependency)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateMergeRequestDependency indicates an expected call of CreateMergeRequestDependency.
func (mr *MockMergeRequestsServiceMockRecorder) CreateMergeRequestDependency(pid, mergeRequest, opts any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opts}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMergeRequestDependency", reflect.TypeOf((*MockMergeRequestsService)(nil).CreateMergeRequestDependency), varargs...)
}

// CreateMergeRequestPipeline mocks base method.
func (m *MockMergeRequestsService) CreateMergeRequestPipeline(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) (*gitlab.PipelineInfo, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMergeRequestPipeline", varargs...)
	ret0, _ := ret[0].(*gitlab.PipelineInfo)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateMergeRequestPipeline indicates an expected call of CreateMergeRequestPipeline.
func (mr *MockMergeRequestsServiceMockRecorder) CreateMergeRequestPipeline(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMergeRequestPipeline", reflect.TypeOf((*MockMergeRequestsService)(nil).CreateMergeRequestPipeline), varargs...)
}

// CreateTodo mocks base method.
func (m *MockMergeRequestsService) CreateTodo(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) (*gitlab.Todo, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTodo", varargs...)
	ret0, _ := ret[0].(*gitlab.Todo)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateTodo indicates an expected call of CreateTodo.
func (mr *MockMergeRequestsServiceMockRecorder) CreateTodo(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockMergeRequestsService)(nil).CreateTodo), varargs...)
}

// DeleteMergeRequest mocks base method.
func (m *MockMergeRequestsService) DeleteMergeRequest(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMergeRequest", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMergeRequest indicates an expected call of DeleteMergeRequest.
func (mr *MockMergeRequestsServiceMockRecorder) DeleteMergeRequest(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMergeRequest", reflect.TypeOf((*MockMergeRequestsService)(nil).DeleteMergeRequest), varargs...)
}

// DeleteMergeRequestDependency mocks base method.
func (m *MockMergeRequestsService) DeleteMergeRequestDependency(pid any, mergeRequest, blockingMergeRequest int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, blockingMergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMergeRequestDependency", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMergeRequestDependency indicates an expected call of DeleteMergeRequestDependency.
func (mr *MockMergeRequestsServiceMockRecorder) DeleteMergeRequestDependency(pid, mergeRequest, blockingMergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, blockingMergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMergeRequestDependency", reflect.TypeOf((*MockMergeRequestsService)(nil).DeleteMergeRequestDependency), varargs...)
}

// GetIssuesClosedOnMerge mocks base method.
func (m *MockMergeRequestsService) GetIssuesClosedOnMerge(pid any, mergeRequest int, opt *gitlab.GetIssuesClosedOnMergeOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Issue, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIssuesClosedOnMerge", varargs...)
	ret0, _ := ret[0].([]*gitlab.Issue)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIssuesClosedOnMerge indicates an expected call of GetIssuesClosedOnMerge.
func (mr *MockMergeRequestsServiceMockRecorder) GetIssuesClosedOnMerge(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssuesClosedOnMerge", reflect.TypeOf((*MockMergeRequestsService)(nil).GetIssuesClosedOnMerge), varargs...)
}

// GetMergeRequest mocks base method.
func (m *MockMergeRequestsService) GetMergeRequest(pid any, mergeRequest int, opt *gitlab.GetMergeRequestsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMergeRequest", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMergeRequest indicates an expected call of GetMergeRequest.
func (mr *MockMergeRequestsServiceMockRecorder) GetMergeRequest(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeRequest", reflect.TypeOf((*MockMergeRequestsService)(nil).GetMergeRequest), varargs...)
}

// GetMergeRequestApprovals mocks base method.
func (m *MockMergeRequestsService) GetMergeRequestApprovals(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequestApprovals, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMergeRequestApprovals", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequestApprovals)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMergeRequestApprovals indicates an expected call of GetMergeRequestApprovals.
func (mr *MockMergeRequestsServiceMockRecorder) GetMergeRequestApprovals(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeRequestApprovals", reflect.TypeOf((*MockMergeRequestsService)(nil).GetMergeRequestApprovals), varargs...)
}

// GetMergeRequestChanges mocks base method.
func (m *MockMergeRequestsService) GetMergeRequestChanges(pid any, mergeRequest int, opt *gitlab.GetMergeRequestChangesOptions, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMergeRequestChanges", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMergeRequestChanges indicates an expected call of GetMergeRequestChanges.
func (mr *MockMergeRequestsServiceMockRecorder) GetMergeRequestChanges(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeRequestChanges", reflect.TypeOf((*MockMergeRequestsService)(nil).GetMergeRequestChanges), varargs...)
}

// GetMergeRequestCommits mocks base method.
func (m *MockMergeRequestsService) GetMergeRequestCommits(pid any, mergeRequest int, opt *gitlab.GetMergeRequestCommitsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Commit, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMergeRequestCommits", varargs...)
	ret0, _ := ret[0].([]*gitlab.Commit)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMergeRequestCommits indicates an expected call of GetMergeRequestCommits.
func (mr *MockMergeRequestsServiceMockRecorder) GetMergeRequestCommits(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeRequestCommits", reflect.TypeOf((*MockMergeRequestsService)(nil).GetMergeRequestCommits), varargs...)
}

// GetMergeRequestDependencies mocks base method.
func (m *MockMergeRequestsService) GetMergeRequestDependencies(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) ([]gitlab.MergeRequestDependency, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMergeRequestDependencies", varargs...)
	ret0, _ := ret[0].([]gitlab.MergeRequestDependency)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMergeRequestDependencies indicates an expected call of GetMergeRequestDependencies.
func (mr *MockMergeRequestsServiceMockRecorder) GetMergeRequestDependencies(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeRequestDependencies", reflect.TypeOf((*MockMergeRequestsService)(nil).GetMergeRequestDependencies), varargs...)
}

// GetMergeRequestDiffVersions mocks base method.
func (m *MockMergeRequestsService) GetMergeRequestDiffVersions(pid any, mergeRequest int, opt *gitlab.GetMergeRequestDiffVersionsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.MergeRequestDiffVersion, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMergeRequestDiffVersions", varargs...)
	ret0, _ := ret[0].([]*gitlab.MergeRequestDiffVersion)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMergeRequestDiffVersions indicates an expected call of GetMergeRequestDiffVersions.
func (mr *MockMergeRequestsServiceMockRecorder) GetMergeRequestDiffVersions(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeRequestDiffVersions", reflect.TypeOf((*MockMergeRequestsService)(nil).GetMergeRequestDiffVersions), varargs...)
}

// GetMergeRequestParticipants mocks base method.
func (m *MockMergeRequestsService) GetMergeRequestParticipants(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) ([]*gitlab.BasicUser, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMergeRequestParticipants", varargs...)
	ret0, _ := ret[0].([]*gitlab.BasicUser)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMergeRequestParticipants indicates an expected call of GetMergeRequestParticipants.
func (mr *MockMergeRequestsServiceMockRecorder) GetMergeRequestParticipants(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeRequestParticipants", reflect.TypeOf((*MockMergeRequestsService)(nil).GetMergeRequestParticipants), varargs...)
}

// GetMergeRequestReviewers mocks base method.
func (m *MockMergeRequestsService) GetMergeRequestReviewers(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) ([]*gitlab.MergeRequestReviewer, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMergeRequestReviewers", varargs...)
	ret0, _ := ret[0].([]*gitlab.MergeRequestReviewer)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMergeRequestReviewers indicates an expected call of GetMergeRequestReviewers.
func (mr *MockMergeRequestsServiceMockRecorder) GetMergeRequestReviewers(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeRequestReviewers", reflect.TypeOf((*MockMergeRequestsService)(nil).GetMergeRequestReviewers), varargs...)
}

// GetSingleMergeRequestDiffVersion mocks base method.
func (m *MockMergeRequestsService) GetSingleMergeRequestDiffVersion(pid any, mergeRequest, version int, opt *gitlab.GetSingleMergeRequestDiffVersionOptions, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequestDiffVersion, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, version, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSingleMergeRequestDiffVersion", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequestDiffVersion)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSingleMergeRequestDiffVersion indicates an expected call of GetSingleMergeRequestDiffVersion.
func (mr *MockMergeRequestsServiceMockRecorder) GetSingleMergeRequestDiffVersion(pid, mergeRequest, version, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, version, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSingleMergeRequestDiffVersion", reflect.TypeOf((*MockMergeRequestsService)(nil).GetSingleMergeRequestDiffVersion), varargs...)
}

// GetTimeSpent mocks base method.
func (m *MockMergeRequestsService) GetTimeSpent(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) (*gitlab.TimeStats, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTimeSpent", varargs...)
	ret0, _ := ret[0].(*gitlab.TimeStats)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTimeSpent indicates an expected call of GetTimeSpent.
func (mr *MockMergeRequestsServiceMockRecorder) GetTimeSpent(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeSpent", reflect.TypeOf((*MockMergeRequestsService)(nil).GetTimeSpent), varargs...)
}

// ListGroupMergeRequests mocks base method.
func (m *MockMergeRequestsService) ListGroupMergeRequests(gid any, opt *gitlab.ListGroupMergeRequestsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.BasicMergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroupMergeRequests", varargs...)
	ret0, _ := ret[0].([]*gitlab.BasicMergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListGroupMergeRequests indicates an expected call of ListGroupMergeRequests.
func (mr *MockMergeRequestsServiceMockRecorder) ListGroupMergeRequests(gid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupMergeRequests", reflect.TypeOf((*MockMergeRequestsService)(nil).ListGroupMergeRequests), varargs...)
}

// ListMergeRequestDiffs mocks base method.
func (m *MockMergeRequestsService) ListMergeRequestDiffs(pid any, mergeRequest int, opt *gitlab.ListMergeRequestDiffsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.MergeRequestDiff, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMergeRequestDiffs", varargs...)
	ret0, _ := ret[0].([]*gitlab.MergeRequestDiff)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMergeRequestDiffs indicates an expected call of ListMergeRequestDiffs.
func (mr *MockMergeRequestsServiceMockRecorder) ListMergeRequestDiffs(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMergeRequestDiffs", reflect.TypeOf((*MockMergeRequestsService)(nil).ListMergeRequestDiffs), varargs...)
}

// ListMergeRequestPipelines mocks base method.
func (m *MockMergeRequestsService) ListMergeRequestPipelines(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) ([]*gitlab.PipelineInfo, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMergeRequestPipelines", varargs...)
	ret0, _ := ret[0].([]*gitlab.PipelineInfo)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMergeRequestPipelines indicates an expected call of ListMergeRequestPipelines.
func (mr *MockMergeRequestsServiceMockRecorder) ListMergeRequestPipelines(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMergeRequestPipelines", reflect.TypeOf((*MockMergeRequestsService)(nil).ListMergeRequestPipelines), varargs...)
}

// ListMergeRequests mocks base method.
func (m *MockMergeRequestsService) ListMergeRequests(opt *gitlab.ListMergeRequestsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.BasicMergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMergeRequests", varargs...)
	ret0, _ := ret[0].([]*gitlab.BasicMergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMergeRequests indicates an expected call of ListMergeRequests.
func (mr *MockMergeRequestsServiceMockRecorder) ListMergeRequests(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMergeRequests", reflect.TypeOf((*MockMergeRequestsService)(nil).ListMergeRequests), varargs...)
}

// ListProjectMergeRequests mocks base method.
func (m *MockMergeRequestsService) ListProjectMergeRequests(pid any, opt *gitlab.ListProjectMergeRequestsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.BasicMergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectMergeRequests", varargs...)
	ret0, _ := ret[0].([]*gitlab.BasicMergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjectMergeRequests indicates an expected call of ListProjectMergeRequests.
func (mr *MockMergeRequestsServiceMockRecorder) ListProjectMergeRequests(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectMergeRequests", reflect.TypeOf((*MockMergeRequestsService)(nil).ListProjectMergeRequests), varargs...)
}

// RebaseMergeRequest mocks base method.
func (m *MockMergeRequestsService) RebaseMergeRequest(pid any, mergeRequest int, opt *gitlab.RebaseMergeRequestOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RebaseMergeRequest", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RebaseMergeRequest indicates an expected call of RebaseMergeRequest.
func (mr *MockMergeRequestsServiceMockRecorder) RebaseMergeRequest(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RebaseMergeRequest", reflect.TypeOf((*MockMergeRequestsService)(nil).RebaseMergeRequest), varargs...)
}

// ResetSpentTime mocks base method.
func (m *MockMergeRequestsService) ResetSpentTime(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) (*gitlab.TimeStats, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResetSpentTime", varargs...)
	ret0, _ := ret[0].(*gitlab.TimeStats)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ResetSpentTime indicates an expected call of ResetSpentTime.
func (mr *MockMergeRequestsServiceMockRecorder) ResetSpentTime(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetSpentTime", reflect.TypeOf((*MockMergeRequestsService)(nil).ResetSpentTime), varargs...)
}

// ResetTimeEstimate mocks base method.
func (m *MockMergeRequestsService) ResetTimeEstimate(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) (*gitlab.TimeStats, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResetTimeEstimate", varargs...)
	ret0, _ := ret[0].(*gitlab.TimeStats)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ResetTimeEstimate indicates an expected call of ResetTimeEstimate.
func (mr *MockMergeRequestsServiceMockRecorder) ResetTimeEstimate(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetTimeEstimate", reflect.TypeOf((*MockMergeRequestsService)(nil).ResetTimeEstimate), varargs...)
}

// SetTimeEstimate mocks base method.
func (m *MockMergeRequestsService) SetTimeEstimate(pid any, mergeRequest int, opt *gitlab.SetTimeEstimateOptions, options ...gitlab.RequestOptionFunc) (*gitlab.TimeStats, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetTimeEstimate", varargs...)
	ret0, _ := ret[0].(*gitlab.TimeStats)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SetTimeEstimate indicates an expected call of SetTimeEstimate.
func (mr *MockMergeRequestsServiceMockRecorder) SetTimeEstimate(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimeEstimate", reflect.TypeOf((*MockMergeRequestsService)(nil).SetTimeEstimate), varargs...)
}

// ShowMergeRequestRawDiffs mocks base method.
func (m *MockMergeRequestsService) ShowMergeRequestRawDiffs(pid any, mergeRequest int, opt *gitlab.ShowMergeRequestRawDiffsOptions, options ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ShowMergeRequestRawDiffs", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ShowMergeRequestRawDiffs indicates an expected call of ShowMergeRequestRawDiffs.
func (mr *MockMergeRequestsServiceMockRecorder) ShowMergeRequestRawDiffs(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowMergeRequestRawDiffs", reflect.TypeOf((*MockMergeRequestsService)(nil).ShowMergeRequestRawDiffs), varargs...)
}

// SubscribeToMergeRequest mocks base method.
func (m *MockMergeRequestsService) SubscribeToMergeRequest(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeToMergeRequest", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SubscribeToMergeRequest indicates an expected call of SubscribeToMergeRequest.
func (mr *MockMergeRequestsServiceMockRecorder) SubscribeToMergeRequest(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToMergeRequest", reflect.TypeOf((*MockMergeRequestsService)(nil).SubscribeToMergeRequest), varargs...)
}

// UnsubscribeFromMergeRequest mocks base method.
func (m *MockMergeRequestsService) UnsubscribeFromMergeRequest(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsubscribeFromMergeRequest", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UnsubscribeFromMergeRequest indicates an expected call of UnsubscribeFromMergeRequest.
func (mr *MockMergeRequestsServiceMockRecorder) UnsubscribeFromMergeRequest(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeFromMergeRequest", reflect.TypeOf((*MockMergeRequestsService)(nil).UnsubscribeFromMergeRequest), varargs...)
}

// UpdateMergeRequest mocks base method.
func (m *MockMergeRequestsService) UpdateMergeRequest(pid any, mergeRequest int, opt *gitlab.UpdateMergeRequestOptions, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateMergeRequest", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateMergeRequest indicates an expected call of UpdateMergeRequest.
func (mr *MockMergeRequestsServiceMockRecorder) UpdateMergeRequest(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMergeRequest", reflect.TypeOf((*MockMergeRequestsService)(nil).UpdateMergeRequest), varargs...)
}
