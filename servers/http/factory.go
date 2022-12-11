package http

import (
	"context"
	validator "github.com/go-playground/validator/v10"
	"tokowiwin/config"
	"tokowiwin/repositories/db"
	products "tokowiwin/usecases/products/usecase"
	users "tokowiwin/usecases/users/usecase"
)

func InitFactoryHTTP(ctx context.Context, cfg *config.AppConfig) {
	dbRepo := db.NewDatabaseRepository(ctx, cfg)
	validate := validator.New()
	h := DeliveryHTTP{}

	//Users Module
	h.UsersAuthentication = AddController(users.UCLogin{}.NewUsecase(ctx, dbRepo), validate)
	h.UsersRegister = AddController(users.UCRegister{}.NewUsecase(ctx, dbRepo), validate)

	//Products Module
	h.ProductsInsert = AddController(products.UCInsert{}.NewUsecase(ctx, dbRepo), validate)
	h.ProductsGet = AddController(products.UCGet{}.NewUsecase(ctx, dbRepo), validate)

	h.InitHTTPClient(cfg)
}
