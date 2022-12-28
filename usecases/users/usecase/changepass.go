package usecase

import (
	"context"
	"github.com/jackc/pgx/v5"
	"tokowiwin/repositories/db"
	"tokowiwin/usecases"
	"tokowiwin/utils/hash"
)

type UCChangePass struct{}
type usecaseBuyerChangePass struct {
	ctx  context.Context
	repo db.RepositoryI
}

type requestChangePass struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type responseChangePass struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

func (c UCChangePass) NewUsecase(ctx context.Context, repo db.RepositoryI) *usecaseBuyerChangePass {
	return &usecaseBuyerChangePass{
		ctx:  ctx,
		repo: repo,
	}
}

func (u usecaseBuyerChangePass) HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error) {
	var (
		req  = new(requestChangePass)
		resp = new(responseChangePass)
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

	if !hash.CheckPasswordHash(req.OldPassword, user.Password) {
		resp.Success = 0
		resp.Message = "Kata sandi lama tidak benar"
		return resp, nil
	}

	user.Password = hash.HashPassword(req.NewPassword)

	err = db.ExecuteWithTx(ctx, data.TxExecutor, func(tx pgx.Tx) error {
		err = u.repo.UpdateUser(ctx, tx, user)
		if err != nil {
			return err
		}
		return nil
	})

	resp.Success = 1
	resp.Message = "Berhasil"

	return resp, nil
}
