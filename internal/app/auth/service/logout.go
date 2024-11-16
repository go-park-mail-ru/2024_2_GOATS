package service

import (
	"context"

	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

func (s *AuthService) Logout(ctx context.Context, cookie string) *errVals.ServiceError {
	_, err := s.authMS.DestroySession(ctx, &auth.DestroySessionRequest{Cookie: cookie})

	if err != nil {
		// return errVals.ToServiceErrorFromRepo(err)
	}

	return nil
}
