package controllers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/st0pcha/go-fiber-example/pkg/response"
)

func GetHelloWorld(ctx fiber.Ctx) error {
	return response.SuccessResponse(ctx, fiber.StatusOK, "Hello world!", nil)
}
