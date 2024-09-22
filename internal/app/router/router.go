package router

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/handler"
	"github.com/gorilla/mux"
)

func Setup(ctx context.Context, api *api.Implementation) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/login", handler.Login(ctx, api, router)).Methods("POST")
	router.Handle("/signup", handler.Register(ctx, api, router)).Methods("POST")
	router.Handle("/movie_collections", handler.MovieCollections(ctx, api, router)).Methods("GET")

	// TODO: Add middleware using for sessions
	http.Handle("/", router)

	return router
}
