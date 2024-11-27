// Code generated by MockGen. DO NOT EDIT.
// Source: delivery.go

// Package mock_delivery is a generated GoMock package.
package mock_delivery

import (
	context "context"
	reflect "reflect"

	errors "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	models "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthServiceInterface is a mock of AuthServiceInterface interface.
type MockAuthServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceInterfaceMockRecorder
}

// MockAuthServiceInterfaceMockRecorder is the mock recorder for MockAuthServiceInterface.
type MockAuthServiceInterfaceMockRecorder struct {
	mock *MockAuthServiceInterface
}

// NewMockAuthServiceInterface creates a new mock instance.
func NewMockAuthServiceInterface(ctrl *gomock.Controller) *MockAuthServiceInterface {
	mock := &MockAuthServiceInterface{ctrl: ctrl}
	mock.recorder = &MockAuthServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServiceInterface) EXPECT() *MockAuthServiceInterfaceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthServiceInterface) Login(ctx context.Context, loginData *models.LoginData) (*models.AuthRespData, *errors.ServiceError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, loginData)
	ret0, _ := ret[0].(*models.AuthRespData)
	ret1, _ := ret[1].(*errors.ServiceError)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthServiceInterfaceMockRecorder) Login(ctx, loginData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthServiceInterface)(nil).Login), ctx, loginData)
}

// Logout mocks base method.
func (m *MockAuthServiceInterface) Logout(ctx context.Context, cookie string) *errors.ServiceError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", ctx, cookie)
	ret0, _ := ret[0].(*errors.ServiceError)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockAuthServiceInterfaceMockRecorder) Logout(ctx, cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthServiceInterface)(nil).Logout), ctx, cookie)
}

// Register mocks base method.
func (m *MockAuthServiceInterface) Register(ctx context.Context, registerData *models.RegisterData) (*models.AuthRespData, *errors.ServiceError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, registerData)
	ret0, _ := ret[0].(*models.AuthRespData)
	ret1, _ := ret[1].(*errors.ServiceError)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAuthServiceInterfaceMockRecorder) Register(ctx, registerData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthServiceInterface)(nil).Register), ctx, registerData)
}

// Session mocks base method.
func (m *MockAuthServiceInterface) Session(ctx context.Context, cookie string) (*models.SessionRespData, *errors.ServiceError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Session", ctx, cookie)
	ret0, _ := ret[0].(*models.SessionRespData)
	ret1, _ := ret[1].(*errors.ServiceError)
	return ret0, ret1
}

// Session indicates an expected call of Session.
func (mr *MockAuthServiceInterfaceMockRecorder) Session(ctx, cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Session", reflect.TypeOf((*MockAuthServiceInterface)(nil).Session), ctx, cookie)
}
