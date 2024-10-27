package movie

import (
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

func ScanMovieConnection(rows *sql.Rows) (*models.MovieInfo, error) {
	mvInfo := &models.MovieInfo{}
	directorInfo := &models.DirectorInfo{}

	for rows.Next() {
		err := rows.Scan(
			&mvInfo.Id,
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
			log.Err(errMsg)

			return nil, errMsg
		}
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Err(fmt.Errorf("cannot close rows while taking movie info: %w", err))
		}
	}()

	mvInfo.Director = directorInfo

	return mvInfo, nil
}

func ScanActorsConnections(rows *sql.Rows) ([]*models.ActorInfo, error) {
	actorInfos := []*models.ActorInfo{}

	for rows.Next() {
		var actorInfo models.ActorInfo

		err := rows.Scan(
			&actorInfo.Id,
			&actorInfo.Name,
			&actorInfo.Surname,
			&actorInfo.Biography,
			&actorInfo.SmallPhotoUrl,
		)

		if err != nil {
			errMsg := fmt.Errorf("error while scanning actors info: %w", err)
			log.Err(errMsg)

			return nil, errMsg
		}

		actorInfos = append(actorInfos, &actorInfo)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Err(fmt.Errorf("cannot close rows while taking movie info: %w", err))
		}
	}()

	return actorInfos, nil
}

func ScanActorMoviesConnections(rows *sql.Rows) ([]*models.MovieShortInfo, error) {
	actMvs := []*models.MovieShortInfo{}

	for rows.Next() {
		var mvShortInfo models.MovieShortInfo

		err := rows.Scan(
			&mvShortInfo.Id,
			&mvShortInfo.Title,
			&mvShortInfo.CardUrl,
			&mvShortInfo.Rating,
			&mvShortInfo.ReleaseDate,
		)

		if err != nil {
			errMsg := fmt.Errorf("error while scanning actor's movies short info: %w", err)
			log.Err(errMsg)

			return nil, errMsg
		}

		actMvs = append(actMvs, &mvShortInfo)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Err(fmt.Errorf("cannot close rows while taking actor's movies short info: %w", err))
		}
	}()

	return actMvs, nil
}
