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

	// Mutiara
	handlerMutiara "vincentcoreapi/modules/mutiara/handler"
	mutiaraMapper "vincentcoreapi/modules/mutiara/mapper"
	repositoryMutiara "vincentcoreapi/modules/mutiara/repository"
	mutiaraUseCase "vincentcoreapi/modules/mutiara/usecase"

	// File Transfer
	handlerFileTransfer "vincentcoreapi/modules/transfer/handler"

	"github.com/joho/godotenv"
)

type Service struct {
	UserHandler         *handlerUser.UserHandler
	AntrianHandler      *handlerAntrian.AntrianHandler
	FarmasiHandler      *handlerFarmasi.FarmasiHandler
	MutiaraHandler      *handlerMutiara.MutiaraHandler
	FileTransferHandler *handlerFileTransfer.TranferHandler
}

func RunApplication() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("env is not loaded properly")
		os.Exit(1)
	}

	db := config.InitMysqlDB()
	// logging := logs.NewLogger()

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
	au := antrianUseCase.NewAntrianUseCase(repoAntrian, mapperAntrian)
	mu := mutiaraUseCase.MutiaraUseCase(repoMutiara, mutiaraMapper)

	// HANDLER
	userHandler := handlerUser.UserHandler{UserUseCase: uu, UserRepository: repoUser}
	farmasiHandler := handlerFarmasi.FarmasiHandler{FarmasiUseCase: fu, FarmasiRepository: repoFarmasi, IFarmasiMapper: mapperFarmasi}
	antrianHandler := handlerAntrian.AntrianHandler{AntrianUseCase: au, AntrianRepository: repoAntrian, IAntrianMapper: mapperAntrian}
	mutiaraHandler := handlerMutiara.MutiaraHandler{MutiaraUseCase: mu}
	fileTransferHandler := handlerFileTransfer.TranferHandler{}

	service := &Service{
		UserHandler:         &userHandler,
		AntrianHandler:      &antrianHandler,
		FarmasiHandler:      &farmasiHandler,
		MutiaraHandler:      &mutiaraHandler,
		FileTransferHandler: &fileTransferHandler,
	}

	// ROUTING APP
	// service.RoutingFiberAndListen(logging)
	service.RoutingAndListen()
}
