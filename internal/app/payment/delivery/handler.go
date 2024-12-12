package delivery

import (
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

// PaymentHandler is a payments handler struct
type PaymentHandler struct {
	paymentService PaymentServiceInterface
}

// NewPaymentHandler returns an instance of PaymentHandlerInterface
func NewPaymentHandler(srv PaymentServiceInterface) handlers.PaymentHandlerInterface {
	return &PaymentHandler{
		paymentService: srv,
	}
}

// NotifyYooMoney processes YooMoney callback
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

		return
	}

	callbackData, err := ph.parsePaymentCallback(r)
	if err != nil {
		logger.Error().Str("callback_error", "cannot parse request").Msg("bad_request")
		api.Response(r.Context(), w, http.StatusBadRequest, nil)

		return
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

func (ph *PaymentHandler) parsePaymentCallback(r *http.Request) (*api.PaymentCallback, error) {
	formValues := r.Form

	amount, err := convertToCopeck(formValues.Get("amount"))
	if err != nil {
		return nil, fmt.Errorf("invalid amount: %w", err)
	}

	wAmount, err := convertToCopeck(formValues.Get("withdraw_amount"))
	if err != nil {
		return nil, fmt.Errorf("invalid withdraw amount: %w", err)
	}

	datetime, err := parseDatetime(formValues.Get("datetime"))
	if err != nil {
		return nil, fmt.Errorf("invalid datetime: %w", err)
	}

	unaccepted, err := strconv.ParseBool(formValues.Get("unaccepted"))
	if err != nil {
		return nil, fmt.Errorf("invalid unaccepted value: %w", err)
	}

	codepro, err := strconv.ParseBool(formValues.Get("codepro"))
	if err != nil {
		return nil, fmt.Errorf("invalid codepro value: %w", err)
	}

	return &api.PaymentCallback{
		NotificationType: formValues.Get("notification_type"),
		OperationID:      formValues.Get("operation_id"),
		Amount:           amount,
		WithdrawAmount:   wAmount,
		Currency:         formValues.Get("currency"),
		DateTime:         datetime,
		Sender:           formValues.Get("sender"),
		Label:            formValues.Get("label"),
		Codepro:          codepro,
		Signature:        formValues.Get("sha1_hash"),
		Unaccepted:       unaccepted,
	}, nil
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
