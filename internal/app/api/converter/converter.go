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
		UserID:               rp.UserID,
		OldPassword:          rp.OldPassword,
		Password:             rp.Password,
		PasswordConfirmation: rp.PasswordConfirmation,
	}
}

func ToServUserData(pr *api.UpdateProfileRequest) *models.User {
	return &models.User{
		ID:         pr.UserID,
		Email:      pr.Email,
		Username:   pr.Username,
		AvatarName: pr.AvatarName,
		AvatarFile: pr.AvatarFile,
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
			ID:        sr.UserData.ID,
			Email:     sr.UserData.Email,
			Username:  sr.UserData.Username,
			AvatarURL: sr.UserData.AvatarURL,
		},
		StatusCode: sr.StatusCode,
	}
}

func ToApiCollectionsResponse(cl *models.CollectionsRespData) *api.CollectionsResponse {
	if cl == nil {
		return nil
	}

	var colls = make([]api.Collection, 0, len(cl.Collections))
	for _, coll := range cl.Collections {
		tempCol := api.Collection{ID: coll.ID, Title: coll.Title, Movies: coll.Movies}
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
		ID:               mv.ID,
		Title:            mv.Title,
		FullDescription:  mv.FullDescription,
		ShortDescription: mv.ShortDescription,
		CardURL:          mv.CardURL,
		AlbumURL:         mv.AlbumURL,
		TitleURL:         mv.TitleURL,
		Rating:           mv.Rating,
		ReleaseDate:      mv.ReleaseDate,
		MovieType:        mv.MovieType,
		Country:          mv.Country,
		VideoURL:         mv.VideoURL,
		Director:         mv.Director.FullName(),
	}

	var actors = make([]*api.ActorInfo, 0, len(mv.Actors))
	for _, actor := range mv.Actors {
		tempSt := &api.ActorInfo{
			ID:       actor.ID,
			FullName: actor.FullName(),
			PhotoURL: actor.SmallPhotoURL,
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
		ID:        ac.ID,
		FullName:  ac.FullName(),
		Biography: ac.Biography,
		PhotoURL:  ac.BigPhotoURL,
		Country:   ac.Country,
		Movies:    ac.Movies,
	}

	if ac.Birthdate.Valid {
		actor.Birthdate = ac.Birthdate.String
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
