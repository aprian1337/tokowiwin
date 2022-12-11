package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"tokowiwin/usecases"
	"tokowiwin/utils/contexts"
)

func (ct Controller) buildParamHandleUsecase(c *fiber.Ctx) (context.Context, usecases.HandleUsecaseData) {
	return contexts.BuildContextApp(), usecases.BuildHandleUsecaseData(c, ct.Validator, ct.TxExecutor)
}
