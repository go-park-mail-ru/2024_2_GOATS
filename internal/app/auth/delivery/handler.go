package delivery

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

func (i *Implementation) Register(ctx context.Context, registerData *authModels.RegisterData) (*authModels.AuthResponse, *models.ErrorResponse) {
	resp, errData := i.authService.Register(ctx, registerData)
	if errData != nil {
		return nil, errData
	}

	return resp, nil
}

func (i *Implementation) Login(ctx context.Context, loginData *authModels.LoginData) (*authModels.AuthResponse, *models.ErrorResponse) {
	resp, errData := i.authService.Login(ctx, loginData)
	if errData != nil {
		return nil, errData
	}

	return resp, nil
}

func (i *Implementation) Session(ctx context.Context, cookie string) (*authModels.SessionResponse, *models.ErrorResponse) {
	resp, errData := i.authService.Session(ctx, cookie)
	if errData != nil {
		return nil, errData
	}

	return resp, nil
}

func (i *Implementation) Logout(ctx context.Context, cookie string) (*authModels.AuthResponse, *models.ErrorResponse) {
	resp, errData := i.authService.Logout(ctx, cookie)
	if errData != nil {
		return nil, errData
	}

	return resp, nil
}
