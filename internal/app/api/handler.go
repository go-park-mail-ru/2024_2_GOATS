package api

import (
	"context"
	"net/url"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

func (i *Implementation) GetCollection(ctx context.Context, query url.Values) (*models.CollectionsResponse, *models.ErrorResponse) {
	colls, errData := i.service.GetCollection(ctx)
	if errData != nil {
		return nil, errData
	}

	return colls, nil
}

func (i *Implementation) Register(ctx context.Context, registerData *authModels.RegisterData) (*authModels.AuthResponse, *models.ErrorResponse) {
	resp, errData := i.service.Register(ctx, registerData)
	if errData != nil {
		return nil, errData
	}

	return resp, nil
}

func (i *Implementation) Login(ctx context.Context, loginData *authModels.LoginData) (*authModels.AuthResponse, *models.ErrorResponse) {
	resp, errData := i.service.Login(ctx, loginData)
	if errData != nil {
		return nil, errData
	}

	return resp, nil
}

func (i *Implementation) Session(ctx context.Context, cookie string) (*authModels.SessionResponse, *models.ErrorResponse) {
	resp, errData := i.service.Session(ctx, cookie)
	if errData != nil {
		return nil, errData
	}

	return resp, nil
}
