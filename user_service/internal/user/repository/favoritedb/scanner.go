package favoritedb

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
)

// ScanConnections scans user favorites from rows
func ScanConnections(rows *sql.Rows) ([]uint64, error) {
	defer func() {
		if err := rows.Close(); err != nil {
			errMsg := fmt.Errorf("cannot close rows while taking user favorites: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())
		}
	}()

	var movieIDs []uint64

	for rows.Next() {
		var mvID uint64
		err := rows.Scan(&mvID)

		if err != nil {
			errMsg := fmt.Errorf("error while scanning favorite movies: %w", err)
			log.Error().Err(errMsg).Msg(errMsg.Error())

			return nil, errMsg
		}

		movieIDs = append(movieIDs, mvID)
	}

	return movieIDs, nil
}
