package delivery

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	mockSrv "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/delivery/mocks"
)

func TestUserHandler_UpdatePassword(t *testing.T) {
	ctx := testContext(t)
	tests := []struct {
		name        string
		reqBody     string
		mockErr     *errVals.ServiceError
		statusCode  int
		resp        string
		skipService bool
	}{
		{
			name:       "Success",
			reqBody:    `{"password": "newpass123", "passwordConfirmation": "newpass123"}`,
			resp:       "",
			statusCode: http.StatusOK,
		},
		{
			name:        "Parse_error",
			reqBody:     `{"password": "newpass123"}`,
			resp:        `{"errors":[{"code":"user_request_parse_error","error":"updateProfile action: Path params err - strconv.Atoi: parsing \"\": invalid syntax"}]}`,
			statusCode:  http.StatusBadRequest,
			skipService: true,
		},
		{
			name:        "Validation error",
			reqBody:     `{"password": "newpass123", "passwordConfirmation": "wrongpass"}`,
			resp:        `{"errors":[{"code":"user_validation_error","error":"updatePassword action: Password err - password doesn't match with passwordConfirmation"}]}`,
			statusCode:  http.StatusBadRequest,
			skipService: true,
		},
		{
			name:       "Service error",
			reqBody:    `{"password": "newpass123", "passwordConfirmation": "newpass123"}`,
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

			req := httptest.NewRequest(http.MethodPost, "/users/1/password", bytes.NewBufferString(test.reqBody))
			if test.name != "Parse_error" {
				req = mux.SetURLVars(req, map[string]string{"id": "1"})
			}
			req.Header.Set("Content-Type", "application/json")

			ms := mockSrv.NewMockUserServiceInterface(ctrl)
			handler := NewUserHandler(ctx, ms)

			if !test.skipService {
				ms.EXPECT().UpdatePassword(gomock.Any(), gomock.Any()).Return(test.mockErr)
			}

			w := httptest.NewRecorder()

			handler.UpdatePassword(w, req)
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, test.statusCode, w.Result().StatusCode)
			if test.resp == "" {
				assert.Equal(t, test.resp, w.Body.String())
			} else {
				assert.JSONEq(t, test.resp, w.Body.String())
			}
		})
	}
}

func TestUserHandler_UpdateProfile(t *testing.T) {
	ctx := testContext(t)

	tests := []struct {
		name        string
		formData    map[string]string
		fileData    []byte
		mockErr     *errVals.ServiceError
		statusCode  int
		resp        string
		skipService bool
	}{
		{
			name: "Success",
			formData: map[string]string{
				"email":    "test@mail.ru",
				"username": "testuser",
			},
			fileData:   []byte("fake image data"),
			resp:       "",
			statusCode: http.StatusOK,
		},
		{
			name: "Parse_error",
			formData: map[string]string{
				"email": "test@mail.ru",
			},
			resp:        `{"errors":[{"code":"user_request_parse_error","error":"updateProfile action: Path params err - strconv.Atoi: parsing \"\": invalid syntax"}]}`,
			statusCode:  http.StatusBadRequest,
			skipService: true,
		},
		{
			name: "Validation error",
			formData: map[string]string{
				"email":    "invalid-email",
				"username": "testuser",
			},
			resp:        `{"errors":[{"code":"user_validation_error","error":"updateProfile action: Email err - email is incorrect"}]}`,
			statusCode:  http.StatusBadRequest,
			skipService: true,
		},
		{
			name: "Service error",
			formData: map[string]string{
				"email":    "test@example.com",
				"username": "testuser",
			},
			fileData:   []byte("fake image data"),
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

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)

			for key, val := range test.formData {
				_ = writer.WriteField(key, val)
			}

			if test.fileData != nil {
				part, _ := writer.CreateFormFile("avatar", "avatar.png")
				part.Write(test.fileData)
			}

			writer.Close()

			req := httptest.NewRequest(http.MethodPost, "/users/1/update_profile", body)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			if test.name != "Parse_error" {
				req = mux.SetURLVars(req, map[string]string{"id": "1"})
			}

			ms := mockSrv.NewMockUserServiceInterface(ctrl)
			handler := NewUserHandler(ctx, ms)

			if !test.skipService {
				ms.EXPECT().UpdateProfile(gomock.Any(), gomock.Any()).Return(test.mockErr)
			}

			w := httptest.NewRecorder()
			handler.UpdateProfile(w, req)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, test.statusCode, res.StatusCode)
			if test.resp == "" {
				assert.Equal(t, test.resp, w.Body.String())
			} else {
				assert.JSONEq(t, test.resp, w.Body.String())
			}
		})
	}
}

