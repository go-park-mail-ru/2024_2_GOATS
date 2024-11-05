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
	ID         int
	Email      string
	Username   string
	Password   string
	AvatarURL  string
	AvatarName string
	AvatarFile multipart.File
}

type Collection struct {
	ID     int               `json:"id"`
	Title  string            `json:"title"`
	Movies []*MovieShortInfo `json:"movies"`
}

type MovieInfo struct {
	ID               int           `json:"id"`
	Title            string        `json:"title"`
	ShortDescription string        `json:"short_description"`
	FullDescription  string        `json:"full_description"`
	CardURL          string        `json:"card_url"`
	AlbumURL         string        `json:"album_url"`
	TitleURL         string        `json:"title_url"`
	Rating           float32       `json:"rating"`
	ReleaseDate      string        `json:"release_date"`
	MovieType        string        `json:"movie_type"`
	Country          string        `json:"country"`
	VideoURL         string        `json:"video_url"`
	Actors           []*ActorInfo  `json:"actors_info"`
	Director         *DirectorInfo `json:"director_info"`
}

type MovieShortInfo struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	CardURL     string  `json:"card_url"`
	AlbumURL    string  `json:"album_url"`
	Rating      float32 `json:"rating"`
	ReleaseDate string  `json:"release_date"`
	MovieType   string  `json:"movie_type"`
	Country     string  `json:"country"`
}

type ActorInfo struct {
	Person
	ID            int               `json:"id"`
	Biography     string            `json:"biography"`
	Post          string            `json:"post"`
	Birthdate     sql.NullString    `json:"birthdate"`
	SmallPhotoURL string            `json:"small_photo_url"`
	BigPhotoURL   string            `json:"big_photo_url"`
	Country       string            `json:"country"`
	Movies        []*MovieShortInfo `json:"movies"`
}

type DirectorInfo struct {
	Person
	ID int
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
	UserID               int
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
