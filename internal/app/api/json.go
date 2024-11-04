package api

import (
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
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(obj)
	if err != nil {
		lg := log.Ctx(ctx)
		lg.Err(fmt.Errorf("error while encoding success response: %v", err))

		errObj := ErrorDetails{
			Message: err.Error(),
			Success: false,
		}

		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(errObj)
		if err != nil {
			lg.Err(fmt.Errorf("error while encoding error details: %v", err))
		}
	}
}

func DecodeBody(w http.ResponseWriter, r *http.Request, obj interface{}) {
	err := json.NewDecoder(r.Body).Decode(obj)
	if err != nil {
		lg := log.Ctx(r.Context())
		lg.Err(fmt.Errorf("cannot parse request: %w", err))
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
	errMsg := fmt.Errorf("request error: %w", err)
	lg := log.Ctx(ctx)
	lg.Error().Err(errMsg).Msg("request error")

	Response(ctx, w, status, PreparedDefaultError(code, errMsg))
}
