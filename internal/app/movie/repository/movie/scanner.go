package movie

import (
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

func ScanConnections(rows *sql.Rows) (*models.MovieInfo, error) {
	mvInfo := &models.MovieInfo{}
	actorsInfo := []*models.StaffInfo{}

	for rows.Next() {
		var actorInfo models.StaffInfo

		err := rows.Scan(&actorInfo.Id, &actorInfo.Name, &actorInfo.Surname,
			&actorInfo.Patronymic, &actorInfo.Biography, &actorInfo.Post, &actorInfo.SmallPhotoUrl,
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
			&mvInfo.Country,
		)

		if err != nil {
			errMsg := fmt.Errorf("error while scanning actors info: %w", err)
			log.Err(errMsg)

			return nil, errMsg
		}

		actorsInfo = append(actorsInfo, &actorInfo)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Err(fmt.Errorf("cannot close rows while taking movie info: %w", err))
		}
	}()

	acInfo := []*models.StaffInfo{}
	directorInfo := []*models.StaffInfo{}

	for _, staff := range actorsInfo {
		if staff.Post == "actor" {
			acInfo = append(acInfo, staff)
		}

		if staff.Post == "director" {
			directorInfo = append(directorInfo, staff)
		}
	}

	mvInfo.Actors = acInfo
	mvInfo.Directors = directorInfo

	return mvInfo, nil
}
