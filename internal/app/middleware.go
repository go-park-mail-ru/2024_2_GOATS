package app

import (
	"net/http"
)

const serviceStoppedMsg = "Services is not started"

func (a *App) AppReadyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !a.AcceptConnections {
			a.Logger.Error().Msg(serviceStoppedMsg)
			http.Error(w, serviceStoppedMsg, http.StatusServiceUnavailable)

			return
		}

		next.ServeHTTP(w, r)
	})
}
