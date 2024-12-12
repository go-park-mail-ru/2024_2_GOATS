package delivery

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	mockSrv "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/subscription/delivery/mocks"
)

func TestSubscriptionHandler_Subscribe(t *testing.T) {
	tests := []struct {
		name        string
		reqBody     string
		userID      int
		resp        string
		mockResp    string
		mockErr     *errVals.ServiceError
		statusCode  int
		skipService bool
	}{
		{
			name:       "Success",
			reqBody:    `{"amount": 100}`,
			userID:     123,
			mockResp:   "1-1",
			statusCode: http.StatusOK,
		},
		{
			name:        "Unauthorized user",
			reqBody:     `{"amount": 100}`,
			userID:      0,
			resp:        `{"errors":[{"code":"check_subscription_request_parse_error","error":"createSubscription action: unauthorized"}]}`,
			statusCode:  http.StatusForbidden,
			skipService: true,
		},
		{
			name:       "Service error",
			reqBody:    `{"amount": 100}`,
			userID:     123,
			mockErr:    errVals.NewServiceError(errVals.ErrServerCode, errors.New("internal server error")),
			resp:       `{"errors":[{"code":"something_went_wrong","error":"internal server error"}]}`,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mockSrv.NewMockSubscriptionServiceInterface(ctrl)

			handler := &SubscriptionHandler{subscriptionService: mockService}

			req := httptest.NewRequest(http.MethodPost, "/subscribe", bytes.NewBufferString(test.reqBody))
			req.Header.Set("Content-Type", "application/json")

			ctx := req.Context()
			ctx = context.WithValue(ctx, config.CurrentUserKey{}, test.userID)
			req = req.WithContext(ctx)

			if !test.skipService {
				mockService.EXPECT().
					Subscribe(gomock.Any(), gomock.Any()).
					Return(test.mockResp, test.mockErr)
			}

			w := httptest.NewRecorder()
			handler.Subscribe(w, req)
			res := w.Result()

			defer func() {
				clErr := res.Body.Close()
				assert.NoError(t, clErr)
			}()

			assert.Equal(t, test.statusCode, res.StatusCode)

			if test.resp != "" {
				body, _ := io.ReadAll(res.Body)
				assert.JSONEq(t, test.resp, string(body))
			}
		})
	}
}
