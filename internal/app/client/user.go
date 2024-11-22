package client

import (
	"context"
	"fmt"
	"io"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
)

type UserClientInterface interface {
	Create(ctx context.Context, regData *models.RegisterData) (int, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByID(ctx context.Context, id uint64) (*models.User, error)
	UpdatePassword(ctx context.Context, passwordData *models.PasswordData) error
	GetFavorites(ctx context.Context, usrID int) ([]uint64, error)
	SetFavorite(ctx context.Context, favData *models.Favorite) error
	ResetFavorite(ctx context.Context, favData *models.Favorite) error
	CheckFavorite(ctx context.Context, favData *models.Favorite) (bool, error)
	UpdateProfile(ctx context.Context, usrData *models.User) error
}

type UserClient struct {
	userMS user.UserRPCClient
}

func NewUserClient(userMS user.UserRPCClient) UserClientInterface {
	return &UserClient{
		userMS: userMS,
	}
}

func (uc *UserClient) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	resp, err := uc.userMS.FindByEmail(ctx, &user.Email{Email: email})
	if err != nil {
		return nil, fmt.Errorf("userClientError#findByEmail: %w", err)
	}

	return &models.User{
		ID:         int(resp.UserID),
		Email:      resp.Email,
		Username:   resp.Username,
		Password:   resp.Password,
		AvatarURL:  resp.AvatarURL,
		AvatarName: resp.AvatarName,
	}, nil
}

func (uc *UserClient) FindByID(ctx context.Context, id uint64) (*models.User, error) {
	resp, err := uc.userMS.FindByID(ctx, &user.ID{ID: id})
	if err != nil {
		return nil, fmt.Errorf("userClientError#findByID: %w", err)
	}

	return &models.User{
		ID:         int(resp.UserID),
		Email:      resp.Email,
		Username:   resp.Username,
		Password:   resp.Password,
		AvatarURL:  resp.AvatarURL,
		AvatarName: resp.AvatarName,
	}, nil
}

func (uc *UserClient) Create(ctx context.Context, regData *models.RegisterData) (int, error) {
	resp, err := uc.userMS.Create(ctx, &user.CreateUserRequest{
		Email:                regData.Email,
		Username:             regData.Username,
		Password:             regData.Password,
		PasswordConfirmation: regData.PasswordConfirmation,
	})

	if err != nil {
		return 0, fmt.Errorf("userClientError#create: %w", err)
	}

	return int(resp.ID), nil
}

func (uc *UserClient) UpdateProfile(ctx context.Context, usrData *models.User) error {
	var fileBytes []byte
	var err error

	if usrData.AvatarFile != nil {
		fileBytes, err = io.ReadAll(usrData.AvatarFile)
		if err != nil && usrData.AvatarFile != nil {
			return fmt.Errorf("userClientError#updateProfile: %w", err)
		}
	} else {
		fileBytes = nil
	}

	req := &user.UserData{
		UserID:     uint64(usrData.ID),
		Email:      usrData.Email,
		Username:   usrData.Username,
		AvatarURL:  usrData.AvatarURL,
		AvatarName: usrData.AvatarName,
		AvatarFile: fileBytes,
	}
	_, err = uc.userMS.UpdateProfile(ctx, req)
	if err != nil {
		return fmt.Errorf("userClientError#updateProfile: %w", err)
	}

	return nil
}

func (uc *UserClient) UpdatePassword(ctx context.Context, passwordData *models.PasswordData) error {
	_, err := uc.userMS.UpdatePassword(ctx, &user.UpdatePasswordRequest{
		UserID:               uint64(passwordData.UserID),
		OldPassword:          passwordData.OldPassword,
		Password:             passwordData.Password,
		PasswordConfirmation: passwordData.PasswordConfirmation,
	})

	if err != nil {
		return fmt.Errorf("userClientError#updatePassword: %w", err)
	}

	return nil
}

func (uc *UserClient) GetFavorites(ctx context.Context, usrID int) ([]uint64, error) {
	resp, err := uc.userMS.GetFavorites(ctx, &user.ID{ID: uint64(usrID)})

	if err != nil {
		return nil, fmt.Errorf("userClientError#getFavorites: %w", err)
	}

	return resp.MovieIDs, nil
}

func (uc *UserClient) SetFavorite(ctx context.Context, favData *models.Favorite) error {
	return uc.toggleFavorite(ctx, favData, "set")
}

func (uc *UserClient) ResetFavorite(ctx context.Context, favData *models.Favorite) error {
	return uc.toggleFavorite(ctx, favData, "reset")
}

func (uc *UserClient) CheckFavorite(ctx context.Context, favData *models.Favorite) (bool, error) {
	resp, err := uc.userMS.CheckFavorite(ctx, &user.HandleFavorite{
		UserID:  uint64(favData.UserID),
		MovieID: uint64(favData.MovieID),
	})

	if err != nil {
		return false, fmt.Errorf("userClientError#checkFavorite: %w", err)
	}

	return resp.Dummy, nil
}

func (uc *UserClient) toggleFavorite(ctx context.Context, favData *models.Favorite, op string) error {
	data := &user.HandleFavorite{
		UserID:  uint64(favData.UserID),
		MovieID: uint64(favData.MovieID),
	}

	var err error
	if op == "set" {
		_, err = uc.userMS.SetFavorite(ctx, data)
	} else {
		_, err = uc.userMS.ResetFavorite(ctx, data)
	}

	if err != nil {
		return fmt.Errorf("userClientError#toggleFavorite: %w", err)
	}

	return nil
}
