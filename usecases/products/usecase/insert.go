package usecase

import (
	"context"
	"github.com/jackc/pgx/v5"
	"tokowiwin/repositories/db"
	"tokowiwin/repositories/model"
	"tokowiwin/usecases"
)

type UCInsert struct{}
type usecaseProductsInsert struct {
	ctx  context.Context
	repo db.RepositoryI
}

type requestInsert struct {
	ProductName  string `json:"product_name"`
	ProductPic   string `json:"product_pic"`
	ProductStock int    `json:"product_stock"`
	ProductPrice int    `json:"product_price"`
}

type responseInsert struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

func (c UCInsert) NewUsecase(ctx context.Context, repo db.RepositoryI) *usecaseProductsInsert {
	return &usecaseProductsInsert{
		ctx:  ctx,
		repo: repo,
	}
}

func (u usecaseProductsInsert) HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error) {
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
