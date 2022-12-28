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

type UCInsert struct{}
type usecaseCartssInsert struct {
	ctx  context.Context
	repo db.RepositoryI
}

type requestInsert struct {
	ProductID int64 `json:"product_id"`
	UserID    int64 `json:"user_id"`
	Qty       int64 `json:"qty"`
}

type responseInsert struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

func (c UCInsert) NewUsecase(ctx context.Context, repo db.RepositoryI) *usecaseCartssInsert {
	return &usecaseCartssInsert{
		ctx:  ctx,
		repo: repo,
	}
}

func (u usecaseCartssInsert) HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error) {
	var (
		err  error
		req  = new(requestInsert)
		resp = new(responseInsert)
	)

	if err = data.HTTPData.BodyParser(req); err != nil {
		return nil, err
	}
	err = data.Validator.Struct(*req)
	if err != nil {
		return nil, err
	}
	_, err = u.repo.GetProductsByID(ctx, req.ProductID)
	if err != nil {
		return nil, fmt.Errorf("[GetProductsByID] %v", err)
	}

	carts, err := u.repo.GetCart(ctx, req.UserID)
	for _, v := range carts {
		if v.ProductID.Valid && v.ProductID.Int64 == req.ProductID {
			resp.Success = 0
			resp.Message = "Barang sudah ada di keranjang"
			return resp, nil
		}
	}

	err = db.ExecuteWithTx(ctx, data.TxExecutor, func(tx pgx.Tx) error {
		err = u.repo.InsertCart(ctx, tx, &model.Carts{
			ProductID: sql.NullInt64{
				Int64: req.ProductID,
				Valid: true,
			},
			UserID: req.UserID,
			Qty: sql.NullInt64{
				Int64: req.Qty,
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
