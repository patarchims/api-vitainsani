package entity

import (
	"context"
	"vincentcoreapi/modules/farmasi"
	"vincentcoreapi/modules/farmasi/dto"
)

type FarmasiRepository interface {
	CekKodeBooking(ctx context.Context, req dto.GetAntreanFarmasiRequest) (res farmasi.AntreanOL, err error)
	InsertAntreanFarmasi(ctx context.Context, cekKodeBooking farmasi.AntreanOL) (res farmasi.AntreanResep, err error)
	CekKodeBookingAntreanResep(ctx context.Context, req dto.GetAntreanFarmasiRequest) (res farmasi.AntreanResep, err error)
	StatusAntreanFarmasi(ctx context.Context) (res farmasi.StatusAntrean, err error)
}

type FarmasiUseCase interface {
	AmbilAntreanFarmasi(ctx context.Context, req dto.GetAntreanFarmasiRequest) (res dto.AmbilAntreanFarmasiResponse, err error)
	StatusAntreanFarmasi(ctx context.Context, req dto.GetAntreanFarmasiRequest) (res dto.StatusAntreanFarmasiResponse, err error)
}
