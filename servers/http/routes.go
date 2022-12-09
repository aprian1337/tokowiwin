package http

import (
	"github.com/gofiber/fiber/v2"
	"tokowiwin/config"
	"tokowiwin/controllers"
)

const (
	GET  = "GET"
	POST = "POST"
)

var app *fiber.App

type HTTPDelivery struct {
	UsersAuthentication *controllers.Controller
}

func (h HTTPDelivery) InitHTTPClient(cfg *config.AppConfig) {
	app = fiber.New()
	app.Get("/login", h.UsersAuthentication.AuthenticationController)
	app.Listen(cfg.Server.Address)
}
