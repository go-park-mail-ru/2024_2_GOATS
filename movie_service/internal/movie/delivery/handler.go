package delivery

import (
	"context"
	movie "github.com/go-park-mail-ru/2024_2_GOATS/movie_service/pkg/movie_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
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

	respp := make([]*movie.MovieInfo, len(movies))
	for i, mov := range movies {
		respp[i] = &movie.MovieInfo{
			Id:          int32(mov.ID),
			CardUrl:     mov.CardURL,
			AlbumUrl:    mov.AlbumURL,
			Rating:      mov.Rating,
			Title:       mov.Title,
			MovieType:   mov.MovieType,
			Country:     mov.Country,
			ReleaseDate: mov.ReleaseDate,
			IsFavorite:  mov.IsFavorite,
			VideoUrl:    mov.VideoURL,

			FullDescription:  mov.FullDescription,
			ShortDescription: mov.ShortDescription,
			TitleUrl:         mov.TitleURL,
		}
	}
	log.Println("resppmovie", respp)
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

	respp := make([]*movie.ActorInfo, len(actors))
	for i, v := range actors {
		respp[i] = &movie.ActorInfo{
			Id:          int32(v.ID),
			BigPhotoUrl: v.BigPhotoURL,
			Name:        v.Person.Name,
			Surname:     v.Person.Surname,
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

func (h *MovieHandler) GetUserRating(ctx context.Context, req *movie.GetUserRatingRequest) (*movie.GetUserRatingResponse, error) {
	if req.MovieId <= 0 || req.UserId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid req.MovieId or req.UserId")
	}

	rating, err := h.movieService.GetUserRating(ctx, int(req.UserId), int(req.MovieId))
	if err != nil {
		return nil, err
	}

	return &movie.GetUserRatingResponse{
		Rating: &movie.UserRating{
			UserId:  req.UserId,
			MovieId: req.MovieId,
			Rating:  rating,
		},
	}, nil
}

func (h *MovieHandler) AddOrUpdateRating(ctx context.Context, req *movie.AddOrUpdateRatingRequest) (*movie.AddOrUpdateRatingResponse, error) {
	if req.MovieId <= 0 || req.UserId <= 0 || req.Rating < 1 || req.Rating > 10 {
		return nil, status.Error(codes.InvalidArgument, "invalid req.UserId or req.Rating req.Rating")
	}

	err := h.movieService.AddOrUpdateRating(ctx, int(req.UserId), int(req.MovieId), float32(req.Rating))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *MovieHandler) DeleteRating(ctx context.Context, req *movie.DeleteRatingRequest) (*movie.DeleteRatingResponse, error) {
	if req.MovieId <= 0 || req.UserId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid movie ID or user ID")
	}

	err := h.movieService.DeleteRating(ctx, int(req.MovieId), int(req.UserId))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
