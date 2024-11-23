package client

import (
	"context"
	"time"

	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

type AuthClientInterface interface {
	CreateSession(ctx context.Context, usrID int) (*models.CookieData, error)
	DestroySession(ctx context.Context, cookie string) error
	Session(ctx context.Context, cookie string) (uint64, error)
}

type AuthClient struct {
	authMS auth.SessionRPCClient
}

func NewAuthClient(authMS auth.SessionRPCClient) AuthClientInterface {
	return &AuthClient{
		authMS: authMS,
	}
}

func (cl *AuthClient) CreateSession(ctx context.Context, usrID int) (*models.CookieData, error) {
	resp, err := cl.authMS.CreateSession(ctx, &auth.CreateSessionRequest{UserID: uint64(usrID)})

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

func (cl *AuthClient) DestroySession(ctx context.Context, cookie string) error {
	_, err := cl.authMS.DestroySession(ctx, &auth.DestroySessionRequest{Cookie: cookie})

	if err != nil {
		return err
	}

	return nil
}

func (cl AuthClient) Session(ctx context.Context, cookie string) (uint64, error) {
	resp, err := cl.authMS.Session(ctx, &auth.GetSessionRequest{Cookie: cookie})

	if err != nil {
		return 0, err
	}

	return resp.UserID, nil
}
