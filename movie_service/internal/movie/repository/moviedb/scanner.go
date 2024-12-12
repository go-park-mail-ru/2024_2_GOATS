package moviedb

import (
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie/repository/dto"
	"github.com/rs/zerolog/log"
)

// ScanMovieConnection scans full movie info from sql Rows
func ScanMovieConnection(rows *sql.Rows) (*models.MovieInfo, error) {
	mvInfo := &models.MovieInfo{}
	directorInfo := &models.DirectorInfo{}
	seasons := make(map[int]models.Season)

	defer func() {
		if err := rows.Close(); err != nil {
			errMsg := fmt.Errorf("cannot close rows while taking movie_service info: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())
		}
	}()

	for rows.Next() {
		var seasonNumber sql.NullInt64
		episode := &models.DBEpisode{}

		err := rows.Scan(
			&mvInfo.ID,
			&mvInfo.Title,
			&mvInfo.ShortDescription,
			&mvInfo.FullDescription,
			&mvInfo.CardURL,
			&mvInfo.AlbumURL,
			&mvInfo.Rating,
			&mvInfo.ReleaseDate,
			&mvInfo.VideoURL,
			&mvInfo.MovieType,
			&mvInfo.TitleURL,
			&directorInfo.Name,
			&directorInfo.Surname,
			&mvInfo.Country,
			&episode.ID,
			&episode.Title,
			&episode.Description,
			&seasonNumber,
			&episode.EpisodeNumber,
			&episode.ReleaseDate,
			&episode.Rating,
			&episode.PreviewURL,
			&episode.VideoURL,
		)

		if episode.ID.Valid && seasonNumber.Valid {
			sn := int(seasonNumber.Int64)
			if _, exists := seasons[sn]; !exists {
				seasons[sn] = models.Season{SeasonNumber: sn, Episodes: []*models.Episode{}}
			}

			season := seasons[sn]
			season.Episodes = append(season.Episodes, converter.ToRepoEpisodeFromDB(episode))

			seasons[sn] = season
		}

		if err != nil {
			errMsg := fmt.Errorf("error while scanning movies info: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())

			return nil, errMsg
		}
	}

	mvInfo.Director = directorInfo
	seasRes := make([]*models.Season, 0, len(seasons))
	for _, season := range seasons {
		seasRes = append(seasRes, &season)
	}

	mvInfo.Seasons = seasRes

	return mvInfo, nil
}

// ScanActorsConnections scans short actors info for movie page from sql Rows
func ScanActorsConnections(rows *sql.Rows) ([]*dto.RepoActor, error) {
	actorInfos := []*dto.RepoActor{}

	defer func() {
		if err := rows.Close(); err != nil {
			errMsg := fmt.Errorf("cannot close rows while taking movie_service info: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())
		}
	}()

	for rows.Next() {
		var actorInfo dto.RepoActor

		err := rows.Scan(
			&actorInfo.ID,
			&actorInfo.Name,
			&actorInfo.Surname,
			&actorInfo.Biography,
			&actorInfo.SmallPhotoURL,
		)

		if err != nil {
			errMsg := fmt.Errorf("error while scanning actors info: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())

			return nil, errMsg
		}

		actorInfos = append(actorInfos, &actorInfo)
	}

	return actorInfos, nil
}

// ScanActorMoviesConnections scans full actors info for actor page from sql Rows
func ScanActorMoviesConnections(rows *sql.Rows) ([]*dto.RepoMovieShortInfo, error) {
	actMvs := []*dto.RepoMovieShortInfo{}

	defer func() {
		if err := rows.Close(); err != nil {
			errMsg := fmt.Errorf("cannot close rows while taking actor's movies short info: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())
		}
	}()

	for rows.Next() {
		var mvShortInfo dto.RepoMovieShortInfo

		err := rows.Scan(
			&mvShortInfo.ID,
			&mvShortInfo.Title,
			&mvShortInfo.CardURL,
			&mvShortInfo.Rating,
			&mvShortInfo.ReleaseDate,
			&mvShortInfo.Country,
		)

		if err != nil {
			errMsg := fmt.Errorf("error while scanning actor's movies short info: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())

			return nil, errMsg
		}

		actMvs = append(actMvs, &mvShortInfo)
	}

	return actMvs, nil
}

// ScanMovieShortConnection scans short movie info from sql Rows
func ScanMovieShortConnection(rows *sql.Rows) ([]*models.MovieShortInfo, error) {
	var movies []*models.MovieShortInfo

	defer func() {
		if err := rows.Close(); err != nil {
			errMsg := fmt.Errorf("cannot close rows while taking favorites: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())
		}
	}()

	for rows.Next() {
		var movie = &models.MovieShortInfo{}

		err := rows.Scan(&movie.ID, &movie.Title, &movie.CardURL, &movie.AlbumURL, &movie.Rating, &movie.ReleaseDate, &movie.MovieType, &movie.Country)

		if err != nil {
			errMsg := fmt.Errorf("error while scanning favorites: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())

			return nil, errMsg
		}

		movies = append(movies, movie)
	}

	return movies, nil
}
