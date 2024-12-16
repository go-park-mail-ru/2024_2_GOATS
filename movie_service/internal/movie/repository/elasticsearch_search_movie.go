package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	zl "github.com/rs/zerolog/log"
)

// SearchMovies search movies via Elasticsearch
func (r *MovieRepo) SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error) {
	logger := zl.Ctx(ctx)
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match_phrase_prefix": map[string]interface{}{
				"title": query,
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		return nil, fmt.Errorf("error encoding search query: %w", err)
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

	var esResponse struct {
		Hits struct {
			Hits []struct {
				Source struct {
					ID       string  `json:"id"`
					Title    string  `json:"title"`
					Rating   float32 `json:"rating"`
					AlbumURL string  `json:"album_url"`
					CardURL  string  `json:"card_url"`
				} `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if decErr := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&esResponse); decErr != nil {
		return nil, fmt.Errorf("error decoding search response: %w", decErr)
	}

	if len(esResponse.Hits.Hits) == 0 {
		return []models.MovieInfo{}, nil
	}

	movies := make([]models.MovieInfo, len(esResponse.Hits.Hits))
	for i, hit := range esResponse.Hits.Hits {
		id := hit.Source.ID
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
			Title:    hit.Source.Title,
			Rating:   hit.Source.Rating,
			CardURL:  hit.Source.CardURL,
			AlbumURL: hit.Source.AlbumURL,
			VerURL:   hit.Source.VerURL,
		}
	}

	return movies, nil
}
