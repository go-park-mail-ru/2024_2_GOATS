package service

import (
	"context"
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/service/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/service/dto"
	mockRepo "github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPaymentService_MarkPaid(t *testing.T) {
	tests := []struct {
		name          string
		paymentID     uint64
		mockSetup     func(mock *mockRepo.MockPaymentRepoInterface)
		expectedError error
	}{
		{
			name:      "Success - Mark Paid",
			paymentID: 123,
			mockSetup: func(mock *mockRepo.MockPaymentRepoInterface) {
				mock.EXPECT().
					MarkPaid(gomock.Any(), uint64(123)).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name:      "Repo Error",
			paymentID: 456,
			mockSetup: func(mock *mockRepo.MockPaymentRepoInterface) {
				mock.EXPECT().
					MarkPaid(gomock.Any(), uint64(456)).
					Return(errors.New("repo error"))
			},
			expectedError: errors.New("paymentService - failed to mark payment as paid: repo error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mockRepo.NewMockPaymentRepoInterface(ctrl)
			test.mockSetup(mockRepo)

			paymentService := &PaymentService{paymentRepo: mockRepo}
			err := paymentService.MarkPaid(context.Background(), test.paymentID)

			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPaymentService_CreatePayment(t *testing.T) {
	tests := []struct {
		name          string
		createData    *dto.CreatePaymentData
		mockSetup     func(mock *mockRepo.MockPaymentRepoInterface)
		expectedID    uint64
		expectedError error
	}{
		{
			name: "Success - Payment Created",
			createData: &dto.CreatePaymentData{
				SubscriptionID: 1,
				Amount:         100,
			},
			mockSetup: func(mock *mockRepo.MockPaymentRepoInterface) {
				mock.EXPECT().
					CreatePayment(gomock.Any(), converter.ConvertToRepoPaymentData(&dto.CreatePaymentData{
						SubscriptionID: 1,
						Amount:         100,
					})).
					Return(uint64(12345), nil)
			},
			expectedID:    12345,
			expectedError: nil,
		},
		{
			name: "Repo Error",
			createData: &dto.CreatePaymentData{
				SubscriptionID: 1,
				Amount:         100,
			},
			mockSetup: func(mock *mockRepo.MockPaymentRepoInterface) {
				mock.EXPECT().
					CreatePayment(gomock.Any(), converter.ConvertToRepoPaymentData(&dto.CreatePaymentData{
						SubscriptionID: 1,
						Amount:         100,
					})).
					Return(uint64(0), errors.New("repo error"))
			},
			expectedID:    0,
			expectedError: errors.New("paymentService - failed to create payment: repo error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mockRepo.NewMockPaymentRepoInterface(ctrl)
			test.mockSetup(mockRepo)

			paymentService := &PaymentService{paymentRepo: mockRepo}
			pID, err := paymentService.CreatePayment(context.Background(), test.createData)

			assert.Equal(t, test.expectedID, pID)
			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
