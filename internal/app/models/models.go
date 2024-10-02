package models

import (
	"time"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type Collection struct {
	Id     int      `json:"id"`
	Title  string   `json:"title"`
	Movies []*Movie `json:"movies"`
}

type Movie struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CardUrl     string    `json:"card_image"`
	AlbumUrl    string    `json:"album_image"`
	Rating      float32   `json:"rating"`
	ReleaseDate time.Time `json:"release_date"`
	MovieType   string    `json:"movie_type"`
	Country     string    `json:"country"`
}

type CollectionsResponse struct {
	Success     bool         `json:"success"`
	Collections []Collection `json:"collections"`
}

type ErrorResponse struct {
	Success    bool               `json:"success"`
	StatusCode int                `json:"-"`
	Errors     []errVals.ErrorObj `json:"errors"`
}
