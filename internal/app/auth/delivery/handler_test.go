package delivery

import (
	"context"
	"errors"
	"testing"

	srvMock "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery/mocks"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDelivery_Register(t *testing.T) {
	tests := []struct {
		name             string
		registerData     *authModels.RegisterData
		mockReturn       *authModels.AuthResponse
		mockErr          *models.ErrorResponse
		expectedResponse *authModels.AuthResponse
		expectedErr      *models.ErrorResponse
	}{
		{
			name: "Success",
			registerData: &authModels.RegisterData{
				Email:                "test@mail.ru",
				Username:             "tester",
				Password:             "some pass",
				PasswordConfirmation: "some pass",
			},
			mockReturn: &authModels.AuthResponse{
				Success: true,
				NewCookie: &authModels.CookieData{
					Name:   "session_id",
					Value:  "cookie value",
					UserID: 1,
				},
				StatusCode: 200,
			},
			expectedResponse: &authModels.AuthResponse{
				Success: true,
				NewCookie: &authModels.CookieData{
					Name:   "session_id",
					Value:  "cookie value",
					UserID: 1,
				},
				StatusCode: 200,
			},
		},
		{
			name: "Service Error",
			registerData: &authModels.RegisterData{
				Email:    "test@mail.ru",
				Username: "tester",
				Password: "some pass",
			},
			mockErr: &models.ErrorResponse{
				Success:    false,
				StatusCode: 500,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrGenerateTokenCode, errVals.CustomError{Err: errors.New("some token error")})},
			},
			expectedErr: &models.ErrorResponse{
				Success:    false,
				StatusCode: 500,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrGenerateTokenCode, errVals.CustomError{Err: errors.New("some token error")})},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			srv := srvMock.NewMockAuthServiceInterface(ctrl)
			imp := NewImplementation(context.Background(), srv)

			srv.EXPECT().Register(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)
			resp, err := imp.Register(context.Background(), test.registerData)

			if test.expectedErr != nil {
				assert.Nil(t, resp)
				assert.Equal(t, test.expectedErr, err)
			} else {
				assert.Equal(t, test.expectedResponse, resp)
				assert.Nil(t, err)
			}
		})
	}
}

func TestDelivery_Login(t *testing.T) {
	tests := []struct {
		name             string
		loginData        *authModels.LoginData
		mockReturn       *authModels.AuthResponse
		mockErr          *models.ErrorResponse
		expectedResponse *authModels.AuthResponse
		expectedErr      *models.ErrorResponse
	}{
		{
			name: "Success",
			loginData: &authModels.LoginData{
				Email:    "test@mail.ru",
				Password: "some pass",
				Cookie:   "some cookie",
			},
			mockReturn: &authModels.AuthResponse{
				Success: true,
				NewCookie: &authModels.CookieData{
					Name:   "session_id",
					Value:  "cookie value",
					UserID: 1,
				},
				StatusCode: 200,
			},
			expectedResponse: &authModels.AuthResponse{
				Success: true,
				NewCookie: &authModels.CookieData{
					Name:   "session_id",
					Value:  "cookie value",
					UserID: 1,
				},
				StatusCode: 200,
			},
		},
		{
			name: "Service Error",
			loginData: &authModels.LoginData{
				Email:    "test@mail.ru",
				Password: "incorrect some pass",
			},
			mockErr: &models.ErrorResponse{
				Success:    false,
				StatusCode: 500,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrInvalidPasswordCode, errVals.ErrInvalidPasswordsMatchText)},
			},
			expectedErr: &models.ErrorResponse{
				Success:    false,
				StatusCode: 500,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrInvalidPasswordCode, errVals.ErrInvalidPasswordsMatchText)},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			srv := srvMock.NewMockAuthServiceInterface(ctrl)
			imp := NewImplementation(context.Background(), srv)

			srv.EXPECT().Login(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)
			resp, err := imp.Login(context.Background(), test.loginData)

			if test.expectedErr != nil {
				assert.Nil(t, resp)
				assert.Equal(t, test.expectedErr, err)
			} else {
				assert.Equal(t, test.expectedResponse, resp)
				assert.Nil(t, err)
			}
		})
	}
}

