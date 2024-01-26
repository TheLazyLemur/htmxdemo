package service

import (
	"context"
	"database/sql"
	"htmxdemo/db/queries"
	"htmxdemo/types"
	"time"
)

type IService interface {
	InsertTransaction(ctx context.Context, req types.Transaction) (err error)
	ListTransactions(ctx context.Context) (txs []types.Transaction, err error)
	SearchTransactions(ctx context.Context, searchValInt int64) (txs []types.Transaction, err error)
	UpdateTransaction(ctx context.Context, id int64) (tx types.Transaction, err error)
}

type Service struct {
	dbc     *sql.DB
	querier queries.Querier
}

func New(dbc *sql.DB, querier queries.Querier) IService {
	return &Service{
		dbc:     dbc,
		querier: querier,
	}
}

func (o *Service) InsertTransaction(ctx context.Context, req types.Transaction) (err error) {
	tx, _ := o.dbc.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	err = o.querier.InsertTransaction(ctx, tx, queries.InsertTransactionParams{
		FromAccount: sql.NullInt64{
			Int64: req.From,
			Valid: true,
		},
		ToAccount: sql.NullInt64{
			Int64: req.To,
			Valid: true,
		},
		Amount: sql.NullInt64{
			Int64: req.Amount,
			Valid: true,
		},
		Status: sql.NullInt64{
			Int64: int64(req.Status),
			Valid: true,
		},
		Flagged: sql.NullBool{
			Bool:  req.Flagged,
			Valid: true,
		},
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (o *Service) ListTransactions(ctx context.Context) (txs []types.Transaction, err error) {
	tx, _ := o.dbc.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	dbResult, err := o.querier.ListTransactions(ctx, tx)
	if err != nil {
		return nil, err
	}

	transactions := []types.Transaction{}
	for _, tx := range dbResult {
		t := types.Transaction{
			ID:      tx.ID,
			From:    tx.FromAccount.Int64,
			To:      tx.ToAccount.Int64,
			Amount:  tx.Amount.Int64,
			Status:  int(tx.Status.Int64),
			Flagged: tx.Flagged.Bool,
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}

func (o *Service) SearchTransactions(ctx context.Context, searchValInt int64) (txs []types.Transaction, err error) {
	tx, _ := o.dbc.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	dbResult, err := o.querier.ListTransactions(ctx, tx)
	if err != nil {
		return nil, err
	}

	transactions := []types.Transaction{}
	for _, tx := range dbResult {
		t := types.Transaction{
			ID:      tx.ID,
			From:    tx.FromAccount.Int64,
			To:      tx.ToAccount.Int64,
			Amount:  tx.Amount.Int64,
			Status:  int(tx.Status.Int64),
			Flagged: tx.Flagged.Bool,
		}
		transactions = append(transactions, t)
	}

	hasMatch := false
	final := []types.Transaction{}
	for i := 0; i < len(transactions); i++ {
		if transactions[i].ID == searchValInt {
			final = append(final, transactions[i])
			hasMatch = true
		}
	}

	if !hasMatch {
		final = transactions
	}

	return final, nil
}

func (o *Service) UpdateTransaction(ctx context.Context, id int64) (txn types.Transaction, err error) {
	tx, _ := o.dbc.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	dbResult, err := o.querier.GetTransaction(ctx, tx, id)
	if err != nil {
		return types.Transaction{}, err
	}

	transaction := types.Transaction{
		ID:      dbResult.ID,
		From:    dbResult.FromAccount.Int64,
		To:      dbResult.ToAccount.Int64,
		Amount:  dbResult.Amount.Int64,
		Status:  int(dbResult.Status.Int64),
		Flagged: dbResult.Flagged.Bool,
	}

	if transaction.Flagged {
		err = o.querier.UnflagTransaction(ctx, tx, transaction.ID)
	} else {
		err = o.querier.FlagTransaction(ctx, tx, transaction.ID)
	}
	if err != nil {
		return types.Transaction{}, err
	}

	transaction.Flagged = !transaction.Flagged
	return transaction, nil
}
