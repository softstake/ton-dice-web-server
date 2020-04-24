package storage

//go:generate salgen -destination=./client.go -package=github.com/tonradar/ton-dice-web-server/storage github.com/tonradar/ton-dice-web-server/storage Store

func (r InitReq) Query() string {
	return `create table IF NOT EXISTS bets (
			id INTEGER PRIMARY KEY,
			player_address VARCHAR (48) not null,
			ref_address VARCHAR (48) not null,
			amount BIGINT not null,
			roll_under SMALLINT not null,
			random_roll SMALLINT not null default 0,
			seed TEXT not null,
			signature TEXT not null default '',
			player_payout BIGINT not null default 0,
			ref_payout BIGINT not null default 0,
			created_at TIMESTAMP WITH TIME ZONE not null,
			create_trx_hash TEXT not null,
			create_trx_lt BIGINT not null,
			resolved_at TIMESTAMP WITH TIME ZONE not null,
			resolve_trx_hash TEXT not null default '',
			resolve_trx_lt BIGINT not null default 0
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
