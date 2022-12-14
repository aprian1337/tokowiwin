package usecase

import (
	"context"
	"github.com/jackc/pgx/v5"
	"tokowiwin/repositories/db"
	"tokowiwin/repositories/model"
	"tokowiwin/usecases"
)

type UCUpdate struct{}
type usecaseProductsUpdate struct {
	ctx  context.Context
	repo db.RepositoryI
}

type requestUpdate struct {
	ID           int64  `json:"id"`
	ProductName  string `json:"product_name"`
	ProductPic   string `json:"product_pic"`
	ProductStock int    `json:"product_stock"`
	ProductPrice int    `json:"product_price"`
}

type responseUpdate struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

func (c UCUpdate) NewUsecase(ctx context.Context, repo db.RepositoryI) *usecaseProductsUpdate {
	return &usecaseProductsUpdate{
		ctx:  ctx,
		repo: repo,
	}
}

func (u usecaseProductsUpdate) HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error) {
	var (
		err  error
		req  = new(requestUpdate)
		resp = new(responseUpdate)
	)

	if err = data.HTTPData.BodyParser(req); err != nil {
		return nil, err
	}
	err = data.Validator.Struct(*req)
	if err != nil {
		return nil, err
	}
	err = db.ExecuteWithTx(ctx, data.TxExecutor, func(tx pgx.Tx) error {
		err = u.repo.UpdateProduct(ctx, tx, &model.Products{
			ID:           req.ID,
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
