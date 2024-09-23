package api

import (
	"context"
	"net/url"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

func (i *Implementation) GetCollection(ctx context.Context, query url.Values) (*models.CollectionsResponse, *models.ErrorResponse) {
	colls, err := i.service.GetCollection(ctx)
	if err != nil {
		return nil, err
	}

	return colls, nil
}

func (i *Implementation) Register(ctx context.Context, registerData *authModels.RegisterData) (*authModels.AuthResponse, *models.ErrorResponse) {
	resp, err := i.service.Register(ctx, registerData)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *Implementation) Login(ctx context.Context, loginData *authModels.LoginData) (*authModels.AuthResponse, *models.ErrorResponse) {
	resp, err := i.service.Login(ctx, loginData)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *Implementation) Session(ctx context.Context, cookie string) (*authModels.SessionResponse, *models.ErrorResponse) {
	resp, err := i.service.Session(ctx, cookie)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
