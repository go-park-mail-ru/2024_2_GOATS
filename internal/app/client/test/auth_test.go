package client_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	mockAuth "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client/mocks"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

var expirationTime = time.Now().Add(24 * time.Hour).Unix()

func TestAuthClient_CreateSession(t *testing.T) {
	tests := []struct {
		name          string
		usrID         int
		mockSetup     func(mock *mockAuth.MockSessionRPCClient)
		expectedResp  *models.CookieData
		expectedError error
	}{
		{
			name:  "Success",
			usrID: 1,
			mockSetup: func(mock *mockAuth.MockSessionRPCClient) {
				mock.EXPECT().
					CreateSession(gomock.Any(), &auth.CreateSessionRequest{UserID: 1}).
					Return(&auth.CreateSessionResponse{
						Cookie: "cookie-token",
						MaxAge: expirationTime,
						Name:   "session",
					}, nil)
			},
			expectedResp: &models.CookieData{
				Name: "session",
				Token: &models.Token{
					UserID:  1,
					TokenID: "cookie-token",
					Expiry:  time.Unix(expirationTime, 0),
				},
			},
			expectedError: nil,
		},
		{
			name:  "Error",
			usrID: 1,
			mockSetup: func(mock *mockAuth.MockSessionRPCClient) {
				mock.EXPECT().
					CreateSession(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("gRPC error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("gRPC error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockAuthRPC := mockAuth.NewMockSessionRPCClient(ctrl)
			test.mockSetup(mockAuthRPC)

			authClient := client.NewAuthClient(mockAuthRPC)
			resp, err := authClient.CreateSession(context.Background(), test.usrID)

			assert.Equal(t, test.expectedResp, resp)
			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAuthClient_DestroySession(t *testing.T) {
	tests := []struct {
		name          string
		cookie        string
		mockSetup     func(mock *mockAuth.MockSessionRPCClient)
		expectedError error
	}{
		{
			name:   "Success",
			cookie: "cookie-token",
			mockSetup: func(mock *mockAuth.MockSessionRPCClient) {
				mock.EXPECT().
					DestroySession(gomock.Any(), &auth.DestroySessionRequest{Cookie: "cookie-token"}).
					Return(&auth.Nothing{Dummy: true}, nil)
			},
			expectedError: nil,
		},
		{
			name:   "Error",
			cookie: "cookie-token",
			mockSetup: func(mock *mockAuth.MockSessionRPCClient) {
				mock.EXPECT().
					DestroySession(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("gRPC error"))
			},
			expectedError: errors.New("gRPC error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockAuthRPC := mockAuth.NewMockSessionRPCClient(ctrl)
			test.mockSetup(mockAuthRPC)

			authClient := client.NewAuthClient(mockAuthRPC)
			err := authClient.DestroySession(context.Background(), test.cookie)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAuthClient_Session(t *testing.T) {
	tests := []struct {
		name          string
		cookie        string
		mockSetup     func(mock *mockAuth.MockSessionRPCClient)
		expectedResp  uint64
		expectedError error
	}{
		{
			name:   "Success",
			cookie: "cookie-token",
			mockSetup: func(mock *mockAuth.MockSessionRPCClient) {
				mock.EXPECT().
					Session(gomock.Any(), &auth.GetSessionRequest{Cookie: "cookie-token"}).
					Return(&auth.GetSessionResponse{
						UserID: 1,
					}, nil)
			},
			expectedResp:  1,
			expectedError: nil,
		},
		{
			name:   "Error",
			cookie: "cookie-token",
			mockSetup: func(mock *mockAuth.MockSessionRPCClient) {
				mock.EXPECT().
					Session(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("gRPC error"))
			},
			expectedResp:  0,
			expectedError: errors.New("gRPC error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockAuthRPC := mockAuth.NewMockSessionRPCClient(ctrl)
			test.mockSetup(mockAuthRPC)

			authClient := client.NewAuthClient(mockAuthRPC)
			resp, err := authClient.Session(context.Background(), test.cookie)

			assert.Equal(t, test.expectedResp, resp)
			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
