// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	dto "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/repository/dto"
	dto0 "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthRepositoryInterface is a mock of AuthRepositoryInterface interface.
type MockAuthRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepositoryInterfaceMockRecorder
}

// MockAuthRepositoryInterfaceMockRecorder is the mock recorder for MockAuthRepositoryInterface.
type MockAuthRepositoryInterfaceMockRecorder struct {
	mock *MockAuthRepositoryInterface
}

// NewMockAuthRepositoryInterface creates a new mock instance.
func NewMockAuthRepositoryInterface(ctrl *gomock.Controller) *MockAuthRepositoryInterface {
	mock := &MockAuthRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockAuthRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepositoryInterface) EXPECT() *MockAuthRepositoryInterfaceMockRecorder {
	return m.recorder
}

// DestroySession mocks base method.
func (m *MockAuthRepositoryInterface) DestroySession(ctx context.Context, cookie string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroySession", ctx, cookie)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroySession indicates an expected call of DestroySession.
func (mr *MockAuthRepositoryInterfaceMockRecorder) DestroySession(ctx, cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroySession", reflect.TypeOf((*MockAuthRepositoryInterface)(nil).DestroySession), ctx, cookie)
}

// GetSessionData mocks base method.
func (m *MockAuthRepositoryInterface) GetSessionData(ctx context.Context, cookie string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionData", ctx, cookie)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionData indicates an expected call of GetSessionData.
func (mr *MockAuthRepositoryInterfaceMockRecorder) GetSessionData(ctx, cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionData", reflect.TypeOf((*MockAuthRepositoryInterface)(nil).GetSessionData), ctx, cookie)
}

// SetCookie mocks base method.
func (m *MockAuthRepositoryInterface) SetCookie(ctx context.Context, token *dto.TokenData) (*dto0.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCookie", ctx, token)
	ret0, _ := ret[0].(*dto0.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetCookie indicates an expected call of SetCookie.
func (mr *MockAuthRepositoryInterfaceMockRecorder) SetCookie(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCookie", reflect.TypeOf((*MockAuthRepositoryInterface)(nil).SetCookie), ctx, token)
}