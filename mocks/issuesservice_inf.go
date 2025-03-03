// Code generated by MockGen. DO NOT EDIT.
// Source: issuesservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=issuesservice_inf.go -destination=mocks/issuesservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockIssuesService is a mock of IssuesService interface.
type MockIssuesService struct {
	ctrl     *gomock.Controller
	recorder *MockIssuesServiceMockRecorder
	isgomock struct{}
}

// MockIssuesServiceMockRecorder is the mock recorder for MockIssuesService.
type MockIssuesServiceMockRecorder struct {
	mock *MockIssuesService
}

// NewMockIssuesService creates a new mock instance.
func NewMockIssuesService(ctrl *gomock.Controller) *MockIssuesService {
	mock := &MockIssuesService{ctrl: ctrl}
	mock.recorder = &MockIssuesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIssuesService) EXPECT() *MockIssuesServiceMockRecorder {
	return m.recorder
}

// AddSpentTime mocks base method.
func (m *MockIssuesService) AddSpentTime(pid any, issue int, opt *gitlab.AddSpentTimeOptions, options ...gitlab.RequestOptionFunc) (*gitlab.TimeStats, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, opt}
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
func (mr *MockIssuesServiceMockRecorder) AddSpentTime(pid, issue, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSpentTime", reflect.TypeOf((*MockIssuesService)(nil).AddSpentTime), varargs...)
}

// CreateIssue mocks base method.
func (m *MockIssuesService) CreateIssue(pid any, opt *gitlab.CreateIssueOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Issue, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateIssue", varargs...)
	ret0, _ := ret[0].(*gitlab.Issue)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateIssue indicates an expected call of CreateIssue.
func (mr *MockIssuesServiceMockRecorder) CreateIssue(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIssue", reflect.TypeOf((*MockIssuesService)(nil).CreateIssue), varargs...)
}

// CreateTodo mocks base method.
func (m *MockIssuesService) CreateTodo(pid any, issue int, options ...gitlab.RequestOptionFunc) (*gitlab.Todo, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue}
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
func (mr *MockIssuesServiceMockRecorder) CreateTodo(pid, issue any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockIssuesService)(nil).CreateTodo), varargs...)
}

