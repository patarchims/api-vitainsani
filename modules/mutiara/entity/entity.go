package entity

import (
	"vincentcoreapi/modules/mutiara"
	"vincentcoreapi/modules/mutiara/dto"
)

type MutiaraRepository interface {
	GetGajiRepository(userID string) (res []mutiara.DGaji, err error)
	GetKaryawanRepository(userID string) (res mutiara.DKaryawan, err error)
	GetPengajarRepository() (res []mutiara.DKaryawan, err error)
}

type MutiaraUseCase interface {
	GetDataKaryawanUsecase(userID string) (data interface{}, err error)
}

type MutiaraMapper interface {
	ToDataGajiMapper(gaji []mutiara.DGaji) (res []dto.ResDataGaji)
}
