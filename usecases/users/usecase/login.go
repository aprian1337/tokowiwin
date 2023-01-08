package usecase

import (
	"context"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
	Success    int    `json:"success"`
	Message    string `json:"message"`
	HeaderText string `json:"header_text"`
	User       *User  `json:"user"`
}

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"fullname"`
	Email string `json:"email"`
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

	resp.User = &User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	resp.Success = 1
	resp.Message = "Berhasil"
	resp.HeaderText = fmt.Sprintf("Hello, %v", cases.Title(language.Indonesian, cases.NoLower).String(user.Name))

	return resp, nil
}
