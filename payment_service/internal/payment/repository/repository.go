package repository

import (
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/service"
)

// PaymentRepo is a payment_service repository layer struct
type PaymentRepo struct {
	Database *sql.DB
}

// NewPaymentRepository returns an instance of PaymentRepoInterface
func NewPaymentRepository(db *sql.DB) service.PaymentRepoInterface {
	return &PaymentRepo{
		Database: db,
	}
}
