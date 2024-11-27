package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user"
)

func main() {
	ua, err := user.New(false)
	if err != nil {
		log.Fatal(err)
	}

	ua.Run()
}
