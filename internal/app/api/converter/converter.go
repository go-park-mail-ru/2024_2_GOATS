package converter

import (
	"database/sql"
	"fmt"
	"strings"
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

	colls := []api.Collection{}
	for _, coll := range cl.Collections {
		tempCol := api.Collection{Id: coll.Id, Title: coll.Title, Movies: &[]api.CollectionMovie{}}
		for _, movie := range coll.Movies {
			tempMv := api.CollectionMovie{
				Id:          movie.Id,
				Title:       movie.Title,
				CardUrl:     movie.CardUrl,
				AlbumUrl:    movie.AlbumUrl,
				Rating:      movie.Rating,
				ReleaseDate: movie.ReleaseDate,
				MovieType:   movie.MovieType,
				Country:     movie.Country,
			}

			*tempCol.Movies = append(*tempCol.Movies, tempMv)
		}

		colls = append(colls, tempCol)
	}

	return &api.CollectionsResponse{
		Success:     true,
		Collections: colls,
		StatusCode:  cl.StatusCode,
	}
}

func ToApiGetMovieResponse(mv *models.MovieInfo) *api.MovieResponse {
	if mv == nil {
		return nil
	}

	mvInfo := &api.MovieInfo{
		Id:               mv.Id,
		Title:            mv.Title,
		FullDescription:  mv.FullDescription,
		ShortDescription: mv.ShortDescription,
		CardUrl:          mv.CardUrl,
		AlbumUrl:         mv.AlbumUrl,
		TitleUrl:         mv.TitleUrl,
		Rating:           mv.Rating,
		ReleaseDate:      mv.ReleaseDate,
		MovieType:        mv.MovieType,
		Country:          mv.Country,
		VideoUrl:         mv.VideoUrl,
	}

	actors := []*api.StaffShortInfo{}
	directors := []*api.StaffShortInfo{}
	staffs := mv.Actors
	staffs = append(staffs, mv.Directors...)

	for _, staff := range staffs {
		tempSt := &api.StaffShortInfo{
			Id:       staff.Id,
			FullName: strings.TrimSpace(fmt.Sprintf("%s %s %s", staff.Name, staff.Surname, staff.Patronymic)),
			PhotoUrl: staff.SmallPhotoUrl,
			Country:  staff.Country,
		}

		if staff.Post == "actor" {
			actors = append(actors, tempSt)
		}

		if staff.Post == "director" {
			directors = append(directors, tempSt)
		}
	}

	mvInfo.Actors = actors
	mvInfo.Directors = directors

	return &api.MovieResponse{
		Success:   true,
		MovieInfo: mvInfo,
	}
}

func ToApiGetActorResponse(ac *models.StaffInfo) *api.ActorResponse {
	if ac == nil {
		return nil
	}

	actor := &api.Actor{
		Id:        ac.Id,
		FullName:  strings.TrimSpace(fmt.Sprintf("%s %s %s", ac.Name, ac.Surname, ac.Patronymic)),
		Biography: ac.Biography,
		PhotoUrl:  ac.BigPhotoUrl,
		Country:   ac.Country,
	}

	if ac.Birthdate.Valid {
		actor.Birthdate = ac.Birthdate.Time.Format("2006-01-02")
	} else {
		actor.Birthdate = ""
	}

	return &api.ActorResponse{
		Success:   true,
		ActorInfo: actor,
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
