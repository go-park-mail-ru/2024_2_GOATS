package converter

import (
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/dto"
)

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

func ToRepoEpisodeFromDB(dbEp *models.DBEpisode) *models.Episode {
	if dbEp == nil {
		return nil
	}

	return &models.Episode{
		ID:            ToIntFromSQLInt(dbEp.ID),
		Title:         ToStringFromSQLString(dbEp.Title),
		Description:   ToStringFromSQLString(dbEp.Description),
		EpisodeNumber: ToIntFromSQLInt(dbEp.EpisodeNumber),
		ReleaseDate:   ToStringFromSQLString(dbEp.ReleaseDate),
		Rating:        ToFloat32FromSQLFloat(dbEp.Rating),
		PreviewURL:    ToStringFromSQLString(dbEp.PreviewURL),
		VideoURL:      ToStringFromSQLString(dbEp.VideoURL),
	}
}

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

func ToIntFromSQLInt(sqlInt sql.NullInt64) int {
	if !sqlInt.Valid {
		return 0
	}

	return int(sqlInt.Int64)
}

func ToFloat32FromSQLFloat(sqlFt sql.NullFloat64) float32 {
	if !sqlFt.Valid {
		return 0
	}

	return float32(sqlFt.Float64)
}

func ToStringFromSQLString(sqlStr sql.NullString) string {
	if !sqlStr.Valid {
		return ""
	}

	return sqlStr.String
}
