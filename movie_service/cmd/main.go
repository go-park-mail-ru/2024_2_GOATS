package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_GOATS/movie_service/internal/movie"
)

func main() {
	a, err := movie.New(false)
	if err != nil {
		log.Fatal(err)
		log.Fatal(err)
	}

	a.Run()
}
