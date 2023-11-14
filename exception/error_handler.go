package exception

import (
	"vincentcoreapi/pkg/helper"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, ok := err.(ValidationError)
	if ok {

		response := helper.APIResponseFailure("BAD_REQUEST", fiber.StatusBadRequest)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)

	}

	response := helper.APIResponseFailure("INTERNAL_SERVER_ERROR", 500)
	return ctx.Status(fiber.StatusInternalServerError).JSON(response)

}
