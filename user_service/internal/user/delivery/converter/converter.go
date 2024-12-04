package converter

import (
	srvDTO "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
)

func ConvertToSrvCreateUser(req *user.CreateUserRequest) *srvDTO.CreateUserData {
	if req == nil {
		return nil
	}

	return &srvDTO.CreateUserData{
		Email:                req.Email,
		Username:             req.Username,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}
}

func ConvertToSrvCreateSubscription(req *user.CreateSubscriptionRequest) *srvDTO.CreateSubscriptionData {
	if req == nil {
		return nil
	}

	return &srvDTO.CreateSubscriptionData{
		UserID: req.UserID,
		Amount: req.Amount,
	}
}

func ConvertToSrvUpdatePassword(req *user.UpdatePasswordRequest) *srvDTO.PasswordData {
	if req == nil {
		return nil
	}

	return &srvDTO.PasswordData{
		UserID:               req.UserID,
		OldPassword:          req.OldPassword,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}
}

func ConvertToSrvUpdateProfile(req *user.UserData) *srvDTO.User {
	if req == nil {
		return nil
	}

	return &srvDTO.User{
		ID:         req.UserID,
		Email:      req.Email,
		Username:   req.Username,
		AvatarURL:  req.AvatarURL,
		AvatarName: req.AvatarName,
		AvatarFile: req.AvatarFile,
	}
}

func ConvertToSrvFavorite(req *user.HandleFavorite) *srvDTO.Favorite {
	if req == nil {
		return nil
	}

	return &srvDTO.Favorite{
		UserID:  req.UserID,
		MovieID: req.MovieID,
	}
}

func ConvertToGRPCUser(su *srvDTO.User) *user.UserData {
	if su == nil {
		return nil
	}

	return &user.UserData{
		UserID:     su.ID,
		Email:      su.Email,
		Username:   su.Username,
		Password:   su.Password,
		AvatarURL:  su.AvatarURL,
		AvatarName: su.AvatarName,
	}
}
