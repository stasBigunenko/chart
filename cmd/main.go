package main

import (
	"chart/internal/config"
	"chart/storage"
	"log"
)

func main() {
	config := config.Set()

	_, err := storage.New(config.ConfigDB)
	if err != nil {
		log.Fatalf("Couldn't connect db: %s\n", err)
	}

}
