package delivery

import (
	"context"
	"errors"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/delivery/converter"
	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
	"github.com/rs/zerolog/log"
)

// "google.golang.org/grpc/codes"
//     "google.golang.org/grpc/status"

type AuthManager struct {
	auth.UnimplementedSessionRPCServer
	SessionSrv AuthServiceInterface
	redisCfg   *config.Redis
}

func NewAuthManager(ctx context.Context, srv AuthServiceInterface) auth.SessionRPCServer {
	return &AuthManager{
		SessionSrv: srv,
		redisCfg:   config.FromRedisContext(ctx),
	}
}

func (am *AuthManager) CreateSession(ctx context.Context, createCookieReq *auth.CreateSessionRequest) (*auth.CreateSessionResponse, error) {
	logger := log.Ctx(ctx)
	ctx = config.WrapRedisContext(ctx, am.redisCfg)
	srvData := converter.ToSrvCreateCookieFromDesc(createCookieReq)
	if srvData == nil {
		return nil, errors.New("bad_request")
	}

	srvResp, srvErr := am.SessionSrv.CreateSession(ctx, srvData)
	resp := converter.ToDescCreateCookieRespFromSrv(srvResp)
	if srvErr != nil {
		logger.Error().Interface("createSessionError", srvErr).Msg("failed_to_create_session")
		return nil, srvErr.Error.Err
	}

	logger.Info().Interface("createSessionSuccess", resp).Msg("successfully_create_session")
	return resp, nil
}

func (am *AuthManager) DestroySession(ctx context.Context, deleteCookieReq *auth.DestroySessionRequest) (*auth.Nothing, error) {
	logger := log.Ctx(ctx)
	srvResp, srvErr := am.SessionSrv.DestroySession(ctx, deleteCookieReq.Cookie)
	if srvErr != nil {
		logger.Error().Interface("destroySessionError", srvErr).Msg("failed_to_destroy_session")
		return nil, srvErr.Error.Err
	}

	logger.Info().Msg("successfully_destroy_session")
	return &auth.Nothing{Dummy: srvResp}, nil
}

func (am *AuthManager) Session(ctx context.Context, checkCookieReq *auth.GetSessionRequest) (*auth.GetSessionResponse, error) {
	logger := log.Ctx(ctx)
	srvResp, srvErr := am.SessionSrv.GetSessionData(ctx, checkCookieReq.Cookie)
	if srvErr != nil {
		logger.Error().Interface("getSessionDataError", srvErr).Msg("failed_to_get_session_data")
		return nil, srvErr.Error.Err
	}

	logger.Info().Interface("getSessionDataSuccess", srvResp).Msg("successfully_get_session_data")
	return &auth.GetSessionResponse{UserID: srvResp}, nil
}
