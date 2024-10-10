package handlers

import (
	"encoding/json"
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
