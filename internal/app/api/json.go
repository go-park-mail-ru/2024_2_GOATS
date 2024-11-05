package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/rs/zerolog/log"
)

type ErrorDetails struct {
	Success bool
	Message string
}

func Response(ctx context.Context, w http.ResponseWriter, code int, obj interface{}) {
	var buf bytes.Buffer
	logger := log.Ctx(ctx)
	err := json.NewEncoder(&buf).Encode(obj)

	if err != nil {
		logger.Error().Err(err).Msg("error while encoding success response")

		errObj := ErrorDetails{
			Message: err.Error(),
			Success: false,
		}

		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(errObj)
		if err != nil {
			logger.Error().Err(err).Msg("error while encoding error details")
		}

		return
	}

	w.WriteHeader(code)
	_, writeErr := w.Write(buf.Bytes())
	if writeErr != nil {
		logger.Error().Err(writeErr).Msg("error while writing response to client")
	}
}

func DecodeBody(w http.ResponseWriter, r *http.Request, obj interface{}) {
	err := json.NewDecoder(r.Body).Decode(obj)
	if err != nil {
		logger := log.Ctx(r.Context())
		logger.Error().Err(err).Msg("cannot parse request")
		Response(r.Context(), w, http.StatusBadRequest, fmt.Errorf("cannot parse request: %w", err))

		return
	}
}

func PreparedDefaultError(code string, err error) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusForbidden,
		Errors: []errVals.ErrorObj{{
			Code:  code,
			Error: errVals.CustomError{Err: err},
		}},
	}
}

func RequestError(ctx context.Context, w http.ResponseWriter, code string, status int, err error) {
	logger := log.Ctx(ctx)
	logger.Error().Err(err).Msg("request error")

	Response(ctx, w, status, PreparedDefaultError(code, err))
}
