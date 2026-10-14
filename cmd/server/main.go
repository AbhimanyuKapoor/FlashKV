package main

import (
	"log"

	"github.com/Hari08-dev/flashkv/internal/server"
)

func main() {
	srv := server.NewServer(":6379")

	log.Println("FlashKV server starting on port 6379...")

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
