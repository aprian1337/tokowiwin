package usecases

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type HandleUsecaseData struct {
	HTTPData  *fiber.Ctx
	Validator *validator.Validate
}

func BuildHandleUsecaseData(c *fiber.Ctx, validate *validator.Validate) HandleUsecaseData {
	return HandleUsecaseData{
		Validator: validate,
		HTTPData:  c,
	}
}
