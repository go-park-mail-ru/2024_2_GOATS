package service

// import (
// 	"context"
// 	"errors"
// 	"os"
// 	"testing"

// 	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
// 	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
// 	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/password"
// 	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service/converter"
// 	mockRep "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service/mocks"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// func TestUserService_UpdatePassword(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mUsrRep := mockRep.NewMockUserRepositoryInterface(ctrl)
// 	userService := NewUserService(mUsrRep)
// 	hashedPasswd, err := password.HashAndSalt(context.Background(), "oldpassword")
// 	assert.NoError(t, err)

// 	tests := []struct {
// 		name          string
// 		passwordData  *models.PasswordData
// 		mockUser      *models.User
// 		mockUserErr   *errVals.RepoError
// 		mockUpdateErr *errVals.RepoError
// 		expectedError *errVals.ServiceError
// 		tryUpdate     bool
// 	}{
// 		{
// 			name: "Success",
// 			passwordData: &models.PasswordData{
// 				UserID:               1,
// 				OldPassword:          "oldpassword",
// 				Password:             "newpassword",
// 				PasswordConfirmation: "newpassword",
// 			},
// 			mockUser: &models.User{
// 				ID:       1,
// 				Password: hashedPasswd,
// 			},
// 			tryUpdate: true,
// 		},
// 		{
// 			name: "User not found",
// 			passwordData: &models.PasswordData{
// 				UserID:               2,
// 				OldPassword:          "somepassword",
// 				Password:             "newpassword",
// 				PasswordConfirmation: "newpassword",
// 			},
// 			mockUserErr: &errVals.RepoError{
// 				Code:  "user_not_found",
// 				Error: errVals.CustomError{Err: errors.New("user not found")},
// 			},
// 			expectedError: &errVals.ServiceError{
// 				Code:  "user_not_found",
// 				Error: errVals.CustomError{Err: errors.New("user not found")},
// 			},
// 		},
// 		{
// 			name: "Invalid old password",
// 			passwordData: &models.PasswordData{
// 				UserID:               1,
// 				OldPassword:          "wrongpassword",
// 				Password:             "newpassword",
// 				PasswordConfirmation: "newpassword",
// 			},
// 			mockUser: &models.User{
// 				ID:       1,
// 				Password: hashedPasswd,
// 			},
// 			expectedError: &errVals.ServiceError{
// 				Code:  errVals.ErrInvalidPasswordCode,
// 				Error: errVals.ErrInvalidOldPassword,
// 			},
// 		},
// 		{
// 			name: "Repository update error",
// 			passwordData: &models.PasswordData{
// 				UserID:               1,
// 				OldPassword:          "oldpassword",
// 				Password:             "newpassword",
// 				PasswordConfirmation: "newpassword",
// 			},
// 			mockUser: &models.User{
// 				ID:       1,
// 				Password: hashedPasswd,
// 			},
// 			mockUpdateErr: &errVals.RepoError{
// 				Code:  "update_failed",
// 				Error: errVals.CustomError{Err: errors.New("failed to update password")},
// 			},
// 			expectedError: &errVals.ServiceError{
// 				Code:  "update_failed",
// 				Error: errVals.CustomError{Err: errors.New("failed to update password")},
// 			},
// 			tryUpdate: true,
// 		},
// 	}

// 	ctx := context.Background()
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.mockUserErr != nil {
// 				mUsrRep.EXPECT().UserByID(ctx, tt.passwordData.UserID).Return(nil, tt.mockUserErr)
// 			} else {
// 				mUsrRep.EXPECT().UserByID(ctx, tt.passwordData.UserID).Return(tt.mockUser, nil)
// 			}

// 			if tt.mockUserErr == nil && tt.tryUpdate {
// 				mUsrRep.EXPECT().UpdatePassword(ctx, tt.passwordData.UserID, gomock.Any()).Return(tt.mockUpdateErr)
// 			}

// 			errData := userService.UpdatePassword(ctx, tt.passwordData)

// 			if tt.expectedError != nil {
// 				assert.Equal(t, tt.expectedError, errData)
// 			} else {
// 				assert.Nil(t, errData)
// 			}
// 		})
// 	}
// }

// func TestUserService_UpdateProfile(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mUsrRep := mockRep.NewMockUserRepositoryInterface(ctrl)
// 	userService := NewUserService(mUsrRep)

// 	tests := []struct {
// 		name          string
// 		usrData       *models.User
// 		mockAvatarErr *errVals.RepoError
// 		mockUpdateErr *errVals.RepoError
// 		expectedError *errVals.ServiceError
// 	}{
// 		// {
// 		// 	name: "Success with avatar",
// 		// 	usrData: &models.User{
// 		// 		ID:         1,
// 		// 		AvatarName: "avatar.png",
// 		// 		Email:      "test@mail.ru",
// 		// 		Username:   "hello world",
// 		// 	},
// 		// },
// 		{
// 			name: "Success without avatar",
// 			usrData: &models.User{
// 				ID:         1,
// 				AvatarName: "",
// 				Email:      "test@mail.ru",
// 				Username:   "hello world",
// 			},
// 		},
// 		// {
// 		// 	name: "Error updating profile",
// 		// 	usrData: &models.User{
// 		// 		ID:         1,
// 		// 		AvatarName: "avatar.png",
// 		// 		Email:      "test@mail.ru",
// 		// 		Username:   "hello world",
// 		// 	},
// 		// 	mockUpdateErr: &errVals.RepoError{
// 		// 		Code:  "update_failed",
// 		// 		Error: errVals.CustomError{Err: errors.New("failed to update profile")},
// 		// 	},
// 		// 	expectedError: &errVals.ServiceError{
// 		// 		Code:  "update_failed",
// 		// 		Error: errVals.CustomError{Err: errors.New("failed to update profile")},
// 		// 	},
// 		// },
// 	}

// 	ctx := context.Background()
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.mockAvatarErr != nil {
// 				mUsrRep.EXPECT().SaveUserAvatar(ctx, converter.ToRepoUserFromUser(tt.usrData)).Return("", tt.mockAvatarErr)
// 			} else {
// 				if tt.usrData.AvatarName != "" {
// 					tempFile, err := os.CreateTemp("", tt.usrData.AvatarName)
// 					require.NoError(t, err)
// 					mUsrRep.EXPECT().SaveUserAvatar(ctx, tt.usrData.AvatarName).Return("http://example.com/avatar.png", tempFile, nil)
// 				}
// 			}

// 			if tt.mockUpdateErr != nil {
// 				mUsrRep.EXPECT().UpdateProfileData(ctx, converter.ToRepoUserFromUser(tt.usrData)).Return(tt.mockUpdateErr)
// 			} else if tt.mockAvatarErr == nil {
// 				mUsrRep.EXPECT().UpdateProfileData(ctx, converter.ToRepoUserFromUser(tt.usrData)).Return(nil)
// 			}

// 			errData := userService.UpdateProfile(ctx, tt.usrData)

// 			if tt.expectedError != nil {
// 				assert.Equal(t, tt.expectedError, errData)
// 			} else {
// 				assert.Nil(t, errData)
// 			}
// 		})
// 	}
// }
