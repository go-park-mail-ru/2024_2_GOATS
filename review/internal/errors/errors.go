package errors

import "encoding/json"

type CustomError struct {
	Err error
}

type ErrorObj struct {
	Code  string
	Error CustomError
}

type SrvErrorObj struct {
	Code  string
	Error CustomError
}

func (ce *CustomError) MarshalJSON() ([]byte, error) {
	return json.Marshal(ce.Err.Error())
}

func NewErrorObj(code string, text CustomError) *ErrorObj {
	return &ErrorObj{
		Code:  code,
		Error: text,
	}
}
