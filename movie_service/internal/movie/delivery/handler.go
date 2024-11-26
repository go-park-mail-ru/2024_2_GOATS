package delivery

import (
	"context"
	movie "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/pkg/movie_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MovieHandler struct {
	movieService MovieServiceInterface
	movie.UnimplementedMovieServiceServer
}

func NewMovieHandler(ctx context.Context, service MovieServiceInterface) *MovieHandler {
	return &MovieHandler{
		movieService: service,
	}
}

func (h *MovieHandler) GetMovieByGenre(ctx context.Context, req *movie.GetMovieByGenreRequest) (*movie.GetMovieByGenreResponse, error) {
	if req.Genre == "" {
		return nil, status.Error(codes.InvalidArgument, "genre is required")
	}

	movies, err := h.movieService.GetMovieByGenre(ctx, req.Genre)
	if err != nil {
		return nil, err.Error
	}

	var respp []*movie.MovieShortInfo
	for i, movie := range movies {
		respp[i].Id = int32(movie.ID)
		respp[i].CardUrl = movie.CardURL
		respp[i].MovieType = movie.MovieType
		respp[i].AlbumUrl = movie.AlbumURL
		respp[i].Title = movie.Title
		respp[i].Country = movie.Country
		respp[i].ReleaseDate = movie.ReleaseDate
		respp[i].Rating = movie.Rating
	}

	return &movie.GetMovieByGenreResponse{Movies: respp}, nil
}

func (h *MovieHandler) GetMovie(ctx context.Context, req *movie.GetMovieRequest) (*movie.GetMovieResponse, error) {
	if req.MovieId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid movie ID")
	}

	movieg, err := h.movieService.GetMovie(ctx, int(req.MovieId))
	if err != nil {
		return nil, err.Error
	}

	var respp *movie.MovieInfo

	respp.Id = int32(movieg.ID)
	respp.CardUrl = movieg.CardURL
	respp.AlbumUrl = movieg.AlbumURL
	respp.Rating = movieg.Rating
	respp.Title = movieg.Title
	respp.MovieType = movieg.MovieType
	respp.Country = movieg.Country
	respp.ReleaseDate = movieg.ReleaseDate
	respp.IsFavorite = movieg.IsFavorite
	respp.VideoUrl = movieg.VideoURL
	respp.DirectorInfo.Id = int32(movieg.Director.ID)

	respp.FullDescription = movieg.FullDescription
	respp.ShortDescription = movieg.ShortDescription
	respp.TitleUrl = movieg.TitleURL
	for j, actor := range movieg.Actors {
		respp.ActorsInfo[j].Person.Name = actor.Person.Name
		respp.ActorsInfo[j].Person.Surname = actor.Person.Surname
		respp.ActorsInfo[j].Id = int32(actor.ID)
		respp.ActorsInfo[j].Biography = actor.Biography
		respp.ActorsInfo[j].Post = actor.Post
		respp.ActorsInfo[j].Birthdate = actor.Birthdate.String
		respp.ActorsInfo[j].SmallPhotoUrl = actor.SmallPhotoURL
		respp.ActorsInfo[j].BigPhotoUrl = actor.BigPhotoURL
		respp.ActorsInfo[j].Country = actor.Country
	}
	for s, season := range movieg.Seasons {
		respp.Seasons[s].SeasonNumber = int32(season.SeasonNumber)
		for g, ep := range season.Episodes {
			respp.Seasons[s].Episodes[g].Id = int64(ep.ID)
			respp.Seasons[s].Episodes[g].Description = ep.Description
			respp.Seasons[s].Episodes[g].EpisodeNumber = int64(ep.EpisodeNumber)
			respp.Seasons[s].Episodes[g].Title = ep.Title
			respp.Seasons[s].Episodes[g].Rating = ep.Rating
			respp.Seasons[s].Episodes[g].ReleaseDate = ep.ReleaseDate
			respp.Seasons[s].Episodes[g].VideoURL = ep.VideoURL
			respp.Seasons[s].Episodes[g].PreviewURL = ep.PreviewURL
		}

	}

	return &movie.GetMovieResponse{Movie: respp}, nil
}

