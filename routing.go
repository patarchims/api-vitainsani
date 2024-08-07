package main

import (
	"os"
	"vincentcoreapi/app/rest"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/gin-contrib/gzip"

	_ "vincentcoreapi/docs"
)

// ROUTING APPLICATION
func (s *Service) RoutingAndListen(Logging *logrus.Logger) {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPathsRegexs([]string{".*"})))

	// USER LOGGER MIDDLEWARE
	router.Use(rest.CORSMiddleware())
	router.Use(static.Serve("/apps/files", static.LocalFile("./files", true)))

	router.Use(rest.CORSMiddleware())
	api := router.Group("")

	// API PUBLIC
	// API PROTECTED
	apiProtected := api.Group("")
	apiProtected.Use(rest.JwtVerify())

	apiPublic := api.Group("")
	apiPublic.GET("/gettoken", s.UserHandler.Login)
	apiProtected.POST("/status-antrean", s.AntrianHandler.GetStatusAntrian)
	apiProtected.POST("/sisa-antrean", s.AntrianHandler.GetSisaAntrian)
	apiProtected.POST("/batal-antrean", s.AntrianHandler.BatalAntrean)
	apiProtected.POST("/check-in", s.AntrianHandler.CheckIn)
	apiProtected.POST("/pasien-baru", s.AntrianHandler.RegisterPasienBaru)
	apiProtected.POST("/get-jadwal-operasi", s.AntrianHandler.GetJadwalOperasi)
	apiProtected.POST("/list-jadwal-operasi", s.AntrianHandler.GetKodeBookingOperasi)
	apiProtected.POST("/ambil-antrean", s.AntrianHandler.AmbilAntrean)

	apiProtected.POST("/ambil-antrean-farmasi", s.FarmasiHandler.AmbilAntreanFarmasi)
	apiProtected.POST("/status-antrean-farmasi", s.FarmasiHandler.StatusAntreanFarmasi)

	// MUTIARA
	apiPublic.GET("/karyawan/:id", s.MutiaraHandler.GetDataGaji)
	apiPublic.GET("/karyawan-all", s.MutiaraHandler.Pengajar)
	apiPublic.GET("/pengajar/detail/:id", s.MutiaraHandler.Pengajar)

	// USER
	apiPublic.POST("/user-update", s.UserHandler.Login)

	// FILETRANFER
	apiProtected.POST("/upload-file", s.FileTransferHandler.UploadFile)
	apiPublic.GET("/file-directories", s.FileTransferHandler.UploadFile)

	// RUN SERVER
	router.Run(os.Getenv("DEPLOY_PORT"))

}
