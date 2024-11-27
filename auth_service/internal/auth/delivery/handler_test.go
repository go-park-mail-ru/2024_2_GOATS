package delivery

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	srvMock "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/delivery/mocks"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/errs"
	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
)

var expirationTime = time.Now().Add(time.Second * 84400)

func TestAuthManager_CreateSession(t *testing.T) {
	tests := []struct {
		name          string
		req           *auth.CreateSessionRequest
		mockSetup     func(mock *srvMock.MockAuthServiceInterface)
		expectedResp  *auth.CreateSessionResponse
		expectedError error
	}{
		{
			name: "Success",
			req:  &auth.CreateSessionRequest{UserID: 1},
			mockSetup: func(mock *srvMock.MockAuthServiceInterface) {
				mock.EXPECT().CreateSession(gomock.Any(), &dto.SrvCreateCookie{UserID: 1}).Return(&dto.Cookie{
					TokenID: "test_cookie",
					Expiry:  expirationTime,
					Name:    "session_cookie",
				}, nil)
			},
			expectedResp: &auth.CreateSessionResponse{
				Cookie: "test_cookie",
				MaxAge: int64(expirationTime.Unix()),
				Name:   "session_cookie",
			},
			expectedError: nil,
		},
		{
			name:          "ValidationError",
			req:           &auth.CreateSessionRequest{UserID: 0},
			mockSetup:     nil,
			expectedResp:  nil,
			expectedError: fmt.Errorf("create_session: %s", errs.ErrInvalidUserID),
		},
		{
			name: "ServiceError",
			req:  &auth.CreateSessionRequest{UserID: 1},
			mockSetup: func(mock *srvMock.MockAuthServiceInterface) {
				mock.EXPECT().CreateSession(gomock.Any(), &dto.SrvCreateCookie{UserID: 1}).Return(nil, errors.New("service error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("service error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockAuthServiceInterface(ctrl)
			if test.mockSetup != nil {
				test.mockSetup(mockService)
			}

			am := NewAuthManager(context.Background(), mockService)

			resp, err := am.CreateSession(context.Background(), test.req)

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

func TestAuthManager_DestroySession(t *testing.T) {
	tests := []struct {
		name          string
		req           *auth.DestroySessionRequest
		mockSetup     func(mock *srvMock.MockAuthServiceInterface)
		expectedResp  *auth.Nothing
		expectedError error
	}{
		{
			name: "Success",
			req:  &auth.DestroySessionRequest{Cookie: "valid_cookie"},
			mockSetup: func(mock *srvMock.MockAuthServiceInterface) {
				mock.EXPECT().DestroySession(gomock.Any(), "valid_cookie").Return(true, nil)
			},
			expectedResp:  &auth.Nothing{Dummy: true},
			expectedError: nil,
		},
		{
			name:          "ValidationError",
			req:           &auth.DestroySessionRequest{Cookie: ""},
			mockSetup:     nil,
			expectedResp:  nil,
			expectedError: fmt.Errorf("destroy_session: %s", errs.ErrInvalidCookie),
		},
		{
			name: "ServiceError",
			req:  &auth.DestroySessionRequest{Cookie: "valid_cookie"},
			mockSetup: func(mock *srvMock.MockAuthServiceInterface) {
				mock.EXPECT().DestroySession(gomock.Any(), "valid_cookie").Return(false, errors.New("service error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("service error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockAuthServiceInterface(ctrl)
			if test.mockSetup != nil {
				test.mockSetup(mockService)
			}

			am := NewAuthManager(context.Background(), mockService)

			resp, err := am.DestroySession(context.Background(), test.req)

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

func TestAuthManager_Session(t *testing.T) {
	tests := []struct {
		name          string
		req           *auth.GetSessionRequest
		mockSetup     func(mock *srvMock.MockAuthServiceInterface)
		expectedResp  *auth.GetSessionResponse
		expectedError error
	}{
		{
			name: "Success",
			req:  &auth.GetSessionRequest{Cookie: "valid_cookie"},
			mockSetup: func(mock *srvMock.MockAuthServiceInterface) {
				mock.EXPECT().GetSessionData(gomock.Any(), "valid_cookie").Return(uint64(1), nil)
			},
			expectedResp:  &auth.GetSessionResponse{UserID: 1},
			expectedError: nil,
		},
		{
			name:          "ValidationError",
			req:           &auth.GetSessionRequest{Cookie: ""},
			mockSetup:     nil,
			expectedResp:  nil,
			expectedError: fmt.Errorf("check_session: %s", errs.ErrInvalidCookie),
		},
		{
			name: "ServiceError",
			req:  &auth.GetSessionRequest{Cookie: "valid_cookie"},
			mockSetup: func(mock *srvMock.MockAuthServiceInterface) {
				mock.EXPECT().GetSessionData(gomock.Any(), "valid_cookie").Return(uint64(0), errors.New("service error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("service error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockAuthServiceInterface(ctrl)
			if test.mockSetup != nil {
				test.mockSetup(mockService)
			}

			am := NewAuthManager(context.Background(), mockService)

			resp, err := am.Session(context.Background(), test.req)

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
