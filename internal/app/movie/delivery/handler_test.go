package delivery

import (
	"bytes"
	"errors"
	"io"
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
		name           string
		mockReturn     *models.MovieInfo
		mockRateReturn int32
		mockErr        *errVals.ServiceError
		statusCode     int
		resp           string
		badReq         bool
	}{
		{
			name:           "Success",
			mockRateReturn: 9,
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
				Genres:           []string{},
			},
			resp:       `{"movie_info":{"id":1,"title":"Test","full_description":"Test desc","short_description":"","card_url":"card_link","album_url":"album_link","title_url":"","rating":7.8,"release_date":"","movie_type":"film","country":"Russia","video_url":"video_link","director":"","actors_info":[],"seasons":null,"is_favorite":false,"with_subscription":false,"rating_user":9, "genres":[]}}`,
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
				ms.EXPECT().GetUserRating(gomock.Any(), gomock.Any()).Return(test.mockRateReturn, test.mockErr)
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

func TestMovieHandler_SearchMovies(t *testing.T) {
	tests := []struct {
		name         string
		queryParam   string
		mockMovies   []models.MovieInfo
		mockErr      error
		expectedResp string
		statusCode   int
	}{
		{
			name:       "Success - Movies Found",
			queryParam: "Inception",
			mockMovies: []models.MovieInfo{
				{
					ID:          1,
					Title:       "Inception",
					CardURL:     "/card.webp",
					AlbumURL:    "/album.webp",
					Rating:      8.8,
					ReleaseDate: "2010-07-16",
					MovieType:   "Science Fiction",
					Country:     "USA",
				},
			},
			expectedResp: `[{"id":1,"title":"Inception","card_url":"/card.webp","album_url":"/album.webp","rating":"8.8","release_date":"2010-07-16","movie_type":"Science Fiction","country":"USA"}]`,
			statusCode:   http.StatusOK,
		},
		{
			name:         "Success - No Movies Found",
			queryParam:   "Unknown",
			mockMovies:   []models.MovieInfo{},
			expectedResp: `[{}]`,
			statusCode:   http.StatusOK,
		},
		{
			name:         "Error - Query Parameter Missing",
			queryParam:   "",
			expectedResp: "query parameter is required\n",
			statusCode:   http.StatusBadRequest,
		},
		{
			name:         "Error - Service Failure",
			queryParam:   "Matrix",
			mockErr:      errors.New("internal server error"),
			expectedResp: "search error: internal server error\n",
			statusCode:   http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockMovieServiceInterface(ctrl)
			handler := NewMovieHandler(mockService)

			if test.queryParam != "" && test.mockErr == nil {
				mockService.EXPECT().
					SearchMovies(gomock.Any(), test.queryParam).
					Return(test.mockMovies, nil)
			} else if test.mockErr != nil {
				mockService.EXPECT().
					SearchMovies(gomock.Any(), test.queryParam).
					Return(nil, test.mockErr)
			}

			req := httptest.NewRequest(http.MethodGet, "/movies/search?query="+test.queryParam, nil)
			w := httptest.NewRecorder()

			handler.SearchMovies(w, req)
			res := w.Result()

			defer func() {
				clErr := res.Body.Close()
				assert.NoError(t, clErr)
			}()

			assert.Equal(t, test.statusCode, res.StatusCode)

			body, _ := io.ReadAll(res.Body)
			if test.statusCode == http.StatusOK {
				assert.JSONEq(t, test.expectedResp, string(body))
			} else {
				assert.Equal(t, test.expectedResp, string(body))
			}
		})
	}
}

func TestMovieHandler_SearchActors(t *testing.T) {
	tests := []struct {
		name         string
		queryParam   string
		mockActors   []models.ActorInfo
		mockErr      error
		expectedResp string
		statusCode   int
	}{
		{
			name:       "Success - Actors Found",
			queryParam: "Leonardo",
			mockActors: []models.ActorInfo{
				{
					ID: 1,
					Person: models.Person{
						Name:    "Leonardo",
						Surname: "DiCaprio",
					},
					BigPhotoURL: "/photo.webp",
					Country:     "USA",
				},
			},
			expectedResp: `[{"id":1,"full_name":"Leonardo DiCaprio","photo_url":"/photo.webp","country":"USA"}]`,
			statusCode:   http.StatusOK,
		},
		{
			name:         "Success - No Actors Found",
			queryParam:   "Unknown",
			mockActors:   []models.ActorInfo{},
			expectedResp: `[{}]`,
			statusCode:   http.StatusOK,
		},
		{
			name:         "Error - Query Parameter Missing",
			queryParam:   "",
			expectedResp: "query parameter is required\n",
			statusCode:   http.StatusBadRequest,
		},
		{
			name:         "Error - Service Failure",
			queryParam:   "Smith",
			mockErr:      errors.New("internal server error"),
			expectedResp: "search error: internal server error\n",
			statusCode:   http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockMovieServiceInterface(ctrl)
			handler := NewMovieHandler(mockService)

			if test.queryParam != "" && test.mockErr == nil {
				mockService.EXPECT().
					SearchActors(gomock.Any(), test.queryParam).
					Return(test.mockActors, nil)
			} else if test.mockErr != nil {
				mockService.EXPECT().
					SearchActors(gomock.Any(), test.queryParam).
					Return(nil, test.mockErr)
			}

			req := httptest.NewRequest(http.MethodGet, "/actors/search?query="+test.queryParam, nil)
			w := httptest.NewRecorder()

			handler.SearchActors(w, req)
			res := w.Result()

			defer func() {
				clErr := res.Body.Close()
				assert.NoError(t, clErr)
			}()

			assert.Equal(t, test.statusCode, res.StatusCode)

			body, _ := io.ReadAll(res.Body)
			if test.statusCode == http.StatusOK {
				assert.JSONEq(t, test.expectedResp, string(body))
			} else {
				assert.Equal(t, test.expectedResp, string(body))
			}
		})
	}
}

