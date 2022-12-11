package usecase

import (
	"context"
	"github.com/jackc/pgx/v5"
	"tokowiwin/repositories/db"
	"tokowiwin/repositories/model"
	"tokowiwin/usecases"
	"tokowiwin/utils/hash"
)

type UCRegister struct{}
type usecaseUserRegister struct {
	ctx  context.Context
	repo db.RepositoryI
}

func (c UCRegister) NewUsecase(ctx context.Context, repo db.RepositoryI) *usecaseUserRegister {
	return &usecaseUserRegister{
		ctx:  ctx,
		repo: repo,
	}
}

type responseRegister struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

type requestRegister struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
}

func (u usecaseUserRegister) HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error) {
	var (
		req = new(requestRegister)
		err error
	)

	if err = data.HTTPData.BodyParser(req); err != nil {
		return nil, err
	}
	err = data.Validator.Struct(*req)
	if err != nil {
		return nil, err
	}

	hashPassword := hash.HashPassword(req.Password)
	err = db.ExecuteWithTx(ctx, data.TxExecutor, func(tx pgx.Tx) error {
		err = u.repo.InsertUser(ctx, tx, &model.Users{
			Name:     req.Name,
			Email:    req.Email,
			Password: hashPassword,
			IsSeller: false,
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return responseRegister{
		Success: 1,
		Message: "Berhasil",
	}, nil
}
