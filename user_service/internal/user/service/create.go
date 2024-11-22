package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
	"github.com/rs/zerolog/log"
)

func (u *UserService) Create(ctx context.Context, createData *dto.CreateUserData) (uint64, error) {
	logger := log.Ctx(ctx)
	repoData := converter.ConvertToRepoCreateData(createData)
	usr, err := u.userRepo.CreateUser(ctx, repoData)
	if err != nil {
		logger.Error().Err(err).Msg("userService - failed to create user")
		return 0, fmt.Errorf("userService - failed to create user: %w", err)
	}

	return usr.ID, nil
}
