package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	servAuthMock "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/service/mocks"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	servUserMock "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_Register(t *testing.T) {
	ctx := testContext(t)

	tests := []struct {
		name string
		args *struct {
			ctx          context.Context
			registerData *models.RegisterData
		}
		mockCreateUser   *models.User
		mockSetCookie    *models.CookieData
		mockUserErr      *errVals.ErrorObj
		mockCookieErr    *errVals.ErrorObj
		expectedResponse *models.AuthRespData
		expectedError    *models.ErrorRespData
		statusCode       int
		WithCookie       bool
	}{
		{
			name: "Success",
			args: &struct {
				ctx          context.Context
				registerData *models.RegisterData
			}{
				ctx: ctx,
				registerData: &models.RegisterData{
					Email:                "test@mail.ru",
					Username:             "tester",
					Password:             "test_password",
					PasswordConfirmation: "test_password",
				},
			},
			mockCreateUser: &models.User{
				ID:       1,
				Email:    "test@mail.ru",
				Username: "tester",
			},
			mockSetCookie: &models.CookieData{
				Name: "session_id",
				Token: &models.Token{
					TokenID: "some_cookie",
					UserID:  1,
				},
			},
			mockUserErr:   nil,
			mockCookieErr: nil,
			expectedResponse: &models.AuthRespData{
				StatusCode: 200,
				NewCookie: &models.CookieData{
					Name: "session_id",
					Token: &models.Token{
						TokenID: "some_cookie",
						UserID:  1,
					},
				},
			},
			expectedError: nil,
			statusCode:    200,
			WithCookie:    true,
		},
		{
			name: "User error",
			args: &struct {
				ctx          context.Context
				registerData *models.RegisterData
			}{
				ctx: ctx,
				registerData: &models.RegisterData{
					Email:                "test@mail.ru",
					Username:             "tester",
					Password:             "test_password",
					PasswordConfirmation: "test_password",
				},
			},
			mockUserErr:      &errVals.ErrorObj{Code: errVals.ErrCreateUserCode, Error: errVals.CustomError{Err: errors.New("cannot create user")}},
			mockCookieErr:    nil,
			expectedResponse: nil,
			expectedError: &models.ErrorRespData{
				StatusCode: 500,
				Errors:     []errVals.ErrorObj{{Code: errVals.ErrCreateUserCode, Error: errVals.CustomError{Err: errors.New("cannot create user")}}},
			},
			statusCode: 500,
		},
		{
			name: "Cookie error",
			args: &struct {
				ctx          context.Context
				registerData *models.RegisterData
			}{
				ctx: ctx,
				registerData: &models.RegisterData{
					Email:                "test@mail.ru",
					Username:             "tester",
					Password:             "test_password",
					PasswordConfirmation: "test_password",
				},
			},
			mockCreateUser: &models.User{
				ID:       1,
				Email:    "test@mail.ru",
				Username: "tester",
			},
			mockCookieErr: errVals.NewErrorObj(
				errVals.ErrCreateUserCode,
				errVals.CustomError{Err: fmt.Errorf("cannot set cookie into redis")},
			),
			expectedResponse: nil,
			expectedError: &models.ErrorRespData{
				StatusCode: 500,
				Errors: []errVals.ErrorObj{*errVals.NewErrorObj(
					errVals.ErrCreateUserCode,
					errVals.CustomError{Err: fmt.Errorf("cannot set cookie into redis")},
				)},
			},
			statusCode: 500,
			WithCookie: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mAuthRepo := servAuthMock.NewMockAuthRepositoryInterface(ctrl)
			mUsrRepo := servUserMock.NewMockUserRepositoryInterface(ctrl)
			s := NewAuthService(mAuthRepo, mUsrRepo)

			mUsrRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(test.mockCreateUser, test.mockUserErr, test.statusCode)

			if test.WithCookie {
				mAuthRepo.EXPECT().SetCookie(gomock.Any(), gomock.Any()).Return(test.mockSetCookie, test.mockCookieErr, test.statusCode)
			}

			response, err := s.Register(ctx, test.args.registerData)

			if test.expectedError != nil {
				assert.Nil(t, response)
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResponse, response)
				assert.Nil(t, err)
			}
		})
	}
}

