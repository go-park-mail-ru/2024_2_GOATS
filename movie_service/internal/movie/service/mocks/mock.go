// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=mocks/mock.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	gomock "go.uber.org/mock/gomock"
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
func (mr *MockMovieRepositoryInterfaceMockRecorder) GetActor(ctx, actorID any) *gomock.Call {
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
func (mr *MockMovieRepositoryInterfaceMockRecorder) GetCollection(ctx, filter any) *gomock.Call {
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
func (mr *MockMovieRepositoryInterfaceMockRecorder) GetMovie(ctx, mvID any) *gomock.Call {
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
func (mr *MockMovieRepositoryInterfaceMockRecorder) GetMovieActors(ctx, mvID any) *gomock.Call {
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
func (mr *MockMovieRepositoryInterfaceMockRecorder) GetMovieByGenre(ctx, genre any) *gomock.Call {
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
func (mr *MockMovieRepositoryInterfaceMockRecorder) SearchActors(ctx, query any) *gomock.Call {
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
func (mr *MockMovieRepositoryInterfaceMockRecorder) SearchMovies(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMovies", reflect.TypeOf((*MockMovieRepositoryInterface)(nil).SearchMovies), ctx, query)
}

// MockUserClientInterface is a mock of UserClientInterface interface.
type MockUserClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserClientInterfaceMockRecorder
}

// MockUserClientInterfaceMockRecorder is the mock recorder for MockUserClientInterface.
type MockUserClientInterfaceMockRecorder struct {
	mock *MockUserClientInterface
}

// NewMockUserClientInterface creates a new mock instance.
func NewMockUserClientInterface(ctrl *gomock.Controller) *MockUserClientInterface {
	mock := &MockUserClientInterface{ctrl: ctrl}
	mock.recorder = &MockUserClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserClientInterface) EXPECT() *MockUserClientInterfaceMockRecorder {
	return m.recorder
}

// CheckFavorite mocks base method.
func (m *MockUserClientInterface) CheckFavorite(ctx context.Context, favData *models.Favorite) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckFavorite", ctx, favData)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckFavorite indicates an expected call of CheckFavorite.
func (mr *MockUserClientInterfaceMockRecorder) CheckFavorite(ctx, favData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckFavorite", reflect.TypeOf((*MockUserClientInterface)(nil).CheckFavorite), ctx, favData)
}

// Create mocks base method.
func (m *MockUserClientInterface) Create(ctx context.Context, regData *models.RegisterData) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, regData)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserClientInterfaceMockRecorder) Create(ctx, regData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserClientInterface)(nil).Create), ctx, regData)
}

// FindByEmail mocks base method.
func (m *MockUserClientInterface) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", ctx, email)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserClientInterfaceMockRecorder) FindByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserClientInterface)(nil).FindByEmail), ctx, email)
}

// FindByID mocks base method.
func (m *MockUserClientInterface) FindByID(ctx context.Context, id uint64) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockUserClientInterfaceMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserClientInterface)(nil).FindByID), ctx, id)
}

// GetFavorites mocks base method.
func (m *MockUserClientInterface) GetFavorites(ctx context.Context, usrID int) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavorites", ctx, usrID)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavorites indicates an expected call of GetFavorites.
func (mr *MockUserClientInterfaceMockRecorder) GetFavorites(ctx, usrID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavorites", reflect.TypeOf((*MockUserClientInterface)(nil).GetFavorites), ctx, usrID)
}

// ResetFavorite mocks base method.
func (m *MockUserClientInterface) ResetFavorite(ctx context.Context, favData *models.Favorite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetFavorite", ctx, favData)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetFavorite indicates an expected call of ResetFavorite.
func (mr *MockUserClientInterfaceMockRecorder) ResetFavorite(ctx, favData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetFavorite", reflect.TypeOf((*MockUserClientInterface)(nil).ResetFavorite), ctx, favData)
}

// SetFavorite mocks base method.
func (m *MockUserClientInterface) SetFavorite(ctx context.Context, favData *models.Favorite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFavorite", ctx, favData)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFavorite indicates an expected call of SetFavorite.
func (mr *MockUserClientInterfaceMockRecorder) SetFavorite(ctx, favData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFavorite", reflect.TypeOf((*MockUserClientInterface)(nil).SetFavorite), ctx, favData)
}

// UpdatePassword mocks base method.
func (m *MockUserClientInterface) UpdatePassword(ctx context.Context, passwordData *models.PasswordData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", ctx, passwordData)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockUserClientInterfaceMockRecorder) UpdatePassword(ctx, passwordData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserClientInterface)(nil).UpdatePassword), ctx, passwordData)
}

// UpdateProfile mocks base method.
func (m *MockUserClientInterface) UpdateProfile(ctx context.Context, usrData *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", ctx, usrData)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockUserClientInterfaceMockRecorder) UpdateProfile(ctx, usrData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUserClientInterface)(nil).UpdateProfile), ctx, usrData)
}
