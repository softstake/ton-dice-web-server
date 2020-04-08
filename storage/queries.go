package storage

//go:generate salgen -destination=./client.go -package=github.com/tonradar/ton-dice-web-server/storage github.com/tonradar/ton-dice-web-server/storage Store

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
			resolved_at TIMESTAMP WITH TIME ZONE,
			resolve_trx_hash TEXT,
			resolve_trx_lt BIGINT,
			PRIMARY KEY(id, game_id)
		)`
}

func (r GetAllBetsReq) Query() string {
	return `SELECT * FROM bets ORDER BY id DESC`
}

func (r CreateBetReq) Query() string {
	return `INSERT INTO bets(game_id, player_address, ref_address, amount, roll_under, seed, create_trx_hash, create_trx_lt, created_at) VALUES (@game_id, @player_address, @ref_address, @amount, @roll_under, @seed, @create_trx_hash, @create_trx_lt, now()) RETURNING id, created_at`
}

func (r UpdateBetReq) Query() string {
	return `UPDATE bets SET random_roll=@random_roll, signature=@signature, player_payout=@player_payout, ref_payout=@ref_payout, resolve_trx_hash=@resolve_trx_hash, resolve_trx_lt=@resolve_trx_lt, resolved_at=now() WHERE id=@id AND game_id=@game_id RETURNING id, resolved_at`
}

func (r *GetPlayerBetsReq) Query() string {
	return "SELECT * FROM bets WHERE player_address=@player_address ORDER BY id DESC"
}

func (r *GetFetchedBetsReq) Query() string {
	return "SELECT * FROM bets WHERE game_id=@game_id AND create_trx_hash=@create_trx_hash AND create_trx_lt=@create_trx_lt ORDER BY id DESC"
}

func (r *GetResolvedBetsReq) Query() string {
	return "SELECT * FROM bets WHERE game_id=@game_id AND resolve_trx_hash=@resolve_trx_hash AND resolve_trx_lt=@resolve_trx_lt ORDER BY id DESC"
}