func TestService_Session(t *testing.T) {
	tests := []struct {
		name string
		args *struct {
			ctx    context.Context
			cookie string
		}
		mockGetFromCookie    string
		mockGetUser          *models.User
		mockGetUserErr       *errVals.ErrorObj
		mockGetFromCookieErr *errVals.ErrorObj
		expectedResponse     *models.SessionRespData
		expectedError        *models.ErrorRespData
		statusCode           int
		WithGetUser          bool
	}{
		{
			name: "Success",
			args: &struct {
				ctx    context.Context
				cookie string
			}{
				ctx:    context.Background(),
				cookie: "some random cookie",
			},
			mockGetFromCookie:    "1",
			mockGetFromCookieErr: nil,
			mockGetUser: &models.User{
				ID:       1,
				Email:    "test@mail.ru",
				Username: "TestUser",
				Password: "secret_password",
			},
			mockGetUserErr: nil,
			expectedResponse: &models.SessionRespData{
				StatusCode: 200,
				UserData: models.User{
					ID:       1,
					Email:    "test@mail.ru",
					Username: "TestUser",
					Password: "secret_password",
				},
			},
			expectedError: nil,
			statusCode:    200,
			WithGetUser:   true,
		},
		{
			name: "Cookie error",
			args: &struct {
				ctx    context.Context
				cookie string
			}{
				ctx:    context.Background(),
				cookie: "some random cookie",
			},
			mockGetFromCookie: "",
			mockGetFromCookieErr: errVals.NewErrorObj(
				errVals.ErrCreateUserCode,
				errVals.CustomError{Err: fmt.Errorf("cannot get cookie from redis")},
			),
			expectedResponse: nil,
			expectedError: &models.ErrorRespData{
				StatusCode: http.StatusForbidden,
				Errors: []errVals.ErrorObj{*errVals.NewErrorObj(
					errVals.ErrCreateUserCode,
					errVals.CustomError{Err: fmt.Errorf("cannot get cookie from redis")},
				)},
			},
			statusCode: http.StatusForbidden,
		},
		{
			name: "User error",
			args: &struct {
				ctx    context.Context
				cookie string
			}{
				ctx:    context.Background(),
				cookie: "some random cookie",
			},
			mockGetFromCookie: "1",
			mockGetUserErr:    errVals.NewErrorObj(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFoundText),
			expectedResponse:  nil,
			expectedError: &models.ErrorRespData{
				StatusCode: http.StatusInternalServerError,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFoundText)},
			},
			statusCode:  http.StatusInternalServerError,
			WithGetUser: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mAuthRepo := servAuthMock.NewMockAuthRepositoryInterface(ctrl)
			mUsrRepo := servUserMock.NewMockUserRepositoryInterface(ctrl)
			s := NewAuthService(mAuthRepo, mUsrRepo)

			mAuthRepo.EXPECT().GetFromCookie(gomock.Any(), gomock.Any()).Return(test.mockGetFromCookie, test.mockGetFromCookieErr, test.statusCode)
			if test.WithGetUser {
				mUsrRepo.EXPECT().UserByID(gomock.Any(), gomock.Any()).Return(test.mockGetUser, test.mockGetUserErr, test.statusCode)
			}

			response, err := s.Session(test.args.ctx, test.args.cookie)

			if test.expectedError != nil {
				assert.Nil(t, response)
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResponse, response)
				assert.Nil(t, err)
			}
		})
	}
}

