package router

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	webSocket "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/ws"
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
}

func SetupUser(delLayer handlers.UserImplementationInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	userRouter := apiMux.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("/{id:[0-9]+}/update_profile", delLayer.UpdateProfile).Methods(http.MethodPost, http.MethodOptions)
	userRouter.HandleFunc("/{id:[0-9]+}/update_password", delLayer.UpdatePassword).Methods(http.MethodPost, http.MethodOptions)
}

//// Настройка маршрутов для комнат и WebSocket
//func SetupRoom(ctx context.Context, hub *webSocket.RoomHub, roomHandler handlers.RoomImplementationInterface, router *mux.Router) {
//	apiMux := router.PathPrefix("/api").Subrouter()
//	roomRouter := apiMux.PathPrefix("/room").Subrouter()
//	roomRouter.HandleFunc("/create", roomHandler.CreateRoom).Methods(http.MethodPost, http.MethodOptions)
//	roomRouter.HandleFunc("/join", roomHandler.JoinRoom).Methods(http.MethodGet)
//
//	// Добавляем обработку WebSocket по пути /ws
//	router.HandleFunc("/ws", hub.HandleConnections).Methods(http.MethodGet)
//
//	//// Запуск сервера
//	//go func() {
//	//	log.Println("HTTP server started on :8000")
//	//	err := http.ListenAndServe(":8000", router)
//	//	if err != nil {
//	//		log.Fatal("ListenAndServe: ", err)
//	//	}
//	//}()
//
//}

func SetupRoom(hub *webSocket.RoomHub, roomHandler handlers.RoomImplementationInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	roomRouter := apiMux.PathPrefix("/room").Subrouter()
	roomRouter.HandleFunc("/create", roomHandler.CreateRoom).Methods(http.MethodPost)
	roomRouter.HandleFunc("/join", roomHandler.JoinRoom).Methods(http.MethodGet)

	router.HandleFunc("/ws", hub.HandleConnections).Methods(http.MethodGet)
}

func ActivateMiddlewares(mx *mux.Router) {
	mx.Use(middleware.AccessLogMiddleware)
	mx.Use(middleware.PanicMiddleware)
	mx.Use(middleware.CorsMiddleware)
}
