package http

import (
	"context"
	validator "github.com/go-playground/validator/v10"
	"tokowiwin/config"
	"tokowiwin/repositories/db"
	carts "tokowiwin/usecases/carts/usecase"
	products "tokowiwin/usecases/products/usecase"
	transactions "tokowiwin/usecases/transactions/usecase"
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
	h.ProductsGet = AddController(products.UCGet{}.NewUsecase(ctx, dbRepo), validate)
	h.ProductsInsert = AddController(carts.UCInsert{}.NewUsecase(ctx, dbRepo), validate)
	h.ProductsUpdate = AddController(products.UCUpdate{}.NewUsecase(ctx, dbRepo), validate)
	h.ProductsDelete = AddController(products.UCDelete{}.NewUsecase(ctx, dbRepo), validate)

	//Products Module
	h.CartsGet = AddController(carts.UCGet{}.NewUsecase(ctx, dbRepo), validate)
	h.CartsInsert = AddController(carts.UCInsert{}.NewUsecase(ctx, dbRepo), validate)
	h.CartsUpdate = AddController(carts.UCUpdate{}.NewUsecase(ctx, dbRepo), validate)
	h.CartsDelete = AddController(carts.UCDelete{}.NewUsecase(ctx, dbRepo), validate)

	//Transactions Module
	h.TransactionsGet = AddController(transactions.UCGet{}.NewUsecase(ctx, dbRepo), validate)
	h.TransactionsInsert = AddController(transactions.UCInsert{}.NewUsecase(ctx, dbRepo), validate)

	h.InitHTTPClient(cfg)
}
