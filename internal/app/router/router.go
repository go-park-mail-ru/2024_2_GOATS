package router

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	webSocket "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/ws"
	csrf_handle "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/secur/csrf/handlers"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/middleware"
	"github.com/gorilla/mux"
)

// SetupAuth setups auth subrouter
func SetupAuth(delLayer handlers.AuthHandlerInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	authRouter := apiMux.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", delLayer.Login).Methods(http.MethodPost, http.MethodOptions).Name("LoginRoute")
	authRouter.HandleFunc("/logout", delLayer.Logout).Methods(http.MethodPost, http.MethodOptions).Name("LogoutRoute")
	authRouter.HandleFunc("/signup", delLayer.Register).Methods(http.MethodPost, http.MethodOptions).Name("SignupRoute")
	authRouter.HandleFunc("/session", delLayer.Session).Methods(http.MethodGet, http.MethodOptions).Name("SessionRoute")
}

// SetupMovie setups movie, actors, genres and movie_collections subrouters
func SetupMovie(delLayer handlers.MovieHandlerInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	movieRouter := apiMux.PathPrefix("/movies").Subrouter()
	actorRouter := apiMux.PathPrefix("/actors").Subrouter()
	genreRouter := apiMux.PathPrefix("/genres").Subrouter()
	movieCollectionsRouter := apiMux.PathPrefix("/movie_collections").Subrouter()

	movieCollectionsRouter.HandleFunc("/", delLayer.GetCollections).Methods(http.MethodGet, http.MethodOptions).Name("CollectionRoute")
	genreRouter.HandleFunc("/", delLayer.GetGenres).Methods(http.MethodGet, http.MethodOptions).Name("GenresRoute")
	movieRouter.HandleFunc("/{movie_id:[0-9]+}", delLayer.GetMovie).Methods(http.MethodGet, http.MethodOptions).Name("MovieRoute")
	// movieRouter.HandleFunc("/", delLayer.GetMovieByGenre).Methods(http.MethodGet, http.MethodOptions).Name("MovieByGenreRoute")
	actorRouter.HandleFunc("/{actor_id:[0-9]+}", delLayer.GetActor).Methods(http.MethodGet, http.MethodOptions).Name("ActorRoute")

	movieRouter.HandleFunc("/movies/search", delLayer.SearchMovies).Methods(http.MethodGet, http.MethodOptions).Name("MovieSearchRoute")
	movieRouter.HandleFunc("/actors/search", delLayer.SearchActors).Methods(http.MethodGet, http.MethodOptions).Name("ActorSearchRoute")
}

// SetupUser setups users subrouter
func SetupUser(delLayer handlers.UserHandlerInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	userRouter := apiMux.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("/{id:[0-9]+}/profile", delLayer.UpdateProfile).Methods(http.MethodPut, http.MethodOptions).Name("UpdateProfileRoute")
	userRouter.HandleFunc("/{id:[0-9]+}/password", delLayer.UpdatePassword).Methods(http.MethodPut, http.MethodOptions).Name("UpdatePasswordRoute")
	userRouter.HandleFunc("/{id:[0-9]+}/favorites", delLayer.GetFavorites).Methods(http.MethodGet, http.MethodOptions).Name("FavoritesRoute")
	userRouter.HandleFunc("/favorites", delLayer.SetFavorite).Methods(http.MethodPost, http.MethodOptions).Name("SetFavoriteRoute")
	userRouter.HandleFunc("/favorites", delLayer.ResetFavorite).Methods(http.MethodDelete, http.MethodOptions).Name("ResetFavoriteRoute")
}

// SetupPayment setups payments subrouter
func SetupPayment(delLayer handlers.PaymentHandlerInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	paymentRouter := apiMux.PathPrefix("/payments").Subrouter()

	paymentRouter.HandleFunc("/notify_yoo_money", delLayer.NotifyYooMoney).Methods(http.MethodPost, http.MethodOptions)
}

// SetupSubscription setups subscription subrouter
func SetupSubscription(delLayer handlers.SubscriptionHandlerInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	subscrRouter := apiMux.PathPrefix("/subscription").Subrouter()

	subscrRouter.HandleFunc("/", delLayer.Subscribe).Methods(http.MethodPost, http.MethodOptions)
}

func SetupRoom(hub *webSocket.RoomHub, roomHandler handlers.RoomImplementationInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	roomRouter := apiMux.PathPrefix("/room").Subrouter()
	roomRouter.HandleFunc("/create", roomHandler.CreateRoom).Methods(http.MethodPost, http.MethodOptions)
	roomRouter.HandleFunc("/join", roomHandler.JoinRoom).Methods(http.MethodGet)
}

// UseCommonMiddlewares activates common middlewares
func UseCommonMiddlewares(mx *mux.Router, authMW *middleware.SessionMiddleware) {
	mx.Use(middleware.AccessLogMiddleware)
	mx.Use(middleware.WithLogger)
	mx.Use(middleware.PanicMiddleware)
	mx.Use(middleware.CorsMiddleware)
	mx.Use(middleware.CsrfMiddleware)
	mx.Use(middleware.XSSMiddleware)
	mx.Use(authMW.AuthMiddleware)
}

// SetupCsrf setups csrf router
func SetupCsrf(router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	apiMux.HandleFunc("/csrf-token", csrf_handle.GenerateCSRFTokenHandler).Methods(http.MethodGet)
}
