package api

import (
	"mime/multipart"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

type RegisterRequest struct {
	Email                string `json:"email"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type UpdateProfileRequest struct {
	UserID     int    `json:"user_id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	AvatarFile multipart.File
	AvatarName string
}

type UpdatePasswordRequest struct {
	UserID               int    `json:"user_id"`
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
	UserData User `json:"user_data"`
}

type AuthResponse struct {
	NewCookie *models.CookieData `json:"-"`
}

type CollectionsResponse struct {
	Collections []Collection `json:"collections"`
}

type Collection struct {
	ID     int                      `json:"id"`
	Title  string                   `json:"title"`
	Movies []*models.MovieShortInfo `json:"movies"`
}

type MovieShortInfos struct {
	Movies []models.MovieShortInfo `json:"movies"`
}

type MovieResponse struct {
	MovieInfo *MovieInfo `json:"movie_info"`
}

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
}

type ActorResponse struct {
	ActorInfo *Actor `json:"actor_info"`
}

type ActorInfo struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	PhotoURL string `json:"photo_url"`
	Country  string `json:"country"`
}
type Actor struct {
	ID        int                      `json:"id"`
	FullName  string                   `json:"full_name"`
	Biography string                   `json:"biography"`
	Birthdate string                   `json:"birthdate"`
	PhotoURL  string                   `json:"photo_url"`
	Country   string                   `json:"country"`
	Movies    []*models.MovieShortInfo `json:"movies"`
}

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

type FavReq struct {
	UserID  int `json:"user_id"`
	MovieID int `json:"movie_id"`
}

type CheckReviewData struct {
	CSAT bool `json:"csat_status"`
	NPS  bool `json:"nps_status"`
	CSI  bool `json:"csi_status"`
}

type GetReviewResponse struct {
	Questions []*ReviewData `json:"questions"`
}

type ReviewData struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	Answers []Answer `json:"answers"`
	Type    string   `json:"type"`
}

type Answer struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type CreateReviewRequest struct {
	Type      string     `json:"type"`
	Questions []Question `json:"questions"`
}

type Question struct {
	ID         int    `json:"id"`
	AnswerID   int    `json:"answer_id"`
	AnswerText string `json:"answer_text"`
}

type Statistic struct {
	Rating   float64  `json:"rating"`
	Type     string   `json:"type"`
	Comments []string `json:"comments"`
}

type Time struct {
	Time int `json:"time"`
}
