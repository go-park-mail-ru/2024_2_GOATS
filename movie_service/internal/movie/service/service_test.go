package service

import (
	"context"
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	servMock "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_GetCollection(t *testing.T) {
	tests := []struct {
		name             string
		mockReturn       []models.Collection
		mockErr          error
		expectedResponse *models.CollectionsRespData
		expectedError    error
	}{
		{
			name: "Success",
			mockReturn: []models.Collection{
				{ID: 1, Title: "Collection 1", Movies: []*models.MovieShortInfo{}},
				{ID: 2, Title: "Collection 2", Movies: []*models.MovieShortInfo{}},
			},
			mockErr: nil,
			expectedResponse: &models.CollectionsRespData{
				Collections: []models.Collection{
					{ID: 1, Title: "Collection 1", Movies: []*models.MovieShortInfo{}},
					{ID: 2, Title: "Collection 2", Movies: []*models.MovieShortInfo{}},
				},
			},
			expectedError: nil,
		},
		{
			name:             "Error",
			mockReturn:       nil,
			mockErr:          errors.New("Database fail"),
			expectedResponse: nil,
			expectedError:    errors.New("movieService.GetCollection: Database fail"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mr := servMock.NewMockMovieRepositoryInterface(ctrl)

			s := NewMovieService(mr)

			mr.EXPECT().GetCollection(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)

			response, err := s.GetCollection(context.Background(), "")

			if test.expectedError != nil {
				assert.Nil(t, response)
				assert.Equal(t, test.expectedError.Error(), err.Error())
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
		actorID          int
		mockReturn       *models.ActorInfo
		mockErr          error
		expectedResponse *models.ActorInfo
		expectedError    error
	}{
		{
			name:    "Success",
			actorID: 1,
			mockReturn: &models.ActorInfo{
				ID: 1,
				Person: models.Person{
					Name:    "Test",
					Surname: "Tester",
				},
			},
			mockErr: nil,
			expectedResponse: &models.ActorInfo{
				ID: 1,
				Person: models.Person{
					Name:    "Test",
					Surname: "Tester",
				},
			},
			expectedError: nil,
		},
		{
			name:             "Error",
			actorID:          1,
			mockReturn:       nil,
			mockErr:          errors.New("error"),
			expectedResponse: nil,
			expectedError:    errors.New("movieService.GetActor: error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mr := servMock.NewMockMovieRepositoryInterface(ctrl)
			s := NewMovieService(mr)

			mr.EXPECT().GetActor(gomock.Any(), test.actorID).Return(test.mockReturn, test.mockErr)

			response, err := s.GetActor(context.Background(), test.actorID)

			if test.expectedError != nil {
				assert.Nil(t, response)
				assert.Equal(t, test.expectedError.Error(), err.Error())
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
		mvID             int
		mockMovieReturn  *models.MovieInfo
		mockMovieErr     error
		mockActorsReturn []*models.ActorInfo
		mockActorsErr    error
		expectedResponse *models.MovieInfo
		expectedError    error
	}{
		{
			name: "Success",
			mvID: 1,
			mockMovieReturn: &models.MovieInfo{
				ID:    1,
				Title: "Test Movie",
			},
			mockMovieErr: nil,
			mockActorsReturn: []*models.ActorInfo{{
				ID: 1,
				Person: models.Person{
					Name:    "Test",
					Surname: "Tester",
				}}},
			mockActorsErr: nil,
			expectedResponse: &models.MovieInfo{
				ID:    1,
				Title: "Test Movie",
				Actors: []*models.ActorInfo{{
					ID: 1,
					Person: models.Person{
						Name:    "Test",
						Surname: "Tester",
					},
				}},
			},
			expectedError: nil,
		},
		{
			name:             "Error_GetMovie",
			mvID:             1,
			mockMovieReturn:  nil,
			mockMovieErr:     errors.New("error"),
			mockActorsReturn: nil,
			mockActorsErr:    nil,
			expectedResponse: nil,
			expectedError:    errors.New("movieService.GetMovie: error"),
		},
		{
			name: "Error_GetMovieActors",
			mvID: 1,
			mockMovieReturn: &models.MovieInfo{
				ID:    1,
				Title: "Test Movie",
			},
			mockMovieErr:     nil,
			mockActorsReturn: nil,
			mockActorsErr:    errors.New("error"),
			expectedResponse: nil,
			expectedError:    errors.New("movieService.GetMovieActors: error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mr := servMock.NewMockMovieRepositoryInterface(ctrl)
			s := NewMovieService(mr)

			mr.EXPECT().GetMovie(gomock.Any(), test.mvID).Return(test.mockMovieReturn, test.mockMovieErr)

			if test.mockMovieErr == nil {
				mr.EXPECT().GetMovieActors(gomock.Any(), test.mvID).Return(test.mockActorsReturn, test.mockActorsErr)
			}

			response, err := s.GetMovie(context.Background(), test.mvID)

			if test.expectedError != nil {
				assert.Nil(t, response)
				assert.Equal(t, test.expectedError.Error(), err.Error())
			} else {
				assert.Equal(t, test.expectedResponse, response)
				assert.Nil(t, err)
			}
		})
	}
}
