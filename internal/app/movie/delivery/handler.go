package delivery

import (
	"context"
	"net/url"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (i *Implementation) GetCollection(ctx context.Context, query url.Values) (*models.CollectionsResponse, *models.ErrorResponse) {
	colls, errData := i.movieService.GetCollection(ctx)
	if errData != nil {
		return nil, errData
	}

	return colls, nil
}
