package router

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	csrf_handle "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/secur/csrf/handlers"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/middleware"
	"github.com/gorilla/mux"
)

func SetupAuth(delLayer handlers.AuthHandlerInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	authRouter := apiMux.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", delLayer.Login).Methods(http.MethodPost, http.MethodOptions)
	authRouter.HandleFunc("/logout", delLayer.Logout).Methods(http.MethodPost, http.MethodOptions)
	authRouter.HandleFunc("/signup", delLayer.Register).Methods(http.MethodPost, http.MethodOptions)
	authRouter.HandleFunc("/session", delLayer.Session).Methods(http.MethodGet, http.MethodOptions)
}

func SetupMovie(delLayer handlers.MovieHandlerInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	movieRouter := apiMux.PathPrefix("/movies").Subrouter()
	actorRouter := apiMux.PathPrefix("/actors").Subrouter()
	genreRouter := apiMux.PathPrefix("/genres").Subrouter()
	movieCollectionsRouter := apiMux.PathPrefix("/movie_collections").Subrouter()

	movieCollectionsRouter.HandleFunc("/", delLayer.GetCollections).Methods(http.MethodGet, http.MethodOptions)
	genreRouter.HandleFunc("/", delLayer.GetGenres).Methods(http.MethodGet, http.MethodOptions)
	movieRouter.HandleFunc("/{movie_id:[0-9]+}", delLayer.GetMovie).Methods(http.MethodGet, http.MethodOptions)
	movieRouter.HandleFunc("/", delLayer.GetMovieByGenre).Methods(http.MethodGet, http.MethodOptions)
	actorRouter.HandleFunc("/{actor_id:[0-9]+}", delLayer.GetActor).Methods(http.MethodGet, http.MethodOptions)
}

func SetupUser(delLayer handlers.UserHandlerInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	userRouter := apiMux.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("/{id:[0-9]+}/profile", delLayer.UpdateProfile).Methods(http.MethodPut, http.MethodOptions)
	userRouter.HandleFunc("/{id:[0-9]+}/password", delLayer.UpdatePassword).Methods(http.MethodPut, http.MethodOptions)
	userRouter.HandleFunc("/{id:[0-9]+}/favorites", delLayer.GetFavorites).Methods(http.MethodGet, http.MethodOptions)
	userRouter.HandleFunc("/favorites", delLayer.SetFavorite).Methods(http.MethodPost, http.MethodOptions)
	userRouter.HandleFunc("/favorites", delLayer.ResetFavorite).Methods(http.MethodDelete, http.MethodOptions)
}

func UseCommonMiddlewares(mx *mux.Router, authMW *middleware.SessionMiddleware) {
	mx.Use(middleware.AccessLogMiddleware)
	mx.Use(middleware.WithLogger)
	mx.Use(middleware.PanicMiddleware)
	mx.Use(middleware.CorsMiddleware)
	mx.Use(middleware.CsrfMiddleware)
	mx.Use(middleware.XssMiddleware)
	mx.Use(authMW.AuthMiddleware)
}

func SetupCsrf(router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	apiMux.HandleFunc("/csrf-token", csrf_handle.GenerateCSRFTokenHandler).Methods(http.MethodGet)
}
