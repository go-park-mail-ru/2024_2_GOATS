package handlers

import (
	"net/http"
)

// MovieHandlerInterface defines movie_handler methods
type MovieHandlerInterface interface {
	GetCollections(w http.ResponseWriter, r *http.Request)
	GetGenres(w http.ResponseWriter, r *http.Request)
	GetMovie(w http.ResponseWriter, r *http.Request)
	GetActor(w http.ResponseWriter, r *http.Request)
	// GetMovieByGenre(w http.ResponseWriter, r *http.Request)
	SearchMovies(w http.ResponseWriter, r *http.Request)
	SearchActors(w http.ResponseWriter, r *http.Request)
	GetUserRating(w http.ResponseWriter, r *http.Request)
	AddOrUpdateRating(w http.ResponseWriter, r *http.Request)
	DeleteRating(w http.ResponseWriter, r *http.Request)
}

// AuthHandlerInterface defines auth_handler methods
type AuthHandlerInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Session(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

// UserHandlerInterface defines user_handler methods
type UserHandlerInterface interface {
	UpdateProfile(w http.ResponseWriter, r *http.Request)
	UpdatePassword(w http.ResponseWriter, r *http.Request)
	GetFavorites(w http.ResponseWriter, r *http.Request)
	SetFavorite(w http.ResponseWriter, r *http.Request)
	ResetFavorite(w http.ResponseWriter, r *http.Request)
}

// RoomImplementationInterface defines room_handler methods
type RoomImplementationInterface interface {
	CreateRoom(w http.ResponseWriter, r *http.Request)
	JoinRoom(w http.ResponseWriter, r *http.Request)
}

// PaymentHandlerInterface defines payment_handler methods
type PaymentHandlerInterface interface {
	NotifyYooMoney(w http.ResponseWriter, r *http.Request)
}

// SubscriptionHandlerInterface defines subscription_handler methods
type SubscriptionHandlerInterface interface {
	Subscribe(w http.ResponseWriter, r *http.Request)
}
