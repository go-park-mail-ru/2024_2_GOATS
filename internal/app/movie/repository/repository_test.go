package repository

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func TestGetActor_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewMovieRepository(db)

	actorID := 1
	expectedActor := &models.ActorInfo{
		ID: actorID,
		Person: models.Person{
			Name:    "John",
			Surname: "Doe",
		},
		Biography: "Some biography",
		Birthdate: sql.NullString{
			String: "1980-03-10",
			Valid:  true,
		},
		BigPhotoURL: "some_photo_url",
		Country:     "USA",
		Movies:      nil,
	}

	// actor.FindByID
	mock.ExpectQuery(`
		SELECT
			actors.id,
			actors.first_name,
			actors.second_name,
			actors.biography,
			actors.birthdate,
			actors.big_photo_url,
			countries.title
		FROM actors
		JOIN countries on countries.id = actors.country_id
		WHERE actors.id = \$1
	`).
		WithArgs(actorID).RowsWillBeClosed().
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "second_name", "biography", "birthdate", "big_photo_url", "title"}).
			AddRow(actorID, "John", "Doe", "Some biography", expectedActor.Birthdate, "some_photo_url", "USA"))

	// movie.FindByActorID
	mock.ExpectQuery(`
		SELECT
			movies.id,
			movies.title,
			movies.card_url,
			movies.rating,
			movies.release_date,
			countries.title
		FROM movies
		JOIN movie_actors ON movie_actors.movie_id = movies.id
		JOIN actors ON movie_actors.actor_id = actors.id
		JOIN countries ON movies.country_id = countries.id
		WHERE actors.id = \$1
	`).
		WithArgs(actorID).RowsWillBeClosed().
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "card_url", "rating", "release_date", "country_title"}).
			AddRow(1, "Movie 1", "https://example.com/movie1.jpg", 8.5, "2020-01-01", "Russia").
			AddRow(2, "Movie 2", "https://example.com/movie2.jpg", 7.9, "2019-06-15", "USA"))

	expectedMovies := []*models.MovieShortInfo{
		{ID: 1, Title: "Movie 1", CardURL: "https://example.com/movie1.jpg", Rating: 8.5, ReleaseDate: "2020-01-01", Country: "Russia"},
		{ID: 2, Title: "Movie 2", CardURL: "https://example.com/movie2.jpg", Rating: 7.9, ReleaseDate: "2019-06-15", Country: "USA"},
	}

	actor, errObj := r.GetActor(context.Background(), actorID)

	assert.Nil(t, errObj)
	assert.Equal(t, expectedActor.Person.Name, actor.Person.Name)
	assert.Equal(t, expectedMovies, actor.Movies)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetActor_FindByIDError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewMovieRepository(db)

	actorID := 1
	mock.ExpectQuery(`
		SELECT
			actors.id,
			actors.first_name,
			actors.second_name,
			actors.biography,
			actors.birthdate,
			actors.big_photo_url,
			countries.title
		FROM actors
		JOIN countries on countries.id = actors.country_id
		WHERE actors.id = \$1
	`).
		WithArgs(actorID).
		WillReturnError(fmt.Errorf("some_error"))

	actor, errObj := r.GetActor(context.Background(), actorID)

	assert.Nil(t, actor)
	assert.NotNil(t, errObj)
	assert.Equal(t, errors.ErrServerCode, errObj.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetActor_FindByActorIDError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewMovieRepository(db)

	actorID := 1
	mock.ExpectQuery(`
		SELECT
			actors.id,
			actors.first_name,
			actors.second_name,
			actors.biography,
			actors.birthdate,
			actors.big_photo_url,
			countries.title
		FROM actors
		JOIN countries on countries.id = actors.country_id
		WHERE actors.id = \$1
	`).
		WithArgs(actorID).RowsWillBeClosed().
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "second_name", "biography", "birthdate", "big_photo_url", "title"}).
			AddRow(actorID, "John", "Doe", "Some biography", "1980-03-10", "some_photo_url", "USA"))

	mock.ExpectQuery(`
		SELECT
			movies.id,
			movies.title,
			movies.card_url,
			movies.rating,
			movies.release_date,
			countries.title
		FROM movies
		JOIN movie_actors ON movie_actors.movie_id = movies.id
		JOIN actors ON movie_actors.actor_id = actors.id
		JOIN countries ON movies.country_id = countries.id
		WHERE actors.id = \$1
	`).
		WithArgs(actorID).
		WillReturnError(fmt.Errorf("some_error"))

	actor, errObj := r.GetActor(context.Background(), actorID)

	assert.Nil(t, actor)
	assert.NotNil(t, errObj)
	assert.Equal(t, errors.ErrServerCode, errObj.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetCollection_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	collID := 1
	expectedCollections := []models.Collection{{
		ID:    collID,
		Title: "Test collection",
		Movies: []*models.MovieShortInfo{
			{
				ID:          1,
				Title:       "test movie",
				CardURL:     "some_card_url",
				AlbumURL:    "some_album_url",
				Rating:      7.6,
				ReleaseDate: "1980-03-10",
				MovieType:   "film",
				Country:     "Russia",
			},
		},
	}}

	r := NewMovieRepository(db)

	mock.ExpectQuery(`
		SELECT
			collections.id,
			collections.title,
			movies.id,
			movies.title,
			movies.card_url,
			movies.album_url,
			movies.rating,
			movies.release_date,
			movies.movie_type,
			countries.title
		FROM collections
		JOIN movie_collections ON movie_collections.collection_id = collections.id
		JOIN movies ON movies.id = movie_collections.movie_id
		JOIN countries ON countries.id = movies.country_id
	`).RowsWillBeClosed().WillReturnRows(
		sqlmock.NewRows([]string{
			"id",
			"title",
			"movie_id",
			"movie_title",
			"card_url",
			"album_url",
			"rating",
			"release_date",
			"movie_type",
			"country_title",
		}).AddRow(
			1,
			"Test collection",
			1,
			"test movie",
			"some_card_url",
			"some_album_url",
			7.6,
			"1980-03-10",
			"film",
			"Russia",
		),
	)

	colls, errObj := r.GetCollection(context.Background(), "")
	assert.NotNil(t, colls)
	assert.Nil(t, errObj)
	assert.Equal(t, expectedCollections, colls)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetCollection_ObtainError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`
		SELECT
			collections.id,
			collections.title,
			movies.id,
			movies.title,
			movies.card_url,
			movies.album_url,
			movies.rating,
			movies.release_date,
			movies.movie_type,
			countries.title
		FROM collections
		JOIN movie_collections ON movie_collections.collection_id = collections.id
		JOIN movies ON movies.id = movie_collections.movie_id
		JOIN countries ON countries.id = movies.country_id
	`).WillReturnError(fmt.Errorf("some_error"))

	r := NewMovieRepository(db)

	colls, errObj := r.GetCollection(context.Background(), "")
	assert.Nil(t, colls)
	assert.NotNil(t, errObj)
	assert.Equal(t, errors.ErrServerCode, errObj.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetMovie_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	movieID := 1
	expectedMovie := &models.MovieInfo{
		ID:               movieID,
		Title:            "test movie",
		ShortDescription: "short desc",
		FullDescription:  "long desc",
		CardURL:          "card url",
		AlbumURL:         "album url",
		Rating:           7.6,
		ReleaseDate:      "1980-03-10",
		VideoURL:         "video url",
		MovieType:        "film",
		TitleURL:         "title url",
		Country:          "Russia",
		Director: &models.DirectorInfo{
			Person: models.Person{
				Name:    "Test",
				Surname: "Tester",
			},
		},
		Seasons: []*models.Season{},
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
			countries.title,
			episodes.id,
   		episodes.title,
   		episodes.description,
   		seasons.season_number,
   		episodes.episode_number,
   		episodes.release_date,
   		episodes.rating,
   		episodes.preview_url,
   		episodes.video_url
		FROM movies
		JOIN directors ON directors.id = movies.director_id
		JOIN countries ON countries.id = movies.country_id
		LEFT JOIN seasons ON seasons.movie_id = movies.id AND movies.movie_type = 'serial'
		LEFT JOIN episodes ON seasons.id = episodes.season_id AND movies.movie_type = 'serial'
		WHERE movies.id = \$1
	`).WithArgs(movieID).RowsWillBeClosed().WillReturnRows(sqlmock.NewRows(
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
			"episode_id",
			"episode_title",
			"episode_description",
			"season_number",
			"episode_number",
			"episode_release_date",
			"episode_rating",
			"episode_preview_url",
			"episode_video_url",
		}).AddRow(
		1,
		"test movie",
		"short desc",
		"long desc",
		"card url",
		"album url",
		7.6,
		"1980-03-10",
		"video url",
		"film",
		"title url",
		"Test",
		"Tester",
		"Russia",
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	),
	)

	r := NewMovieRepository(db)

	movie, errObj := r.GetMovie(context.Background(), movieID)

	assert.NotNil(t, movie)
	assert.Nil(t, errObj)
	assert.Equal(t, expectedMovie, movie)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetMovie_FindByIDError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	movieID := 1

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
			countries.title,
			episodes.id,
   		episodes.title,
   		episodes.description,
   		seasons.season_number,
   		episodes.episode_number,
   		episodes.release_date,
   		episodes.rating,
   		episodes.preview_url,
   		episodes.video_url
		FROM movies
		JOIN directors ON directors.id = movies.director_id
		JOIN countries ON countries.id = movies.country_id
		LEFT JOIN seasons ON seasons.movie_id = movies.id AND movies.movie_type = 'serial'
		LEFT JOIN episodes ON seasons.id = episodes.season_id AND movies.movie_type = 'serial'
		WHERE movies.id = \$1
	`).WithArgs(movieID).WillReturnError(fmt.Errorf("some error"))

	r := NewMovieRepository(db)

	movie, errObj := r.GetMovie(context.Background(), movieID)

	assert.Nil(t, movie)
	assert.NotNil(t, errObj)
	assert.Equal(t, errors.ErrServerCode, errObj.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetMovieActors_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	movieID := 1
	expectedActors := []*models.ActorInfo{{
		ID: 1,
		Person: models.Person{
			Name:    "Test",
			Surname: "Tester",
		},
		Biography:     "some bio",
		SmallPhotoURL: "some_small_photo_link",
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
	`).WithArgs(movieID).RowsWillBeClosed().WillReturnRows(sqlmock.NewRows(
		[]string{
			"id",
			"name",
			"surname",
			"biography",
			"small_photo_url",
		}).AddRow(1, "Test", "Tester", "some bio", "some_small_photo_link"))

	r := NewMovieRepository(db)

	actors, errObj := r.GetMovieActors(context.Background(), movieID)

	assert.NotNil(t, actors)
	assert.Nil(t, errObj)
	assert.Equal(t, expectedActors, actors)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetMovieActors_DbError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	movieID := 1

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
	`).WithArgs(movieID).WillReturnError(fmt.Errorf("some error"))

	r := NewMovieRepository(db)

	actors, errObj := r.GetMovieActors(context.Background(), movieID)

	assert.Nil(t, actors)
	assert.NotNil(t, errObj)
	assert.Equal(t, errors.ErrServerCode, errObj.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}
