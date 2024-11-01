// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=service_mock.go -package=service
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	errors "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	models "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	models0 "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/model"
	gomock "go.uber.org/mock/gomock"
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

// GetActor mocks base method.
func (m *MockMovieServiceInterface) GetActor(ctx context.Context, actorId int) (*models.StaffInfo, *models.ErrorRespData) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActor", ctx, actorId)
	ret0, _ := ret[0].(*models.StaffInfo)
	ret1, _ := ret[1].(*models.ErrorRespData)
	return ret0, ret1
}

// GetActor indicates an expected call of GetActor.
func (mr *MockMovieServiceInterfaceMockRecorder) GetActor(ctx, actorId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActor", reflect.TypeOf((*MockMovieServiceInterface)(nil).GetActor), ctx, actorId)
}

// GetCollection mocks base method.
func (m *MockMovieServiceInterface) GetCollection(ctx context.Context) (*models.CollectionsRespData, *models.ErrorRespData) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollection", ctx)
	ret0, _ := ret[0].(*models.CollectionsRespData)
	ret1, _ := ret[1].(*models.ErrorRespData)
	return ret0, ret1
}

// GetCollection indicates an expected call of GetCollection.
func (mr *MockMovieServiceInterfaceMockRecorder) GetCollection(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollection", reflect.TypeOf((*MockMovieServiceInterface)(nil).GetCollection), ctx)
}

// GetMovie mocks base method.
func (m *MockMovieServiceInterface) GetMovie(ctx context.Context, mvId int) (*models.MovieInfo, *models.ErrorRespData) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovie", ctx, mvId)
	ret0, _ := ret[0].(*models.MovieInfo)
	ret1, _ := ret[1].(*models.ErrorRespData)
	return ret0, ret1
}

// GetMovie indicates an expected call of GetMovie.
func (mr *MockMovieServiceInterfaceMockRecorder) GetMovie(ctx, mvId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovie", reflect.TypeOf((*MockMovieServiceInterface)(nil).GetMovie), ctx, mvId)
}

// MockRoomRepositoryInterface is a mock of RoomRepositoryInterface interface.
type MockRoomRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRoomRepositoryInterfaceMockRecorder
}

// MockRoomRepositoryInterfaceMockRecorder is the mock recorder for MockRoomRepositoryInterface.
type MockRoomRepositoryInterfaceMockRecorder struct {
	mock *MockRoomRepositoryInterface
}

// NewMockRoomRepositoryInterface creates a new mock instance.
func NewMockRoomRepositoryInterface(ctrl *gomock.Controller) *MockRoomRepositoryInterface {
	mock := &MockRoomRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRoomRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomRepositoryInterface) EXPECT() *MockRoomRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateRoom mocks base method.
func (m *MockRoomRepositoryInterface) CreateRoom(ctx context.Context, room *models0.RoomState) (*models0.RoomState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", ctx, room)
	ret0, _ := ret[0].(*models0.RoomState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockRoomRepositoryInterfaceMockRecorder) CreateRoom(ctx, room any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockRoomRepositoryInterface)(nil).CreateRoom), ctx, room)
}

// GetFromCookie mocks base method.
func (m *MockRoomRepositoryInterface) GetFromCookie(ctx context.Context, cookie string) (string, *errors.ErrorObj, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromCookie", ctx, cookie)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*errors.ErrorObj)
	ret2, _ := ret[2].(int)
	return ret0, ret1, ret2
}

// GetFromCookie indicates an expected call of GetFromCookie.
func (mr *MockRoomRepositoryInterfaceMockRecorder) GetFromCookie(ctx, cookie any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromCookie", reflect.TypeOf((*MockRoomRepositoryInterface)(nil).GetFromCookie), ctx, cookie)
}

// GetRoomState mocks base method.
func (m *MockRoomRepositoryInterface) GetRoomState(ctx context.Context, roomID string) (*models0.RoomState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomState", ctx, roomID)
	ret0, _ := ret[0].(*models0.RoomState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomState indicates an expected call of GetRoomState.
func (mr *MockRoomRepositoryInterfaceMockRecorder) GetRoomState(ctx, roomID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomState", reflect.TypeOf((*MockRoomRepositoryInterface)(nil).GetRoomState), ctx, roomID)
}

// UpdateRoomState mocks base method.
func (m *MockRoomRepositoryInterface) UpdateRoomState(ctx context.Context, roomID string, state *models0.RoomState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoomState", ctx, roomID, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRoomState indicates an expected call of UpdateRoomState.
func (mr *MockRoomRepositoryInterfaceMockRecorder) UpdateRoomState(ctx, roomID, state any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoomState", reflect.TypeOf((*MockRoomRepositoryInterface)(nil).UpdateRoomState), ctx, roomID, state)
}

// UserById mocks base method.
func (m *MockRoomRepositoryInterface) UserById(ctx context.Context, userId string) (*models0.User, *errors.ErrorObj, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserById", ctx, userId)
	ret0, _ := ret[0].(*models0.User)
	ret1, _ := ret[1].(*errors.ErrorObj)
	ret2, _ := ret[2].(int)
	return ret0, ret1, ret2
}

// UserById indicates an expected call of UserById.
func (mr *MockRoomRepositoryInterfaceMockRecorder) UserById(ctx, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserById", reflect.TypeOf((*MockRoomRepositoryInterface)(nil).UserById), ctx, userId)
}
