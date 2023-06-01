package main

import (
	"chart/internal/config"
	"chart/storage"
	"fmt"
	"log"
)

func main() {
	config := config.Set()

	s, err := storage.New(config.ConfigDB)
	if err != nil {
		log.Fatalf("Couldn't connect db: %s\n", err)
	}

	fmt.Printf("%s", s.GetDB())

}
