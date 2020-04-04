package bets

import (
	"fmt"
	"log"
	"net"

	"github.com/tonradar/ton-dice-web-server/config"
	pb "github.com/tonradar/ton-dice-web-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	betsService *BetsService
	config      *config.TonWebServerConfig
}

func NewGRPCServer(s *BetsService, cfg *config.TonWebServerConfig) *GRPCServer {
	return &GRPCServer{
		betsService: s,
		config:      cfg,
	}
}

func (s *GRPCServer) Start() {
	listener, err := net.Listen("tcp", fmt.Sprint(":%d", s.config.RPCListenPort))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	rpcServer := grpc.NewServer()

	pb.RegisterBetsServer(rpcServer, pb.BetsServer(s.betsService))
	reflection.Register(rpcServer)

	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve", err)
	}
}
