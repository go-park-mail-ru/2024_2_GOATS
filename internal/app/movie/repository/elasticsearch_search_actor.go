package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (r *MovieRepo) SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error) {
	// Создаём запрос ElasticSearch
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query,
				"fields": []string{"title^2", "description"},
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

	if res.IsError() {
		return nil, fmt.Errorf("search error: %s", res.String())
	}

	// Обрабатываем ответ
	var esResponse struct {
		Hits struct {
			Hits []struct {
				Source models.ActorInfo `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&esResponse); err != nil {
		return nil, fmt.Errorf("error decoding search response: %w", err)
	}

	// Собираем результаты
	actors := make([]models.ActorInfo, len(esResponse.Hits.Hits))
	for i, hit := range esResponse.Hits.Hits {
		actors[i] = hit.Source
	}

	return actors, nil
}
