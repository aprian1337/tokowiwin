package usecase

import (
	"context"
	"fmt"
	"tokowiwin/repositories/db"
	"tokowiwin/repositories/model"
	"tokowiwin/usecases"
)

type UCBuyerLogin struct{}
type usecaseBuyerLogin struct {
	ctx  context.Context
	repo db.PsqlRepo
}

type loginResponse struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

func (c UCBuyerLogin) NewUsecase(ctx context.Context, repo db.PsqlRepo) *usecaseBuyerLogin {
	return &usecaseBuyerLogin{
		ctx:  ctx,
		repo: repo,
	}
}

func (u usecaseBuyerLogin) HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error) {
	err := u.repo.InsertEmail(context.Background(), nil, &model.Users{
		Name:     "B",
		Email:    "B",
		Password: "B",
	})
	if err != nil {
		fmt.Println("error email :", err)
	}
	a := loginResponse{
		Success: 1,
		Message: "Berhasil",
	}
	return a, nil
}
