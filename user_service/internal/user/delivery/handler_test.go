package delivery

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/errs"
	srvMock "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/delivery/mocks"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_Create(t *testing.T) {
	tests := []struct {
		name          string
		req           *user.CreateUserRequest
		mockSetup     func(mock *srvMock.MockUserServiceInterface)
		expectedResp  *user.ID
		expectedError error
	}{
		{
			name: "Success",
			req: &user.CreateUserRequest{
				Email:                "test@example.com",
				Username:             "testuser",
				Password:             "password123",
				PasswordConfirmation: "password123",
			},
			mockSetup: func(mock *srvMock.MockUserServiceInterface) {
				mock.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(uint64(1), nil)
			},
			expectedResp:  &user.ID{ID: 1},
			expectedError: nil,
		},
		{
			name: "ValidationError",
			req: &user.CreateUserRequest{
				Email: "invalid",
			},
			mockSetup:     nil,
			expectedResp:  nil,
			expectedError: errors.New("validation error"),
		},
		{
			name: "ServiceError",
			req: &user.CreateUserRequest{
				Email:                "test@example.com",
				Username:             "testuser",
				Password:             "password123",
				PasswordConfirmation: "password123",
			},
			mockSetup: func(mock *srvMock.MockUserServiceInterface) {
				mock.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(uint64(0), errors.New("service error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("service error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockUserServiceInterface(ctrl)
			if test.mockSetup != nil {
				test.mockSetup(mockService)
			}

			handler := NewUserHandler(testContext(t), mockService)

			resp, err := handler.Create(context.Background(), test.req)

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

func TestUserHandler_UpdateProfile(t *testing.T) {
	tests := []struct {
		name          string
		req           *user.UserData
		mockSetup     func(mock *srvMock.MockUserServiceInterface)
		expectedResp  *user.Nothing
		expectedError error
	}{
		{
			name: "Success",
			req: &user.UserData{
				UserID:    1,
				Email:     "test@example.com",
				Username:  "testuser",
				Password:  "",
				AvatarURL: "avatar_url",
			},
			mockSetup: func(mock *srvMock.MockUserServiceInterface) {
				mock.EXPECT().
					UpdateProfile(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedResp:  &user.Nothing{Dummy: true},
			expectedError: nil,
		},
		{
			name: "ServiceError",
			req: &user.UserData{
				UserID:    1,
				Email:     "test@example.com",
				Username:  "testuser",
				Password:  "",
				AvatarURL: "avatar_url",
			},
			mockSetup: func(mock *srvMock.MockUserServiceInterface) {
				mock.EXPECT().
					UpdateProfile(gomock.Any(), gomock.Any()).
					Return(errors.New("update profile error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("update profile error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockUserServiceInterface(ctrl)
			if test.mockSetup != nil {
				test.mockSetup(mockService)
			}

			handler := NewUserHandler(testContext(t), mockService)

			resp, err := handler.UpdateProfile(context.Background(), test.req)

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

func TestUserHandler_UpdatePassword(t *testing.T) {
	tests := []struct {
		name          string
		req           *user.UpdatePasswordRequest
		mockSetup     func(mock *srvMock.MockUserServiceInterface)
		expectedResp  *user.Nothing
		expectedError error
	}{
		{
			name: "Success",
			req: &user.UpdatePasswordRequest{
				UserID:               1,
				OldPassword:          "oldpassword",
				Password:             "newpassword",
				PasswordConfirmation: "newpassword",
			},
			mockSetup: func(mock *srvMock.MockUserServiceInterface) {
				mock.EXPECT().
					UpdatePassword(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedResp:  &user.Nothing{Dummy: true},
			expectedError: nil,
		},
		{
			name: "ValidationError",
			req: &user.UpdatePasswordRequest{
				UserID: 1,
			},
			mockSetup:     nil,
			expectedResp:  nil,
			expectedError: errors.New("validation error"),
		},
		{
			name: "ServiceError",
			req: &user.UpdatePasswordRequest{
				UserID:               1,
				OldPassword:          "oldpassss",
				Password:             "newpassssssss",
				PasswordConfirmation: "newpassssssss",
			},
			mockSetup: func(mock *srvMock.MockUserServiceInterface) {
				mock.EXPECT().
					UpdatePassword(gomock.Any(), gomock.Any()).
					Return(errors.New("update password error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("update password error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockUserServiceInterface(ctrl)
			if test.mockSetup != nil {
				test.mockSetup(mockService)
			}

			handler := NewUserHandler(testContext(t), mockService)
			resp, err := handler.UpdatePassword(context.Background(), test.req)

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

func TestUserHandler_GetFavorites(t *testing.T) {
	tests := []struct {
		name          string
		req           *user.ID
		mockSetup     func(mock *srvMock.MockUserServiceInterface)
		expectedResp  *user.GetFavoritesResponse
		expectedError error
	}{
		{
			name: "Success",
			req:  &user.ID{ID: 1},
			mockSetup: func(mock *srvMock.MockUserServiceInterface) {
				mock.EXPECT().
					GetFavorites(gomock.Any(), uint64(1)).
					Return([]uint64{1, 2, 3}, nil)
			},
			expectedResp:  &user.GetFavoritesResponse{MovieIDs: []uint64{1, 2, 3}},
			expectedError: nil,
		},
		{
			name:          "BadRequest",
			req:           &user.ID{ID: 0},
			mockSetup:     nil,
			expectedResp:  nil,
			expectedError: errs.ErrBadRequest,
		},
		{
			name: "ServiceError",
			req:  &user.ID{ID: 1},
			mockSetup: func(mock *srvMock.MockUserServiceInterface) {
				mock.EXPECT().
					GetFavorites(gomock.Any(), uint64(1)).
					Return(nil, errors.New("get favorites error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("get favorites error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockUserServiceInterface(ctrl)
			if test.mockSetup != nil {
				test.mockSetup(mockService)
			}

			handler := NewUserHandler(testContext(t), mockService)

			resp, err := handler.GetFavorites(context.Background(), test.req)

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

func TestUserHandler_SetFavorite(t *testing.T) {
	tests := []struct {
		name          string
		req           *user.HandleFavorite
		mockSetup     func(mock *srvMock.MockUserServiceInterface)
		expectedResp  *user.Nothing
		expectedError error
	}{
		{
			name: "Success",
			req: &user.HandleFavorite{
				UserID:  1,
				MovieID: 3,
			},
			mockSetup: func(mock *srvMock.MockUserServiceInterface) {
				mock.EXPECT().
					SetFavorite(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedResp:  &user.Nothing{Dummy: true},
			expectedError: nil,
		},
		{
			name: "ValidationError",
			req: &user.HandleFavorite{
				UserID: 0,
			},
			mockSetup:     nil,
			expectedResp:  nil,
			expectedError: errors.New("incorrect_favorite_params"),
		},
		{
			name: "ServiceError",
			req: &user.HandleFavorite{
				UserID:  1,
				MovieID: 3,
			},
			mockSetup: func(mock *srvMock.MockUserServiceInterface) {
				mock.EXPECT().
					SetFavorite(gomock.Any(), gomock.Any()).
					Return(errors.New("set favorite error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("set favorite error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockUserServiceInterface(ctrl)
			if test.mockSetup != nil {
				test.mockSetup(mockService)
			}

			handler := NewUserHandler(testContext(t), mockService)
			resp, err := handler.SetFavorite(context.Background(), test.req)

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

func TestUserHandler_CheckFavorite(t *testing.T) {
	tests := []struct {
		name          string
		req           *user.HandleFavorite
		mockSetup     func(mock *srvMock.MockUserServiceInterface)
		expectedResp  *user.Nothing
		expectedError error
	}{
		{
			name: "Success",
			req: &user.HandleFavorite{
				UserID:  1,
				MovieID: 3,
			},
			mockSetup: func(mock *srvMock.MockUserServiceInterface) {
				mock.EXPECT().
					CheckFavorite(gomock.Any(), gomock.Any()).
					Return(true, nil)
			},
			expectedResp:  &user.Nothing{Dummy: true},
			expectedError: nil,
		},
		{
			name: "ValidationError",
			req: &user.HandleFavorite{
				UserID: 0,
			},
			mockSetup:     nil,
			expectedResp:  nil,
			expectedError: errors.New("incorrect_favorite_params"),
		},
		{
			name: "ServiceError",
			req: &user.HandleFavorite{
				UserID:  1,
				MovieID: 3,
			},
			mockSetup: func(mock *srvMock.MockUserServiceInterface) {
				mock.EXPECT().
					CheckFavorite(gomock.Any(), gomock.Any()).
					Return(false, errors.New("check favorite error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("check favorite error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockUserServiceInterface(ctrl)
			if test.mockSetup != nil {
				test.mockSetup(mockService)
			}

			handler := NewUserHandler(testContext(t), mockService)

			resp, err := handler.CheckFavorite(context.Background(), test.req)

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

func TestUserHandler_FindByID(t *testing.T) {
	tests := []struct {
		name          string
		req           *user.ID
		mockSetup     func(mock *srvMock.MockUserServiceInterface)
		expectedResp  *user.UserData
		expectedError error
	}{
		{
			name: "Success",
			req:  &user.ID{ID: 1},
			mockSetup: func(mock *srvMock.MockUserServiceInterface) {
				mock.EXPECT().
					FindByID(gomock.Any(), uint64(1)).
					Return(&dto.User{
						ID:         1,
						Email:      "test@example.com",
						Username:   "testuser",
						Password:   "hashedpassword",
						AvatarURL:  "avatar_url",
						AvatarName: "avatar_name",
					}, nil)
			},
			expectedResp: &user.UserData{
				UserID:     1,
				Email:      "test@example.com",
				Username:   "testuser",
				Password:   "hashedpassword",
				AvatarURL:  "avatar_url",
				AvatarName: "avatar_name",
			},
			expectedError: nil,
		},
		{
			name:          "BadRequest",
			req:           &user.ID{ID: 0},
			mockSetup:     nil,
			expectedResp:  nil,
			expectedError: errs.ErrBadRequest,
		},
		{
			name: "ServiceError",
			req:  &user.ID{ID: 1},
			mockSetup: func(mock *srvMock.MockUserServiceInterface) {
				mock.EXPECT().
					FindByID(gomock.Any(), uint64(1)).
					Return(nil, errors.New("find user by ID error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("find user by ID error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockUserServiceInterface(ctrl)
			if test.mockSetup != nil {
				test.mockSetup(mockService)
			}

			handler := NewUserHandler(testContext(t), mockService)

			resp, err := handler.FindByID(context.Background(), test.req)

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

func TestUserHandler_FindByEmail(t *testing.T) {
	tests := []struct {
		name          string
		req           *user.Email
		mockSetup     func(mock *srvMock.MockUserServiceInterface)
		expectedResp  *user.UserData
		expectedError error
	}{
		{
			name: "Success",
			req:  &user.Email{Email: "test@example.com"},
			mockSetup: func(mock *srvMock.MockUserServiceInterface) {
				mock.EXPECT().
					FindByEmail(gomock.Any(), "test@example.com").
					Return(&dto.User{
						ID:         1,
						Email:      "test@example.com",
						Username:   "testuser",
						Password:   "hashedpassword",
						AvatarURL:  "avatar_url",
						AvatarName: "avatar_name",
					}, nil)
			},
			expectedResp: &user.UserData{
				UserID:     1,
				Email:      "test@example.com",
				Username:   "testuser",
				Password:   "hashedpassword",
				AvatarURL:  "avatar_url",
				AvatarName: "avatar_name",
			},
			expectedError: nil,
		},
		{
			name:          "BadRequest",
			req:           &user.Email{Email: ""},
			mockSetup:     nil,
			expectedResp:  nil,
			expectedError: errs.ErrBadRequest,
		},
		{
			name: "ServiceError",
			req:  &user.Email{Email: "test@example.com"},
			mockSetup: func(mock *srvMock.MockUserServiceInterface) {
				mock.EXPECT().
					FindByEmail(gomock.Any(), "test@example.com").
					Return(nil, errors.New("find user by email error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("find user by email error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockUserServiceInterface(ctrl)
			if test.mockSetup != nil {
				test.mockSetup(mockService)
			}

			handler := NewUserHandler(testContext(t), mockService)

			resp, err := handler.FindByEmail(context.Background(), test.req)

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

func testContext(t *testing.T) context.Context {
	require.NoError(t, os.Chdir("../../.."), "failed to change directory")

	cfg, err := config.New(true)
	require.NoError(t, err, "failed to read config from auth handler_test")

	return config.WrapContext(context.Background(), cfg)
}
