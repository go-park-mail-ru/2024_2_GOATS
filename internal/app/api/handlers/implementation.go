package handlers

import (
	"net/http"
)

type MovieHandlerInterface interface {
	GetCollections(w http.ResponseWriter, r *http.Request)
	GetGenres(w http.ResponseWriter, r *http.Request)
	GetMovie(w http.ResponseWriter, r *http.Request)
	GetActor(w http.ResponseWriter, r *http.Request)
	// GetMovieByGenre(w http.ResponseWriter, r *http.Request)
	SearchMovies(w http.ResponseWriter, r *http.Request)
	SearchActors(w http.ResponseWriter, r *http.Request)
}

type AuthHandlerInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Session(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type UserHandlerInterface interface {
	UpdateProfile(w http.ResponseWriter, r *http.Request)
	UpdatePassword(w http.ResponseWriter, r *http.Request)
	GetFavorites(w http.ResponseWriter, r *http.Request)
	SetFavorite(w http.ResponseWriter, r *http.Request)
	ResetFavorite(w http.ResponseWriter, r *http.Request)
	GetWatchedMovies(w http.ResponseWriter, r *http.Request)
	AddWatchedMovie(w http.ResponseWriter, r *http.Request)
}

type RoomImplementationInterface interface {
	CreateRoom(w http.ResponseWriter, r *http.Request)
	JoinRoom(w http.ResponseWriter, r *http.Request)
}

type PaymentHandlerInterface interface {
	NotifyYooMoney(w http.ResponseWriter, r *http.Request)
}

type SubscriptionHandlerInterface interface {
	Subscribe(w http.ResponseWriter, r *http.Request)
}
