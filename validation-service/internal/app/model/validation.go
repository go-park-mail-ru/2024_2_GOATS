package model

type UserRegisterData struct {
	Email           string
	Password        string
	PasswordConfirm string
	Sex             string
	Birthday        int
}

type ErrorResponse struct {
	Code  string
	Error string
}

type ValidationResponse struct {
	Success bool
	Errors  []ErrorResponse
}
