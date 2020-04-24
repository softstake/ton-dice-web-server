// Code generated by SalGen. DO NOT EDIT.
package storage

import (
	"context"
	"database/sql"
	"github.com/go-gad/sal"
	"github.com/pkg/errors"
)

type SalStore struct {
	Store
	handler  sal.QueryHandler
	parent   sal.QueryHandler
	ctrl     *sal.Controller
	txOpened bool
}

func NewStore(h sal.QueryHandler, options ...sal.ClientOption) *SalStore {
	s := &SalStore{
		handler:  h,
		ctrl:     sal.NewController(options...),
		txOpened: false,
	}

	return s
}

func (s *SalStore) BeginTx(ctx context.Context, opts *sql.TxOptions) (Store, error) {
	dbConn, ok := s.handler.(sal.TransactionBegin)
	if !ok {
		return nil, errors.New("handler doesn't satisfy the interface TransactionBegin")
	}
	var (
		err error
		tx  *sql.Tx
	)

	ctx = context.WithValue(ctx, sal.ContextKeyTxOpened, s.txOpened)
	ctx = context.WithValue(ctx, sal.ContextKeyOperationType, "Begin")
	ctx = context.WithValue(ctx, sal.ContextKeyMethodName, "BeginTx")

	for _, fn := range s.ctrl.BeforeQuery {
		var fnz sal.FinalizerFunc
		ctx, fnz = fn(ctx, "BEGIN", nil)
		if fnz != nil {
			defer func() { fnz(ctx, err) }()
		}
	}

	tx, err = dbConn.BeginTx(ctx, opts)
	if err != nil {
		err = errors.Wrap(err, "failed to start tx")
		return nil, err
	}

	newClient := &SalStore{
		handler:  tx,
		parent:   s.handler,
		ctrl:     s.ctrl,
		txOpened: true,
	}

	return newClient, nil
}

func (s *SalStore) Tx() sal.Transaction {
	if tx, ok := s.handler.(sal.SqlTx); ok {
		return sal.NewWrappedTx(tx, s.ctrl)
	}
	return nil
}
func (s *SalStore) GetAllBets(ctx context.Context, req GetAllBetsReq) (GetAllBetsResp, error) {
	var (
		err      error
		rawQuery = req.Query()
		reqMap   = make(sal.RowMap)
	)
	reqMap.AppendTo("limit", &req.Limit)

	ctx = context.WithValue(ctx, sal.ContextKeyTxOpened, s.txOpened)
	ctx = context.WithValue(ctx, sal.ContextKeyOperationType, "Query")
	ctx = context.WithValue(ctx, sal.ContextKeyMethodName, "GetAllBets")

	pgQuery, args := sal.ProcessQueryAndArgs(rawQuery, reqMap)

	stmt, err := s.ctrl.PrepareStmt(ctx, s.parent, s.handler, pgQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, fn := range s.ctrl.BeforeQuery {
		var fnz sal.FinalizerFunc
		ctx, fnz = fn(ctx, rawQuery, req)
		if fnz != nil {
			defer func() { fnz(ctx, err) }()
		}
	}

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute Query")
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch columns")
	}

	var list = make(GetAllBetsResp, 0)

	for rows.Next() {
		var resp Bet
		var respMap = make(sal.RowMap)
		respMap.AppendTo("id", &resp.ID)
		respMap.AppendTo("amount", &resp.Amount)
		respMap.AppendTo("roll_under", &resp.RollUnder)
		respMap.AppendTo("player_address", &resp.PlayerAddress)
		respMap.AppendTo("ref_address", &resp.RefAddress)
		respMap.AppendTo("seed", &resp.Seed)
		respMap.AppendTo("created_at", &resp.CreatedAt)
		respMap.AppendTo("create_trx_hash", &resp.CreateTrxHash)
		respMap.AppendTo("create_trx_lt", &resp.CreateTrxLt)
		respMap.AppendTo("signature", &resp.Signature)
		respMap.AppendTo("random_roll", &resp.RandomRoll)
		respMap.AppendTo("player_payout", &resp.PlayerPayout)
		respMap.AppendTo("ref_payout", &resp.RefPayout)
		respMap.AppendTo("resolved_at", &resp.ResolvedAt)
		respMap.AppendTo("resolve_trx_hash", &resp.ResolveTrxHash)
		respMap.AppendTo("resolve_trx_lt", &resp.ResolveTrxLt)

		dest := sal.GetDests(cols, respMap)

		if err = rows.Scan(dest...); err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}

		list = append(list, &resp)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "something failed during iteration")
	}

	return list, nil
}

