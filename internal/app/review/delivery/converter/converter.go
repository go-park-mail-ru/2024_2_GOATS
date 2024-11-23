package converter

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/review/service/dto"
)

func ConvertCheckReviewToAPI(srvData *dto.CheckReviewData) *api.CheckReviewData {
	if srvData == nil {
		return nil
	}

	return &api.CheckReviewData{
		CSAT: srvData.CSAT,
		NPS:  srvData.NPS,
		CSI:  srvData.CSI,
	}
}

func ConvertReviewDataToAPI(data []*dto.ReviewData) *api.GetReviewResponse {
	if len(data) == 0 {
		return nil
	}

	var resp api.GetReviewResponse

	for _, rd := range data {
		var answers []api.Answer
		for _, a := range rd.Answers {
			currAns := api.Answer{
				ID:      a.ID,
				Content: a.Content,
			}

			answers = append(answers, currAns)
		}

		review := &api.ReviewData{
			ID:      rd.ID,
			Title:   rd.Title,
			Answers: answers,
		}

		resp.Questions = append(resp.Questions, review)
	}

	return &resp
}

func ConvertCreateReqToSrv(req *api.CreateReviewRequest) []*dto.CreateReviewData {
	if req == nil {
		return nil
	}

	var res []*dto.CreateReviewData

	for _, q := range req.Questions {
		cur := &dto.CreateReviewData{
			ID:         q.ID,
			AnswerID:   q.AnswerID,
			AnswerText: q.AnswerText,
		}

		res = append(res, cur)
	}

	return res
}

func ConvertStatisticToAPI(srvData *dto.Statistic) *api.Statistic {
	if srvData == nil {
		return nil
	}

	return &api.Statistic{
		Rating:   srvData.Rating,
		Type:     srvData.Type,
		Comments: srvData.Comments,
	}
}
