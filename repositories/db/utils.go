package db

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Tx func(tx pgx.Tx) error

type TxExecutor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

func ExecuteWithTx(ctx context.Context, executor TxExecutor, fn Tx) error {
	var (
		err error
	)

	tx, err := executor.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	err = fn(tx)
	if err == nil {
		err = tx.Commit(ctx)
		if err != nil {
			return err
		}
	} else {
		_ = tx.Rollback(ctx)
		return err
	}
	return nil
}
