package service

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
// 	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
// 	servMock "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie_service/service/mocks"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestService_GetCollection(t *testing.T) {
// 	tests := []struct {
// 		name             string
// 		mockReturn       []models.Collection
// 		mockErr          *errVals.RepoError
// 		expectedResponse *models.CollectionsRespData
// 		expectedError    *errVals.ServiceError
// 	}{
// 		{
// 			name: "Success",
// 			mockReturn: []models.Collection{
// 				{ID: 1, Title: "Collection 1", Movies: []*models.MovieShortInfo{}},
// 				{ID: 2, Title: "Collection 2", Movies: []*models.MovieShortInfo{}},
// 			},
// 			mockErr: nil,
// 			expectedResponse: &models.CollectionsRespData{
// 				Collections: []models.Collection{
// 					{ID: 1, Title: "Collection 1", Movies: []*models.MovieShortInfo{}},
// 					{ID: 2, Title: "Collection 2", Movies: []*models.MovieShortInfo{}},
// 				},
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			name:             "Error",
// 			mockReturn:       nil,
// 			mockErr:          &errVals.RepoError{Code: errVals.ErrServerCode, Error: errVals.CustomError{Err: errors.New("Database fail")}},
// 			expectedResponse: nil,
// 			expectedError: &errVals.ServiceError{
// 				Code:  errVals.ErrServerCode,
// 				Error: errVals.CustomError{Err: errors.New("Database fail")},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()

// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			mr := servMock.NewMockMovieRepositoryInterface(ctrl)
// 			s := NewMovieService(mr)

// 			mr.EXPECT().GetCollection(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)

// 			response, err := s.GetCollection(context.Background(), "")

// 			if test.expectedError != nil {
// 				assert.Nil(t, response)
// 				assert.Equal(t, test.expectedError, err)
// 			} else {
// 				assert.Equal(t, test.expectedResponse, response)
// 				assert.Nil(t, err)
// 			}
// 		})
// 	}
// }

// func TestService_GetActor(t *testing.T) {
// 	tests := []struct {
// 		name             string
// 		actorID          int
// 		mockReturn       *models.ActorInfo
// 		mockErr          *errVals.RepoError
// 		expectedResponse *models.ActorInfo
// 		expectedError    *errVals.ServiceError
// 	}{
// 		{
// 			name:    "Success",
// 			actorID: 1,
// 			mockReturn: &models.ActorInfo{
// 				ID: 1,
// 				Person: models.Person{
// 					Name:    "Test",
// 					Surname: "Tester",
// 				},
// 			},
// 			mockErr: nil,
// 			expectedResponse: &models.ActorInfo{
// 				ID: 1,
// 				Person: models.Person{
// 					Name:    "Test",
// 					Surname: "Tester",
// 				},
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			name:       "Error",
// 			actorID:    1,
// 			mockReturn: nil,
// 			mockErr: &errVals.RepoError{
// 				Code:  "actor_not_found",
// 				Error: errVals.CustomError{Err: errors.New("actor not found")},
// 			},
// 			expectedResponse: nil,
// 			expectedError: &errVals.ServiceError{
// 				Code:  "actor_not_found",
// 				Error: errVals.CustomError{Err: errors.New("actor not found")},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()

// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			mr := servMock.NewMockMovieRepositoryInterface(ctrl)
// 			s := NewMovieService(mr)

// 			mr.EXPECT().GetActor(gomock.Any(), test.actorID).Return(test.mockReturn, test.mockErr)

// 			response, err := s.GetActor(context.Background(), test.actorID)

// 			if test.expectedError != nil {
// 				assert.Nil(t, response)
// 				assert.Equal(t, test.expectedError, err)
// 			} else {
// 				assert.Equal(t, test.expectedResponse, response)
// 				assert.Nil(t, err)
// 			}
// 		})
// 	}
// }

// func TestService_GetMovie(t *testing.T) {
// 	tests := []struct {
// 		name             string
// 		mvID             int
// 		mockMovieReturn  *models.MovieInfo
// 		mockMovieErr     *errVals.RepoError
// 		mockActorsReturn []*models.ActorInfo
// 		mockActorsErr    *errVals.RepoError
// 		expectedResponse *models.MovieInfo
// 		expectedError    *errVals.ServiceError
// 	}{
// 		{
// 			name: "Success",
// 			mvID: 1,
// 			mockMovieReturn: &models.MovieInfo{
// 				ID:    1,
// 				Title: "Test Movie",
// 			},
// 			mockMovieErr: nil,
// 			mockActorsReturn: []*models.ActorInfo{{
// 				ID: 1,
// 				Person: models.Person{
// 					Name:    "Test",
// 					Surname: "Tester",
// 				}}},
// 			mockActorsErr: nil,
// 			expectedResponse: &models.MovieInfo{
// 				ID:    1,
// 				Title: "Test Movie",
// 				Actors: []*models.ActorInfo{{
// 					ID: 1,
// 					Person: models.Person{
// 						Name:    "Test",
// 						Surname: "Tester",
// 					},
// 				}},
// 			},
// 			expectedError: nil,
// 		},
// 		{
// 			name:             "Error_GetMovie",
// 			mvID:             1,
// 			mockMovieReturn:  nil,
// 			mockMovieErr:     &errVals.RepoError{Code: "movie_not_found", Error: errVals.CustomError{Err: errors.New("movie_service not found")}},
// 			mockActorsReturn: nil,
// 			mockActorsErr:    nil,
// 			expectedResponse: nil,
// 			expectedError: &errVals.ServiceError{
// 				Code:  "movie_not_found",
// 				Error: errVals.CustomError{Err: errors.New("movie_service not found")},
// 			},
// 		},
// 		{
// 			name: "Error_GetMovieActors",
// 			mvID: 1,
// 			mockMovieReturn: &models.MovieInfo{
// 				ID:    1,
// 				Title: "Test Movie",
// 			},
// 			mockMovieErr:     nil,
// 			mockActorsReturn: nil,
// 			mockActorsErr:    &errVals.RepoError{Code: "actors_fetch_error", Error: errVals.CustomError{Err: errors.New("failed to fetch actors")}},
// 			expectedResponse: nil,
// 			expectedError: &errVals.ServiceError{
// 				Code:  "actors_fetch_error",
// 				Error: errVals.CustomError{Err: errors.New("failed to fetch actors")},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()

// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			mr := servMock.NewMockMovieRepositoryInterface(ctrl)
// 			s := NewMovieService(mr)

// 			mr.EXPECT().GetMovie(gomock.Any(), test.mvID).Return(test.mockMovieReturn, test.mockMovieErr)

// 			if test.mockMovieErr == nil {
// 				mr.EXPECT().GetMovieActors(gomock.Any(), test.mvID).Return(test.mockActorsReturn, test.mockActorsErr)
// 			}

// 			response, err := s.GetMovie(context.Background(), test.mvID)

// 			if test.expectedError != nil {
// 				assert.Nil(t, response)
// 				assert.Equal(t, test.expectedError, err)
// 			} else {
// 				assert.Equal(t, test.expectedResponse, response)
// 				assert.Nil(t, err)
// 			}
// 		})
// 	}
// }
