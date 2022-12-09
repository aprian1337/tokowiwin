package controllers

import "github.com/gofiber/fiber/v2"

type errorResponse struct {
	ErrorMessage string `json:"error_message"`
	Status       int    `json:"status"`
}

func BaseErrorResponse(fctx *fiber.Ctx, err error, statusCode int) error {
	return fctx.JSON(errorResponse{
		ErrorMessage: err.Error(),
		Status:       statusCode,
	})
}
