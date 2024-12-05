package delivery

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

type PaymentServiceInterface interface {
	ProcessCallback(ctx context.Context, data *models.PaymentCallbackData) *errVals.ServiceError
}