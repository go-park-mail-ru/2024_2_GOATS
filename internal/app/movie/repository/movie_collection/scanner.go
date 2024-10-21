package movie_collection

import (
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

func ScanConnections(rows *sql.Rows) (map[int]models.Collection, error) {
	collections := map[int]models.Collection{}

	for rows.Next() {
		var collectionID int
		var collectionTitle string
		var movie models.MovieInfo

		err := rows.Scan(&collectionID, &collectionTitle,
			&movie.Id, &movie.Title, &movie.CardUrl, &movie.AlbumUrl, &movie.Rating, &movie.MovieType, &movie.ReleaseDate, &movie.Country)

		if err != nil {
			errMsg := fmt.Errorf("error while scanning movie collections: %w", err)
			log.Err(errMsg)

			return nil, errMsg
		}

		if _, exists := collections[collectionID]; !exists {
			collections[collectionID] = models.Collection{Id: collectionID, Title: collectionTitle, Movies: []*models.MovieInfo{}}
		}

		tempCollection := collections[collectionID]
		tempCollection.Movies = append(tempCollection.Movies, &movie)

		collections[collectionID] = tempCollection
	}

	return collections, nil
}
