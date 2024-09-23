package router

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	authHandlers "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/handlers/auth"
	movieCollectionHandlers "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/handlers/movie_collections"
	"github.com/gorilla/mux"
)

func Setup(ctx context.Context, api *api.Implementation) *mux.Router {
	router := mux.NewRouter()

	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.Handle("/login", authHandlers.Login(ctx, api, router)).Methods("POST")
	authRouter.Handle("/signup", authHandlers.Register(ctx, api, router)).Methods("POST")

	movieCollectionsRouter := router.PathPrefix("/movie_collections").Subrouter()
	movieCollectionsRouter.Handle("", movieCollectionHandlers.MovieCollections(ctx, api, router)).Methods("GET")

	// TODO: Add middleware using for sessions
	http.Handle("/", router)

	return router
}
