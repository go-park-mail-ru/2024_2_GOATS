package router

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/middleware"
	"github.com/gorilla/mux"
)

func SetupAuth(ctx context.Context, delLayer handlers.AuthImplementationInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	authRouter := apiMux.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", delLayer.Login).Methods(http.MethodPost, http.MethodOptions)
	authRouter.HandleFunc("/logout", delLayer.Logout).Methods(http.MethodPost, http.MethodOptions)
	authRouter.HandleFunc("/signup", delLayer.Register).Methods(http.MethodPost, http.MethodOptions)
	authRouter.HandleFunc("/session", delLayer.Session).Methods(http.MethodGet, http.MethodOptions)
}

func SetupMovie(ctx context.Context, delLayer handlers.MovieImplementationInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	movieCollectionsRouter := apiMux.PathPrefix("/movie_collections").Subrouter()

	movieCollectionsRouter.HandleFunc("/", delLayer.GetCollections).Methods(http.MethodGet, http.MethodOptions)
}

func ActivateMiddlewares(mx *mux.Router) {
	mx.Use(middleware.CorsMiddleware)
	mx.Use(middleware.PanicMiddleware)
}
