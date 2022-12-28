package usecase

import (
	"context"
	"github.com/jackc/pgx/v5"
	"tokowiwin/repositories/db"
	"tokowiwin/usecases"
)

type UCDelete struct{}
type usecaseCartsDelete struct {
	ctx  context.Context
	repo db.RepositoryI
}

type requestDelete struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}

type responseDelete struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

func (c UCDelete) NewUsecase(ctx context.Context, repo db.RepositoryI) *usecaseCartsDelete {
	return &usecaseCartsDelete{
		ctx:  ctx,
		repo: repo,
	}
}

func (u usecaseCartsDelete) HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error) {
	var (
		err  error
		req  = new(requestDelete)
		resp = new(responseDelete)
	)

	if err = data.HTTPData.BodyParser(req); err != nil {
		return nil, err
	}
	err = data.Validator.Struct(*req)
	if err != nil {
		return nil, err
	}
	err = db.ExecuteWithTx(ctx, data.TxExecutor, func(tx pgx.Tx) error {
		err = u.repo.DeleteCart(ctx, tx, req.ID, req.UserID)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	resp.Success = 1
	resp.Message = "Berhasil"

	return resp, nil
}
