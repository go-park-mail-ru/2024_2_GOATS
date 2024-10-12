package service

import (
	"context"
	"errors"
	"net/http"
	"testing"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	servMock "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_GetCollection(t *testing.T) {
	tests := []struct {
		name             string
		mockReturn       []models.Collection
		mockErr          *errVals.ErrorObj
		expectedResponse *models.CollectionsResponse
		expectedError    *models.ErrorResponse
		statusCode       int
	}{
		{
			name: "Success",
			mockReturn: []models.Collection{
				{Id: 1, Title: "Collection 1", Movies: []*models.Movie{}},
				{Id: 2, Title: "Collection 2", Movies: []*models.Movie{}},
			},
			mockErr: nil,
			expectedResponse: &models.CollectionsResponse{
				Collections: []models.Collection{
					{Id: 1, Title: "Collection 1", Movies: []*models.Movie{}},
					{Id: 2, Title: "Collection 2", Movies: []*models.Movie{}},
				},
				StatusCode: 200,
				Success:    true,
			},
			expectedError: nil,
			statusCode:    200,
		},
		{
			name:             "Error",
			mockReturn:       nil,
			mockErr:          &errVals.ErrorObj{Code: "something_went_wrong", Error: errVals.CustomError{Err: errors.New("Database fail")}},
			expectedResponse: nil,
			expectedError: &models.ErrorResponse{
				Success:    false,
				StatusCode: http.StatusUnprocessableEntity,
				Errors:     []errVals.ErrorObj{{Code: "something_went_wrong", Error: errVals.CustomError{Err: errors.New("Database fail")}}},
			},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := servMock.NewMockRepositoryInterface(ctrl)
			s := &Service{repository: repo}

			repo.EXPECT().GetCollection(gomock.Any()).Return(test.mockReturn, test.mockErr, test.statusCode)

			// Параллельное выполнение теста
			t.Parallel()

			response, err := s.GetCollection(context.Background())

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

func TestService_Register(t *testing.T) {
	tests := []struct {
		name string
		args *struct {
			ctx          context.Context
			registerData *authModels.RegisterData
		}
		mockReturn       *authModels.CookieData
		mockErr          *errVals.ErrorObj
		expectedResponse *authModels.AuthResponse
		expectedError    *models.ErrorResponse
		statusCode       int
		isValidation     bool
	}{
		{
			name: "Success",
			args: &struct {
				ctx          context.Context
				registerData *authModels.RegisterData
			}{
				ctx: context.Background(),
				registerData: &authModels.RegisterData{
					Email:                "test@mail.ru",
					Username:             "tester",
					Password:             "test_password",
					PasswordConfirmation: "test_password",
				},
			},
			mockReturn: &authModels.CookieData{
				Name:   "session_id",
				Value:  "test cookie value",
				UserID: 1,
			},
			mockErr: nil,
			expectedResponse: &authModels.AuthResponse{
				Success:    true,
				StatusCode: 200,
				NewCookie: &authModels.CookieData{
					Name:   "session_id",
					Value:  "test cookie value",
					UserID: 1,
				},
				ExpCookie: nil,
			},
			expectedError: nil,
			statusCode:    200,
		},
		{
			name: "Repo error",
			args: &struct {
				ctx          context.Context
				registerData *authModels.RegisterData
			}{
				ctx: context.Background(),
				registerData: &authModels.RegisterData{
					Email:                "test@mail.ru",
					Username:             "tester",
					Password:             "test_password",
					PasswordConfirmation: "test_password",
				},
			},
			mockReturn:       nil,
			mockErr:          &errVals.ErrorObj{Code: errVals.ErrCreateUserCode, Error: errVals.CustomError{Err: errors.New("cannot create user")}},
			expectedResponse: nil,
			expectedError: &models.ErrorResponse{
				Success:    false,
				StatusCode: 409,
				Errors:     []errVals.ErrorObj{{Code: errVals.ErrCreateUserCode, Error: errVals.CustomError{Err: errors.New("cannot create user")}}},
			},
			statusCode: 409,
		},
		{
			name: "Validation Error",
			args: &struct {
				ctx          context.Context
				registerData *authModels.RegisterData
			}{
				ctx:          context.Background(),
				registerData: &authModels.RegisterData{},
			},
			mockReturn:       nil,
			mockErr:          nil,
			expectedResponse: nil,
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
				ctx: context.Background(),
				registerData: &authModels.RegisterData{
					Email:                "test@mail.ru",
					Username:             "tester",
					Password:             "test_password",
					PasswordConfirmation: "test_password_wrong",
				},
			},
			mockReturn:       nil,
			mockErr:          nil,
			expectedResponse: nil,
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

			repo := servMock.NewMockRepositoryInterface(ctrl)
			s := &Service{repository: repo}

			if !test.isValidation {
				repo.EXPECT().Register(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr, test.statusCode)
			}

			t.Parallel()

			response, err := s.Register(test.args.ctx, test.args.registerData)

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
		mockReturn       *models.User
		mockErr          *errVals.ErrorObj
		expectedResponse *authModels.SessionResponse
		expectedError    *models.ErrorResponse
		statusCode       int
		isValidation     bool
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
			mockReturn: &models.User{
				Id:       1,
				Email:    "test@mail.ru",
				Username: "TestUser",
				Password: "secret_password",
			},
			mockErr: nil,
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
		},
		{
			name: "Repo error",
			args: &struct {
				ctx    context.Context
				cookie string
			}{
				ctx:    context.Background(),
				cookie: "some random cookie",
			},
			mockReturn:       nil,
			mockErr:          &errVals.ErrorObj{errVals.ErrUnauthorizedCode, errVals.CustomError{Err: errors.New("expired cookie")}},
			expectedResponse: nil,
			expectedError: &models.ErrorResponse{
				Success:    false,
				StatusCode: http.StatusForbidden,
				Errors:     []errVals.ErrorObj{{errVals.ErrUnauthorizedCode, errVals.CustomError{Err: errors.New("expired cookie")}}},
			},
			statusCode: http.StatusForbidden,
		},
		{
			name: "Validation error",
			args: &struct {
				ctx    context.Context
				cookie string
			}{
				ctx:    context.Background(),
				cookie: "",
			},
			mockReturn:       nil,
			mockErr:          nil,
			expectedResponse: nil,
			expectedError: &models.ErrorResponse{
				Success:    false,
				StatusCode: http.StatusForbidden,
				Errors:     []errVals.ErrorObj{{errVals.ErrBrokenCookieCode, errVals.CustomError{Err: errVals.ErrBrokenCookieText.Err}}},
			},
			statusCode:   http.StatusForbidden,
			isValidation: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := servMock.NewMockRepositoryInterface(ctrl)
			s := &Service{repository: repo}

			if !test.isValidation {
				repo.EXPECT().Session(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr, test.statusCode)
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

	tests := []struct {
		name             string
		args             *args
		mockReturn       []*authModels.CookieData
		mockErr          *errVals.ErrorObj
		expectedResponse *authModels.AuthResponse
		expectedError    *models.ErrorResponse
		statusCode       int
		isValidation     bool
	}{
		{
			name: "Success",
			args: &args{
				ctx: context.Background(),
				loginData: &authModels.LoginData{
					Email:    "test@mail.ru",
					Password: "test_password",
					Cookie:   "some_cookie",
				},
			},
			mockReturn: []*authModels.CookieData{
				&authModels.CookieData{
					Name:   "session_id",
					Value:  "exp_cookie",
					UserID: 1,
				},
				&authModels.CookieData{
					Name:   "session_id",
					Value:  "new_cookie",
					UserID: 1,
				},
			},
			mockErr: nil,
			expectedResponse: &authModels.AuthResponse{
				Success: true,
				NewCookie: &authModels.CookieData{
					Name:   "session_id",
					Value:  "new_cookie",
					UserID: 1,
				},
				ExpCookie: &authModels.CookieData{
					Name:   "session_id",
					Value:  "exp_cookie",
					UserID: 1,
				},
				StatusCode: 200,
			},
			expectedError: nil,
			statusCode:    200,
		},
		{
			name: "Repo error",
			args: &args{
				ctx: context.Background(),
				loginData: &authModels.LoginData{
					Email:    "test@mail.ru",
					Password: "test_password",
					Cookie:   "some_cookie",
				},
			},
			mockReturn:       nil,
			mockErr:          errVals.NewErrorObj(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFoundText),
			expectedResponse: nil,
			expectedError: &models.ErrorResponse{
				Success:    false,
				StatusCode: 404,
				Errors:     []errVals.ErrorObj{{errVals.ErrUserNotFoundCode, errVals.CustomError{Err: errVals.ErrUserNotFoundText.Err}}},
			},
			statusCode: 404,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := servMock.NewMockRepositoryInterface(ctrl)
			s := &Service{repository: repo}

			t.Parallel()

			repo.EXPECT().Login(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr, test.statusCode)
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
