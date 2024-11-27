package converter

import (
	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth/service/dto"
	auth "github.com/go-park-mail-ru/2024_2_GOATS/auth_service/pkg/auth_v1"
)

func ToSrvCreateCookieFromDesc(createCookieReq *auth.CreateSessionRequest) *dto.SrvCreateCookie {
	if createCookieReq == nil {
		return nil
	}

	return &dto.SrvCreateCookie{
		UserID: createCookieReq.UserID,
	}
}

func ToDescCreateCookieRespFromSrv(cookie *dto.Cookie) *auth.CreateSessionResponse {
	if cookie == nil {
		return nil
	}

	return &auth.CreateSessionResponse{
		Name:   cookie.Name,
		Cookie: cookie.TokenID,
		MaxAge: cookie.Expiry.Unix(),
	}
}
