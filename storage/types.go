package storage

import (
	"context"
	"time"
)

// represent bet state in code
type BetState int

const (
	BetSaved = iota
	BetSent
	BetResolved
)

func BetStateFromInt32(state int32) BetState {
	var out BetState
	switch state {
	case 0:
		out = BetSaved
	case 1:
		out = BetSent
	case 2:
		out = BetResolved
	}
	return out
}

type Store interface {
	Init(ctx context.Context, req *InitReq) error
	SaveBet(ctx context.Context, req SaveBetReq) (*SaveBetResp, error)
	UpdateBet(ctx context.Context, req UpdateBetReq) (*UpdateBetResp, error)
	GetAllBets(ctx context.Context, req GetAllBetsReq) (GetAllBetsResp, error)
	GetPlayerBets(ctx context.Context, req GetPlayerBetsReq) (GetPlayerBetsResp, error)
	GetBet(ctx context.Context, req GetBetReq) (GetBetResp, error)
}

type Bet struct {
	ID             int32     `sql:"id"`
	Amount         int64     `sql:"amount"`
	State          BetState  `sql:"state"`
	RollUnder      int8      `sql:"roll_under"`
	PlayerAddress  string    `sql:"player_address"`
	RefAddress     string    `sql:"ref_address"`
	Seed           string    `sql:"seed"`
	Signature      string    `sql:"signature"`
	RandomRoll     int8      `sql:"random_roll"`
	PlayerPayout   int64     `sql:"player_payout"`
	RefPayout      int64     `sql:"ref_payout"`
	CreatedAt      time.Time `sql:"created_at"`
	CreateTrxHash  string    `sql:"create_trx_hash"`
	CreateTrxLt    int64     `sql:"create_trx_lt"`
	UpdatedAt      time.Time `sql:"updated_at"`
	ResolveTrxHash string    `sql:"resolve_trx_hash"`
	ResolveTrxLt   int64     `sql:"resolve_trx_lt"`
}

type InitReq struct{}

type SaveBetReq struct {
	ID            int32  `sql:"id"`
	PlayerAddress string `sql:"player_address"`
	RefAddress    string `sql:"ref_address"`
	Amount        int64  `sql:"amount"`
	RollUnder     int8   `sql:"roll_under"`
	Seed          string `sql:"seed"`
	CreateTrxHash string `sql:"create_trx_hash"`
	CreateTrxLt   int64  `sql:"create_trx_lt"`
}

type SaveBetResp struct {
	ID        int32     `sql:"id"`
	State     BetState  `sql:"state"`
	UpdatedAt time.Time `sql:"updated_at"`
}

type UpdateBetReq struct {
	ID             int32    `sql:"id"`
	State          BetState `sql:"state"`
	RandomRoll     int8     `sql:"random_roll"`
	Signature      string   `sql:"signature"`
	PlayerPayout   int64    `sql:"player_payout"`
	RefPayout      int64    `sql:"ref_payout"`
	ResolveTrxHash string   `sql:"resolve_trx_hash"`
	ResolveTrxLt   int64    `sql:"resolve_trx_lt"`
}

type UpdateBetResp struct {
	ID        int32     `sql:"id"`
	State     BetState  `sql:"state"`
	UpdatedAt time.Time `sql:"updated_at"`
}

type GetAllBetsReq struct {
	Limit uint64 `sql:"limit"`
}

type GetAllBetsResp []*Bet

type GetPlayerBetsReq struct {
	PlayerAddress string `sql:"player_address"`
	Limit         uint64 `sql:"limit"`
}

type GetPlayerBetsResp []*Bet

type GetBetReq struct {
	ID int32 `sql:"id"`
}

type GetBetResp Bet
