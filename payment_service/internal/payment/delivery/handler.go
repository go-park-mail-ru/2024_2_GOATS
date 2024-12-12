package delivery

import (
	"context"
	"errors"

	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/errs"
	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/delivery/converter"
	payment "github.com/go-park-mail-ru/2024_2_GOATS/payment_service/pkg/payment_v1"
	"github.com/rs/zerolog/log"
)

// PaymentHandler grpc payments handler
type PaymentHandler struct {
	payment.UnimplementedPaymentRPCServer
	paymentService PaymentServiceInterface
}

// NewPaymentHandler returns an instance of PaymentRPCServer
func NewPaymentHandler(usrSrv PaymentServiceInterface) payment.PaymentRPCServer {
	return &PaymentHandler{
		paymentService: usrSrv,
	}
}

// Create grpc create payment handler
func (uh *PaymentHandler) Create(ctx context.Context, createPayReq *payment.CreateRequest) (*payment.PaymentID, error) {
	logger := log.Ctx(ctx)
	if createPayReq.Amount == 0 || createPayReq.SubscriptionID == 0 {
		err := errors.New("incorrect params given")
		logger.Error().Err(err).Msg("bad_request")
		return nil, err
	}

	srvData := converter.ConvertToSrvPayment(createPayReq)
	if srvData == nil {
		logger.Error().Msgf("convert error")
		return nil, errs.ErrBadRequest
	}

	pID, err := uh.paymentService.CreatePayment(ctx, srvData)
	if err != nil {
		logger.Error().Interface("createPaymentError", err).Msg("failed_to_create_payment")
		return nil, err
	}

	logger.Info().Uint64("createPaymentSuccess", pID).Msg("successfully_create_payment")
	return &payment.PaymentID{ID: pID}, nil
}

// MarkPaid grpc mark payment as paid handler
func (uh *PaymentHandler) MarkPaid(ctx context.Context, req *payment.PaymentID) (*payment.Nothing, error) {
	logger := log.Ctx(ctx)
	if req.ID == 0 {
		err := errors.New("incorrect params given")
		logger.Error().Err(err).Msg("bad_request")
		return nil, err
	}

	err := uh.paymentService.MarkPaid(ctx, req.ID)
	if err != nil {
		logger.Error().Interface("markPaidError", err).Msg("failed_to_mark_payment_paid")
		return nil, err
	}

	logger.Info().Msg("successfully_marked_payment_paid")
	return &payment.Nothing{Dummy: true}, nil
}
