package main

import (
	"io/ioutil"
	"os"
	"vincentcoreapi/app/rest"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/gzip"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "vincentcoreapi/docs"
)

// ROUTING APPLICATION
func (s *Service) RoutingAndListen() {

	gin.DefaultWriter = ioutil.Discard
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

	// NEW FITUR, ANTREAN FARMASI
	apiProtected.POST("/ambil-antrean-farmasi", s.FarmasiHandler.AmbilAntreanFarmasi)
	apiProtected.POST("/status-antrean-farmasi", s.FarmasiHandler.StatusAntreanFarmasi)

	// MUTIARA
	apiPublic.GET("/karyawan/:id", s.MutiaraHandler.GetDataGaji)
	apiPublic.GET("/karyawan-all", s.MutiaraHandler.Pengajar)
	apiPublic.GET("/pengajar/detail/:id", s.MutiaraHandler.Pengajar)

	// FILETRANFER
	apiProtected.POST("/upload-file", s.FileTransferHandler.UploadFile)
	apiPublic.GET("/file-directories", s.FileTransferHandler.UploadFile)
	api.GET("api/prod/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:6060/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	// RUN SERVER
	router.Run(os.Getenv("DEPLOY_PORT"))

}
