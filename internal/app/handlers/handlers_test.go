package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/handlers"
	hMock "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/handlers/mocks"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetCollections_Success(t *testing.T) {
	tests := []struct {
		name             string
		mockReturn       *models.CollectionsResponse
		mockErr          *models.ErrorResponse
		expectedResponse *models.CollectionsResponse
	}{
		{
			name: "Success",
			mockReturn: &models.CollectionsResponse{
				StatusCode: 200,
				Success:    true,
				Collections: []models.Collection{
					{
						Id:    1,
						Title: "test collection",
						Movies: []*models.Movie{
							{
								Id:          1,
								Title:       "test movie",
								Description: "test description",
							},
						},
					},
				},
			},
			expectedResponse: &models.CollectionsResponse{
				Success: true,
				Collections: []models.Collection{
					{
						Id:    1,
						Title: "test collection",
						Movies: []*models.Movie{
							{
								Id:          1,
								Title:       "test movie",
								Description: "test description",
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			del := hMock.NewMockMovieImplementationInterface(ctrl)
			cfg := &config.Config{} // Инициализируйте вашу конфигурацию
			handler := handlers.NewMovieHandler(del, cfg)

			del.EXPECT().GetCollection(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)

			req := httptest.NewRequest(http.MethodGet, "/collections", nil)
			w := httptest.NewRecorder()

			handler.GetCollections(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)

			var actualResponse *models.CollectionsResponse
			err := json.NewDecoder(w.Body).Decode(&actualResponse)
			assert.NoError(t, err)

			assert.Equal(t, test.expectedResponse, actualResponse)
		})
	}
}

// func TestGetCollections_Error(t *testing.T) {
// 	mockApi := &MockApiLayer{
// 		GetCollectionFunc: func(ctx context.Context, query url.Values) (*models.CollectionsResponse, *models.ErrorResponse) {
// 			// Здесь создаем и возвращаем ошибку
// 			return nil, &models.ErrorResponse{StatusCode: http.StatusInternalServerError, Message: "Internal Server Error"}
// 		},
// 	}
// 	cfg := &config.Config{} // Инициализируйте вашу конфигурацию
// 	handler := handlers.NewMovieHandler(mockApi, cfg)

// 	req := httptest.NewRequest(http.MethodGet, "/collections", nil)
// 	w := httptest.NewRecorder()

// 	handler.GetCollections(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusInternalServerError, w.Code)
// 	// Дополнительные проверки на содержимое ответа...
// }
