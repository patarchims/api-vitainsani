package main

import (
	"os"
	"vincentcoreapi/app/rest"
	"vincentcoreapi/config"
	"vincentcoreapi/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

// ROUTING FIBER APPLICATION
func (s *Service) RoutingFiberAndListen(Logging *logrus.Logger) {
	app := fiber.New(config.NewFiberConfig())

	app.Use(cors.New())

	app.Use(recover.New())

	app.Use(favicon.New(favicon.Config{
		File: "./favicon.ico",
		URL:  "/favicon.ico",
	}))

	app.Use(logger.New())

	api := app.Group("/")

	api.Get("gettoken", s.UserHandler.LoginFiberHandler)

	api.Post("status-antrean", rest.JWTVeifyHandler(Logging), s.AntrianHandler.GetStatusAntrianFiberHandler)
	api.Post("sisa-antrean", rest.JWTVeifyHandler(Logging), s.AntrianHandler.GetSisaAntrianFiberHandler)
	api.Post("batal-antrean", rest.JWTVeifyHandler(Logging), s.AntrianHandler.BatalAntreanFiberHandler)
	api.Post("check-in", rest.JWTVeifyHandler(Logging), s.AntrianHandler.CheckInFiberHandler)
	// PASIEN BARU
	api.Post("pasien-baru", rest.JWTVeifyHandler(Logging), s.AntrianHandler.RegisterPasienBaruFiberHandler)
	api.Post("get-jadwal-operasi", rest.JWTVeifyHandler(Logging), s.AntrianHandler.GetJadwalOperasiFiberHandler)
	api.Post("list-jadwal-operasi", rest.JWTVeifyHandler(Logging), s.AntrianHandler.GetKodeBookingOperasiFiberHandler)
	// AMBIL ANTREAN
	api.Post("ambil-antrean", rest.JWTVeifyHandler(Logging), s.AntrianHandler.AmbilAntreanFiberHandler)

	// NEW FITUR, ANTREAN FARMASI
	api.Post("ambil-antrean-farmasi", rest.JWTVeifyHandler(Logging), s.FarmasiHandler.AmbilAntreanFarmasiFiberHandler)
	api.Post("status-antrean-farmasi", rest.JWTVeifyHandler(Logging), s.FarmasiHandler.StatusAntreanFarmasiFiberHandler)
	api.Post("profile-pasien", rest.JWTVeifyHandler(Logging), s.UserHandler.ProfilePasienFiberHandler)

	// UPDATE DATA USER
	api.Post("user-update", s.UserHandler.OnChagedUserIDFiberHandler)

	err := app.Listen(os.Getenv("DEPLOY_PORT"))

	exception.PanicIfNeeded(err)

}
