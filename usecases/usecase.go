package usecases

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"tokowiwin/repositories/db"
)

type HandleUsecaseData struct {
	HTTPData   *fiber.Ctx
	Validator  *validator.Validate
	TxExecutor db.TxExecutor
}

func BuildHandleUsecaseData(c *fiber.Ctx, validate *validator.Validate, txExec db.TxExecutor) HandleUsecaseData {
	return HandleUsecaseData{
		Validator:  validate,
		HTTPData:   c,
		TxExecutor: txExec,
	}
}
