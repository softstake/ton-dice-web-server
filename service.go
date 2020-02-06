package bets

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/tonradar/ton-dice-web-server/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	pb "github.com/tonradar/ton-dice-web-server/proto"
)

type BetService struct {
	store *storage.SalStore
}

func NewBetService(store *storage.SalStore) *BetService {
	return &BetService{store: store}
}

func (s *BetService) CreateBet(ctx context.Context, in *pb.CreateBetRequest) (*pb.CreateBetResponse, error) {
	req := storage.CreateBetReq{
		GameID:        in.GameId,
		PlayerAddress: in.PlayerAddress,
		RefAddress:    in.RefAddress,
		Amount:        in.Amount,
		RollUnder:     int8(in.RollUnder),
		RandomRoll:    int8(in.RandomRoll),
		Seed:          in.Seed,
		//Signature:     in.Signature,
		//PlayerPayout:  in.PlayerPayout,
		//RefPayout:     in.RefPayout,
	}

	fmt.Printf("req: %v:", req)

	resp, err := s.store.CreateBet(context.Background(), req)
	if err != nil {
		return nil, err
	}

	pts, err := ptypes.TimestampProto(resp.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &pb.CreateBetResponse{Id: resp.ID, CreatedAt: pts}, nil
}

func (s *BetService) Start() {
	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	rpcserv := grpc.NewServer()

	pb.RegisterBetsServer(rpcserv, &BetService{})
	reflection.Register(rpcserv)

	err = rpcserv.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve", err)
	}
}
