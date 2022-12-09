package http

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"tokowiwin/config"
	"tokowiwin/controllers"
)

type DeliveryHTTP struct {
	UsersAuthentication *controllers.Controller
	UsersRegister       *controllers.Controller
}

func (h DeliveryHTTP) InitHTTPClient(cfg *config.AppConfig) {
	app := fiber.New()
	app.Get("/login", h.UsersAuthentication.AuthenticationController)
	app.Get("/register", h.UsersRegister.RegisterController)
	err := app.Listen(cfg.Server.Address)
	if err != nil {
		panic(fmt.Sprintf("error while start the http client, err=%v", err))
	}
}

func AddController(ctx context.Context, uc controllers.HandlerI, validate *validator.Validate) *controllers.Controller {
	return controllers.NewController(ctx, uc, validate)
}
