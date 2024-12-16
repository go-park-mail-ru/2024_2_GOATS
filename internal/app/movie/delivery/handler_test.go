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
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const (
	mvCollPath = "/api/movie_service/movie_collections"
	actorsPath = "/api/actors/1"
	moviePath  = "/api/movies/1"
)

func TestDelivery_GetCollection(t *testing.T) {
	tests := []struct {
		name       string
		mockReturn *models.CollectionsRespData
		mockErr    *errVals.ServiceError
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
								Title: "test movie_service",
							},
						},
					},
				},
			},
			resp:       `{"collections":[{"id":1,"title":"Test collection","movies":[{"id":1,"title":"test movie_service","card_url":"", "album_url":"", "rating":0,"release_date":"","movie_type":"","country":""}]}]}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "Service Error",
			mockErr:    errVals.NewServiceError(errVals.ErrServerCode, errors.New("Some database error")),
			resp:       `{"errors":[{"code":"something_went_wrong","error":"Some database error"}]}`,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ms := srvMock.NewMockMovieServiceInterface(ctrl)
			handler := NewMovieHandler(ms)

			ms.EXPECT().GetCollection(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)

			r := mux.NewRouter()
			r.HandleFunc(mvCollPath, handler.GetCollections)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", mvCollPath, bytes.NewBufferString(""))

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
		mockErr    *errVals.ServiceError
		statusCode int
		resp       string
		badReq     bool
	}{
		{
			name: "Success",
			mockReturn: &models.MovieInfo{
				ID:               1,
				Title:            "Test",
				FullDescription:  "Test desc",
				CardURL:          "card_link",
				AlbumURL:         "album_link",
				Rating:           7.8,
				MovieType:        "film",
				Country:          "Russia",
				VideoURL:         "video_link",
				Director:         &models.DirectorInfo{},
				WithSubscription: false,
			},
			resp:       `{"movie_info":{"id":1,"title":"Test","full_description":"Test desc","short_description":"","card_url":"card_link","album_url":"album_link","title_url":"","rating":7.8,"release_date":"","movie_type":"film","country":"Russia","video_url":"video_link","director":"","actors_info":[],"seasons":null,"is_favorite":false,"with_subscription":false}}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "Service Error",
			mockErr:    errVals.NewServiceError(errVals.ErrServerCode, errors.New("Some database error")),
			resp:       `{"errors":[{"code":"something_went_wrong","error":"Some database error"}]}`,
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "Bad Request",
			resp:       `{"errors":[{"code":"bad_request","error":"getMovie action: Bad request - strconv.Atoi: parsing \"\": invalid syntax"}]}`,
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

			ms := srvMock.NewMockMovieServiceInterface(ctrl)
			handler := NewMovieHandler(ms)

			r := mux.NewRouter()
			r.HandleFunc(moviePath, handler.GetMovie)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", moviePath, nil)

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
		mockErr    *errVals.ServiceError
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
			resp:       `{"actor_info":{"id":1,"full_name":"Tester Testov","biography":"","birthdate":"","photo_url":"","country":"", "movies":null}}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "Service Error",
			mockErr:    errVals.NewServiceError(errVals.ErrServerCode, errors.New("Some database error")),
			resp:       `{"errors":[{"code":"something_went_wrong","error":"Some database error"}]}`,
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "Bad Request",
			resp:       `{"errors":[{"code":"bad_request","error":"getActor action: Bad request - strconv.Atoi: parsing \"\": invalid syntax"}]}`,
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

			ms := srvMock.NewMockMovieServiceInterface(ctrl)
			handler := NewMovieHandler(ms)

			r := mux.NewRouter()
			r.HandleFunc(actorsPath, handler.GetActor)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", actorsPath, nil)

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
