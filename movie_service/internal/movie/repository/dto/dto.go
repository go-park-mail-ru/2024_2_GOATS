package dto

import "database/sql"

// RepoMovieShortInfo repo layer movie short info
type RepoMovieShortInfo struct {
	ID          int
	Title       string
	CardURL     string
	AlbumURL    string
	Rating      float32
	ReleaseDate string
	MovieType   string
	Country     string
}

// RepoActor repo layer actor info
type RepoActor struct {
	ID            int
	Name          string
	Surname       string
	Biography     string
	Post          string
	Birthdate     sql.NullString
	SmallPhotoURL string
	BigPhotoURL   string
	Country       string
}
