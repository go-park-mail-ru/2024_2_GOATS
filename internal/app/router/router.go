package router

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	webSocket "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/ws"
	csrf_handle "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/secur/csrf/handlers"
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
	actorRouter := apiMux.PathPrefix("/actors").Subrouter()

	movieCollectionsRouter.HandleFunc("/", delLayer.GetCollections).Methods(http.MethodGet, http.MethodOptions)
	movieRouter.HandleFunc("/{movie_id:[0-9]+}", delLayer.GetMovie).Methods(http.MethodGet, http.MethodOptions)
	actorRouter.HandleFunc("/{actor_id:[0-9]+}", delLayer.GetActor).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc("/movies/search", delLayer.SearchMovies).Methods("GET")

	router.HandleFunc("/actors/search", delLayer.SearchActors).Methods("GET")
}

func SetupUser(delLayer handlers.UserImplementationInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	userRouter := apiMux.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("/{id:[0-9]+}/update_profile", delLayer.UpdateProfile).Methods(http.MethodPost, http.MethodOptions)
	userRouter.HandleFunc("/{id:[0-9]+}/update_password", delLayer.UpdatePassword).Methods(http.MethodPost, http.MethodOptions)
}

func SetupRoom(hub *webSocket.RoomHub, roomHandler handlers.RoomImplementationInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	roomRouter := apiMux.PathPrefix("/room").Subrouter()
	roomRouter.HandleFunc("/create", roomHandler.CreateRoom).Methods(http.MethodPost, http.MethodOptions)
	roomRouter.HandleFunc("/join", roomHandler.JoinRoom).Methods(http.MethodGet)
}

func ActivateMiddlewares(mx *mux.Router) {
	//mx.Use(middleware.CsrfMiddleware())
	//mx.Use(middleware.XssMiddleware)
	mx.Use(middleware.AccessLogMiddleware)
	mx.Use(middleware.PanicMiddleware)
	mx.Use(middleware.CorsMiddleware)
}

func SetupCsrf(router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	apiMux.HandleFunc("/csrf-token", csrf_handle.GenerateCSRFTokenHandler).Methods(http.MethodGet)
}
