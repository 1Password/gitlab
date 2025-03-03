// Code generated by MockGen. DO NOT EDIT.
// Source: mergerequestapprovalsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=mergerequestapprovalsservice_inf.go -destination=mocks/mergerequestapprovalsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockMergeRequestApprovalsService is a mock of MergeRequestApprovalsService interface.
type MockMergeRequestApprovalsService struct {
	ctrl     *gomock.Controller
	recorder *MockMergeRequestApprovalsServiceMockRecorder
	isgomock struct{}
}

// MockMergeRequestApprovalsServiceMockRecorder is the mock recorder for MockMergeRequestApprovalsService.
type MockMergeRequestApprovalsServiceMockRecorder struct {
	mock *MockMergeRequestApprovalsService
}

// NewMockMergeRequestApprovalsService creates a new mock instance.
func NewMockMergeRequestApprovalsService(ctrl *gomock.Controller) *MockMergeRequestApprovalsService {
	mock := &MockMergeRequestApprovalsService{ctrl: ctrl}
	mock.recorder = &MockMergeRequestApprovalsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMergeRequestApprovalsService) EXPECT() *MockMergeRequestApprovalsServiceMockRecorder {
	return m.recorder
}

// ApproveMergeRequest mocks base method.
func (m *MockMergeRequestApprovalsService) ApproveMergeRequest(pid any, mr int, opt *gitlab.ApproveMergeRequestOptions, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequestApprovals, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mr, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ApproveMergeRequest", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequestApprovals)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ApproveMergeRequest indicates an expected call of ApproveMergeRequest.
func (mr_2 *MockMergeRequestApprovalsServiceMockRecorder) ApproveMergeRequest(pid, mr, opt any, options ...any) *gomock.Call {
	mr_2.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mr, opt}, options...)
	return mr_2.mock.ctrl.RecordCallWithMethodType(mr_2.mock, "ApproveMergeRequest", reflect.TypeOf((*MockMergeRequestApprovalsService)(nil).ApproveMergeRequest), varargs...)
}

