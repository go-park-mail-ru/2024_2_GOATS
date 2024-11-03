package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/logger"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func TestGetActor_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := &Repo{Database: db}

	actorID := 1
	expectedActor := &models.ActorInfo{
		Id: actorID,
		Person: models.Person{
			Name:    "John",
			Surname: "Doe",
		},
		Biography: "Some biography",
		Birthdate: sql.NullTime{
			Time:  time.Date(1980, time.March, 10, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
		BigPhotoUrl: "some_photo_url",
		Country:     "USA",
		Movies:      nil,
	}

	// actor.FindById
	mock.ExpectQuery(`SELECT actors.id, actors.first_name, actors.second_name, actors.biography, actors.birthdate, actors.big_photo_url, countries.title FROM actors JOIN countries on countries.id = actors.country_id WHERE actors.id = \$1`).
		WithArgs(actorID).RowsWillBeClosed().
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "second_name", "biography", "birthdate", "big_photo_url", "title"}).
			AddRow(actorID, "John", "Doe", "Some biography", expectedActor.Birthdate, "some_photo_url", "USA"))

	// movie.FindByActorId
	mock.ExpectQuery(`SELECT movies.id, movies.title, movies.card_url, movies.rating, movies.release_date, countries.title FROM movies JOIN movie_actors ON movie_actors.movie_id = movies.id JOIN actors ON movie_actors.actor_id = actors.id JOIN countries ON movies.country_id = countries.id WHERE actors.id = \$1`).
		WithArgs(actorID).RowsWillBeClosed().
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "card_url", "rating", "release_date", "country_title"}).
			AddRow(1, "Movie 1", "https://example.com/movie1.jpg", 8.5, time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC), "Russia").
			AddRow(2, "Movie 2", "https://example.com/movie2.jpg", 7.9, time.Date(2019, time.June, 15, 0, 0, 0, 0, time.UTC), "USA"))

	expectedMovies := []*models.MovieShortInfo{
		{Id: 1, Title: "Movie 1", CardUrl: "https://example.com/movie1.jpg", Rating: 8.5, ReleaseDate: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC), Country: "Russia"},
		{Id: 2, Title: "Movie 2", CardUrl: "https://example.com/movie2.jpg", Rating: 7.9, ReleaseDate: time.Date(2019, time.June, 15, 0, 0, 0, 0, time.UTC), Country: "USA"},
	}

	actor, errObj, statusCode := r.GetActor(testContext(), actorID)

	assert.Nil(t, errObj)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, expectedActor.Person.Name, actor.Person.Name)
	assert.Equal(t, expectedMovies, actor.Movies)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetActor_FindByIdError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := &Repo{Database: db}

	actorID := 1
	mock.ExpectQuery(`SELECT actors.id, actors.first_name, actors.second_name, actors.biography, actors.birthdate, actors.big_photo_url, countries.title FROM actors JOIN countries on countries.id = actors.country_id WHERE actors.id = \$1`).
		WithArgs(actorID).
		WillReturnError(fmt.Errorf("some_error"))

	actor, errObj, statusCode := r.GetActor(testContext(), actorID)

	assert.Nil(t, actor)
	assert.NotNil(t, errObj)
	assert.Equal(t, http.StatusUnprocessableEntity, statusCode)
	assert.Equal(t, errors.ErrServerCode, errObj.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetActor_FindByActorIdError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := &Repo{Database: db}

	actorID := 1
	mock.ExpectQuery(`SELECT actors.id, actors.first_name, actors.second_name, actors.biography, actors.birthdate, actors.big_photo_url, countries.title FROM actors JOIN countries on countries.id = actors.country_id WHERE actors.id = \$1`).
		WithArgs(actorID).RowsWillBeClosed().
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "second_name", "biography", "birthdate", "big_photo_url", "title"}).
			AddRow(actorID, "John", "Doe", "Some biography", time.Date(1980, time.March, 10, 0, 0, 0, 0, time.UTC), "some_photo_url", "USA"))

	mock.ExpectQuery(`SELECT movies.id, movies.title, movies.card_url, movies.rating, movies.release_date, countries.title FROM movies JOIN movie_actors ON movie_actors.movie_id = movies.id JOIN actors ON movie_actors.actor_id = actors.id JOIN countries ON movies.country_id = countries.id WHERE actors.id = \$1`).
		WithArgs(actorID).
		WillReturnError(fmt.Errorf("some_error"))

	actor, errObj, statusCode := r.GetActor(testContext(), actorID)

	assert.Nil(t, actor)
	assert.NotNil(t, errObj)
	assert.Equal(t, http.StatusUnprocessableEntity, statusCode)
	assert.Equal(t, errors.ErrServerCode, errObj.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetCollection_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	collId := 1
	expectedCollections := []models.Collection{{
		Id:    collId,
		Title: "Test collection",
		Movies: []*models.MovieShortInfo{
			{
				Id:          1,
				Title:       "test movie",
				CardUrl:     "some_card_url",
				AlbumUrl:    "some_album_url",
				Rating:      7.6,
				ReleaseDate: time.Date(1980, time.March, 10, 0, 0, 0, 0, time.UTC),
				MovieType:   "film",
				Country:     "Russia",
			},
		},
	}}

	r := &Repo{Database: db}

	mock.ExpectQuery(`SELECT collections.id, collections.title, movies.id, movies.title, movies.card_url, movies.album_url, movies.rating, movies.release_date, movies.movie_type, countries.title FROM collections
		JOIN movie_collections ON movie_collections.collection_id = collections.id
		JOIN movies ON movies.id = movie_collections.movie_id
		JOIN countries ON countries.id = movies.country_id`).RowsWillBeClosed().WillReturnRows(sqlmock.NewRows([]string{"id", "title", "movie_id", "movie_title", "card_url", "album_url", "rating", "release_date", "movie_type", "country_title"}).
		AddRow(1, "Test collection", 1, "test movie", "some_card_url", "some_album_url", 7.6, time.Date(1980, time.March, 10, 0, 0, 0, 0, time.UTC), "film", "Russia"))

	colls, errObj, statusCode := r.GetCollection(testContext())
	assert.NotNil(t, colls)
	assert.Nil(t, errObj)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, expectedCollections, colls)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetCollection_ObtainError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT collections.id, collections.title, movies.id, movies.title, movies.card_url, movies.album_url, movies.rating, movies.release_date, movies.movie_type, countries.title FROM collections
	JOIN movie_collections ON movie_collections.collection_id = collections.id
	JOIN movies ON movies.id = movie_collections.movie_id
	JOIN countries ON countries.id = movies.country_id`).WillReturnError(fmt.Errorf("some_error"))

	r := &Repo{Database: db}

	colls, errObj, statusCode := r.GetCollection(testContext())
	assert.Nil(t, colls)
	assert.NotNil(t, errObj)
	assert.Equal(t, http.StatusUnprocessableEntity, statusCode)
	assert.Equal(t, errors.ErrServerCode, errObj.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetMovie_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	movieId := 1
	expectedMovie := &models.MovieInfo{
		Id:               movieId,
		Title:            "test movie",
		ShortDescription: "short desc",
		FullDescription:  "long desc",
		CardUrl:          "card url",
		AlbumUrl:         "album url",
		Rating:           7.6,
		ReleaseDate:      time.Date(1980, time.March, 10, 0, 0, 0, 0, time.UTC),
		VideoUrl:         "video url",
		MovieType:        "film",
		TitleUrl:         "title url",
		Country:          "Russia",
	}

	mock.ExpectQuery(`
		SELECT
			movies.id,
			movies.title,
			movies.short_description,
			movies.long_description,
			movies.card_url,
			movies.album_url,
			movies.rating,
			movies.release_date,
			movies.video_url,
			movies.movie_type,
			movies.title_url,
			directors.first_name,
			directors.second_name,
			countries.title
		FROM movies
		JOIN directors ON directors.id = movies.director_id
		JOIN countries ON countries.id = movies.country_id
		WHERE movies.id = \$1
	`).WithArgs(movieId).RowsWillBeClosed().WillReturnRows(sqlmock.NewRows(
		[]string{
			"id",
			"title",
			"short_description",
			"long_description",
			"card_url",
			"album_url",
			"rating",
			"release_date",
			"video_url",
			"movie_type",
			"title_url",
			"directors_name",
			"directors_surname",
			"country_title",
		}).AddRow(1, "test movie", "short desc", "long desc", "card url", "album url", 7.6, time.Date(1980, time.March, 10, 0, 0, 0, 0, time.UTC), "video url", "film", "title url", "Test", "Tester", "Russia"))

	r := &Repo{Database: db}

	movie, errObj, statusCode := r.GetMovie(testContext(), movieId)
	expectedMovie.Director = movie.Director

	assert.NotNil(t, movie)
	assert.Nil(t, errObj)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, expectedMovie, movie)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetMovie_FindByIdError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	movieId := 1

	mock.ExpectQuery(`
		SELECT
			movies.id,
			movies.title,
			movies.short_description,
			movies.long_description,
			movies.card_url,
			movies.album_url,
			movies.rating,
			movies.release_date,
			movies.video_url,
			movies.movie_type,
			movies.title_url,
			directors.first_name,
			directors.second_name,
			countries.title
		FROM movies
		JOIN directors ON directors.id = movies.director_id
		JOIN countries ON countries.id = movies.country_id
		WHERE movies.id = \$1
	`).WithArgs(movieId).WillReturnError(fmt.Errorf("some error"))

	r := &Repo{Database: db}

	movie, errObj, statusCode := r.GetMovie(testContext(), movieId)

	assert.Nil(t, movie)
	assert.NotNil(t, errObj)
	assert.Equal(t, http.StatusUnprocessableEntity, statusCode)
	assert.Equal(t, errors.ErrServerCode, errObj.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetMovieActors_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	movieId := 1
	expectedActors := []*models.ActorInfo{{
		Id: 1,
		Person: models.Person{
			Name:    "Test",
			Surname: "Tester",
		},
		Biography:     "some bio",
		SmallPhotoUrl: "some_small_photo_link",
	}}

	mock.ExpectQuery(`
		SELECT
			actors.id,
			actors.first_name,
			actors.second_name,
			actors.biography,
			actors.small_photo_url
		FROM actors
		JOIN movie_actors on movie_actors.actor_id = actors.id
		JOIN movies on movie_actors.movie_id = movies.id
		WHERE movies.id = \$1
	`).WithArgs(movieId).RowsWillBeClosed().WillReturnRows(sqlmock.NewRows(
		[]string{
			"id",
			"name",
			"surname",
			"biography",
			"small_photo_url",
		}).AddRow(1, "Test", "Tester", "some bio", "some_small_photo_link"))

	r := &Repo{Database: db}

	actors, errObj, statusCode := r.GetMovieActors(testContext(), movieId)

	assert.NotNil(t, actors)
	assert.Nil(t, errObj)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, expectedActors, actors)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetMovieActors_DbError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	movieId := 1

	mock.ExpectQuery(`
		SELECT
			actors.id,
			actors.first_name,
			actors.second_name,
			actors.biography,
			actors.small_photo_url
		FROM actors
		JOIN movie_actors on movie_actors.actor_id = actors.id
		JOIN movies on movie_actors.movie_id = movies.id
		WHERE movies.id = \$1
	`).WithArgs(movieId).WillReturnError(fmt.Errorf("some error"))

	r := &Repo{Database: db}

	actors, errObj, statusCode := r.GetMovieActors(testContext(), movieId)

	assert.Nil(t, actors)
	assert.NotNil(t, errObj)
	assert.Equal(t, http.StatusUnprocessableEntity, statusCode)
	assert.Equal(t, errors.ErrServerCode, errObj.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func testContext() context.Context {
	ctx := context.WithValue(context.Background(), "request-id", "some-request-id")
	return config.WrapLoggerContext(ctx, logger.NewLogger())
}
