package delivery

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	srvMock "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/delivery/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestDelivery_GetCollection(t *testing.T) {
	tests := []struct {
		name       string
		mockReturn *models.CollectionsRespData
		mockErr    *models.ErrorRespData
		statusCode int
		resp       string
	}{
		{
			name: "Success",
			mockReturn: &models.CollectionsRespData{
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
				StatusCode: http.StatusOK,
			},
			resp:       `{"success":true,"collections":[{"Id":1,"Title":"Test collection","Movies":[{"Id":1,"Title":"test movie","Description":"some interesting movie","CardUrl":"","AlbumUrl":"","Rating":0,"ReleaseDate":"0001-01-01T00:00:00Z","MovieType":"","Country":""}]}]}`,
			statusCode: http.StatusOK,
		},
		{
			name: "Service Error",
			mockErr: &models.ErrorRespData{
				StatusCode: http.StatusInternalServerError,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: errors.New("Some database error")})},
			},
			resp:       `{"success":false,"errors":[{"Code":"something_went_wrong","Error":"Some database error"}]}`,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "/api/movie/movie_collections"
			srv := srvMock.NewMockMovieServiceInterface(ctrl)
			handler := NewMovieHandler(srv, GetCfg())

			srv.EXPECT().GetCollection(gomock.Any()).Return(test.mockReturn, test.mockErr)

			r := mux.NewRouter()
			r.HandleFunc(path, handler.GetCollections)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", path, bytes.NewBufferString(""))

			r.ServeHTTP(w, req)

			assert.Equal(t, test.statusCode, w.Result().StatusCode)
			assert.JSONEq(t, test.resp, w.Body.String())
		})
	}
}

func GetCfg() *config.Config {
	err := os.Chdir("../../../..")
	if err != nil {
		log.Fatalf("failed to change directory: %v", err)
	}

	cfg, err := config.New(false)
	if err != nil {
		log.Fatalf("failed to read config from Register test: %v", err)
	}

	return cfg
}
