package delivery

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	authSrvMock "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery/mocks"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	usrSrvMock "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/delivery/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

const (
	registerPath = "/api/auth/signup"
	loginPath    = "/api/auth/login"
	logoutPath   = "/api/auth/logout"
	sessionPath  = "/api/auth/session"
)

func TestDelivery_Register(t *testing.T) {
	tests := []struct {
		name         string
		req          string
		resp         string
		mockReturn   *models.AuthRespData
		mockErr      *errVals.ServiceError
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
			},
			resp:       `{}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "Service Error",
			req:        `{"email": "test@mail.ru", "username": "tester", "password": "123456789", "passwordConfirmation": "123456789"}`,
			mockErr:    errVals.NewServiceError(errVals.ErrGenerateTokenCode, errors.New("some token error")),
			statusCode: http.StatusInternalServerError,
			resp:       `{"errors":[{"code":"auth_token_generation_error","error":"some token error"}]}`,
		},
		{
			name:         "Validation Error",
			req:          `{"email": "invalid", "username": "tester", "password": "short", "passwordConfirmation": "short"}`,
			statusCode:   http.StatusBadRequest,
			isValidation: true,
			resp:         `{"errors":[{"code":"invalid_password","error":"password is too short. The minimal len is 8"},{"code":"invalid_email","error":"email is incorrect"}]}`,
		},
		{
			name: "Password mismatch",
			req:  `{"email": "test@mail.ru", "username": "tester", "password": "123456789", "passwordConfirmation": "12345678910"}`,
			mockErr: &errVals.ServiceError{
				Code:  errVals.ErrInvalidPasswordCode,
				Error: errVals.ErrInvalidPasswordsMatch,
			},
			statusCode:   http.StatusBadRequest,
			isValidation: true,
			resp:         `{"errors":[{"code":"invalid_password","error":"password doesn't match with passwordConfirmation"}]}`,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mAuthSrv := authSrvMock.NewMockAuthServiceInterface(ctrl)
			mUsrSrv := usrSrvMock.NewMockUserServiceInterface(ctrl)
			handler := NewAuthHandler(mAuthSrv, mUsrSrv)

			if !test.isValidation {
				mAuthSrv.EXPECT().Register(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)
			}

			r := mux.NewRouter()
			r.HandleFunc(registerPath, handler.Register)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", registerPath, bytes.NewBufferString(test.req))

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
		mockErr    *errVals.ServiceError
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
			},
			statusCode: http.StatusOK,
			resp:       `{}`,
		},
		{
			req:        `{"email": "ashurov@mail.rs", "password": "A123456bb"}`,
			name:       "Service Error",
			mockErr:    errVals.NewServiceError(errVals.ErrInvalidPasswordCode, errVals.ErrInvalidPasswordsMatch),
			resp:       `{"errors":[{"code":"invalid_password","error":"password doesn't match with passwordConfirmation"}]}`,
			statusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mAuthSrv := authSrvMock.NewMockAuthServiceInterface(ctrl)
			mUsrSrv := usrSrvMock.NewMockUserServiceInterface(ctrl)
			handler := NewAuthHandler(mAuthSrv, mUsrSrv)

			mAuthSrv.EXPECT().Login(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)

			r := mux.NewRouter()
			r.HandleFunc(loginPath, handler.Login)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", loginPath, bytes.NewBufferString(test.req))

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
		mockErr      *errVals.ServiceError
		statusCode   int
		isValidation bool
		emptyCookie  bool
		noCookie     bool
	}{
		{
			name:       "Success",
			mockReturn: &models.AuthRespData{},
			statusCode: http.StatusOK,
			resp:       ``,
		},
		{
			name:       "Service Error",
			mockErr:    errVals.NewServiceError(errVals.ErrRedisClearCode, errors.New("some redis error")),
			statusCode: http.StatusInternalServerError,
			resp:       `{"errors":[{"code":"failed_delete_from_redis","error":"some redis error"}]}`,
		},
		{
			name:         "Validation Error",
			mockErr:      errVals.NewServiceError(errVals.ErrBrokenCookieCode, errVals.ErrBrokenCookie),
			statusCode:   http.StatusBadRequest,
			isValidation: true,
			emptyCookie:  true,
			resp:         `{"errors":[{"code":"auth_validation_error","error":"logout action: Invalid cookie err - broken cookie was given"}]}`,
		},
		{
			name:         "No cookie provided",
			statusCode:   http.StatusBadRequest,
			isValidation: true,
			noCookie:     true,
			resp:         `{"errors":[{"code":"auth_request_parse_error","error":"logout action: No cookie err - http: named cookie not present"}]}`,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mAuthSrv := authSrvMock.NewMockAuthServiceInterface(ctrl)
			mUsrSrv := usrSrvMock.NewMockUserServiceInterface(ctrl)
			handler := NewAuthHandler(mAuthSrv, mUsrSrv)

			r := mux.NewRouter()
			r.HandleFunc(logoutPath, handler.Logout)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", logoutPath, bytes.NewBufferString(""))

			if !test.isValidation {
				req.Header.Set("Cookie", "session_id=some_cookie")
				mAuthSrv.EXPECT().Logout(gomock.Any(), gomock.Any()).Return(test.mockErr)
			} else if test.emptyCookie {
				req.Header.Set("Cookie", "session_id=")
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, test.statusCode, w.Result().StatusCode)

			if test.resp != "" {
				assert.JSONEq(t, test.resp, w.Body.String())
			} else {
				assert.Equal(t, test.resp, w.Body.String())
			}
		})
	}
}

func TestDelivery_Session(t *testing.T) {
	tests := []struct {
		name       string
		mockReturn *models.SessionRespData
		mockErr    *errVals.ServiceError
		resp       string
		statusCode int
		noCookie   bool
	}{
		{
			name: "Success",
			mockReturn: &models.SessionRespData{
				UserData: models.User{
					ID:       1,
					Email:    "test@mail.ru",
					Username: "Tester",
				},
			},
			resp:       `{"user_data":{"id":1,"email":"test@mail.ru","username":"Tester","avatar_url":"","subscription_expiration_date":"", "subscription_status":false}}`,
			statusCode: http.StatusOK,
		},
		{
			name: "Service Error",
			mockErr: errVals.NewServiceError(
				errVals.ErrCreateUserCode,
				errors.New("cannot get cookie from redis"),
			),
			resp:       `{"errors":[{"code":"create_user_error","error":"cannot get cookie from redis"}]}`,
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "Forbidden",
			resp:       `{"errors":[{"code":"auth_request_parse_error","error":"session action: No cookie err - http: named cookie not present"}]}`,
			statusCode: http.StatusForbidden,
			noCookie:   true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mAuthSrv := authSrvMock.NewMockAuthServiceInterface(ctrl)
			mUsrSrv := usrSrvMock.NewMockUserServiceInterface(ctrl)
			handler := NewAuthHandler(mAuthSrv, mUsrSrv)

			r := mux.NewRouter()
			r.HandleFunc(sessionPath, handler.Session)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", sessionPath, bytes.NewBufferString(""))

			if !test.noCookie {
				req.Header.Set("Cookie", "session_id=some_cookie")
				mAuthSrv.EXPECT().Session(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)
			}

			r.ServeHTTP(w, req)

			assert.Equal(t, test.statusCode, w.Result().StatusCode)
			assert.JSONEq(t, test.resp, w.Body.String())
		})
	}
}
