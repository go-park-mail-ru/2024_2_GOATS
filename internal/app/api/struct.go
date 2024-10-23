package api

import (
	"mime/multipart"
	"time"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

type RegisterRequest struct {
	Email                string `json:"email"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type UpdateProfileRequest struct {
	UserId     int    `json:"user_id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Birthdate  string `json:"birthdate"`
	Sex        string `json:"sex"`
	Avatar     multipart.File
	AvatarName string
}

type UpdatePasswordRequest struct {
	UserId               int    `json:"user_id"`
	OldPassword          string `json:"oldPassword"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Cookie   string `json:"-"`
}
type SessionResponse struct {
	Success    bool `json:"success"`
	UserData   User `json:"user_data"`
	StatusCode int  `json:"-"`
}

type UpdateUserResponse struct {
	Success    bool `json:"success"`
	StatusCode int  `json:"-"`
}

type AuthResponse struct {
	Success    bool               `json:"success"`
	NewCookie  *models.CookieData `json:"-"`
	StatusCode int                `json:"-"`
}

type CollectionsResponse struct {
	Success     bool         `json:"success"`
	Collections []Collection `json:"collections"`
	StatusCode  int          `json:"-"`
}

type Collection struct {
	Id     int                `json:"id"`
	Title  string             `json:"title"`
	Movies *[]CollectionMovie `json:"movies"`
}

type CollectionMovie struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	CardUrl     string    `json:"card_url"`
	AlbumUrl    string    `json:"album_url"`
	Rating      float32   `json:"rating"`
	ReleaseDate time.Time `json:"release_date"`
	MovieType   string    `json:"movie_type"`
	Country     string    `json:"country"`
}

type MovieResponse struct {
	Success   bool       `json:"success"`
	MovieInfo *MovieInfo `json:"movie_info"`
}

type MovieInfo struct {
	Id               int               `json:"id"`
	Title            string            `json:"title"`
	FullDescription  string            `json:"full_description"`
	ShortDescription string            `json:"short_description"`
	CardUrl          string            `json:"card_url"`
	AlbumUrl         string            `json:"album_url"`
	TitleUrl         string            `json:"title_url"`
	Rating           float32           `json:"rating"`
	ReleaseDate      time.Time         `json:"release_date"`
	MovieType        string            `json:"movie_type"`
	Country          string            `json:"country"`
	VideoUrl         string            `json:"video_url"`
	Actors           []*StaffShortInfo `json:"actors_info"`
	Directors        []*StaffShortInfo `json:"directors_info"`
}

type ActorResponse struct {
	Success   bool   `json:"success"`
	ActorInfo *Actor `json:"actor_info"`
}

type StaffShortInfo struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	PhotoUrl string `json:"photo_url"`
	Country  string `json:"country"`
}

type Actor struct {
	Id        int    `json:"id"`
	FullName  string `json:"full_name"`
	Biography string `json:"biography"`
	Birthdate string `json:"birthdate"`
	PhotoUrl  string `json:"photo_url"`
	Country   string `json:"country"`
}

type ErrorResponse struct {
	Success    bool               `json:"success"`
	StatusCode int                `json:"-"`
	Errors     []errVals.ErrorObj `json:"errors"`
}

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Birthdate string `json:"birthdate"`
	Sex       string `json:"sex"`
	AvatarUrl string `json:"avatar_url"`
}
