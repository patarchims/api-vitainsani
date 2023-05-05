package main

import (
	"os"
	"vincentcoreapi/app/rest"

	"github.com/gin-gonic/gin"
)

// ROUTING APPLICATION
func (s *Service) RoutingAndListen() {

	router := gin.Default()

	router.Use(rest.CORSMiddleware())
	// API VERSI 1
	api := router.Group("")

	// API PUBLIC
	apiPublic := api.Group("")
	apiPublic.GET("/gettoken", s.UserHandler.Login)

	// API PROTECTED
	apiProtected := api.Group("")
	apiProtected.Use(rest.JwtVerify())

	apiProtected.POST("/status-antrean", s.AntrianHandler.GetStatusAntrian)
	apiProtected.POST("/sisa-antrean", s.AntrianHandler.GetSisaAntrian)
	// apiProtected.POST("/antrean/batal", s.AntrianHandler.BatalAntrean)
	apiProtected.POST("/batal-antrean", s.AntrianHandler.BatalAntrean)
	apiProtected.POST("/check-in", s.AntrianHandler.CheckIn)
	apiProtected.POST("/pasien-baru", s.AntrianHandler.RegisterPasienBaru)
	apiProtected.POST("/get-jadwal-operasi", s.AntrianHandler.GetJadwalOperasi)
	apiProtected.POST("/list-jadwal-operasi", s.AntrianHandler.GetKodeBookingOperasi)
	apiProtected.POST("/ambil-antrean", s.AntrianHandler.AmbilAntrean)

	// NEW FITUR, ANTREAN FARMASI
	apiProtected.POST("/ambil-antrean-farmasi", s.FarmasiHandler.AmbilAntreanFarmasi)
	apiProtected.POST("/status-antrean-farmasi", s.FarmasiHandler.StatusAntreanFarmasi)

	// RUN SERVER
	router.Run(os.Getenv("DEPLOY_PORT"))

}
