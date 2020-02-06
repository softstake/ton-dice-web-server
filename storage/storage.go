package storage

import (
	"context"
	"time"
)

//go:generate salgen -destination=./client.go -package=github.com/tonradar/ton-dice-web-server/storage github.com/tonradar/ton-dice-web-server/storage Store
type Store interface {
	CreateBet(ctx context.Context, req CreateBetReq) (*CreateBetResp, error)
	//BetsByPlayer(ctx context.Context, req BetsByPlayerReq) ([]*BetsByPlayerResp, error)
}

type CreateBetReq struct {
	GameID        int32  `sql:"game_id"`
	PlayerAddress string `sql:"player_address"`
	RefAddress    string `sql:"ref_address"`
	Amount        int64  `sql:"amount"`
	RollUnder     int8   `sql:"roll_under"`
	RandomRoll    int8   `sql:"random_roll"`
	Seed          string `sql:"seed"`
	//Signature     string `sql:"signature"`
	//PlayerPayout  int64  `sql:"player_payout"`
	//RefPayout     int64  `sql:"ref_payout"`
}

func (r CreateBetReq) Query() string {
	return `INSERT INTO bets(game_id, player_address, ref_address, amount, roll_under, random_roll, seed, created_at) VALUES (@game_id, @player_address, @ref_address, @amount, @roll_under, @random_roll, @seed, now()) RETURNING id, created_at`
}

type CreateBetResp struct {
	ID        int64     `sql:"id"`
	CreatedAt time.Time `sql:"created_at"`
}
