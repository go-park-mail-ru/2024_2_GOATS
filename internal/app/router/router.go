package router

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	authApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/auth/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/handlers"
	movieApi "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/delivery"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/middleware"
	"github.com/gorilla/mux"
)

func SetupAuth(ctx context.Context, delLayer *authApi.Implementation, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	authRouter := apiMux.PathPrefix("/auth").Subrouter()

	authHandler := handlers.NewAuthHandler(delLayer, config.FromContext(ctx))
	authRouter.Handle("/login", authHandler.Login(router)).Methods(http.MethodPost, http.MethodOptions)
	authRouter.Handle("/logout", authHandler.Logout(router)).Methods(http.MethodPost, http.MethodOptions)
	authRouter.Handle("/signup", authHandler.Register(router)).Methods(http.MethodPost, http.MethodOptions)
	authRouter.Handle("/session", authHandler.Session(router)).Methods(http.MethodGet, http.MethodOptions)
}

func SetupMovie(ctx context.Context, delLayer *movieApi.Implementation, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	movieCollectionsRouter := apiMux.PathPrefix("/movie_collections").Subrouter()

	movieHandler := handlers.NewMovieHandler(delLayer, config.FromContext(ctx))
	movieCollectionsRouter.Handle("/", movieHandler.GetCollections(router)).Methods(http.MethodGet, http.MethodOptions)
}

func ActivateMiddlewares(mx *mux.Router) {
	mx.Use(middleware.CorsMiddleware)
	mx.Use(middleware.PanicMiddleware)
}
