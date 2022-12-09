package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"tokowiwin/usecases"
)

type Controller struct {
	Ctx     context.Context
	Handler ControllerI
}

type ControllerI interface {
	HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error)
}

func NewController(ctx context.Context, handler ControllerI) *Controller {
	return &Controller{
		Ctx:     ctx,
		Handler: handler,
	}
}

func (ct *Controller) AuthenticationController(c *fiber.Ctx) error {
	data, err := ct.Handler.HandleUsecase(ct.Ctx, usecases.BuildHandleUsecaseData(c))
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(data)
}
