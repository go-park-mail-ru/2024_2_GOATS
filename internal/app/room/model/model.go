package models

import (
	"database/sql"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	model "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"mime/multipart"
)

// TODO раскоментить к 4му РК

type Room struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Movie   string `json:"movie_service"`
	AdminID string `json:"admin_id"`
}

//type Movie struct {
//	Id               int    `json:"id"`
//	Title            string `json:"title"`
//	TitleImage       string `json:"titleImage"`
//	ShortDescription string `json:"short_description"`
//	LongDescription  string `json:"full_description"`
//	Image            string `json:"image"`
//	Rating           int    `json:"rating"`
//	ReleaseDate      string `json:"releaseDate"`
//	Country          string `json:"country"`
//	Director         string `json:"director"`
//	IsSerial         bool   `json:"isSerial"`
//	Video            string `json:"video"`
//}

// ////////////Movie теперь такой
//type MovieInfo struct {
//	Id               int          `json:"id"`
//	Title            string       `json:"title"`
//	ShortDescription string       `json:"short_description"`
//	FullDescription  string       `json:"full_description"`
//	CardUrl          string       `json:"card_url"`
//	AlbumUrl         string       `json:"album_url"`
//	TitleUrl         string       `json:"title_url"`
//	Rating           float32      `json:"rating"`
//	ReleaseDate      time.Time    `json:"release_date"`
//	MovieType        string       `json:"movie_type"`
//	Country          string       `json:"country"`
//	VideoUrl         string       `json:"video_url"`
//	Actors           []*StaffInfo `json:"actors_info"`
//	Directors        []*StaffInfo `json:"directors_info"`
//}

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

type StaffInfo struct {
	Id            int          `json:"id"`
	Name          string       `json:"name"`
	Surname       string       `json:"surname"`
	Patronymic    string       `json:"patronymic"`
	Biography     string       `json:"biography"`
	Post          string       `json:"post"`
	Birthdate     sql.NullTime `json:"birthdate"`
	SmallPhotoUrl string       `json:"small_photo_url"`
	BigPhotoUrl   string       `json:"big_photo_url"`
	Country       string       `json:"country"`
}

type ActorInfo struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	PhotoURL string `json:"photo_url"`
	Country  string `json:"country"`
}

type RoomState struct {
	Id         string    `json:"id"`
	Status     string    `json:"status"` // paused, playing
	TimeCode   float64   `json:"time_code"`
	Movie      MovieInfo `json:"movie"`
	Message    Msg       `json:"message"`
	Duration   int       `json:"duration"`
	SeasonNow  int       `json:"season_now"`
	EpisodeNow int       `json:"episode_now"`
	//TimerQuit chan struct{} `json:"timerQuit"`
}

type Action struct {
	Name       string    `json:"name"` // pause, play, rewind
	TimeCode   float64   `json:"time_code"`
	Message    Msg       `json:"message"`
	MovieId    int       `json:"movie_id"`
	Movie      MovieInfo `json:"movie"`
	SeasonNow  int       `json:"season_number"`
	EpisodeNow int       `json:"episode_number"`
	Duration   int       `json:"duration"`
}

//	type ActionMsg struct {
//		Name     string  `json:"name"` // pause, play, rewind
//		Msg  string  `json:"message"`
//	}
type Msg struct {
	Text   string `json:"text"` // pause, play, rewind
	Sender string `json:"sender"`
	Avatar string `json:"avatar"`
}

//type User struct {
//	Id         int            `json:"id"`
//	Email      string         `json:"email"`
//	Username   string         `json:"username"`
//	Password   string         `json:"password"`
//	Birthdate  sql.NullTime   `json:"birthdate"`
//	AvatarUrl  string         `json:"avatar_url"`
//	AvatarName string         `json:"avatar_name"`
//	Avatar     multipart.File `json:"avatar"`
//	Sex        sql.NullString `json:"sex"`
//}

//type User struct {
//	ID         int
//	Email      string
//	Username   string
//	Password   string
//	AvatarURL  string
//	AvatarName string
//	AvatarFile multipart.File
//}

type User struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	AvatarURL  string `json:"avatar_url"`
	AvatarName string `json:"avatar_name"`
	AvatarFile multipart.File
}

type SessionRespData struct {
	UserData model.User `json:"user_data"`
}

type ErrorRespData struct {
	StatusCode int
	Errors     []errVals.RepoError
}
