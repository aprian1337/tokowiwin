package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"tokowiwin/repositories/db"
	"tokowiwin/repositories/model"
	"tokowiwin/usecases"
)

type UCUpdate struct{}
type usecaseCartssUpdate struct {
	ctx  context.Context
	repo db.RepositoryI
}

type requestUpdate struct {
	ProductID int64 `query:"product_id"`
	UserID    int64 `query:"user_id"`
	Qty       int   `query:"qty"`
}

type responseUpdate struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

func (c UCUpdate) NewUsecase(ctx context.Context, repo db.RepositoryI) *usecaseCartssUpdate {
	return &usecaseCartssUpdate{
		ctx:  ctx,
		repo: repo,
	}
}

func (u usecaseCartssUpdate) HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error) {
	var (
		err  error
		req  = new(requestUpdate)
		resp = new(responseUpdate)
	)

	if err = data.HTTPData.QueryParser(req); err != nil {
		return nil, err
	}
	err = data.Validator.Struct(*req)

	product, err := u.repo.GetProductsByID(ctx, req.ProductID)
	if err != nil {
		return nil, err
	}
	if req.Qty > product.ProductStock {
		resp.Success = 0
		resp.Message = fmt.Sprintf("Gagal, stock barang yang tersedia : %v", product.ProductStock)
		return resp, nil
	}

	if err != nil {
		return nil, err
	}
	err = db.ExecuteWithTx(ctx, data.TxExecutor, func(tx pgx.Tx) error {
		err = u.repo.UpdateCart(ctx, tx, &model.Carts{
			ProductID: sql.NullInt64{
				Int64: req.ProductID,
				Valid: true,
			},
			UserID: req.UserID,
			Qty: sql.NullInt64{
				Int64: int64(req.Qty),
				Valid: true,
			},
		})
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