func (s *SalStore) GetBet(ctx context.Context, req GetBetReq) (GetBetResp, error) {
	var (
		err      error
		rawQuery = req.Query()
		reqMap   = make(sal.RowMap)
	)
	reqMap.AppendTo("id", &req.ID)

	ctx = context.WithValue(ctx, sal.ContextKeyTxOpened, s.txOpened)
	ctx = context.WithValue(ctx, sal.ContextKeyOperationType, "Query")
	ctx = context.WithValue(ctx, sal.ContextKeyMethodName, "GetBet")

	pgQuery, args := sal.ProcessQueryAndArgs(rawQuery, reqMap)

	stmt, err := s.ctrl.PrepareStmt(ctx, s.parent, s.handler, pgQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, fn := range s.ctrl.BeforeQuery {
		var fnz sal.FinalizerFunc
		ctx, fnz = fn(ctx, rawQuery, req)
		if fnz != nil {
			defer func() { fnz(ctx, err) }()
		}
	}

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute Query")
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch columns")
	}

	var list = make(GetBetResp, 0)

	for rows.Next() {
		var resp Bet
		var respMap = make(sal.RowMap)
		respMap.AppendTo("id", &resp.ID)
		respMap.AppendTo("amount", &resp.Amount)
		respMap.AppendTo("roll_under", &resp.RollUnder)
		respMap.AppendTo("player_address", &resp.PlayerAddress)
		respMap.AppendTo("ref_address", &resp.RefAddress)
		respMap.AppendTo("seed", &resp.Seed)
		respMap.AppendTo("created_at", &resp.CreatedAt)
		respMap.AppendTo("create_trx_hash", &resp.CreateTrxHash)
		respMap.AppendTo("create_trx_lt", &resp.CreateTrxLt)
		respMap.AppendTo("signature", &resp.Signature)
		respMap.AppendTo("random_roll", &resp.RandomRoll)
		respMap.AppendTo("player_payout", &resp.PlayerPayout)
		respMap.AppendTo("ref_payout", &resp.RefPayout)
		respMap.AppendTo("resolved_at", &resp.ResolvedAt)
		respMap.AppendTo("resolve_trx_hash", &resp.ResolveTrxHash)
		respMap.AppendTo("resolve_trx_lt", &resp.ResolveTrxLt)

		dest := sal.GetDests(cols, respMap)

		if err = rows.Scan(dest...); err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}

		list = append(list, &resp)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "something failed during iteration")
	}

	return list, nil
}

func (s *SalStore) GetPlayerBets(ctx context.Context, req GetPlayerBetsReq) (GetPlayerBetsResp, error) {
	var (
		err      error
		rawQuery = req.Query()
		reqMap   = make(sal.RowMap)
	)
	reqMap.AppendTo("player_address", &req.PlayerAddress)
	reqMap.AppendTo("limit", &req.Limit)

	ctx = context.WithValue(ctx, sal.ContextKeyTxOpened, s.txOpened)
	ctx = context.WithValue(ctx, sal.ContextKeyOperationType, "Query")
	ctx = context.WithValue(ctx, sal.ContextKeyMethodName, "GetPlayerBets")

	pgQuery, args := sal.ProcessQueryAndArgs(rawQuery, reqMap)

	stmt, err := s.ctrl.PrepareStmt(ctx, s.parent, s.handler, pgQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, fn := range s.ctrl.BeforeQuery {
		var fnz sal.FinalizerFunc
		ctx, fnz = fn(ctx, rawQuery, req)
		if fnz != nil {
			defer func() { fnz(ctx, err) }()
		}
	}

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute Query")
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch columns")
	}

	var list = make(GetPlayerBetsResp, 0)

	for rows.Next() {
		var resp Bet
		var respMap = make(sal.RowMap)
		respMap.AppendTo("id", &resp.ID)
		respMap.AppendTo("amount", &resp.Amount)
		respMap.AppendTo("roll_under", &resp.RollUnder)
		respMap.AppendTo("player_address", &resp.PlayerAddress)
		respMap.AppendTo("ref_address", &resp.RefAddress)
		respMap.AppendTo("seed", &resp.Seed)
		respMap.AppendTo("created_at", &resp.CreatedAt)
		respMap.AppendTo("create_trx_hash", &resp.CreateTrxHash)
		respMap.AppendTo("create_trx_lt", &resp.CreateTrxLt)
		respMap.AppendTo("signature", &resp.Signature)
		respMap.AppendTo("random_roll", &resp.RandomRoll)
		respMap.AppendTo("player_payout", &resp.PlayerPayout)
		respMap.AppendTo("ref_payout", &resp.RefPayout)
		respMap.AppendTo("resolved_at", &resp.ResolvedAt)
		respMap.AppendTo("resolve_trx_hash", &resp.ResolveTrxHash)
		respMap.AppendTo("resolve_trx_lt", &resp.ResolveTrxLt)

		dest := sal.GetDests(cols, respMap)

		if err = rows.Scan(dest...); err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}

		list = append(list, &resp)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "something failed during iteration")
	}

	return list, nil
}

