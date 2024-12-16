package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	easyjson "github.com/mailru/easyjson"
	"github.com/rs/zerolog/log"
)

// ErrorDetails custom error struct
type ErrorDetails struct {
	Success bool
	Message string
}

// Response prepares http response
func Response(ctx context.Context, w http.ResponseWriter, code int, obj interface{}) {
	if obj == nil {
		w.WriteHeader(code)
		return
	}

	logger := log.Ctx(ctx)

	marshaler, ok := obj.(easyjson.Marshaler)
	if !ok {
		logger.Error().Msg("object does not implement easyjson.Marshaler")

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(obj); err != nil {
			logger.Error().Err(err).Msg("error while encoding response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(code)
		_, wErr := w.Write(buf.Bytes())
		if wErr != nil {
			logger.Error().Err(wErr).Msg("error while writing response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		return
	}

	var buf bytes.Buffer
	if _, err := easyjson.MarshalToWriter(marshaler, &buf); err != nil {
		logger.Error().Err(err).Msg("error while encoding success response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	_, writeErr := w.Write(buf.Bytes())
	if writeErr != nil {
		logger.Error().Err(writeErr).Msg("error while writing response to client")
	}
}

// DecodeBody decodes http request body
func DecodeBody(w http.ResponseWriter, r *http.Request, obj interface{}) bool {
	logger := log.Ctx(r.Context())

	unmarshaler, ok := obj.(easyjson.Unmarshaler)
	if !ok {
		logger.Error().Msg("object does not implement easyjson.Unmarshaler")
		if err := json.NewDecoder(r.Body).Decode(obj); err != nil {
			logger.Error().Err(err).Msg("cannot parse request")
			Response(r.Context(), w, http.StatusBadRequest, fmt.Errorf("cannot parse request: %w", err))
		}
		return false
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error().Err(err).Msg("cannot read request body")
		Response(r.Context(), w, http.StatusInternalServerError, fmt.Errorf("cannot read request body: %w", err))
		return false
	}

	if err := easyjson.Unmarshal(body, unmarshaler); err != nil {
		logger.Error().Err(err).Msg("cannot parse request")
		Response(r.Context(), w, http.StatusBadRequest, fmt.Errorf("cannot parse request: %w", err))
		return false
	}

	return true
}

// PreparedDefaultError is a prepared default http error
func PreparedDefaultError(code string, err error) *errVals.DeliveryError {
	errs := []errVals.ErrorItem{errVals.NewErrorItem(code, errVals.NewCustomError(err.Error()))}
	return errVals.NewDeliveryError(http.StatusForbidden, errs)
}

// RequestError is a common http request error method
func RequestError(ctx context.Context, w http.ResponseWriter, code string, status int, err error) {
	logger := log.Ctx(ctx)
	logger.Error().Err(err).Msg("request error")

	Response(ctx, w, status, PreparedDefaultError(code, err))
}