func (h *MovieHandler) GetActor(ctx context.Context, req *movie.GetActorRequest) (*movie.GetActorResponse, error) {
	if req.ActorId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid actor ID")
	}

	actor, err := h.movieService.GetActor(ctx, int(req.ActorId))
	if err != nil {
		return nil, err.Error
	}

	var respp *movie.ActorInfo
	respp.Id = int32(actor.ID)
	respp.Birthdate = actor.Birthdate.String
	respp.Country = actor.Country
	respp.BigPhotoUrl = actor.BigPhotoURL
	respp.Biography = actor.Biography
	respp.Person.Name = actor.Person.Name
	respp.Person.Surname = actor.Person.Surname
	respp.Post = actor.Post
	respp.SmallPhotoUrl = actor.SmallPhotoURL
	respp.Person.Name = actor.Person.Name
	respp.Person.Surname = actor.Person.Surname
	for j, mov := range actor.Movies {
		respp.Movies[j].Id = int32(mov.ID)
		respp.Movies[j].Title = mov.Title
		respp.Movies[j].Rating = mov.Rating
		respp.Movies[j].ReleaseDate = mov.ReleaseDate
		respp.Movies[j].Country = mov.Country
		respp.Movies[j].MovieType = mov.MovieType
		respp.Movies[j].CardUrl = mov.CardURL
		respp.Movies[j].AlbumUrl = mov.AlbumURL
	}

	return &movie.GetActorResponse{Actor: respp}, nil
}

func (h *MovieHandler) SearchMovies(ctx context.Context, req *movie.SearchMoviesRequest) (*movie.SearchMoviesResponse, error) {
	if req.Query == "" {
		return nil, status.Error(codes.InvalidArgument, "query is required")
	}

	movies, err := h.movieService.SearchMovies(ctx, req.Query)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var respp []*movie.MovieInfo
	for i, movie := range movies {
		respp[i].Id = int32(movie.ID)
		respp[i].CardUrl = movie.CardURL
		respp[i].AlbumUrl = movie.AlbumURL
		respp[i].Rating = movie.Rating
		respp[i].Title = movie.Title
		respp[i].MovieType = movie.MovieType
		respp[i].Country = movie.Country
		respp[i].ReleaseDate = movie.ReleaseDate
		respp[i].IsFavorite = movie.IsFavorite
		respp[i].VideoUrl = movie.VideoURL
		respp[i].DirectorInfo.Id = int32(movie.Director.ID)

		respp[i].FullDescription = movie.FullDescription
		respp[i].ShortDescription = movie.ShortDescription
		respp[i].TitleUrl = movie.TitleURL
		for j, actor := range movie.Actors {
			respp[i].ActorsInfo[j].Person.Name = actor.Person.Name
			respp[i].ActorsInfo[j].Person.Surname = actor.Person.Surname
			respp[i].ActorsInfo[j].Id = int32(actor.ID)
			respp[i].ActorsInfo[j].Biography = actor.Biography
			respp[i].ActorsInfo[j].Post = actor.Post
			respp[i].ActorsInfo[j].Birthdate = actor.Birthdate.String
			respp[i].ActorsInfo[j].SmallPhotoUrl = actor.SmallPhotoURL
			respp[i].ActorsInfo[j].BigPhotoUrl = actor.BigPhotoURL
			respp[i].ActorsInfo[j].Country = actor.Country
		}
		for s, season := range movie.Seasons {
			respp[i].Seasons[s].SeasonNumber = int32(season.SeasonNumber)
			for g, ep := range season.Episodes {
				respp[i].Seasons[s].Episodes[g].Id = int64(ep.ID)
				respp[i].Seasons[s].Episodes[g].Description = ep.Description
				respp[i].Seasons[s].Episodes[g].EpisodeNumber = int64(ep.EpisodeNumber)
				respp[i].Seasons[s].Episodes[g].Title = ep.Title
				respp[i].Seasons[s].Episodes[g].Rating = ep.Rating
				respp[i].Seasons[s].Episodes[g].ReleaseDate = ep.ReleaseDate
				respp[i].Seasons[s].Episodes[g].VideoURL = ep.VideoURL
				respp[i].Seasons[s].Episodes[g].PreviewURL = ep.PreviewURL
			}

		}
	}

	return &movie.SearchMoviesResponse{Movies: respp}, nil
}

