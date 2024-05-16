package config

import (
	"time"
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
		AppName:       "Aplikasi Mobile Vita Insani API v1.0.1",
		WriteTimeout:  15 * time.Second,
		ReadTimeout:   15 * time.Second,
	}
}
