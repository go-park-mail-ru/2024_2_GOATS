package repository

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strconv"

	zl "github.com/rs/zerolog/log"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
)

// SearchActors search_actors in elasticsearch
func (r *MovieRepo) SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error) {
	logger := zl.Ctx(ctx)
	searchQuery := models.SearchActorQuery{
		ActorQuery: models.ActorQuery{
			MatchActorPhrasePrefix: models.MatchActorPhrasePrefix{
				FullName: query,
			},
		},
	}

	var buf bytes.Buffer
	data, err := searchQuery.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("error encoding search query: %w", err)
	}

	if _, err := buf.Write(data); err != nil {
		return nil, fmt.Errorf("error writing to buffer: %w", err)
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

	var esResponse models.ActorESResponse

	if err := esResponse.UnmarshalJSON(bodyBytes); err != nil {
		return nil, fmt.Errorf("error decoding search response: %w", err)
	}

	if len(esResponse.ActorHits.ActorHits) == 0 {
		return []models.ActorInfo{}, nil
	}

	actors := make([]models.ActorInfo, len(esResponse.ActorHits.ActorHits))
	for i, hit := range esResponse.ActorHits.ActorHits {
		id, err := strconv.Atoi(hit.ActorSource.ID)
		if err != nil {
			return nil, fmt.Errorf("error converting id to int: %w", err)
		}
		actors[i] = models.ActorInfo{
			ID:          id,
			BigPhotoURL: hit.ActorSource.PhotoBigURL,
			Person: models.Person{
				Name: hit.ActorSource.Name,
			},
		}
	}

	return actors, nil
}
