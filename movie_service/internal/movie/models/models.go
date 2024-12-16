package models

import (
	"database/sql"
	"fmt"
	"strings"
)

// CollectionsRespData represents collections response
type CollectionsRespData struct {
	Collections []Collection
}

// Collection represents collection full info
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

// DBEpisode represents db episode data
type DBEpisode struct {
	ID            sql.NullInt64
	Title         sql.NullString
	Description   sql.NullString
	EpisodeNumber sql.NullInt64
	ReleaseDate   sql.NullString
	Rating        sql.NullFloat64
	PreviewURL    sql.NullString
	VideoURL      sql.NullString
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

// DirectorInfo represents director full info
type DirectorInfo struct {
	Person
	ID int
}

// Person represents person's name and surname
type Person struct {
	Name    string
	Surname string
}

// Favorite represents favorite's relations
type Favorite struct {
	UserID  int
	MovieID int
}

// FullName returns the person's fullname
func (p Person) FullName() string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", p.Name, p.Surname))
}

// UserRating структура рейтинга
type UserRating struct {
	UserID  int
	MovieID int
	Rating  float64
}
