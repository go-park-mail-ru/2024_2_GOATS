package models

import (
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
	Id        int
	Email     string
	Username  string
	Password  string
	Birthdate string
	AvatarUrl string
	Sex       string
}

type Collection struct {
	Id     int
	Title  string
	Movies []*Movie
}

type Movie struct {
	Id          int
	Title       string
	Description string
	CardUrl     string
	AlbumUrl    string
	Rating      float32
	ReleaseDate time.Time
	MovieType   string
	Country     string
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
