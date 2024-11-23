package delivery

import (
	"context"
	dto "github.com/go-park-mail-ru/2024_2_GOATS/review/internal/review/service/dto"
)

//go:generate mockgen -source=delivery.go -destination=mocks/mock.go
type ReviewServiceInterface interface {
	Create(ctx context.Context, userID int64, data []*dto.DataDTO) error
	GetQuestionData(ctx context.Context) ([]*dto.DataDTO, float64, error)
	CheckPass(ctx context.Context, userID int64) (bool, error)
	CreateFront(ctx context.Context) ([]*dto.QuestionDTO, error)
}
