package repository

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (r *Repo) GetCollection(ctx context.Context) ([]models.Collection, *errVals.ErrorObj, int) {
	sqlStatement := `
		SELECT collections.id, collections.title, movies.id, movies.title, movies.card_url, movies.album_url, movies.rating, movies.release_date, countries.title FROM collections
		JOIN movie_collections ON movie_collections.collection_id = collections.id
		JOIN movies ON movies.id = movie_collections.movie_id
		JOIN countries ON countries.id = movies.country_id
	`

	rows, err := r.Database.QueryContext(ctx, sqlStatement)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	collections := map[int]models.Collection{}

	for rows.Next() {
		var collectionID int
		var collectionTitle string
		var movie models.Movie

		err := rows.Scan(&collectionID, &collectionTitle,
			&movie.Id, &movie.Title, &movie.CardUrl, &movie.AlbumUrl, &movie.Rating, &movie.ReleaseDate, &movie.Country)

		if err != nil {
			return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
		}

		if _, exists := collections[collectionID]; !exists {
			collections[collectionID] = models.Collection{Id: collectionID, Title: collectionTitle, Movies: []*models.Movie{}}
		}

		tempCollection := collections[collectionID]
		tempCollection.Movies = append(tempCollection.Movies, &movie)

		collections[collectionID] = tempCollection
	}

	result := make([]models.Collection, 0, len(collections))
	for _, collection := range collections {
		result = append(result, collection)
	}

	return result, nil, http.StatusOK
}
