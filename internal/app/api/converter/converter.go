package converter

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
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
	birthdate, err := StringToNullTime(pr.Birthdate)
	if err != nil {
		return nil
	}

	return &models.User{
		Id:         pr.UserId,
		Email:      pr.Email,
		Username:   pr.Username,
		Birthdate:  birthdate,
		Sex:        StringToNullString(pr.Sex),
		AvatarName: pr.AvatarName,
		Avatar:     pr.Avatar,
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

	resp := &api.SessionResponse{
		Success: true,
		UserData: api.User{
			Id:        sr.UserData.Id,
			Email:     sr.UserData.Email,
			Username:  sr.UserData.Username,
			AvatarUrl: sr.UserData.AvatarUrl,
		},
		StatusCode: sr.StatusCode,
	}

	if sr.UserData.Birthdate.Valid {
		resp.UserData.Birthdate = sr.UserData.Birthdate.Time.Format("2006-01-02")
	} else {
		resp.UserData.Birthdate = ""
	}

	if sr.UserData.Sex.Valid {
		resp.UserData.Sex = sr.UserData.Sex.String
	} else {
		resp.UserData.Sex = ""
	}

	return resp
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

func StringToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func StringToNullTime(s string) (sql.NullTime, error) {
	if s == "" {
		return sql.NullTime{Valid: false}, nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		errMsg := fmt.Errorf("cannot parse given string to date: %w", err)
		log.Err(errMsg)

		return sql.NullTime{}, errMsg
	}
	return sql.NullTime{Time: t, Valid: true}, nil
}
