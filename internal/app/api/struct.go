package api

import (
	"mime/multipart"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

// RegisterRequest json struct
//
//go:generate easyjson -all struct.go
type RegisterRequest struct {
	Email                string `json:"email"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

// UpdateProfileRequest json struct
type UpdateProfileRequest struct {
	UserID     int            `json:"user_id"`
	Email      string         `json:"email"`
	Username   string         `json:"username"`
	AvatarFile multipart.File `json:"-"`
	AvatarName string         `json:"-"`
}

// UpdatePasswordRequest json struct
type UpdatePasswordRequest struct {
	UserID               int    `json:"user_id"`
	OldPassword          string `json:"oldPassword"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

// LoginRequest json struct
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Cookie   string `json:"-"`
}

// SessionResponse json struct
type SessionResponse struct {
	UserData User `json:"user_data"`
}

// AuthResponse json struct
type AuthResponse struct {
	NewCookie *models.CookieData `json:"-"`
}

// CollectionsResponse json struct
type CollectionsResponse struct {
	Collections []Collection `json:"collections"`
}

// Collection json struct
type Collection struct {
	ID     int                      `json:"id"`
	Title  string                   `json:"title"`
	Movies []*models.MovieShortInfo `json:"movies"`
}

// MovieShortInfos json struct
type MovieShortInfos struct {
	Movies []models.MovieShortInfo `json:"movies"`
}

// MovieResponse json struct
type MovieResponse struct {
	MovieInfo *MovieInfo `json:"movie_info"`
	//Rating    int64      `json:"rating_info"`
}

// MovieInfo json struct
type MovieInfo struct {
	ID               int              `json:"id"`
	Title            string           `json:"title"`
	FullDescription  string           `json:"full_description"`
	ShortDescription string           `json:"short_description"`
	CardURL          string           `json:"card_url"`
	AlbumURL         string           `json:"album_url"`
	TitleURL         string           `json:"title_url"`
	Rating           float32          `json:"rating"`
	ReleaseDate      string           `json:"release_date"`
	MovieType        string           `json:"movie_type"`
	Country          string           `json:"country"`
	VideoURL         string           `json:"video_url"`
	Director         string           `json:"director"`
	Actors           []*ActorInfo     `json:"actors_info"`
	Seasons          []*models.Season `json:"seasons"`
	IsFavorite       bool             `json:"is_favorite"`
	WithSubscription bool             `json:"with_subscription"`
	RatingUser       int64            `json:"rating_user"`
}

// ActorResponse json struct
type ActorResponse struct {
	ActorInfo *Actor `json:"actor_info"`
}

// ActorInfo json struct
type ActorInfo struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	PhotoURL string `json:"photo_url"`
	Country  string `json:"country"`
}

// Actor json struct
type Actor struct {
	ID        int                      `json:"id"`
	FullName  string                   `json:"full_name"`
	Biography string                   `json:"biography"`
	Birthdate string                   `json:"birthdate"`
	PhotoURL  string                   `json:"photo_url"`
	Country   string                   `json:"country"`
	Movies    []*models.MovieShortInfo `json:"movies"`
}

// User json struct
type User struct {
	ID                         int    `json:"id"`
	Email                      string `json:"email"`
	Username                   string `json:"username"`
	AvatarURL                  string `json:"avatar_url"`
	SubscriptionStatus         bool   `json:"subscription_status"`
	SubscriptionExpirationDate string `json:"subscription_expiration_date"`
}

// FavReq json struct
type FavReq struct {
	UserID  int `json:"user_id"`
	MovieID int `json:"movie_id"`
}

// PaymentCallback json struct
type PaymentCallback struct {
	NotificationType string    `json:"notification_type"`
	OperationID      string    `json:"operation_id"`
	Amount           int64     `json:"amount"`
	WithdrawAmount   int64     `json:"withdraw_amount"`
	Currency         string    `json:"currency"`
	DateTime         time.Time `json:"date_time"`
	Sender           string    `json:"sender"`
	Codepro          bool      `json:"codepro"`
	Label            string    `json:"label"`
	Signature        string    `json:"sha1_hash"`
	Unaccepted       bool      `json:"unaccepted"`
}

// SubscriptionStatus json struct
type SubscriptionStatus struct {
	Status string `json:"status"`
}

// SubscribeRequest json struct
type SubscribeRequest struct {
	Amount int64 `json:"amount"`
}

// SubscribeResponse json struct
type SubscribeResponse struct {
	SubscriptionIDP string `json:"subscription_idp"`
}

// MovieSearchData json struct
type MovieSearchData struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	CardURL     string `json:"card_url"`
	AlbumURL    string `json:"album_url"`
	Rating      string `json:"rating"`
	ReleaseDate string `json:"release_date"`
	MovieType   string `json:"movie_type"`
	Country     string `json:"country"`
}

// MovieSearchList type
// easyjson:json
type MovieSearchList []MovieSearchData

// ActorSearchData json struct
type ActorSearchData struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	PhotoURL string `json:"photo_url"`
	Country  string `json:"country"`
}

// ActorSearchList type
// easyjson:json
type ActorSearchList []ActorSearchData

// AddOrUpdateRatingReq json struct
type AddOrUpdateRatingReq struct {
	Rating int `json:"rating"`
}
