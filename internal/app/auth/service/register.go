package service

import (
	"context"
	"fmt"
	"strings"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

// Register user by given params
func (s *AuthService) Register(ctx context.Context, registerData *models.RegisterData) (*models.AuthRespData, *errVals.ServiceError) {
	usrID, err := s.userClient.Create(ctx, registerData)
	if err != nil {
		if strings.Contains(err.Error(), errVals.DuplicateErrCode) {
			return nil, errVals.NewServiceError(errVals.DuplicateErrCode, fmt.Errorf("failed to register: %w", err))
		}

		return nil, errVals.NewServiceError(errVals.ErrCreateUserCode, fmt.Errorf("failed to register: %w", err))
	}

	ckData, err := s.authClient.CreateSession(ctx, usrID)
	if err != nil {
		return nil, errVals.NewServiceError(errVals.ErrCreateSessionCode, fmt.Errorf("failed to regisyer: %w", err))
	}

	return &models.AuthRespData{
		NewCookie: ckData,
	}, nil
}
