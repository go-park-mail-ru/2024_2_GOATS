package converter

import (
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/dto"
)

// ToActorInfoFromRepo converts from dto Actor to models Actor
func ToActorInfoFromRepo(ac *dto.RepoActor) *models.ActorInfo {
	if ac == nil {
		return nil
	}

	return &models.ActorInfo{
		ID: ac.ID,
		Person: models.Person{
			Name:    ac.Name,
			Surname: ac.Surname,
		},
		Biography:     ac.Biography,
		Post:          ac.Post,
		Birthdate:     ac.Birthdate,
		SmallPhotoURL: ac.SmallPhotoURL,
		BigPhotoURL:   ac.BigPhotoURL,
		Country:       ac.Country,
	}
}

// ToRepoEpisodeFromDB converts from db Episode to models Episode
func ToRepoEpisodeFromDB(dbEp *models.DBEpisode) *models.Episode {
	if dbEp == nil {
		return nil
	}

	return &models.Episode{
		ID:            toIntFromSQLInt(dbEp.ID),
		Title:         toStringFromSQLString(dbEp.Title),
		Description:   toStringFromSQLString(dbEp.Description),
		EpisodeNumber: toIntFromSQLInt(dbEp.EpisodeNumber),
		ReleaseDate:   toStringFromSQLString(dbEp.ReleaseDate),
		Rating:        toFloat32FromSQLFloat(dbEp.Rating),
		PreviewURL:    toStringFromSQLString(dbEp.PreviewURL),
		VideoURL:      toStringFromSQLString(dbEp.VideoURL),
	}
}

// ToMovieShortInfoFromRepo converts from repo MovieShortInfo to models MovieShortInfo
func ToMovieShortInfoFromRepo(m *dto.RepoMovieShortInfo) *models.MovieShortInfo {
	if m == nil {
		return nil
	}

	return &models.MovieShortInfo{
		ID:          m.ID,
		Title:       m.Title,
		CardURL:     m.CardURL,
		AlbumURL:    m.AlbumURL,
		Rating:      m.Rating,
		ReleaseDate: m.ReleaseDate,
		MovieType:   m.MovieType,
		Country:     m.Country,
	}
}

func toIntFromSQLInt(sqlInt sql.NullInt64) int {
	if !sqlInt.Valid {
		return 0
	}

	return int(sqlInt.Int64)
}

func toFloat32FromSQLFloat(sqlFt sql.NullFloat64) float32 {
	if !sqlFt.Valid {
		return 0
	}

	return float32(sqlFt.Float64)
}

func toStringFromSQLString(sqlStr sql.NullString) string {
	if !sqlStr.Valid {
		return ""
	}

	return sqlStr.String
}
