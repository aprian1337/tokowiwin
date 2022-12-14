package usecase

import (
	"context"
	"tokowiwin/repositories/db"
	"tokowiwin/repositories/model"
	"tokowiwin/usecases"
	"tokowiwin/utils/formatter"
)

type UCGet struct{}
type usecaseProductsGet struct {
	ctx  context.Context
	repo db.RepositoryI
}

type requestGet struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type responseGet struct {
	ID           int64  `json:"id"`
	ProductName  string `json:"product_name"`
	ProductStock int    `json:"product_stock"`
	ProductPrice string `json:"product_price"`
	ProductPic   string `json:"product_pic"`
}

func (c UCGet) NewUsecase(ctx context.Context, repo db.RepositoryI) *usecaseProductsGet {
	return &usecaseProductsGet{
		ctx:  ctx,
		repo: repo,
	}
}

func (u usecaseProductsGet) HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error) {
	var (
		req = new(requestGet)
		err error
	)

	if data.HTTPData.Query("id") != "" {
		if err = data.HTTPData.QueryParser(req); err != nil {
			return nil, err
		}
		err = data.Validator.Struct(*req)
		if err != nil {
			return nil, err
		}
	}

	if req.ID == 0 {
		m, err := u.repo.GetProductsAll(ctx)
		if err != nil {
			return nil, err
		}
		return u.buildArrayResponse(m), nil
	}

	m, err := u.repo.GetProductsByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return u.buildResponse(m), nil
}

func (u usecaseProductsGet) buildResponse(m *model.Products) *responseGet {
	return &responseGet{
		ID:           m.ID,
		ProductName:  m.ProductName,
		ProductStock: m.ProductStock,
		ProductPrice: formatter.Int64ToRupiah(int64(m.ProductPrice)),
		ProductPic:   m.ProductPic,
	}
}

func (u usecaseProductsGet) buildArrayResponse(m []*model.Products) []*responseGet {
	resp := make([]*responseGet, 0)
	for _, v := range m {
		resp = append(resp, u.buildResponse(v))
	}

	return resp
}
