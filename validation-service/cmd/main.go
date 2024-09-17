package main

import (
	"context"
	"log"

	"github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/app"
)

const serverPort = 5050

func main() {
	ctx := context.Background()
	a, err := app.New(ctx, serverPort)

	if err != nil {
		log.Fatal(err)
	}

	a.Run()
}
