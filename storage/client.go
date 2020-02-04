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
func (s *SalStore) CreateBet(ctx context.Context, req CreateBetReq) (*CreateBetResp, error) {
	var (
		err      error
		rawQuery = req.Query()
		reqMap   = make(sal.RowMap)
	)
	reqMap.AppendTo("game_id", &req.GameID)
	reqMap.AppendTo("player_address", &req.PlayerAddress)
	reqMap.AppendTo("ref_address", &req.RefAddress)
	reqMap.AppendTo("amount", &req.Amount)
	reqMap.AppendTo("roll_under", &req.RollUnder)
	reqMap.AppendTo("random_roll", &req.RandomRoll)
	reqMap.AppendTo("seed", &req.Seed)
	reqMap.AppendTo("signature", &req.Signature)
	reqMap.AppendTo("player_payout", &req.PlayerPayout)
	reqMap.AppendTo("ref_payout", &req.RefPayout)

	ctx = context.WithValue(ctx, sal.ContextKeyTxOpened, s.txOpened)
	ctx = context.WithValue(ctx, sal.ContextKeyOperationType, "QueryRow")
	ctx = context.WithValue(ctx, sal.ContextKeyMethodName, "CreateBet")

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

	var resp CreateBetResp
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

// compile time checks
var _ Store = &SalStore{}
