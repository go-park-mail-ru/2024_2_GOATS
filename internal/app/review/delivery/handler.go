package delivery

import (
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/review/delivery/converter"
	"github.com/rs/zerolog/log"
)

type ReviewHandler struct {
	reviewSrv ReviewServiceInterface
}

func NewReviewHandler(reviewSRV ReviewServiceInterface) handlers.ReviewHandlerInterface {
	return &ReviewHandler{
		reviewSrv: reviewSRV,
	}
}

func (h *ReviewHandler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	srvData, srvErr := h.reviewSrv.GetQuestions(r.Context())

	if srvErr != nil {
		err := errors.ToDeliveryErrorFromService(srvErr)
		errMsg := fmt.Errorf("failed to get csat questions: %w", srvErr.Error.Err)
		logger.Error().Err(errMsg).Interface("getCSATErr", err).Msg("request_failed")
		api.Response(r.Context(), w, err.HTTPStatus, err)
	}

	apiData := converter.ConvertReviewDataToAPI(srvData)
	logger.Info().Interface("getCSATSuccess", apiData).Msg("get csta success")
	api.Response(r.Context(), w, http.StatusOK, apiData)
}

func (h *ReviewHandler) CreateCSAT(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	req := &api.CreateReviewRequest{}
	api.DecodeBody(w, r, req)

	fmt.Println(req)
	srvErr := h.reviewSrv.CreateCSAT(r.Context(), converter.ConvertCreateReqToSrv(req))

	if srvErr != nil {
		err := errors.ToDeliveryErrorFromService(srvErr)
		errMsg := fmt.Errorf("failed to check review's statuses: %w", srvErr.Error.Err)
		logger.Error().Err(errMsg).Interface("checkReviewErr", err).Msg("request_failed")
		api.Response(r.Context(), w, err.HTTPStatus, err)
	}

	logger.Info().Msg("createCSAT success")
	api.Response(r.Context(), w, http.StatusOK, nil)
}

func (h *ReviewHandler) CheckReview(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	usrID := config.CurrentUserID(r.Context())
	if usrID == 0 {
		logger.Error().Msg("unathorized")
		api.Response(r.Context(), w, http.StatusForbidden, "unathorized")
	}

	srvData, srvErr := h.reviewSrv.CheckReview(r.Context(), usrID)

	if srvErr != nil {
		err := errors.ToDeliveryErrorFromService(srvErr)
		errMsg := fmt.Errorf("failed to check review's statuses: %w", srvErr.Error.Err)
		logger.Error().Err(errMsg).Interface("checkReviewErr", err).Msg("request_failed")
		api.Response(r.Context(), w, err.HTTPStatus, err)
	}

	apiData := converter.ConvertCheckReviewToAPI(srvData)
	logger.Info().Interface("checkReviewSuccess", apiData).Msg("checkReview success")
	api.Response(r.Context(), w, http.StatusOK, apiData)
}

func (h *ReviewHandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	srvData, srvErr := h.reviewSrv.GetStatistics(r.Context())

	if srvErr != nil {
		err := errors.ToDeliveryErrorFromService(srvErr)
		errMsg := fmt.Errorf("failed to get statistic: %w", srvErr.Error.Err)
		logger.Error().Err(errMsg).Interface("getStatisticErr", err).Msg("request_failed")
		api.Response(r.Context(), w, err.HTTPStatus, err)
	}

	apiData := converter.ConvertStatisticToAPI(srvData)
	logger.Info().Interface("getStatisticSuccess", apiData).Msg("getStatistic success")
	api.Response(r.Context(), w, http.StatusOK, apiData)
}
