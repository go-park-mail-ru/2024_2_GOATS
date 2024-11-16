package converter

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/dto"
)

func ToDBUserFromUser(u *models.User) *dto.DBUser {
	if u == nil {
		return nil
	}

	return &dto.DBUser{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		AvatarURL: u.AvatarURL,
	}
}

func ToDBFavoriteFromFavorite(f *models.Favorite) *dto.DBFavorite {
	if f == nil {
		return nil
	}

	return &dto.DBFavorite{
		UserID:  f.UserID,
		MovieID: f.MovieID,
	}
}
