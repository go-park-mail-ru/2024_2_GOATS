package repository

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/password"
)

func TestCreateUser_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewUserRepository(db)

	usrID := 1
	usrEmail := "test@mail.ru"
	usrUsername := "mr tester"
	pass := "test_password"
	usrPassword, err := password.HashAndSalt(context.Background(), pass)

	expectedUsr := &models.User{
		ID:       usrID,
		Email:    usrEmail,
		Password: usrPassword,
		Username: usrUsername,
	}

	regData := &dto.RepoRegisterData{
		Email:                usrEmail,
		Username:             usrUsername,
		Password:             pass,
		PasswordConfirmation: pass,
	}

	mock.ExpectQuery(`INSERT INTO users \(email, username, password_hash\) VALUES \(\$1, \$2, \$3\) RETURNING id, email`).
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
	defer db.Close()

	r := NewUserRepository(db)

	usrEmail := "test@mail.ru"
	usrUsername := "mr tester"
	pass := "test_password"

	regData := &dto.RepoRegisterData{
		Email:                usrEmail,
		Username:             usrUsername,
		Password:             pass,
		PasswordConfirmation: pass,
	}

	mock.ExpectQuery(`INSERT INTO users \(email, username, password_hash\) VALUES \(\$1, \$2, \$3\) RETURNING id, email`).
		WithArgs(usrEmail, usrUsername, sqlmock.AnyArg()).WillReturnError(fmt.Errorf("some_error"))

	usr, errObj := r.CreateUser(context.Background(), regData)

	assert.NotNil(t, errObj)
	assert.Nil(t, usr)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdatePassword_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewUserRepository(db)

	usrID := 1
	pass := "test_password"

	mock.ExpectExec(`UPDATE users SET password_hash = \$1, updated_at = \$2 WHERE id = \$3`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), usrID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	errObj := r.UpdatePassword(context.Background(), usrID, pass)

	assert.Nil(t, errObj)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdatePassword_DbError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewUserRepository(db)

	usrID := 1
	pass := "test_password"

	mock.ExpectExec(`UPDATE users SET password_hash = \$1, updated_at = \$2 WHERE id = \$3`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), usrID).
		WillReturnError(fmt.Errorf("some_error"))

	errObj := r.UpdatePassword(context.Background(), usrID, pass)

	assert.NotNil(t, errObj)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserByEmail_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewUserRepository(db)

	usrID := 1
	usrEmail := "test@mail.ru"
	usrUsername := "mr tester"
	passHash, err := password.HashAndSalt(context.Background(), "test_password")
	assert.NoError(t, err)

	mock.ExpectQuery(`SELECT id, email, username, password_hash FROM USERS WHERE email = \$1`).
		WithArgs(usrEmail).WillReturnRows(sqlmock.NewRows([]string{"id", "email", "username", "password_hash"}).
		AddRow(usrID, usrEmail, usrUsername, passHash))

	usr, errObj := r.UserByEmail(context.Background(), usrEmail)

	assert.NotNil(t, usr)
	assert.Nil(t, errObj)
	assert.Equal(t, usrID, usr.ID)
	assert.Equal(t, usrEmail, usr.Email)
	assert.Equal(t, usrUsername, usr.Username)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserByEmail_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewUserRepository(db)

	usrEmail := "test@mail.ru"

	mock.ExpectQuery(`SELECT id, email, username, password_hash FROM USERS WHERE email = \$1`).
		WithArgs(usrEmail).WillReturnError(sql.ErrNoRows)

	usr, errObj := r.UserByEmail(context.Background(), usrEmail)

	assert.Nil(t, usr)
	assert.NotNil(t, errObj)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserByID_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewUserRepository(db)

	usrID := 1
	usrEmail := "test@mail.ru"
	usrUsername := "mr tester"
	passHash, err := password.HashAndSalt(context.Background(), "test_password")
	assert.NoError(t, err)

	mock.ExpectQuery(`SELECT id, email, username, password_hash, avatar_url FROM USERS WHERE id = \$1`).
		WithArgs(usrID).WillReturnRows(sqlmock.NewRows([]string{"id", "email", "username", "password_hash", "avatar_url"}).
		AddRow(usrID, usrEmail, usrUsername, passHash, "test.png"))

	usr, errObj := r.UserByID(context.Background(), usrID)

	assert.NotNil(t, usr)
	assert.Nil(t, errObj)
	assert.Equal(t, usrID, usr.ID)
	assert.Equal(t, usrEmail, usr.Email)
	assert.Equal(t, usrUsername, usr.Username)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewUserRepository(db)

	usrID := 1

	mock.ExpectQuery(`SELECT id, email, username, password_hash, avatar_url FROM USERS WHERE id = \$1`).
		WithArgs(usrID).WillReturnError(sql.ErrNoRows)

	usr, errObj := r.UserByID(context.Background(), usrID)

	assert.Nil(t, usr)
	assert.NotNil(t, errObj)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateProfileData_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewUserRepository(db)
	profileData := &dto.RepoUser{
		ID:        1,
		Email:     "test@mail.ru",
		Username:  "testuser",
		AvatarURL: "some_avatar_url",
	}

	mock.ExpectExec(`UPDATE users SET email = \$1, username = \$2, avatar_url = \$3, updated_at = \$4 WHERE id = \$5`).
		WithArgs(profileData.Email, profileData.Username, profileData.AvatarURL, sqlmock.AnyArg(), profileData.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	errObj := r.UpdateProfileData(context.Background(), profileData)

	assert.Nil(t, errObj)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateProfileData_DbError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewUserRepository(db)
	profileData := &dto.RepoUser{
		ID:        1,
		Email:     "test@mail.ru",
		Username:  "testuser",
		AvatarURL: "some_avatar_url",
	}

	mock.ExpectExec(`UPDATE users SET email = \$1, username = \$2, avatar_url = \$3, updated_at = \$4 WHERE id = \$5`).
		WithArgs(profileData.Email, profileData.Username, profileData.AvatarURL, sqlmock.AnyArg(), profileData.ID).
		WillReturnError(fmt.Errorf("some database error"))

	errObj := r.UpdateProfileData(context.Background(), profileData)

	assert.NotNil(t, errObj)
	assert.Equal(t, "update_profile_error", errObj.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}
