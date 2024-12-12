package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service"
	clMock "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service/mocks"
)

func TestUpdatePassword(t *testing.T) {
	tests := []struct {
		name           string
		setupMocks     func(mockUserClient *clMock.MockUserClientInterface, ctx context.Context, passwordData *models.PasswordData)
		passwordData   *models.PasswordData
		expectedErr    error
		expectedErrMsg string
	}{
		{
			name: "Success",
			setupMocks: func(mockUserClient *clMock.MockUserClientInterface, ctx context.Context, passwordData *models.PasswordData) {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("old_password"), bcrypt.DefaultCost)
				mockUser := &models.User{
					ID:       1,
					Password: string(hashedPassword),
				}
				mockUserClient.EXPECT().FindByID(ctx, uint64(1)).Return(mockUser, nil)
				mockUserClient.EXPECT().UpdatePassword(ctx, passwordData).Return(nil)
			},
			passwordData: &models.PasswordData{
				UserID:               1,
				OldPassword:          "old_password",
				Password:             "new_password",
				PasswordConfirmation: "new_password",
			},
			expectedErr: nil,
		},
		{
			name: "InvalidOldPassword",
			setupMocks: func(mockUserClient *clMock.MockUserClientInterface, ctx context.Context, _ *models.PasswordData) {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("old_password"), bcrypt.DefaultCost)
				mockUser := &models.User{
					ID:       1,
					Password: string(hashedPassword),
				}
				mockUserClient.EXPECT().FindByID(ctx, uint64(1)).Return(mockUser, nil)
			},
			passwordData: &models.PasswordData{
				UserID:               1,
				OldPassword:          "wrong_password",
				Password:             "new_password",
				PasswordConfirmation: "new_password",
			},
			expectedErr:    errVals.ErrInvalidOldPassword,
			expectedErrMsg: errVals.ErrInvalidOldPassword.Error(),
		},
		{
			name: "FailureOnFindByID",
			setupMocks: func(mockUserClient *clMock.MockUserClientInterface, ctx context.Context, _ *models.PasswordData) {
				mockUserClient.EXPECT().FindByID(ctx, uint64(1)).Return(nil, errors.New("user not found"))
			},
			passwordData: &models.PasswordData{
				UserID:               1,
				OldPassword:          "old_password",
				Password:             "new_password",
				PasswordConfirmation: "new_password",
			},
			expectedErr:    errors.New("user not found"),
			expectedErrMsg: "user not found",
		},
		{
			name: "FailureOnUpdatePassword",
			setupMocks: func(mockUserClient *clMock.MockUserClientInterface, ctx context.Context, passwordData *models.PasswordData) {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("old_password"), bcrypt.DefaultCost)
				mockUser := &models.User{
					ID:       1,
					Password: string(hashedPassword),
				}
				mockUserClient.EXPECT().FindByID(ctx, uint64(1)).Return(mockUser, nil)
				mockUserClient.EXPECT().UpdatePassword(ctx, passwordData).Return(errors.New("update failed"))
			},
			passwordData: &models.PasswordData{
				UserID:               1,
				OldPassword:          "old_password",
				Password:             "new_password",
				PasswordConfirmation: "new_password",
			},
			expectedErr:    errors.New("update failed"),
			expectedErrMsg: "update failed",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserClient := clMock.NewMockUserClientInterface(ctrl)
			mockMvClient := clMock.NewMockMovieClientInterface(ctrl)
			userService := service.NewUserService(mockUserClient, mockMvClient)

			ctx := context.Background()
			tt.setupMocks(mockUserClient, ctx, tt.passwordData)

			err := userService.UpdatePassword(ctx, tt.passwordData)

			if tt.expectedErr == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error.Error(), tt.expectedErrMsg)
			}
		})
	}
}

