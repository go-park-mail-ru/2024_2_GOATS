package movie_collectiondb

import (
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

func ScanConnections(rows *sql.Rows) (map[int]models.Collection, error) {
	collections := make(map[int]models.Collection, 3)

	defer func() {
		if err := rows.Close(); err != nil {
			errMsg := fmt.Errorf("cannot close rows while taking movie_collections: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())
		}
	}()

	for rows.Next() {
		var collectionID int
		var collectionTitle string
		var movie models.MovieShortInfo

		err := rows.Scan(&collectionID, &collectionTitle,
			&movie.ID, &movie.Title, &movie.CardUrl, &movie.AlbumUrl, &movie.Rating, &movie.ReleaseDate, &movie.MovieType, &movie.Country)

		if err != nil {
			errMsg := fmt.Errorf("error while scanning movie collections: %w", err)
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
