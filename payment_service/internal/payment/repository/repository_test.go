package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/repository/dto"
	"github.com/stretchr/testify/assert"
)

const (
	createSQL = `
		INSERT INTO payments \(subscription_id, requested_amount\)
		VALUES \(\$1, \$2\)
		RETURNING id
	`

	updateSQL      = `UPDATE payments SET captured_amount = requested_amount, captured_at = \$1 WHERE id = \$2`
	amount         = 100
	subscriptionID = 1
	paymentID      = 1
)

func TestCreatePayment_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewPaymentRepository(db)
	expectedID := uint64(paymentID)
	paymentData := &dto.RepoPaymentData{
		SubscriptionID: subscriptionID,
		Amount:         amount,
	}

	mockRows := sqlmock.NewRows([]string{"id"}).AddRow(expectedID)
	mock.ExpectPrepare(createSQL).
		ExpectQuery().
		WithArgs(paymentData.SubscriptionID, paymentData.Amount).
		WillReturnRows(mockRows)

	pID, errObj := r.CreatePayment(context.Background(), paymentData)

	assert.Nil(t, errObj)
	assert.Equal(t, expectedID, pID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreatePayment_PrepareError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewPaymentRepository(db)
	paymentData := &dto.RepoPaymentData{
		SubscriptionID: subscriptionID,
		Amount:         amount,
	}

	mock.ExpectPrepare(createSQL).
		WillReturnError(errors.New("prepare statement error"))

	pID, errObj := r.CreatePayment(context.Background(), paymentData)

	assert.Equal(t, uint64(0), pID)
	assert.NotNil(t, errObj)
	assert.Contains(t, errObj.Error(), "prepareStatement#createPayment: prepare statement error")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreatePayment_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewPaymentRepository(db)
	paymentData := &dto.RepoPaymentData{
		SubscriptionID: subscriptionID,
		Amount:         amount,
	}

	mock.ExpectPrepare(createSQL).
		ExpectQuery().
		WithArgs(paymentData.SubscriptionID, paymentData.Amount).
		WillReturnError(errors.New("execution error"))

	pID, errObj := r.CreatePayment(context.Background(), paymentData)

	assert.Equal(t, uint64(0), pID)
	assert.NotNil(t, errObj)
	assert.Contains(t, errObj.Error(), "postgres: error while creating payment - execution error")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMarkPaid_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewPaymentRepository(db)
	paymentID := uint64(paymentID)

	mock.ExpectPrepare(updateSQL).
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), paymentID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	errObj := r.MarkPaid(context.Background(), paymentID)

	assert.Nil(t, errObj)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMarkPaid_PrepareError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewPaymentRepository(db)
	paymentID := uint64(paymentID)

	mock.ExpectPrepare(updateSQL).
		WillReturnError(errors.New("prepare statement error"))

	errObj := r.MarkPaid(context.Background(), paymentID)

	assert.NotNil(t, errObj)
	assert.Contains(t, errObj.Error(), "prepareStatement#markPaid: prepare statement error")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMarkPaid_ExecError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewPaymentRepository(db)
	paymentID := uint64(paymentID)

	mock.ExpectPrepare(updateSQL).
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), paymentID).
		WillReturnError(errors.New("execution error"))

	errObj := r.MarkPaid(context.Background(), paymentID)

	assert.NotNil(t, errObj)
	assert.Contains(t, errObj.Error(), "postgres: error while marking payment as paid - execution error")
	assert.NoError(t, mock.ExpectationsWereMet())
}
