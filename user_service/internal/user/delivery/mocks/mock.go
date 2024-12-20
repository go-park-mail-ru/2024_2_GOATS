// Code generated by MockGen. DO NOT EDIT.
// Source: delivery.go

// Package mock_delivery is a generated GoMock package.
package mock_delivery

import (
	context "context"
	reflect "reflect"

	dto "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockUserServiceInterface is a mock of UserServiceInterface interface.
type MockUserServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceInterfaceMockRecorder
}

// MockUserServiceInterfaceMockRecorder is the mock recorder for MockUserServiceInterface.
type MockUserServiceInterfaceMockRecorder struct {
	mock *MockUserServiceInterface
}

// NewMockUserServiceInterface creates a new mock instance.
func NewMockUserServiceInterface(ctrl *gomock.Controller) *MockUserServiceInterface {
	mock := &MockUserServiceInterface{ctrl: ctrl}
	mock.recorder = &MockUserServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceInterface) EXPECT() *MockUserServiceInterfaceMockRecorder {
	return m.recorder
}

// CheckFavorite mocks base method.
func (m *MockUserServiceInterface) CheckFavorite(ctx context.Context, favData *dto.Favorite) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckFavorite", ctx, favData)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckFavorite indicates an expected call of CheckFavorite.
func (mr *MockUserServiceInterfaceMockRecorder) CheckFavorite(ctx, favData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckFavorite", reflect.TypeOf((*MockUserServiceInterface)(nil).CheckFavorite), ctx, favData)
}

// Create mocks base method.
func (m *MockUserServiceInterface) Create(ctx context.Context, createData *dto.CreateUserData) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, createData)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserServiceInterfaceMockRecorder) Create(ctx, createData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserServiceInterface)(nil).Create), ctx, createData)
}

// CreateSubscription mocks base method.
func (m *MockUserServiceInterface) CreateSubscription(ctx context.Context, createData *dto.CreateSubscriptionData) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscription", ctx, createData)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubscription indicates an expected call of CreateSubscription.
func (mr *MockUserServiceInterfaceMockRecorder) CreateSubscription(ctx, createData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscription", reflect.TypeOf((*MockUserServiceInterface)(nil).CreateSubscription), ctx, createData)
}

// FindByEmail mocks base method.
func (m *MockUserServiceInterface) FindByEmail(ctx context.Context, email string) (*dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", ctx, email)
	ret0, _ := ret[0].(*dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserServiceInterfaceMockRecorder) FindByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserServiceInterface)(nil).FindByEmail), ctx, email)
}

// FindByID mocks base method.
func (m *MockUserServiceInterface) FindByID(ctx context.Context, usrID uint64) (*dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, usrID)
	ret0, _ := ret[0].(*dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockUserServiceInterfaceMockRecorder) FindByID(ctx, usrID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserServiceInterface)(nil).FindByID), ctx, usrID)
}

// GetFavorites mocks base method.
func (m *MockUserServiceInterface) GetFavorites(ctx context.Context, usrID uint64) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavorites", ctx, usrID)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavorites indicates an expected call of GetFavorites.
func (mr *MockUserServiceInterfaceMockRecorder) GetFavorites(ctx, usrID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavorites", reflect.TypeOf((*MockUserServiceInterface)(nil).GetFavorites), ctx, usrID)
}

// ResetFavorite mocks base method.
func (m *MockUserServiceInterface) ResetFavorite(ctx context.Context, favData *dto.Favorite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetFavorite", ctx, favData)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetFavorite indicates an expected call of ResetFavorite.
func (mr *MockUserServiceInterfaceMockRecorder) ResetFavorite(ctx, favData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetFavorite", reflect.TypeOf((*MockUserServiceInterface)(nil).ResetFavorite), ctx, favData)
}

// SetFavorite mocks base method.
func (m *MockUserServiceInterface) SetFavorite(ctx context.Context, favData *dto.Favorite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFavorite", ctx, favData)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFavorite indicates an expected call of SetFavorite.
func (mr *MockUserServiceInterfaceMockRecorder) SetFavorite(ctx, favData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFavorite", reflect.TypeOf((*MockUserServiceInterface)(nil).SetFavorite), ctx, favData)
}

// UpdatePassword mocks base method.
func (m *MockUserServiceInterface) UpdatePassword(ctx context.Context, passwordData *dto.PasswordData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", ctx, passwordData)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockUserServiceInterfaceMockRecorder) UpdatePassword(ctx, passwordData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserServiceInterface)(nil).UpdatePassword), ctx, passwordData)
}

// UpdateProfile mocks base method.
func (m *MockUserServiceInterface) UpdateProfile(ctx context.Context, usrData *dto.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", ctx, usrData)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockUserServiceInterfaceMockRecorder) UpdateProfile(ctx, usrData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUserServiceInterface)(nil).UpdateProfile), ctx, usrData)
}

// UpdateSubscribtionStatus mocks base method.
func (m *MockUserServiceInterface) UpdateSubscribtionStatus(ctx context.Context, subID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubscribtionStatus", ctx, subID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSubscribtionStatus indicates an expected call of UpdateSubscribtionStatus.
func (mr *MockUserServiceInterfaceMockRecorder) UpdateSubscribtionStatus(ctx, subID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubscribtionStatus", reflect.TypeOf((*MockUserServiceInterface)(nil).UpdateSubscribtionStatus), ctx, subID)
}
