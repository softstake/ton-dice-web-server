package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/tonradar/ton-dice-web-server/bets"
	"github.com/tonradar/ton-dice-web-server/storage"
	"github.com/tonradar/ton-dice-web-server/webserver"
	"log"
)

func main() {
	connStr := "user=postgres password=docker dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	store := storage.NewStore(db)
	s := bets.NewBetService(store)
	err = s.Init()
	if err != nil {
		log.Fatal("failed to init storage: %v", err)
	}

	grpc := bets.NewGRPCServer(s)
	go grpc.Start()

	server := webserver.NewWebserver(s)
	server.Start()
}
