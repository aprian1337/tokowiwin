package http

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"tokowiwin/config"
	"tokowiwin/controllers"
	"tokowiwin/repositories/db"
)

type DeliveryHTTP struct {
	UsersAuthentication *controllers.Controller
	UsersRegister       *controllers.Controller
	ProductsGet         *controllers.Controller
	ProductsInsert      *controllers.Controller
	ProductsUpdate      *controllers.Controller
	ProductsDelete      *controllers.Controller
}

func (h DeliveryHTTP) InitHTTPClient(cfg *config.AppConfig) {
	app := fiber.New()
	app.Post("/login", h.UsersAuthentication.AuthenticationController)
	app.Post("/register", h.UsersRegister.RegisterController)

	app.Get("/products", h.ProductsGet.ProductsGetController)
	app.Post("/products", h.ProductsInsert.ProductsInsertController)
	app.Delete("/products", h.ProductsDelete.ProductsDeleteController)
	app.Put("/products", h.ProductsUpdate.ProductsUpdateController)
	err := app.Listen(cfg.Server.Address)
	if err != nil {
		panic(fmt.Sprintf("error while start the http client, err=%v", err))
	}
}

func AddController(uc controllers.HandlerI, validate *validator.Validate) *controllers.Controller {
	return controllers.NewController(uc, validate, db.GetDBClient())
}
