package handlers

import (
	"net/http"
)

//go:generate mockgen -source=interface.go -destination=mocks/mock.go
type MovieImplementationInterface interface {
	GetCollections(w http.ResponseWriter, r *http.Request)
}

type AuthImplementationInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Session(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type RoomImplementationInterface interface {
	CreateRoom(w http.ResponseWriter, r *http.Request)
	JoinRoom(w http.ResponseWriter, r *http.Request)
}
