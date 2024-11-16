package service

import (
	"context"
	"time"

	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
)

func (s *AuthService) Register(ctx context.Context, registerData *models.RegisterData) (*models.AuthRespData, *errVals.ServiceError) {
	createResp, err := s.userMS.Create(ctx, &user.CreateUserRequest{
		Email:                registerData.Email,
		Username:             registerData.Username,
		Password:             registerData.Password,
		PasswordConfirmation: registerData.PasswordConfirmation,
	})

	if err != nil {
		// return nil, errVals.ToServiceErrorFromRepo(err)
	}

	resp, msErr := s.authMS.CreateSession(ctx, &auth.CreateSessionRequest{UserID: createResp.ID})

	if msErr != nil {
		// return nil, errVals.ToServiceErrorFromRepo(msErr)
	}

	return &models.AuthRespData{
		NewCookie: &models.CookieData{
			Name: resp.Name,
			Token: &models.Token{
				UserID:  int(createResp.ID),
				TokenID: resp.Cookie,
				Expiry:  time.Unix(resp.MaxAge, 0),
			},
		},
	}, nil
}
