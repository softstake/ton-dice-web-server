package storage

import (
	"context"
	"time"
)

type Store interface {
	Init(ctx context.Context, req *InitReq) error
	CreateBet(ctx context.Context, req CreateBetReq) (*CreateBetResp, error)
	UpdateBet(ctx context.Context, req UpdateBetReq) (*UpdateBetResp, error)
	GetAllBets(ctx context.Context, req GetAllBetsReq) (GetAllBetsResp, error)
	GetPlayerBets(ctx context.Context, req GetPlayerBetsReq) (GetPlayerBetsResp, error)
	GetFetchedBet(ctx context.Context, req GetFetchedBetsReq) (GetFetchedBetsResp, error)
	GetResolvedBet(ctx context.Context, req GetResolvedBetsReq) (GetResolvedBetsResp, error)
}

type Bet struct {
	ID             int32     `sql:"id"`
	GameID         int32     `sql:"game_id"`
	PlayerAddress  string    `sql:"player_address"`
	RefAddress     string    `sql:"ref_address"`
	Amount         int64     `sql:"amount"`
	RollUnder      int8      `sql:"roll_under"`
	RandomRoll     int8      `sql:"random_roll"`
	Seed           string    `sql:"seed"`
	Signature      string    `sql:"signature"`
	PlayerPayout   int64     `sql:"player_payout"`
	RefPayout      int64     `sql:"ref_payout"`
	CreatedAt      time.Time `sql:"created_at"`
	CreateTrxHash  string    `sql:"create_trx_hash"`
	CreateTrxLt    int64     `sql:"create_trx_lt"`
	ResolvedAt     time.Time `sql:"resolved_at"`
	ResolveTrxHash string    `sql:"resolve_trx_hash"`
	ResolveTrxLt   int64     `sql:"resolve_trx_lt"`
}

type InitReq struct{}

type CreateBetReq struct {
	GameID        int32  `sql:"game_id"`
	PlayerAddress string `sql:"player_address"`
	RefAddress    string `sql:"ref_address"`
	Amount        int64  `sql:"amount"`
	RollUnder     int8   `sql:"roll_under"`
	Seed          string `sql:"seed"`
	CreateTrxHash string `sql:"create_trx_hash"`
	CreateTrxLt   int64  `sql:"create_trx_lt"`
}

type CreateBetResp struct {
	ID        int64     `sql:"id"`
	CreatedAt time.Time `sql:"created_at"`
}

type UpdateBetReq struct {
	ID             int32  `sql:"id"`
	GameID         int32  `sql:"game_id"`
	RandomRoll     int8   `sql:"random_roll"`
	Signature      string `sql:"signature"`
	PlayerPayout   int64  `sql:"player_payout"`
	RefPayout      int64  `sql:"ref_payout"`
	ResolveTrxHash string `sql:"resolve_trx_hash"`
	ResolveTrxLt   int64  `sql:"resolve_trx_lt"`
}

type UpdateBetResp struct {
	ID         int64     `sql:"id"`
	ResolvedAt time.Time `sql:"resolved_at"`
}

type GetAllBetsReq struct{}

type GetAllBetsResp []*Bet

type GetPlayerBetsReq struct {
	PlayerAddress string `sql:"player_address"`
}

type GetPlayerBetsResp []*Bet

type GetFetchedBetsReq struct {
	GameID        int32  `sql:"game_id"`
	CreateTrxHash string `sql:"create_trx_hash"`
	CreateTrxLt   int64  `sql:"create_trx_lt"`
}

type GetFetchedBetsResp []*Bet

type GetResolvedBetsReq struct {
	GameID         int32  `sql:"game_id"`
	ResolveTrxHash string `sql:"resolve_trx_hash"`
	ResolveTrxLt   int64  `sql:"resolve_trx_lt"`
}

type GetResolvedBetsResp []*Bet
