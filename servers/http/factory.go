package http

import (
	"context"
	validator "github.com/go-playground/validator/v10"
	"tokowiwin/config"
	"tokowiwin/repositories/db"
	users "tokowiwin/usecases/users/usecase"
)

func InitFactoryHTTP(ctx context.Context, cfg *config.AppConfig) {
	dbRepo := db.NewDatabaseRepository(ctx, cfg)
	validate := validator.New()
	h := DeliveryHTTP{}

	h.UsersAuthentication = AddController(ctx, users.UCBuyerLogin{}.NewUsecase(ctx, dbRepo), validate)
	h.UsersRegister = AddController(ctx, users.UCUserRegister{}.NewUsecase(ctx, dbRepo), validate)

	h.InitHTTPClient(cfg)
}
