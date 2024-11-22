package delivery

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/delivery/converter"
	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
	"github.com/rs/zerolog/log"
)

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

	srvResp, err := am.SessionSrv.CreateSession(ctx, srvData)
	fmt.Println("SERVICE RESPONSE", srvResp)
	resp := converter.ToDescCreateCookieRespFromSrv(srvResp)
	fmt.Println("DEL RESPONSE", resp)
	if err != nil {
		logger.Error().Interface("createSessionError", err).Msg("failed_to_create_session")
		return nil, err
	}

	logger.Info().Interface("createSessionSuccess", resp).Msg("successfully_create_session")
	return resp, nil
}

func (am *AuthManager) DestroySession(ctx context.Context, deleteCookieReq *auth.DestroySessionRequest) (*auth.Nothing, error) {
	logger := log.Ctx(ctx)
	srvResp, err := am.SessionSrv.DestroySession(ctx, deleteCookieReq.Cookie)
	if err != nil {
		logger.Error().Interface("destroySessionError", err).Msg("failed_to_destroy_session")
		return nil, err
	}

	logger.Info().Msg("successfully_destroy_session")
	return &auth.Nothing{Dummy: srvResp}, nil
}

func (am *AuthManager) Session(ctx context.Context, checkCookieReq *auth.GetSessionRequest) (*auth.GetSessionResponse, error) {
	logger := log.Ctx(ctx)
	srvResp, err := am.SessionSrv.GetSessionData(ctx, checkCookieReq.Cookie)
	if err != nil {
		logger.Error().Interface("getSessionDataError", err).Msg("failed_to_get_session_data")
		return nil, err
	}

	logger.Info().Interface("getSessionDataSuccess", srvResp).Msg("successfully_get_session_data")
	return &auth.GetSessionResponse{UserID: srvResp}, nil
}
