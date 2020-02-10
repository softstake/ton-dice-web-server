package bets

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	bets "ton-dice-web-server/proto"
	"ton-dice-web-server/storage"
)

type BetService struct {
	store *storage.SalStore
}

func NewBetService(store *storage.SalStore) bets.BetsServer {
	return &BetService{store: store}
}

func (s *BetService) CreateBet(ctx context.Context, in *bets.CreateBetRequest) (*bets.CreateBetResponse, error) {
	req := storage.CreateBetReq{
		GameID:        in.GameId,
		PlayerAddress: in.PlayerAddress,
		RefAddress:    in.RefAddress,
		Amount:        in.Amount,
		RollUnder:     int8(in.RollUnder),
		RandomRoll:    int8(in.RandomRoll),
		Seed:          in.Seed,
	}

	resp, err := s.store.CreateBet(ctx, req)
	if err != nil {
		return nil, err
	}

	pts, err := ptypes.TimestampProto(resp.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &bets.CreateBetResponse{Id: resp.ID, CreatedAt: pts}, nil
}
