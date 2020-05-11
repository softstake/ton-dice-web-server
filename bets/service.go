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
	ctx := context.Background()

	err := s.Store.Init(ctx, &storage.InitReq{})
	_ = s.Store.PrepareExpressionFirst(ctx, &storage.PrepareExpressionFirstReq{})
	_ = s.Store.PrepareExpressionSecond(ctx, &storage.PrepareExpressionSecondReq{})

	if err != nil {
		return err
	}
	return nil
}

// SaveBet - used by GRPC
func (s *BetsService) SaveBet(ctx context.Context, in *pb.SaveBetRequest) (*pb.SaveBetResponse, error) {
	req := storage.SaveBetReq{
		ID:            in.Id,
		PlayerAddress: in.PlayerAddress,
		RefAddress:    in.RefAddress,
		Amount:        in.Amount,
		RollUnder:     int8(in.RollUnder),
		Seed:          in.Seed,
		CreateTrxHash: in.CreateTrxHash,
		CreateTrxLt:   in.CreateTrxLt,
	}

	resp, err := s.Store.SaveBet(ctx, req)
	if err != nil {
		log.Printf("save bet in DB failed with %s\n", err)
		return nil, err
	}

	getBetReq := storage.GetBetReq{
		ID: in.Id,
	}
	bet, err := s.Store.GetBet(ctx, getBetReq)
	if err != nil {
		log.Printf("get bet from DB after saving is failed: %s\n", err)
		return nil, err
	}

	pbUpdatedAt, err := ptypes.TimestampProto(bet.UpdatedAt)
	if err != nil {
		log.Println("Convert datetime error", err)
		return nil, err
	}
	log.Printf("bet with id %d successfully saved (date: %s)", resp.ID, bet.CreatedAt)

	return &pb.SaveBetResponse{Id: resp.ID, State: int32(bet.State), UpdatedAt: pbUpdatedAt}, nil
}

// UpdateBet - used by GRPC
func (s *BetsService) UpdateBet(ctx context.Context, in *pb.UpdateBetRequest) (*pb.UpdateBetResponse, error) {

	req := storage.UpdateBetReq{
		ID:             in.Id,
		State:          storage.BetStateFromInt32(in.State),
		RandomRoll:     int8(in.RandomRoll),
		Signature:      in.Signature,
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

	pbUpdatedAt, err := ptypes.TimestampProto(resp.UpdatedAt)
	if err != nil {
		log.Println("Convert datetime error", err)
		return nil, err
	}
	log.Printf("bet with id %d successfully updated (date: %s)", resp.ID, resp.UpdatedAt)

	return &pb.UpdateBetResponse{Id: resp.ID, State: int32(resp.State), UpdatedAt: pbUpdatedAt}, nil
}

// IsBetCreated - used by GRPC
func (s *BetsService) IsBetSaved(ctx context.Context, in *pb.IsBetSavedRequest) (*pb.IsBetSavedResponse, error) {
	req := storage.GetBetReq{
		ID: in.Id,
	}

	resp, err := s.Store.GetBet(ctx, req)
	if err != nil {
		log.Printf("get bet failed with %s\n", err)
		return nil, err
	}

	isSaved := false
	// check for empty struct
	if (resp != storage.GetBetResp{}) {
		isSaved = true
	}

	return &pb.IsBetSavedResponse{IsSaved: isSaved}, nil
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
	if resp.State == storage.BetResolved {
		isResolved = true
	}

	return &pb.IsBetResolvedResponse{IsResolved: isResolved}, nil
}

// GetBet - used by GRPC
func (s *BetsService) GetBet(ctx context.Context, in *pb.GetBetRequest) (*pb.GetBetResponse, error) {
	req := storage.GetBetReq{
		ID: in.Id,
	}

	bet, err := s.Store.GetBet(ctx, req)
	if err != nil {
		log.Printf("get bet from DB failed with %s\n", err)
		return nil, err
	}

	pbCreatedAt, err := ptypes.TimestampProto(bet.CreatedAt)
	if err != nil {
		log.Println("Convert datetime error", err)
		return nil, err
	}

	pbUpdatedAt, err := ptypes.TimestampProto(bet.UpdatedAt)
	if err != nil {
		log.Println("Convert datetime error", err)
		return nil, err
	}

	return &pb.GetBetResponse{
		Id:             bet.ID,
		Amount:         bet.Amount,
		State:          int32(bet.State),
		RollUnder:      int32(bet.RollUnder),
		PlayerAddress:  bet.PlayerAddress,
		RefAddress:     bet.RefAddress,
		Seed:           bet.Seed,
		Signature:      bet.Signature,
		RandomRoll:     int32(bet.RandomRoll),
		PlayerPayout:   bet.PlayerPayout,
		RefPayout:      bet.RefPayout,
		CreatedAt:      pbCreatedAt,
		CreateTrxHash:  bet.CreateTrxHash,
		CreateTrxLt:    bet.CreateTrxLt,
		UpdatedAt:      pbUpdatedAt,
		ResolveTrxHash: bet.ResolveTrxHash,
		ResolveTrxLt:   bet.ResolveTrxLt,
	}, nil
}
