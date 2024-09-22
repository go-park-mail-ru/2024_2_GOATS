package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app"
)

func main() {
	a, ctx, err := app.New()
	if err != nil {
		log.Fatalf(err.Error())
	}

	a.Run(ctx)
}
