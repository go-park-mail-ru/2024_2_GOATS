package api

import (
	"context"
	"log"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
)

func (i *Implementation) GetCollection(ctx context.Context) {
	log.Println("From api: ", config.GetConfigFromContext(ctx))
	i.service.GetCollection(ctx)
}
