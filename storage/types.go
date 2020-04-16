package storage

import (
	"context"
	"database/sql"
	"time"
)

type Store interface {
	Init(ctx context.Context, req *InitReq) error
	CreateBet(ctx context.Context, req CreateBetReq) (*CreateBetResp, error)
	UpdateBet(ctx context.Context, req UpdateBetReq) (*UpdateBetResp, error)
	GetAllBets(ctx context.Context, req GetAllBetsReq) (GetAllBetsResp, error)
	GetPlayerBets(ctx context.Context, req GetPlayerBetsReq) (GetPlayerBetsResp, error)
	GetBet(ctx context.Context, req GetBetReq) (GetBetResp, error)
}

type Bet struct {
	ID            int32     `sql:"id"`
	Amount        int64     `sql:"amount"`
	RollUnder     int8      `sql:"roll_under"`
	PlayerAddress string    `sql:"player_address"`
	RefAddress    string    `sql:"ref_address"`
	Seed          string    `sql:"seed"`
	CreatedAt     time.Time `sql:"created_at"`
	CreateTrxHash string    `sql:"create_trx_hash"`
	CreateTrxLt   int64     `sql:"create_trx_lt"`

	Signature      sql.NullString `sql:"signature"`
	RandomRoll     sql.NullInt32  `sql:"random_roll"`
	PlayerPayout   sql.NullInt64  `sql:"player_payout"`
	RefPayout      sql.NullInt64  `sql:"ref_payout"`
	ResolvedAt     sql.NullTime   `sql:"resolved_at"`
	ResolveTrxHash sql.NullString `sql:"resolve_trx_hash"`
	ResolveTrxLt   sql.NullInt64  `sql:"resolve_trx_lt"`
}

type InitReq struct{}

type CreateBetReq struct {
	ID            int32  `sql:"id"`
	PlayerAddress string `sql:"player_address"`
	RefAddress    string `sql:"ref_address"`
	Amount        int64  `sql:"amount"`
	RollUnder     int8   `sql:"roll_under"`
	Seed          string `sql:"seed"`
	CreateTrxHash string `sql:"create_trx_hash"`
	CreateTrxLt   int64  `sql:"create_trx_lt"`
}

type CreateBetResp struct {
	ID        int32     `sql:"id"`
	CreatedAt time.Time `sql:"created_at"`
}

type UpdateBetReq struct {
	ID             int32  `sql:"id"`
	RandomRoll     int8   `sql:"random_roll"`
	Signature      string `sql:"signature"`
	PlayerPayout   int64  `sql:"player_payout"`
	RefPayout      int64  `sql:"ref_payout"`
	ResolveTrxHash string `sql:"resolve_trx_hash"`
	ResolveTrxLt   int64  `sql:"resolve_trx_lt"`
}

type UpdateBetResp struct {
	ID         int32     `sql:"id"`
	ResolvedAt time.Time `sql:"resolved_at"`
}

type GetAllBetsReq struct{}

type GetAllBetsResp []*Bet

type GetPlayerBetsReq struct {
	PlayerAddress string `sql:"player_address"`
}

type GetPlayerBetsResp []*Bet

type GetBetReq struct {
	ID int32 `sql:"id"`
}

type GetBetResp []*Bet
