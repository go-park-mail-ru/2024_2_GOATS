package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

func (s *AuthService) SetActiveSessionTime(ctx context.Context, cookie string, seconds int) *errVals.ServiceError {
	usrID, idErr := s.authRepository.GetFromCookie(ctx, cookie)
	if idErr != nil {
		return errVals.ToServiceErrorFromRepo(idErr)
	}

	err := s.authRepository.SetActiveSessionTime(ctx, usrID, seconds)
	if err != nil {
		return errVals.NewServiceError("set_active_time_error", errVals.CustomError{Err: err})
	}

	return nil
}
