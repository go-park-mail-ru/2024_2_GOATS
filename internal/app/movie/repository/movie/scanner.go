package movie

import (
	"database/sql"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/rs/zerolog/log"
)

func ScanConnections(rows *sql.Rows) ([]*models.ActorInfo, error) {
	actorsInfo := []*models.ActorInfo{}

	for rows.Next() {
		var actorInfo models.ActorInfo

		err := rows.Scan(&actorInfo.Id, &actorInfo.Name, &actorInfo.Surname, &actorInfo.Patronymic, &actorInfo.PhotoUrl)

		if err != nil {
			errMsg := fmt.Errorf("error while scanning actors info: %w", err)
			log.Err(errMsg)

			return nil, errMsg
		}

		actorsInfo = append(actorsInfo, &actorInfo)
	}

	return actorsInfo, nil
}
