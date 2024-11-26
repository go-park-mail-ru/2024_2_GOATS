package client

import (
	"context"
	"database/sql"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	movie "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/pkg/movie_v1"
)

type MovieClientInterface interface {
	GetMovieActors(ctx context.Context, mvID int) ([]*models.ActorInfo, error)
	GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, error)
	GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, error)
	GetActor(ctx context.Context, actorID int) (*models.ActorInfo, error)
	SearchMovies(ctx context.Context, query string) ([]models.MovieInfo, error)
	SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error)
	GetCollection(ctx context.Context, filter string) ([]models.Collection, error)
}

type MovieClient struct {
	movieMS movie.MovieServiceServer
}

func NewMovieClient(movieMS movie.MovieServiceServer) MovieClientInterface {
	return &MovieClient{
		movieMS: movieMS,
	}
}

func (m MovieClient) GetCollection(ctx context.Context, filter string) ([]models.Collection, error) {
	resp, err := m.movieMS.GetCollections(ctx, &movie.GetCollectionsRequest{Filter: filter})

	if err != nil {
		return nil, err
	}

	var ans []models.Collection
	for i, v := range resp.Collections {
		ans[i].ID = int(v.Id)
		ans[i].Title = v.Title
		for j, vv := range ans[i].Movies {
			ans[i].Movies[j].ID = vv.ID
			ans[i].Movies[j].MovieType = vv.MovieType
			ans[i].Movies[j].Rating = vv.Rating
			ans[i].Movies[j].ReleaseDate = vv.ReleaseDate
			ans[i].Movies[j].Country = vv.Country
			ans[i].Movies[j].Title = vv.Title
			ans[i].Movies[j].AlbumURL = vv.AlbumURL
			ans[i].Movies[j].CardURL = vv.CardURL
		}
	}
	return ans, nil
}

func (m MovieClient) GetMovie(ctx context.Context, mvID int) (*models.MovieInfo, error) {
	resp, err := m.movieMS.GetMovie(ctx, &movie.GetMovieRequest{MovieId: int32(mvID)})
	if err != nil {
		return nil, err
	}

	respMov := resp.Movie

	var respp *models.MovieInfo

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
	respp.Director.ID = int(respMov.DirectorInfo.Id)

	respp.FullDescription = respMov.FullDescription
	respp.ShortDescription = respMov.ShortDescription
	respp.TitleURL = respMov.TitleUrl
	for j, actor := range respMov.ActorsInfo {
		respp.Actors[j].Person.Name = actor.Person.Name
		respp.Actors[j].Person.Surname = actor.Person.Surname
		respp.Actors[j].ID = int(actor.Id)
		respp.Actors[j].Biography = actor.Biography
		respp.Actors[j].Post = actor.Post
		respp.Actors[j].Birthdate = sql.NullString{String: actor.Birthdate}
		respp.Actors[j].SmallPhotoURL = actor.SmallPhotoUrl
		respp.Actors[j].BigPhotoURL = actor.BigPhotoUrl
		respp.Actors[j].Country = actor.Country
	}
	for s, season := range respMov.Seasons {
		respp.Seasons[s].SeasonNumber = int(season.SeasonNumber)
		for g, ep := range season.Episodes {
			respp.Seasons[s].Episodes[g].ID = int(ep.Id)
			respp.Seasons[s].Episodes[g].Description = ep.Description
			respp.Seasons[s].Episodes[g].EpisodeNumber = int(ep.EpisodeNumber)
			respp.Seasons[s].Episodes[g].Title = ep.Title
			respp.Seasons[s].Episodes[g].Rating = ep.Rating
			respp.Seasons[s].Episodes[g].ReleaseDate = ep.ReleaseDate
			respp.Seasons[s].Episodes[g].VideoURL = ep.VideoURL
			respp.Seasons[s].Episodes[g].PreviewURL = ep.PreviewURL
		}

	}
	return respp, nil
}

