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
	srvMock "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery/mocks"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestDelivery_Register(t *testing.T) {
	tests := []struct {
		name         string
		req          string
		resp         string
		mockReturn   *models.AuthRespData
		mockErr      *models.ErrorRespData
		statusCode   int
		isValidation bool
	}{
		{
			name: "Success",
			req:  `{"email": "test@mail.ru", "username": "tester", "password": "123456789", "passwordConfirmation": "123456789"}`,
			mockReturn: &models.AuthRespData{
				NewCookie: &models.CookieData{
					Name: "session_id",
					Token: &models.Token{
						TokenID: "cookie value",
						UserID:  1,
					},
				},
				StatusCode: http.StatusOK,
			},
			resp:       `{"success":true}`,
			statusCode: http.StatusOK,
		},
		{
			name: "Service Error",
			req:  `{"email": "test@mail.ru", "username": "tester", "password": "123456789", "passwordConfirmation": "123456789"}`,
			mockErr: &models.ErrorRespData{
				StatusCode: http.StatusInternalServerError,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrGenerateTokenCode, errVals.CustomError{Err: errors.New("some token error")})},
			},
			statusCode: http.StatusInternalServerError,
			resp:       `{"success":false,"errors":[{"Code":"auth_token_generation_error","Error":"some token error"}]}`,
		},
		{
			name: "Validation Error",
			req:  `{"email": "invalid", "username": "tester", "password": "short", "passwordConfirmation": "short"}`,
			mockErr: &models.ErrorRespData{
				StatusCode: http.StatusBadRequest,
				Errors: []errVals.ErrorObj{
					{
						Code: errVals.ErrInvalidPasswordCode, Error: errVals.CustomError{Err: errVals.ErrInvalidPasswordText.Err},
					},
					{
						Code: errVals.ErrInvalidEmailCode, Error: errVals.CustomError{Err: errVals.ErrInvalidEmailText.Err},
					},
				},
			},
			statusCode:   http.StatusBadRequest,
			isValidation: true,
			resp:         `{"success":false,"errors":[{"Code":"invalid_password","Error":"password is too short. The minimal len is 8"},{"Code":"invalid_email","Error":"email is incorrect"}]}`,
		},
		{
			name: "Password mismatch",
			req:  `{"email": "test@mail.ru", "username": "tester", "password": "123456789", "passwordConfirmation": "12345678910"}`,
			mockErr: &models.ErrorRespData{
				StatusCode: http.StatusBadRequest,
				Errors: []errVals.ErrorObj{
					{
						Code: errVals.ErrInvalidPasswordCode, Error: errVals.CustomError{Err: errVals.ErrInvalidPasswordsMatchText.Err},
					},
				},
			},
			statusCode:   http.StatusBadRequest,
			isValidation: true,
			resp:         `{"success":false,"errors":[{"Code":"invalid_password","Error":"password doesnt match with passwordConfirmation"}]}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "/api/auth/signup"
			srv := srvMock.NewMockAuthServiceInterface(ctrl)
			handler := NewAuthHandler(srv, GetCfg())

			if !test.isValidation {
				srv.EXPECT().Register(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)
			}

			r := mux.NewRouter()
			r.HandleFunc(path, handler.Register)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", path, bytes.NewBufferString(test.req))

			r.ServeHTTP(w, req)

			assert.Equal(t, test.statusCode, w.Result().StatusCode)
			assert.JSONEq(t, test.resp, w.Body.String())
		})
	}
}

func TestDelivery_Login(t *testing.T) {
	tests := []struct {
		name       string
		req        string
		resp       string
		mockReturn *models.AuthRespData
		mockErr    *models.ErrorRespData
		statusCode int
	}{
		{
			name: "Success",
			req:  `{"email": "ashurov@mail.rs", "password": "A123456bb"}`,
			mockReturn: &models.AuthRespData{
				NewCookie: &models.CookieData{
					Name: "session_id",
					Token: &models.Token{
						TokenID: "cookie value",
						UserID:  1,
					},
				},
				StatusCode: http.StatusOK,
			},
			statusCode: http.StatusOK,
			resp:       `{"success": true}`,
		},
		{
			req:  `{"email": "ashurov@mail.rs", "password": "A123456bb"}`,
			name: "Service Error",
			mockErr: &models.ErrorRespData{
				StatusCode: http.StatusInternalServerError,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrInvalidPasswordCode, errVals.ErrInvalidPasswordsMatchText)},
			},
			resp:       `{"success":false,"errors":[{"Code":"invalid_password","Error":"password doesnt match with passwordConfirmation"}]}`,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "/api/auth/login"
			srv := srvMock.NewMockAuthServiceInterface(ctrl)
			handler := NewAuthHandler(srv, GetCfg())

			srv.EXPECT().Login(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)

			r := mux.NewRouter()
			r.HandleFunc(path, handler.Login)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", path, bytes.NewBufferString(test.req))

			r.ServeHTTP(w, req)

			assert.Equal(t, test.statusCode, w.Result().StatusCode)
			assert.JSONEq(t, test.resp, w.Body.String())
		})
	}
}

func TestDelivery_Logout(t *testing.T) {
	tests := []struct {
		name         string
		resp         string
		mockReturn   *models.AuthRespData
		mockErr      *models.ErrorRespData
		statusCode   int
		isValidation bool
		emptyCookie  bool
		noCookie     bool
	}{
		{
			name: "Success",
			mockReturn: &models.AuthRespData{
				StatusCode: http.StatusOK,
			},
			statusCode: http.StatusOK,
			resp:       `{"success": true}`,
		},
		{
			name: "Service Error",
			mockErr: &models.ErrorRespData{
				StatusCode: http.StatusInternalServerError,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrRedisClearCode, errVals.CustomError{Err: errors.New("some redis error")})},
			},
			statusCode: http.StatusInternalServerError,
			resp:       `{"success":false,"errors":[{"Code":"failed_delete_from_redis","Error":"some redis error"}]}`,
		},
		{
			name: "Validation Error",
			mockErr: &models.ErrorRespData{
				StatusCode: http.StatusBadRequest,
				Errors:     []errVals.ErrorObj{{Code: errVals.ErrBrokenCookieCode, Error: errVals.CustomError{Err: errVals.ErrBrokenCookieText.Err}}},
			},
			statusCode:   http.StatusBadRequest,
			isValidation: true,
			emptyCookie:  true,
			resp:         `{"success":false,"errors":[{"Code":"cookie_validation_error","Error":"Logout action: Invalid cookie err - broken cookie was given"}]}`,
		},
		{
			name:         "No cookie provided",
			statusCode:   http.StatusForbidden,
			isValidation: true,
			noCookie:     true,
			resp:         `{"success":false,"errors":[{"Code":"no_cookie_provided","Error":"Logout action: No cookie err - http: named cookie not present"}]}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "/api/auth/logout"
			srv := srvMock.NewMockAuthServiceInterface(ctrl)
			handler := NewAuthHandler(srv, GetCfg())

			r := mux.NewRouter()
			r.HandleFunc(path, handler.Logout)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", path, bytes.NewBufferString(""))

			if !test.isValidation {
				req.Header.Set("Cookie", "session_id=some_cookie")
				srv.EXPECT().Logout(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)
			} else if test.emptyCookie {
				req.Header.Set("Cookie", "session_id=")
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, test.statusCode, w.Result().StatusCode)
			assert.JSONEq(t, test.resp, w.Body.String())
		})
	}
}

