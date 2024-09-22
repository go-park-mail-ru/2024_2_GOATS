package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
)

func (r *Repo) GetCollection(ctx context.Context) {
	fmt.Println("From repo:", config.FromContext(ctx))
	r.Database.QueryContext(ctx, "SELECT * FROM movie_collections")
}
