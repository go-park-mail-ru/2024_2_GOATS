package client

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
)

// UserClientInterface defines the methods for interacting with user-related data via an RPC client
//
//go:generate mockgen -source=user.go -destination=../user/service/mocks/mock.go
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
	CreateSubscription(ctx context.Context, data *models.SubscriptionData) (int, error)
	UpdateSubscriptionStatus(ctx context.Context, subID int) error
}

// UserClient implements UserClientInterface
type UserClient struct {
	UserMS user.UserRPCClient
}

// NewUserClient returns an instance of UserClientInterface
func NewUserClient(userMS user.UserRPCClient) UserClientInterface {
	return &UserClient{
		UserMS: userMS,
	}
}

// FindByEmail makes find_by_email request to user microservice
func (uc *UserClient) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	start := time.Now()
	method := "FindByEmail"

	resp, err := uc.UserMS.FindByEmail(ctx, &user.Email{Email: email})
	saveMetric(start, userClient, method, err)

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

// FindByID makes find_by_id request to user microservice
func (uc *UserClient) FindByID(ctx context.Context, id uint64) (*models.User, error) {
	start := time.Now()
	method := "FindByID"

	resp, err := uc.UserMS.FindByID(ctx, &user.ID{ID: id})
	saveMetric(start, userClient, method, err)

	if err != nil {
		return nil, fmt.Errorf("userClientError#findByID: %w", err)
	}

	return &models.User{
		ID:                         int(resp.UserID),
		Email:                      resp.Email,
		Username:                   resp.Username,
		Password:                   resp.Password,
		AvatarURL:                  resp.AvatarURL,
		AvatarName:                 resp.AvatarName,
		SubscriptionStatus:         resp.SubscriptionStatus,
		SubscriptionExpirationDate: resp.SubscriptionExpirationDate,
	}, nil
}

// Create makes create_user request to user microservice
func (uc *UserClient) Create(ctx context.Context, regData *models.RegisterData) (int, error) {
	start := time.Now()
	method := "CreateUser"

	resp, err := uc.UserMS.Create(ctx, &user.CreateUserRequest{
		Email:                regData.Email,
		Username:             regData.Username,
		Password:             regData.Password,
		PasswordConfirmation: regData.PasswordConfirmation,
	})

	saveMetric(start, userClient, method, err)

	if err != nil {
		return 0, fmt.Errorf("userClientError#create: %w", err)
	}

	return int(resp.ID), nil
}

// UpdateProfile makes find_by_id request to user microservice
func (uc *UserClient) UpdateProfile(ctx context.Context, usrData *models.User) error {
	var fileBytes []byte
	var err error

	start := time.Now()
	method := "UpdateProfile"

	if usrData.AvatarFile != nil {
		fileBytes, err = io.ReadAll(usrData.AvatarFile)
		if err != nil && usrData.AvatarFile != nil {
			saveMetric(start, userClient, method, err)
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

	_, err = uc.UserMS.UpdateProfile(ctx, req)
	saveMetric(start, userClient, method, err)

	if err != nil {
		return fmt.Errorf("userClientError#updateProfile: %w", err)
	}

	return nil
}

// UpdatePassword makes update_user_password request to user microservice
func (uc *UserClient) UpdatePassword(ctx context.Context, passwordData *models.PasswordData) error {
	start := time.Now()
	method := "UpdatePassword"

	_, err := uc.UserMS.UpdatePassword(ctx, &user.UpdatePasswordRequest{
		UserID:               uint64(passwordData.UserID),
		OldPassword:          passwordData.OldPassword,
		Password:             passwordData.Password,
		PasswordConfirmation: passwordData.PasswordConfirmation,
	})

	saveMetric(start, userClient, method, err)

	if err != nil {
		return fmt.Errorf("userClientError#updatePassword: %w", err)
	}

	return nil
}

// GetFavorites collects movie_ids from user's favorites from user microservice
func (uc *UserClient) GetFavorites(ctx context.Context, usrID int) ([]uint64, error) {
	start := time.Now()
	method := "GetFavorites"

	resp, err := uc.UserMS.GetFavorites(ctx, &user.ID{ID: uint64(usrID)})
	saveMetric(start, userClient, method, err)

	if err != nil {
		return nil, fmt.Errorf("userClientError#getFavorites: %w", err)
	}

	return resp.MovieIDs, nil
}

// SetFavorite create new user favorite in user microservice
func (uc *UserClient) SetFavorite(ctx context.Context, favData *models.Favorite) error {
	start := time.Now()
	method := "SetFavorite"

	err := uc.toggleFavorite(ctx, favData, "set")
	saveMetric(start, userClient, method, err)

	return err
}

// ResetFavorite destroy user favorite in user microservice
func (uc *UserClient) ResetFavorite(ctx context.Context, favData *models.Favorite) error {
	start := time.Now()
	method := "ResetFavorite"

	err := uc.toggleFavorite(ctx, favData, "reset")
	saveMetric(start, userClient, method, err)

	return err
}

// CheckFavorite checks favorite existense for given user in user microservice
func (uc *UserClient) CheckFavorite(ctx context.Context, favData *models.Favorite) (bool, error) {
	start := time.Now()
	method := "CheckFavorite"

	resp, err := uc.UserMS.CheckFavorite(ctx, &user.HandleFavorite{
		UserID:  uint64(favData.UserID),
		MovieID: uint64(favData.MovieID),
	})

	saveMetric(start, userClient, method, err)

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
		_, err = uc.UserMS.SetFavorite(ctx, data)
	} else {
		_, err = uc.UserMS.ResetFavorite(ctx, data)
	}

	if err != nil {
		return fmt.Errorf("userClientError#toggleFavorite: %w", err)
	}

	return nil
}

// CreateSubscription makes create_subscription request to user microservice
func (uc *UserClient) CreateSubscription(ctx context.Context, data *models.SubscriptionData) (int, error) {
	start := time.Now()
	method := "CreateSubscription"

	resp, err := uc.UserMS.Subscribe(ctx, &user.CreateSubscriptionRequest{
		UserID: uint64(data.UserID),
		Amount: uint64(data.Amount),
	})

	saveMetric(start, userClient, method, err)

	if err != nil {
		return 0, fmt.Errorf("userClientError#createSubscription: %w", err)
	}

	return int(resp.ID), nil
}

// UpdateSubscriptionStatus updates user subscription status in user microservice
func (uc *UserClient) UpdateSubscriptionStatus(ctx context.Context, subID int) error {
	start := time.Now()
	method := "UpdateSubscriptionStatus"

	_, err := uc.UserMS.UpdateSubscribtionStatus(ctx, &user.SubscriptionID{ID: uint64(subID)})

	saveMetric(start, userClient, method, err)

	if err != nil {
		return fmt.Errorf("userClientError#updateSubscriptionStatus: %w", err)
	}

	return nil
}
