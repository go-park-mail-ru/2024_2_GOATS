// Code generated by MockGen. DO NOT EDIT.
// Source: auth.go
//
// Generated by this command:
//
//	mockgen -source=auth.go -destination=../auth/service/mocks/mock.go
//

// Package mock_client is a generated GoMock package.
package mock_client

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthClientInterface is a mock of AuthClientInterface interface.
type MockAuthClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAuthClientInterfaceMockRecorder
}

// MockAuthClientInterfaceMockRecorder is the mock recorder for MockAuthClientInterface.
type MockAuthClientInterfaceMockRecorder struct {
	mock *MockAuthClientInterface
}

// NewMockAuthClientInterface creates a new mock instance.
func NewMockAuthClientInterface(ctrl *gomock.Controller) *MockAuthClientInterface {
	mock := &MockAuthClientInterface{ctrl: ctrl}
	mock.recorder = &MockAuthClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthClientInterface) EXPECT() *MockAuthClientInterfaceMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockAuthClientInterface) CreateSession(ctx context.Context, usrID int) (*models.CookieData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, usrID)
	ret0, _ := ret[0].(*models.CookieData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockAuthClientInterfaceMockRecorder) CreateSession(ctx, usrID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockAuthClientInterface)(nil).CreateSession), ctx, usrID)
}

// DestroySession mocks base method.
func (m *MockAuthClientInterface) DestroySession(ctx context.Context, cookie string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroySession", ctx, cookie)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroySession indicates an expected call of DestroySession.
func (mr *MockAuthClientInterfaceMockRecorder) DestroySession(ctx, cookie any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroySession", reflect.TypeOf((*MockAuthClientInterface)(nil).DestroySession), ctx, cookie)
}

// Session mocks base method.
func (m *MockAuthClientInterface) Session(ctx context.Context, cookie string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Session", ctx, cookie)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Session indicates an expected call of Session.
func (mr *MockAuthClientInterfaceMockRecorder) Session(ctx, cookie any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Session", reflect.TypeOf((*MockAuthClientInterface)(nil).Session), ctx, cookie)
}
