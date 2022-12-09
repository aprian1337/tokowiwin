package usecases

import "github.com/gofiber/fiber/v2"

type HandleUsecaseData struct {
	HTTPData *fiber.Request
}

func BuildHandleUsecaseData(c *fiber.Ctx) HandleUsecaseData {
	return HandleUsecaseData{
		HTTPData: c.Request(),
	}
}
