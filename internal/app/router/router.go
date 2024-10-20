package router

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/middleware"
	"github.com/gorilla/mux"
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
	movieRouter := apiMux.PathPrefix("/movies").Subrouter()

	movieCollectionsRouter.HandleFunc("/", delLayer.GetCollections).Methods(http.MethodGet, http.MethodOptions)
	movieRouter.HandleFunc("/{movie_id:[0-9]+}", delLayer.GetMovie).Methods(http.MethodGet, http.MethodOptions)
}

func SetupUser(delLayer handlers.UserImplementationInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	userRouter := apiMux.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("/{id:[0-9]+}/update_profile", delLayer.UpdateProfile).Methods(http.MethodPost, http.MethodOptions)
	userRouter.HandleFunc("/{id:[0-9]+}/update_password", delLayer.UpdatePassword).Methods(http.MethodPost, http.MethodOptions)
}

func ActivateMiddlewares(mx *mux.Router) {
	mx.Use(middleware.AccessLogMiddleware)
	mx.Use(middleware.PanicMiddleware)
	mx.Use(middleware.CorsMiddleware)
}
