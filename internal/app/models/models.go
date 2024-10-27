package models

import (
	"database/sql"
	"fmt"
	"mime/multipart"
	"strings"
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
	Id     int               `json:"id"`
	Title  string            `json:"title"`
	Movies []*MovieShortInfo `json:"movies"`
}

type MovieInfo struct {
	Id               int           `json:"id"`
	Title            string        `json:"title"`
	ShortDescription string        `json:"short_description"`
	FullDescription  string        `json:"full_description"`
	CardUrl          string        `json:"card_url"`
	AlbumUrl         string        `json:"album_url"`
	TitleUrl         string        `json:"title_url"`
	Rating           float32       `json:"rating"`
	ReleaseDate      time.Time     `json:"release_date"`
	MovieType        string        `json:"movie_type"`
	Country          string        `json:"country"`
	VideoUrl         string        `json:"video_url"`
	Actors           []*ActorInfo  `json:"actors_info"`
	Director         *DirectorInfo `json:"director_info"`
}

type MovieShortInfo struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	CardUrl     string    `json:"card_url"`
	Rating      float32   `json:"rating"`
	ReleaseDate time.Time `json:"release_date"`
	MovieType   string    `json:"movie_type"`
	Country     string    `json:"country"`
}

type ActorInfo struct {
	Person
	Id            int               `json:"id"`
	Biography     string            `json:"biography"`
	Post          string            `json:"post"`
	Birthdate     sql.NullTime      `json:"birthdate"`
	SmallPhotoUrl string            `json:"small_photo_url"`
	BigPhotoUrl   string            `json:"big_photo_url"`
	Country       string            `json:"country"`
	Movies        []*MovieShortInfo `json:"movies"`
}

type DirectorInfo struct {
	Person
	Id int
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

type Person struct {
	Name    string
	Surname string
}

func (p Person) FullName() string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", p.Name, p.Surname))
}
