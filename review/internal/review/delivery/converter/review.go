package converter

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/review/internal/review/service/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/review/models"
	review "github.com/go-park-mail-ru/2024_2_GOATS/review/pkg/review_v1"
)

// Convert gRPC Data slice to Service DTO slice
func ToSrvDataSlice(grpcData []*review.Data) []*dto.DataDTO {
	if grpcData == nil {
		return nil
	}

	srvDataSlice := make([]*dto.DataDTO, 0, len(grpcData))
	for _, d := range grpcData {
		if d == nil {
			continue
		}
		srvDataSlice = append(srvDataSlice, &dto.DataDTO{
			QuestionID: d.QuestionId,
			AnswerID:   d.AnswerId,
			Answer:     d.Answer,
			Rating:     int64(d.Rating),
		})
	}
	return srvDataSlice
}

// Convert Service DTO slice to gRPC Data slice
func ToGRPCDataSlice(srvData []*dto.DataDTO, rating float64) []*review.Data {
	if srvData == nil {
		return nil
	}

	grpcDataSlice := make([]*review.Data, 0, len(srvData))
	for _, d := range srvData {
		if d == nil {
			continue
		}
		grpcDataSlice = append(grpcDataSlice, &review.Data{
			QuestionId: d.QuestionID,
			AnswerId:   d.AnswerID,
			Answer:     d.Answer,
			Rating:     float32(rating),
		})
	}
	return grpcDataSlice
}

// Convert Service DTO slice to gRPC Data slice
func ToGRPCQuestionsSlice(srvData []*dto.QuestionDTO) []*review.Question {
	if srvData == nil {
		return nil
	}

	grpcQuestionsSlice := make([]*review.Question, 0, len(srvData))
	for _, d := range srvData {
		if d == nil {
			continue
		}

		var ans []*review.Answer
		for _, a := range d.Answers {
			curr := &review.Answer{
				AnswersId: int64(a.ID),
				Answers:   a.Content,
			}

			ans = append(ans, curr)
		}

		grpcQuestionsSlice = append(grpcQuestionsSlice, &review.Question{
			IsActive:   d.IsActive,
			QuestionId: d.QuestionID,
			Question:   d.Question,
			Answers:    ans,
		})
	}
	return grpcQuestionsSlice
}

// From Model Data to Service DTO
func ToSrvData(modelData *models.Data) *dto.DataDTO {
	if modelData == nil {
		return nil
	}

	return &dto.DataDTO{
		QuestionID: modelData.QuestionID,
		AnswerID:   modelData.AnswerID,
		Answer:     modelData.Answer,
	}
}

// From Service DTO to Model Data
func ToModelData(srvData *dto.DataDTO) *models.Data {
	if srvData == nil {
		return nil
	}

	return &models.Data{
		QuestionID: srvData.QuestionID,
		AnswerID:   srvData.AnswerID,
		Answer:     srvData.Answer,
	}
}

// From Model Question to Service DTO
func ToSrvQuestion(modelQuestion *models.Question) *dto.QuestionDTO {
	if modelQuestion == nil {
		return nil
	}

	var ans []dto.AnswerDTO
	for _, a := range modelQuestion.Answers {
		cur := dto.AnswerDTO{
			ID:      int(a.ID),
			Content: a.Content,
		}

		ans = append(ans, cur)
	}

	return &dto.QuestionDTO{
		IsActive:   modelQuestion.IsActive,
		QuestionID: modelQuestion.QuestionID,
		Question:   modelQuestion.Question,
		Answers:    ans,
	}
}

// From Service DTO to Model Question
func ToModelQuestion(srvQuestion *dto.QuestionDTO) *models.Question {
	if srvQuestion == nil {
		return nil
	}

	var ans []models.Answer
	for _, a := range srvQuestion.Answers {
		cur := models.Answer{
			ID:      int64(a.ID),
			Content: a.Content,
		}

		ans = append(ans, cur)
	}

	return &models.Question{
		IsActive:   srvQuestion.IsActive,
		QuestionID: srvQuestion.QuestionID,
		Question:   srvQuestion.Question,
		Answers:    ans,
	}
}
