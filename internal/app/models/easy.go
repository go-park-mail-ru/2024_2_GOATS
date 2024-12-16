package models

import "database/sql"

// Collection represents collection full info
//
//go:generate easyjson -all easy.go
type Collection struct {
	ID     int               `json:"id"`
	Title  string            `json:"title"`
	Movies []*MovieShortInfo `json:"movies"`
}

// MovieInfo represents movie full info
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
	WithSubscription bool          `json:"with_subscription"`
}

// Season represents movie's season full info
type Season struct {
	SeasonNumber int        `json:"season_number"`
	Episodes     []*Episode `json:"episodes"`
}

// Episode represents season's episode full info
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

// MovieShortInfo represents movie_short_info data
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

// ActorInfo represents actor full info
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
