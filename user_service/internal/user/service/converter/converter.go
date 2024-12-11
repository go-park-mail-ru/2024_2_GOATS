package converter

import (
	repoDTO "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
)

// ConvertToRepoUser converts service dto User to repo DTO
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

// ConvertToRepoFavorite converts service dto Favorite to repo DTO
func ConvertToRepoFavorite(f *dto.Favorite) *repoDTO.RepoFavorite {
	if f == nil {
		return nil
	}

	return &repoDTO.RepoFavorite{
		UserID:  f.UserID,
		MovieID: f.MovieID,
	}
}

// ConvertToRepoCreateData converts service dto CreateUserData to repo DTO
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

// ConvertToRepoCreateSubData converts service dto SubscriptionData to repo DTO
func ConvertToRepoCreateSubData(cr *dto.CreateSubscriptionData) *repoDTO.RepoCreateSubscriptionData {
	if cr == nil {
		return nil
	}

	return &repoDTO.RepoCreateSubscriptionData{
		UserID: cr.UserID,
		Amount: cr.Amount,
	}
}
