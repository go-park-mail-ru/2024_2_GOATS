package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/repository/dto"
	srvDTO "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/dto"
)

var expirationTime = time.Now().Add(24 * time.Hour)
var closeRDBError = "cannot_close_redis_db"

func TestAuthRepository_SetCookie(t *testing.T) {
	tests := []struct {
		name          string
		token         *dto.TokenData
		mockSetup     func(mock redismock.ClientMock)
		expectedResp  *srvDTO.Cookie
		expectedError error
	}{
		{
			name: "Success",
			token: &dto.TokenData{
				TokenID: "valid_token",
				UserID:  1,
				Expiry:  expirationTime,
			},
			mockSetup: func(mock redismock.ClientMock) {
				mock.ExpectSet("valid_token", "1", 24*time.Hour).SetVal("OK")
			},
			expectedResp: &srvDTO.Cookie{
				Name:    "session_id",
				UserID:  1,
				TokenID: "valid_token",
				Expiry:  expirationTime,
			},
			expectedError: nil,
		},
		{
			name: "RedisError",
			token: &dto.TokenData{
				TokenID: "invalid_token",
				UserID:  1,
				Expiry:  expirationTime,
			},
			mockSetup: func(mock redismock.ClientMock) {
				mock.ExpectSet("invalid_token", "1", 24*time.Hour).SetErr(errors.New("redis error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("redis: cannot set cookie into redis - redis error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rdb, mock := redismock.NewClientMock()
			defer func() {
				if err := rdb.Close(); err != nil {
					t.Errorf("%s:%v", closeRDBError, err)
				}
			}()

			test.mockSetup(mock)
			repo := NewAuthRepository(rdb)
			resp, err := repo.SetCookie(testContext(t), test.token)

			assert.Equal(t, test.expectedResp, resp)
			if test.expectedError != nil {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestAuthRepository_DestroySession(t *testing.T) {
	tests := []struct {
		name          string
		cookie        string
		mockSetup     func(mock redismock.ClientMock)
		expectedError error
	}{
		{
			name:   "Success",
			cookie: "valid_cookie",
			mockSetup: func(mock redismock.ClientMock) {
				mock.ExpectDel("valid_cookie").SetVal(1)
			},
			expectedError: nil,
		},
		{
			name:   "RedisError",
			cookie: "invalid_cookie",
			mockSetup: func(mock redismock.ClientMock) {
				mock.ExpectDel("invalid_cookie").SetErr(errors.New("redis error"))
			},
			expectedError: fmt.Errorf("redis: failed to destroy session. Error - redis error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rdb, mock := redismock.NewClientMock()
			defer func() {
				if err := rdb.Close(); err != nil {
					t.Errorf("%s:%v", closeRDBError, err)
				}
			}()

			test.mockSetup(mock)
			repo := NewAuthRepository(rdb)
			err := repo.DestroySession(testContext(t), test.cookie)

			if test.expectedError != nil {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestAuthRepository_GetSessionData(t *testing.T) {
	tests := []struct {
		name           string
		cookie         string
		mockSetup      func(mock redismock.ClientMock)
		expectedUserID string
		expectedError  error
	}{
		{
			name:   "Success",
			cookie: "valid_cookie",
			mockSetup: func(mock redismock.ClientMock) {
				mock.ExpectGet("valid_cookie").SetVal("12345")
			},
			expectedUserID: "12345",
			expectedError:  nil,
		},
		{
			name:   "RedisError",
			cookie: "invalid_cookie",
			mockSetup: func(mock redismock.ClientMock) {
				mock.ExpectGet("invalid_cookie").SetErr(errors.New("redis error"))
			},
			expectedUserID: "",
			expectedError:  fmt.Errorf("redis: cannot get cookie from redis - redis error"),
		},
		{
			name:   "EmptyValue",
			cookie: "empty_cookie",
			mockSetup: func(mock redismock.ClientMock) {
				mock.ExpectGet("empty_cookie").SetVal("")
			},
			expectedUserID: "",
			expectedError:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rdb, mock := redismock.NewClientMock()
			defer func() {
				if err := rdb.Close(); err != nil {
					t.Errorf("%s:%v", closeRDBError, err)
				}
			}()

			test.mockSetup(mock)
			repo := NewAuthRepository(rdb)
			userID, err := repo.GetSessionData(testContext(t), test.cookie)

			assert.Equal(t, test.expectedUserID, userID)
			if test.expectedError != nil {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func testContext(t *testing.T) context.Context {
	require.NoError(t, os.Chdir("../../.."), "failed to change directory")

	cfg, err := config.New(true)
	require.NoError(t, err, "failed to read config from app_test")

	return config.WrapRedisContext(context.Background(), &cfg.Databases.Redis)
}
