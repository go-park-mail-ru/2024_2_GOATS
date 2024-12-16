package models

// Movie data for elasticsearch
type Movie struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	CardURL     string  `json:"card_url"`
	AlbumURL    string  `json:"album_url"`
	Rating      float64 `json:"rating"`
	ReleaseDate string  `json:"release_date"`
	MovieType   string  `json:"movie_type"`
	Country     string  `json:"country"`
}

// Actor data for elasticsearch
type Actor struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name"`
	PhotoURL    string `json:"photo_url"`
	PhotoBigURL string `json:"photo_big_url"`
	Biography   string `json:"biography"`
	Country     string `json:"country"`
	BirthDate   string `json:"birth_date"`
}
