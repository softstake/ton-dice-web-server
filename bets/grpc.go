package bets

import (
	pb "github.com/tonradar/ton-dice-web-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GRPCServer struct {
	BetService *BetService
}

func NewGRPCServer(s *BetService) *GRPCServer {
	return &GRPCServer{
		BetService: s,
	}
}

func (s *GRPCServer) Start() {
	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	rpcserv := grpc.NewServer()

	pb.RegisterBetsServer(rpcserv, pb.BetsServer(s.BetService))
	reflection.Register(rpcserv)

	err = rpcserv.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve", err)
	}
}
