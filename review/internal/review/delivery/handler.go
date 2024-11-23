package delivery

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/review/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/review/internal/review/delivery/converter"
	review "github.com/go-park-mail-ru/2024_2_GOATS/review/pkg/review_v1"
	"github.com/rs/zerolog/log"
)

type ReviewHandler struct {
	review.UnimplementedReviewServer
	ReviewService ReviewServiceInterface
	redisCfg      *config.Redis
}

func NewReviewHandler(ctx context.Context, srv ReviewServiceInterface) review.ReviewServer {
	return &ReviewHandler{
		ReviewService: srv,
		redisCfg:      config.FromRedisContext(ctx),
	}
}

func (rm *ReviewHandler) Create(ctx context.Context, createReq *review.CreateRequest) (*review.CreateResponse, error) {
	logger := log.Ctx(ctx)
	ctx = config.WrapRedisContext(ctx, rm.redisCfg)
	srvData := converter.ToSrvDataSlice(createReq.Data)
	if srvData == nil {
		return nil, errors.New("bad_request")
	}
	srvUserData := createReq.UserId

	err := rm.ReviewService.Create(ctx, srvUserData, srvData)
	if err != nil {
		logger.Error().Interface("createSessionError", err).Msg("failed_to_create_session")
		return nil, err
	}

	return &review.CreateResponse{Status: "1"}, nil
}

func (rm *ReviewHandler) GetQuestionData(ctx context.Context, getQuestionDataReq *review.GetQuestionDataRequest) (*review.GetQuestionDataResponse, error) {
	logger := log.Ctx(ctx)
	srvResp, rating, err := rm.ReviewService.GetQuestionData(ctx)
	dataSrvResp := converter.ToGRPCDataSlice(srvResp, rating)

	if err != nil {
		logger.Error().Interface("destroySessionError", err).Msg("failed_to_destroy_session")
		return nil, err
	}

	logger.Info().Msg("successfully_destroy_session")
	return &review.GetQuestionDataResponse{Data: dataSrvResp}, nil
}

func (rm *ReviewHandler) CheckPass(ctx context.Context, checkPassRequest *review.CheckPassRequest) (*review.CheckPassResponse, error) {
	logger := log.Ctx(ctx)
	srvUserData := checkPassRequest.UserId
	srvResp, err := rm.ReviewService.CheckPass(ctx, srvUserData)
	if err != nil {
		logger.Error().Interface("getSessionDataError", err).Msg("failed_to_get_session_data")
		return nil, err
	}

	logger.Info().Interface("getSessionDataSuccess", srvResp).Msg("successfully_get_session_data")
	return &review.CheckPassResponse{Passed: srvResp}, nil
}

func (rm *ReviewHandler) CreateFront(ctx context.Context, createFrontReq *review.CreateFrontRequest) (*review.CreateFrontResponse, error) {
	logger := log.Ctx(ctx)
	srvResp, err := rm.ReviewService.CreateFront(ctx)
	dataSrvResp := converter.ToGRPCQuestionsSlice(srvResp)

	if err != nil {
		logger.Error().Interface("getSessionDataError", err).Msg("failed_to_get_session_data")
		return nil, err
	}

	logger.Info().Interface("getSessionDataSuccess", srvResp).Msg("successfully_get_session_data")
	return &review.CreateFrontResponse{Questions: dataSrvResp}, nil
}
