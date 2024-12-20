package service

import (
	"context"
	"errors"
	"testing"

	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
	mockRepo "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_CheckFavorite(t *testing.T) {
	tests := []struct {
		name          string
		favData       *dto.Favorite
		mockSetup     func(mock *mockRepo.MockUserRepoInterface)
		expectedResp  bool
		expectedError error
	}{
		{
			name: "Success - Is Favorite",
			favData: &dto.Favorite{
				UserID:  1,
				MovieID: 100,
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					CheckFavorite(gomock.Any(), converter.ConvertToRepoFavorite(&dto.Favorite{
						UserID:  1,
						MovieID: 100,
					})).
					Return(true, nil)
			},
			expectedResp:  true,
			expectedError: nil,
		},
		{
			name: "Success - Not Favorite",
			favData: &dto.Favorite{
				UserID:  2,
				MovieID: 200,
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					CheckFavorite(gomock.Any(), converter.ConvertToRepoFavorite(&dto.Favorite{
						UserID:  2,
						MovieID: 200,
					})).
					Return(false, nil)
			},
			expectedResp:  false,
			expectedError: nil,
		},
		{
			name: "Repo Error",
			favData: &dto.Favorite{
				UserID:  1,
				MovieID: 100,
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					CheckFavorite(gomock.Any(), converter.ConvertToRepoFavorite(&dto.Favorite{
						UserID:  1,
						MovieID: 100,
					})).
					Return(false, errors.New("database error"))
			},
			expectedResp:  false,
			expectedError: errors.New("userService - failed to checkFavorite: database error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mockRepo.NewMockUserRepoInterface(ctrl)
			test.mockSetup(mockRepo)

			userService := &UserService{userRepo: mockRepo}
			resp, err := userService.CheckFavorite(context.Background(), test.favData)

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

func TestUserService_Create(t *testing.T) {
	tests := []struct {
		name          string
		createData    *dto.CreateUserData
		mockSetup     func(mock *mockRepo.MockUserRepoInterface)
		expectedResp  uint64
		expectedError error
	}{
		{
			name: "Success",
			createData: &dto.CreateUserData{
				Email:    "test@example.com",
				Password: "password",
				Username: "testuser",
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					CreateUser(gomock.Any(), converter.ConvertToRepoCreateData(&dto.CreateUserData{
						Email:    "test@example.com",
						Password: "password",
						Username: "testuser",
					})).
					Return(&dto.User{
						ID:       1,
						Email:    "test@example.com",
						Username: "testuser",
					}, nil)
			},
			expectedResp:  1,
			expectedError: nil,
		},
		{
			name: "Repo Error",
			createData: &dto.CreateUserData{
				Email:    "test@example.com",
				Password: "password",
				Username: "testuser",
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					CreateUser(gomock.Any(), converter.ConvertToRepoCreateData(&dto.CreateUserData{
						Email:    "test@example.com",
						Password: "password",
						Username: "testuser",
					})).
					Return(nil, errors.New("database error"))
			},
			expectedResp:  0,
			expectedError: errors.New("userService - failed to create user: database error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mockRepo.NewMockUserRepoInterface(ctrl)
			test.mockSetup(mockRepo)

			userService := &UserService{userRepo: mockRepo}
			resp, err := userService.Create(context.Background(), test.createData)

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

func TestUserService_FindByEmail(t *testing.T) {
	tests := []struct {
		name          string
		email         string
		mockSetup     func(mock *mockRepo.MockUserRepoInterface)
		expectedResp  *dto.User
		expectedError error
	}{
		{
			name:  "Success",
			email: "test@example.com",
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					UserByEmail(gomock.Any(), "test@example.com").
					Return(&dto.User{
						ID:       1,
						Email:    "test@example.com",
						Username: "testuser",
					}, nil)
			},
			expectedResp: &dto.User{
				ID:       1,
				Email:    "test@example.com",
				Username: "testuser",
			},
			expectedError: nil,
		},
		{
			name:  "Repo Error",
			email: "unknown@example.com",
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					UserByEmail(gomock.Any(), "unknown@example.com").
					Return(nil, errors.New("user not found"))
			},
			expectedResp:  nil,
			expectedError: errors.New("userService - failed to get user by email: user not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mockRepo.NewMockUserRepoInterface(ctrl)
			test.mockSetup(mockRepo)

			userService := &UserService{userRepo: mockRepo}
			resp, err := userService.FindByEmail(context.Background(), test.email)

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

func TestUserService_FindByID(t *testing.T) {
	tests := []struct {
		name          string
		usrID         uint64
		mockSetup     func(mock *mockRepo.MockUserRepoInterface)
		expectedResp  *dto.User
		expectedError error
	}{
		{
			name:  "Success",
			usrID: 1,
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					UserByID(gomock.Any(), uint64(1)).
					Return(&dto.User{
						ID:       1,
						Email:    "test@example.com",
						Username: "testuser",
					}, nil)
			},
			expectedResp: &dto.User{
				ID:       1,
				Email:    "test@example.com",
				Username: "testuser",
			},
			expectedError: nil,
		},
		{
			name:  "Repo Error",
			usrID: 2,
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					UserByID(gomock.Any(), uint64(2)).
					Return(nil, errors.New("user not found"))
			},
			expectedResp:  nil,
			expectedError: errors.New("userService - failed to get user by id: user not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mockRepo.NewMockUserRepoInterface(ctrl)
			test.mockSetup(mockRepo)

			userService := &UserService{userRepo: mockRepo}
			resp, err := userService.FindByID(context.Background(), test.usrID)

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

func TestUserService_GetFavorites(t *testing.T) {
	tests := []struct {
		name          string
		usrID         uint64
		mockSetup     func(mock *mockRepo.MockUserRepoInterface)
		expectedResp  []uint64
		expectedError error
	}{
		{
			name:  "Success",
			usrID: 1,
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					GetFavorites(gomock.Any(), uint64(1)).
					Return([]uint64{101, 102, 103}, nil)
			},
			expectedResp:  []uint64{101, 102, 103},
			expectedError: nil,
		},
		{
			name:  "Empty Favorites",
			usrID: 2,
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					GetFavorites(gomock.Any(), uint64(2)).
					Return([]uint64{}, nil)
			},
			expectedResp:  []uint64{},
			expectedError: nil,
		},
		{
			name:  "Repo Error",
			usrID: 3,
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					GetFavorites(gomock.Any(), uint64(3)).
					Return(nil, errors.New("database error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("userService - failed to get user favorites: database error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mockRepo.NewMockUserRepoInterface(ctrl)
			test.mockSetup(mockRepo)

			userService := &UserService{userRepo: mockRepo}
			resp, err := userService.GetFavorites(context.Background(), test.usrID)

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

func TestUserService_ResetFavorite(t *testing.T) {
	tests := []struct {
		name          string
		favData       *dto.Favorite
		mockSetup     func(mock *mockRepo.MockUserRepoInterface)
		expectedError error
	}{
		{
			name: "Success",
			favData: &dto.Favorite{
				UserID:  1,
				MovieID: 101,
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					ResetFavorite(gomock.Any(), converter.ConvertToRepoFavorite(&dto.Favorite{
						UserID:  1,
						MovieID: 101,
					})).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Repo Error",
			favData: &dto.Favorite{
				UserID:  2,
				MovieID: 202,
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					ResetFavorite(gomock.Any(), converter.ConvertToRepoFavorite(&dto.Favorite{
						UserID:  2,
						MovieID: 202,
					})).
					Return(errors.New("database error"))
			},
			expectedError: errors.New("userService failed to reset favorite: database error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mockRepo.NewMockUserRepoInterface(ctrl)
			test.mockSetup(mockRepo)

			userService := &UserService{userRepo: mockRepo}
			err := userService.ResetFavorite(context.Background(), test.favData)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserService_SetFavorite(t *testing.T) {
	tests := []struct {
		name          string
		favData       *dto.Favorite
		mockSetup     func(mock *mockRepo.MockUserRepoInterface)
		expectedError error
	}{
		{
			name: "Success",
			favData: &dto.Favorite{
				UserID:  1,
				MovieID: 101,
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					SetFavorite(gomock.Any(), converter.ConvertToRepoFavorite(&dto.Favorite{
						UserID:  1,
						MovieID: 101,
					})).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Repo Error",
			favData: &dto.Favorite{
				UserID:  2,
				MovieID: 202,
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					SetFavorite(gomock.Any(), converter.ConvertToRepoFavorite(&dto.Favorite{
						UserID:  2,
						MovieID: 202,
					})).
					Return(errors.New("database error"))
			},
			expectedError: errors.New("userService - failed to setFavorite: database error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mockRepo.NewMockUserRepoInterface(ctrl)
			test.mockSetup(mockRepo)

			userService := &UserService{userRepo: mockRepo}
			err := userService.SetFavorite(context.Background(), test.favData)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserService_UpdatePassword(t *testing.T) {
	tests := []struct {
		name          string
		passwordData  *dto.PasswordData
		mockSetup     func(mock *mockRepo.MockUserRepoInterface)
		expectedError error
	}{
		{
			name: "Success",
			passwordData: &dto.PasswordData{
				UserID:      1,
				OldPassword: "oldPassword123",
				Password:    "newPassword123",
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				pass, err := bcrypt.GenerateFromPassword([]byte("oldPassword123"), bcrypt.DefaultCost)
				assert.NoError(t, err)

				mock.EXPECT().
					UserByID(gomock.Any(), uint64(1)).
					Return(&dto.User{ID: 1, Password: string(pass)}, nil)
				mock.EXPECT().
					UpdatePassword(gomock.Any(), uint64(1), "newPassword123").
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Repo Error - UserByID",
			passwordData: &dto.PasswordData{
				UserID:      2,
				OldPassword: "oldPassword123",
				Password:    "newPassword123",
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					UserByID(gomock.Any(), uint64(2)).
					Return(nil, errors.New("user not found"))
			},
			expectedError: fmt.Errorf("userService - failed to update password: user not found"),
		},
		{
			name: "Invalid Old Password",
			passwordData: &dto.PasswordData{
				UserID:      3,
				OldPassword: "wrongPassword",
				Password:    "newPassword123",
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				pass, err := bcrypt.GenerateFromPassword([]byte("oldPassword123"), bcrypt.DefaultCost)
				assert.NoError(t, err)

				mock.EXPECT().
					UserByID(gomock.Any(), uint64(3)).
					Return(&dto.User{ID: 3, Password: string(pass)}, nil)
			},
			expectedError: fmt.Errorf("userService failed to update password: %s: %w", errVals.ErrInvalidPasswordCode, errVals.ErrInvalidOldPassword),
		},
		{
			name: "Repo Error - UpdatePassword",
			passwordData: &dto.PasswordData{
				UserID:      4,
				OldPassword: "oldPassword123",
				Password:    "newPassword123",
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				pass, err := bcrypt.GenerateFromPassword([]byte("oldPassword123"), bcrypt.DefaultCost)
				assert.NoError(t, err)

				mock.EXPECT().
					UserByID(gomock.Any(), uint64(4)).
					Return(&dto.User{ID: 4, Password: string(pass)}, nil)
				mock.EXPECT().
					UpdatePassword(gomock.Any(), uint64(4), "newPassword123").
					Return(errors.New("database error"))
			},
			expectedError: fmt.Errorf("userService failed to update password: database error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mockRepo.NewMockUserRepoInterface(ctrl)
			test.mockSetup(mockRepo)

			userService := &UserService{userRepo: mockRepo}
			err := userService.UpdatePassword(context.Background(), test.passwordData)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserService_UpdateProfile(t *testing.T) {
	tests := []struct {
		name          string
		usrData       *dto.User
		mockSetup     func(mock *mockRepo.MockUserRepoInterface)
		expectedError error
	}{
		{
			name: "Success with Avatar Update",
			usrData: &dto.User{
				ID:         1,
				Username:   "testuser",
				Email:      "test@example.com",
				AvatarName: "avatar.png",
				AvatarFile: []byte("mock avatar file data"),
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					SaveUserAvatar(gomock.Any(), "avatar.png").
					Return("http://example.com/avatar.png", nil)
				mock.EXPECT().
					UpdateProfileData(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Success without Avatar Update",
			usrData: &dto.User{
				ID:       1,
				Username: "testuser",
				Email:    "test@example.com",
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					UpdateProfileData(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Error during Avatar Save",
			usrData: &dto.User{
				ID:         1,
				Username:   "testuser",
				Email:      "test@example.com",
				AvatarName: "avatar.png",
				AvatarFile: []byte("mock avatar file data"),
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					SaveUserAvatar(gomock.Any(), "avatar.png").
					Return("", errors.New("avatar save error"))
			},
			expectedError: fmt.Errorf("userService - failed to updateProfile: avatar save error"),
		},
		{
			name: "Error during Profile Update",
			usrData: &dto.User{
				ID:         1,
				Username:   "testuser",
				Email:      "test@example.com",
				AvatarName: "avatar.png",
				AvatarFile: []byte("mock avatar file data"),
			},
			mockSetup: func(mock *mockRepo.MockUserRepoInterface) {
				mock.EXPECT().
					SaveUserAvatar(gomock.Any(), "avatar.png").
					Return("http://example.com/avatar.png", nil)
				mock.EXPECT().
					UpdateProfileData(gomock.Any(), gomock.Any()).
					Return(errors.New("profile update error"))
			},
			expectedError: fmt.Errorf("userService failed to updateProfile: profile update error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mockRepo.NewMockUserRepoInterface(ctrl)
			test.mockSetup(mockRepo)

			userService := &UserService{userRepo: mockRepo}
			err := userService.UpdateProfile(context.Background(), test.usrData)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
