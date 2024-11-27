package converter

import (
	repoDTO "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
)

func ConvertToRepoUser(u *dto.User) *repoDTO.RepoUser {
	if u == nil {
		return nil
	}

	return &repoDTO.RepoUser{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		AvatarURL: u.AvatarURL,
	}
}

func ConvertToRepoFavorite(f *dto.Favorite) *repoDTO.RepoFavorite {
	if f == nil {
		return nil
	}

	return &repoDTO.RepoFavorite{
		UserID:  f.UserID,
		MovieID: f.MovieID,
	}
}

func ConvertToRepoCreateData(cr *dto.CreateUserData) *repoDTO.RepoCreateData {
	if cr == nil {
		return nil
	}

	return &repoDTO.RepoCreateData{
		Email:                cr.Email,
		Username:             cr.Username,
		Password:             cr.Password,
		PasswordConfirmation: cr.PasswordConfirmation,
	}
}
