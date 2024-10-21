package delivery

import (
	"bytes"
	"context"
	"errors"
	"fmt"
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
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
						Movies: []*models.MovieInfo{
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
			resp:       `{"success":true,"collections":[{"id":1,"title":"Test collection","movies":[{"id":1,"title":"test movie","card_url":"","album_url":"","rating":0,"release_date":"0001-01-01T00:00:00Z","movie_type":"","country":""}]}]}`,
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
			handler := NewMovieHandler(testContext(), srv)

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

func testContext() context.Context {
	err := os.Chdir("../../../..")
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("failed to change directory: %v", err))
	}

	cfg, err := config.New(zerolog.Logger{}, false, nil)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("failed to read config: %v", err))
	}

	ctx := config.WrapContext(context.Background(), cfg)

	return ctx
}
