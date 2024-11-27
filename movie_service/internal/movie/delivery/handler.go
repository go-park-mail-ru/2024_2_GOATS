package delivery

import (
	"context"
	"log"

	movie "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/pkg/movie_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MovieHandler struct {
	movieService MovieServiceInterface
	movie.UnimplementedMovieServiceServer
}

func NewMovieHandler(service MovieServiceInterface) *MovieHandler {
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
		return nil, err
	}

	respp := make([]*movie.MovieShortInfo, len(movies))
	for i, mov := range movies {
		respp[i] = &movie.MovieShortInfo{
			Id:          int32(mov.ID),
			CardUrl:     mov.CardURL,
			MovieType:   mov.MovieType,
			AlbumUrl:    mov.AlbumURL,
			Title:       mov.Title,
			Country:     mov.Country,
			ReleaseDate: mov.ReleaseDate,
			Rating:      mov.Rating,
		}
	}

	return &movie.GetMovieByGenreResponse{Movies: respp}, nil
}

func (h *MovieHandler) GetMovie(ctx context.Context, req *movie.GetMovieRequest) (*movie.GetMovieResponse, error) {
	if req.MovieId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid movie ID")
	}

	movieg, err := h.movieService.GetMovie(ctx, int(req.MovieId))
	if err != nil {
		return nil, err
	}

	var respp movie.MovieInfo

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
	respp.DirectorInfo = &movie.DirectorInfo{
		Name:    movieg.Director.Name,
		Surname: movieg.Director.Surname,
	}
	respp.FullDescription = movieg.FullDescription
	respp.ShortDescription = movieg.ShortDescription
	respp.TitleUrl = movieg.TitleURL

	var actors []*movie.ActorInfo
	for _, actor := range movieg.Actors {
		cur := &movie.ActorInfo{
			Id:            int32(actor.ID),
			Name:          actor.Person.Name,
			Surname:       actor.Person.Surname,
			Biography:     actor.Biography,
			Post:          actor.Post,
			Birthdate:     actor.Birthdate.String,
			SmallPhotoUrl: actor.SmallPhotoURL,
			BigPhotoUrl:   actor.BigPhotoURL,
			Country:       actor.Country,
		}

		actors = append(actors, cur)
	}

	respp.ActorsInfo = actors

	var seasons []*movie.Season
	for _, season := range movieg.Seasons {
		sn := season.SeasonNumber
		var eps []*movie.Episode
		for _, ep := range season.Episodes {
			cur := &movie.Episode{
				Id:            int64(ep.ID),
				Description:   ep.Description,
				EpisodeNumber: int64(ep.EpisodeNumber),
				Title:         ep.Title,
				Rating:        ep.Rating,
				ReleaseDate:   ep.ReleaseDate,
				VideoURL:      ep.VideoURL,
				PreviewURL:    ep.PreviewURL,
			}

			eps = append(eps, cur)
		}

		curSeas := &movie.Season{
			SeasonNumber: int32(sn),
			Episodes:     eps,
		}

		seasons = append(seasons, curSeas)
	}

	respp.Seasons = seasons

	return &movie.GetMovieResponse{Movie: &respp}, nil
}

