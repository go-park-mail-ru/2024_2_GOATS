package service

import (
	"context"
	"errors"
	"net/http"
	"testing"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/password"
	mockRep "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_UpdatePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mUsrRep := mockRep.NewMockUserRepositoryInterface(ctrl)
	userService := NewUserService(mUsrRep)
	hashedPasswd, err := password.HashAndSalt(context.Background(), "oldpassword")
	assert.NoError(t, err)

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
				UserID:               1,
				OldPassword:          "oldpassword",
				Password:             "newpassword",
				PasswordConfirmation: "newpassword",
			},
			mockUser: &models.User{
				ID:       1,
				Password: hashedPasswd,
			},
			expectedStatus: http.StatusOK,
			tryUpdate:      true,
		},
		{
			name: "User not found",
			passwordData: &models.PasswordData{
				UserID:               2,
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
				UserID:               1,
				OldPassword:          "wrongpassword",
				Password:             "newpassword",
				PasswordConfirmation: "newpassword",
			},
			mockUser: &models.User{
				ID:       1,
				Password: hashedPasswd,
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
				UserID:               1,
				OldPassword:          "oldpassword",
				Password:             "newpassword",
				PasswordConfirmation: "newpassword",
			},
			mockUser: &models.User{
				ID:       1,
				Password: hashedPasswd,
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

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockUserErr != nil {
				mUsrRep.EXPECT().UserByID(ctx, tt.passwordData.UserID).Return(nil, tt.mockUserErr, tt.expectedStatus)
			} else {
				mUsrRep.EXPECT().UserByID(ctx, tt.passwordData.UserID).Return(tt.mockUser, nil, http.StatusOK)
			}

			if tt.mockUserErr == nil && tt.tryUpdate {
				mUsrRep.EXPECT().UpdatePassword(ctx, tt.passwordData.UserID, gomock.Any()).Return(tt.mockUpdateErr, tt.expectedStatus)
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

	mUsrRep := mockRep.NewMockUserRepositoryInterface(ctrl)
	userService := NewUserService(mUsrRep)

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
				ID:         1,
				AvatarName: "avatar.png",
				Email:      "test@mail.ru",
				Username:   "hello world",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Success without avatar",
			usrData: &models.User{
				ID:         1,
				AvatarName: "",
				Email:      "test@mail.ru",
				Username:   "hello world",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Error saving avatar",
			usrData: &models.User{
				ID:         1,
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
				ID:         1,
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

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockAvatarErr != nil {
				mUsrRep.EXPECT().SaveUserAvatar(ctx, tt.usrData).Return("", tt.mockAvatarErr)
			} else {
				if tt.usrData.AvatarName != "" {
					mUsrRep.EXPECT().SaveUserAvatar(ctx, tt.usrData).Return("http://example.com/avatar.png", nil)
				}
			}

			if tt.mockUpdateErr != nil {
				mUsrRep.EXPECT().UpdateProfileData(ctx, tt.usrData).Return(tt.mockUpdateErr, http.StatusInternalServerError)
			} else if tt.mockAvatarErr == nil {
				mUsrRep.EXPECT().UpdateProfileData(ctx, tt.usrData).Return(nil, http.StatusOK)
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
