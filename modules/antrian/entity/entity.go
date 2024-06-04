package entity

import (
	"vincentcoreapi/modules/antrian"
	"vincentcoreapi/modules/antrian/dto"
)

// AntrianUseCase
type AntrianUseCase interface {
	GetStatusAntreanUsecase(payload *dto.StatusAntrianRequest, detailPoli antrian.Kpoli) (res dto.StatusAntreanDTO, err error)
	GetMobileJknByKodebookingUsecase(req dto.GetSisaAntrianRequest) (res map[string]any, err error)
	BatalAntreanUsecase(req dto.BatalAntreanRequest) (isSuccessBatal bool, err error)
	CheckedInUsecase(req dto.CheckInRequest) (isSuccess bool, err error)
	RegisterPasienBaruUsecase(req dto.RegisterPasienBaruRequest) (rres dto.ResPasienBaru, err error)
	GetKodeBookingOperasiByNoPesertaUsecase(req dto.JadwalOperasiPasienRequest) (res map[string]any, err error)
	AmbilAntreanUsecase(req dto.GetAntrianRequest, detailPoli antrian.Kpoli, detaiProfilPasien antrian.Dprofilpasien) (response dto.InsertPasienDTO, err error)
	AmbilAntreanUsecaseV2(req dto.GetAntrianRequestV2, detailPoli antrian.Kpoli, detaiProfilPasien antrian.Dprofilpasien) (response dto.InsertPasienDTO, err error)
	ValidasiDateUsecase(req string) (isTrue bool)
	GetStatusAntreanUsecaseV2(payload *dto.StatusAntrianRequestV2, detailPoli antrian.Kpoli) (res dto.StatusAntreanDTO, err error)
	BatalAntreanUsecaseV2(req dto.BatalAntreanRequestV2) (isSuccessBatal bool, err error)
	GetKodeBookingOperasiByNoPesertaUsecaseV2(req dto.JadwalOperasiPasienRequestV2) (res map[string]any, err error)
}

// AntrianRepository
type AntrianRepository interface {
	CheckPasienDProfilePasienRepository(noka string) (isDuplicate bool)
	CekPoliRepository(value string) (isTrue bool, err error)
	GetJadwalOperasiRepository(tanggalAwal, tanggalAkhir string) (jadopOls []antrian.JadopOl, err error)
	DetailPoliRepository(params map[string]interface{}) (res map[string]interface{}, err error)
	LastCalledRepository(payload *dto.StatusAntrianRequest) (res antrian.LastCalled, err error)
	GetKodeDokterRsRepository(params map[string]interface{}) (res map[string]interface{}, err error)
	JmlAntreanRepository(payload *dto.StatusAntrianRequest, kodeDokter string) (res int, err error)
	GetKodeDokterJadwalRsRepository(day string, params map[string]interface{}) (res bool, err error)
	DetailTaripDokterByMapingAntrolRepository(mapingAntrol int) (res antrian.KtaripDokter, err error)
	CariPoliRepository(kdPoli string) (res antrian.Kpoli, err error)
	GetMobileJknByKodebookingRepository(kodebooking string) (res dto.GetMobileJknByKodebookingDTO, err error)
	GetSisaAntreanRepository(req dto.GetSisaAntrianRequest) (res dto.SisaAntreanResnonse, err error)
	GetAntreanByKodeBookingRepository(kodeBooking string) (antrianOL antrian.AntrianOl, err error)
	BatalAntreanRepository(kodeBooking, keterangan string) (isSuccessBatal bool)
	CheckInRepository(kodeBooking string, waktu int64) (isSuccess bool)
	CheckPasienDuplikatRepository(noka string) (isDuplicate bool)
	InsertPasienBaruRepository(pasienBaru antrian.AntrianOl) (isSuccess bool, norm string)
	GetDokterNameRepository(kodeDokter int) (dokter antrian.KtaripDokter, err error)
	GetKodeBookingOperasiByNoPesertaRepository(noPeserta string) (jadopOls []antrian.JadopOl, err error)
	CheckAntreanRepository(nomorKartu, tglPeriksa, kodePoli string) (jumlah int64, err error)
	CheckDokterLiburRepository(tglPeriksa string, kodeDokter string) (dokterLiburs antrian.LiburOl, err error)
	CheckJadwalPraktekRepository(tglPeriksa string, idDokter string) (jadwal int64, err error)
	CheckKuotaRepository(tglPeriksa string, idDokter string, kuotaToday int) (isAvailable bool)
	CheckMedrekRepository(noRM string) (dprofilpasien antrian.Dprofilpasien, err error)

	InsertAntreanMjknRepository(req dto.GetAntrianRequest, detailKTaripDokter antrian.KtaripDokter, kotaHariIni int, detailPoli antrian.Kpoli, detaiProfilPasien antrian.Dprofilpasien) (response dto.InsertPasienDTO, err error)
	GetSisaAntreanRepositoryV2(req dto.GetSisaAntrianRequestV2) (res dto.SisaAntreanResnonse, err error)
	InsertAntreanMjknRepositoryV2(req dto.GetAntrianRequestV2, detailKTaripDokter antrian.KtaripDokter, kotaHariIni int, detailPoli antrian.Kpoli, detaiProfilPasien antrian.Dprofilpasien, umum string) (response dto.InsertPasienDTO, err error)
	JmlAntreanRepositoryV2(payload *dto.StatusAntrianRequestV2, kodeDokter string) (res int, err error)
	LastCalledRepositoryV2(payload *dto.StatusAntrianRequestV2) (res antrian.LastCalled, err error)
	ListAntrianTodayRepository() (res []antrian.AntrianOl, err error)

	GetNormPasienRepository() (res antrian.IDPasien, err error)
	DetailKtaripDokter2AntrolRepository(idDokter string) (res antrian.KtaripDokter2, err error)
	InsertPasienBaruDprofilePasien(pasienBaru antrian.Dprofilpasien) (res antrian.Dprofilpasien, err error)
}
