package client

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/review/service/dto"
	review "github.com/go-park-mail-ru/2024_2_GOATS/review/pkg/review_v1"
)

type ReviewClientInterface interface {
	Create(ctx context.Context, usrID int, data []*dto.CreateReviewData) error
	CheckReview(ctx context.Context, usrID int) (bool, error)
	GetStatistics(ctx context.Context) (*dto.Statistic, error)
	GetQuestions(ctx context.Context) ([]*dto.ReviewData, error)
}

type ReviewClient struct {
	reviewMS review.ReviewClient
}

func NewReviewClient(reviewMS review.ReviewClient) ReviewClientInterface {
	return &ReviewClient{
		reviewMS: reviewMS,
	}
}

func (rc *ReviewClient) Create(ctx context.Context, usrID int, data []*dto.CreateReviewData) error {
	var grpcData []*review.Data
	for _, rd := range data {
		curr := review.Data{
			QuestionId: int64(rd.ID),
			AnswerId:   int64(rd.AnswerID),
			Answer:     rd.AnswerText,
		}

		grpcData = append(grpcData, &curr)
	}

	grpcReq := review.CreateRequest{
		UserId: int64(usrID),
		Data:   grpcData,
	}

	_, err := rc.reviewMS.Create(ctx, &grpcReq)
	if err != nil {
		return fmt.Errorf("reviewClient: %w", err)
	}

	return nil
}

func (rc *ReviewClient) CheckReview(ctx context.Context, usrID int) (bool, error) {
	set, err := rc.reviewMS.CheckPass(ctx, &review.CheckPassRequest{UserId: int64(usrID)})
	if err != nil {
		return false, fmt.Errorf("cannot check review existence: %w", err)
	}

	return set.Passed, nil
}

func (rc *ReviewClient) GetQuestions(ctx context.Context) ([]*dto.ReviewData, error) {
	resp, err := rc.reviewMS.CreateFront(ctx, &review.CreateFrontRequest{})
	if err != nil {
		return nil, fmt.Errorf("cannot check review existence: %w", err)
	}

	var rd []*dto.ReviewData

	for _, q := range resp.Questions {
		var ans []dto.Answer
		for _, a := range q.Answers {
			curr := dto.Answer{
				ID:      int(a.AnswersId),
				Content: a.Answers,
			}

			ans = append(ans, curr)
		}

		curr := &dto.ReviewData{
			ID:      int(q.QuestionId),
			Title:   q.Question,
			Answers: ans,
			// Type:    q.Type,
		}

		rd = append(rd, curr)
	}

	return rd, nil
}

func (rc *ReviewClient) GetStatistics(ctx context.Context) (*dto.Statistic, error) {
	resp, err := rc.reviewMS.GetQuestionData(ctx, &review.GetQuestionDataRequest{})
	if err != nil {
		return nil, fmt.Errorf("cannot get statistics: %w", err)
	}

	var comments []string
	var rating float32
	for _, d := range resp.Data {
		comments = append(comments, d.Answer)
		rating = d.Rating
	}

	return &dto.Statistic{
		Rating:   float64(rating),
		Type:     "csat",
		Comments: comments,
	}, nil
}
