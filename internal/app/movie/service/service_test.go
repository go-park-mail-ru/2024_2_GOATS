package service

import (
	"context"
	"errors"
	"net/http"
	"testing"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	servMock "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_GetCollection(t *testing.T) {
	tests := []struct {
		name             string
		mockReturn       []models.Collection
		mockErr          *errVals.ErrorObj
		expectedResponse *models.CollectionsRespData
		expectedError    *models.ErrorRespData
		statusCode       int
	}{
		{
			name: "Success",
			mockReturn: []models.Collection{
				{Id: 1, Title: "Collection 1", Movies: []*models.MovieShortInfo{}},
				{Id: 2, Title: "Collection 2", Movies: []*models.MovieShortInfo{}},
			},
			mockErr: nil,
			expectedResponse: &models.CollectionsRespData{
				Collections: []models.Collection{
					{Id: 1, Title: "Collection 1", Movies: []*models.MovieShortInfo{}},
					{Id: 2, Title: "Collection 2", Movies: []*models.MovieShortInfo{}},
				},
				StatusCode: http.StatusOK,
			},
			expectedError: nil,
			statusCode:    http.StatusOK,
		},
		{
			name:             "Error",
			mockReturn:       nil,
			mockErr:          &errVals.ErrorObj{Code: errVals.ErrServerCode, Error: errVals.CustomError{Err: errors.New("Database fail")}},
			expectedResponse: nil,
			expectedError: &models.ErrorRespData{
				StatusCode: http.StatusUnprocessableEntity,
				Errors:     []errVals.ErrorObj{{Code: errVals.ErrServerCode, Error: errVals.CustomError{Err: errors.New("Database fail")}}},
			},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := servMock.NewMockMovieRepositoryInterface(ctrl)
			s := NewService(repo)

			repo.EXPECT().GetCollection(gomock.Any()).Return(test.mockReturn, test.mockErr, test.statusCode)

			t.Parallel()

			response, err := s.GetCollection(context.Background())

			if test.expectedError != nil {
				assert.Nil(t, response)
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResponse, response)
				assert.Nil(t, err)
			}
		})
	}
}

func TestService_GetActor(t *testing.T) {
	tests := []struct {
		name             string
		actorId          int
		mockReturn       *models.ActorInfo
		mockErr          *errVals.ErrorObj
		expectedResponse *models.ActorInfo
		expectedError    *models.ErrorRespData
		statusCode       int
	}{
		{
			name:    "Success",
			actorId: 1,
			mockReturn: &models.ActorInfo{
				Id: 1,
				Person: models.Person{
					Name:    "Test",
					Surname: "Tester",
				},
			},
			mockErr: nil,
			expectedResponse: &models.ActorInfo{
				Id: 1,
				Person: models.Person{
					Name:    "Test",
					Surname: "Tester",
				},
			},
			expectedError: nil,
			statusCode:    http.StatusOK,
		},
		{
			name:       "Error",
			actorId:    1,
			mockReturn: nil,
			mockErr: &errVals.ErrorObj{
				Code:  "actor_not_found",
				Error: errVals.CustomError{Err: errors.New("actor not found")},
			},
			expectedResponse: nil,
			expectedError: &models.ErrorRespData{
				StatusCode: http.StatusNotFound,
				Errors: []errVals.ErrorObj{
					{Code: "actor_not_found", Error: errVals.CustomError{Err: errors.New("actor not found")}},
				},
			},
			statusCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := servMock.NewMockMovieRepositoryInterface(ctrl)
			s := NewService(repo)

			repo.EXPECT().GetActor(gomock.Any(), test.actorId).Return(test.mockReturn, test.mockErr, test.statusCode)

			response, err := s.GetActor(context.Background(), test.actorId)

			if test.expectedError != nil {
				assert.Nil(t, response)
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResponse, response)
				assert.Nil(t, err)
			}
		})
	}
}

func TestService_GetMovie(t *testing.T) {
	tests := []struct {
		name             string
		mvId             int
		mockMovieReturn  *models.MovieInfo
		mockMovieErr     *errVals.ErrorObj
		mockActorsReturn []*models.ActorInfo
		mockActorsErr    *errVals.ErrorObj
		expectedResponse *models.MovieInfo
		expectedError    *models.ErrorRespData
		statusCode       int
	}{
		{
			name: "Success",
			mvId: 1,
			mockMovieReturn: &models.MovieInfo{
				Id:    1,
				Title: "Test Movie",
			},
			mockMovieErr: nil,
			mockActorsReturn: []*models.ActorInfo{{
				Id: 1,
				Person: models.Person{
					Name:    "Test",
					Surname: "Tester",
				}}},
			mockActorsErr: nil,
			expectedResponse: &models.MovieInfo{
				Id:    1,
				Title: "Test Movie",
				Actors: []*models.ActorInfo{{
					Id: 1,
					Person: models.Person{
						Name:    "Test",
						Surname: "Tester",
					},
				}},
			},
			expectedError: nil,
			statusCode:    http.StatusOK,
		},
		{
			name:             "Error_GetMovie",
			mvId:             1,
			mockMovieReturn:  nil,
			mockMovieErr:     &errVals.ErrorObj{Code: "movie_not_found", Error: errVals.CustomError{Err: errors.New("movie not found")}},
			mockActorsReturn: nil,
			mockActorsErr:    nil,
			expectedResponse: nil,
			expectedError: &models.ErrorRespData{
				StatusCode: http.StatusNotFound,
				Errors:     []errVals.ErrorObj{{Code: "movie_not_found", Error: errVals.CustomError{Err: errors.New("movie not found")}}},
			},
			statusCode: http.StatusNotFound,
		},
		{
			name: "Error_GetMovieActors",
			mvId: 1,
			mockMovieReturn: &models.MovieInfo{
				Id:    1,
				Title: "Test Movie",
			},
			mockMovieErr:     nil,
			mockActorsReturn: nil,
			mockActorsErr:    &errVals.ErrorObj{Code: "actors_fetch_error", Error: errVals.CustomError{Err: errors.New("failed to fetch actors")}},
			expectedResponse: nil,
			expectedError: &models.ErrorRespData{
				StatusCode: http.StatusInternalServerError,
				Errors:     []errVals.ErrorObj{{Code: "actors_fetch_error", Error: errVals.CustomError{Err: errors.New("failed to fetch actors")}}},
			},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := servMock.NewMockMovieRepositoryInterface(ctrl)
			s := NewService(repo)

			repo.EXPECT().GetMovie(gomock.Any(), test.mvId).Return(test.mockMovieReturn, test.mockMovieErr, test.statusCode)

			if test.mockMovieErr == nil {
				repo.EXPECT().GetMovieActors(gomock.Any(), test.mvId).Return(test.mockActorsReturn, test.mockActorsErr, test.statusCode)
			}

			response, err := s.GetMovie(context.Background(), test.mvId)

			if test.expectedError != nil {
				assert.Nil(t, response)
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResponse, response)
				assert.Nil(t, err)
			}
		})
	}
}
