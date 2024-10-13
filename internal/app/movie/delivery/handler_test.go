package delivery

import (
	"context"
	"errors"
	"testing"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	srvMock "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/delivery/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDelivery_GetCollection(t *testing.T) {
	tests := []struct {
		name             string
		mockReturn       *models.CollectionsResponse
		mockErr          *models.ErrorResponse
		expectedResponse *models.CollectionsResponse
		expectedErr      *models.ErrorResponse
	}{
		{
			name: "Success",
			mockReturn: &models.CollectionsResponse{
				Success: true,
				Collections: []models.Collection{
					{
						Id:    1,
						Title: "Test collection",
						Movies: []*models.Movie{
							{
								Id:          1,
								Title:       "test movie",
								Description: "some interesting movie",
							},
						},
					},
				},
				StatusCode: 200,
			},
			expectedResponse: &models.CollectionsResponse{
				Success: true,
				Collections: []models.Collection{
					{
						Id:    1,
						Title: "Test collection",
						Movies: []*models.Movie{
							{
								Id:          1,
								Title:       "test movie",
								Description: "some interesting movie",
							},
						},
					},
				},
				StatusCode: 200,
			},
		},
		{
			name: "Service Error",
			mockErr: &models.ErrorResponse{
				Success:    false,
				StatusCode: 500,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: errors.New("Some database error")})},
			},
			expectedErr: &models.ErrorResponse{
				Success:    false,
				StatusCode: 500,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: errors.New("Some database error")})},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			srv := srvMock.NewMockMovieServiceInterface(ctrl)
			imp := NewImplementation(context.Background(), srv)

			srv.EXPECT().GetCollection(gomock.Any()).Return(test.mockReturn, test.mockErr)
			resp, err := imp.GetCollection(context.Background(), nil)

			if test.expectedErr != nil {
				assert.Nil(t, resp)
				assert.Equal(t, test.expectedErr, err)
			} else {
				assert.Equal(t, test.expectedResponse, resp)
				assert.Nil(t, err)
			}
		})
	}
}
