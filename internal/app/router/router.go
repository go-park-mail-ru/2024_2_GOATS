package router

import (
	"context"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/gorilla/mux"
)

func SetupRouter(ctx context.Context, api *api.Implementation) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/login", LoginHandler(ctx, api, router)).Methods("POST")
	router.Handle("/signup", RegisterHandler(ctx, api, router)).Methods("POST")
	router.Handle("/movie_collections", MovieCollectionsHandler(ctx, api, router)).Methods("GET")

	// TODO: Add middleware using for sessions
	http.Handle("/", router)

	return router
}

func LoginHandler(ctx context.Context, api *api.Implementation, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func RegisterHandler(ctx context.Context, api *api.Implementation, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func MovieCollectionsHandler(ctx context.Context, api *api.Implementation, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handler")
		api.GetCollection(ctx)
	})
}
