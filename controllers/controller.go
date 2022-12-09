package controllers

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"tokowiwin/usecases"
)

type Controller struct {
	Ctx       context.Context
	Handler   HandlerI
	Validator *validator.Validate
}

type HandlerI interface {
	HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error)
}

func NewController(ctx context.Context, handler HandlerI, validator *validator.Validate) *Controller {
	return &Controller{
		Ctx:       ctx,
		Handler:   handler,
		Validator: validator,
	}
}

func (ct *Controller) AuthenticationController(c *fiber.Ctx) error {
	data, err := ct.Handler.HandleUsecase(ct.Ctx, usecases.BuildHandleUsecaseData(c, ct.Validator))
	if err != nil {
		return BaseErrorResponse(c, err, http.StatusInternalServerError)
	}
	return c.JSON(data)
}

func (ct *Controller) RegisterController(c *fiber.Ctx) error {
	data, err := ct.Handler.HandleUsecase(ct.Ctx, usecases.BuildHandleUsecaseData(c, ct.Validator))
	if err != nil {
		return BaseErrorResponse(c, err, http.StatusInternalServerError)
	}
	return c.JSON(data)
}
