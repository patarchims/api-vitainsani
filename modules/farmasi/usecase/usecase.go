package usecase

import (
	"context"
	"errors"
	"fmt"
	"vincentcoreapi/modules/farmasi/dto"
	"vincentcoreapi/modules/farmasi/entity"
	"vincentcoreapi/modules/farmasi/mapper"
)

type farmasiUseCase struct {
	farmasiRepository entity.FarmasiRepository
	IFarmasiMapper    mapper.IFarmasiMapper
}

func FarmasiUseCase(fr entity.FarmasiRepository, IFarmasiMapper mapper.IFarmasiMapper) entity.FarmasiUseCase {
	return &farmasiUseCase{
		farmasiRepository: fr,
		IFarmasiMapper:    IFarmasiMapper,
	}
}

func (fr *farmasiUseCase) AmbilAntreanFarmasi(ctx context.Context, req dto.GetAntreanFarmasiRequest) (res dto.AmbilAntreanFarmasiResponse, err error) {
	// CEK APAKAH KODE BOOKING ADA DI REKAM ANTERAN OL,
	cekKodeBooking, err := fr.farmasiRepository.CekKodeBooking(ctx, req)

	if err != nil || cekKodeBooking.NoBook == "" {
		message := fmt.Sprintf("Kodebooking %s tidak ditemukan", req.Kodebooking)
		return res, errors.New(message)
	}

	// CEK APAKAH SUDAH PERNAH MENGAMBIL ANTREAN
	cekAmbilAntran, _ := fr.farmasiRepository.CekKodeBookingAntreanResep(ctx, req)
	if len(cekAmbilAntran.KodeBooking) > 0 {
		message := fmt.Sprintf("Antrean dengan Kodebooking %s, hanya dapat diambil satu kali", req.Kodebooking)
		return res, errors.New(message)
	}

	// JIKA DI TEMUKAN,
	// SIMPAN ANTREAN PADA TABEL AMBIL ANTREAN FARMASI
	// POSFAR, ANTRAN RESEP
	insertData, errs := fr.farmasiRepository.InsertAntreanFarmasi(ctx, cekKodeBooking)

	if errs != nil || insertData.Mrn == "" {
		message := fmt.Sprintf("Kodebooking %s, gagal ambil antrean", req.Kodebooking)
		return res, errors.New(message)
	}

	// MAPPING INSERT ANTREAN FARMASI
	mapper := fr.IFarmasiMapper.ToFarmasiAntreanResep(insertData)

	return mapper, nil
}

func (fr *farmasiUseCase) StatusAntreanFarmasi(ctx context.Context, req dto.GetAntreanFarmasiRequest) (res dto.StatusAntreanFarmasiResponse, err error) {
	// CEK APAKAH KODE BOOKING ADA DI ANTREAN RESEP
	cekKodeBooking, err := fr.farmasiRepository.CekKodeBookingAntreanResep(ctx, req)

	if err != nil || cekKodeBooking.KodeBooking == "" {
		message := fmt.Sprintf("Kodebooking %s tidak ditemukan", req.Kodebooking)
		return res, errors.New(message)
	}

	// QUERY STATUS ANTREAN
	statusAntrea, _ := fr.farmasiRepository.StatusAntreanFarmasi(ctx)

	// MAPPING INSERT ANTREAN FARMASI
	mapper := fr.IFarmasiMapper.ToStatusAntranFarmasiResponse(cekKodeBooking, statusAntrea)

	return mapper, nil
}
