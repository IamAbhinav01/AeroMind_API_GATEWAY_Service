package main

import (
	"AeromindGO/app"
	config "AeromindGO/config/env"
	"log"
)


func main() {
	config.Load()

	application := app.NewApplication()
	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}