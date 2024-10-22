package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app"
)

func main() {
	a, err := app.New(false)
	if err != nil {
		log.Fatal(err.Error())
	}

	a.Run()
}
