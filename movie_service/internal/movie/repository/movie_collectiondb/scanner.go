package movie_collectiondb

import (
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	"github.com/rs/zerolog/log"
)

// ScanConnections scans movie_collections rows
func ScanConnections(rows *sql.Rows) (map[int]models.Collection, error) {
	defer closeRows(rows, "movie_collections")
	collections := make(map[int]models.Collection, 3)

	for rows.Next() {
		var collectionID int
		var collectionTitle string
		var movie models.MovieShortInfo

		err := rows.Scan(&collectionID, &collectionTitle,
			&movie.ID, &movie.Title, &movie.CardURL, &movie.AlbumURL, &movie.Rating, &movie.ReleaseDate, &movie.MovieType, &movie.Country)

		if err != nil {
			errMsg := fmt.Errorf("error while scanning movie_service collections: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())

			return nil, errMsg
		}

		if _, exists := collections[collectionID]; !exists {
			collections[collectionID] = models.Collection{ID: collectionID, Title: collectionTitle, Movies: []*models.MovieShortInfo{}}
		}

		collection := collections[collectionID]
		collection.Movies = append(collection.Movies, &movie)

		collections[collectionID] = collection
	}

	return collections, nil
}

// ScanMovieShortInfo scans movie short info rows
func ScanMovieShortInfo(rows *sql.Rows) ([]models.MovieShortInfo, error) {
	defer closeRows(rows, "movie_short_info")
	var movies []models.MovieShortInfo

	for rows.Next() {
		var movie models.MovieShortInfo

		err := rows.Scan(&movie.ID, &movie.Title, &movie.CardURL, &movie.AlbumURL, &movie.Rating, &movie.ReleaseDate, &movie.MovieType, &movie.Country)

		if err != nil {
			errMsg := fmt.Errorf("error while scanning favorite movies: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())

			return nil, errMsg
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func closeRows(rows *sql.Rows, entity string) {
	if err := rows.Close(); err != nil {
		errMsg := fmt.Errorf("cannot close rows while taking %s: %w", entity, err)
		log.Error().Err(errMsg).Msg(errMsg.Error())
	}
}
