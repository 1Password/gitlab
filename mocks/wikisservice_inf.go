// Code generated by MockGen. DO NOT EDIT.
// Source: wikisservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=wikisservice_inf.go -destination=mocks/wikisservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockWikisService is a mock of WikisService interface.
type MockWikisService struct {
	ctrl     *gomock.Controller
	recorder *MockWikisServiceMockRecorder
}

// MockWikisServiceMockRecorder is the mock recorder for MockWikisService.
type MockWikisServiceMockRecorder struct {
	mock *MockWikisService
}

// NewMockWikisService creates a new mock instance.
func NewMockWikisService(ctrl *gomock.Controller) *MockWikisService {
	mock := &MockWikisService{ctrl: ctrl}
	mock.recorder = &MockWikisServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWikisService) EXPECT() *MockWikisServiceMockRecorder {
	return m.recorder
}

// CreateWikiPage mocks base method.
func (m *MockWikisService) CreateWikiPage(pid any, opt *gitlab.CreateWikiPageOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Wiki, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateWikiPage", varargs...)
	ret0, _ := ret[0].(*gitlab.Wiki)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateWikiPage indicates an expected call of CreateWikiPage.
func (mr *MockWikisServiceMockRecorder) CreateWikiPage(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWikiPage", reflect.TypeOf((*MockWikisService)(nil).CreateWikiPage), varargs...)
}

// DeleteWikiPage mocks base method.
func (m *MockWikisService) DeleteWikiPage(pid any, slug string, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, slug}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteWikiPage", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteWikiPage indicates an expected call of DeleteWikiPage.
func (mr *MockWikisServiceMockRecorder) DeleteWikiPage(pid, slug any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, slug}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWikiPage", reflect.TypeOf((*MockWikisService)(nil).DeleteWikiPage), varargs...)
}

// EditWikiPage mocks base method.
func (m *MockWikisService) EditWikiPage(pid any, slug string, opt *gitlab.EditWikiPageOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Wiki, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, slug, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditWikiPage", varargs...)
	ret0, _ := ret[0].(*gitlab.Wiki)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EditWikiPage indicates an expected call of EditWikiPage.
func (mr *MockWikisServiceMockRecorder) EditWikiPage(pid, slug, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, slug, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditWikiPage", reflect.TypeOf((*MockWikisService)(nil).EditWikiPage), varargs...)
}

// GetWikiPage mocks base method.
func (m *MockWikisService) GetWikiPage(pid any, slug string, opt *gitlab.GetWikiPageOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Wiki, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, slug, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWikiPage", varargs...)
	ret0, _ := ret[0].(*gitlab.Wiki)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetWikiPage indicates an expected call of GetWikiPage.
func (mr *MockWikisServiceMockRecorder) GetWikiPage(pid, slug, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, slug, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWikiPage", reflect.TypeOf((*MockWikisService)(nil).GetWikiPage), varargs...)
}

// ListWikis mocks base method.
func (m *MockWikisService) ListWikis(pid any, opt *gitlab.ListWikisOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Wiki, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWikis", varargs...)
	ret0, _ := ret[0].([]*gitlab.Wiki)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListWikis indicates an expected call of ListWikis.
func (mr *MockWikisServiceMockRecorder) ListWikis(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWikis", reflect.TypeOf((*MockWikisService)(nil).ListWikis), varargs...)
}
