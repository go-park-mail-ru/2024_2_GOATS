package moviedb

import (
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository/converter"
	"github.com/rs/zerolog/log"
)

func ScanMovieConnection(rows *sql.Rows) (*models.MovieInfo, error) {
	mvInfo := &models.MovieInfo{}
	directorInfo := &models.DirectorInfo{}
	seasons := make(map[int]models.Season)

	defer func() {
		if err := rows.Close(); err != nil {
			errMsg := fmt.Errorf("cannot close rows while taking movie info: %w", err)
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

func ScanActorsConnections(rows *sql.Rows) ([]*models.ActorInfo, error) {
	actorInfos := []*models.ActorInfo{}

	defer func() {
		if err := rows.Close(); err != nil {
			errMsg := fmt.Errorf("cannot close rows while taking movie info: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())
		}
	}()

	for rows.Next() {
		var actorInfo models.ActorInfo

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

func ScanActorMoviesConnections(rows *sql.Rows) ([]*models.MovieShortInfo, error) {
	actMvs := []*models.MovieShortInfo{}

	defer func() {
		if err := rows.Close(); err != nil {
			errMsg := fmt.Errorf("cannot close rows while taking actor's movies short info: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())
		}
	}()

	for rows.Next() {
		var mvShortInfo models.MovieShortInfo

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
