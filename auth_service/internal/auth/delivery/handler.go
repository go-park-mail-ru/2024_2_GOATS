package delivery

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/delivery/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/errs"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/validation"
	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
	"github.com/rs/zerolog/log"
)

// AuthManager is a auth grpc handler
type AuthManager struct {
	auth.UnimplementedSessionRPCServer
	SessionSrv AuthServiceInterface
	redisCfg   *config.Redis
}

// NewAuthManager returns an instance of SessionRPCServer
func NewAuthManager(ctx context.Context, srv AuthServiceInterface) auth.SessionRPCServer {
	return &AuthManager{
		SessionSrv: srv,
		redisCfg:   config.FromRedisContext(ctx),
	}
}

// CreateSession create_session grpc handler
func (am *AuthManager) CreateSession(ctx context.Context, createCookieReq *auth.CreateSessionRequest) (*auth.CreateSessionResponse, error) {
	logger := log.Ctx(ctx)
	err := validation.ValidateUserID(createCookieReq.UserID)
	if err != nil {
		err = fmt.Errorf("create_session: %w", err)
		logger.Error().Err(err).Msg("failed_to_create_session")

		return nil, err
	}

	ctx = config.WrapRedisContext(ctx, am.redisCfg)
	srvData := converter.ToSrvCreateCookieFromDesc(createCookieReq)
	if srvData == nil {
		return nil, errs.ErrBadRequest
	}

	srvResp, err := am.SessionSrv.CreateSession(ctx, srvData)
	resp := converter.ToDescCreateCookieRespFromSrv(srvResp)
	if err != nil {
		logger.Error().Interface("createSessionError", err).Msg("failed_to_create_session")
		return nil, err
	}

	logger.Info().Interface("createSessionSuccess", resp).Msg("successfully_create_session")
	return resp, nil
}

// DestroySession destroy_session grpc handler
func (am *AuthManager) DestroySession(ctx context.Context, deleteCookieReq *auth.DestroySessionRequest) (*auth.Nothing, error) {
	logger := log.Ctx(ctx)
	err := validation.ValidateCookie(deleteCookieReq.Cookie)
	if err != nil {
		err = fmt.Errorf("destroy_session: %w", err)
		logger.Error().Err(err).Msg("failed_to_destroy_session")

		return nil, err
	}

	srvResp, err := am.SessionSrv.DestroySession(ctx, deleteCookieReq.Cookie)
	if err != nil {
		logger.Error().Interface("destroySessionError", err).Msg("failed_to_destroy_session")
		return nil, err
	}

	logger.Info().Msg("successfully_destroy_session")
	return &auth.Nothing{Dummy: srvResp}, nil
}

// Session get_session_data grpc handler
func (am *AuthManager) Session(ctx context.Context, checkCookieReq *auth.GetSessionRequest) (*auth.GetSessionResponse, error) {
	logger := log.Ctx(ctx)
	err := validation.ValidateCookie(checkCookieReq.Cookie)
	if err != nil {
		err = fmt.Errorf("check_session: %w", err)
		logger.Error().Err(err).Msg("failed_to_get_session_data")

		return nil, err
	}

	srvResp, err := am.SessionSrv.GetSessionData(ctx, checkCookieReq.Cookie)
	if err != nil {
		logger.Error().Interface("getSessionDataError", err).Msg("failed_to_get_session_data")
		return nil, err
	}

	logger.Info().Interface("getSessionDataSuccess", srvResp).Msg("successfully_get_session_data")
	return &auth.GetSessionResponse{UserID: srvResp}, nil
}
