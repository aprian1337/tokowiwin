package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"strconv"
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
	ID              int64  `json:"id"`
	UserID          int64  `json:"user_id"`
	ReceiverName    string `json:"receiver_name"`
	ReceiverPhone   string `json:"receiver_phone"`
	ReceiverAddress string `json:"receiver_address"`
	PaymentType     string `json:"payment_type"`
}

type responseInsert struct {
	ID          int64  `json:"id"`
	PaymentType string `json:"payment_type"`
	Success     int    `json:"success"`
	Message     string `json:"message"`
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
		idTx int64
	)

	if err = data.HTTPData.BodyParser(req); err != nil {
		return nil, err
	}
	err = data.Validator.Struct(*req)
	if err != nil {
		return nil, err
	}

	if req.ID == 0 {
		IDStr := fmt.Sprintf("%v%v", req.UserID, time.Now().UnixMilli())
		IDNum, _ := strconv.ParseInt(IDStr, 10, 64)
		req.ID = IDNum
	}

	if req.PaymentType != model.PaymentTypeCOD && req.PaymentType != model.PaymentTypeManualTransfer {
		return nil, errors.New(fmt.Sprintf("payment type only accept for %v, %v", model.PaymentTypeCOD, model.PaymentTypeManualTransfer))
	}

	carts, err := u.repo.GetCart(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	if len(carts) == 0 {
		return nil, errors.New("cart is empty, no product need to checkout")
	}

	productIDs := helper.GetProductIDs(carts)

	products, err := u.repo.GetProductsByIDsMapped(ctx, productIDs)
	if err != nil {
		return nil, err
	}

	err = db.ExecuteWithTx(ctx, data.TxExecutor, func(tx pgx.Tx) error {
		var (
			err    error
			status = model.StatusBelumDibayar
		)

		if req.PaymentType == model.PaymentTypeCOD {
			status = model.PaymentTypeCOD
		}

		idTx, err = u.repo.InsertTransaction(ctx, tx, &model.Transactions{
			ID:              req.ID,
			UserID:          req.UserID,
			ReceiverName:    req.ReceiverName,
			ReceiverPhone:   req.ReceiverPhone,
			ReceiverAddress: req.ReceiverAddress,
			PaymentType:     req.PaymentType,
			Status:          status,
			CreatedDate:     time.Now().UTC(),
		})
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
				TransactionID: idTx,
				ProductPic:    productPic,
				ProductName:   productName,
				ProductPrice:  productPrice,
				Qty:           productQty,
			})
			if err != nil {
				return err
			}

			err = u.repo.UpdateProduct(ctx, tx, &model.Products{
				ID:           v.ProductID.Int64,
				ProductName:  productName,
				ProductStock: products[v.ProductID.Int64].ProductStock - int(productQty),
				ProductPrice: productPrice,
				ProductPic:   productPic,
			})
			if err != nil {
				return err
			}

		}

		if err != nil {
			return err
		}

		err = u.repo.DeleteAllCartByUserID(ctx, tx, req.UserID)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	resp.ID = idTx
	resp.PaymentType = req.PaymentType
	resp.Success = 1
	resp.Message = "Berhasil"

	return resp, nil
}
