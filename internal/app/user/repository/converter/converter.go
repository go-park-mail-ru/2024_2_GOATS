package converter

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/dto"
)

func ToUserFromRepoUser(u *dto.RepoUser) *models.User {
	if u == nil {
		return nil
	}

	return &models.User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		AvatarURL: u.AvatarURL,
	}
}

func ToMovieShortInfoFromRepo(m *dto.RepoMovieShortInfo) *models.MovieShortInfo {
	if m == nil {
		return nil
	}

	return &models.MovieShortInfo{
		ID:          m.ID,
		Title:       m.Title,
		CardURL:     m.CardURL,
		AlbumURL:    m.AlbumURL,
		Rating:      m.Rating,
		ReleaseDate: m.ReleaseDate,
		MovieType:   m.MovieType,
		Country:     m.Country,
	}
}
