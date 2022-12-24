package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type errorResponse struct {
	ErrorMessage     string `json:"error_message"`
	DeveloperMessage string `json:"developer_message"`
	Status           int    `json:"status"`
}

type successResponse struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

func BaseErrorResponse(fctx *fiber.Ctx, err error, statusCode int) error {
	return fctx.JSON(errorResponse{
		ErrorMessage:     "Aplikasi sedang ada masalah, silakan coba lagi secara berkala..",
		DeveloperMessage: err.Error(),
		Status:           statusCode,
	})
}

func BaseSuccessResponse(fctx *fiber.Ctx, data interface{}) error {
	return fctx.JSON(successResponse{
		Data:   data,
		Status: http.StatusOK,
	})
}
