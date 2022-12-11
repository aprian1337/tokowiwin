package usecase

import (
	"context"
	"tokowiwin/repositories/db"
	"tokowiwin/usecases"
	"tokowiwin/utils/hash"
)

type UCLogin struct{}
type usecaseBuyerLogin struct {
	ctx  context.Context
	repo db.RepositoryI
}

type requestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type responseLogin struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

func (c UCLogin) NewUsecase(ctx context.Context, repo db.RepositoryI) *usecaseBuyerLogin {
	return &usecaseBuyerLogin{
		ctx:  ctx,
		repo: repo,
	}
}

func (u usecaseBuyerLogin) HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error) {
	var (
		req  = new(requestLogin)
		resp = new(responseLogin)
		err  error
	)

	if err = data.HTTPData.BodyParser(req); err != nil {
		return nil, err
	}
	err = data.Validator.Struct(*req)
	if err != nil {
		return nil, err
	}

	user, err := u.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !hash.CheckPasswordHash(req.Password, user.Password) {
		resp.Success = 0
		resp.Message = "Login tidak berhasil"
		return resp, nil
	}

	resp.Success = 1
	resp.Message = "Berhasil"

	return resp, nil
}