package main

import (
	"context"
	"log"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app"
)

func main() {
	ctx := context.Background()

	a, authImpl, err := app.New(ctx)
	if err != nil {
		log.Fatalf("can't create app: %s", err)
	}

	if err = a.Run(ctx, authImpl); err != nil {
		log.Fatalf("can't run app: %s", err)
	}
}
