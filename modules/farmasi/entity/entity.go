package entity

import (
	"vincentcoreapi/modules/farmasi"
	"vincentcoreapi/modules/farmasi/dto"
)

type FarmasiRepository interface {
	CekKodeBookingRepository(req dto.GetAntreanFarmasiRequest) (res farmasi.AntreanOL, err error)
	CekKodeBookingRepositoryV2(req dto.GetAntreanFarmasiRequestV2) (res farmasi.AntreanOL, err error)
	InsertAntreanFarmasiRepository(cekKodeBooking farmasi.AntreanOL) (res farmasi.AntreanResep, err error)
	CekKodeBookingAntreanResepRepository(req dto.GetAntreanFarmasiRequest) (res farmasi.AntreanResep, err error)
	StatusAntreanFarmasiRepository() (res farmasi.StatusAntrean, err error)
	CekKodeBookingAntreanResepRepositoryV2(req dto.GetAntreanFarmasiRequestV2) (res farmasi.AntreanResep, err error)
}

type FarmasiUseCase interface {
	AmbilAntreanFarmasiUsecase(req dto.GetAntreanFarmasiRequest) (res dto.AmbilAntreanFarmasiResponse, err error)
	AmbilAntreanFarmasiUsecaseV2(req dto.GetAntreanFarmasiRequestV2) (res dto.AmbilAntreanFarmasiResponse, err error)
	StatusAntreanFarmasiUsecase(req dto.GetAntreanFarmasiRequest) (res dto.StatusAntreanFarmasiResponse, err error)
	StatusAntreanFarmasiUsecaseV2(req dto.GetAntreanFarmasiRequestV2) (res dto.StatusAntreanFarmasiResponse, err error)
}
