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
				{Id: 1, Title: "Collection 1", Movies: []*models.MovieBaseInfo{}},
				{Id: 2, Title: "Collection 2", Movies: []*models.MovieBaseInfo{}},
			},
			mockErr: nil,
			expectedResponse: &models.CollectionsRespData{
				Collections: []models.Collection{
					{Id: 1, Title: "Collection 1", Movies: []*models.MovieBaseInfo{}},
					{Id: 2, Title: "Collection 2", Movies: []*models.MovieBaseInfo{}},
				},
				StatusCode: http.StatusOK,
			},
			expectedError: nil,
			statusCode:    http.StatusOK,
		},
		{
			name:             "Error",
			mockReturn:       nil,
			mockErr:          &errVals.ErrorObj{Code: "something_went_wrong", Error: errVals.CustomError{Err: errors.New("Database fail")}},
			expectedResponse: nil,
			expectedError: &models.ErrorRespData{
				StatusCode: http.StatusUnprocessableEntity,
				Errors:     []errVals.ErrorObj{{Code: "something_went_wrong", Error: errVals.CustomError{Err: errors.New("Database fail")}}},
			},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := servMock.NewMockMovieRepositoryInterface(ctrl)
			s := MovieService{movieRepository: repo}

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
