package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/gommon/log"
)

type ErrorDetails struct {
	Success bool
	Message string
}

func Response(w http.ResponseWriter, code int, obj interface{}) {
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(obj)
	if err != nil {
		log.Errorf("error while encoding success auth response: %v", err)

		errObj := ErrorDetails{
			Message: err.Error(),
			Success: false,
		}

		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(errObj)
		if err != nil {
			log.Errorf("не удалось закодировать детали ошибки: %v", err)
		}
	}
}

func DecodeBody(w http.ResponseWriter, r *http.Request, obj interface{}) {
	err := json.NewDecoder(r.Body).Decode(obj)
	if err != nil {
		Response(w, http.StatusBadRequest, fmt.Errorf("cannot parse request: %w", err))
		return
	}
}
