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
		Id:         pr.UserId,
		Email:      pr.Email,
		Username:   pr.Username,
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

	return resp
}

func ToApiCollectionsResponse(cl *models.CollectionsRespData) *api.CollectionsResponse {
	if cl == nil {
		return nil
	}

	colls := []api.Collection{}
	for _, coll := range cl.Collections {
		tempCol := api.Collection{Id: coll.Id, Title: coll.Title, Movies: coll.Movies}
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
		Director:         mv.Director.FullName(),
	}

	actors := []*api.ActorInfo{}
	for _, actor := range mv.Actors {
		tempSt := &api.ActorInfo{
			Id:       actor.Id,
			FullName: actor.FullName(),
			PhotoUrl: actor.SmallPhotoUrl,
			Country:  actor.Country,
		}

		actors = append(actors, tempSt)
	}

	mvInfo.Actors = actors

	return &api.MovieResponse{
		Success:   true,
		MovieInfo: mvInfo,
	}
}

func ToApiGetActorResponse(ac *models.ActorInfo) *api.ActorResponse {
	if ac == nil {
		return nil
	}

	actor := &api.Actor{
		Id:        ac.Id,
		FullName:  ac.FullName(),
		Biography: ac.Biography,
		PhotoUrl:  ac.BigPhotoUrl,
		Country:   ac.Country,
		Movies:    ac.Movies,
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