func TestDelivery_Session(t *testing.T) {
	tests := []struct {
		name       string
		mockReturn *models.SessionRespData
		mockErr    *models.ErrorRespData
		resp       string
		statusCode int
		noCookie   bool
	}{
		{
			name: "Success",
			mockReturn: &models.SessionRespData{
				UserData: models.User{
					Id:       1,
					Email:    "test@mail.ru",
					Username: "Tester",
				},
				StatusCode: http.StatusOK,
			},
			resp:       `{"success":true,"user_data":{"id":1,"email":"test@mail.ru","username":"Tester"}}`,
			statusCode: http.StatusOK,
		},
		{
			name: "Service Error",
			mockErr: &models.ErrorRespData{
				StatusCode: http.StatusInternalServerError,
				Errors: []errVals.ErrorObj{*errVals.NewErrorObj(
					errVals.ErrCreateUserCode,
					errVals.CustomError{Err: errors.New("cannot get cookie from redis")},
				)},
			},
			resp:       `{"success":false,"errors":[{"Code":"create_user_error","Error":"cannot get cookie from redis"}]}`,
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "Forbidden",
			resp:       `{"success":false,"errors":[{"Code":"no_cookie_provided","Error":"Session action: No cookie err - http: named cookie not present"}]}`,
			statusCode: http.StatusForbidden,
			noCookie:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "/api/auth/session"
			srv := srvMock.NewMockAuthServiceInterface(ctrl)
			handler := NewAuthHandler(srv, GetCfg())

			r := mux.NewRouter()
			r.HandleFunc(path, handler.Session)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", path, bytes.NewBufferString(""))

			if !test.noCookie {
				req.Header.Set("Cookie", "session_id=some_cookie")
				srv.EXPECT().Session(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)
			}

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

	cfg, err := config.New(false, nil)
	if err != nil {
		log.Fatalf("failed to read config from Register test: %v", err)
	}

	return cfg
}
