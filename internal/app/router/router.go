package router

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	auth "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/handlers/auth"
	movie "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/handlers/movie_collections"
	"github.com/gorilla/mux"
)

func Setup(ctx context.Context, api *api.Implementation) *mux.Router {
	router := mux.NewRouter()

	authRouter := router.PathPrefix("/auth").Subrouter()
	authHandler := auth.NewHandler(ctx, api)
	authRouter.Handle("/login", authHandler.Login(router)).Methods("POST")
	authRouter.Handle("/signup", authHandler.Register(router)).Methods("POST")
	authRouter.Handle("/session", authHandler.Session(router)).Methods("GET")

	movieCollectionsRouter := router.PathPrefix("/movie_collections").Subrouter()
	movieHandler := movie.NewHandler(ctx, api)
	movieCollectionsRouter.Handle("", movieHandler.GetCollections(router)).Methods("GET")

	// TODO: Add middleware using for sessions
	http.Handle("/", router)

	return router
}
