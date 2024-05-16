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
	"vincentcoreapi/pkg/logs"

	// Mutiara
	handlerMutiara "vincentcoreapi/modules/mutiara/handler"
	mutiaraMapper "vincentcoreapi/modules/mutiara/mapper"
	repositoryMutiara "vincentcoreapi/modules/mutiara/repository"
	mutiaraUseCase "vincentcoreapi/modules/mutiara/usecase"

	// File Transfer
	handlerFileTransfer "vincentcoreapi/modules/transfer/handler"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Service struct {
	UserHandler         *handlerUser.UserHandler
	AntrianHandler      *handlerAntrian.AntrianHandler
	FarmasiHandler      *handlerFarmasi.FarmasiHandler
	MutiaraHandler      *handlerMutiara.MutiaraHandler
	FileTransferHandler *handlerFileTransfer.TranferHandler
	Logging             *logrus.Logger
}

func RunApplication() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("env is not loaded properly")
		os.Exit(1)
	}

	db := config.InitMysqlDB()
	logging := logs.NewLogger()

	repoUser := repositoryUser.NewUserRepository(db)
	repoAntrian := repositoryAntrian.NewAntrianRepository(db)
	repoFarmasi := repoFarmasi.NewFarmasiRepository(db)
	repoMutiara := repositoryMutiara.NewMutiaraRepository(db)

	// MAPPER LAYER
	mapperAntrian := mapperAntrian.NewAntrianMapperImpl()
	mapperFarmasi := mapperFarmasi.NewAntrianMapperImpl()
	mutiaraMapper := mutiaraMapper.NewMutiaranMapperImpl()

	// USECASE LAYER
	uu := userUseCase.NewUserUseCase(repoUser)
	fu := farmasiUseCase.FarmasiUseCase(repoFarmasi, mapperFarmasi)
	au := antrianUseCase.NewAntrianUseCase(repoAntrian, mapperAntrian, logging)
	mu := mutiaraUseCase.MutiaraUseCase(repoMutiara, mutiaraMapper)

	// HANDLER
	userHandler := handlerUser.UserHandler{UserUseCase: uu, UserRepository: repoUser, Logging: logging}
	farmasiHandler := handlerFarmasi.FarmasiHandler{FarmasiUseCase: fu, FarmasiRepository: repoFarmasi, IFarmasiMapper: mapperFarmasi, Logging: logging}
	antrianHandler := handlerAntrian.AntrianHandler{AntrianUseCase: au, AntrianRepository: repoAntrian, IAntrianMapper: mapperAntrian, Logging: logging}
	mutiaraHandler := handlerMutiara.MutiaraHandler{MutiaraUseCase: mu, MutiaraRepository: repoMutiara, Logging: logging}
	fileTransferHandler := handlerFileTransfer.TranferHandler{}

	service := &Service{
		UserHandler:         &userHandler,
		AntrianHandler:      &antrianHandler,
		FarmasiHandler:      &farmasiHandler,
		MutiaraHandler:      &mutiaraHandler,
		FileTransferHandler: &fileTransferHandler,
		Logging:             logging,
	}

	// ROUTING APP
	// service.RoutingAndListen(logging)
	service.RoutingFiberAndListen(logging)
}