func TestDelivery_Logout(t *testing.T) {
	tests := []struct {
		name             string
		cookie           string
		mockReturn       *authModels.AuthResponse
		mockErr          *models.ErrorResponse
		expectedResponse *authModels.AuthResponse
		expectedErr      *models.ErrorResponse
	}{
		{
			name:   "Success",
			cookie: "some cookie",
			mockReturn: &authModels.AuthResponse{
				Success:    true,
				StatusCode: 200,
			},
			expectedResponse: &authModels.AuthResponse{
				Success:    true,
				StatusCode: 200,
			},
		},
		{
			name: "Service Error",
			mockErr: &models.ErrorResponse{
				Success:    false,
				StatusCode: 500,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrRedisClearCode, errVals.CustomError{Err: errors.New("some redis error")})},
			},
			expectedErr: &models.ErrorResponse{
				Success:    false,
				StatusCode: 500,
				Errors:     []errVals.ErrorObj{*errVals.NewErrorObj(errVals.ErrRedisClearCode, errVals.CustomError{Err: errors.New("some redis error")})},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			srv := srvMock.NewMockAuthServiceInterface(ctrl)
			imp := NewImplementation(context.Background(), srv)

			srv.EXPECT().Logout(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)
			resp, err := imp.Logout(context.Background(), test.cookie)

			if test.expectedErr != nil {
				assert.Nil(t, resp)
				assert.Equal(t, test.expectedErr, err)
			} else {
				assert.Equal(t, test.expectedResponse, resp)
				assert.Nil(t, err)
			}
		})
	}
}

func TestDelivery_Session(t *testing.T) {
	tests := []struct {
		name             string
		cookie           string
		mockReturn       *authModels.SessionResponse
		mockErr          *models.ErrorResponse
		expectedResponse *authModels.SessionResponse
		expectedErr      *models.ErrorResponse
	}{
		{
			name:   "Success",
			cookie: "some cookie",
			mockReturn: &authModels.SessionResponse{
				Success: true,
				UserData: models.User{
					Id:       1,
					Email:    "test@mail.ru",
					Username: "Tester",
				},
				StatusCode: 200,
			},
			expectedResponse: &authModels.SessionResponse{
				Success: true,
				UserData: models.User{
					Id:       1,
					Email:    "test@mail.ru",
					Username: "Tester",
				},
				StatusCode: 200,
			},
		},
		{
			name: "Service Error",
			mockErr: &models.ErrorResponse{
				Success:    false,
				StatusCode: 500,
				Errors: []errVals.ErrorObj{*errVals.NewErrorObj(
					errVals.ErrCreateUserCode,
					errVals.CustomError{Err: errors.New("cannot get cookie from redis")},
				)},
			},
			expectedErr: &models.ErrorResponse{
				Success:    false,
				StatusCode: 500,
				Errors: []errVals.ErrorObj{*errVals.NewErrorObj(
					errVals.ErrCreateUserCode,
					errVals.CustomError{Err: errors.New("cannot get cookie from redis")},
				)},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			srv := srvMock.NewMockAuthServiceInterface(ctrl)
			imp := NewImplementation(context.Background(), srv)

			srv.EXPECT().Session(gomock.Any(), gomock.Any()).Return(test.mockReturn, test.mockErr)
			resp, err := imp.Session(context.Background(), test.cookie)

			if test.expectedErr != nil {
				assert.Nil(t, resp)
				assert.Equal(t, test.expectedErr, err)
			} else {
				assert.Equal(t, test.expectedResponse, resp)
				assert.Nil(t, err)
			}
		})
	}
}
