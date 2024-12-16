package client

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	movie "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/pkg/movie_v1"
)

// MovieClientInterface defines client methods to transmit to Movie Microservice
//
//go:generate mockgen -source=movie.go -destination=../user/service/mocks/movie_mock.go
//go:generate mockgen -source=movie.go -destination=../movie/service/mocks/mock.go

// MovieClientInterface интерфейс клиента фильмов
type MovieClientInterface interface {
	// GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, error)
	GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, error)
	GetActor(ctx context.Context, actorID int) (*models.ActorInfo, error)
	SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error)
	SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error)
	GetCollection(ctx context.Context, filter string) ([]models.Collection, error)
	GetFavorites(ctx context.Context, mvIDs []uint64) ([]models.MovieShortInfo, error)
	GetUserRating(ctx context.Context, movieID, userID int32) (int32, error)
	AddOrUpdateRating(ctx context.Context, movieID, userID, rating int32) error
	DeleteUserRating(ctx context.Context, movieID, userID int32) error
}

// MovieClient struct implements MovieClientInterface
type MovieClient struct {
	movieMS movie.MovieServiceClient
}

// NewMovieClient returns an instance of MovieClientInterface
func NewMovieClient(movieMS movie.MovieServiceClient) MovieClientInterface {
	return &MovieClient{
		movieMS: movieMS,
	}
}

// GetCollection get movie collections
func (m MovieClient) GetCollection(ctx context.Context, filter string) ([]models.Collection, error) {
	start := time.Now()
	method := "GetCollection"
	resp, err := m.movieMS.GetCollections(ctx, &movie.GetCollectionsRequest{Filter: filter})
	saveMetric(start, movieClient, method, err)

	if err != nil {
		return nil, err
	}

	var ans = make([]models.Collection, 0, len(resp.Collections))

	for _, col := range resp.Collections {
		var mvs []*models.MovieShortInfo
		for _, mv := range col.Movies {
			curMV := &models.MovieShortInfo{
				ID:          int(mv.Id),
				Title:       mv.Title,
				CardURL:     mv.CardUrl,
				AlbumURL:    mv.AlbumUrl,
				Rating:      mv.Rating,
				ReleaseDate: mv.ReleaseDate,
				MovieType:   mv.MovieType,
				Country:     mv.Country,
			}

			mvs = append(mvs, curMV)
		}
		cur := models.Collection{
			ID:     int(col.Id),
			Title:  col.Title,
			Movies: mvs,
		}

		ans = append(ans, cur)
	}

	return ans, nil
}

// GetMovie get movie
func (m MovieClient) GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, error) {
	start := time.Now()
	method := "GetMovie"

	resp, err := m.movieMS.GetMovie(ctx, &movie.GetMovieRequest{MovieId: int32(mvID)})

	saveMetric(start, movieClient, method, err)

	if err != nil {
		return nil, err
	}

	respMov := resp.Movie
	var respp = &models.MovieInfo{}

	respp.ID = int(respMov.Id)
	respp.CardURL = respMov.CardUrl
	respp.AlbumURL = respMov.AlbumUrl
	respp.Rating = respMov.Rating
	respp.Title = respMov.Title
	respp.MovieType = respMov.MovieType
	respp.Country = respMov.Country
	respp.ReleaseDate = respMov.ReleaseDate
	respp.IsFavorite = respMov.IsFavorite
	respp.VideoURL = respMov.VideoUrl
	respp.Director = &models.DirectorInfo{
		Person: models.Person{
			Name:    respMov.DirectorInfo.Name,
			Surname: respMov.DirectorInfo.Surname,
		},
	}
	respp.FullDescription = respMov.FullDescription
	respp.ShortDescription = respMov.ShortDescription
	respp.TitleURL = respMov.TitleUrl
	respp.WithSubscription = respMov.WithSubscription

	var actors []*models.ActorInfo
	for _, actor := range respMov.ActorsInfo {
		cur := &models.ActorInfo{
			ID: int(actor.Id),
			Person: models.Person{
				Name:    actor.Name,
				Surname: actor.Surname,
			},
			Biography:     actor.Biography,
			Post:          actor.Post,
			Birthdate:     sql.NullString{String: actor.Birthdate, Valid: true},
			SmallPhotoURL: actor.SmallPhotoUrl,
			BigPhotoURL:   actor.BigPhotoUrl,
			Country:       actor.Country,
		}

		actors = append(actors, cur)
	}

	respp.Actors = actors

	var seasons []*models.Season

	for _, season := range respMov.Seasons {
		sn := season.SeasonNumber
		var eps []*models.Episode
		for _, ep := range season.Episodes {
			cur := &models.Episode{
				ID:            int(ep.Id),
				Description:   ep.Description,
				EpisodeNumber: int(ep.EpisodeNumber),
				Title:         ep.Title,
				Rating:        ep.Rating,
				ReleaseDate:   ep.ReleaseDate,
				VideoURL:      ep.VideoURL,
				PreviewURL:    ep.PreviewURL,
			}

			eps = append(eps, cur)
		}

		curSeas := &models.Season{
			SeasonNumber: int(sn),
			Episodes:     eps,
		}

		seasons = append(seasons, curSeas)
	}

	respp.Seasons = seasons

	return respp, nil
}

