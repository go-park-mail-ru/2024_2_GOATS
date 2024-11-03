package service

import (
	"context"
	"errors"
	"net/http"
	"testing"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	mockRep "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_UpdatePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockRep.NewMockUserRepositoryInterface(ctrl)
	userService := NewUserService(mockUserRepo)

	tests := []struct {
		name           string
		passwordData   *models.PasswordData
		mockUser       *models.User
		mockUserErr    *errVals.ErrorObj
		mockUpdateErr  *errVals.ErrorObj
		expectedStatus int
		expectedError  *models.ErrorRespData
		tryUpdate      bool
	}{
		{
			name: "Success",
			passwordData: &models.PasswordData{
				UserId:               1,
				OldPassword:          "oldpassword",
				Password:             "newpassword",
				PasswordConfirmation: "newpassword",
			},
			mockUser: &models.User{
				Id:       1,
				Password: hashPassword("oldpassword"),
			},
			expectedStatus: http.StatusOK,
			tryUpdate:      true,
		},
		{
			name: "User not found",
			passwordData: &models.PasswordData{
				UserId:               2,
				OldPassword:          "somepassword",
				Password:             "newpassword",
				PasswordConfirmation: "newpassword",
			},
			mockUserErr: &errVals.ErrorObj{
				Code:  "user_not_found",
				Error: errVals.CustomError{Err: errors.New("user not found")},
			},
			expectedStatus: http.StatusNotFound,
			expectedError: &models.ErrorRespData{
				StatusCode: http.StatusNotFound,
				Errors:     []errVals.ErrorObj{{Code: "user_not_found", Error: errVals.CustomError{Err: errors.New("user not found")}}},
			},
		},
		{
			name: "Invalid old password",
			passwordData: &models.PasswordData{
				UserId:               1,
				OldPassword:          "wrongpassword",
				Password:             "newpassword",
				PasswordConfirmation: "newpassword",
			},
			mockUser: &models.User{
				Id:       1,
				Password: hashPassword("oldpassword"),
			},
			expectedStatus: http.StatusConflict,
			expectedError: &models.ErrorRespData{
				StatusCode: http.StatusConflict,
				Errors: []errVals.ErrorObj{
					*errVals.NewErrorObj(errVals.ErrInvalidPasswordCode, errVals.ErrInvalidOldPasswordText),
				},
			},
		},
		{
			name: "Repository update error",
			passwordData: &models.PasswordData{
				UserId:               1,
				OldPassword:          "oldpassword",
				Password:             "newpassword",
				PasswordConfirmation: "newpassword",
			},
			mockUser: &models.User{
				Id:       1,
				Password: hashPassword("oldpassword"),
			},
			mockUpdateErr: &errVals.ErrorObj{
				Code:  "update_failed",
				Error: errVals.CustomError{Err: errors.New("failed to update password")},
			},
			expectedStatus: http.StatusInternalServerError,
			expectedError: &models.ErrorRespData{
				StatusCode: http.StatusInternalServerError,
				Errors: []errVals.ErrorObj{
					{Code: "update_failed", Error: errVals.CustomError{Err: errors.New("failed to update password")}},
				},
			},
			tryUpdate: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			if tt.mockUserErr != nil {
				mockUserRepo.EXPECT().UserById(ctx, tt.passwordData.UserId).Return(nil, tt.mockUserErr, tt.expectedStatus)
			} else {
				mockUserRepo.EXPECT().UserById(ctx, tt.passwordData.UserId).Return(tt.mockUser, nil, http.StatusOK)
			}

			if tt.mockUserErr == nil && tt.tryUpdate {
				mockUserRepo.EXPECT().UpdatePassword(ctx, tt.passwordData.UserId, gomock.Any()).Return(tt.mockUpdateErr, tt.expectedStatus)
			}

			respData, errData := userService.UpdatePassword(ctx, tt.passwordData)

			if tt.expectedError != nil {
				assert.Nil(t, respData)
				assert.Equal(t, tt.expectedError, errData)
			} else {
				assert.Nil(t, errData)
				assert.Equal(t, tt.expectedStatus, respData.StatusCode)
			}
		})
	}
}

func TestUserService_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockRep.NewMockUserRepositoryInterface(ctrl)
	userService := NewUserService(mockUserRepo)

	tests := []struct {
		name           string
		usrData        *models.User
		mockAvatarErr  *errVals.ErrorObj
		mockUpdateErr  *errVals.ErrorObj
		expectedStatus int
		expectedError  *models.ErrorRespData
	}{
		{
			name: "Success with avatar",
			usrData: &models.User{
				Id:         1,
				AvatarName: "avatar.png",
				Email:      "test@mail.ru",
				Username:   "hello world",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Success without avatar",
			usrData: &models.User{
				Id:         1,
				AvatarName: "",
				Email:      "test@mail.ru",
				Username:   "hello world",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Error saving avatar",
			usrData: &models.User{
				Id:         1,
				AvatarName: "avatar.png",
				Email:      "test@mail.ru",
				Username:   "hello world",
			},
			mockAvatarErr: &errVals.ErrorObj{
				Code:  "avatar_save_failed",
				Error: errVals.CustomError{Err: errors.New("failed to save avatar")},
			},
			expectedStatus: http.StatusInternalServerError,
			expectedError: &models.ErrorRespData{
				StatusCode: http.StatusInternalServerError,
				Errors: []errVals.ErrorObj{
					{Code: "avatar_save_failed", Error: errVals.CustomError{Err: errors.New("failed to save avatar")}},
				},
			},
		},
		{
			name: "Error updating profile",
			usrData: &models.User{
				Id:         1,
				AvatarName: "avatar.png",
				Email:      "test@mail.ru",
				Username:   "hello world",
			},
			mockUpdateErr: &errVals.ErrorObj{
				Code:  "update_failed",
				Error: errVals.CustomError{Err: errors.New("failed to update profile")},
			},
			expectedStatus: http.StatusInternalServerError,
			expectedError: &models.ErrorRespData{
				StatusCode: http.StatusInternalServerError,
				Errors: []errVals.ErrorObj{
					{Code: "update_failed", Error: errVals.CustomError{Err: errors.New("failed to update profile")}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			if tt.mockAvatarErr != nil {
				mockUserRepo.EXPECT().SaveAvatar(ctx, tt.usrData).Return("", tt.mockAvatarErr)
			} else {
				if tt.usrData.AvatarName != "" {
					mockUserRepo.EXPECT().SaveAvatar(ctx, tt.usrData).Return("http://example.com/avatar.png", nil)
				}
			}

			if tt.mockUpdateErr != nil {
				mockUserRepo.EXPECT().UpdateProfileData(ctx, tt.usrData).Return(tt.mockUpdateErr, http.StatusInternalServerError)
			} else if tt.mockAvatarErr == nil {
				mockUserRepo.EXPECT().UpdateProfileData(ctx, tt.usrData).Return(nil, http.StatusOK)
			}

			respData, errData := userService.UpdateProfile(ctx, tt.usrData)

			if tt.expectedError != nil {
				assert.Nil(t, respData)
				assert.Equal(t, tt.expectedError, errData)
			} else {
				assert.Nil(t, errData)
				assert.Equal(t, tt.expectedStatus, respData.StatusCode)
			}
		})
	}
}

// Вспомогательная функция для хеширования пароля
func hashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
