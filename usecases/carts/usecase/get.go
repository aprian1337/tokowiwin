package usecase

import (
	"context"
	"tokowiwin/repositories/db"
	"tokowiwin/repositories/model"
	"tokowiwin/usecases"
)

type UCGet struct{}
type usecaseCartsGet struct {
	ctx  context.Context
	repo db.RepositoryI
}

type requestGet struct {
	UserID int64 `json:"user_id" validate:"required,numeric"`
}

type responseGet struct {
	ID           int64  `json:"id"`
	ProductName  string `json:"product_name"`
	ProductStock int    `json:"product_stock"`
	ProductPrice int    `json:"product_price"`
	ProductPic   string `json:"product_pic"`
	Qty          int64  `json:"qty"`
}

func (c UCGet) NewUsecase(ctx context.Context, repo db.RepositoryI) *usecaseCartsGet {
	return &usecaseCartsGet{
		ctx:  ctx,
		repo: repo,
	}
}

func (u usecaseCartsGet) HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error) {
	var (
		req = new(requestGet)
		err error
	)

	if len(data.HTTPData.Body()) != 0 {
		if err = data.HTTPData.BodyParser(req); err != nil {
			return nil, err
		}
		err = data.Validator.Struct(*req)
		if err != nil {
			return nil, err
		}
	}

	m, err := u.repo.GetCart(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	var productIDs = make([]int64, 0)
	for _, v := range m {
		if v.ProductID.Valid {
			productIDs = append(productIDs, v.ProductID.Int64)
		}
	}

	p, err := u.repo.GetProductsByIDsMapped(ctx, productIDs)

	return u.buildArrayResponse(m, p), nil
}

func (u usecaseCartsGet) buildResponse(p *model.Products, qty int64) *responseGet {
	return &responseGet{
		ID:           p.ID,
		ProductName:  p.ProductName,
		ProductStock: p.ProductStock,
		ProductPrice: p.ProductPrice,
		ProductPic:   p.ProductPic,
		Qty:          qty,
	}
}

func (u usecaseCartsGet) buildArrayResponse(c []*model.Carts, p map[int64]*model.Products) []*responseGet {
	resp := make([]*responseGet, 0)
	for _, v := range c {
		if !v.ProductID.Valid || (v.ProductID.Valid && p[v.ProductID.Int64] == nil) {
			continue
		}
		if v.Qty.Valid {
			resp = append(resp, u.buildResponse(p[v.ProductID.Int64], v.Qty.Int64))
		}
	}

	return resp
}
