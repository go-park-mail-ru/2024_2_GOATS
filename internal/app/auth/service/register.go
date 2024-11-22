package service

import (
	"context"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *AuthService) Register(ctx context.Context, registerData *models.RegisterData) (*models.AuthRespData, *errVals.ServiceError) {
	usrId, err := s.userClient.Create(ctx, registerData)
	if err != nil {
		return nil, errVals.NewServiceError(errVals.ErrCreateUserCode, fmt.Errorf("failed to register: %w", err))
	}

	ckData, err := s.authClient.CreateSession(ctx, usrId)
	if err != nil {
		return nil, errVals.NewServiceError(errVals.ErrCreateSessionCode, fmt.Errorf("failed to regisyer: %w", err))
	}

	return &models.AuthRespData{
		NewCookie: ckData,
	}, nil
}
