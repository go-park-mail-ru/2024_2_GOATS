package favoritedb

import (
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/dto"
	"github.com/rs/zerolog/log"
)

func ScanConnections(rows *sql.Rows) ([]*dto.DBMovieShortInfo, error) {
	defer func() {
		if err := rows.Close(); err != nil {
			errMsg := fmt.Errorf("cannot close rows while taking user favorites: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())
		}
	}()

	var movies []*dto.DBMovieShortInfo

	for rows.Next() {
		var movie dto.DBMovieShortInfo

		err := rows.Scan(&movie.ID, &movie.Title, &movie.CardURL, &movie.AlbumURL, &movie.Rating, &movie.ReleaseDate, &movie.MovieType, &movie.Country)

		if err != nil {
			errMsg := fmt.Errorf("error while scanning favorite movies: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())

			return nil, errMsg
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}
