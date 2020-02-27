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
	"time"
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

	for {
		err = s.Init()
		if err != nil {
			log.Printf("failed to init storage: %v, retrying...", err)
			time.Sleep(3000 * time.Millisecond)
			continue
		}
		break
	}

	grpc := bets.NewGRPCServer(s)
	go grpc.Start()

	server := webserver.NewWebserver(s)
	server.Start()
}
