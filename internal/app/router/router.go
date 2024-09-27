package router

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/handlers"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/middleware"
	"github.com/gorilla/mux"
)

func Setup(ctx context.Context, api *api.Implementation) *mux.Router {
	router := mux.NewRouter()
	apiMux := router.PathPrefix("/api").Subrouter()

	authRouter := apiMux.PathPrefix("/auth").Subrouter()
	authHandler := handlers.NewAuthHandler(api)
	authRouter.Handle("/login", authHandler.Login(router)).Methods(http.MethodPost, http.MethodOptions)
	authRouter.Handle("/signup", authHandler.Register(router)).Methods(http.MethodPost, http.MethodOptions)
	authRouter.Handle("/session", authHandler.Session(router)).Methods(http.MethodGet, http.MethodOptions)

	movieCollectionsRouter := apiMux.PathPrefix("/movie_collections").Subrouter()
	movieHandler := handlers.NewMovieHandler(api)
	movieCollectionsRouter.Handle("/", movieHandler.GetCollections(router)).Methods(http.MethodGet, http.MethodOptions)

	apiMux.Use(middleware.CorsMiddleware)

	http.Handle("/", router)

	return router
}
