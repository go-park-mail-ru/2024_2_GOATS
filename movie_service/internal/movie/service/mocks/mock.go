// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	gomock "github.com/golang/mock/gomock"
)

// MockMovieRepositoryInterface is a mock of MovieRepositoryInterface interface.
type MockMovieRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMovieRepositoryInterfaceMockRecorder
}

// MockMovieRepositoryInterfaceMockRecorder is the mock recorder for MockMovieRepositoryInterface.
type MockMovieRepositoryInterfaceMockRecorder struct {
	mock *MockMovieRepositoryInterface
}

// NewMockMovieRepositoryInterface creates a new mock instance.
func NewMockMovieRepositoryInterface(ctrl *gomock.Controller) *MockMovieRepositoryInterface {
	mock := &MockMovieRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockMovieRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieRepositoryInterface) EXPECT() *MockMovieRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetActor mocks base method.
func (m *MockMovieRepositoryInterface) GetActor(ctx context.Context, actorID int) (*models.ActorInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActor", ctx, actorID)
	ret0, _ := ret[0].(*models.ActorInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActor indicates an expected call of GetActor.
func (mr *MockMovieRepositoryInterfaceMockRecorder) GetActor(ctx, actorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActor", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).GetActor), ctx, actorID)
}

// GetCollection mocks base method.
func (m *MockMovieRepositoryInterface) GetCollection(ctx context.Context, filter string) ([]models.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollection", ctx, filter)
	ret0, _ := ret[0].([]models.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollection indicates an expected call of GetCollection.
func (mr *MockMovieRepositoryInterfaceMockRecorder) GetCollection(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollection", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).GetCollection), ctx, filter)
}

// GetMovie mocks base method.
func (m *MockMovieRepositoryInterface) GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovie", ctx, mvID)
	ret0, _ := ret[0].(*models.MovieInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovie indicates an expected call of GetMovie.
func (mr *MockMovieRepositoryInterfaceMockRecorder) GetMovie(ctx, mvID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovie", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).GetMovie), ctx, mvID)
}

// GetMovieActors mocks base method.
func (m *MockMovieRepositoryInterface) GetMovieActors(ctx context.Context, mvID int) ([]*models.ActorInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieActors", ctx, mvID)
	ret0, _ := ret[0].([]*models.ActorInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieActors indicates an expected call of GetMovieActors.
func (mr *MockMovieRepositoryInterfaceMockRecorder) GetMovieActors(ctx, mvID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieActors", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).GetMovieActors), ctx, mvID)
}

// GetMovieByGenre mocks base method.
func (m *MockMovieRepositoryInterface) GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieByGenre", ctx, genre)
	ret0, _ := ret[0].([]models.MovieShortInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieByGenre indicates an expected call of GetMovieByGenre.
func (mr *MockMovieRepositoryInterfaceMockRecorder) GetMovieByGenre(ctx, genre interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieByGenre", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).GetMovieByGenre), ctx, genre)
}

// SearchActors mocks base method.
func (m *MockMovieRepositoryInterface) SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchActors", ctx, query)
	ret0, _ := ret[0].([]models.ActorInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchActors indicates an expected call of SearchActors.
func (mr *MockMovieRepositoryInterfaceMockRecorder) SearchActors(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchActors", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).SearchActors), ctx, query)
}

// SearchMovies mocks base method.
func (m *MockMovieRepositoryInterface) SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchMovies", ctx, query)
	ret0, _ := ret[0].([]models.MovieInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMovies indicates an expected call of SearchMovies.
func (mr *MockMovieRepositoryInterfaceMockRecorder) SearchMovies(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMovies", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).SearchMovies), ctx, query)
}
