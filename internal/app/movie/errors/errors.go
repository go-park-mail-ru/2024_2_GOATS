package errors

import (
	"encoding/json"
	errs "errors"
	"net/http"

	"github.com/lib/pq"
)

const (
	DuplicateErrKey        = "23505"
	DuplicateErrCode       = "db_duplicate_entry"
	ErrBrokenCookieCode    = "broken_cookie"
	ErrNoCookieCode        = "no_cookie_provided"
	ErrRedisGetCode        = "failed_read_from_redis"
	ErrTransformationCode  = "transformation_error"
	ErrServerCode          = "something_went_wrong"
	ErrGenerateTokenCode   = "auth_token_generation_error"
	ErrCreateUserCode      = "create_user_error"
	ErrUnauthorizedCode    = "user_unauthorized"
	ErrRedisClearCode      = "failed_delete_from_redis"
	ErrRedisWriteCode      = "failed_write_into_redis"
	ErrFileUploadCode      = "file_upload_err"
	ErrUpdatePasswordCode  = "update_password_error"
	ErrUpdateProfileCode   = "update_profile_error"
	ErrInvalidEmailCode    = "invalid_email"
	ErrInvalidPasswordCode = "invalid_password"
	ErrUserNotFoundCode    = "user_not_found"
	ErrInvalidUsernameCode = "invalid_username"
	ErrGenCSRFCode         = "failed_to_generate_csrf_token"
	ErrSaveFileCode        = "failed_to_save_image"
	ErrCreateFavorite      = "create_favorite_error"
	ErrResetFavorite       = "destroy_favorite_error"
	ErrGetFavorites        = "get_favorite_error"
	ErrGetUserCode         = "cannot_find_user_by_given_params"
	ErrCreateSessionCode   = "failed_to_create_session"
	ErrDestroySessionCode  = "failed_to_destroy_session"
	ErrCheckSessionCode    = "failed_to_get_session_data"
)

var ErrorCodeToHTTPStatus = map[string]int{
	DuplicateErrCode:       http.StatusConflict,
	ErrBrokenCookieCode:    http.StatusBadRequest,
	ErrNoCookieCode:        http.StatusBadRequest,
	ErrRedisGetCode:        http.StatusInternalServerError,
	ErrTransformationCode:  http.StatusInternalServerError,
	ErrServerCode:          http.StatusInternalServerError,
	ErrGenerateTokenCode:   http.StatusInternalServerError,
	ErrCreateUserCode:      http.StatusInternalServerError,
	ErrUnauthorizedCode:    http.StatusUnauthorized,
	ErrRedisClearCode:      http.StatusInternalServerError,
	ErrRedisWriteCode:      http.StatusInternalServerError,
	ErrFileUploadCode:      http.StatusInternalServerError,
	ErrUpdatePasswordCode:  http.StatusInternalServerError,
	ErrUpdateProfileCode:   http.StatusInternalServerError,
	ErrInvalidEmailCode:    http.StatusBadRequest,
	ErrInvalidPasswordCode: http.StatusBadRequest,
	ErrUserNotFoundCode:    http.StatusNotFound,
	ErrInvalidUsernameCode: http.StatusBadRequest,
	ErrGenCSRFCode:         http.StatusInternalServerError,
	ErrSaveFileCode:        http.StatusInternalServerError,
	ErrCreateFavorite:      http.StatusInternalServerError,
	ErrResetFavorite:       http.StatusInternalServerError,
	ErrGetFavorites:        http.StatusInternalServerError,
	ErrGetUserCode:         http.StatusNotFound,
	ErrCreateSessionCode:   http.StatusInternalServerError,
	ErrDestroySessionCode:  http.StatusInternalServerError,
	ErrCheckSessionCode:    http.StatusInternalServerError,
}

var (
	ErrInvalidEmail          = NewCustomError("email is incorrect")
	ErrInvalidPassword       = NewCustomError("password is too short. The minimal len is 8")
	ErrInvalidUsername       = NewCustomError("username is too short. The minimal len is 6")
	ErrInvalidPasswordsMatch = NewCustomError("password doesn't match with passwordConfirmation")
	ErrInvalidOldPassword    = NewCustomError("invalid old password")
	ErrUserNotFound          = NewCustomError("cannot find user by given params")
	ErrBrokenCookie          = NewCustomError("broken cookie was given")
	ErrSaveFile              = NewCustomError("cannot save file")
)

type CustomError struct {
	Err error
}

type ErrorItem struct {
	Code  string      `json:"code"`
	Error CustomError `json:"error"`
}

type DeliveryError struct {
	HTTPStatus int         `json:"-"`
	Errors     []ErrorItem `json:"errors"`
}

type ServiceError struct {
	Code  string
	Error error
}

type RepoError struct {
	Code  string
	Error CustomError
}

// Функции для работы с ошибками
func IsDuplicateError(err error) bool {
	var pqErr *pq.Error
	return errs.As(err, &pqErr) && pqErr.Code == DuplicateErrKey
}

func (ce *CustomError) MarshalJSON() ([]byte, error) {
	return json.Marshal(ce.Err.Error())
}

func NewCustomError(message string) CustomError {
	return CustomError{Err: errs.New(message)}
}

func NewRepoError(code string, err CustomError) *RepoError {
	return &RepoError{
		Code:  code,
		Error: err,
	}
}

func NewServiceError(code string, err error) *ServiceError {
	return &ServiceError{
		Code:  code,
		Error: err,
	}
}

func NewDeliveryError(status int, errs []ErrorItem) *DeliveryError {
	return &DeliveryError{
		HTTPStatus: status,
		Errors:     errs,
	}
}

func NewErrorItem(code string, err CustomError) ErrorItem {
	return ErrorItem{
		Code:  code,
		Error: err,
	}
}

func matchStatus(code string) int {
	if status, exists := ErrorCodeToHTTPStatus[code]; exists {
		return status
	}
	return http.StatusInternalServerError
}

func ToDeliveryErrorFromService(se *ServiceError) *DeliveryError {
	if se == nil {
		return nil
	}
	return &DeliveryError{
		HTTPStatus: matchStatus(se.Code),
		Errors:     []ErrorItem{NewErrorItem(se.Code, NewCustomError(se.Error.Error()))},
	}
}
