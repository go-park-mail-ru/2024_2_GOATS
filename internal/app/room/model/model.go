package models

import errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"

type Room struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Movie   string `json:"movie"` // ID фильма или URL
	AdminID string `json:"admin_id"`
}

type Movie struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	TitleImage       string `json:"titleImage"`
	ShortDescription string `json:"shortDescription"`
	LongDescription  string `json:"longDescription"`
	Image            string `json:"image"`
	Rating           int    `json:"rating"`
	ReleaseDate      string `json:"releaseDate"`
	Country          string `json:"country"`
	Director         string `json:"director"`
	IsSerial         bool   `json:"isSerial"`
	Video            string `json:"video"`
}

type RoomState struct {
	Id       string  `json:"id"`        // paused, playing
	Status   string  `json:"status"`    // paused, playing
	TimeCode float64 `json:"time_code"` // Текущий момент фильма
	Movie    Movie   `json:"movie"`
}

type Action struct {
	Name     string  `json:"name"`      // pause, play, rewind
	TimeCode float64 `json:"time_code"` // Время, на которое надо перемотать
}

type User struct {
	Id       int
	Email    string
	Username string
	Password string
}

type SessionRespData struct {
	UserData   User
	StatusCode int
}

type ErrorRespData struct {
	StatusCode int
	Errors     []errVals.ErrorObj
}
