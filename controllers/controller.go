package controllers

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"tokowiwin/repositories/db"
	"tokowiwin/usecases"
)

type Controller struct {
	Ctx        context.Context
	Handler    HandlerI
	Validator  *validator.Validate
	TxExecutor db.TxExecutor
}

type HandlerI interface {
	HandleUsecase(ctx context.Context, data usecases.HandleUsecaseData) (interface{}, error)
}

func NewController(handler HandlerI, validator *validator.Validate, tx db.TxExecutor) *Controller {
	return &Controller{
		Handler:    handler,
		Validator:  validator,
		TxExecutor: tx,
	}
}

func (ct *Controller) AuthenticationController(c *fiber.Ctx) error {
	data, err := ct.Handler.HandleUsecase(ct.buildParamHandleUsecase(c))
	if err != nil {
		return BaseErrorResponse(c, err, http.StatusInternalServerError)
	}
	return c.JSON(data)
}

func (ct *Controller) RegisterController(c *fiber.Ctx) error {
	data, err := ct.Handler.HandleUsecase(ct.buildParamHandleUsecase(c))
	if err != nil {
		return BaseErrorResponse(c, err, http.StatusInternalServerError)
	}
	return c.JSON(data)
}

func (ct *Controller) ProductsGetController(c *fiber.Ctx) error {
	data, err := ct.Handler.HandleUsecase(ct.buildParamHandleUsecase(c))
	if err != nil {
		return BaseErrorResponse(c, err, http.StatusInternalServerError)
	}
	return c.JSON(data)
}

func (ct *Controller) ProductsInsertController(c *fiber.Ctx) error {
	data, err := ct.Handler.HandleUsecase(ct.buildParamHandleUsecase(c))
	if err != nil {
		return BaseErrorResponse(c, err, http.StatusInternalServerError)
	}
	return c.JSON(data)
}

func (ct *Controller) ProductsDeleteController(c *fiber.Ctx) error {
	data, err := ct.Handler.HandleUsecase(ct.buildParamHandleUsecase(c))
	if err != nil {
		return BaseErrorResponse(c, err, http.StatusInternalServerError)
	}
	return c.JSON(data)
}

func (ct *Controller) ProductsUpdateController(c *fiber.Ctx) error {
	data, err := ct.Handler.HandleUsecase(ct.buildParamHandleUsecase(c))
	if err != nil {
		return BaseErrorResponse(c, err, http.StatusInternalServerError)
	}
	return c.JSON(data)
}

func (ct *Controller) CartsGetController(c *fiber.Ctx) error {
	data, err := ct.Handler.HandleUsecase(ct.buildParamHandleUsecase(c))
	if err != nil {
		return BaseErrorResponse(c, err, http.StatusInternalServerError)
	}
	return c.JSON(data)
}

func (ct *Controller) CartsInsertController(c *fiber.Ctx) error {
	data, err := ct.Handler.HandleUsecase(ct.buildParamHandleUsecase(c))
	if err != nil {
		return BaseErrorResponse(c, err, http.StatusInternalServerError)
	}
	return c.JSON(data)
}

func (ct *Controller) CartsDeleteController(c *fiber.Ctx) error {
	data, err := ct.Handler.HandleUsecase(ct.buildParamHandleUsecase(c))
	if err != nil {
		return BaseErrorResponse(c, err, http.StatusInternalServerError)
	}
	return c.JSON(data)
}

func (ct *Controller) CartsUpdateController(c *fiber.Ctx) error {
	data, err := ct.Handler.HandleUsecase(ct.buildParamHandleUsecase(c))
	if err != nil {
		return BaseErrorResponse(c, err, http.StatusInternalServerError)
	}
	return c.JSON(data)
}

func (ct *Controller) TransactionsUpdateController(c *fiber.Ctx) error {
	data, err := ct.Handler.HandleUsecase(ct.buildParamHandleUsecase(c))
	if err != nil {
		return BaseErrorResponse(c, err, http.StatusInternalServerError)
	}
	return c.JSON(data)
}

func (ct *Controller) TransactionsInsertController(c *fiber.Ctx) error {
	data, err := ct.Handler.HandleUsecase(ct.buildParamHandleUsecase(c))
	if err != nil {
		return BaseErrorResponse(c, err, http.StatusInternalServerError)
	}
	return c.JSON(data)
}

func (ct *Controller) TransactionsGetController(c *fiber.Ctx) error {
	data, err := ct.Handler.HandleUsecase(ct.buildParamHandleUsecase(c))
	if err != nil {
		return BaseErrorResponse(c, err, http.StatusInternalServerError)
	}
	return c.JSON(data)
}
