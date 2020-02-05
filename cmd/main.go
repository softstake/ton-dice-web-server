package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"

	service "github.com/tonradar/ton-dice-web-server"
	"github.com/tonradar/ton-dice-web-server/storage"
)

func main() {
	connStr := "user=postgres password=docker dbname=postgres sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	store := storage.NewStore(db)
	s := service.NewBetService(store)
	s.Start()
}
