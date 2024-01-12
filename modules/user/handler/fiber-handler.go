package handler

import (
	"net/http"
	"vincentcoreapi/app/rest"
	"vincentcoreapi/helper"

	"github.com/gofiber/fiber/v2"
)

func (uh *UserHandler) LoginFiberHandler(c *fiber.Ctx) error {

	var username = c.Get("x-username")
	var password = c.Get("x-password")

	if username == "" {
		response := helper.APIResponseFailure("Username kosong", http.StatusCreated)
		uh.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	if password == "" {
		response := helper.APIResponseFailure("Password kosong", http.StatusCreated)
		uh.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	user, exist := uh.UserRepository.GetByUserRepository(username)

	if !exist {
		response := helper.APIResponseFailure("Username atau Password Tidak Sesuai", http.StatusCreated)
		uh.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	if user.Password != password {
		uh.Logging.Info("Username atau Password Tidak Sesuai")
		response := helper.APIResponseFailure("Username atau Password Tidak Sesuai", http.StatusCreated)
		uh.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	m, _ := rest.GenerateTokenPair(user)
	response := helper.APIResponse("Ok", http.StatusOK, m)
	uh.Logging.Info(response)
	return c.Status(fiber.StatusOK).JSON(response)
}
