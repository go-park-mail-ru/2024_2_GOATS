package service

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/config"
	rdto "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/dto"
	repoMock "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/mocks"
)

func TestAuthService_GetSessionData(t *testing.T) {
	tests := []struct {
		name           string
		cookie         string
		mockSetup      func(mock *repoMock.MockAuthRepositoryInterface)
		expectedUserID uint64
		expectedError  error
	}{
		{
			name:   "Success",
			cookie: "valid_cookie",
			mockSetup: func(mock *repoMock.MockAuthRepositoryInterface) {
				mock.EXPECT().GetSessionData(gomock.Any(), "valid_cookie").Return("12345", nil)
			},
			expectedUserID: 12345,
			expectedError:  nil,
		},
		{
			name:   "RepositoryError",
			cookie: "invalid_cookie",
			mockSetup: func(mock *repoMock.MockAuthRepositoryInterface) {
				mock.EXPECT().GetSessionData(gomock.Any(), "invalid_cookie").Return("", errors.New("repository error"))
			},
			expectedUserID: 0,
			expectedError:  errors.New("failed to getSessionData: repository error"),
		},
		{
			name:   "InvalidUserIDFormat",
			cookie: "bad_cookie",
			mockSetup: func(mock *repoMock.MockAuthRepositoryInterface) {
				mock.EXPECT().GetSessionData(gomock.Any(), "bad_cookie").Return("invalid_id", nil)
			},
			expectedUserID: 0,
			expectedError:  errors.New("failed to getSessionData: failed to convert string into integer: strconv.ParseUint: parsing \"invalid_id\": invalid syntax"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := repoMock.NewMockAuthRepositoryInterface(ctrl)
			if test.mockSetup != nil {
				test.mockSetup(mockRepo)
			}

			authService := NewAuthService(mockRepo)

			userID, err := authService.GetSessionData(context.Background(), test.cookie)

			assert.Equal(t, test.expectedUserID, userID)
			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAuthService_DestroySession(t *testing.T) {
	tests := []struct {
		name           string
		cookie         string
		mockSetup      func(mock *repoMock.MockAuthRepositoryInterface)
		expectedResult bool
		expectedError  error
	}{
		{
			name:   "Success",
			cookie: "valid_cookie",
			mockSetup: func(mock *repoMock.MockAuthRepositoryInterface) {
				mock.EXPECT().DestroySession(gomock.Any(), "valid_cookie").Return(nil)
			},
			expectedResult: true,
			expectedError:  nil,
		},
		{
			name:   "RepositoryError",
			cookie: "invalid_cookie",
			mockSetup: func(mock *repoMock.MockAuthRepositoryInterface) {
				mock.EXPECT().DestroySession(gomock.Any(), "invalid_cookie").Return(errors.New("repository error"))
			},
			expectedResult: false,
			expectedError:  errors.New("failed to destroySession: repository error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := repoMock.NewMockAuthRepositoryInterface(ctrl)
			if test.mockSetup != nil {
				test.mockSetup(mockRepo)
			}

			authService := NewAuthService(mockRepo)

			result, err := authService.DestroySession(context.Background(), test.cookie)

			assert.Equal(t, test.expectedResult, result)
			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAuthService_CreateSession(t *testing.T) {
	tests := []struct {
		name          string
		inputData     *dto.SrvCreateCookie
		mockSetup     func(mockRepo *repoMock.MockAuthRepositoryInterface)
		expectedResp  *dto.Cookie
		expectedError error
	}{
		{
			name:      "Success",
			inputData: &dto.SrvCreateCookie{UserID: 1},
			mockSetup: func(mockRepo *repoMock.MockAuthRepositoryInterface) {
				mockRepo.EXPECT().SetCookie(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, token *rdto.TokenData) (*dto.Cookie, error) {
					return &dto.Cookie{
						Name:    "session_cookie",
						UserID:  token.UserID,
						TokenID: token.TokenID,
						Expiry:  token.Expiry,
					}, nil
				})
			},
			expectedResp: &dto.Cookie{
				Name:    "session_cookie",
				UserID:  1,
				TokenID: "mocked_token_id",
				Expiry:  time.Now().Add(24 * time.Hour),
			},
			expectedError: nil,
		},
		{
			name:      "RepositoryError",
			inputData: &dto.SrvCreateCookie{UserID: 1},
			mockSetup: func(mockRepo *repoMock.MockAuthRepositoryInterface) {
				mockRepo.EXPECT().SetCookie(gomock.Any(), gomock.Any()).Return(nil, errors.New("repository error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("failed to createSession: repository error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := repoMock.NewMockAuthRepositoryInterface(ctrl)
			if test.mockSetup != nil {
				test.mockSetup(mockRepo)
			}

			authService := NewAuthService(mockRepo)

			resp, err := authService.CreateSession(testContext(t), test.inputData)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, resp.TokenID)
			}
		})
	}
}

func testContext(t *testing.T) context.Context {
	require.NoError(t, os.Chdir("../../.."), "failed to change directory")

	cfg, err := config.New(true)
	require.NoError(t, err, "failed to read config from app_test")

	return config.WrapRedisContext(context.Background(), &cfg.Databases.Redis)
}
