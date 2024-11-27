package converter

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	srvDTO "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
)

func ToUserFromRepoUser(u *dto.RepoUser) *srvDTO.User {
	if u == nil {
		return nil
	}

	return &srvDTO.User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		AvatarURL: u.AvatarURL,
	}
}
