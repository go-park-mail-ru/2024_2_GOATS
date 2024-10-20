package converter

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func ToServLoginData(lg *api.LoginRequest) *models.LoginData {
	if lg == nil {
		return nil
	}

	return &models.LoginData{
		Email:    lg.Email,
		Password: lg.Password,
		Cookie:   lg.Cookie,
	}
}

func ToServRegisterData(rg *api.RegisterRequest) *models.RegisterData {
	if rg == nil {
		return nil
	}

	return &models.RegisterData{
		Email:                rg.Email,
		Username:             rg.Username,
		Password:             rg.Password,
		PasswordConfirmation: rg.PasswordConfirmation,
	}
}

func ToServPasswordData(rp *api.UpdatePasswordRequest) *models.PasswordData {
	return &models.PasswordData{
		UserId:               rp.UserId,
		OldPassword:          rp.OldPassword,
		Password:             rp.Password,
		PasswordConfirmation: rp.PasswordConfirmation,
	}
}

func ToServUserData(pr *api.UpdateProfileRequest) *models.User {
	return &models.User{
		Id:        pr.UserId,
		Email:     pr.Email,
		Username:  pr.Username,
		Birthdate: pr.Birthdate,
		Sex:       pr.Sex,
		AvatarUrl: pr.Avatar,
	}
}

func ToApiAuthResponse(ld *models.AuthRespData) *api.AuthResponse {
	if ld == nil {
		return nil
	}

	return &api.AuthResponse{
		Success:    true,
		NewCookie:  ld.NewCookie,
		StatusCode: ld.StatusCode,
	}
}

func ToApiUpdateUserResponse(ud *models.UpdateUserRespData) *api.UpdateUserResponse {
	if ud == nil {
		return nil
	}

	return &api.UpdateUserResponse{
		Success:    true,
		StatusCode: ud.StatusCode,
	}
}

func ToApiSessionResponse(sr *models.SessionRespData) *api.SessionResponse {
	if sr == nil {
		return nil
	}

	return &api.SessionResponse{
		Success: true,
		UserData: api.User{
			Id:        sr.UserData.Id,
			Email:     sr.UserData.Email,
			Username:  sr.UserData.Username,
			Birthdate: sr.UserData.Birthdate,
			Sex:       sr.UserData.Sex,
		},
		StatusCode: sr.StatusCode,
	}
}

func ToApiCollectionsResponse(cl *models.CollectionsRespData) *api.CollectionsResponse {
	if cl == nil {
		return nil
	}

	return &api.CollectionsResponse{
		Success:     true,
		Collections: cl.Collections,
		StatusCode:  cl.StatusCode,
	}
}

func ToApiErrorResponse(e *models.ErrorRespData) *api.ErrorResponse {
	if e == nil {
		return nil
	}

	return &api.ErrorResponse{
		Success:    false,
		StatusCode: e.StatusCode,
		Errors:     e.Errors,
	}
}
