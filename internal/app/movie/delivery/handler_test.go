package delivery

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

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
						ID:    1,
						Title: "Test collection",
						Movies: []*models.MovieShortInfo{
							{
								ID:    1,
								Title: "test movie",
							},
						},
					},
				},
				StatusCode: http.StatusOK,
			},
			resp:       `{"success":true,"collections":[{"id":1,"title":"Test collection","movies":[{"id":1,"title":"test movie","card_url":"", "album_url":"", "rating":0,"release_date":"0001-01-01T00:00:00Z","movie_type":"","country":""}]}]}`,
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
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "/api/movie/movie_collections"
			ms := srvMock.NewMockMovieServiceInterface(ctrl)
			handler := NewMovieHandler(ms)

			ms.EXPECT().GetCollection(gomock.Any()).Return(test.mockReturn, test.mockErr)

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

func TestDelivery_GetMovie(t *testing.T) {
	tests := []struct {
		name       string
		mockReturn *models.MovieInfo
		mockErr    *models.ErrorRespData
		statusCode int
		resp       string
		badReq     bool
	}{
		{
			name: "Success",
			mockReturn: &models.MovieInfo{
				ID:              1,
				Title:           "Test",
				FullDescription: "Test desc",
				CardUrl:         "card_link",
				AlbumUrl:        "album_link",
				Rating:          7.8,
				MovieType:       "film",
				Country:         "Russia",
				VideoUrl:        "video_link",
				Director:        &models.DirectorInfo{},
			},
			resp:       `{"success":true,"movie_info":{"id":1,"title":"Test","full_description":"Test desc","short_description":"","card_url":"card_link","album_url":"album_link","title_url":"","rating":7.8,"release_date":"0001-01-01T00:00:00Z","movie_type":"film", "director": "", "country":"Russia","video_url":"video_link","actors_info":null}}`,
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
		{
			name:       "Bad Request",
			resp:       `{"success":false,"errors":[{"Code":"bad_request","Error":"getMovie action: Bad request - strconv.Atoi: parsing \"\": invalid syntax"}]}`,
			statusCode: http.StatusBadRequest,
			badReq:     true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "/api/movies/1"
			ms := srvMock.NewMockMovieServiceInterface(ctrl)
			handler := NewMovieHandler(ms)

			r := mux.NewRouter()
			r.HandleFunc(path, handler.GetMovie)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", path, nil)

			vars := map[string]string{}
			if !test.badReq {
				ms.EXPECT().GetMovie(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)
				vars["movie_id"] = "1"
			}

			req = mux.SetURLVars(req, vars)

			handler.GetMovie(w, req)

			assert.Equal(t, test.statusCode, w.Result().StatusCode)
			assert.JSONEq(t, test.resp, w.Body.String())
		})
	}
}

func TestDelivery_GetActor(t *testing.T) {
	tests := []struct {
		name       string
		mockReturn *models.ActorInfo
		mockErr    *models.ErrorRespData
		statusCode int
		resp       string
		badReq     bool
	}{
		{
			name: "Success",
			mockReturn: &models.ActorInfo{
				ID: 1,
				Person: models.Person{
					Name:    "Tester",
					Surname: "Testov",
				},
			},
			resp:       `{"success":true,"actor_info":{"id":1,"full_name":"Tester Testov","biography":"","birthdate":"","photo_url":"","country":"", "movies":null}}`,
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
		{
			name:       "Bad Request",
			resp:       `{"success":false,"errors":[{"Code":"bad_request","Error":"getActor action: Bad request - strconv.Atoi: parsing \"\": invalid syntax"}]}`,
			statusCode: http.StatusBadRequest,
			badReq:     true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "/api/actors/1"
			ms := srvMock.NewMockMovieServiceInterface(ctrl)
			handler := NewMovieHandler(ms)

			r := mux.NewRouter()
			r.HandleFunc(path, handler.GetActor)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", path, nil)

			vars := map[string]string{}
			if !test.badReq {
				ms.EXPECT().GetActor(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)
				vars["actor_id"] = "1"
			}

			req = mux.SetURLVars(req, vars)

			handler.GetActor(w, req)

			assert.Equal(t, test.statusCode, w.Result().StatusCode)
			assert.JSONEq(t, test.resp, w.Body.String())
		})
	}
}
