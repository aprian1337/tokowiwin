package usecase

import (
	"context"
	"github.com/jackc/pgx/v5"
	"time"
	"tokowiwin/repositories/db"
	"tokowiwin/repositories/model"
	"tokowiwin/repositories/model/helper"
	"tokowiwin/usecases"
)

type UCInsert struct{}
type usecaseTransactionsInsert struct {
	ctx  context.Context
	repo db.RepositoryI
}

type requestInsert struct {
	ID              int64  `db:"id"`
	UserID          int64  `db:"user_id"`
	ReceiverName    string `db:"receiver_name"`
	ReceiverPhone   string `db:"receiver_phone"`
	ReceiverAddress string `db:"receiver_address"`
	PaymentType     string `db:"payment_type"`
}

type responseInsert struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

func (c UCInsert) NewUsecase(ctx context.Context, repo db.RepositoryI) *usecaseTransactionsInsert {
	return &usecaseTransactionsInsert{
		ctx:  ctx,
		repo: repo,
	}
}

func (u usecaseTransactionsInsert) HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error) {
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
		var (
			err error
			id  int64
		)

		id, err = u.repo.InsertTransaction(ctx, tx, &model.Transactions{
			UserID:          req.UserID,
			ReceiverName:    req.ReceiverName,
			ReceiverPhone:   req.ReceiverPhone,
			ReceiverAddress: req.ReceiverAddress,
			PaymentType:     req.PaymentType,
			Status:          model.StatusBelumDibayar,
			Date:            time.Now().UTC(),
		})
		if err != nil {
			return err
		}

		carts, err := u.repo.GetCart(ctx, req.UserID)
		if err != nil {
			return err
		}

		productIDs := helper.GetProductIDs(carts)

		products, err := u.repo.GetProductsByIDsMapped(ctx, productIDs)
		if err != nil {
			return err
		}

		for _, v := range carts {
			var (
				productPic   string
				productName  string
				productPrice int
				productQty   int64
			)

			if v.ProductID.Valid && products[v.ProductID.Int64] != nil {
				productPic = products[v.ProductID.Int64].ProductPic
				productName = products[v.ProductID.Int64].ProductName
				productPrice = products[v.ProductID.Int64].ProductPrice
			}

			if v.Qty.Valid {
				productQty = v.Qty.Int64
			}

			err = u.repo.InsertSnapshot(ctx, tx, &model.Snapshots{
				TransactionID: id,
				ProductPic:    productPic,
				ProductName:   productName,
				ProductPrice:  productPrice,
				Qty:           productQty,
			})

			if err != nil {
				return err
			}
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
