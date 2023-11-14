package config

import (
	"vincentcoreapi/exception"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{

		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		ErrorHandler:  exception.ErrorHandler,
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Aplikasi Mobile API v1.0.1",
	}
}
