package usecase

import (
	"context"
	"tokowiwin/repositories/db"
	"tokowiwin/repositories/model"
	"tokowiwin/usecases"
	"tokowiwin/utils/hash"
)

type UCUserRegister struct{}
type usecaseUserRegister struct {
	ctx  context.Context
	repo db.PsqlRepo
}

func (c UCUserRegister) NewUsecase(ctx context.Context, repo db.PsqlRepo) *usecaseUserRegister {
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
	err = u.repo.InsertEmail(context.Background(), nil, &model.Users{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashPassword,
	})
	if err != nil {
		return nil, err
	}

	return responseRegister{
		Success: 1,
		Message: "Berhasil",
	}, nil
}
