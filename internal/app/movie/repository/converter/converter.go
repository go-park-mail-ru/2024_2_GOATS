package converter

import (
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

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

	return string(sqlStr.String)
}