func TestUpdateProfile(t *testing.T) {
	tests := []struct {
		name           string
		setupMocks     func(mockUserClient *clMock.MockUserClientInterface, ctx context.Context, usrData *models.User)
		usrData        *models.User
		expectedErr    error
		expectedErrMsg string
	}{
		{
			name: "Success",
			setupMocks: func(mockUserClient *clMock.MockUserClientInterface, ctx context.Context, usrData *models.User) {
				mockUserClient.EXPECT().UpdateProfile(ctx, usrData).Return(nil)
			},
			usrData: &models.User{
				ID:         1,
				Email:      "test@example.com",
				Username:   "testuser",
				AvatarURL:  "/static/avatars/avatar.png",
				AvatarName: "avatar.png",
			},
			expectedErr: nil,
		},
		{
			name: "UpdateProfileError",
			setupMocks: func(mockUserClient *clMock.MockUserClientInterface, ctx context.Context, usrData *models.User) {
				mockUserClient.EXPECT().UpdateProfile(ctx, usrData).Return(errors.New("update failed"))
			},
			usrData: &models.User{
				ID:         2,
				Email:      "user2@example.com",
				Username:   "user2",
				AvatarURL:  "/static/avatars/user2.png",
				AvatarName: "user2.png",
			},
			expectedErr:    errors.New("update failed"),
			expectedErrMsg: "update_profile_error",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserClient := clMock.NewMockUserClientInterface(ctrl)
			mockMvClient := clMock.NewMockMovieClientInterface(ctrl)
			userService := service.NewUserService(mockUserClient, mockMvClient)

			ctx := context.Background()
			tt.setupMocks(mockUserClient, ctx, tt.usrData)

			err := userService.UpdateProfile(ctx, tt.usrData)

			if tt.expectedErr == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tt.expectedErrMsg, err.Code)
			}
		})
	}
}

func TestResetFavorite(t *testing.T) {
	tests := []struct {
		name           string
		setupMocks     func(mockUserClient *clMock.MockUserClientInterface, ctx context.Context, favData *models.Favorite)
		favData        *models.Favorite
		expectedErr    error
		expectedErrMsg string
	}{
		{
			name: "Success",
			setupMocks: func(mockUserClient *clMock.MockUserClientInterface, ctx context.Context, favData *models.Favorite) {
				mockUserClient.EXPECT().ResetFavorite(ctx, favData).Return(nil)
			},
			favData: &models.Favorite{
				UserID:  1,
				MovieID: 100,
			},
			expectedErr: nil,
		},
		{
			name: "ResetFavoriteError",
			setupMocks: func(mockUserClient *clMock.MockUserClientInterface, ctx context.Context, favData *models.Favorite) {
				mockUserClient.EXPECT().ResetFavorite(ctx, favData).Return(errors.New("reset failed"))
			},
			favData: &models.Favorite{
				UserID:  2,
				MovieID: 200,
			},
			expectedErr:    errors.New("reset failed"),
			expectedErrMsg: "failed_to_reset_favorite",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserClient := clMock.NewMockUserClientInterface(ctrl)
			mockMvClient := clMock.NewMockMovieClientInterface(ctrl)
			userService := service.NewUserService(mockUserClient, mockMvClient)

			ctx := context.Background()
			tt.setupMocks(mockUserClient, ctx, tt.favData)

			err := userService.ResetFavorite(ctx, tt.favData)

			if tt.expectedErr == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tt.expectedErrMsg, err.Code)
			}
		})
	}
}

func TestAddFavorite(t *testing.T) {
	tests := []struct {
		name           string
		setupMocks     func(mockUserClient *clMock.MockUserClientInterface, ctx context.Context, favData *models.Favorite)
		favData        *models.Favorite
		expectedErr    error
		expectedErrMsg string
	}{
		{
			name: "Success",
			setupMocks: func(mockUserClient *clMock.MockUserClientInterface, ctx context.Context, favData *models.Favorite) {
				mockUserClient.EXPECT().SetFavorite(ctx, favData).Return(nil)
			},
			favData: &models.Favorite{
				UserID:  1,
				MovieID: 1,
			},
			expectedErr: nil,
		},
		{
			name: "SetFavoriteError",
			setupMocks: func(mockUserClient *clMock.MockUserClientInterface, ctx context.Context, favData *models.Favorite) {
				mockUserClient.EXPECT().SetFavorite(ctx, favData).Return(errors.New("set favorite failed"))
			},
			favData: &models.Favorite{
				UserID:  2,
				MovieID: 2,
			},
			expectedErr:    errors.New("set favorite failed"),
			expectedErrMsg: "failed_to_set_favorite",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserClient := clMock.NewMockUserClientInterface(ctrl)
			mockMvClient := clMock.NewMockMovieClientInterface(ctrl)
			userService := service.NewUserService(mockUserClient, mockMvClient)

			ctx := context.Background()
			tt.setupMocks(mockUserClient, ctx, tt.favData)

			err := userService.AddFavorite(ctx, tt.favData)

			if tt.expectedErr == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tt.expectedErrMsg, err.Code)
			}
		})
	}
}
