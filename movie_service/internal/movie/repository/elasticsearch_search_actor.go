package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	zl "github.com/rs/zerolog/log"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
)

// SearchActors search_actors in elasticsearch
func (r *MovieRepo) SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error) {
	logger := zl.Ctx(ctx)
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match_phrase_prefix": map[string]interface{}{
				"full_name": query,
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		return nil, fmt.Errorf("error encoding search query: %w", err)
	}

	res, err := r.Elasticsearch.Search(
		r.Elasticsearch.Search.WithContext(ctx),
		r.Elasticsearch.Search.WithIndex("actors"),
		r.Elasticsearch.Search.WithBody(&buf),
		r.Elasticsearch.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, fmt.Errorf("error executing search query: %w", err)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			logger.Error().Err(err).Msg("cannot close searchActors body")
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
					ID          string `json:"id"`
					Name        string `json:"full_name"`
					PhotoBigURL string `json:"photo_big_url"`
				} `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&esResponse); err != nil {
		return nil, fmt.Errorf("error decoding search response: %w", err)
	}

	if len(esResponse.Hits.Hits) == 0 {
		return []models.ActorInfo{}, nil
	}

	actors := make([]models.ActorInfo, len(esResponse.Hits.Hits))
	for i, hit := range esResponse.Hits.Hits {
		id, err := strconv.Atoi(hit.Source.ID)
		if err != nil {
			return nil, fmt.Errorf("error converting id to int: %w", err)
		}
		actors[i] = models.ActorInfo{
			ID:          id,
			BigPhotoURL: hit.Source.PhotoBigURL,
			Person: models.Person{
				Name: hit.Source.Name,
			},
		}
	}

	return actors, nil
}
