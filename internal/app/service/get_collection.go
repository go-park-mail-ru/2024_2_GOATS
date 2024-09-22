package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
)

func (s *Service) GetCollection(ctx context.Context) {
	fmt.Println("From service: ", config.GetConfigFromContext(ctx))
	s.repository.GetCollection(ctx)
}