func TestService_Login(t *testing.T) {
	type args struct {
		ctx       context.Context
		loginData *models.LoginData
	}

	ctx := testContext(t)

	tests := []struct {
		name                  string
		args                  *args
		mockUser              *models.User
		mockSetCookie         *models.CookieData
		mockUserErr           *errVals.ErrorObj
		mockDestroySessionErr *errVals.ErrorObj
		mockSetCookieErr      *errVals.ErrorObj
		expectedResponse      *models.AuthRespData
		expectedError         *models.ErrorRespData
		statusCode            int
		withCookieDestruction bool
		withCookieSetting     bool
	}{
		{
			name: "Success",
			args: &args{
				ctx: ctx,
				loginData: &models.LoginData{
					Email:    "test@mail.ru",
					Password: "A123456bb",
					Cookie:   "some_cookie",
				},
			},
			mockUser: &models.User{
				ID:       1,
				Email:    "test@mail.ru",
				Password: "$2a$10$wfvAfweY9mrak.zBcnvY1eneItl0nWftZiH0/HH5IK5l/6LgC/fpe",
				Username: "test",
			},
			mockUserErr: nil,
			mockSetCookie: &models.CookieData{
				Name: "session_id",
				Token: &models.Token{
					TokenID: "new_cookie",
					UserID:  1,
				},
			},
			expectedResponse: &models.AuthRespData{
				NewCookie: &models.CookieData{
					Name: "session_id",
					Token: &models.Token{
						TokenID: "new_cookie",
						UserID:  1,
					},
				},
				StatusCode: 200,
			},
			expectedError:         nil,
			statusCode:            200,
			withCookieDestruction: true,
			withCookieSetting:     true,
		},
		{
			name: "Failed to destroy session",
			args: &args{
				ctx: ctx,
				loginData: &models.LoginData{
					Email:    "test@mail.ru",
					Password: "A123456bb",
					Cookie:   "some_cookie",
				},
			},
			mockUser: &models.User{
				ID:       1,
				Email:    "test@mail.ru",
				Password: "$2a$10$wfvAfweY9mrak.zBcnvY1eneItl0nWftZiH0/HH5IK5l/6LgC/fpe",
				Username: "test",
			},
			mockUserErr:           nil,
			mockDestroySessionErr: errVals.NewErrorObj(errVals.ErrRedisClearCode, errVals.CustomError{Err: errors.New("some err")}),
			expectedResponse:      nil,
			expectedError: &models.ErrorRespData{
				StatusCode: 500,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrRedisClearCode, errVals.CustomError{Err: errors.New("some err")})},
			},
			statusCode:            500,
			withCookieDestruction: true,
		},
		{
			name: "Failed to set new cookie",
			args: &args{
				ctx: ctx,
				loginData: &models.LoginData{
					Email:    "test@mail.ru",
					Password: "A123456bb",
					Cookie:   "some_cookie",
				},
			},
			mockUser: &models.User{
				ID:       1,
				Email:    "test@mail.ru",
				Password: "$2a$10$wfvAfweY9mrak.zBcnvY1eneItl0nWftZiH0/HH5IK5l/6LgC/fpe",
				Username: "test",
			},
			mockUserErr: nil,
			mockSetCookieErr: errVals.NewErrorObj(
				errVals.ErrCreateUserCode,
				errVals.CustomError{Err: fmt.Errorf("cannot set cookie into redis")}),
			expectedResponse: nil,
			expectedError: &models.ErrorRespData{
				StatusCode: 500,
				Errors: []errVals.ErrorObj{*errVals.NewErrorObj(
					errVals.ErrCreateUserCode,
					errVals.CustomError{Err: fmt.Errorf("cannot set cookie into redis")})},
			},
			statusCode:            500,
			withCookieDestruction: true,
			withCookieSetting:     true,
		},
		{
			name: "Failed to get user",
			args: &args{
				ctx: ctx,
				loginData: &models.LoginData{
					Email:    "test@mail.ru",
					Password: "A123456bb",
					Cookie:   "some_cookie",
				},
			},
			mockUserErr:      errVals.NewErrorObj(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFoundText),
			expectedResponse: nil,
			expectedError: &models.ErrorRespData{
				StatusCode: 404,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFoundText)},
			},
			statusCode: 404,
		},
		{
			name: "Wrong password error",
			args: &args{
				ctx: ctx,
				loginData: &models.LoginData{
					Email:    "test@mail.ru",
					Password: "some different password",
					Cookie:   "some_cookie",
				},
			},
			mockUser: &models.User{
				ID:       1,
				Email:    "test@mail.ru",
				Password: "$2a$10$wfvAfweY9mrak.zBcnvY1eneItl0nWftZiH0/HH5IK5l/6LgC/fpe",
				Username: "test",
			},
			mockUserErr:      nil,
			expectedResponse: nil,
			expectedError: &models.ErrorRespData{
				StatusCode: 409,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrInvalidPasswordCode, errVals.ErrInvalidPasswordsMatchText)},
			},
			statusCode: 409,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mAuthRepo := servAuthMock.NewMockAuthRepositoryInterface(ctrl)
			mUsrRepo := servUserMock.NewMockUserRepositoryInterface(ctrl)
			s := NewAuthService(mAuthRepo, mUsrRepo)

			mUsrRepo.EXPECT().UserByEmail(gomock.Any(), gomock.Any()).Return(test.mockUser, test.mockUserErr, test.statusCode)
			if test.withCookieDestruction {
				mAuthRepo.EXPECT().DestroySession(gomock.Any(), gomock.Any()).Return(test.mockDestroySessionErr, test.statusCode)
			}

			if test.withCookieSetting {
				mAuthRepo.EXPECT().SetCookie(gomock.Any(), gomock.Any()).Return(test.mockSetCookie, test.mockSetCookieErr, test.statusCode)
			}

			response, err := s.Login(test.args.ctx, test.args.loginData)

			if test.expectedError != nil {
				assert.Nil(t, response)
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResponse, response)
				assert.Nil(t, err)
			}
		})
	}
}