// ChangeAllowedApprovers mocks base method.
func (m *MockMergeRequestApprovalsService) ChangeAllowedApprovers(pid any, mergeRequest int, opt *gitlab.ChangeMergeRequestAllowedApproversOptions, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeAllowedApprovers", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChangeAllowedApprovers indicates an expected call of ChangeAllowedApprovers.
func (mr *MockMergeRequestApprovalsServiceMockRecorder) ChangeAllowedApprovers(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeAllowedApprovers", reflect.TypeOf((*MockMergeRequestApprovalsService)(nil).ChangeAllowedApprovers), varargs...)
}

// ChangeApprovalConfiguration mocks base method.
func (m *MockMergeRequestApprovalsService) ChangeApprovalConfiguration(pid any, mergeRequest int, opt *gitlab.ChangeMergeRequestApprovalConfigurationOptions, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeApprovalConfiguration", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChangeApprovalConfiguration indicates an expected call of ChangeApprovalConfiguration.
func (mr *MockMergeRequestApprovalsServiceMockRecorder) ChangeApprovalConfiguration(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeApprovalConfiguration", reflect.TypeOf((*MockMergeRequestApprovalsService)(nil).ChangeApprovalConfiguration), varargs...)
}

// CreateApprovalRule mocks base method.
func (m *MockMergeRequestApprovalsService) CreateApprovalRule(pid any, mergeRequest int, opt *gitlab.CreateMergeRequestApprovalRuleOptions, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequestApprovalRule, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateApprovalRule", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequestApprovalRule)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateApprovalRule indicates an expected call of CreateApprovalRule.
func (mr *MockMergeRequestApprovalsServiceMockRecorder) CreateApprovalRule(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApprovalRule", reflect.TypeOf((*MockMergeRequestApprovalsService)(nil).CreateApprovalRule), varargs...)
}

// DeleteApprovalRule mocks base method.
func (m *MockMergeRequestApprovalsService) DeleteApprovalRule(pid any, mergeRequest, approvalRule int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, approvalRule}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteApprovalRule", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteApprovalRule indicates an expected call of DeleteApprovalRule.
func (mr *MockMergeRequestApprovalsServiceMockRecorder) DeleteApprovalRule(pid, mergeRequest, approvalRule any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, approvalRule}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApprovalRule", reflect.TypeOf((*MockMergeRequestApprovalsService)(nil).DeleteApprovalRule), varargs...)
}

// GetApprovalRules mocks base method.
func (m *MockMergeRequestApprovalsService) GetApprovalRules(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) ([]*gitlab.MergeRequestApprovalRule, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetApprovalRules", varargs...)
	ret0, _ := ret[0].([]*gitlab.MergeRequestApprovalRule)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetApprovalRules indicates an expected call of GetApprovalRules.
func (mr *MockMergeRequestApprovalsServiceMockRecorder) GetApprovalRules(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApprovalRules", reflect.TypeOf((*MockMergeRequestApprovalsService)(nil).GetApprovalRules), varargs...)
}

// GetApprovalState mocks base method.
func (m *MockMergeRequestApprovalsService) GetApprovalState(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequestApprovalState, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetApprovalState", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequestApprovalState)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetApprovalState indicates an expected call of GetApprovalState.
func (mr *MockMergeRequestApprovalsServiceMockRecorder) GetApprovalState(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApprovalState", reflect.TypeOf((*MockMergeRequestApprovalsService)(nil).GetApprovalState), varargs...)
}

// GetConfiguration mocks base method.
func (m *MockMergeRequestApprovalsService) GetConfiguration(pid any, mr int, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequestApprovals, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mr}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfiguration", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequestApprovals)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetConfiguration indicates an expected call of GetConfiguration.
func (mr_2 *MockMergeRequestApprovalsServiceMockRecorder) GetConfiguration(pid, mr any, options ...any) *gomock.Call {
	mr_2.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mr}, options...)
	return mr_2.mock.ctrl.RecordCallWithMethodType(mr_2.mock, "GetConfiguration", reflect.TypeOf((*MockMergeRequestApprovalsService)(nil).GetConfiguration), varargs...)
}

// ResetApprovalsOfMergeRequest mocks base method.
func (m *MockMergeRequestApprovalsService) ResetApprovalsOfMergeRequest(pid any, mr int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mr}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResetApprovalsOfMergeRequest", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetApprovalsOfMergeRequest indicates an expected call of ResetApprovalsOfMergeRequest.
func (mr_2 *MockMergeRequestApprovalsServiceMockRecorder) ResetApprovalsOfMergeRequest(pid, mr any, options ...any) *gomock.Call {
	mr_2.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mr}, options...)
	return mr_2.mock.ctrl.RecordCallWithMethodType(mr_2.mock, "ResetApprovalsOfMergeRequest", reflect.TypeOf((*MockMergeRequestApprovalsService)(nil).ResetApprovalsOfMergeRequest), varargs...)
}

// UnapproveMergeRequest mocks base method.
func (m *MockMergeRequestApprovalsService) UnapproveMergeRequest(pid any, mr int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mr}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnapproveMergeRequest", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnapproveMergeRequest indicates an expected call of UnapproveMergeRequest.
func (mr_2 *MockMergeRequestApprovalsServiceMockRecorder) UnapproveMergeRequest(pid, mr any, options ...any) *gomock.Call {
	mr_2.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mr}, options...)
	return mr_2.mock.ctrl.RecordCallWithMethodType(mr_2.mock, "UnapproveMergeRequest", reflect.TypeOf((*MockMergeRequestApprovalsService)(nil).UnapproveMergeRequest), varargs...)
}

// UpdateApprovalRule mocks base method.
func (m *MockMergeRequestApprovalsService) UpdateApprovalRule(pid any, mergeRequest, approvalRule int, opt *gitlab.UpdateMergeRequestApprovalRuleOptions, options ...gitlab.RequestOptionFunc) (*gitlab.MergeRequestApprovalRule, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, approvalRule, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateApprovalRule", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeRequestApprovalRule)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateApprovalRule indicates an expected call of UpdateApprovalRule.
func (mr *MockMergeRequestApprovalsServiceMockRecorder) UpdateApprovalRule(pid, mergeRequest, approvalRule, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, approvalRule, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApprovalRule", reflect.TypeOf((*MockMergeRequestApprovalsService)(nil).UpdateApprovalRule), varargs...)
}
