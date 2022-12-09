package usecase

import (
	"context"
	"fmt"
	"tokowiwin/repositories/db"
	"tokowiwin/usecases"
	"tokowiwin/usecases/users"
)

type UCBuyerLogin struct{}
type usecase struct {
	ctx  context.Context
	repo *db.DatabaseRepository
}

func (c UCBuyerLogin) NewUsecase(ctx context.Context, repo *db.DatabaseRepository) *usecase {
	return &usecase{
		ctx:  ctx,
		repo: repo,
	}
}

func (u usecase) HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error) {
	user, err := u.repo.GetUserByEmail(ctx, "z")
	if err != nil {
		fmt.Println("error email :", err)
	}
	a := users.AuthenticationResponse{
		Success: 1,
		Message: user.Name,
	}
	return a, nil
}
