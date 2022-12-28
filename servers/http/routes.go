package http

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"time"
	"tokowiwin/config"
	"tokowiwin/controllers"
	"tokowiwin/repositories/db"
)

type DeliveryHTTP struct {
	UsersAuthentication *controllers.Controller
	UsersRegister       *controllers.Controller

	ProductsGet    *controllers.Controller
	ProductsInsert *controllers.Controller
	ProductsUpdate *controllers.Controller
	ProductsDelete *controllers.Controller

	CartsGet    *controllers.Controller
	CartsInsert *controllers.Controller
	CartsUpdate *controllers.Controller
	CartsDelete *controllers.Controller

	TransactionsGet    *controllers.Controller
	TransactionsInsert *controllers.Controller
	TransactionsUpdate *controllers.Controller
}

func (h DeliveryHTTP) InitHTTPClient(cfg *config.AppConfig) {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Next:         nil,
		Done:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path} - ${reqHeaders} - ${body} - ${resBody}\n",
		TimeFormat:   "02-Jan-2006 15:04:05",
		TimeZone:     "Asia/Jakarta",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stdout,
	}))

	app.Post("/login", h.UsersAuthentication.AuthenticationController)
	app.Post("/register", h.UsersRegister.RegisterController)

	app.Get("/products", h.ProductsGet.ProductsGetController)
	app.Post("/products", h.ProductsInsert.ProductsInsertController)
	app.Delete("/products", h.ProductsDelete.ProductsDeleteController)
	app.Put("/products", h.ProductsUpdate.ProductsUpdateController)

	app.Get("/carts", h.CartsGet.CartsGetController)
	app.Post("/carts", h.CartsInsert.CartsInsertController)
	app.Delete("/carts", h.CartsDelete.CartsDeleteController)
	app.Put("/carts", h.CartsUpdate.CartsUpdateController)

	app.Get("/transactions", h.TransactionsGet.TransactionsGetController)
	app.Post("/transactions", h.TransactionsInsert.TransactionsInsertController)

	err := app.Listen(cfg.Server.Address)
	if err != nil {
		panic(fmt.Sprintf("error while start the http client, err=%v", err))
	}
}

func AddController(uc controllers.HandlerI, validate *validator.Validate) *controllers.Controller {
	return controllers.NewController(uc, validate, db.GetDBClient())
}
