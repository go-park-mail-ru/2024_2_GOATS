package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_GOATS/auth_service/internal/auth"
)

func main() {
	a, err := auth.New(false)
	if err != nil {
		log.Fatal(err)
	}

	a.Run()
}
