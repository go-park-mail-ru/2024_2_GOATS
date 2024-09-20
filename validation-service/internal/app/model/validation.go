package model

type UserRegisterData struct {
	Email           string
	Password        string
	PasswordConfirm string
	Sex             int32
	Birthday        int
}

type ErrorResponse struct {
	Code     string
	ErrorObj error
}

type ValidationResponse struct {
	Success bool
	Errors  []ErrorResponse
}
