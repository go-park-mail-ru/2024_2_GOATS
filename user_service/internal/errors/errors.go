package errors

import (
	errs "errors"

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

var (
	ErrInvalidEmail          = errs.New("email is incorrect")
	ErrInvalidPassword       = errs.New("password is too short. The minimal len is 8")
	ErrInvalidUsername       = errs.New("username is too short. The minimal len is 6")
	ErrInvalidPasswordsMatch = errs.New("password doesn't match with passwordConfirmation")
	ErrInvalidOldPassword    = errs.New("invalid old password")
	ErrUserNotFound          = errs.New("cannot find user by given params")
	ErrBrokenCookie          = errs.New("broken cookie was given")
	ErrSaveFile              = errs.New("cannot save file")
)

func IsDuplicateError(err error) bool {
	var pqErr *pq.Error
	return errs.As(err, &pqErr) && pqErr.Code == DuplicateErrKey
}
