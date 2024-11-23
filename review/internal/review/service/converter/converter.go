package converter

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/dto"
	repoDTO "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/review/repository/dto"
)

func ToRepoTokenFromSrv(token *dto.Token) *repoDTO.TokenData {
	if token == nil {
		return nil
	}

	return &repoDTO.TokenData{
		UserID:  token.UserID,
		Expiry:  token.Expiry,
		TokenID: token.TokenID,
	}
}
