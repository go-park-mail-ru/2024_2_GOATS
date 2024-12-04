package delivery

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/rs/zerolog/log"
)

type SubscriptionHandler struct {
	subscriptionService SubscriptionServiceInterface
}

func NewSubscriptionHandler(ctx context.Context, srv SubscriptionServiceInterface) handlers.SubscriptionHandlerInterface {
	return &SubscriptionHandler{
		subscriptionService: srv,
	}
}

func (sh *SubscriptionHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	createSubReq := &api.SubscribeRequest{}
	api.DecodeBody(w, r, createSubReq)

	usrID := config.CurrentUserID(r.Context())
	if usrID == 0 {
		errMsg := errors.New("createSubscription action: unauthorized")
		api.RequestError(r.Context(), w, "check_subscription_request_parse_error", http.StatusForbidden, errMsg)

		return
	}

	srvData := converter.ToServCreateSubscription(createSubReq, usrID)
	subIDP, srvErr := sh.subscriptionService.Subscribe(r.Context(), srvData)
	respErr := errVals.ToDeliveryErrorFromService(srvErr)

	if respErr != nil {
		logger.Error().Err(srvErr.Error).Interface("createSubError", srvErr).Msg("request_failed")
		api.Response(r.Context(), w, respErr.HTTPStatus, respErr)
	}

	logger.Info().Str("subIDP", subIDP).Msg("successfully check subscription status")
	api.Response(r.Context(), w, http.StatusOK, &api.SubscribeResponse{SubscriptionIDP: subIDP})
}
