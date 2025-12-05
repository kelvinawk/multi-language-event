package main

import (
	"log"
	"multi-language-event/config"
	"multi-language-event/db"
	store "multi-language-event/internal/storage"
	"multi-language-event/server"
)

func main() {
	cfg := config.Load()

	db, err := db.NewPostgresConn(cfg.Db_dsn)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer db.Close()

	store := store.NewStore(db)
	server := server.NewHTTPServer(store, cfg)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
