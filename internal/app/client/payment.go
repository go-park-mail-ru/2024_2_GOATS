package client

import (
	"context"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	payment "github.com/go-park-mail-ru/2024_2_GOATS/payment_service/pkg/payment_v1"
)

//go:generate mockgen -source=auth.go -destination=../auth/service/mocks/mock.go
type PaymentClientInterface interface {
	MarkPaid(ctx context.Context, pID int) error
	CreatePayment(ctx context.Context, data *models.CreatePaymentData) (int, error)
}

type PaymentClient struct {
	paymentMS payment.PaymentRPCClient
}

func NewPaymentClient(paymentMS payment.PaymentRPCClient) PaymentClientInterface {
	return &PaymentClient{
		paymentMS: paymentMS,
	}
}

func (pc *PaymentClient) MarkPaid(ctx context.Context, pID int) error {
	start := time.Now()
	method := "MarkPaid"

	_, err := pc.paymentMS.MarkPaid(ctx, &payment.PaymentID{ID: uint64(pID)})

	saveMetric(start, userClient, method, err)

	if err != nil {
		return fmt.Errorf("paymentClientError#markPaid: %w", err)
	}

	return nil
}

func (pc *PaymentClient) CreatePayment(ctx context.Context, data *models.CreatePaymentData) (int, error) {
	start := time.Now()
	method := "CreatePayment"

	resp, err := pc.paymentMS.Create(ctx, &payment.CreateRequest{
		SubscriptionID: uint64(data.SubscriptionID),
		Amount:         data.Amount,
	})

	saveMetric(start, userClient, method, err)

	if err != nil {
		return 0, fmt.Errorf("paymentClientError#createPayment: %w", err)
	}

	return int(resp.ID), nil
}
