package repository

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	zl "github.com/rs/zerolog/log"
)

// SearchMovies search movies via Elasticsearch
func (r *MovieRepo) SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error) {
	logger := zl.Ctx(ctx)
	searchQuery := models.SearchMovieQuery{
		MovieQuery: models.MovieQuery{
			MatchMoviePhrasePrefix: models.MatchMoviePhrasePrefix{
				Title: query,
			},
		},
	}

	data, err := searchQuery.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("error encoding search query: %w", err)
	}

	var buf bytes.Buffer
	if _, err := buf.Write(data); err != nil {
		return nil, fmt.Errorf("error writing to buffer: %w", err)
	}

	res, err := r.Elasticsearch.Search(
		r.Elasticsearch.Search.WithContext(ctx),
		r.Elasticsearch.Search.WithIndex("movies"),
		r.Elasticsearch.Search.WithBody(&buf),
		r.Elasticsearch.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, fmt.Errorf("error executing search query: %w", err)
	}

	defer func() {
		if clErr := res.Body.Close(); clErr != nil {
			logger.Error().Err(clErr).Msg("cannot close searchMovies body")
		}
	}()

	bodyBytes, _ := io.ReadAll(res.Body)

	if res.IsError() {
		return nil, fmt.Errorf("search error: %s", res.String())
	}

	var esResponse models.MovieESResponse

	if err := esResponse.UnmarshalJSON(bodyBytes); err != nil {
		return nil, fmt.Errorf("error decoding search response: %w", err)
	}

	if len(esResponse.MovieHits.MovieHits) == 0 {
		return []models.MovieInfo{}, nil
	}

	movies := make([]models.MovieInfo, len(esResponse.MovieHits.MovieHits))
	for i, hit := range esResponse.MovieHits.MovieHits {
		id := hit.MovieSource.ID
		if err != nil {
			return nil, fmt.Errorf("error converting id to int: %w", err)
		}
		if err != nil {
			return nil, fmt.Errorf("error converting rating to float: %w", err)
		}
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return nil, fmt.Errorf("error converting rating to int: %w", err)
		}
		movies[i] = models.MovieInfo{
			ID:       idInt,
			Title:    hit.MovieSource.Title,
			Rating:   hit.MovieSource.Rating,
			CardURL:  hit.MovieSource.CardURL,
			AlbumURL: hit.MovieSource.AlbumURL,
		}
	}

	return movies, nil
}
