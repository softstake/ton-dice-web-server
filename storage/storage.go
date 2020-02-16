package storage

import (
	"context"
	"time"
)

//go:generate salgen -destination=./client.go -package=github.com/tonradar/ton-dice-web-server/storage github.com/tonradar/ton-dice-web-server/storage Store
type Store interface {
	Init(ctx context.Context, req *InitReq) error
	CreateBet(ctx context.Context, req CreateBetReq) (*CreateBetResp, error)
	GetAllBets(ctx context.Context, req GetAllBetsReq) (GetAllBetsResp, error)
	//BetsByPlayer(ctx context.Context, req BetsByPlayerReq) ([]*BetsByPlayerResp, error)
}

type InitReq struct{}

func (r InitReq) Query() string {
	return `create table IF NOT EXISTS bets (
			id SERIAL PRIMARY KEY,
			game_id INTEGER PRIMARY KEY not null,
			player_address VARCHAR (48) not null,
			ref_address VARCHAR (48) not null,
			amount BIGINT not null,
			roll_under SMALLINT not null,
			random_roll SMALLINT,
			seed text,
			signature text,
			player_payout BIGINT,
			ref_payout BIGINT,
			created_at TIMESTAMP WITH TIME ZONE not null
		)`
}

type GetAllBetsReq struct{}

func (r GetAllBetsReq) Query() string {
	return `SELECT * FROM bets ORDER BY id DESC`
}

type Bet struct {
	ID            int32     `sql:"id"`
	GameID        int32     `sql:"game_id"`
	PlayerAddress string    `sql:"player_address"`
	RefAddress    string    `sql:"ref_address"`
	Amount        int64     `sql:"amount"`
	RollUnder     int8      `sql:"roll_under"`
	RandomRoll    int8      `sql:"random_roll"`
	Seed          string    `sql:"seed"`
	Signature     string    `sql:"signature"`
	PlayerPayout  int64     `sql:"player_payout"`
	RefPayout     int64     `sql:"ref_payout"`
	CreatedAt     time.Time `sql:"created_at"`
}

type GetAllBetsResp []*Bet

type CreateBetReq struct {
	GameID        int32  `sql:"game_id"`
	PlayerAddress string `sql:"player_address"`
	RefAddress    string `sql:"ref_address"`
	Amount        int64  `sql:"amount"`
	RollUnder     int8   `sql:"roll_under"`
	RandomRoll    int8   `sql:"random_roll"`
	Seed          string `sql:"seed"`
	Signature     string `sql:"signature"`
	PlayerPayout  int64  `sql:"player_payout"`
	RefPayout     int64  `sql:"ref_payout"`
}

func (r CreateBetReq) Query() string {
	return `INSERT INTO bets(game_id, player_address, ref_address, amount, roll_under, random_roll, seed, signature, player_payout, ref_payout, created_at) VALUES (@game_id, @player_address, @ref_address, @amount, @roll_under, @random_roll, @seed, @signature, @player_payout, @ref_payout, now()) RETURNING id, created_at`
}

type CreateBetResp struct {
	ID        int64     `sql:"id"`
	CreatedAt time.Time `sql:"created_at"`
}
