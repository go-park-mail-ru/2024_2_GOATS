package delivery

import (
	"context"
	"errors"
	"testing"

	srvMock "github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment/delivery/mocks"
	payment "github.com/go-park-mail-ru/2024_2_GOATS/payment_service/pkg/payment_v1"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPaymentHandler_Create(t *testing.T) {
	tests := []struct {
		name          string
		req           *payment.CreateRequest
		mockSetup     func(mock *srvMock.MockPaymentServiceInterface)
		expectedResp  *payment.PaymentID
		expectedError error
	}{
		{
			name: "Success",
			req: &payment.CreateRequest{
				SubscriptionID: 1,
				Amount:         100,
			},
			mockSetup: func(mock *srvMock.MockPaymentServiceInterface) {
				mock.EXPECT().
					CreatePayment(gomock.Any(), gomock.Any()).
					Return(uint64(1), nil)
			},
			expectedResp:  &payment.PaymentID{ID: 1},
			expectedError: nil,
		},
		{
			name: "ValidationError",
			req: &payment.CreateRequest{
				SubscriptionID: 0,
				Amount:         0,
			},
			mockSetup:     nil,
			expectedResp:  nil,
			expectedError: errors.New("incorrect params given"),
		},
		{
			name: "ServiceError",
			req: &payment.CreateRequest{
				SubscriptionID: 1,
				Amount:         100,
			},
			mockSetup: func(mock *srvMock.MockPaymentServiceInterface) {
				mock.EXPECT().
					CreatePayment(gomock.Any(), gomock.Any()).
					Return(uint64(0), errors.New("service error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("service error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockPaymentServiceInterface(ctrl)
			if test.mockSetup != nil {
				test.mockSetup(mockService)
			}

			handler := NewPaymentHandler(mockService)

			resp, err := handler.Create(context.Background(), test.req)

			assert.Equal(t, test.expectedResp, resp)
			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPaymentHandler_MarkPaid(t *testing.T) {
	tests := []struct {
		name          string
		req           *payment.PaymentID
		mockSetup     func(mock *srvMock.MockPaymentServiceInterface)
		expectedResp  *payment.Nothing
		expectedError error
	}{
		{
			name: "Success",
			req: &payment.PaymentID{
				ID: 1,
			},
			mockSetup: func(mock *srvMock.MockPaymentServiceInterface) {
				mock.EXPECT().
					MarkPaid(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedResp:  &payment.Nothing{Dummy: true},
			expectedError: nil,
		},
		{
			name: "ValidationError",
			req: &payment.PaymentID{
				ID: 0,
			},
			mockSetup:     nil,
			expectedResp:  nil,
			expectedError: errors.New("incorrect params given"),
		},
		{
			name: "ServiceError",
			req: &payment.PaymentID{
				ID: 1,
			},
			mockSetup: func(mock *srvMock.MockPaymentServiceInterface) {
				mock.EXPECT().
					MarkPaid(gomock.Any(), gomock.Any()).
					Return(errors.New("mark payment as paid error"))
			},
			expectedResp:  nil,
			expectedError: errors.New("mark payment as paid error"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := srvMock.NewMockPaymentServiceInterface(ctrl)
			if test.mockSetup != nil {
				test.mockSetup(mockService)
			}

			handler := NewPaymentHandler(mockService)

			resp, err := handler.MarkPaid(context.Background(), test.req)

			assert.Equal(t, test.expectedResp, resp)
			if test.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
