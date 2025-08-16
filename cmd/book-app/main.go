package main

import (
	"book-app/internal/app"
	"book-app/pkg/config"
	"log"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}
	app.Run()
}
