package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/password"
)

var usrEmail = "test@mail.ru"
var usrUsername = "mr tester"
var pass = "test_password"

func TestCreateUser_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)

	usrID := 1
	usrPassword, err := password.HashAndSalt(context.Background(), pass)
	assert.NoError(t, err)

	expectedUsr := &dto.RepoUser{
		ID:       uint64(usrID),
		Email:    usrEmail,
		Password: usrPassword,
		Username: usrUsername,
	}

	regData := &dto.RepoCreateData{
		Email:                usrEmail,
		Username:             usrUsername,
		Password:             pass,
		PasswordConfirmation: pass,
	}

	mock.ExpectPrepare(`INSERT INTO users \(email, username, password_hash\) VALUES \(\$1, \$2, \$3\) RETURNING id, email`).
		ExpectQuery().
		WithArgs(usrEmail, usrUsername, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email"}).
			AddRow(usrID, usrEmail))

	usr, errObj := r.CreateUser(context.Background(), regData)

	assert.Nil(t, errObj)
	assert.Equal(t, expectedUsr.Email, usr.Email)
	assert.Equal(t, expectedUsr.ID, usr.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateUser_DbError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)

	regData := &dto.RepoCreateData{
		Email:                usrEmail,
		Username:             usrUsername,
		Password:             pass,
		PasswordConfirmation: pass,
	}

	mock.ExpectPrepare(`INSERT INTO users \(email, username, password_hash\) VALUES \(\$1, \$2, \$3\) RETURNING id, email`).
		ExpectQuery().
		WithArgs(usrEmail, usrUsername, sqlmock.AnyArg()).WillReturnError(fmt.Errorf("some_error"))

	usr, errObj := r.CreateUser(context.Background(), regData)

	assert.NotNil(t, errObj)
	assert.Nil(t, usr)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdatePassword_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)

	usrID := 1

	mock.ExpectPrepare(`UPDATE users SET password_hash = \$1, updated_at = \$2 WHERE id = \$3`).
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), usrID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	errObj := r.UpdatePassword(context.Background(), uint64(usrID), pass)

	assert.Nil(t, errObj)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdatePassword_DbError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)

	usrID := 1

	mock.ExpectPrepare(`UPDATE users SET password_hash = \$1, updated_at = \$2 WHERE id = \$3`).
		ExpectExec().
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), usrID).
		WillReturnError(fmt.Errorf("some_error"))

	errObj := r.UpdatePassword(context.Background(), uint64(usrID), pass)

	assert.NotNil(t, errObj)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserByEmail_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)

	usrID := 1
	passHash, err := password.HashAndSalt(context.Background(), "test_password")
	assert.NoError(t, err)

	mock.ExpectPrepare(`SELECT id, email, username, password_hash FROM users WHERE email = \$1`).
		ExpectQuery().
		WithArgs(usrEmail).WillReturnRows(sqlmock.NewRows([]string{"id", "email", "username", "password_hash"}).
		AddRow(usrID, usrEmail, usrUsername, passHash))

	usr, errObj := r.UserByEmail(context.Background(), usrEmail)

	assert.NotNil(t, usr)
	assert.Nil(t, errObj)
	assert.Equal(t, uint64(usrID), usr.ID)
	assert.Equal(t, usrEmail, usr.Email)
	assert.Equal(t, usrUsername, usr.Username)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserByEmail_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)

	mock.ExpectPrepare(`SELECT id, email, username, password_hash FROM users WHERE email = \$1`).
		ExpectQuery().
		WithArgs(usrEmail).WillReturnError(sql.ErrNoRows)

	usr, errObj := r.UserByEmail(context.Background(), usrEmail)

	assert.Nil(t, usr)
	assert.NotNil(t, errObj)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserByID_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)

	usrID := 1
	passHash, err := password.HashAndSalt(context.Background(), "test_password")
	assert.NoError(t, err)

	mock.ExpectPrepare(`SELECT id, email, username, password_hash, avatar_url FROM users WHERE id = \$1`).
		ExpectQuery().
		WithArgs(usrID).WillReturnRows(sqlmock.NewRows([]string{"id", "email", "username", "password_hash", "avatar_url"}).
		AddRow(usrID, usrEmail, usrUsername, passHash, "test.png"))

	mock.ExpectPrepare(`SELECT status, expiration_date FROM subscriptions WHERE user_id = \$1 and expiration_date > \$2 and status = \$3`).
		ExpectQuery().
		WithArgs(usrID, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"status", "expiration_date"}))

	usr, errObj := r.UserByID(context.Background(), uint64(usrID))

	assert.NotNil(t, usr)
	assert.Nil(t, errObj)
	assert.Equal(t, uint64(usrID), usr.ID)
	assert.Equal(t, usrEmail, usr.Email)
	assert.Equal(t, usrUsername, usr.Username)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)

	usrID := 1

	mock.ExpectPrepare(`SELECT id, email, username, password_hash, avatar_url FROM users WHERE id = \$1`).
		ExpectQuery().
		WithArgs(usrID).WillReturnError(sql.ErrNoRows)

	usr, errObj := r.UserByID(context.Background(), uint64(usrID))

	assert.Nil(t, usr)
	assert.NotNil(t, errObj)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateProfileData_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)
	profileData := &dto.RepoUser{
		ID:        1,
		Email:     "test@mail.ru",
		Username:  "testuser",
		AvatarURL: "some_avatar_url",
	}

	mock.ExpectPrepare(`UPDATE users SET email = \$1, username = \$2, avatar_url = \$3, updated_at = \$4 WHERE id = \$5`).
		ExpectExec().
		WithArgs(profileData.Email, profileData.Username, profileData.AvatarURL, sqlmock.AnyArg(), profileData.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	errObj := r.UpdateProfileData(context.Background(), profileData)

	assert.Nil(t, errObj)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateProfileData_DbError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)
	profileData := &dto.RepoUser{
		ID:        1,
		Email:     "test@mail.ru",
		Username:  "testuser",
		AvatarURL: "some_avatar_url",
	}

	mock.ExpectPrepare(`UPDATE users SET email = \$1, username = \$2, avatar_url = \$3, updated_at = \$4 WHERE id = \$5`).
		ExpectExec().
		WithArgs(profileData.Email, profileData.Username, profileData.AvatarURL, sqlmock.AnyArg(), profileData.ID).
		WillReturnError(fmt.Errorf("some database error"))

	errObj := r.UpdateProfileData(context.Background(), profileData)

	assert.NotNil(t, errObj)
	assert.Equal(t, "update_profile_error: postgres: error while updating user profile - some database error", errObj.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCheckFavorite_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)
	favData := &dto.RepoFavorite{
		UserID:  1,
		MovieID: 1,
	}

	mockRows := sqlmock.NewRows([]string{"movie_id"}).AddRow(1)
	mock.ExpectPrepare(`SELECT movie_id FROM favorites WHERE user_id = \$1 and movie_id = \$2`).
		ExpectQuery().
		WithArgs(favData.UserID, favData.MovieID).
		WillReturnRows(mockRows)

	present, errObj := r.CheckFavorite(context.Background(), favData)

	assert.Nil(t, errObj)
	assert.True(t, present)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCheckFavorite_FalseSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)
	favData := &dto.RepoFavorite{
		UserID:  1,
		MovieID: 1,
	}

	mockRows := sqlmock.NewRows([]string{"movie_id"})
	mock.ExpectPrepare(`SELECT movie_id FROM favorites WHERE user_id = \$1 and movie_id = \$2`).
		ExpectQuery().
		WithArgs(favData.UserID, favData.MovieID).
		WillReturnRows(mockRows)

	present, errObj := r.CheckFavorite(context.Background(), favData)

	assert.Nil(t, errObj)
	assert.False(t, present)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCheckFavorite_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)
	favData := &dto.RepoFavorite{
		UserID:  1,
		MovieID: 1,
	}

	mock.ExpectPrepare(`SELECT movie_id FROM favorites WHERE user_id = \$1 and movie_id = \$2`).
		ExpectQuery().
		WithArgs(favData.UserID, favData.MovieID).
		WillReturnError(errors.New("some_database_error"))

	present, errObj := r.CheckFavorite(context.Background(), favData)

	assert.NotNil(t, errObj)
	assert.False(t, present)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetFavorites_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)
	expectedFavs := []uint64{1, 2}
	usrID := 1

	mockRows := sqlmock.NewRows([]string{"movie_id"}).AddRow(1).AddRow(2)
	mock.ExpectPrepare(`SELECT movie_id FROM favorites WHERE user_id = \$1`).
		ExpectQuery().
		WithArgs(usrID).
		WillReturnRows(mockRows)

	favs, errObj := r.GetFavorites(context.Background(), uint64(usrID))

	assert.Nil(t, errObj)
	assert.Equal(t, expectedFavs, favs)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetFavorites_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)
	usrID := 1

	mock.ExpectPrepare(`SELECT movie_id FROM favorites WHERE user_id = \$1`).
		ExpectQuery().
		WithArgs(usrID).
		WillReturnError(errors.New("some_database_error"))

	favs, errObj := r.GetFavorites(context.Background(), uint64(usrID))

	assert.Nil(t, favs)
	assert.NotNil(t, errObj)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSetFavorite_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)
	favData := &dto.RepoFavorite{
		UserID:  1,
		MovieID: 1,
	}

	mock.ExpectPrepare(`INSERT INTO favorites \(user_id, movie_id\) VALUES \(\$1, \$2\)`).
		ExpectQuery().
		WithArgs(favData.UserID, favData.MovieID).
		WillReturnRows(sqlmock.NewRows(nil))

	err = r.SetFavorite(context.Background(), favData)

	assert.Nil(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSetFavorite_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)
	favData := &dto.RepoFavorite{
		UserID:  1,
		MovieID: 1,
	}

	mock.ExpectPrepare(`INSERT INTO favorites \(user_id, movie_id\) VALUES \(\$1, \$2\)`).
		ExpectQuery().
		WithArgs(favData.UserID, favData.MovieID).
		WillReturnError(errors.New("some_database_error"))

	err = r.SetFavorite(context.Background(), favData)

	assert.NotNil(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestResetFavorite_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)
	favData := &dto.RepoFavorite{
		UserID:  1,
		MovieID: 1,
	}

	mock.ExpectPrepare(`DELETE FROM favorites WHERE user_id = \$1 and movie_id = \$2`).
		ExpectQuery().
		WithArgs(favData.UserID, favData.MovieID).
		WillReturnRows(sqlmock.NewRows(nil))

	err = r.ResetFavorite(context.Background(), favData)

	assert.Nil(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestResetFavorite_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	r := NewUserRepository(db)
	favData := &dto.RepoFavorite{
		UserID:  1,
		MovieID: 1,
	}

	mock.ExpectPrepare(`DELETE FROM favorites WHERE user_id = \$1 and movie_id = \$2`).
		ExpectQuery().
		WithArgs(favData.UserID, favData.MovieID).
		WillReturnError(errors.New("some_database_error"))

	err = r.ResetFavorite(context.Background(), favData)

	assert.NotNil(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
