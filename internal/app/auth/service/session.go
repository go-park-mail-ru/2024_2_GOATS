package service

import (
	"context"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

// Session checks user session by cookie
func (s *AuthService) Session(ctx context.Context, cookie string) (*models.SessionRespData, *errVals.ServiceError) {
	usrID, err := s.authClient.Session(ctx, cookie)

	if err != nil || usrID == 0 {
		return nil, errVals.NewServiceError(errVals.ErrCheckSessionCode, fmt.Errorf("failed to get session data: %w", err))
	}

	usr, err := s.userClient.FindByID(ctx, usrID)
	if err != nil {
		return nil, errVals.NewServiceError(errVals.ErrGetUserCode, fmt.Errorf("failed to get session data: %w", err))
	}

	return &models.SessionRespData{
		UserData: *usr,
	}, nil
}