func TestMovieHandler_GetUserRating(t *testing.T) {
	tests := []struct {
		name         string
		queryParam   string
		mockRating   int32
		mockErr      *errVals.ServiceError
		expectedResp string
		statusCode   int
	}{
		{
			name:         "Success - Valid Rating",
			queryParam:   "1",
			mockRating:   int32(5),
			expectedResp: `{"rating":5}`,
			statusCode:   http.StatusOK,
		},
		{
			name:         "Error - Invalid Movie ID",
			queryParam:   "invalid",
			expectedResp: `{"errors":[{"code":"bad_request","error":"strconv.Atoi: parsing \"invalid\": invalid syntax"}]}`,
			statusCode:   http.StatusBadRequest,
		},
		{
			name:         "Error - Service Failure",
			queryParam:   "2",
			mockErr:      errVals.NewServiceError("internal_error", errors.New("internal server error")),
			expectedResp: `{"errors":[{"code":"internal_error","error":"internal server error"}]}`,
			statusCode:   http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockMovieServiceInterface(ctrl)
			handler := NewMovieHandler(mockService)

			if test.mockErr == nil && test.queryParam != "invalid" {
				mockService.EXPECT().
					GetUserRating(gomock.Any(), gomock.Eq(int32(1))).
					Return(test.mockRating, nil)
			} else if test.mockErr != nil {
				mockService.EXPECT().
					GetUserRating(gomock.Any(), gomock.Any()).
					Return(int32(0), test.mockErr)
			}

			req := httptest.NewRequest(http.MethodGet, "/movies/rating?movie_id="+test.queryParam, nil)
			w := httptest.NewRecorder()

			handler.GetUserRating(w, req)
			res := w.Result()

			defer func() {
				clErr := res.Body.Close()
				assert.NoError(t, clErr)
			}()

			assert.Equal(t, test.statusCode, res.StatusCode)

			body, _ := io.ReadAll(res.Body)
			if test.statusCode == http.StatusOK {
				assert.JSONEq(t, test.expectedResp, string(body))
			} else {
				assert.Equal(t, test.expectedResp, string(body))
			}
		})
	}
}

func TestMovieHandler_AddOrUpdateRating(t *testing.T) {
	tests := []struct {
		name         string
		movieID      string
		reqBody      string
		mockErr      *errVals.ServiceError
		expectedResp string
		statusCode   int
		skipService  bool
	}{
		{
			name:         "Success - Rating Updated",
			movieID:      "1",
			reqBody:      `{"rating": 8}`,
			expectedResp: `{"message":"rating updated"}`,
			statusCode:   http.StatusOK,
		},
		{
			name:         "Error - Invalid Movie ID",
			movieID:      "invalid",
			reqBody:      `{"rating": 8}`,
			expectedResp: `{"errors":[{"code":"bad_request","error":"getMovie action: Bad request - strconv.Atoi: parsing \"invalid\": invalid syntax"}]}`,
			statusCode:   http.StatusBadRequest,
		},
		{
			name:         "Error - Invalid Rating",
			movieID:      "1",
			reqBody:      `{"rating": 15}`,
			expectedResp: `{"errors":[{"code":"bad_request","error":"rating must be between 1 and 10"}]}`,
			statusCode:   http.StatusBadRequest,
			skipService:  true,
		},
		{
			name:         "Error - Service Failure",
			movieID:      "1",
			reqBody:      `{"rating": 7}`,
			mockErr:      errVals.NewServiceError("something_went_wrong", errors.New("internal server error")),
			expectedResp: `{"errors":[{"code":"something_went_wrong","error":"internal server error"}]}`,
			statusCode:   http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockMovieServiceInterface(ctrl)
			handler := NewMovieHandler(mockService)

			if test.mockErr == nil && test.movieID != "invalid" && !test.skipService {
				mockService.EXPECT().
					AddOrUpdateRating(gomock.Any(), gomock.Eq(int32(1)), gomock.Eq(int32(8))).
					Return(nil)
			} else if test.mockErr != nil {
				mockService.EXPECT().
					AddOrUpdateRating(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(test.mockErr)
			}

			req := httptest.NewRequest(http.MethodPost, "/movies/1/rating", bytes.NewBufferString(test.reqBody))
			req = mux.SetURLVars(req, map[string]string{"movie_id": test.movieID})
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			handler.AddOrUpdateRating(w, req)
			res := w.Result()

			defer func() {
				clErr := res.Body.Close()
				assert.NoError(t, clErr)
			}()

			assert.Equal(t, test.statusCode, res.StatusCode)

			body, _ := io.ReadAll(res.Body)
			if test.statusCode == http.StatusOK {
				assert.JSONEq(t, test.expectedResp, string(body))
			} else {
				assert.Equal(t, test.expectedResp, string(body))
			}
		})
	}
}
