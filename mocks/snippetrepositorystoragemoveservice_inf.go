// Code generated by MockGen. DO NOT EDIT.
// Source: snippetrepositorystoragemoveservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=snippetrepositorystoragemoveservice_inf.go -destination=mocks/snippetrepositorystoragemoveservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockSnippetRepositoryStorageMoveService is a mock of SnippetRepositoryStorageMoveService interface.
type MockSnippetRepositoryStorageMoveService struct {
	ctrl     *gomock.Controller
	recorder *MockSnippetRepositoryStorageMoveServiceMockRecorder
}

// MockSnippetRepositoryStorageMoveServiceMockRecorder is the mock recorder for MockSnippetRepositoryStorageMoveService.
type MockSnippetRepositoryStorageMoveServiceMockRecorder struct {
	mock *MockSnippetRepositoryStorageMoveService
}

// NewMockSnippetRepositoryStorageMoveService creates a new mock instance.
func NewMockSnippetRepositoryStorageMoveService(ctrl *gomock.Controller) *MockSnippetRepositoryStorageMoveService {
	mock := &MockSnippetRepositoryStorageMoveService{ctrl: ctrl}
	mock.recorder = &MockSnippetRepositoryStorageMoveServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSnippetRepositoryStorageMoveService) EXPECT() *MockSnippetRepositoryStorageMoveServiceMockRecorder {
	return m.recorder
}

// GetStorageMove mocks base method.
func (m *MockSnippetRepositoryStorageMoveService) GetStorageMove(repositoryStorage int, options ...gitlab.RequestOptionFunc) (*gitlab.SnippetRepositoryStorageMove, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{repositoryStorage}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStorageMove", varargs...)
	ret0, _ := ret[0].(*gitlab.SnippetRepositoryStorageMove)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStorageMove indicates an expected call of GetStorageMove.
func (mr *MockSnippetRepositoryStorageMoveServiceMockRecorder) GetStorageMove(repositoryStorage any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{repositoryStorage}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageMove", reflect.TypeOf((*MockSnippetRepositoryStorageMoveService)(nil).GetStorageMove), varargs...)
}

// GetStorageMoveForSnippet mocks base method.
func (m *MockSnippetRepositoryStorageMoveService) GetStorageMoveForSnippet(snippet, repositoryStorage int, options ...gitlab.RequestOptionFunc) (*gitlab.SnippetRepositoryStorageMove, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{snippet, repositoryStorage}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStorageMoveForSnippet", varargs...)
	ret0, _ := ret[0].(*gitlab.SnippetRepositoryStorageMove)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStorageMoveForSnippet indicates an expected call of GetStorageMoveForSnippet.
func (mr *MockSnippetRepositoryStorageMoveServiceMockRecorder) GetStorageMoveForSnippet(snippet, repositoryStorage any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{snippet, repositoryStorage}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageMoveForSnippet", reflect.TypeOf((*MockSnippetRepositoryStorageMoveService)(nil).GetStorageMoveForSnippet), varargs...)
}

// RetrieveAllStorageMoves mocks base method.
func (m *MockSnippetRepositoryStorageMoveService) RetrieveAllStorageMoves(opts gitlab.RetrieveAllSnippetStorageMovesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.SnippetRepositoryStorageMove, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveAllStorageMoves", varargs...)
	ret0, _ := ret[0].([]*gitlab.SnippetRepositoryStorageMove)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RetrieveAllStorageMoves indicates an expected call of RetrieveAllStorageMoves.
func (mr *MockSnippetRepositoryStorageMoveServiceMockRecorder) RetrieveAllStorageMoves(opts any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opts}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveAllStorageMoves", reflect.TypeOf((*MockSnippetRepositoryStorageMoveService)(nil).RetrieveAllStorageMoves), varargs...)
}

// RetrieveAllStorageMovesForSnippet mocks base method.
func (m *MockSnippetRepositoryStorageMoveService) RetrieveAllStorageMovesForSnippet(snippet int, opts gitlab.RetrieveAllSnippetStorageMovesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.SnippetRepositoryStorageMove, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{snippet, opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveAllStorageMovesForSnippet", varargs...)
	ret0, _ := ret[0].([]*gitlab.SnippetRepositoryStorageMove)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RetrieveAllStorageMovesForSnippet indicates an expected call of RetrieveAllStorageMovesForSnippet.
func (mr *MockSnippetRepositoryStorageMoveServiceMockRecorder) RetrieveAllStorageMovesForSnippet(snippet, opts any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{snippet, opts}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveAllStorageMovesForSnippet", reflect.TypeOf((*MockSnippetRepositoryStorageMoveService)(nil).RetrieveAllStorageMovesForSnippet), varargs...)
}

// ScheduleAllStorageMoves mocks base method.
func (m *MockSnippetRepositoryStorageMoveService) ScheduleAllStorageMoves(opts gitlab.ScheduleAllSnippetStorageMovesOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScheduleAllStorageMoves", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScheduleAllStorageMoves indicates an expected call of ScheduleAllStorageMoves.
func (mr *MockSnippetRepositoryStorageMoveServiceMockRecorder) ScheduleAllStorageMoves(opts any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opts}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleAllStorageMoves", reflect.TypeOf((*MockSnippetRepositoryStorageMoveService)(nil).ScheduleAllStorageMoves), varargs...)
}

// ScheduleStorageMoveForSnippet mocks base method.
func (m *MockSnippetRepositoryStorageMoveService) ScheduleStorageMoveForSnippet(snippet int, opts gitlab.ScheduleStorageMoveForSnippetOptions, options ...gitlab.RequestOptionFunc) (*gitlab.SnippetRepositoryStorageMove, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{snippet, opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScheduleStorageMoveForSnippet", varargs...)
	ret0, _ := ret[0].(*gitlab.SnippetRepositoryStorageMove)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ScheduleStorageMoveForSnippet indicates an expected call of ScheduleStorageMoveForSnippet.
func (mr *MockSnippetRepositoryStorageMoveServiceMockRecorder) ScheduleStorageMoveForSnippet(snippet, opts any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{snippet, opts}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleStorageMoveForSnippet", reflect.TypeOf((*MockSnippetRepositoryStorageMoveService)(nil).ScheduleStorageMoveForSnippet), varargs...)
}
