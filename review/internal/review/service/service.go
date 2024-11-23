package service

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/review/internal/review/service/dto"
)

type ReviewRepositoryInterface interface {
	SaveSurveyData(ctx context.Context, userID int64, data []*dto.DataDTO) error
	FetchSurveyStatistics(ctx context.Context) ([]*dto.DataDTO, error)
	HasUserPassedSurvey(ctx context.Context, userID int64) (bool, error)
	FetchActiveQuestions(ctx context.Context) ([]*dto.QuestionDTO, error)
}

type ReviewService struct {
	repo ReviewRepositoryInterface
}

func NewReviewService(repo ReviewRepositoryInterface) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) Create(ctx context.Context, userID int64, data []*dto.DataDTO) error {
	if userID == 0 || len(data) == 0 {
		return errors.New("invalid_input")
	}

	return s.repo.SaveSurveyData(ctx, userID, data)
}

func (s *ReviewService) GetQuestionData(ctx context.Context) ([]*dto.DataDTO, error) {
	return s.repo.FetchSurveyStatistics(ctx)
}

func (s *ReviewService) CheckPass(ctx context.Context, userID int64) (bool, error) {
	if userID == 0 {
		return false, errors.New("invalid_user_id")
	}

	return s.repo.HasUserPassedSurvey(ctx, userID)
}

func (s *ReviewService) CreateFront(ctx context.Context) ([]*dto.QuestionDTO, error) {
	return s.repo.FetchActiveQuestions(ctx)
}
