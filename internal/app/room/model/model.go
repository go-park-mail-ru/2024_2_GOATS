package models

import (
	"database/sql"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"mime/multipart"
)

type Room struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Movie   string `json:"movie"` // ID фильма или URL
	AdminID string `json:"admin_id"`
}

type Movie struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	TitleImage       string `json:"titleImage"`
	ShortDescription string `json:"shortDescription"`
	LongDescription  string `json:"longDescription"`
	Image            string `json:"image"`
	Rating           int    `json:"rating"`
	ReleaseDate      string `json:"releaseDate"`
	Country          string `json:"country"`
	Director         string `json:"director"`
	IsSerial         bool   `json:"isSerial"`
	Video            string `json:"video"`
}

type RoomState struct {
	Id       string  `json:"id"`
	Status   string  `json:"status"` // paused, playing
	TimeCode float64 `json:"time_code"`
	Movie    Movie   `json:"movie"`
	Message  string  `json:"message"`
}

type Action struct {
	Name     string  `json:"name"` // pause, play, rewind
	TimeCode float64 `json:"time_code"`
	Message  string  `json:"message"`
}

type User struct {
	Id         int            `json:"id"`
	Email      string         `json:"email"`
	Username   string         `json:"username"`
	Password   string         `json:"password"`
	Birthdate  sql.NullTime   `json:"birthdate"`
	AvatarUrl  string         `json:"avatar_url"`
	AvatarName string         `json:"avatar_name"`
	Avatar     multipart.File `json:"avatar"`
	Sex        sql.NullString `json:"sex"`
}

type SessionRespData struct {
	UserData   User `json:"user_data"`
	StatusCode int  `json:"status_code"`
}

type ErrorRespData struct {
	StatusCode int
	Errors     []errVals.ErrorObj
}
