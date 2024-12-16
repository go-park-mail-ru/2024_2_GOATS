package models

import (
	"database/sql"
	"mime/multipart"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	model "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

// Room структура комнаты
type Room struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Movie   string `json:"movie_service"`
	AdminID string `json:"admin_id"`
}

// MovieInfo структура фильмов
type MovieInfo struct {
	ID               int          `json:"id"`
	Title            string       `json:"title"`
	ShortDescription string       `json:"short_description"`
	FullDescription  string       `json:"full_description"`
	CardURL          string       `json:"card_url"`
	AlbumURL         string       `json:"album_url"`
	TitleURL         string       `json:"title_url"`
	Rating           float32      `json:"rating"`
	ReleaseDate      string       `json:"release_date"`
	MovieType        string       `json:"movie_type"`
	Country          string       `json:"country"`
	VideoURL         string       `json:"video_url"`
	Actors           []*StaffInfo `json:"actors_info"`
	Director         *StaffInfo   `json:"director_info"`
	Seasons          []*Season    `json:"seasons"`
	IsFavorite       bool         `json:"is_favorite"`
}

// Season структура сезонов
type Season struct {
	SeasonNumber int        `json:"season_number"`
	Episodes     []*Episode `json:"episodes"`
}

// Episode структура эпизодов
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

// StaffInfo структура актеров
type StaffInfo struct {
	ID            int          `json:"id"`
	Name          string       `json:"name"`
	Surname       string       `json:"surname"`
	Patronymic    string       `json:"patronymic"`
	Biography     string       `json:"biography"`
	Post          string       `json:"post"`
	Birthdate     sql.NullTime `json:"birthdate"`
	SmallPhotoURL string       `json:"small_photo_url"`
	BigPhotoURL   string       `json:"big_photo_url"`
	Country       string       `json:"country"`
}

// RoomState структура данных о комнате
type RoomState struct {
	ID         string    `json:"id"`
	Status     string    `json:"status"` // paused, playing
	TimeCode   float64   `json:"time_code"`
	Movie      MovieInfo `json:"movie"`
	Message    Msg       `json:"message"`
	Duration   int       `json:"duration"`
	SeasonNow  int       `json:"season_now"`
	EpisodeNow int       `json:"episode_now"`
	//TimerQuit chan struct{} `json:"timerQuit"`
}

// Action структура события
type Action struct {
	Name       string    `json:"name"` // pause, play, rewind
	TimeCode   float64   `json:"time_code"`
	Message    Msg       `json:"message"`
	MovieID    int       `json:"movie_id"`
	Movie      MovieInfo `json:"movie"`
	SeasonNow  int       `json:"season_number"`
	EpisodeNow int       `json:"episode_number"`
	Duration   int       `json:"duration"`
}

// Msg структура сообщения
type Msg struct {
	Text   string `json:"text"` // pause, play, rewind
	Sender string `json:"sender"`
	Avatar string `json:"avatar"`
}

// User структура юзера
type User struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	AvatarURL  string `json:"avatar_url"`
	AvatarName string `json:"avatar_name"`
	AvatarFile multipart.File
}

// SessionRespData структура ответа
type SessionRespData struct {
	UserData model.User `json:"user_data"`
}

// ErrorRespData структура ошибки
type ErrorRespData struct {
	StatusCode int
	Errors     []errVals.RepoError
}
