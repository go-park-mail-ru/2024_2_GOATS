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

// MockPaymentServiceInterface is a mock of PaymentServiceInterface interface.
type MockPaymentServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentServiceInterfaceMockRecorder
}

// MockPaymentServiceInterfaceMockRecorder is the mock recorder for MockPaymentServiceInterface.
type MockPaymentServiceInterfaceMockRecorder struct {
	mock *MockPaymentServiceInterface
}

// NewMockPaymentServiceInterface creates a new mock instance.
func NewMockPaymentServiceInterface(ctrl *gomock.Controller) *MockPaymentServiceInterface {
	mock := &MockPaymentServiceInterface{ctrl: ctrl}
	mock.recorder = &MockPaymentServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentServiceInterface) EXPECT() *MockPaymentServiceInterfaceMockRecorder {
	return m.recorder
}

// ProcessCallback mocks base method.
func (m *MockPaymentServiceInterface) ProcessCallback(ctx context.Context, data *models.PaymentCallbackData) *errors.ServiceError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessCallback", ctx, data)
	ret0, _ := ret[0].(*errors.ServiceError)
	return ret0
}

// ProcessCallback indicates an expected call of ProcessCallback.
func (mr *MockPaymentServiceInterfaceMockRecorder) ProcessCallback(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessCallback", reflect.TypeOf((*MockPaymentServiceInterface)(nil).ProcessCallback), ctx, data)
}