func (h *MovieHandler) SearchActors(ctx context.Context, req *movie.SearchActorsRequest) (*movie.SearchActorsResponse, error) {
	if req.Query == "" {
		return nil, status.Error(codes.InvalidArgument, "query is required")
	}

	actors, err := h.movieService.SearchActors(ctx, req.Query)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var respp []*movie.ActorInfo
	for i, v := range actors {
		respp[i].Id = int32(v.ID)
		respp[i].Birthdate = v.Birthdate.String
		respp[i].Country = v.Country
		respp[i].BigPhotoUrl = v.BigPhotoURL
		respp[i].Biography = v.Biography
		respp[i].Person.Name = v.Person.Name
		respp[i].Person.Surname = v.Person.Surname
		respp[i].Post = v.Post
		respp[i].SmallPhotoUrl = v.SmallPhotoURL
		respp[i].Person.Name = v.Person.Name
		respp[i].Person.Surname = v.Person.Surname
		for j, mov := range v.Movies {
			respp[i].Movies[j].Id = int32(mov.ID)
			respp[i].Movies[j].Title = mov.Title
			respp[i].Movies[j].Rating = mov.Rating
			respp[i].Movies[j].ReleaseDate = mov.ReleaseDate
			respp[i].Movies[j].Country = mov.Country
			respp[i].Movies[j].MovieType = mov.MovieType
			respp[i].Movies[j].CardUrl = mov.CardURL
			respp[i].Movies[j].AlbumUrl = mov.AlbumURL
		}

	}
	return &movie.SearchActorsResponse{Actors: respp}, nil
}

func (h *MovieHandler) GetCollections(ctx context.Context, req *movie.GetCollectionsRequest) (*movie.GetCollectionsResponse, error) {
	collections, err := h.movieService.GetCollection(ctx, req.Filter)
	if err != nil {
		return nil, err.Error
	}

	var respp []*movie.Collection
	for i, collection := range collections.Collections {
		respp[i].Id = int32(collection.ID)
		respp[i].Title = collection.Title
		for j, movie := range collection.Movies {
			respp[i].Movies[j].MovieType = movie.MovieType
			respp[i].Movies[j].ReleaseDate = movie.ReleaseDate
			respp[i].Movies[j].Country = movie.Country
			respp[i].Movies[j].Title = movie.Title
			respp[i].Movies[j].AlbumUrl = movie.AlbumURL
			respp[i].Movies[j].CardUrl = movie.CardURL
			respp[i].Movies[j].Rating = movie.Rating
			respp[i].Movies[j].Id = int32(movie.ID)
		}
	}
	return &movie.GetCollectionsResponse{Collections: respp}, nil
}

func (h *MovieHandler) GetMovieActors(ctx context.Context, req *movie.GetMovieActorsRequest) (*movie.GetMovieActorsResponse, error) {
	actors, err := h.movieService.GetMovieActors(ctx, int(req.MovieId))
	if err != nil {
		return nil, err
	}

	var respp []*movie.ActorInfo
	for i, v := range actors {
		respp[i].Id = int32(v.ID)
		respp[i].Birthdate = v.Name
		respp[i].Country = v.Country
		respp[i].BigPhotoUrl = v.BigPhotoURL
		respp[i].Biography = v.Biography
		respp[i].Person.Name = v.Person.Name
		respp[i].Person.Surname = v.Person.Surname
		respp[i].Post = v.Post
		respp[i].SmallPhotoUrl = v.SmallPhotoURL
		respp[i].Person.Name = v.Person.Name
		respp[i].Person.Surname = v.Person.Surname
		for j, mov := range v.Movies {
			respp[i].Movies[j].Id = int32(mov.ID)
			respp[i].Movies[j].Title = mov.Title
			respp[i].Movies[j].Rating = mov.Rating
			respp[i].Movies[j].ReleaseDate = mov.ReleaseDate
			respp[i].Movies[j].Country = mov.Country
			respp[i].Movies[j].MovieType = mov.MovieType
			respp[i].Movies[j].CardUrl = mov.CardURL
			respp[i].Movies[j].AlbumUrl = mov.AlbumURL
		}

	}
	return &movie.GetMovieActorsResponse{ActorsInfo: respp}, nil
}
