package storage

import (
	"context"
	"time"
)

//go:generate salgen -destination=./client.go -package=github.com/tonradar/ton-dice-web-server/storage github.com/tonradar/ton-dice-web-server/storage Store
type Store interface {
	Init(ctx context.Context, req *InitReq) error
	CreateBet(ctx context.Context, req CreateBetReq) (*CreateBetResp, error)
	UpdateBet(ctx context.Context, req UpdateBetReq) (*UpdateBetResp, error)
	GetAllBets(ctx context.Context, req GetAllBetsReq) (GetAllBetsResp, error)
	GetPlayerBets(ctx context.Context, req GetPlayerBetsReq) (GetPlayerBetsResp, error)
	GetFetchedBet(ctx context.Context, req GetFetchedBetReq) (GetBetResp, error)
	GetResolvedBet(ctx context.Context, req GetFetchedBetReq) (GetBetResp, error)
}

type InitReq struct{}

func (r InitReq) Query() string {
	return `create table IF NOT EXISTS bets (
			id SERIAL not null,
			game_id INTEGER not null,
			player_address VARCHAR (48) not null,
			ref_address VARCHAR (48) not null,
			amount BIGINT not null,
			roll_under SMALLINT not null,
			random_roll SMALLINT,
			seed TEXT,
			signature TEXT,
			player_payout BIGINT,
			ref_payout BIGINT,
			created_at TIMESTAMP WITH TIME ZONE not null,
			create_trx_hash TEXT not null,
			create_trx_lt BIGINT not null,
			resolved_at TIMESTAMP WITH TIME ZONE not null,
			resolve_trx_hash TEXT not null,
			resolve_trx_lt BIGINT not null,
			PRIMARY KEY(id, game_id)
		)`
}

type GetAllBetsReq struct{}

func (r GetAllBetsReq) Query() string {
	return `SELECT * FROM bets ORDER BY id DESC`
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

type GetAllBetsResp []*Bet

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

func (r CreateBetReq) Query() string {
	return `INSERT INTO bets(game_id, player_address, ref_address, amount, roll_under, seed, create_trx_hash, create_trx_lt, created_at) VALUES (@game_id, @player_address, @ref_address, @amount, @roll_under, @seed, @create_trx_hash, @create_trx_lt, now()) RETURNING id, created_at`
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

func (r UpdateBetReq) Query() string {
	return `UPDATE bets SET random_roll=@random_roll, signature=@signature, player_payout=@player_payout, ref_payout=@ref_payout, resolve_trx_hash=@resolve_trx_hash, resolve_trx_lt=@resolve_trx_lt, resolved_at=now() WHERE id=@id AND game_id=@game_id RETURNING id, resolved_at`
}

type UpdateBetResp struct {
	ID         int64     `sql:"id"`
	ResolvedAt time.Time `sql:"resolved_at"`
}

type GetPlayerBetsReq struct {
	PlayerAddress string `sql:"player_address"`
}

func (r *GetPlayerBetsReq) Query() string {
	return "SELECT * FROM bets WHERE player_address=@player_address ORDER BY id DESC"
}

type GetPlayerBetsResp []*Bet

type GetFetchedBetReq struct {
	GameID        int32  `sql:"game_id"`
	CreateTrxHash string `sql:"create_trx_hash"`
	CreateTrxLt   int64  `sql:"create_trx_lt"`
}

func (r *GetFetchedBetReq) Query() string {
	return "SELECT * FROM bets WHERE game_id=@game_id AND create_trx_hash=@create_trx_hash AND create_trx_lt=@create_trx_lt"
}

type GetResolvedBetReq struct {
	GameID         int32  `sql:"game_id"`
	ResolveTrxHash string `sql:"resolve_trx_hash"`
	ResolveTrxLt   int64  `sql:"resolve_trx_lt"`
}

func (r *GetResolvedBetReq) Query() string {
	return "SELECT * FROM bets WHERE game_id=@game_id AND resolve_trx_hash=@resolve_trx_hash AND resolve_trx_lt=@resolve_trx_lt"
}

type GetBetResp []*Bet
