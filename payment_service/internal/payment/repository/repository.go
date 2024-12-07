package repository

import (
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/service"
)

type PaymentRepo struct {
	Database *sql.DB
}

func NewPaymentRepository(db *sql.DB) service.PaymentRepoInterface {
	return &PaymentRepo{
		Database: db,
	}
}
