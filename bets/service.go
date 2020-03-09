package bets

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	pb "github.com/tonradar/ton-dice-web-server/proto"
	"github.com/tonradar/ton-dice-web-server/storage"
	"log"
)

type BetService struct {
	Store *storage.SalStore
}

func NewBetService(store *storage.SalStore) *BetService {
	return &BetService{Store: store}
}

func (s *BetService) Init() error {
	err := s.Store.Init(context.Background(), &storage.InitReq{})
	if err != nil {
		return err
	}
	return nil
}

func (s *BetService) CreateBet(ctx context.Context, in *pb.CreateBetRequest) (*pb.CreateBetResponse, error) {
	req := storage.CreateBetReq{
		GameID:        in.GameId,
		PlayerAddress: in.PlayerAddress,
		RefAddress:    in.RefAddress,
		Amount:        in.Amount,
		RollUnder:     int8(in.RollUnder),
		RandomRoll:    int8(in.RandomRoll),
		PlayerPayout:  in.PlayerPayout,
		Seed:          in.Seed,
		TrxHash:       in.TrxHash,
		TrxLt:         in.TrxLt,
	}

	resp, err := s.Store.CreateBet(ctx, req)
	if err != nil {
		log.Printf("save bet in DB failed with %s\n", err)
		return nil, err
	}

	pts, err := ptypes.TimestampProto(resp.CreatedAt)
	if err != nil {
		return nil, err
	}
	log.Printf("bet with id %d successfully saved (date: %s)", resp.ID, resp.CreatedAt)

	return &pb.CreateBetResponse{Id: resp.ID, CreatedAt: pts}, nil
}
