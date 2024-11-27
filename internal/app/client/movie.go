package client

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	movie "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/pkg/movie_v1"
)

//go:generate mockgen -source=movie.go -destination=../user/service/mocks/movie_mock.go
//go:generate mockgen -source=movie.go -destination=../movie/service/mocks/mock.go
type MovieClientInterface interface {
	GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, error)
	GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, error)
	GetActor(ctx context.Context, actorID int) (*models.ActorInfo, error)
	SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error)
	SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error)
	GetCollection(ctx context.Context, filter string) ([]models.Collection, error)
	GetFavorites(ctx context.Context, mvIDs []uint64) ([]models.MovieShortInfo, error)
}

type MovieClient struct {
	movieMS movie.MovieServiceClient
}

func NewMovieClient(movieMS movie.MovieServiceClient) MovieClientInterface {
	return &MovieClient{
		movieMS: movieMS,
	}
}

func (m MovieClient) GetCollection(ctx context.Context, filter string) ([]models.Collection, error) {
	resp, err := m.movieMS.GetCollections(ctx, &movie.GetCollectionsRequest{Filter: filter})

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

func (m MovieClient) GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, error) {
	resp, err := m.movieMS.GetMovie(ctx, &movie.GetMovieRequest{MovieId: int32(mvID)})
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

func (m MovieClient) GetActor(ctx context.Context, actorID int) (*models.ActorInfo, error) {
	resp, err := m.movieMS.GetActor(ctx, &movie.GetActorRequest{ActorId: int32(actorID)})
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

func (m MovieClient) GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, error) {
	resp, err := m.movieMS.GetMovieByGenre(ctx, &movie.GetMovieByGenreRequest{Genre: genre})
	if err != nil {
		return nil, err
	}

	respMovie := resp.Movies

	var respp = make([]models.MovieShortInfo, 0, len(respMovie))

	for i, movie := range respMovie {
		respp[i].ID = int(movie.Id)
		respp[i].CardURL = movie.CardUrl
		respp[i].MovieType = movie.MovieType
		respp[i].AlbumURL = movie.AlbumUrl
		respp[i].Title = movie.Title
		respp[i].Country = movie.Country
		respp[i].ReleaseDate = movie.ReleaseDate
		respp[i].Rating = movie.Rating
	}
	return respp, nil
}

func (m MovieClient) SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error) {
	resp, err := m.movieMS.SearchMovies(ctx, &movie.SearchMoviesRequest{Query: query})
	if err != nil {
		return nil, err
	}

	respActor := resp.Movies

	var respp = make([]models.MovieInfo, 0, len(respActor))

	for i, movie := range respActor {
		respp[i].ID = int(movie.Id)
		respp[i].CardURL = movie.CardUrl
		respp[i].AlbumURL = movie.AlbumUrl
		respp[i].Rating = movie.Rating
		respp[i].Title = movie.Title
		respp[i].MovieType = movie.MovieType
		respp[i].Country = movie.Country
		respp[i].ReleaseDate = movie.ReleaseDate
		respp[i].IsFavorite = movie.IsFavorite
		respp[i].VideoURL = movie.VideoUrl

		respp[i].FullDescription = movie.FullDescription
		respp[i].ShortDescription = movie.ShortDescription
		respp[i].TitleURL = movie.TitleUrl
		for j, actor := range movie.ActorsInfo {
			respp[i].Actors[j].Person.Name = actor.Name
			respp[i].Actors[j].Person.Surname = actor.Surname
			respp[i].Actors[j].ID = int(actor.Id)
			respp[i].Actors[j].Biography = actor.Biography
			respp[i].Actors[j].Post = actor.Post
			respp[i].Actors[j].Birthdate = sql.NullString{String: actor.Birthdate}
			respp[i].Actors[j].SmallPhotoURL = actor.SmallPhotoUrl
			respp[i].Actors[j].BigPhotoURL = actor.BigPhotoUrl
			respp[i].Actors[j].Country = actor.Country
		}
		for s, season := range movie.Seasons {
			respp[i].Seasons[s].SeasonNumber = int(season.SeasonNumber)
			for g, ep := range season.Episodes {
				respp[i].Seasons[s].Episodes[g].ID = int(ep.Id)
				respp[i].Seasons[s].Episodes[g].Description = ep.Description
				respp[i].Seasons[s].Episodes[g].EpisodeNumber = int(ep.EpisodeNumber)
				respp[i].Seasons[s].Episodes[g].Title = ep.Title
				respp[i].Seasons[s].Episodes[g].Rating = ep.Rating
				respp[i].Seasons[s].Episodes[g].ReleaseDate = ep.ReleaseDate
				respp[i].Seasons[s].Episodes[g].VideoURL = ep.VideoURL
				respp[i].Seasons[s].Episodes[g].PreviewURL = ep.PreviewURL
			}

		}
	}
	return respp, nil
}

func (m MovieClient) SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error) {
	resp, err := m.movieMS.SearchActors(ctx, &movie.SearchActorsRequest{Query: query})
	if err != nil {
		return nil, err
	}

	respActor := resp.Actors

	var respp = make([]models.ActorInfo, 0, len(respActor))

	for i, v := range respActor {
		respp[i].ID = int(v.Id)
		respp[i].Birthdate = sql.NullString{String: v.Birthdate}
		respp[i].Country = v.Country
		respp[i].BigPhotoURL = v.BigPhotoUrl
		respp[i].Biography = v.Biography
		respp[i].Person.Name = v.Name
		respp[i].Person.Surname = v.Surname
		respp[i].Post = v.Post
		respp[i].SmallPhotoURL = v.SmallPhotoUrl
		for j, mov := range v.Movies {
			respp[i].Movies[j].ID = int(mov.Id)
			respp[i].Movies[j].Title = mov.Title
			respp[i].Movies[j].Rating = mov.Rating
			respp[i].Movies[j].ReleaseDate = mov.ReleaseDate
			respp[i].Movies[j].Country = mov.Country
			respp[i].Movies[j].MovieType = mov.MovieType
			respp[i].Movies[j].CardURL = mov.CardUrl
			respp[i].Movies[j].AlbumURL = mov.AlbumUrl
		}
	}
	return respp, nil
}

func (m MovieClient) GetFavorites(ctx context.Context, mvIDs []uint64) ([]models.MovieShortInfo, error) {
	resp, err := m.movieMS.GetFavorites(ctx, &movie.GetFavoritesRequest{MovieIds: mvIDs})

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
