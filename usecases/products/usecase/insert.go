package usecase

import (
	"context"
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
	ProductName  string `json:"product_name" validate:"required"`
	ProductPrice int    `json:"product_price" validate:"required"`
	ProductStock int    `json:"product_stock" validate:"required"`
	ProductPic   string `json:"product_pic" validate:"required"`
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

	err = db.ExecuteWithTx(ctx, data.TxExecutor, func(tx pgx.Tx) error {
		err = u.repo.InsertProduct(ctx, tx, &model.Products{
			ProductName:  req.ProductName,
			ProductStock: req.ProductStock,
			ProductPrice: req.ProductPrice,
			ProductPic:   req.ProductPic,
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
