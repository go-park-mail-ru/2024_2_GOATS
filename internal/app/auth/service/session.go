package service

import (
	"context"

	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
)

func (s *AuthService) Session(ctx context.Context, cookie string) (*models.SessionRespData, *errVals.ServiceError) {
	aMsResp, msErr := s.authMS.Session(ctx, &auth.GetSessionRequest{Cookie: cookie})

	if msErr != nil || aMsResp.UserID != 0 {
		// return nil, errVals.ToServiceErrorFromRepo(msErr)
	}

	usr, err := s.userMS.FindByID(ctx, &user.ID{ID: aMsResp.UserID})
	if err != nil {
		// return nil, errVals.ToServiceErrorFromRepo(err)
	}

	return &models.SessionRespData{
		UserData: models.User{
			ID:         int(usr.UserID),
			Email:      usr.Email,
			Username:   usr.Username,
			AvatarURL:  usr.AvatarURL,
			AvatarName: usr.AvatarName,
		},
	}, nil
}
