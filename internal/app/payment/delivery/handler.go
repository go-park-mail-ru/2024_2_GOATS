package delivery

import (
	"context"
	"crypto/sha1"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type PaymentHandler struct {
	paymentService PaymentServiceInterface
}

func NewPaymentHandler(ctx context.Context, srv PaymentServiceInterface) handlers.PaymentHandlerInterface {
	return &PaymentHandler{
		paymentService: srv,
	}
}

func (ph *PaymentHandler) NotifyYooMoney(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Ошибка при разборе данных формы", http.StatusBadRequest)
		return
	}

	if !checkSignature(r) {
		logger.Error().Str("callback_error", "signature doesnt match").Msg("request_failed")
		api.Response(r.Context(), w, http.StatusBadRequest, nil)
	}

	return
	amount, _ := convertToCopeck(r.FormValue("amount"))
	wAmount, _ := convertToCopeck(r.FormValue("withdraw_amount"))
	datetime, _ := parseDatetime(r.FormValue("datetime"))
	unaccepted, _ := strconv.ParseBool(r.FormValue("unaccepted"))
	codepro, _ := strconv.ParseBool(r.FormValue("codepro"))

	callbackData := &api.PaymentCallback{
		NotificationType: r.FormValue("notification_type"),
		OperationID:      r.FormValue("operation_id"),
		Amount:           amount,
		WithdrawAmount:   wAmount,
		Currency:         r.FormValue("currency"),
		DateTime:         datetime,
		Sender:           r.FormValue("sender"),
		Label:            r.FormValue("label"),
		Codepro:          codepro,
		Signature:        r.FormValue("sha1_hash"),
		Unaccepted:       unaccepted,
	}

	srvData := converter.ToServPaymentCallback(callbackData)
	srvErr := ph.paymentService.ProcessCallback(r.Context(), srvData)
	respErr := errVals.ToDeliveryErrorFromService(srvErr)

	if respErr != nil {
		logger.Error().Err(srvErr.Error).Interface("callback_error", srvErr).Msg("request_failed")
		api.Response(r.Context(), w, respErr.HTTPStatus, respErr)
	}

	logger.Info().Interface("payment_callback", callbackData).Msg("got callback")

	api.Response(r.Context(), w, http.StatusOK, nil)
}

func convertToCopeck(amountStr string) (int64, error) {
	amountFloat, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return 0, fmt.Errorf("некорректное значение суммы: %w", err)
	}

	cents := int64(math.Round(amountFloat * 100))
	return cents, nil
}

func parseDatetime(datetimeStr string) (time.Time, error) {
	const layout = time.RFC3339
	parsedTime, err := time.Parse(layout, datetimeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("не удалось разобрать дату: %w", err)
	}
	return parsedTime, nil
}

func checkSignature(r *http.Request) bool {
	params := []string{
		"notification_type",
		"operation_id",
		"amount",
		"currency",
		"datetime",
		"sender",
		"codepro",
		"notification_secret",
		"label",
	}

	var values []string
	for _, param := range params {
		if param == "notification_secret" {
			values = append(values, viper.GetString("CALLBACK_SECRET"))
			continue
		}

		value := r.FormValue(param)
		values = append(values, value)
	}

	result := strings.Join(values, "&")

	return calculateSHA1(result) == r.FormValue("sha1_hash")
}

func calculateSHA1(data string) string {
	h := sha1.New()
	_, err := h.Write([]byte(data))
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
