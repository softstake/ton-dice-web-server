package storage

//go:generate salgen -destination=./client.go -package=github.com/tonradar/ton-dice-web-server/storage github.com/tonradar/ton-dice-web-server/storage Store

func (r InitReq) Query() string {
	return `create table IF NOT EXISTS bets (
			id INTEGER PRIMARY KEY,
			player_address VARCHAR (48) NOT NULL,
			ref_address VARCHAR (48) NOT NULL,
			amount BIGINT NOT NULL,
			roll_under SMALLINT NOT NULL,
			random_roll SMALLINT NOT NULL DEFAULT 0,
			seed TEXT NOT NULL,
			signature TEXT NOT NULL DEFAULT '',
			player_payout BIGINT NOT NULL DEFAULT 0,
			ref_payout BIGINT NOT NULL DEFAULT 0,
			created_at TIMESTAMP WITH TIME ZONE not null,
			create_trx_hash TEXT NOT NULL,
			create_trx_lt BIGINT NOT NULL,
			resolved_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
			resolve_trx_hash TEXT NOT NULL DEFAULT '',
			resolve_trx_lt BIGINT NOT NULL DEFAULT 0
		)`
}

func (r SaveBetReq) Query() string {
	return `INSERT INTO bets(id, player_address, ref_address, amount, roll_under, seed, create_trx_hash, create_trx_lt, created_at) VALUES (@id, @player_address, @ref_address, @amount, @roll_under, @seed, @create_trx_hash, @create_trx_lt, now()) RETURNING id, created_at`
}

func (r UpdateBetReq) Query() string {
	return `UPDATE bets SET random_roll=@random_roll, signature=@signature, player_payout=@player_payout, ref_payout=@ref_payout, resolve_trx_hash=@resolve_trx_hash, resolve_trx_lt=@resolve_trx_lt, resolved_at=@resolved_at WHERE id=@id RETURNING id, resolved_at`
}

func (r GetAllBetsReq) Query() string {
	return `SELECT * FROM bets ORDER BY id DESC limit @limit`
}

func (r *GetPlayerBetsReq) Query() string {
	return "SELECT * FROM bets WHERE player_address=@player_address ORDER BY id DESC limit @limit"
}

func (r *GetBetReq) Query() string {
	return "SELECT * FROM bets WHERE id=@id LIMIT 1"
}
