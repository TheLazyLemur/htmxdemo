// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package queries

import (
	"context"
)

type Querier interface {
	FlagTransaction(ctx context.Context, db DBTX, id int64) error
	GetTransaction(ctx context.Context, db DBTX, id int64) (Transaction, error)
	InsertTransaction(ctx context.Context, db DBTX, arg InsertTransactionParams) error
	ListTransactions(ctx context.Context, db DBTX) ([]Transaction, error)
	UnflagTransaction(ctx context.Context, db DBTX, id int64) error
}

var _ Querier = (*Queries)(nil)
