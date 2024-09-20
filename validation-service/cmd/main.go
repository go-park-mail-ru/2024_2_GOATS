package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/app"
)

func main() {
	a, ctx, err := app.New()

	if err != nil {
		log.Fatal(err)
	}

	a.Run(ctx)
}
