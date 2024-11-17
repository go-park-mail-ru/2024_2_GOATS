package converter

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/dto"
)

func ToRepoUserFromUser(u *models.User) *dto.RepoUser {
	if u == nil {
		return nil
	}

	return &dto.RepoUser{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		AvatarURL: u.AvatarURL,
	}
}

func ToRepoFavoriteFromFavorite(f *models.Favorite) *dto.RepoFavorite {
	if f == nil {
		return nil
	}

	return &dto.RepoFavorite{
		UserID:  f.UserID,
		MovieID: f.MovieID,
	}
}