func (s *SalStore) Init(ctx context.Context, req *InitReq) error {
	var (
		err      error
		rawQuery = req.Query()
		reqMap   = make(sal.RowMap)
	)

	ctx = context.WithValue(ctx, sal.ContextKeyTxOpened, s.txOpened)
	ctx = context.WithValue(ctx, sal.ContextKeyOperationType, "Exec")
	ctx = context.WithValue(ctx, sal.ContextKeyMethodName, "Init")

	pgQuery, args := sal.ProcessQueryAndArgs(rawQuery, reqMap)

	stmt, err := s.ctrl.PrepareStmt(ctx, s.parent, s.handler, pgQuery)
	if err != nil {
		return errors.WithStack(err)
	}

	for _, fn := range s.ctrl.BeforeQuery {
		var fnz sal.FinalizerFunc
		ctx, fnz = fn(ctx, rawQuery, req)
		if fnz != nil {
			defer func() { fnz(ctx, err) }()
		}
	}

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return errors.Wrap(err, "failed to execute Exec")
	}

	return nil
}

func (s *SalStore) SaveBet(ctx context.Context, req SaveBetReq) (*SaveBetResp, error) {
	var (
		err      error
		rawQuery = req.Query()
		reqMap   = make(sal.RowMap)
	)
	reqMap.AppendTo("id", &req.ID)
	reqMap.AppendTo("player_address", &req.PlayerAddress)
	reqMap.AppendTo("ref_address", &req.RefAddress)
	reqMap.AppendTo("amount", &req.Amount)
	reqMap.AppendTo("roll_under", &req.RollUnder)
	reqMap.AppendTo("seed", &req.Seed)
	reqMap.AppendTo("create_trx_hash", &req.CreateTrxHash)
	reqMap.AppendTo("create_trx_lt", &req.CreateTrxLt)

	ctx = context.WithValue(ctx, sal.ContextKeyTxOpened, s.txOpened)
	ctx = context.WithValue(ctx, sal.ContextKeyOperationType, "QueryRow")
	ctx = context.WithValue(ctx, sal.ContextKeyMethodName, "SaveBet")

	pgQuery, args := sal.ProcessQueryAndArgs(rawQuery, reqMap)

	stmt, err := s.ctrl.PrepareStmt(ctx, s.parent, s.handler, pgQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, fn := range s.ctrl.BeforeQuery {
		var fnz sal.FinalizerFunc
		ctx, fnz = fn(ctx, rawQuery, req)
		if fnz != nil {
			defer func() { fnz(ctx, err) }()
		}
	}

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute Query")
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch columns")
	}

	if !rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, errors.Wrap(err, "rows error")
		}
		return nil, sql.ErrNoRows
	}

	var resp SaveBetResp
	var respMap = make(sal.RowMap)
	respMap.AppendTo("id", &resp.ID)
	respMap.AppendTo("created_at", &resp.CreatedAt)

	dest := sal.GetDests(cols, respMap)

	if err = rows.Scan(dest...); err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "something failed during iteration")
	}

	return &resp, nil
}

func (s *SalStore) UpdateBet(ctx context.Context, req UpdateBetReq) (*UpdateBetResp, error) {
	var (
		err      error
		rawQuery = req.Query()
		reqMap   = make(sal.RowMap)
	)
	reqMap.AppendTo("id", &req.ID)
	reqMap.AppendTo("random_roll", &req.RandomRoll)
	reqMap.AppendTo("signature", &req.Signature)
	reqMap.AppendTo("player_payout", &req.PlayerPayout)
	reqMap.AppendTo("ref_payout", &req.RefPayout)
	reqMap.AppendTo("resolved_at", &req.ResolvedAt)
	reqMap.AppendTo("resolve_trx_hash", &req.ResolveTrxHash)
	reqMap.AppendTo("resolve_trx_lt", &req.ResolveTrxLt)

	ctx = context.WithValue(ctx, sal.ContextKeyTxOpened, s.txOpened)
	ctx = context.WithValue(ctx, sal.ContextKeyOperationType, "QueryRow")
	ctx = context.WithValue(ctx, sal.ContextKeyMethodName, "UpdateBet")

	pgQuery, args := sal.ProcessQueryAndArgs(rawQuery, reqMap)

	stmt, err := s.ctrl.PrepareStmt(ctx, s.parent, s.handler, pgQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, fn := range s.ctrl.BeforeQuery {
		var fnz sal.FinalizerFunc
		ctx, fnz = fn(ctx, rawQuery, req)
		if fnz != nil {
			defer func() { fnz(ctx, err) }()
		}
	}

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute Query")
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch columns")
	}

	if !rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, errors.Wrap(err, "rows error")
		}
		return nil, sql.ErrNoRows
	}

	var resp UpdateBetResp
	var respMap = make(sal.RowMap)
	respMap.AppendTo("id", &resp.ID)
	respMap.AppendTo("resolved_at", &resp.ResolvedAt)

	dest := sal.GetDests(cols, respMap)

	if err = rows.Scan(dest...); err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "something failed during iteration")
	}

	return &resp, nil
}

// compile time checks
var _ Store = &SalStore{}
