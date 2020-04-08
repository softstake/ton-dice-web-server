package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/tonradar/ton-dice-web-server/bets"
	"github.com/tonradar/ton-dice-web-server/config"
	"github.com/tonradar/ton-dice-web-server/storage"
	"github.com/tonradar/ton-dice-web-server/webserver"
)

func main() {

	cfg := config.GetConfig()

	pgConnString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=allow",
		cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPwd, cfg.PgName)

	db, err := sql.Open("postgres", pgConnString)
	if err != nil {
		log.Fatal(err)
	}

	store := storage.NewStore(db)

	betsService := bets.NewBetsService(store, &cfg)

	for {
		err = betsService.Init()
		if err != nil {
			log.Printf("failed to init storage: %v, retrying...", err)
			time.Sleep(3000 * time.Millisecond)
			continue
		}
		break
	}

	gRPCServer := bets.NewGRPCServer(betsService, &cfg)
	go gRPCServer.Start()

	webServer := webserver.NewWebserver(betsService, &cfg)
	webServer.Start()
}