func TestService_Logout(t *testing.T) {
	type args struct {
		ctx    context.Context
		cookie string
	}

	ctx := testContext(t)

	tests := []struct {
		name                  string
		args                  *args
		mockDestroySessionErr *errVals.ErrorObj
		expectedResponse      *models.AuthRespData
		expectedError         *models.ErrorRespData
		statusCode            int
	}{
		{
			name: "Success",
			args: &args{
				ctx:    ctx,
				cookie: "some cookie",
			},
			expectedResponse: &models.AuthRespData{
				StatusCode: 200,
			},
			statusCode: 200,
		},
		{
			name: "Destroy session error",
			args: &args{
				ctx:    ctx,
				cookie: "some cookie",
			},
			mockDestroySessionErr: errVals.NewErrorObj(errVals.ErrRedisClearCode, errVals.CustomError{Err: errors.New("some redis error")}),
			expectedError: &models.ErrorRespData{
				StatusCode: 500,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrRedisClearCode, errVals.CustomError{Err: errors.New("some redis error")})},
			},
			statusCode: 500,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mAuthRepo := servAuthMock.NewMockAuthRepositoryInterface(ctrl)
			mUsrRepo := servUserMock.NewMockUserRepositoryInterface(ctrl)
			s := NewAuthService(mAuthRepo, mUsrRepo)

			mAuthRepo.EXPECT().DestroySession(gomock.Any(), gomock.Any()).Return(test.mockDestroySessionErr, test.statusCode)

			response, err := s.Logout(test.args.ctx, test.args.cookie)

			if test.expectedError != nil {
				assert.Nil(t, response)
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResponse, response)
				assert.Nil(t, err)
			}
		})
	}
}

func testContext(t *testing.T) context.Context {
	require.NoError(t, os.Chdir("../../../.."), "failed to change directory")

	cfg, err := config.New(true)
	require.NoError(t, err, "failed to read config from auth service_test")

	return config.WrapRedisContext(context.Background(), &cfg.Databases.Redis)
}
