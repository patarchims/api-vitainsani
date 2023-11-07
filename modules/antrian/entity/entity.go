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
	ValidasiDateUsecase(req string) (isTrue bool)
}

// AntrianRepository
type AntrianRepository interface {
	ListAntrianTodayRepository() (res []antrian.AntrianOl, err error)
	CekPoliRepository(value string) (isTrue bool, err error)
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
	GetJadwalOperasiRepository(tanggalAwal, tanggalAkhir string) (jadopOls []antrian.JadopOl, err error)
	GetKodeBookingOperasiByNoPesertaRepository(noPeserta string) (jadopOls []antrian.JadopOl, err error)
	CheckAntreanRepository(nomorKartu, tglPeriksa, kodePoli string) (jumlah int64, err error)
	CheckDokterLiburRepository(tglPeriksa string, kodeDokter string) (dokterLiburs antrian.LiburOl, err error)
	CheckJadwalPraktekRepository(tglPeriksa string, idDokter string) (jadwal int64, err error)
	CheckKuotaRepository(tglPeriksa string, idDokter string, kuotaToday int) (isAvailable bool)
	CheckMedrekRepository(noRM string) (dprofilpasien antrian.Dprofilpasien, err error)
	InsertAntreanMjknRepository(req dto.GetAntrianRequest, detailKTaripDokter antrian.KtaripDokter, kotaHariIni int, detailPoli antrian.Kpoli, detaiProfilPasien antrian.Dprofilpasien) (response dto.InsertPasienDTO, err error)
}
