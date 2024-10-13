package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	servMock "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/service/mocks"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_Register(t *testing.T) {
	err := os.Chdir("../../../..")
	if err != nil {
		t.Fatalf("failed to change directory: %v", err)
	}

	cfg, _ := config.New(false, nil)
	ctx := config.WrapContext(context.Background(), cfg)

	tests := []struct {
		name string
		args *struct {
			ctx          context.Context
			registerData *authModels.RegisterData
		}
		mockCreateUser   *models.User
		mockSetCookie    *authModels.CookieData
		mockUserErr      *errVals.ErrorObj
		mockCookieErr    *errVals.ErrorObj
		expectedResponse *authModels.AuthResponse
		expectedError    *models.ErrorResponse
		statusCode       int
		isValidation     bool
		WithCookie       bool
	}{
		{
			name: "Success",
			args: &struct {
				ctx          context.Context
				registerData *authModels.RegisterData
			}{
				ctx: ctx,
				registerData: &authModels.RegisterData{
					Email:                "test@mail.ru",
					Username:             "tester",
					Password:             "test_password",
					PasswordConfirmation: "test_password",
				},
			},
			mockCreateUser: &models.User{
				Id:       1,
				Email:    "test@mail.ru",
				Username: "tester",
			},
			mockSetCookie: &authModels.CookieData{
				Name:   "session_id",
				Value:  "some_cookie",
				UserID: 1,
			},
			mockUserErr:   nil,
			mockCookieErr: nil,
			expectedResponse: &authModels.AuthResponse{
				Success:    true,
				StatusCode: 200,
				NewCookie: &authModels.CookieData{
					Name:   "session_id",
					Value:  "some_cookie",
					UserID: 1,
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
				registerData *authModels.RegisterData
			}{
				ctx: ctx,
				registerData: &authModels.RegisterData{
					Email:                "test@mail.ru",
					Username:             "tester",
					Password:             "test_password",
					PasswordConfirmation: "test_password",
				},
			},
			mockUserErr:      &errVals.ErrorObj{Code: errVals.ErrCreateUserCode, Error: errVals.CustomError{Err: errors.New("cannot create user")}},
			mockCookieErr:    nil,
			expectedResponse: nil,
			expectedError: &models.ErrorResponse{
				Success:    false,
				StatusCode: 500,
				Errors:     []errVals.ErrorObj{{Code: errVals.ErrCreateUserCode, Error: errVals.CustomError{Err: errors.New("cannot create user")}}},
			},
			statusCode: 500,
		},
		{
			name: "Cookie error",
			args: &struct {
				ctx          context.Context
				registerData *authModels.RegisterData
			}{
				ctx: ctx,
				registerData: &authModels.RegisterData{
					Email:                "test@mail.ru",
					Username:             "tester",
					Password:             "test_password",
					PasswordConfirmation: "test_password",
				},
			},
			mockCreateUser: &models.User{
				Id:       1,
				Email:    "test@mail.ru",
				Username: "tester",
			},
			mockCookieErr: errVals.NewErrorObj(
				errVals.ErrCreateUserCode,
				errVals.CustomError{Err: fmt.Errorf("cannot set cookie into redis: %w", err)},
			),
			expectedResponse: nil,
			expectedError: &models.ErrorResponse{
				Success:    false,
				StatusCode: 500,
				Errors: []errVals.ErrorObj{*errVals.NewErrorObj(
					errVals.ErrCreateUserCode,
					errVals.CustomError{Err: fmt.Errorf("cannot set cookie into redis: %w", err)},
				)},
			},
			statusCode: 500,
			WithCookie: true,
		},
		{
			name: "Validation Error",
			args: &struct {
				ctx          context.Context
				registerData *authModels.RegisterData
			}{
				ctx:          ctx,
				registerData: &authModels.RegisterData{},
			},
			mockUserErr:   nil,
			mockCookieErr: nil,
			expectedError: &models.ErrorResponse{
				Success:    false,
				StatusCode: 422,
				Errors: []errVals.ErrorObj{
					{
						Code: errVals.ErrInvalidPasswordCode, Error: errVals.CustomError{Err: errVals.ErrInvalidPasswordText.Err},
					},
					{
						Code: errVals.ErrInvalidEmailCode, Error: errVals.CustomError{Err: errVals.ErrInvalidEmailText.Err},
					},
				},
			},
			isValidation: true,
			statusCode:   422,
		},
		{
			name: "Password missmatch",
			args: &struct {
				ctx          context.Context
				registerData *authModels.RegisterData
			}{
				ctx: ctx,
				registerData: &authModels.RegisterData{
					Email:                "test@mail.ru",
					Username:             "tester",
					Password:             "test_password",
					PasswordConfirmation: "test_password_wrong",
				},
			},
			mockUserErr:   nil,
			mockCookieErr: nil,
			expectedError: &models.ErrorResponse{
				Success:    false,
				StatusCode: 422,
				Errors: []errVals.ErrorObj{
					{Code: errVals.ErrInvalidPasswordCode, Error: errVals.CustomError{Err: errVals.ErrInvalidPasswordsMatchText.Err}},
				},
			},
			isValidation: true,
			statusCode:   422,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := servMock.NewMockAuthRepositoryInterface(ctrl)
			s := NewService(repo)

			if !test.isValidation {
				repo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(test.mockCreateUser, test.mockUserErr, test.statusCode)

				if test.WithCookie {
					repo.EXPECT().SetCookie(gomock.Any(), gomock.Any()).Return(test.mockSetCookie, test.mockCookieErr, test.statusCode)
				}
			}

			t.Parallel()

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
		expectedResponse     *authModels.SessionResponse
		expectedError        *models.ErrorResponse
		statusCode           int
		isValidation         bool
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
				Id:       1,
				Email:    "test@mail.ru",
				Username: "TestUser",
				Password: "secret_password",
			},
			mockGetUserErr: nil,
			expectedResponse: &authModels.SessionResponse{
				Success:    true,
				StatusCode: 200,
				UserData: models.User{
					Id:       1,
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
			expectedError: &models.ErrorResponse{
				Success:    false,
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
			expectedError: &models.ErrorResponse{
				Success:    false,
				StatusCode: http.StatusInternalServerError,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFoundText)},
			},
			statusCode:  http.StatusInternalServerError,
			WithGetUser: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := servMock.NewMockAuthRepositoryInterface(ctrl)
			s := NewService(repo)

			if !test.isValidation {
				repo.EXPECT().GetFromCookie(gomock.Any(), gomock.Any()).Return(test.mockGetFromCookie, test.mockGetFromCookieErr, test.statusCode)

				if test.WithGetUser {
					repo.EXPECT().UserById(gomock.Any(), gomock.Any()).Return(test.mockGetUser, test.mockGetUserErr, test.statusCode)
				}
			}

			t.Parallel()

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
		loginData *authModels.LoginData
	}

	err := os.Chdir("../../../..")
	if err != nil {
		t.Fatalf("failed to change directory: %v", err)
	}

	cfg, _ := config.New(false, nil)
	ctx := config.WrapContext(context.Background(), cfg)

	tests := []struct {
		name                  string
		args                  *args
		mockUser              *models.User
		mockSetCookie         *authModels.CookieData
		mockUserErr           *errVals.ErrorObj
		mockDestroySessionErr *errVals.ErrorObj
		mockSetCookieErr      *errVals.ErrorObj
		expectedResponse      *authModels.AuthResponse
		expectedError         *models.ErrorResponse
		isValidation          bool
		statusCode            int
		withCookieDestruction bool
		withCookieSetting     bool
	}{
		{
			name: "Success",
			args: &args{
				ctx: ctx,
				loginData: &authModels.LoginData{
					Email:    "test@mail.ru",
					Password: "A123456bb",
					Cookie:   "some_cookie",
				},
			},
			mockUser: &models.User{
				Id:       1,
				Email:    "test@mail.ru",
				Password: "$2a$10$wfvAfweY9mrak.zBcnvY1eneItl0nWftZiH0/HH5IK5l/6LgC/fpe",
				Username: "test",
			},
			mockUserErr: nil,
			mockSetCookie: &authModels.CookieData{
				Name:   "session_id",
				Value:  "new_cookie",
				UserID: 1,
			},
			expectedResponse: &authModels.AuthResponse{
				Success: true,
				NewCookie: &authModels.CookieData{
					Name:   "session_id",
					Value:  "new_cookie",
					UserID: 1,
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
				loginData: &authModels.LoginData{
					Email:    "test@mail.ru",
					Password: "A123456bb",
					Cookie:   "some_cookie",
				},
			},
			mockUser: &models.User{
				Id:       1,
				Email:    "test@mail.ru",
				Password: "$2a$10$wfvAfweY9mrak.zBcnvY1eneItl0nWftZiH0/HH5IK5l/6LgC/fpe",
				Username: "test",
			},
			mockUserErr:           nil,
			mockDestroySessionErr: errVals.NewErrorObj(errVals.ErrRedisClearCode, errVals.CustomError{Err: err}),
			expectedResponse:      nil,
			expectedError: &models.ErrorResponse{
				Success:    false,
				StatusCode: 500,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrRedisClearCode, errVals.CustomError{Err: err})},
			},
			statusCode:            500,
			withCookieDestruction: true,
		},
		{
			name: "Failed to set new cookie",
			args: &args{
				ctx: ctx,
				loginData: &authModels.LoginData{
					Email:    "test@mail.ru",
					Password: "A123456bb",
					Cookie:   "some_cookie",
				},
			},
			mockUser: &models.User{
				Id:       1,
				Email:    "test@mail.ru",
				Password: "$2a$10$wfvAfweY9mrak.zBcnvY1eneItl0nWftZiH0/HH5IK5l/6LgC/fpe",
				Username: "test",
			},
			mockUserErr: nil,
			mockSetCookieErr: errVals.NewErrorObj(
				errVals.ErrCreateUserCode,
				errVals.CustomError{Err: fmt.Errorf("cannot set cookie into redis: %w", err)}),
			expectedResponse: nil,
			expectedError: &models.ErrorResponse{
				Success:    false,
				StatusCode: 500,
				Errors: []errVals.ErrorObj{*errVals.NewErrorObj(
					errVals.ErrCreateUserCode,
					errVals.CustomError{Err: fmt.Errorf("cannot set cookie into redis: %w", err)})},
			},
			statusCode:            500,
			withCookieDestruction: true,
			withCookieSetting:     true,
		},
		{
			name: "Failed to get user",
			args: &args{
				ctx: ctx,
				loginData: &authModels.LoginData{
					Email:    "test@mail.ru",
					Password: "A123456bb",
					Cookie:   "some_cookie",
				},
			},
			mockUserErr:      errVals.NewErrorObj(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFoundText),
			expectedResponse: nil,
			expectedError: &models.ErrorResponse{
				Success:    false,
				StatusCode: 404,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFoundText)},
			},
			statusCode: 404,
		},
		{
			name: "Wrong password error",
			args: &args{
				ctx: ctx,
				loginData: &authModels.LoginData{
					Email:    "test@mail.ru",
					Password: "some different password",
					Cookie:   "some_cookie",
				},
			},
			mockUser: &models.User{
				Id:       1,
				Email:    "test@mail.ru",
				Password: "$2a$10$wfvAfweY9mrak.zBcnvY1eneItl0nWftZiH0/HH5IK5l/6LgC/fpe",
				Username: "test",
			},
			mockUserErr:      nil,
			expectedResponse: nil,
			expectedError: &models.ErrorResponse{
				Success:    false,
				StatusCode: 409,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrInvalidPasswordCode, errVals.ErrInvalidPasswordsMatchText)},
			},
			statusCode: 409,
		},
		{
			name: "Validation error",
			args: &args{
				ctx: ctx,
				loginData: &authModels.LoginData{
					Email:    "test@mail.ru",
					Password: "some different password",
					Cookie:   "",
				},
			},
			expectedError: &models.ErrorResponse{
				Success:    false,
				StatusCode: 400,
				Errors:     []errVals.ErrorObj{{Code: errVals.ErrBrokenCookieCode, Error: errVals.ErrBrokenCookieText}},
			},
			statusCode:   400,
			isValidation: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := servMock.NewMockAuthRepositoryInterface(ctrl)
			s := NewService(repo)

			t.Parallel()

			if !test.isValidation {
				repo.EXPECT().UserByEmail(gomock.Any(), gomock.Any()).Return(test.mockUser, test.mockUserErr, test.statusCode)
				if test.withCookieDestruction {
					repo.EXPECT().DestroySession(gomock.Any(), gomock.Any()).Return(test.mockDestroySessionErr, test.statusCode)
				}

				if test.withCookieSetting {
					repo.EXPECT().SetCookie(gomock.Any(), gomock.Any()).Return(test.mockSetCookie, test.mockSetCookieErr, test.statusCode)
				}
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

	err := os.Chdir("../../../..")
	if err != nil {
		t.Fatalf("failed to change directory: %v", err)
	}

	cfg, _ := config.New(false, nil)
	ctx := config.WrapContext(context.Background(), cfg)

	tests := []struct {
		name                  string
		args                  *args
		mockDestroySessionErr *errVals.ErrorObj
		expectedResponse      *authModels.AuthResponse
		expectedError         *models.ErrorResponse
		statusCode            int
	}{
		{
			name: "Success",
			args: &args{
				ctx:    ctx,
				cookie: "some cookie",
			},
			expectedResponse: &authModels.AuthResponse{
				Success:    true,
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
			expectedError: &models.ErrorResponse{
				Success:    false,
				StatusCode: 500,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrRedisClearCode, errVals.CustomError{Err: errors.New("some redis error")})},
			},
			statusCode: 500,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := servMock.NewMockAuthRepositoryInterface(ctrl)
			s := NewService(repo)

			repo.EXPECT().DestroySession(gomock.Any(), gomock.Any()).Return(test.mockDestroySessionErr, test.statusCode)

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
