package bets

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes"
	"github.com/tonradar/ton-dice-web-server/config"
	pb "github.com/tonradar/ton-dice-web-server/proto"
	"github.com/tonradar/ton-dice-web-server/storage"
)

type BetsService struct {
	Store  *storage.SalStore
	config *config.TonWebServerConfig
}

func NewBetsService(store *storage.SalStore, cfg *config.TonWebServerConfig) *BetsService {
	return &BetsService{
		Store:  store,
		config: cfg,
	}
}

func (s *BetsService) Init() error {
	err := s.Store.Init(context.Background(), &storage.InitReq{})
	if err != nil {
		return err
	}
	return nil
}

// CreateBet - used by GRPC
func (s *BetsService) CreateBet(ctx context.Context, in *pb.CreateBetRequest) (*pb.CreateBetResponse, error) {
	req := storage.CreateBetReq{
		GameID:        in.GameId,
		PlayerAddress: in.PlayerAddress,
		RefAddress:    in.RefAddress,
		Amount:        in.Amount,
		RollUnder:     int8(in.RollUnder),
		Seed:          in.Seed,
		CreateTrxHash: in.CreateTrxHash,
		CreateTrxLt:   in.CreateTrxLt,
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

// UpdateBet - used by GRPC
func (s *BetsService) UpdateBet(ctx context.Context, in *pb.UpdateBetRequest) (*pb.UpdateBetResponse, error) {
	req := storage.UpdateBetReq{
		ID:             in.Id,
		GameID:         in.GameId,
		RandomRoll:     int8(in.RandomRoll),
		PlayerPayout:   in.PlayerPayout,
		RefPayout:      in.RefPayout,
		ResolveTrxHash: in.ResolveTrxHash,
		ResolveTrxLt:   in.ResolveTrxLt,
	}

	resp, err := s.Store.UpdateBet(ctx, req)
	if err != nil {
		log.Printf("update bet in DB failed with %s\n", err)
		return nil, err
	}

	pts, err := ptypes.TimestampProto(resp.ResolvedAt)
	if err != nil {
		return nil, err
	}
	log.Printf("bet with id %d successfully updated (date: %s)", resp.ID, resp.ResolvedAt)

	return &pb.UpdateBetResponse{Id: resp.ID, ResolvedAt: pts}, nil
}

// IsBetFetched - used by GRPC
func (s *BetsService) IsBetFetched(ctx context.Context, in *pb.IsBetFetchedRequest) (*pb.IsBetFetchedResponse, error) {
	req := storage.GetFetchedBetsReq{
		GameID:        in.GameId,
		CreateTrxHash: in.CreateTrxHash,
		CreateTrxLt:   in.CreateTrxLt,
	}

	resp, err := s.Store.GetFetchedBet(ctx, req)
	if err != nil {
		log.Printf("get bet failed with %s\n", err)
		return nil, err
	}

	isBetExist := false
	if len(resp) > 0 {
		isBetExist = true
	}

	return &pb.IsBetFetchedResponse{Yes: isBetExist, Id: resp[0].ID}, nil
}

// IsBetResolved - used by GRPC
func (s *BetsService) IsBetResolved(ctx context.Context, in *pb.IsBetResolvedRequest) (*pb.IsBetResolvedResponse, error) {
	req := storage.GetResolvedBetsReq{
		GameID:         in.GameId,
		ResolveTrxHash: in.ResolveTrxHash,
		ResolveTrxLt:   in.ResolveTrxLt,
	}

	resp, err := s.Store.GetResolvedBet(ctx, req)
	if err != nil {
		log.Printf("get bet failed with %s\n", err)
		return nil, err
	}

	isBetExist := false
	if len(resp) > 0 {
		isBetExist = true
	}

	return &pb.IsBetResolvedResponse{Yes: isBetExist, Id: resp[0].ID}, nil
}