func (m MovieClient) GetActor(ctx context.Context, actorID int) (*models.ActorInfo, error) {
	resp, err := m.movieMS.GetActor(ctx, &movie.GetActorRequest{ActorId: int32(actorID)})
	if err != nil {
		return nil, err
	}

	respActor := resp.Actor

	var respp *models.ActorInfo

	respp.ID = int(respActor.Id)
	respp.Birthdate = sql.NullString{String: respActor.Birthdate}
	respp.Country = respActor.Country
	respp.BigPhotoURL = respActor.BigPhotoUrl
	respp.Biography = respActor.Biography
	respp.Person.Name = respActor.Person.Name
	respp.Person.Surname = respActor.Person.Surname
	respp.Post = respActor.Post
	respp.SmallPhotoURL = respActor.SmallPhotoUrl
	respp.Person.Name = respActor.Person.Name
	respp.Person.Surname = respActor.Person.Surname
	for j, mov := range respActor.Movies {
		respp.Movies[j].ID = int(mov.Id)
		respp.Movies[j].Title = mov.Title
		respp.Movies[j].Rating = mov.Rating
		respp.Movies[j].ReleaseDate = mov.ReleaseDate
		respp.Movies[j].Country = mov.Country
		respp.Movies[j].MovieType = mov.MovieType
		respp.Movies[j].CardURL = mov.CardUrl
		respp.Movies[j].AlbumURL = mov.AlbumUrl
	}

	return respp, nil
}

func (m MovieClient) GetMovieActors(ctx context.Context, mvID int) ([]*models.ActorInfo, error) {
	resp, err := m.movieMS.GetMovieActors(ctx, &movie.GetMovieActorsRequest{MovieId: int32(mvID)})
	if err != nil {
		return nil, err
	}

	respActorInfo := resp.ActorsInfo

	var respp []*models.ActorInfo

	for i, v := range respActorInfo {
		respp[i].ID = int(v.Id)
		respp[i].Birthdate = sql.NullString{String: v.Birthdate}
		respp[i].Country = v.Country
		respp[i].BigPhotoURL = v.BigPhotoUrl
		respp[i].Biography = v.Biography
		respp[i].Person.Name = v.Person.Name
		respp[i].Person.Surname = v.Person.Surname
		respp[i].Post = v.Post
		respp[i].SmallPhotoURL = v.SmallPhotoUrl
		respp[i].Person.Name = v.Person.Name
		respp[i].Person.Surname = v.Person.Surname
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

func (m MovieClient) GetMovieByGenre(ctx context.Context, genre string) ([]models.MovieShortInfo, error) {
	resp, err := m.movieMS.GetMovieByGenre(ctx, &movie.GetMovieByGenreRequest{Genre: genre})
	if err != nil {
		return nil, err
	}

	respMovie := resp.Movies

	var respp []models.MovieShortInfo

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

	var respp []models.MovieInfo

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
		respp[i].Director.ID = int(movie.DirectorInfo.Id)

		respp[i].FullDescription = movie.FullDescription
		respp[i].ShortDescription = movie.ShortDescription
		respp[i].TitleURL = movie.TitleUrl
		for j, actor := range movie.ActorsInfo {
			respp[i].Actors[j].Person.Name = actor.Person.Name
			respp[i].Actors[j].Person.Surname = actor.Person.Surname
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

	var respp []models.ActorInfo

	for i, v := range respActor {
		respp[i].ID = int(v.Id)
		respp[i].Birthdate = sql.NullString{String: v.Birthdate}
		respp[i].Country = v.Country
		respp[i].BigPhotoURL = v.BigPhotoUrl
		respp[i].Biography = v.Biography
		respp[i].Person.Name = v.Person.Name
		respp[i].Person.Surname = v.Person.Surname
		respp[i].Post = v.Post
		respp[i].SmallPhotoURL = v.SmallPhotoUrl
		respp[i].Person.Name = v.Person.Name
		respp[i].Person.Surname = v.Person.Surname
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
