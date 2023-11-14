package usecase

import (
	"errors"
	"fmt"
	"vincentcoreapi/modules/farmasi/dto"
)

func (fr *farmasiUseCase) AmbilAntreanFarmasiUsecaseV2(req dto.GetAntreanFarmasiRequestV2) (res dto.AmbilAntreanFarmasiResponse, err error) {
	// CEK APAKAH KODE BOOKING ADA DI REKAM ANTERAN OL,
	cekKodeBooking, err := fr.farmasiRepository.CekKodeBookingRepositoryV2(req)

	if err != nil || cekKodeBooking.NoBook == "" {
		message := fmt.Sprintf("Kodebooking %s tidak ditemukan", req.Kodebooking)
		return res, errors.New(message)
	}

	// CEK APAKAH SUDAH PERNAH MENGAMBIL ANTREAN
	cekAmbilAntran, _ := fr.farmasiRepository.CekKodeBookingAntreanResepRepositoryV2(req)
	if len(cekAmbilAntran.KodeBooking) > 0 {
		message := fmt.Sprintf("Antrean dengan Kodebooking %s, hanya dapat diambil satu kali", req.Kodebooking)
		return res, errors.New(message)
	}

	// JIKA DI TEMUKAN,
	// SIMPAN ANTREAN PADA TABEL AMBIL ANTREAN FARMASI
	// POSFAR, ANTRAN RESEP
	insertData, errs := fr.farmasiRepository.InsertAntreanFarmasiRepository(cekKodeBooking)

	if errs != nil || insertData.Mrn == "" {
		message := fmt.Sprintf("Kodebooking %s, gagal ambil antrean", req.Kodebooking)
		return res, errors.New(message)
	}

	// MAPPING INSERT ANTREAN FARMASI
	mapper := fr.IFarmasiMapper.ToFarmasiAntreanResep(insertData)

	return mapper, nil
}

func (fr *farmasiUseCase) StatusAntreanFarmasiUsecaseV2(req dto.GetAntreanFarmasiRequestV2) (res dto.StatusAntreanFarmasiResponse, err error) {
	// CEK APAKAH KODE BOOKING ADA DI ANTREAN RESEP
	cekKodeBooking, err := fr.farmasiRepository.CekKodeBookingAntreanResepRepositoryV2(req)

	if err != nil || cekKodeBooking.KodeBooking == "" {
		message := fmt.Sprintf("Kodebooking %s tidak ditemukan", req.Kodebooking)
		return res, errors.New(message)
	}

	// QUERY STATUS ANTREAN
	statusAntrea, _ := fr.farmasiRepository.StatusAntreanFarmasiRepository()

	// MAPPING INSERT ANTREAN FARMASI
	mapper := fr.IFarmasiMapper.ToStatusAntranFarmasiResponse(cekKodeBooking, statusAntrea)

	return mapper, nil
}
