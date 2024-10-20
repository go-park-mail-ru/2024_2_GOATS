package router

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func SetupAuth(delLayer handlers.AuthImplementationInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	authRouter := apiMux.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", delLayer.Login).Methods(http.MethodPost, http.MethodOptions)
	authRouter.HandleFunc("/logout", delLayer.Logout).Methods(http.MethodPost, http.MethodOptions)
	authRouter.HandleFunc("/signup", delLayer.Register).Methods(http.MethodPost, http.MethodOptions)
	authRouter.HandleFunc("/session", delLayer.Session).Methods(http.MethodGet, http.MethodOptions)
}

func SetupMovie(delLayer handlers.MovieImplementationInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	movieCollectionsRouter := apiMux.PathPrefix("/movie_collections").Subrouter()

	movieCollectionsRouter.HandleFunc("/", delLayer.GetCollections).Methods(http.MethodGet, http.MethodOptions)
}

func ActivateMiddlewares(mx *mux.Router, logger *zerolog.Logger) {
	mx.Use(middleware.AccessLogMiddleware(logger))
	mx.Use(middleware.PanicMiddleware(logger))
	mx.Use(middleware.CorsMiddleware(logger))
}
