package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *MovieService) GetMovie(ctx context.Context, mvId int) (*models.MovieInfo, *models.ErrorRespData) {
	mv, err, code := s.movieRepository.GetMovie(ctx, mvId)

	if err != nil {
		errs := make([]errVals.ErrorObj, 1)
		errs[0] = *err

		return nil, &models.ErrorRespData{
			StatusCode: code,
			Errors:     errs,
		}
	}

	staffs, err, code := s.movieRepository.GetStaffInfo(ctx, mv.Id)

	if err != nil {
		errs := make([]errVals.ErrorObj, 1)
		errs[0] = *err

		return nil, &models.ErrorRespData{
			StatusCode: code,
			Errors:     errs,
		}
	}

	acInfo := []*models.StaffInfo{}
	directorInfo := []*models.StaffInfo{}

	for _, staff := range staffs {
		if staff.Post == "actor" {
			acInfo = append(acInfo, staff)
		}

		if staff.Post == "director" {
			directorInfo = append(directorInfo, staff)
		}
	}

	mv.Actors = acInfo
	mv.Directors = directorInfo

	return mv, nil
}
