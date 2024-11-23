package repository

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_GOATS/review/internal/review/service/dto"
)

type ReviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

// SaveSurveyData сохраняет данные опроса
func (r *ReviewRepository) SaveSurveyData(ctx context.Context, userID int64, data []*dto.DataDTO) error {
	query := `
  INSERT INTO "survey_data" ("user_id", "question_id", "answer_id", "answer")
  VALUES ($1, $2, $3, $4)
 `

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, item := range data {
		_, err := tx.ExecContext(ctx, query, userID, item.QuestionID, item.AnswerID, item.Answer)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// FetchSurveyStatistics возвращает статистику опросов
func (r *ReviewRepository) FetchSurveyStatistics(ctx context.Context) ([]*dto.DataDTO, error) {
	query := `
  SELECT "question_id", "answer_id", "answer"
  FROM "survey_data"
 `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*dto.DataDTO
	for rows.Next() {
		var data dto.DataDTO
		if err := rows.Scan(&data.QuestionID, &data.AnswerID, &data.Answer); err != nil {
			return nil, err
		}
		result = append(result, &data)
	}

	return result, nil
}

// HasUserPassedSurvey проверяет, прошел ли пользователь опрос
func (r *ReviewRepository) HasUserPassedSurvey(ctx context.Context, userID int64) (bool, error) {
	query := `
  SELECT COUNT(*)
  FROM "survey_data"
  WHERE "user_id" = $1
 `

	var count int
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// FetchActiveQuestions возвращает активные вопросы
func (r *ReviewRepository) FetchActiveQuestions(ctx context.Context) ([]*dto.QuestionDTO, error) {
	query := `
  SELECT "question_id", "question", "is_active"
  FROM "questions"
  WHERE "is_active" = TRUE
 `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*dto.QuestionDTO
	for rows.Next() {
		var question dto.QuestionDTO
		if err := rows.Scan(&question.QuestionID, &question.Question, &question.IsActive); err != nil {
			return nil, err
		}
		result = append(result, &question)
	}

	return result, nil
}