func (h *MovieHandler) GetActor(ctx context.Context, req *movie.GetActorRequest) (*movie.GetActorResponse, error) {
	if req.ActorId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid actor ID")
	}

	actor, err := h.movieService.GetActor(ctx, int(req.ActorId))
	if err != nil {
		return nil, err
	}

	log.Println("actorDel", actor)

	var respp movie.ActorInfo
	respp.Id = int32(actor.ID)
	respp.Birthdate = actor.Birthdate.String
	respp.Country = actor.Country
	respp.BigPhotoUrl = actor.BigPhotoURL
	respp.Biography = actor.Biography
	respp.Name = actor.Person.Name
	respp.Surname = actor.Person.Surname
	respp.Post = actor.Post
	respp.SmallPhotoUrl = actor.SmallPhotoURL

	var mvs []*movie.MovieShortInfo
	for _, mv := range actor.Movies {
		cur := &movie.MovieShortInfo{
			Id:          int32(mv.ID),
			Title:       mv.Title,
			CardUrl:     mv.CardURL,
			AlbumUrl:    mv.AlbumURL,
			Rating:      mv.Rating,
			ReleaseDate: mv.ReleaseDate,
			MovieType:   mv.MovieType,
			Country:     mv.Country,
		}

		mvs = append(mvs, cur)
	}

	respp.Movies = mvs

	return &movie.GetActorResponse{Actor: &respp}, nil
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

		respp[i].FullDescription = movie.FullDescription
		respp[i].ShortDescription = movie.ShortDescription
		respp[i].TitleUrl = movie.TitleURL
		// for j, actor := range movie.Actors {
		// 	respp[i].ActorsInfo[j].Person.Name = actor.Person.Name
		// 	respp[i].ActorsInfo[j].Person.Surname = actor.Person.Surname
		// 	respp[i].ActorsInfo[j].Id = int32(actor.ID)
		// 	respp[i].ActorsInfo[j].Biography = actor.Biography
		// 	respp[i].ActorsInfo[j].Post = actor.Post
		// 	respp[i].ActorsInfo[j].Birthdate = actor.Birthdate.String
		// 	respp[i].ActorsInfo[j].SmallPhotoUrl = actor.SmallPhotoURL
		// 	respp[i].ActorsInfo[j].BigPhotoUrl = actor.BigPhotoURL
		// 	respp[i].ActorsInfo[j].Country = actor.Country
		// }
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
	log.Println("del", actors)
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
		respp[i].Name = v.Person.Name
		respp[i].Surname = v.Person.Surname
		respp[i].Post = v.Post
		respp[i].SmallPhotoUrl = v.SmallPhotoURL
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
	log.Println("respp", respp)
	return &movie.SearchActorsResponse{Actors: respp}, nil
}

func (h *MovieHandler) GetCollections(ctx context.Context, req *movie.GetCollectionsRequest) (*movie.GetCollectionsResponse, error) {
	collections, err := h.movieService.GetCollection(ctx, req.Filter)
	if err != nil {
		return nil, err
	}

	var respp []*movie.Collection
	for _, col := range collections.Collections {
		var mvs []*movie.MovieShortInfo
		for _, mv := range col.Movies {
			curMV := &movie.MovieShortInfo{
				Id:          int32(mv.ID),
				Title:       mv.Title,
				CardUrl:     mv.CardURL,
				AlbumUrl:    mv.AlbumURL,
				Rating:      mv.Rating,
				ReleaseDate: mv.ReleaseDate,
				MovieType:   mv.MovieType,
				Country:     mv.Country,
			}

			mvs = append(mvs, curMV)
		}
		cur := &movie.Collection{
			Id:     int32(col.ID),
			Title:  col.Title,
			Movies: mvs,
		}

		respp = append(respp, cur)
	}

	return &movie.GetCollectionsResponse{Collections: respp}, nil
}

func (h *MovieHandler) GetFavorites(ctx context.Context, req *movie.GetFavoritesRequest) (*movie.GetFavoritesResponse, error) {
	favs, err := h.movieService.GetFavorites(ctx, req.MovieIds)
	if err != nil {
		return nil, err
	}

	var respp []*movie.MovieShortInfo
	for _, fav := range favs {
		curMV := &movie.MovieShortInfo{
			Id:          int32(fav.ID),
			Title:       fav.Title,
			CardUrl:     fav.CardURL,
			AlbumUrl:    fav.AlbumURL,
			Rating:      fav.Rating,
			ReleaseDate: fav.ReleaseDate,
			MovieType:   fav.MovieType,
			Country:     fav.Country,
		}

		respp = append(respp, curMV)
	}

	return &movie.GetFavoritesResponse{Movies: respp}, nil
}