// DeleteIssue mocks base method.
func (m *MockIssuesService) DeleteIssue(pid any, issue int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteIssue", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteIssue indicates an expected call of DeleteIssue.
func (mr *MockIssuesServiceMockRecorder) DeleteIssue(pid, issue any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIssue", reflect.TypeOf((*MockIssuesService)(nil).DeleteIssue), varargs...)
}

// GetIssue mocks base method.
func (m *MockIssuesService) GetIssue(pid any, issue int, options ...gitlab.RequestOptionFunc) (*gitlab.Issue, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIssue", varargs...)
	ret0, _ := ret[0].(*gitlab.Issue)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIssue indicates an expected call of GetIssue.
func (mr *MockIssuesServiceMockRecorder) GetIssue(pid, issue any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssue", reflect.TypeOf((*MockIssuesService)(nil).GetIssue), varargs...)
}

// GetIssueByID mocks base method.
func (m *MockIssuesService) GetIssueByID(issue int, options ...gitlab.RequestOptionFunc) (*gitlab.Issue, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{issue}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIssueByID", varargs...)
	ret0, _ := ret[0].(*gitlab.Issue)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIssueByID indicates an expected call of GetIssueByID.
func (mr *MockIssuesServiceMockRecorder) GetIssueByID(issue any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{issue}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueByID", reflect.TypeOf((*MockIssuesService)(nil).GetIssueByID), varargs...)
}

// GetParticipants mocks base method.
func (m *MockIssuesService) GetParticipants(pid any, issue int, options ...gitlab.RequestOptionFunc) ([]*gitlab.BasicUser, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetParticipants", varargs...)
	ret0, _ := ret[0].([]*gitlab.BasicUser)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetParticipants indicates an expected call of GetParticipants.
func (mr *MockIssuesServiceMockRecorder) GetParticipants(pid, issue any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParticipants", reflect.TypeOf((*MockIssuesService)(nil).GetParticipants), varargs...)
}

// GetTimeSpent mocks base method.
func (m *MockIssuesService) GetTimeSpent(pid any, issue int, options ...gitlab.RequestOptionFunc) (*gitlab.TimeStats, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue}
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
func (mr *MockIssuesServiceMockRecorder) GetTimeSpent(pid, issue any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeSpent", reflect.TypeOf((*MockIssuesService)(nil).GetTimeSpent), varargs...)
}

// ListGroupIssues mocks base method.
func (m *MockIssuesService) ListGroupIssues(pid any, opt *gitlab.ListGroupIssuesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Issue, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroupIssues", varargs...)
	ret0, _ := ret[0].([]*gitlab.Issue)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListGroupIssues indicates an expected call of ListGroupIssues.
func (mr *MockIssuesServiceMockRecorder) ListGroupIssues(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupIssues", reflect.TypeOf((*MockIssuesService)(nil).ListGroupIssues), varargs...)
}

// ListIssues mocks base method.
func (m *MockIssuesService) ListIssues(opt *gitlab.ListIssuesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Issue, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIssues", varargs...)
	ret0, _ := ret[0].([]*gitlab.Issue)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListIssues indicates an expected call of ListIssues.
func (mr *MockIssuesServiceMockRecorder) ListIssues(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIssues", reflect.TypeOf((*MockIssuesService)(nil).ListIssues), varargs...)
}

// ListMergeRequestsClosingIssue mocks base method.
func (m *MockIssuesService) ListMergeRequestsClosingIssue(pid any, issue int, opt *gitlab.ListMergeRequestsClosingIssueOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.BasicMergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMergeRequestsClosingIssue", varargs...)
	ret0, _ := ret[0].([]*gitlab.BasicMergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMergeRequestsClosingIssue indicates an expected call of ListMergeRequestsClosingIssue.
func (mr *MockIssuesServiceMockRecorder) ListMergeRequestsClosingIssue(pid, issue, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMergeRequestsClosingIssue", reflect.TypeOf((*MockIssuesService)(nil).ListMergeRequestsClosingIssue), varargs...)
}

// ListMergeRequestsRelatedToIssue mocks base method.
func (m *MockIssuesService) ListMergeRequestsRelatedToIssue(pid any, issue int, opt *gitlab.ListMergeRequestsRelatedToIssueOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.BasicMergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMergeRequestsRelatedToIssue", varargs...)
	ret0, _ := ret[0].([]*gitlab.BasicMergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMergeRequestsRelatedToIssue indicates an expected call of ListMergeRequestsRelatedToIssue.
func (mr *MockIssuesServiceMockRecorder) ListMergeRequestsRelatedToIssue(pid, issue, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMergeRequestsRelatedToIssue", reflect.TypeOf((*MockIssuesService)(nil).ListMergeRequestsRelatedToIssue), varargs...)
}

// ListProjectIssues mocks base method.
func (m *MockIssuesService) ListProjectIssues(pid any, opt *gitlab.ListProjectIssuesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Issue, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectIssues", varargs...)
	ret0, _ := ret[0].([]*gitlab.Issue)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjectIssues indicates an expected call of ListProjectIssues.
func (mr *MockIssuesServiceMockRecorder) ListProjectIssues(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectIssues", reflect.TypeOf((*MockIssuesService)(nil).ListProjectIssues), varargs...)
}

// MoveIssue mocks base method.
func (m *MockIssuesService) MoveIssue(pid any, issue int, opt *gitlab.MoveIssueOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Issue, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MoveIssue", varargs...)
	ret0, _ := ret[0].(*gitlab.Issue)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MoveIssue indicates an expected call of MoveIssue.
func (mr *MockIssuesServiceMockRecorder) MoveIssue(pid, issue, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveIssue", reflect.TypeOf((*MockIssuesService)(nil).MoveIssue), varargs...)
}

// ReorderIssue mocks base method.
func (m *MockIssuesService) ReorderIssue(pid any, issue int, opt *gitlab.ReorderIssueOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Issue, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReorderIssue", varargs...)
	ret0, _ := ret[0].(*gitlab.Issue)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReorderIssue indicates an expected call of ReorderIssue.
func (mr *MockIssuesServiceMockRecorder) ReorderIssue(pid, issue, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReorderIssue", reflect.TypeOf((*MockIssuesService)(nil).ReorderIssue), varargs...)
}

// ResetSpentTime mocks base method.
func (m *MockIssuesService) ResetSpentTime(pid any, issue int, options ...gitlab.RequestOptionFunc) (*gitlab.TimeStats, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue}
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
func (mr *MockIssuesServiceMockRecorder) ResetSpentTime(pid, issue any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetSpentTime", reflect.TypeOf((*MockIssuesService)(nil).ResetSpentTime), varargs...)
}

// ResetTimeEstimate mocks base method.
func (m *MockIssuesService) ResetTimeEstimate(pid any, issue int, options ...gitlab.RequestOptionFunc) (*gitlab.TimeStats, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue}
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
func (mr *MockIssuesServiceMockRecorder) ResetTimeEstimate(pid, issue any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetTimeEstimate", reflect.TypeOf((*MockIssuesService)(nil).ResetTimeEstimate), varargs...)
}

// SetTimeEstimate mocks base method.
func (m *MockIssuesService) SetTimeEstimate(pid any, issue int, opt *gitlab.SetTimeEstimateOptions, options ...gitlab.RequestOptionFunc) (*gitlab.TimeStats, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, opt}
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
func (mr *MockIssuesServiceMockRecorder) SetTimeEstimate(pid, issue, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimeEstimate", reflect.TypeOf((*MockIssuesService)(nil).SetTimeEstimate), varargs...)
}

// SubscribeToIssue mocks base method.
func (m *MockIssuesService) SubscribeToIssue(pid any, issue int, options ...gitlab.RequestOptionFunc) (*gitlab.Issue, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeToIssue", varargs...)
	ret0, _ := ret[0].(*gitlab.Issue)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SubscribeToIssue indicates an expected call of SubscribeToIssue.
func (mr *MockIssuesServiceMockRecorder) SubscribeToIssue(pid, issue any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToIssue", reflect.TypeOf((*MockIssuesService)(nil).SubscribeToIssue), varargs...)
}

// UnsubscribeFromIssue mocks base method.
func (m *MockIssuesService) UnsubscribeFromIssue(pid any, issue int, options ...gitlab.RequestOptionFunc) (*gitlab.Issue, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsubscribeFromIssue", varargs...)
	ret0, _ := ret[0].(*gitlab.Issue)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UnsubscribeFromIssue indicates an expected call of UnsubscribeFromIssue.
func (mr *MockIssuesServiceMockRecorder) UnsubscribeFromIssue(pid, issue any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeFromIssue", reflect.TypeOf((*MockIssuesService)(nil).UnsubscribeFromIssue), varargs...)
}

// UpdateIssue mocks base method.
func (m *MockIssuesService) UpdateIssue(pid any, issue int, opt *gitlab.UpdateIssueOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Issue, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateIssue", varargs...)
	ret0, _ := ret[0].(*gitlab.Issue)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateIssue indicates an expected call of UpdateIssue.
func (mr *MockIssuesServiceMockRecorder) UpdateIssue(pid, issue, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIssue", reflect.TypeOf((*MockIssuesService)(nil).UpdateIssue), varargs...)
}
