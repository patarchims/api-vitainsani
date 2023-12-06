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

	// var datas = dto.RequestHeader{Username: username, Password: password}

	// data, _ := json.Marshal(datas)

	if username == "" {
		uh.Logging.Info("Username kosong")
		response := helper.APIResponseFailure("Username kosong", http.StatusCreated)
		// telegram.RunFailureMessageFiber("GET TOKEN", response, c, data)
		uh.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	if password == "" {
		uh.Logging.Info("Password kosong")
		response := helper.APIResponseFailure("Password kosong", http.StatusCreated)
		// telegram.RunFailureMessageFiber("GET TOKEN", response, c, data)
		uh.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	// => Be Produktif <= //

	user, exist := uh.UserRepository.GetByUserRepository(username)

	if !exist {
		uh.Logging.Info("Username atau Password Tidak Sesuai")
		response := helper.APIResponseFailure("Username atau Password Tidak Sesuai", http.StatusCreated)
		// telegram.RunFailureMessageFiber("GET TOKEN", response, c, data)
		uh.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	if user.Password != password {
		uh.Logging.Info("Username atau Password Tidak Sesuai")
		response := helper.APIResponseFailure("Username atau Password Tidak Sesuai", http.StatusCreated)
		// telegram.RunFailureMessageFiber("GET TOKEN", response, c, data)
		uh.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	m, _ := rest.GenerateTokenPair(user)
	response := helper.APIResponse("Ok", http.StatusOK, m)
	uh.Logging.Info(response)
	// telegram.RunSuccessMessageFiber("GET TOKEN", response, c, data)
	return c.Status(fiber.StatusOK).JSON(response)
}
