package usecase

import (
	"context"
	"encoding/json"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"tokowiwin/config"
	"tokowiwin/repositories/db"
	"tokowiwin/usecases"
	"tokowiwin/usecases/transactions"
)

type UCPayment struct{}
type usecaseTransactionsPayment struct {
	ctx  context.Context
	repo db.RepositoryI
}

type requestForSnap struct {
	CustomerDetails struct {
		BillingAddress struct {
			Address string `json:"address"`
		} `json:"billing_address"`
		CustomerIdentifier string `json:"customer_identifier"`
		Email              string `json:"email"`
		FirstName          string `json:"first_name"`
		Phone              string `json:"phone"`
		ShippingAddress    struct {
			Address string `json:"address"`
		} `json:"shipping_address"`
	} `json:"customer_details"`
	CreditCard struct {
		Authentication string `json:"authentication"`
		SaveCard       bool   `json:"save_card"`
		Secure         bool   `json:"secure"`
	} `json:"credit_card"`
	ItemDetails        []interface{} `json:"item_details"`
	TransactionDetails struct {
		Currency    string `json:"currency"`
		GrossAmount int64  `json:"gross_amount"`
		OrderID     string `json:"order_id"`
	} `json:"transaction_details"`
	UserID string `json:"user_id"`
}

type requestPayment struct {
	CustomerDetails struct {
		BillingAddress struct {
			Address string `json:"address"`
		} `json:"billing_address"`
		CustomerIdentifier string `json:"customer_identifier"`
		Email              string `json:"email"`
		FirstName          string `json:"first_name"`
		Phone              string `json:"phone"`
		ShippingAddress    struct {
			Address string `json:"address"`
		} `json:"shipping_address"`
	} `json:"customer_details"`
	CreditCard struct {
		Authentication string `json:"authentication"`
		SaveCard       bool   `json:"save_card"`
		Secure         bool   `json:"secure"`
	} `json:"credit_card"`
	ItemDetails        []interface{} `json:"item_details"`
	TransactionDetails struct {
		Currency    string  `json:"currency"`
		GrossAmount float64 `json:"gross_amount"`
		OrderID     string  `json:"order_id"`
	} `json:"transaction_details"`
	UserID string `json:"user_id"`
}

type responsePayment struct {
	Data []*transactions.TransactionList `json:"data"`
}

func (c UCPayment) NewUsecase(ctx context.Context, repo db.RepositoryI) *usecaseTransactionsPayment {
	return &usecaseTransactionsPayment{
		ctx:  ctx,
		repo: repo,
	}
}

func (u usecaseTransactionsPayment) HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error) {
	var (
		reqHttp       = new(requestPayment)
		reqSnap       = new(requestForSnap)
		reqSnapClient = new(snap.Request)
		err           error
	)
	midtrans.ServerKey = config.GetConfig().Gateway.Key
	midtrans.Environment = midtrans.Sandbox

	err = json.Unmarshal(data.HTTPData.Body(), reqHttp)
	if err != nil {
		return nil, err
	}

	reqSnap = &requestForSnap{
		CustomerDetails: struct {
			BillingAddress struct {
				Address string `json:"address"`
			} `json:"billing_address"`
			CustomerIdentifier string `json:"customer_identifier"`
			Email              string `json:"email"`
			FirstName          string `json:"first_name"`
			Phone              string `json:"phone"`
			ShippingAddress    struct {
				Address string `json:"address"`
			} `json:"shipping_address"`
		}{
			BillingAddress: struct {
				Address string `json:"address"`
			}{
				Address: reqHttp.CustomerDetails.BillingAddress.Address,
			},
			CustomerIdentifier: reqHttp.CustomerDetails.CustomerIdentifier,
			Email:              reqHttp.CustomerDetails.Email,
			FirstName:          reqHttp.CustomerDetails.FirstName,
			Phone:              reqHttp.CustomerDetails.Phone,
			ShippingAddress: struct {
				Address string `json:"address"`
			}{
				Address: reqHttp.CustomerDetails.ShippingAddress.Address,
			},
		},
		CreditCard: struct {
			Authentication string `json:"authentication"`
			SaveCard       bool   `json:"save_card"`
			Secure         bool   `json:"secure"`
		}{
			Authentication: reqHttp.CreditCard.Authentication,
			SaveCard:       reqHttp.CreditCard.SaveCard,
			Secure:         reqHttp.CreditCard.Secure,
		},
		ItemDetails: nil,
		TransactionDetails: struct {
			Currency    string `json:"currency"`
			GrossAmount int64  `json:"gross_amount"`
			OrderID     string `json:"order_id"`
		}{
			Currency:    reqHttp.TransactionDetails.Currency,
			GrossAmount: int64(reqHttp.TransactionDetails.GrossAmount),
			OrderID:     reqHttp.TransactionDetails.OrderID,
		},
		UserID: reqHttp.UserID,
	}

	js, err := json.Marshal(reqSnap)
	err = json.Unmarshal(js, reqSnapClient)
	if err != nil {
		return nil, err
	}

	snapResp, _ := snap.CreateTransaction(reqSnapClient)
	return snapResp, nil
}
