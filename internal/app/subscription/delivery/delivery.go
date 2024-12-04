package delivery

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

type SubscriptionServiceInterface interface {
	Subscribe(ctx context.Context, data *models.SubscriptionData) (string, *errVals.ServiceError)
}
