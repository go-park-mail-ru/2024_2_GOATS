package client_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client"
	mockUser "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/client/mocks"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
)

func TestUserClient_FindByEmail(t *testing.T) {
	tests := []struct {
		name          string
		email         string
		mockSetup     func(mock *mockUser.MockUserRPCClient)
		expectedResp  *models.User
		expectedError error
	}{
		{
			name:  "Success",
			email: "test@example.com",
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					FindByEmail(gomock.Any(), &user.Email{Email: "test@example.com"}).
					Return(&user.UserData{
						UserID:     1,
						Email:      "test@example.com",
						Username:   "testuser",
						Password:   "hashed_password",
						AvatarURL:  "avatar_url",
						AvatarName: "avatar_name",
					}, nil)
			},
			expectedResp: &models.User{
				ID:         1,
				Email:      "test@example.com",
				Username:   "testuser",
				Password:   "hashed_password",
				AvatarURL:  "avatar_url",
				AvatarName: "avatar_name",
			},
			expectedError: nil,
		},
		{
			name:  "Error",
			email: "test@example.com",
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					FindByEmail(gomock.Any(), &user.Email{Email: "test@example.com"}).
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

			mockUserRPC := mockUser.NewMockUserRPCClient(ctrl)
			test.mockSetup(mockUserRPC)

			userClient := client.NewUserClient(mockUserRPC)
			resp, err := userClient.FindByEmail(context.Background(), test.email)

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

func TestUserClient_FindByID(t *testing.T) {
	tests := []struct {
		name          string
		id            uint64
		mockSetup     func(mock *mockUser.MockUserRPCClient)
		expectedResp  *models.User
		expectedError error
	}{
		{
			name: "Success",
			id:   1,
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					FindByID(gomock.Any(), &user.ID{ID: 1}).
					Return(&user.UserData{
						UserID:     1,
						Email:      "test@example.com",
						Username:   "testuser",
						Password:   "hashed_password",
						AvatarURL:  "avatar_url",
						AvatarName: "avatar_name",
					}, nil)
			},
			expectedResp: &models.User{
				ID:         1,
				Email:      "test@example.com",
				Username:   "testuser",
				Password:   "hashed_password",
				AvatarURL:  "avatar_url",
				AvatarName: "avatar_name",
			},
			expectedError: nil,
		},
		{
			name: "Error",
			id:   1,
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					FindByID(gomock.Any(), &user.ID{ID: 1}).
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

			mockUserRPC := mockUser.NewMockUserRPCClient(ctrl)
			test.mockSetup(mockUserRPC)

			userClient := client.NewUserClient(mockUserRPC)
			resp, err := userClient.FindByID(context.Background(), test.id)

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

func TestUserClient_GetFavorites(t *testing.T) {
	tests := []struct {
		name          string
		usrID         int
		mockSetup     func(mock *mockUser.MockUserRPCClient)
		expectedResp  []uint64
		expectedError error
	}{
		{
			name:  "Success",
			usrID: 1,
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					GetFavorites(gomock.Any(), &user.ID{ID: 1}).
					Return(&user.GetFavoritesResponse{
						MovieIDs: []uint64{1, 2, 3},
					}, nil)
			},
			expectedResp:  []uint64{1, 2, 3},
			expectedError: nil,
		},
		{
			name:  "Error",
			usrID: 1,
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					GetFavorites(gomock.Any(), &user.ID{ID: 1}).
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

			mockUserRPC := mockUser.NewMockUserRPCClient(ctrl)
			test.mockSetup(mockUserRPC)

			userClient := client.NewUserClient(mockUserRPC)
			resp, err := userClient.GetFavorites(context.Background(), test.usrID)

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

func TestUserClient_SetFavorite(t *testing.T) {
	tests := []struct {
		name          string
		favData       *models.Favorite
		mockSetup     func(mock *mockUser.MockUserRPCClient)
		expectedError error
	}{
		{
			name: "Success",
			favData: &models.Favorite{
				UserID:  1,
				MovieID: 2,
			},
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					SetFavorite(gomock.Any(), &user.HandleFavorite{
						UserID:  1,
						MovieID: 2,
					}).
					Return(&user.Nothing{Dummy: true}, nil)
			},
			expectedError: nil,
		},
		{
			name: "Error",
			favData: &models.Favorite{
				UserID:  1,
				MovieID: 2,
			},
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					SetFavorite(gomock.Any(), gomock.Any()).
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

			mockUserRPC := mockUser.NewMockUserRPCClient(ctrl)
			test.mockSetup(mockUserRPC)

			userClient := client.NewUserClient(mockUserRPC)
			err := userClient.SetFavorite(context.Background(), test.favData)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserClient_ResetFavorite(t *testing.T) {
	tests := []struct {
		name          string
		favData       *models.Favorite
		mockSetup     func(mock *mockUser.MockUserRPCClient)
		expectedError error
	}{
		{
			name: "Success",
			favData: &models.Favorite{
				UserID:  1,
				MovieID: 2,
			},
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					ResetFavorite(gomock.Any(), &user.HandleFavorite{
						UserID:  1,
						MovieID: 2,
					}).
					Return(&user.Nothing{Dummy: true}, nil)
			},
			expectedError: nil,
		},
		{
			name: "Error",
			favData: &models.Favorite{
				UserID:  1,
				MovieID: 2,
			},
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					ResetFavorite(gomock.Any(), gomock.Any()).
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

			mockUserRPC := mockUser.NewMockUserRPCClient(ctrl)
			test.mockSetup(mockUserRPC)

			userClient := client.NewUserClient(mockUserRPC)
			err := userClient.ResetFavorite(context.Background(), test.favData)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserClient_CheckFavorite(t *testing.T) {
	tests := []struct {
		name          string
		favData       *models.Favorite
		mockSetup     func(mock *mockUser.MockUserRPCClient)
		expectedResp  bool
		expectedError error
	}{
		{
			name: "Success",
			favData: &models.Favorite{
				UserID:  1,
				MovieID: 2,
			},
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					CheckFavorite(gomock.Any(), &user.HandleFavorite{
						UserID:  1,
						MovieID: 2,
					}).
					Return(&user.Nothing{Dummy: true}, nil)
			},
			expectedResp:  true,
			expectedError: nil,
		},
		{
			name: "Error",
			favData: &models.Favorite{
				UserID:  1,
				MovieID: 2,
			},
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					CheckFavorite(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("gRPC error"))
			},
			expectedResp:  false,
			expectedError: errors.New("gRPC error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRPC := mockUser.NewMockUserRPCClient(ctrl)
			test.mockSetup(mockUserRPC)

			userClient := client.NewUserClient(mockUserRPC)
			resp, err := userClient.CheckFavorite(context.Background(), test.favData)

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

func TestUserClient_UpdateProfile(t *testing.T) {
	tests := []struct {
		name          string
		usrData       *models.User
		mockSetup     func(mock *mockUser.MockUserRPCClient)
		expectedError error
	}{
		{
			name: "Success",
			usrData: &models.User{
				ID:         1,
				Email:      "test@example.com",
				Username:   "testuser",
				AvatarURL:  "avatar_url",
				AvatarName: "avatar_name",
			},
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					UpdateProfile(gomock.Any(), &user.UserData{
						UserID:     1,
						Email:      "test@example.com",
						Username:   "testuser",
						AvatarURL:  "avatar_url",
						AvatarName: "avatar_name",
					}).
					Return(&user.Nothing{Dummy: true}, nil)
			},
			expectedError: nil,
		},
		{
			name: "Error",
			usrData: &models.User{
				ID:         1,
				Email:      "test@example.com",
				Username:   "testuser",
				AvatarURL:  "avatar_url",
				AvatarName: "avatar_name",
			},
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					UpdateProfile(gomock.Any(), gomock.Any()).
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

			mockUserRPC := mockUser.NewMockUserRPCClient(ctrl)
			test.mockSetup(mockUserRPC)

			userClient := client.NewUserClient(mockUserRPC)
			err := userClient.UpdateProfile(context.Background(), test.usrData)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserClient_Create(t *testing.T) {
	tests := []struct {
		name          string
		regData       *models.RegisterData
		mockSetup     func(mock *mockUser.MockUserRPCClient)
		expectedResp  int
		expectedError error
	}{
		{
			name: "Success",
			regData: &models.RegisterData{
				Email:                "test@example.com",
				Username:             "testuser",
				Password:             "password123",
				PasswordConfirmation: "password123",
			},
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					Create(gomock.Any(), &user.CreateUserRequest{
						Email:                "test@example.com",
						Username:             "testuser",
						Password:             "password123",
						PasswordConfirmation: "password123",
					}).
					Return(&user.ID{ID: 1}, nil)
			},
			expectedResp:  1,
			expectedError: nil,
		},
		{
			name: "Error",
			regData: &models.RegisterData{
				Email:                "test@example.com",
				Username:             "testuser",
				Password:             "password123",
				PasswordConfirmation: "password123",
			},
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					Create(gomock.Any(), gomock.Any()).
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

			mockUserRPC := mockUser.NewMockUserRPCClient(ctrl)
			test.mockSetup(mockUserRPC)

			userClient := client.NewUserClient(mockUserRPC)
			resp, err := userClient.Create(context.Background(), test.regData)

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

func TestUserClient_UpdatePassword(t *testing.T) {
	tests := []struct {
		name          string
		passwordData  *models.PasswordData
		mockSetup     func(mock *mockUser.MockUserRPCClient)
		expectedError error
	}{
		{
			name: "Success",
			passwordData: &models.PasswordData{
				UserID:               1,
				OldPassword:          "oldpassword",
				Password:             "newpassword",
				PasswordConfirmation: "newpassword",
			},
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					UpdatePassword(gomock.Any(), &user.UpdatePasswordRequest{
						UserID:               1,
						OldPassword:          "oldpassword",
						Password:             "newpassword",
						PasswordConfirmation: "newpassword",
					}).
					Return(&user.Nothing{Dummy: true}, nil)
			},
			expectedError: nil,
		},
		{
			name: "Error",
			passwordData: &models.PasswordData{
				UserID:               1,
				OldPassword:          "oldpassword",
				Password:             "newpassword",
				PasswordConfirmation: "newpassword",
			},
			mockSetup: func(mock *mockUser.MockUserRPCClient) {
				mock.EXPECT().
					UpdatePassword(gomock.Any(), gomock.Any()).
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

			mockUserRPC := mockUser.NewMockUserRPCClient(ctrl)
			test.mockSetup(mockUserRPC)

			userClient := client.NewUserClient(mockUserRPC)
			err := userClient.UpdatePassword(context.Background(), test.passwordData)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
