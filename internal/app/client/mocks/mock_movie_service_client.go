// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2024_2_GOATS/movie_service/pkg/movie_v1 (interfaces: MovieServiceClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	__ "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/pkg/movie_v1"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockMovieServiceClient is a mock of MovieServiceClient interface.
type MockMovieServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockMovieServiceClientMockRecorder
}

// MockMovieServiceClientMockRecorder is the mock recorder for MockMovieServiceClient.
type MockMovieServiceClientMockRecorder struct {
	mock *MockMovieServiceClient
}

// NewMockMovieServiceClient creates a new mock instance.
func NewMockMovieServiceClient(ctrl *gomock.Controller) *MockMovieServiceClient {
	mock := &MockMovieServiceClient{ctrl: ctrl}
	mock.recorder = &MockMovieServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieServiceClient) EXPECT() *MockMovieServiceClientMockRecorder {
	return m.recorder
}

// AddOrUpdateRating mocks base method.
func (m *MockMovieServiceClient) AddOrUpdateRating(arg0 context.Context, arg1 *__.AddOrUpdateRatingRequest, arg2 ...grpc.CallOption) (*__.AddOrUpdateRatingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddOrUpdateRating", varargs...)
	ret0, _ := ret[0].(*__.AddOrUpdateRatingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrUpdateRating indicates an expected call of AddOrUpdateRating.
func (mr *MockMovieServiceClientMockRecorder) AddOrUpdateRating(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrUpdateRating", reflect.TypeOf((*MockMovieServiceClient)(nil).AddOrUpdateRating), varargs...)
}

// DeleteRating mocks base method.
func (m *MockMovieServiceClient) DeleteRating(arg0 context.Context, arg1 *__.DeleteRatingRequest, arg2 ...grpc.CallOption) (*__.DeleteRatingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteRating", varargs...)
	ret0, _ := ret[0].(*__.DeleteRatingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRating indicates an expected call of DeleteRating.
func (mr *MockMovieServiceClientMockRecorder) DeleteRating(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRating", reflect.TypeOf((*MockMovieServiceClient)(nil).DeleteRating), varargs...)
}

// GetActor mocks base method.
func (m *MockMovieServiceClient) GetActor(arg0 context.Context, arg1 *__.GetActorRequest, arg2 ...grpc.CallOption) (*__.GetActorResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetActor", varargs...)
	ret0, _ := ret[0].(*__.GetActorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActor indicates an expected call of GetActor.
func (mr *MockMovieServiceClientMockRecorder) GetActor(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActor", reflect.TypeOf((*MockMovieServiceClient)(nil).GetActor), varargs...)
}

// GetCollections mocks base method.
func (m *MockMovieServiceClient) GetCollections(arg0 context.Context, arg1 *__.GetCollectionsRequest, arg2 ...grpc.CallOption) (*__.GetCollectionsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCollections", varargs...)
	ret0, _ := ret[0].(*__.GetCollectionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollections indicates an expected call of GetCollections.
func (mr *MockMovieServiceClientMockRecorder) GetCollections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollections", reflect.TypeOf((*MockMovieServiceClient)(nil).GetCollections), varargs...)
}

// GetFavorites mocks base method.
func (m *MockMovieServiceClient) GetFavorites(arg0 context.Context, arg1 *__.GetFavoritesRequest, arg2 ...grpc.CallOption) (*__.GetFavoritesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFavorites", varargs...)
	ret0, _ := ret[0].(*__.GetFavoritesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavorites indicates an expected call of GetFavorites.
func (mr *MockMovieServiceClientMockRecorder) GetFavorites(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavorites", reflect.TypeOf((*MockMovieServiceClient)(nil).GetFavorites), varargs...)
}

// GetMovie mocks base method.
func (m *MockMovieServiceClient) GetMovie(arg0 context.Context, arg1 *__.GetMovieRequest, arg2 ...grpc.CallOption) (*__.GetMovieResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMovie", varargs...)
	ret0, _ := ret[0].(*__.GetMovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovie indicates an expected call of GetMovie.
func (mr *MockMovieServiceClientMockRecorder) GetMovie(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovie", reflect.TypeOf((*MockMovieServiceClient)(nil).GetMovie), varargs...)
}

// GetMovieActors mocks base method.
func (m *MockMovieServiceClient) GetMovieActors(arg0 context.Context, arg1 *__.GetMovieActorsRequest, arg2 ...grpc.CallOption) (*__.GetMovieActorsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMovieActors", varargs...)
	ret0, _ := ret[0].(*__.GetMovieActorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieActors indicates an expected call of GetMovieActors.
func (mr *MockMovieServiceClientMockRecorder) GetMovieActors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieActors", reflect.TypeOf((*MockMovieServiceClient)(nil).GetMovieActors), varargs...)
}

// GetMovieByGenre mocks base method.
func (m *MockMovieServiceClient) GetMovieByGenre(arg0 context.Context, arg1 *__.GetMovieByGenreRequest, arg2 ...grpc.CallOption) (*__.GetMovieByGenreResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMovieByGenre", varargs...)
	ret0, _ := ret[0].(*__.GetMovieByGenreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieByGenre indicates an expected call of GetMovieByGenre.
func (mr *MockMovieServiceClientMockRecorder) GetMovieByGenre(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieByGenre", reflect.TypeOf((*MockMovieServiceClient)(nil).GetMovieByGenre), varargs...)
}

// GetUserRating mocks base method.
func (m *MockMovieServiceClient) GetUserRating(arg0 context.Context, arg1 *__.GetUserRatingRequest, arg2 ...grpc.CallOption) (*__.GetUserRatingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserRating", varargs...)
	ret0, _ := ret[0].(*__.GetUserRatingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRating indicates an expected call of GetUserRating.
func (mr *MockMovieServiceClientMockRecorder) GetUserRating(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRating", reflect.TypeOf((*MockMovieServiceClient)(nil).GetUserRating), varargs...)
}

// SearchActors mocks base method.
func (m *MockMovieServiceClient) SearchActors(arg0 context.Context, arg1 *__.SearchActorsRequest, arg2 ...grpc.CallOption) (*__.SearchActorsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchActors", varargs...)
	ret0, _ := ret[0].(*__.SearchActorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchActors indicates an expected call of SearchActors.
func (mr *MockMovieServiceClientMockRecorder) SearchActors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchActors", reflect.TypeOf((*MockMovieServiceClient)(nil).SearchActors), varargs...)
}

// SearchMovies mocks base method.
func (m *MockMovieServiceClient) SearchMovies(arg0 context.Context, arg1 *__.SearchMoviesRequest, arg2 ...grpc.CallOption) (*__.SearchMoviesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchMovies", varargs...)
	ret0, _ := ret[0].(*__.SearchMoviesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMovies indicates an expected call of SearchMovies.
func (mr *MockMovieServiceClientMockRecorder) SearchMovies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMovies", reflect.TypeOf((*MockMovieServiceClient)(nil).SearchMovies), varargs...)
}
