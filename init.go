package main

import (
	"fmt"
	"os"
	"vincentcoreapi/config"
	handlerAntrian "vincentcoreapi/modules/antrian/handler"
	mapperAntrian "vincentcoreapi/modules/antrian/mapper"
	repositoryAntrian "vincentcoreapi/modules/antrian/repository"
	antrianUseCase "vincentcoreapi/modules/antrian/usecase"
	handlerFarmasi "vincentcoreapi/modules/farmasi/handler"
	mapperFarmasi "vincentcoreapi/modules/farmasi/mapper"
	repoFarmasi "vincentcoreapi/modules/farmasi/repository"
	farmasiUseCase "vincentcoreapi/modules/farmasi/usecase"
	handlerUser "vincentcoreapi/modules/user/handler"
	repositoryUser "vincentcoreapi/modules/user/repository"
	userUseCase "vincentcoreapi/modules/user/usecase"

	"github.com/joho/godotenv"
)

type Service struct {
	UserHandler    *handlerUser.UserHandler
	AntrianHandler *handlerAntrian.AntrianHandler
	FarmasiHandler *handlerFarmasi.FarmasiHandler
}

func RunApplication() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(1)
	}

	db := config.InitMysqlDB()

	repoUser := repositoryUser.NewUserRepository(db)
	repoAntrian := repositoryAntrian.NewAntrianRepository(db)
	repoFarmasi := repoFarmasi.NewFarmasiRepository(db)

	// MAPPER LAYER
	mapperAntrian := mapperAntrian.NewAntrianMapperImpl()
	mapperFarmasi := mapperFarmasi.NewAntrianMapperImpl()

	// USECASE LAYER
	uu := userUseCase.NewUserUseCase(repoUser)
	fu := farmasiUseCase.FarmasiUseCase(repoFarmasi, mapperFarmasi)
	au := antrianUseCase.NewAntrianUseCase(repoAntrian, mapperAntrian)

	// HANDLER
	userHandler := handlerUser.UserHandler{UserUseCase: uu, UserRepository: repoUser}
	farmasiHandler := handlerFarmasi.FarmasiHandler{FarmasiUseCase: fu, FarmasiRepository: repoFarmasi, IFarmasiMapper: mapperFarmasi}
	antrianHandler := handlerAntrian.AntrianHandler{AntrianUseCase: au, AntrianRepository: repoAntrian, IAntrianMapper: mapperAntrian}

	service := &Service{
		UserHandler:    &userHandler,
		AntrianHandler: &antrianHandler,
		FarmasiHandler: &farmasiHandler,
	}

	// ROUTING APP
	service.RoutingAndListen()
}
