package usecase

import (
	"context"
	"tokowiwin/repositories/db"
	"tokowiwin/repositories/model"
	"tokowiwin/usecases"
	"tokowiwin/usecases/carts"
	"tokowiwin/utils/formatter"
)

type UCGet struct{}
type usecaseCartsGet struct {
	ctx  context.Context
	repo db.RepositoryI
}

type requestGet struct {
	ID int64 `json:"id"`
}

type responseGet struct {
	Data         []*carts.Carts `json:"data"`
	TotalTagihan string         `json:"total_tagihan"`
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

	if data.HTTPData.Query("id") != "" {
		if err = data.HTTPData.QueryParser(req); err != nil {
			return nil, err
		}
		err = data.Validator.Struct(*req)
		if err != nil {
			return nil, err
		}
	}

	m, err := u.repo.GetCart(ctx, req.ID)
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

func (u usecaseCartsGet) buildResponse(p *model.Products, qty int64) *carts.Carts {
	return &carts.Carts{
		ID:                p.ID,
		ProductName:       p.ProductName,
		ProductStock:      p.ProductStock,
		ProductPrice:      formatter.Int64ToRupiah(int64(p.ProductPrice)),
		ProductTotalPrice: formatter.Int64ToRupiah(int64(p.ProductPrice) * qty),
		ProductPic:        p.ProductPic,
		Qty:               qty,
	}
}

func (u usecaseCartsGet) buildArrayResponse(c []*model.Carts, p map[int64]*model.Products) *responseGet {
	cartsTemp := make([]*carts.Carts, 0)
	resp := &responseGet{}
	var totalTagihan int64 = 0
	for _, v := range c {
		if !v.ProductID.Valid || (v.ProductID.Valid && p[v.ProductID.Int64] == nil) {
			continue
		}
		if v.Qty.Valid {
			cartsTemp = append(cartsTemp, u.buildResponse(p[v.ProductID.Int64], v.Qty.Int64))
			totalTagihan += v.Qty.Int64 * int64(p[v.ProductID.Int64].ProductPrice)
		}
	}
	resp.Data = cartsTemp
	resp.TotalTagihan = formatter.Int64ToRupiah(totalTagihan)

	return resp
}
