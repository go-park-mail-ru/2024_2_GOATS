package client

import (
	"context"
	"time"

	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

// AuthClientInterface defines client methods to transmit to Auth microservice
//
//go:generate mockgen -source=auth.go -destination=../auth/service/mocks/mock.go
type AuthClientInterface interface {
	CreateSession(ctx context.Context, usrID int) (*models.CookieData, error)
	DestroySession(ctx context.Context, cookie string) error
	Session(ctx context.Context, cookie string) (uint64, error)
}

// AuthClient struct implements AuthClientInterface
type AuthClient struct {
	authMS auth.SessionRPCClient
}

// NewAuthClient returns an instance of AuthClientInterface
func NewAuthClient(authMS auth.SessionRPCClient) AuthClientInterface {
	return &AuthClient{
		authMS: authMS,
	}
}

// CreateSession creates new session
func (cl *AuthClient) CreateSession(ctx context.Context, usrID int) (*models.CookieData, error) {
	start := time.Now()
	method := "CreateSession"

	resp, err := cl.authMS.CreateSession(ctx, &auth.CreateSessionRequest{UserID: uint64(usrID)})
	saveMetric(start, authClient, method, err)

	if err != nil {
		return nil, err
	}

	return &models.CookieData{
		Name: resp.Name,
		Token: &models.Token{
			UserID:  usrID,
			TokenID: resp.Cookie,
			Expiry:  time.Unix(resp.MaxAge, 0),
		},
	}, nil
}

// DestroySession destroys session
func (cl *AuthClient) DestroySession(ctx context.Context, cookie string) error {
	start := time.Now()
	method := "DestroySession"

	_, err := cl.authMS.DestroySession(ctx, &auth.DestroySessionRequest{Cookie: cookie})
	saveMetric(start, authClient, method, err)

	if err != nil {
		return err
	}

	return nil
}

// Session checks active session
func (cl *AuthClient) Session(ctx context.Context, cookie string) (uint64, error) {
	start := time.Now()
	method := "Session"

	resp, err := cl.authMS.Session(ctx, &auth.GetSessionRequest{Cookie: cookie})
	saveMetric(start, authClient, method, err)

	if err != nil {
		return 0, err
	}

	return resp.UserID, nil
}
