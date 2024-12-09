package models

import (
	"database/sql"
	"fmt"
	"mime/multipart"
	"strings"
	"time"
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
	UserData User
}

type AuthRespData struct {
	NewCookie *CookieData
}

type CollectionsRespData struct {
	Collections []Collection
}

type User struct {
	ID                         int
	Email                      string
	Username                   string
	Password                   string
	AvatarURL                  string
	AvatarName                 string
	AvatarFile                 multipart.File
	SubscriptionStatus         bool
	SubscriptionExpirationDate string
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
	Seasons          []*Season     `json:"seasons"`
	IsFavorite       bool          `json:"is_favorite"`
}

type Season struct {
	SeasonNumber int        `json:"season_number"`
	Episodes     []*Episode `json:"episodes"`
}

type Episode struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	EpisodeNumber int     `json:"episode_number"`
	ReleaseDate   string  `json:"release_date"`
	Rating        float32 `json:"rating"`
	PreviewURL    string  `json:"preview_url"`
	VideoURL      string  `json:"video_url"`
}

type DBEpisode struct {
	ID            sql.NullInt64   `json:"id"`
	Title         sql.NullString  `json:"title"`
	Description   sql.NullString  `json:"description"`
	EpisodeNumber sql.NullInt64   `json:"episode_number"`
	ReleaseDate   sql.NullString  `json:"release_date"`
	Rating        sql.NullFloat64 `json:"rating"`
	PreviewURL    sql.NullString  `json:"preview_url"`
	VideoURL      sql.NullString  `json:"video_url"`
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

type Favorite struct {
	UserID  int
	MovieID int
}

type SubscriptionData struct {
	UserID int
	Amount uint64
}

type PaymentCallbackData struct {
	NotificationType string
	OperationID      string
	Amount           int64
	Currency         string
	Sender           string
	Label            string
	Unaccepted       bool
}

type CreatePaymentData struct {
	SubscriptionID int
	Amount         uint64
}

func (p Person) FullName() string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", p.Name, p.Surname))
}

type WatchedMovieInfo struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	AlbumURL      string `json:"album_url"`
	TimeCode      int64  `json:"timecode"`
	Duration      int64  `json:"duration"`
	SavingSeconds int64  `json:"saving_seconds"`
}

type OwnWatchedMovie struct {
	UserID       int
	WatchedMovie WatchedMovieInfo
}
