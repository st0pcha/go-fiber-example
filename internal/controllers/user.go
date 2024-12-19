package controllers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/st0pcha/go-fiber-example/internal/dal"
	"github.com/st0pcha/go-fiber-example/pkg/response"
)

func GetUserByID(ctx fiber.Ctx) error {
	id := ctx.Params("id")

	if id == "1" {
		return response.ErrorResponse(ctx, fiber.StatusNotFound, "user not found")
	}

	user := dal.NewUser(id, "test")
	return response.SuccessResponse(ctx, fiber.StatusOK, "success", user)
}
