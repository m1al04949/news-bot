package main

import (
	"log"

	"github.com/m1al04949/news-bot/internal/app"
)

func main() {
	if err := app.RunBot(); err != nil {
		log.Fatal(err)
	}
}