// GetActor get actor
func (m MovieClient) GetActor(ctx context.Context, actorID int) (*models.ActorInfo, error) {
	start := time.Now()
	method := "GetActor"

	resp, err := m.movieMS.GetActor(ctx, &movie.GetActorRequest{ActorId: int32(actorID)})

	saveMetric(start, movieClient, method, err)

	if err != nil {
		return nil, err
	}

	respActor := resp.Actor
	var respp = &models.ActorInfo{}

	respp.ID = int(respActor.Id)
	respp.Birthdate = sql.NullString{String: respActor.Birthdate, Valid: true}
	respp.Country = respActor.Country
	respp.BigPhotoURL = respActor.BigPhotoUrl
	respp.Biography = respActor.Biography
	respp.Name = respActor.Name
	respp.Surname = respActor.Surname
	respp.Post = respActor.Post
	respp.SmallPhotoURL = respActor.SmallPhotoUrl

	var mvs []*models.MovieShortInfo
	for _, mv := range respActor.Movies {
		cur := &models.MovieShortInfo{
			ID:          int(mv.Id),
			Title:       mv.Title,
			CardURL:     mv.CardUrl,
			AlbumURL:    mv.AlbumUrl,
			Rating:      mv.Rating,
			ReleaseDate: mv.ReleaseDate,
			MovieType:   mv.MovieType,
			Country:     mv.Country,
		}

		mvs = append(mvs, cur)
	}

	respp.Movies = mvs

	return respp, nil
}

// func (m MovieClient) GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, error) {
// resp, err := m.movieMS.GetMovieByGenre(ctx, &movie.GetMovieByGenreRequest{Genre: genre})
// if err != nil {
// 	return nil, err
// }

// 	respMovie := resp.Movies

// var respp = make([]models.MovieShortInfo, 0, len(respMovie))

// 	for i, movie := range respMovie {
// 		respp[i].ID = int(movie.ID)
// 		respp[i].CardURL = movie.CardUrl
// 		respp[i].MovieType = movie.MovieType
// 		respp[i].AlbumURL = movie.AlbumUrl
// 		respp[i].Title = movie.Title
// 		respp[i].Country = movie.Country
// 		respp[i].ReleaseDate = movie.ReleaseDate
// 		respp[i].Rating = movie.Rating
// 	}
// 	return respp, nil
// }

// SearchMovies search movies
func (m MovieClient) SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error) {
	start := time.Now()
	method := "SearchMovies"

	resp, err := m.movieMS.SearchMovies(ctx, &movie.SearchMoviesRequest{Query: query})

	saveMetric(start, movieClient, method, err)

	if err != nil {
		return nil, err
	}

	respMovie := resp.Movies

	var respp = make([]models.MovieInfo, len(respMovie))

	for i, mov := range respMovie {
		respp[i] = models.MovieInfo{
			ID:       int(mov.Id),
			CardURL:  mov.CardUrl,
			AlbumURL: mov.AlbumUrl,
			Rating:   mov.Rating,
			Title:    mov.Title,
		}
	}
	return respp, nil
}

// SearchActors search actors
func (m MovieClient) SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error) {
	start := time.Now()
	method := "SearchActors"

	resp, err := m.movieMS.SearchActors(ctx, &movie.SearchActorsRequest{Query: query})

	saveMetric(start, movieClient, method, err)

	if err != nil {
		return nil, err
	}

	respActor := resp.Actors

	var respp = make([]models.ActorInfo, len(respActor))

	for i, v := range respActor {
		respp[i] = models.ActorInfo{
			ID:          int(v.Id),
			BigPhotoURL: v.BigPhotoUrl,
			Person: models.Person{
				Name:    v.Name,
				Surname: v.Surname,
			},
		}
	}
	return respp, nil
}

// GetFavorites get favorites info
func (m MovieClient) GetFavorites(ctx context.Context, mvIDs []uint64) ([]models.MovieShortInfo, error) {
	start := time.Now()
	method := "GetFavorites"

	resp, err := m.movieMS.GetFavorites(ctx, &movie.GetFavoritesRequest{MovieIds: mvIDs})

	saveMetric(start, movieClient, method, err)

	if err != nil {
		return nil, err
	}

	var ans = make([]models.MovieShortInfo, 0, len(resp.Movies))

	for _, mv := range resp.Movies {
		curMV := models.MovieShortInfo{
			ID:          int(mv.Id),
			Title:       mv.Title,
			CardURL:     mv.CardUrl,
			AlbumURL:    mv.AlbumUrl,
			Rating:      mv.Rating,
			ReleaseDate: mv.ReleaseDate,
			MovieType:   mv.MovieType,
			Country:     mv.Country,
		}

		ans = append(ans, curMV)
	}

	return ans, nil
}

// GetUserRating получение рейтинга
func (m MovieClient) GetUserRating(ctx context.Context, movieID, userID int32) (int32, error) {
	start := time.Now()
	method := "GetUserRating"

	resp, err := m.movieMS.GetUserRating(ctx, &movie.GetUserRatingRequest{
		MovieId: movieID,
		UserId:  userID,
	})

	saveMetric(start, movieClient, method, err)

	if err != nil {
		return 0, err
	}

	return int32(resp.Rating.Rating), nil
}

// AddOrUpdateRating добавление рейтинга
func (m MovieClient) AddOrUpdateRating(ctx context.Context, movieID, userID, rating int32) error {
	start := time.Now()
	method := "AddOrUpdateRating"

	_, err := m.movieMS.AddOrUpdateRating(ctx, &movie.AddOrUpdateRatingRequest{
		MovieId: movieID,
		UserId:  userID,
		Rating:  rating,
	})

	saveMetric(start, movieClient, method, err)

	return err
}

// DeleteUserRating удаление рейтинга
func (m *MovieClient) DeleteUserRating(ctx context.Context, userID, movieID int32) error {
	start := time.Now()
	method := "AddOrUpdateRating"

	_, err := m.movieMS.DeleteRating(ctx, &movie.DeleteRatingRequest{
		MovieId: int32(movieID),
		UserId:  int32(userID),
	})

	saveMetric(start, movieClient, method, err)

	return err
}
