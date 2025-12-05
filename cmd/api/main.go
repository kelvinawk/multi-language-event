package main

import (
	"log"
	"multi-language-event/config"
	store "multi-language-event/internal/storage"
	"multi-language-event/server"
)

func main() {
	cfg := config.Load()
	store := store.NewStore(nil)
	server := server.NewHTTPServer(store, cfg)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
