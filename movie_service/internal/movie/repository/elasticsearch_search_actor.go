package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	"io"
	"log"
	"strconv"
)

func (r *MovieRepo) SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error) {
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
	defer res.Body.Close()

	bodyBytes, _ := io.ReadAll(res.Body)
	log.Println("ElasticSearch Response:", string(bodyBytes))

	if res.IsError() {
		return nil, fmt.Errorf("search error: %s", res.String())
	}

	var esResponse struct {
		Hits struct {
			Hits []struct {
				Source struct {
					ID          string `json:"id"`
					Name        string `json:"full_name"`
					PhotoBigUrl string `json:"photo_big_url"`
				} `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&esResponse); err != nil {
		return nil, fmt.Errorf("error decoding search response: %w", err)
	}

	log.Println("Hits:", esResponse.Hits.Hits)

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
			BigPhotoURL: hit.Source.PhotoBigUrl,
			Person: models.Person{
				Name: hit.Source.Name,
			},
		}
	}

	return actors, nil
}
