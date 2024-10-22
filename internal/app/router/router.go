package router

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	webSocket "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/room/ws"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/middleware"

	"github.com/gorilla/mux"
)

// Настройка маршрутов для аутентификации
func SetupAuth(ctx context.Context, delLayer handlers.AuthImplementationInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	authRouter := apiMux.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", delLayer.Login).Methods(http.MethodPost, http.MethodOptions)
	authRouter.HandleFunc("/logout", delLayer.Logout).Methods(http.MethodPost, http.MethodOptions)
	authRouter.HandleFunc("/signup", delLayer.Register).Methods(http.MethodPost, http.MethodOptions)
	authRouter.HandleFunc("/session", delLayer.Session).Methods(http.MethodGet, http.MethodOptions)
}

// Настройка маршрутов для фильмов
func SetupMovie(ctx context.Context, delLayer handlers.MovieImplementationInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	movieCollectionsRouter := apiMux.PathPrefix("/movie_collections").Subrouter()

	movieCollectionsRouter.HandleFunc("/", delLayer.GetCollections).Methods(http.MethodGet, http.MethodOptions)
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

func SetupRoom(ctx context.Context, hub *webSocket.RoomHub, roomHandler handlers.RoomImplementationInterface, router *mux.Router) {
	apiMux := router.PathPrefix("/api").Subrouter()
	roomRouter := apiMux.PathPrefix("/room").Subrouter()
	roomRouter.HandleFunc("/create", roomHandler.CreateRoom).Methods(http.MethodPost)
	roomRouter.HandleFunc("/join", roomHandler.JoinRoom).Methods(http.MethodGet)

	router.HandleFunc("/ws", hub.HandleConnections).Methods(http.MethodGet)
}

func ActivateMiddlewares(mx *mux.Router) {
	mx.Use(middleware.CorsMiddleware)
	mx.Use(middleware.PanicMiddleware)
}
