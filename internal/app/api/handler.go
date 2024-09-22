package api

import (
	"context"
	"log"
	"net/url"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
)

func (i *Implementation) GetCollection(ctx context.Context, query url.Values) {
	log.Println("From api: ", config.FromContext(ctx))
	// r := i.service.GetCollection(ctx)

	// return json.Marshal(r)
}
