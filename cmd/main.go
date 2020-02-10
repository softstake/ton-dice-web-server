package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	bets "ton-dice-web-server/proto"

	service "ton-dice-web-server"
	"ton-dice-web-server/storage"
)

func main() {
	connStr := "user=postgres password=docker dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	store := storage.NewStore(db)
	s := service.NewBetService(store)

	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	rpcserv := grpc.NewServer()

	bets.RegisterBetsServer(rpcserv, s)
	reflection.Register(rpcserv)

	err = rpcserv.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve", err)
	}
}
