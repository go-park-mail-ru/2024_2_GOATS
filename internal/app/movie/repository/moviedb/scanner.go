package moviedb

import (
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

func ScanMovieConnection(rows *sql.Rows) (*models.MovieInfo, error) {
	mvInfo := &models.MovieInfo{}
	directorInfo := &models.DirectorInfo{}

	defer func() {
		if err := rows.Close(); err != nil {
			errMsg := fmt.Errorf("cannot close rows while taking movie info: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())
		}
	}()

	for rows.Next() {
		err := rows.Scan(
			&mvInfo.ID,
			&mvInfo.Title,
			&mvInfo.ShortDescription,
			&mvInfo.FullDescription,
			&mvInfo.CardUrl,
			&mvInfo.AlbumUrl,
			&mvInfo.Rating,
			&mvInfo.ReleaseDate,
			&mvInfo.VideoUrl,
			&mvInfo.MovieType,
			&mvInfo.TitleUrl,
			&directorInfo.Name,
			&directorInfo.Surname,
			&mvInfo.Country,
		)

		if err != nil {
			errMsg := fmt.Errorf("error while scanning actors info: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())

			return nil, errMsg
		}
	}

	mvInfo.Director = directorInfo

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
			&actorInfo.SmallPhotoUrl,
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
			&mvShortInfo.CardUrl,
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
