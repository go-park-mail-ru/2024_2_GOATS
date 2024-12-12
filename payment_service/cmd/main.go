package main

import (
	"log"

	"github.com/go-park-mail-ru/2024_2_GOATS/payment_service/internal/payment"
)

func main() {
	ua, err := payment.New(false)
	if err != nil {
		log.Fatal(err)
	}

	ua.Run()
}
