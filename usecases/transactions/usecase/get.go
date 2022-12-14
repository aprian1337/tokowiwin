package usecase

import (
	"context"
	"fmt"
	"tokowiwin/repositories/db"
	"tokowiwin/repositories/model"
	"tokowiwin/repositories/model/helper"
	"tokowiwin/usecases"
	"tokowiwin/usecases/transactions"
	"tokowiwin/utils/formatter"
)

type UCGet struct{}
type usecaseTransactionsGet struct {
	ctx  context.Context
	repo db.RepositoryI
}

type requestGet struct {
	UserID int64 `json:"user_id" validate:"required,numeric"`
}

type responseGet struct {
	Data []*transactions.TransactionList `json:"data"`
}

func (c UCGet) NewUsecase(ctx context.Context, repo db.RepositoryI) *usecaseTransactionsGet {
	return &usecaseTransactionsGet{
		ctx:  ctx,
		repo: repo,
	}
}

func (u usecaseTransactionsGet) HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error) {
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

	txs, err := u.repo.GetTransaction(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	txIDs := helper.GetTransactionIDs(txs)

	snapshots, err := u.repo.GetSnapshotsByIDsMapped(ctx, txIDs)
	if err != nil {
		return nil, err
	}

	return u.buildResponse(txs, snapshots), nil
}

func (u usecaseTransactionsGet) buildResponse(t []*model.Transactions, s map[int64][]*model.Snapshots) *responseGet {
	var (
		txList = make([]*transactions.TransactionList, 0)
	)

	for _, v := range t {
		var (
			tx         = new(transactions.TransactionList)
			txDetail   = new(transactions.TransactionDetail)
			txProducts = make([]*transactions.TransactionProduct, 0)
			totalBill  int64
		)

		for i, x := range s[v.ID] {
			if i == 0 {
				tx.ProductName = x.ProductName
				tx.ProductPic = x.ProductPic
				if len(s[v.ID]) > 1 {
					tx.AnotherProduct = fmt.Sprintf("%v produk lainnya", len(s[v.ID])-1)
				} else {
					tx.AnotherProduct = fmt.Sprintf("Kamu hanya membeli produk ini")
				}
			}

			txProductsTemp := new(transactions.TransactionProduct)
			txProductsTemp.ProductName = x.ProductName
			txProductsTemp.ProductPrice = formatter.Int64ToRupiah(int64(x.ProductPrice))
			txProductsTemp.ProductPic = x.ProductPic
			txProductsTemp.ProductQty = x.Qty
			txProductsTemp.ProductTotalPrice = formatter.Int64ToRupiah(int64(x.ProductPrice) * x.Qty)
			txProducts = append(txProducts, txProductsTemp)

			totalBill += int64(x.ProductPrice) * x.Qty
		}

		txDetail.PaymentType = v.PaymentType
		txDetail.TotalBill = formatter.Int64ToRupiah(totalBill)
		txDetail.ReceiverName = v.ReceiverName
		txDetail.ReceiverPhone = v.ReceiverPhone
		txDetail.ReceiverAddress = v.ReceiverAddress
		txDetail.TransactionProducts = txProducts

		tx.TransactionID = v.ID
		tx.Date = formatter.ToTimezoneJakarta(v.CreatedDate).Format("2 January 2006, 15:04:05")
		tx.TransactionDetails = txDetail

		txList = append(txList, tx)
	}

	return &responseGet{
		Data: txList,
	}
}
