// Code generated by MockGen. DO NOT EDIT.
// Source: implementation.go

// Package mock_delivery is a generated GoMock package.
package mock_delivery

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	gomock "github.com/golang/mock/gomock"
)

// MockMovieServiceInterface is a mock of MovieServiceInterface interface.
type MockMovieServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMovieServiceInterfaceMockRecorder
}

// MockMovieServiceInterfaceMockRecorder is the mock recorder for MockMovieServiceInterface.
type MockMovieServiceInterfaceMockRecorder struct {
	mock *MockMovieServiceInterface
}

// NewMockMovieServiceInterface creates a new mock instance.
func NewMockMovieServiceInterface(ctrl *gomock.Controller) *MockMovieServiceInterface {
	mock := &MockMovieServiceInterface{ctrl: ctrl}
	mock.recorder = &MockMovieServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieServiceInterface) EXPECT() *MockMovieServiceInterfaceMockRecorder {
	return m.recorder
}

// GetCollection mocks base method.
func (m *MockMovieServiceInterface) GetCollection(ctx context.Context) (*models.CollectionsResponse, *models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollection", ctx)
	ret0, _ := ret[0].(*models.CollectionsResponse)
	ret1, _ := ret[1].(*models.ErrorResponse)
	return ret0, ret1
}

// GetCollection indicates an expected call of GetCollection.
func (mr *MockMovieServiceInterfaceMockRecorder) GetCollection(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollection", reflect.TypeOf((*MockMovieServiceInterface)(nil).GetCollection), ctx)
}
