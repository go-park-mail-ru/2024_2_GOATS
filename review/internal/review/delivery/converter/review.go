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
		})
	}
	return srvDataSlice
}

// Convert Service DTO slice to gRPC Data slice
func ToGRPCDataSlice(srvData []*dto.DataDTO) []*review.Data {
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
		grpcQuestionsSlice = append(grpcQuestionsSlice, &review.Question{
			IsActive:   d.IsActive,
			QuestionId: d.QuestionID,
			Question:   d.Question,
			Answer:     d.Answers,
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

	return &dto.QuestionDTO{
		IsActive:   modelQuestion.IsActive,
		QuestionID: modelQuestion.QuestionID,
		Question:   modelQuestion.Question,
		Answers:    modelQuestion.Answers,
	}
}

// From Service DTO to Model Question
func ToModelQuestion(srvQuestion *dto.QuestionDTO) *models.Question {
	if srvQuestion == nil {
		return nil
	}

	return &models.Question{
		IsActive:   srvQuestion.IsActive,
		QuestionID: srvQuestion.QuestionID,
		Question:   srvQuestion.Question,
		Answers:    srvQuestion.Answers,
	}
}
