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
		ID:            in.Id,
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

// IsBetCreated - used by GRPC
func (s *BetsService) IsBetCreated(ctx context.Context, in *pb.IsBetCreatedRequest) (*pb.IsBetCreatedResponse, error) {
	req := storage.GetBetReq{
		ID: in.Id,
	}

	resp, err := s.Store.GetBet(ctx, req)
	if err != nil {
		log.Printf("get bet failed with %s\n", err)
		return nil, err
	}

	isCreated := false
	if len(resp) > 0 {
		isCreated = true
	}

	return &pb.IsBetCreatedResponse{Yes: isCreated}, nil
}

// IsBetResolved - used by GRPC
func (s *BetsService) IsBetResolved(ctx context.Context, in *pb.IsBetResolvedRequest) (*pb.IsBetResolvedResponse, error) {
	req := storage.GetBetReq{
		ID: in.Id,
	}

	resp, err := s.Store.GetBet(ctx, req)
	if err != nil {
		log.Printf("get bet failed with %s\n", err)
		return nil, err
	}

	isResolved := false
	if resp[0].RandomRoll.Int32 > 0 {
		isResolved = true
	}

	return &pb.IsBetResolvedResponse{Yes: isResolved}, nil
}
