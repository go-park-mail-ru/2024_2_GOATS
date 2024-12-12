package delivery

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	mockSrv "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/payment/delivery/mocks"
)

func TestPaymentHandler_NotifyYooMoney(t *testing.T) {
	tests := []struct {
		name        string
		reqBody     url.Values
		mockErr     *errVals.ServiceError
		statusCode  int
		resp        string
		skipService bool
	}{
		{
			name: "Success",
			reqBody: url.Values{
				"notification_type": {"p2p-incoming"},
				"operation_id":      {"1234567"},
				"amount":            {"300.00"},
				"withdraw_amount":   {"100.00"},
				"datetime":          {"2011-07-01T09:00:00.000+04:00"},
				"unaccepted":        {"false"},
				"codepro":           {"false"},
				"sender":            {"41001XXXXXXXX"},
				"currency":          {"643"},
				"sha1_hash":         {"0fed47d1c25dfdadff6edfeed77b3405db5ffac0"},
				"label":             {"1-1"},
			},
			resp:       "",
			statusCode: http.StatusOK,
		},
		{
			name: "Invalid Signature",
			reqBody: url.Values{
				"notification_type": {"p2p-incoming"},
				"operation_id":      {"1234567"},
				"amount":            {"300.00"},
				"withdraw_amount":   {"100.00"},
				"datetime":          {"2011-07-01T09:00:00.000+04:00"},
				"unaccepted":        {"false"},
				"codepro":           {"false"},
				"sender":            {"41001XXXXXXXX"},
				"currency":          {"643"},
				"sha1_hash":         {"invalid_signature"},
				"label":             {"1-1"},
			},
			resp:        "",
			statusCode:  http.StatusBadRequest,
			skipService: true,
		},
		{
			name: "ProcessCallback Error",
			reqBody: url.Values{
				"notification_type": {"p2p-incoming"},
				"operation_id":      {"1234567"},
				"amount":            {"300.00"},
				"withdraw_amount":   {"100.00"},
				"datetime":          {"2011-07-01T09:00:00.000+04:00"},
				"unaccepted":        {"false"},
				"codepro":           {"false"},
				"sender":            {"41001XXXXXXXX"},
				"currency":          {"643"},
				"sha1_hash":         {"0fed47d1c25dfdadff6edfeed77b3405db5ffac0"},
				"label":             {"1-1"},
			},
			mockErr:    errVals.NewServiceError(errVals.ErrServerCode, errors.New("processing error")),
			resp:       `{"errors":[{"code":"something_went_wrong","error":"processing error"}]}`,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			req := httptest.NewRequest(http.MethodPost, "/notify_yoo_money", bytes.NewBufferString(test.reqBody.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			mockService := mockSrv.NewMockPaymentServiceInterface(ctrl)
			handler := &PaymentHandler{paymentService: mockService}

			if !test.skipService {
				setupTestEnv(t)
				mockService.EXPECT().
					ProcessCallback(gomock.Any(), gomock.Any()).
					Return(test.mockErr)
			}

			w := httptest.NewRecorder()
			handler.NotifyYooMoney(w, req)
			res := w.Result()

			defer func() {
				if err := res.Body.Close(); err != nil {
					t.Errorf("cannot close NotifyYooMoney response body")
				}
			}()

			assert.Equal(t, test.statusCode, res.StatusCode)
			if test.resp == "" {
				assert.Equal(t, test.resp, w.Body.String())
			} else {
				assert.JSONEq(t, test.resp, w.Body.String())
			}
		})
	}
}

func setupTestEnv(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_env_*.env")
	assert.NoError(t, err)
	defer func() {
		clErr := os.Remove(tmpFile.Name())
		assert.NoError(t, clErr)
	}()

	envContent := `CALLBACK_SECRET=callback_secret`
	_, err = tmpFile.Write([]byte(envContent))
	assert.NoError(t, err)

	err = tmpFile.Close()
	assert.NoError(t, err)

	viper.SetConfigFile(tmpFile.Name())
	viper.SetConfigType("env")
	err = viper.ReadInConfig()
	assert.NoError(t, err)

	assert.Equal(t, "callback_secret", viper.GetString("CALLBACK_SECRET"))
}
