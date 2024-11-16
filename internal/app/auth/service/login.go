package service

import (
	"context"
	"time"

	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(ctx context.Context, loginData *models.LoginData) (*models.AuthRespData, *errVals.ServiceError) {
	logger := log.Ctx(ctx)
	usr, err := s.userMS.FindByEmail(ctx, &user.Email{Email: loginData.Email})

	if err != nil {
		// return nil, errVals.ToServiceErrorFromRepo(err)
	}

	cryptErr := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(loginData.Password))
	if cryptErr != nil {
		logger.Err(cryptErr).Msg("BCrypt: password missmatched.")

		return nil, errVals.NewServiceError(errVals.ErrInvalidPasswordCode, errVals.NewCustomError(errVals.ErrInvalidPasswordsMatch.Err.Error()))
	}

	_, desErr := s.authMS.DestroySession(ctx, &auth.DestroySessionRequest{Cookie: loginData.Cookie})
	if desErr != nil {
		// return nil, errVals.ToServiceErrorFromRepo(desErr)
	}

	ckResp, setErr := s.authMS.CreateSession(ctx, &auth.CreateSessionRequest{UserID: usr.UserID})
	if setErr != nil {
		// return nil, errVals.ToServiceErrorFromRepo(setErr)
	}

	return &models.AuthRespData{
		NewCookie: &models.CookieData{
			Name: ckResp.Name,
			Token: &models.Token{
				UserID:  int(usr.UserID),
				TokenID: ckResp.Cookie,
				Expiry:  time.Unix(ckResp.MaxAge, 0),
			},
		},
	}, nil
}