func TestUserHandler_SetFavorite(t *testing.T) {
	ctx := testContext(t)
	tests := []struct {
		name       string
		reqBody    string
		mockErr    *errVals.ServiceError
		statusCode int
		resp       string
	}{
		{
			name:       "Success",
			reqBody:    `{"movieID": 101}`,
			resp:       "",
			statusCode: http.StatusOK,
		},
		{
			name:       "Service error",
			reqBody:    `{"movieID": 101}`,
			mockErr:    errVals.NewServiceError(errVals.ErrServerCode, errors.New("database error")),
			resp:       `{"errors":[{"code":"something_went_wrong","error":"database error"}]}`,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			req := httptest.NewRequest(http.MethodPost, "/favorites", bytes.NewBufferString(test.reqBody))
			req.Header.Set("Content-Type", "application/json")

			ms := mockSrv.NewMockUserServiceInterface(ctrl)
			handler := NewUserHandler(ctx, ms)

			ms.EXPECT().AddFavorite(gomock.Any(), gomock.Any()).Return(test.mockErr)

			w := httptest.NewRecorder()

			handler.SetFavorite(w, req)
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, test.statusCode, w.Result().StatusCode)
			if test.resp == "" {
				assert.Equal(t, test.resp, w.Body.String())
			} else {
				assert.JSONEq(t, test.resp, w.Body.String())
			}
		})
	}
}

func TestUserHandler_ResetFavorite(t *testing.T) {
	ctx := testContext(t)
	tests := []struct {
		name       string
		reqBody    string
		mockErr    *errVals.ServiceError
		statusCode int
		resp       string
	}{
		{
			name:       "Success",
			reqBody:    `{"movieID": 101}`,
			resp:       "",
			statusCode: http.StatusOK,
		},
		{
			name:       "Service error",
			reqBody:    `{"movieID": 101}`,
			mockErr:    errVals.NewServiceError(errVals.ErrServerCode, errors.New("database error")),
			resp:       `{"errors":[{"code":"something_went_wrong","error":"database error"}]}`,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			req := httptest.NewRequest(http.MethodDelete, "/favorites", bytes.NewBufferString(test.reqBody))
			req.Header.Set("Content-Type", "application/json")

			ms := mockSrv.NewMockUserServiceInterface(ctrl)
			handler := NewUserHandler(ctx, ms)

			ms.EXPECT().ResetFavorite(gomock.Any(), gomock.Any()).Return(test.mockErr)

			w := httptest.NewRecorder()

			handler.ResetFavorite(w, req)
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, test.statusCode, w.Result().StatusCode)
			if test.resp == "" {
				assert.Equal(t, test.resp, w.Body.String())
			} else {
				assert.JSONEq(t, test.resp, w.Body.String())
			}
		})
	}
}

func TestUserHandler_GetFavorites(t *testing.T) {
	ctx := testContext(t)
	tests := []struct {
		name       string
		usrID      string
		mockResp   []models.MovieShortInfo
		mockErr    *errVals.ServiceError
		statusCode int
		resp       string
	}{
		{
			name:       "Success",
			usrID:      "1",
			mockResp:   []models.MovieShortInfo{{ID: 1, Title: "Test"}},
			resp:       `{"movies":[{"id":1,"title":"Test","card_url":"","album_url":"","rating":0,"release_date":"","movie_type":"","country":""}]}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "Service error",
			usrID:      "1",
			mockErr:    errVals.NewServiceError(errVals.ErrServerCode, errors.New("database error")),
			resp:       `{"errors":[{"code":"something_went_wrong","error":"database error"}]}`,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			req := httptest.NewRequest(http.MethodGet, "/users/1/favorites", nil)
			req = mux.SetURLVars(req, map[string]string{"id": test.usrID})

			ms := mockSrv.NewMockUserServiceInterface(ctrl)
			handler := NewUserHandler(ctx, ms)

			ms.EXPECT().GetFavorites(gomock.Any(), gomock.Any()).Return(test.mockResp, test.mockErr)

			w := httptest.NewRecorder()

			handler.GetFavorites(w, req)
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, test.statusCode, w.Result().StatusCode)
			if test.resp == "" {
				assert.Equal(t, test.resp, w.Body.String())
			} else {
				assert.JSONEq(t, test.resp, w.Body.String())
			}
		})
	}
}

func testContext(t *testing.T) context.Context {
	require.NoError(t, os.Chdir("../../../.."), "failed to change directory")

	cfg, err := config.New(true)
	require.NoError(t, err, "failed to read config from user handler_test")

	return config.WrapContext(context.Background(), cfg)
}
