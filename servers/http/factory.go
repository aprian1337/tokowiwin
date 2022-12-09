package http

import (
	"context"
	"tokowiwin/config"
	"tokowiwin/controllers"
	"tokowiwin/repositories/db"
	usecase_users "tokowiwin/usecases/users/usecase"
)

func InitFactoryHTTP(ctx context.Context, cfg *config.AppConfig) {
	dbRepo := db.NewDatabaseRepository(ctx, cfg)

	ucAuth := usecase_users.UCBuyerLogin{}.NewUsecase(ctx, dbRepo)
	authDelivery := controllers.NewController(ctx, ucAuth)

	h := HTTPDelivery{
		UsersAuthentication: authDelivery,
	}
	h.InitHTTPClient(cfg)
}
