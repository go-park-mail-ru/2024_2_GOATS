// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	errors "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	models "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepositoryInterface is a mock of UserRepositoryInterface interface.
type MockUserRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryInterfaceMockRecorder
}

// MockUserRepositoryInterfaceMockRecorder is the mock recorder for MockUserRepositoryInterface.
type MockUserRepositoryInterfaceMockRecorder struct {
	mock *MockUserRepositoryInterface
}

// NewMockUserRepositoryInterface creates a new mock instance.
func NewMockUserRepositoryInterface(ctrl *gomock.Controller) *MockUserRepositoryInterface {
	mock := &MockUserRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepositoryInterface) EXPECT() *MockUserRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepositoryInterface) CreateUser(ctx context.Context, registerData *models.RegisterData) (*models.User, *errors.ErrorObj, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, registerData)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(*errors.ErrorObj)
	ret2, _ := ret[2].(int)
	return ret0, ret1, ret2
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryInterfaceMockRecorder) CreateUser(ctx, registerData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).CreateUser), ctx, registerData)
}

// SaveAvatar mocks base method.
func (m *MockUserRepositoryInterface) SaveAvatar(ctx context.Context, usrData *models.User) (string, *errors.ErrorObj) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAvatar", ctx, usrData)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*errors.ErrorObj)
	return ret0, ret1
}

// SaveAvatar indicates an expected call of SaveAvatar.
func (mr *MockUserRepositoryInterfaceMockRecorder) SaveAvatar(ctx, usrData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAvatar", reflect.TypeOf((*MockUserRepositoryInterface)(nil).SaveAvatar), ctx, usrData)
}

// UpdatePassword mocks base method.
func (m *MockUserRepositoryInterface) UpdatePassword(ctx context.Context, usrId int, pass string) (*errors.ErrorObj, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", ctx, usrId, pass)
	ret0, _ := ret[0].(*errors.ErrorObj)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockUserRepositoryInterfaceMockRecorder) UpdatePassword(ctx, usrId, pass interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserRepositoryInterface)(nil).UpdatePassword), ctx, usrId, pass)
}

// UpdateProfileData mocks base method.
func (m *MockUserRepositoryInterface) UpdateProfileData(ctx context.Context, usrData *models.User) (*errors.ErrorObj, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfileData", ctx, usrData)
	ret0, _ := ret[0].(*errors.ErrorObj)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// UpdateProfileData indicates an expected call of UpdateProfileData.
func (mr *MockUserRepositoryInterfaceMockRecorder) UpdateProfileData(ctx, usrData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfileData", reflect.TypeOf((*MockUserRepositoryInterface)(nil).UpdateProfileData), ctx, usrData)
}

// UserByEmail mocks base method.
func (m *MockUserRepositoryInterface) UserByEmail(ctx context.Context, email string) (*models.User, *errors.ErrorObj, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserByEmail", ctx, email)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(*errors.ErrorObj)
	ret2, _ := ret[2].(int)
	return ret0, ret1, ret2
}

// UserByEmail indicates an expected call of UserByEmail.
func (mr *MockUserRepositoryInterfaceMockRecorder) UserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserByEmail", reflect.TypeOf((*MockUserRepositoryInterface)(nil).UserByEmail), ctx, email)
}

// UserById mocks base method.
func (m *MockUserRepositoryInterface) UserById(ctx context.Context, userId int) (*models.User, *errors.ErrorObj, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserById", ctx, userId)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(*errors.ErrorObj)
	ret2, _ := ret[2].(int)
	return ret0, ret1, ret2
}

// UserById indicates an expected call of UserById.
func (mr *MockUserRepositoryInterfaceMockRecorder) UserById(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserById", reflect.TypeOf((*MockUserRepositoryInterface)(nil).UserById), ctx, userId)
}