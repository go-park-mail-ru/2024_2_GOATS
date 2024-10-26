package repository

import (
	"context"
	"net/http"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/repository/movie"
)

func (r *Repo) GetStaffInfo(ctx context.Context, mvId int) ([]*models.StaffInfo, *errVals.ErrorObj, int) {
	rows, err := movie.GetStaff(ctx, mvId, r.Database)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	staffInfos, err := movie.ScanStaffConnections(rows)
	if err != nil {
		return nil, errVals.NewErrorObj(errVals.ErrServerCode, errVals.CustomError{Err: err}), http.StatusUnprocessableEntity
	}

	return staffInfos, nil, http.StatusOK
}
