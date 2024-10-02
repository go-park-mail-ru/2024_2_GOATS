package app

import (
	"net/http"
)

func (a *App) AppReadyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !a.AcceptConnections {
			http.Error(w, "Services is not started", http.StatusServiceUnavailable)
			return
		}

		next.ServeHTTP(w, r)
	})
}
