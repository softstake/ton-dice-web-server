package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/tonradar/ton-dice-web-server/bets"
	"github.com/tonradar/ton-dice-web-server/storage"
	"github.com/tonradar/ton-dice-web-server/webserver"
	"log"
	"os"
)

func main() {
	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgName := os.Getenv("PG_NAME")
	pgUser := os.Getenv("PG_USER")
	pgPwd := os.Getenv("PG_PWD")

	if pgHost == "" || pgPort == "" || pgName == "" || pgUser == "" || pgPwd == "" {
		log.Fatalln("Some of required ENV vars are empty. The vars are: PG_HOST, PG_PORT, PG_NAME, PG_USER, PG_PWD")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pgHost, pgPort, pgUser, pgPwd, pgName)

	db, err := sql.Open("postgres", psqlInfo)
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
