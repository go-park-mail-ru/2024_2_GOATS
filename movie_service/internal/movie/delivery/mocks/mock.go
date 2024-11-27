// Code generated by MockGen. DO NOT EDIT.
// Source: delivery.go

// Package mock_delivery is a generated GoMock package.
package mock_delivery

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
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

// GetActor mocks base method.
func (m *MockMovieServiceInterface) GetActor(ctx context.Context, actorID int) (*models.ActorInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActor", ctx, actorID)
	ret0, _ := ret[0].(*models.ActorInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActor indicates an expected call of GetActor.
func (mr *MockMovieServiceInterfaceMockRecorder) GetActor(ctx, actorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActor", reflect.TypeOf((*MockMovieServiceInterface)(nil).GetActor), ctx, actorID)
}

// GetCollection mocks base method.
func (m *MockMovieServiceInterface) GetCollection(ctx context.Context, filter string) (*models.CollectionsRespData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollection", ctx, filter)
	ret0, _ := ret[0].(*models.CollectionsRespData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollection indicates an expected call of GetCollection.
func (mr *MockMovieServiceInterfaceMockRecorder) GetCollection(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollection", reflect.TypeOf((*MockMovieServiceInterface)(nil).GetCollection), ctx, filter)
}

// GetFavorites mocks base method.
func (m *MockMovieServiceInterface) GetFavorites(ctx context.Context, mvIDs []uint64) ([]*models.MovieShortInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavorites", ctx, mvIDs)
	ret0, _ := ret[0].([]*models.MovieShortInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavorites indicates an expected call of GetFavorites.
func (mr *MockMovieServiceInterfaceMockRecorder) GetFavorites(ctx, mvIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavorites", reflect.TypeOf((*MockMovieServiceInterface)(nil).GetFavorites), ctx, mvIDs)
}

// GetMovie mocks base method.
func (m *MockMovieServiceInterface) GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovie", ctx, mvID)
	ret0, _ := ret[0].(*models.MovieInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovie indicates an expected call of GetMovie.
func (mr *MockMovieServiceInterfaceMockRecorder) GetMovie(ctx, mvID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovie", reflect.TypeOf((*MockMovieServiceInterface)(nil).GetMovie), ctx, mvID)
}

// GetMovieActors mocks base method.
func (m *MockMovieServiceInterface) GetMovieActors(ctx context.Context, mvID int) ([]*models.ActorInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieActors", ctx, mvID)
	ret0, _ := ret[0].([]*models.ActorInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieActors indicates an expected call of GetMovieActors.
func (mr *MockMovieServiceInterfaceMockRecorder) GetMovieActors(ctx, mvID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieActors", reflect.TypeOf((*MockMovieServiceInterface)(nil).GetMovieActors), ctx, mvID)
}

// GetMovieByGenre mocks base method.
func (m *MockMovieServiceInterface) GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieByGenre", ctx, genre)
	ret0, _ := ret[0].([]models.MovieShortInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieByGenre indicates an expected call of GetMovieByGenre.
func (mr *MockMovieServiceInterfaceMockRecorder) GetMovieByGenre(ctx, genre interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieByGenre", reflect.TypeOf((*MockMovieServiceInterface)(nil).GetMovieByGenre), ctx, genre)
}

// SearchActors mocks base method.
func (m *MockMovieServiceInterface) SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchActors", ctx, query)
	ret0, _ := ret[0].([]models.ActorInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchActors indicates an expected call of SearchActors.
func (mr *MockMovieServiceInterfaceMockRecorder) SearchActors(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchActors", reflect.TypeOf((*MockMovieServiceInterface)(nil).SearchActors), ctx, query)
}

// SearchMovies mocks base method.
func (m *MockMovieServiceInterface) SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchMovies", ctx, query)
	ret0, _ := ret[0].([]models.MovieInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMovies indicates an expected call of SearchMovies.
func (mr *MockMovieServiceInterfaceMockRecorder) SearchMovies(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMovies", reflect.TypeOf((*MockMovieServiceInterface)(nil).SearchMovies), ctx, query)
}
