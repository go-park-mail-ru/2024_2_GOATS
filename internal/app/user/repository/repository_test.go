package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/logger"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/password"
)

func TestCreateUser_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewRepository(db)

	usrId := 1
	usrEmail := "test@mail.ru"
	usrUsername := "mr tester"
	pass := "test_password"
	usrPassword, err := password.HashAndSalt(pass)
	assert.NoError(t, err)

	expectedUsr := &models.User{
		Id:       usrId,
		Email:    usrEmail,
		Password: usrPassword,
		Username: usrUsername,
	}

	regData := &models.RegisterData{
		Email:                usrEmail,
		Username:             usrUsername,
		Password:             pass,
		PasswordConfirmation: pass,
	}

	mock.ExpectQuery(`INSERT INTO users \(email, username, password_hash\) VALUES \(\$1, \$2, \$3\) RETURNING id, email`).
		WithArgs(usrEmail, usrUsername, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email"}).
			AddRow(usrId, usrEmail))

	usr, errObj, statusCode := r.CreateUser(testContext(), regData)

	assert.Nil(t, errObj)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, expectedUsr.Email, usr.Email)
	assert.Equal(t, expectedUsr.Id, usr.Id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateUser_DbError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewRepository(db)

	usrEmail := "test@mail.ru"
	usrUsername := "mr tester"
	pass := "test_password"
	assert.NoError(t, err)

	regData := &models.RegisterData{
		Email:                usrEmail,
		Username:             usrUsername,
		Password:             pass,
		PasswordConfirmation: pass,
	}

	mock.ExpectQuery(`INSERT INTO users \(email, username, password_hash\) VALUES \(\$1, \$2, \$3\) RETURNING id, email`).
		WithArgs(usrEmail, usrUsername, sqlmock.AnyArg()).WillReturnError(fmt.Errorf("some_error"))

	usr, errObj, statusCode := r.CreateUser(testContext(), regData)

	assert.NotNil(t, errObj)
	assert.Nil(t, usr)
	assert.Equal(t, http.StatusConflict, statusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdatePassword_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewRepository(db)

	usrId := 1
	pass := "test_password"

	mock.ExpectExec(`UPDATE users SET password_hash = \$1, updated_at = \$2 WHERE id = \$3`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), usrId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	errObj, statusCode := r.UpdatePassword(testContext(), usrId, pass)

	assert.Nil(t, errObj)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdatePassword_DbError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewRepository(db)

	usrId := 1
	pass := "test_password"

	mock.ExpectExec(`UPDATE users SET password_hash = \$1, updated_at = \$2 WHERE id = \$3`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), usrId).
		WillReturnError(fmt.Errorf("some_error"))

	errObj, statusCode := r.UpdatePassword(testContext(), usrId, pass)

	assert.NotNil(t, errObj)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserByEmail_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewRepository(db)

	usrId := 1
	usrEmail := "test@mail.ru"
	usrUsername := "mr tester"
	passHash, err := password.HashAndSalt("test_password")
	assert.NoError(t, err, "failed to hash pass")

	mock.ExpectQuery(`SELECT id, email, username, password_hash FROM USERS WHERE email = \$1`).
		WithArgs(usrEmail).WillReturnRows(sqlmock.NewRows([]string{"id", "email", "username", "password_hash"}).
		AddRow(usrId, usrEmail, usrUsername, passHash))

	usr, errObj, statusCode := r.UserByEmail(testContext(), usrEmail)

	assert.NotNil(t, usr)
	assert.Nil(t, errObj)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, usrId, usr.Id)
	assert.Equal(t, usrEmail, usr.Email)
	assert.Equal(t, usrUsername, usr.Username)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserByEmail_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewRepository(db)

	usrEmail := "test@mail.ru"

	mock.ExpectQuery(`SELECT id, email, username, password_hash FROM USERS WHERE email = \$1`).
		WithArgs(usrEmail).WillReturnError(sql.ErrNoRows)

	usr, errObj, statusCode := r.UserByEmail(testContext(), usrEmail)

	assert.Nil(t, usr)
	assert.NotNil(t, errObj)
	assert.Equal(t, http.StatusNotFound, statusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserById_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewRepository(db)

	usrId := 1
	usrEmail := "test@mail.ru"
	usrUsername := "mr tester"
	passHash, err := password.HashAndSalt("test_password")
	assert.NoError(t, err, "failed to hash pass")

	mock.ExpectQuery(`SELECT id, email, username, password_hash, avatar_url FROM USERS WHERE id = \$1`).
		WithArgs(usrId).WillReturnRows(sqlmock.NewRows([]string{"id", "email", "username", "password_hash", "avatar_url"}).
		AddRow(usrId, usrEmail, usrUsername, passHash, "test.png"))

	usr, errObj, statusCode := r.UserById(testContext(), usrId)

	assert.NotNil(t, usr)
	assert.Nil(t, errObj)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, usrId, usr.Id)
	assert.Equal(t, usrEmail, usr.Email)
	assert.Equal(t, usrUsername, usr.Username)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserById_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewRepository(db)

	usrId := 1

	mock.ExpectQuery(`SELECT id, email, username, password_hash, avatar_url FROM USERS WHERE id = \$1`).
		WithArgs(usrId).WillReturnError(sql.ErrNoRows)

	usr, errObj, statusCode := r.UserById(testContext(), usrId)

	assert.Nil(t, usr)
	assert.NotNil(t, errObj)
	assert.Equal(t, http.StatusNotFound, statusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateProfileData_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewRepository(db)
	profileData := &models.User{
		Id:        1,
		Email:     "test@mail.ru",
		Username:  "testuser",
		AvatarUrl: "some_avatar_url",
	}

	mock.ExpectExec(`UPDATE users SET email = \$1, username = \$2, avatar_url = \$3, updated_at = \$4 WHERE id = \$5`).
		WithArgs(profileData.Email, profileData.Username, profileData.AvatarUrl, sqlmock.AnyArg(), profileData.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	errObj, status := r.UpdateProfileData(testContext(), profileData)

	assert.Nil(t, errObj)
	assert.Equal(t, http.StatusOK, status)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateProfileData_DbError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewRepository(db)
	profileData := &models.User{
		Id:        1,
		Email:     "test@mail.ru",
		Username:  "testuser",
		AvatarUrl: "some_avatar_url",
	}

	mock.ExpectExec(`UPDATE users SET email = \$1, username = \$2, avatar_url = \$3, updated_at = \$4 WHERE id = \$5`).
		WithArgs(profileData.Email, profileData.Username, profileData.AvatarUrl, sqlmock.AnyArg(), profileData.Id).
		WillReturnError(fmt.Errorf("some database error"))

	errObj, status := r.UpdateProfileData(testContext(), profileData)

	assert.NotNil(t, errObj)
	assert.Equal(t, "update_profile_error", errObj.Code)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func testContext() context.Context {
	ctx := context.WithValue(context.Background(), "request-id", "some-request-id")
	return config.WrapLoggerContext(ctx, logger.NewLogger())
}
