package converter

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/repository/dto"
	srvDTO "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/dto"
)

func ToCookieFromRepo(title string, token *dto.TokenData) *srvDTO.Cookie {
	if token == nil {
		return nil
	}

	return &srvDTO.Cookie{
		Name:    title,
		UserID:  token.UserID,
		TokenID: token.TokenID,
		Expiry:  token.Expiry,
	}
}
