package converter

import (
	srvDTO "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
	user "github.com/go-park-mail-ru/2024_2_GOATS/user_service/pkg/user_v1"
	"github.com/microcosm-cc/bluemonday"
)

func ConvertToSrvCreateUser(req *user.CreateUserRequest) *srvDTO.CreateUserData {
	if req == nil {
		return nil
	}

	return &srvDTO.CreateUserData{
		Email:                sanitizeInput(req.Email),
		Username:             sanitizeInput(req.Username),
		Password:             sanitizeInput(req.Password),
		PasswordConfirmation: sanitizeInput(req.PasswordConfirmation),
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
		OldPassword:          sanitizeInput(req.OldPassword),
		Password:             sanitizeInput(req.Password),
		PasswordConfirmation: sanitizeInput(req.PasswordConfirmation),
	}
}

func ConvertToSrvUpdateProfile(req *user.UserData) *srvDTO.User {
	if req == nil {
		return nil
	}

	return &srvDTO.User{
		ID:         req.UserID,
		Email:      sanitizeInput(req.Email),
		Username:   sanitizeInput(req.Username),
		AvatarURL:  sanitizeInput(req.AvatarURL),
		AvatarName: sanitizeInput(req.AvatarName),
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
		UserID:                     su.ID,
		Email:                      su.Email,
		Username:                   su.Username,
		Password:                   su.Password,
		AvatarURL:                  su.AvatarURL,
		AvatarName:                 su.AvatarName,
		SubscriptionStatus:         su.SubscriptionStatus,
		SubscriptionExpirationDate: su.SubscriptionExpirationDate,
	}
}

func sanitizeInput(input string) string {
	policy := bluemonday.UGCPolicy()
	return policy.Sanitize(input)
}
