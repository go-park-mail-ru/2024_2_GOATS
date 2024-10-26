package models

import (
	"database/sql"
	"mime/multipart"
	"time"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

type LoginData struct {
	Email    string
	Password string
	Cookie   string
}

type RegisterData struct {
	Email                string
	Username             string
	Password             string
	PasswordConfirmation string
}

type SessionRespData struct {
	UserData   User
	StatusCode int
}

type AuthRespData struct {
	NewCookie  *CookieData
	StatusCode int
}

type UpdateUserRespData struct {
	StatusCode int
}

type CollectionsRespData struct {
	Collections []Collection
	StatusCode  int
}

type ErrorRespData struct {
	StatusCode int
	Errors     []errVals.ErrorObj
}

type User struct {
	Id         int
	Email      string
	Username   string
	Password   string
	Birthdate  sql.NullTime
	AvatarUrl  string
	AvatarName string
	Avatar     multipart.File
	Sex        sql.NullString
}

type Collection struct {
	Id     int      `json:"id"`
	Title  string   `json:"title"`
	Movies []*Movie `json:"movies"`
}

type Movie struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CardUrl     string    `json:"card_url"`
	AlbumUrl    string    `json:"album_url"`
	Rating      float32   `json:"rating"`
	ReleaseDate time.Time `json:"release_date"`
	MovieType   string    `json:"movie_type"`
	Country     string    `json:"country"`
}

type CookieData struct {
	Name  string
	Token *Token
}

type Token struct {
	UserID  int
	TokenID string
	Expiry  time.Time
}

type PasswordData struct {
	UserId               int
	OldPassword          string
	Password             string
	PasswordConfirmation string
}
