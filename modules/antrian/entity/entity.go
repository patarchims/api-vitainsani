package entity

import (
	"context"
	"vincentcoreapi/modules/antrian"
	"vincentcoreapi/modules/antrian/dto"
)

// AntrianUseCase
type AntrianUseCase interface {
	GetStatusAntrean(ctx context.Context, payload *dto.StatusAntrianRequest, detailPoli antrian.Kpoli) (res dto.StatusAntreanDTO, err error)
	GetMobileJknByKodebooking(ctx context.Context, req dto.GetSisaAntrianRequest) (res map[string]any, err error)
	BatalAntrean(ctx context.Context, req dto.BatalAntreanRequest) (isSuccessBatal bool, err error)
	CheckedIn(ctx context.Context, req dto.CheckInRequest) (isSuccess bool, err error)
	RegisterPasienBaru(ctx context.Context, req dto.RegisterPasienBaruRequest) (rres dto.ResPasienBaru, err error)
	GetKodeBookingOperasiByNoPeserta(ctx context.Context, req dto.JadwalOperasiPasienRequest) (res map[string]any, err error)
	AmbilAntrean(ctx context.Context, req dto.GetAntrianRequest, detailPoli antrian.Kpoli, detaiProfilPasien antrian.Dprofilpasien) (response dto.InsertPasienDTO, err error)
	ValidasiDate(ctx context.Context, req string) (isTrue bool)
}

// AntrianRepository
type AntrianRepository interface {
	CekPoli(ctx context.Context, value string) (isTrue bool, err error)
	DetailPoli(ctx context.Context, params map[string]interface{}) (res map[string]interface{}, err error)
	LastCalled(ctx context.Context, payload *dto.StatusAntrianRequest) (res antrian.LastCalled, err error)
	GetKodeDokterRs(ctx context.Context, params map[string]interface{}) (res map[string]interface{}, err error)
	JmlAntrean(ctx context.Context, payload *dto.StatusAntrianRequest, kodeDokter string) (res int, err error)
	GetKodeDokterJadwalRs(ctx context.Context, day string, params map[string]interface{}) (res bool, err error)
	DetailTaripDokterByMapingAntrol(ctx context.Context, mapingAntrol int) (res antrian.KtaripDokter, err error)
	CariPoli(ctx context.Context, kdPoli string) (res antrian.Kpoli, err error)

	GetMobileJknByKodebooking(ctx context.Context, kodebooking string) (
		res dto.GetMobileJknByKodebookingDTO, err error)
	GetSisaAntrean(ctx context.Context, req dto.GetSisaAntrianRequest) (
		res dto.SisaAntreanResnonse, err error)
	GetAntreanByKodeBooking(ctx context.Context, kodeBooking string) (
		antrianOL antrian.AntrianOl, err error)
	BatalAntrean(ctx context.Context, kodeBooking, keterangan string) (isSuccessBatal bool)
	CheckIn(ctx context.Context, kodeBooking string, waktu int64) (isSuccess bool)
	CheckPasienDuplikat(ctx context.Context, noka string) (isDuplicate bool)
	InsertPasienBaru(ctx context.Context, pasienBaru antrian.AntrianOl) (
		isSuccess bool, norm string)
	GetJadwalOperasi(ctx context.Context, tanggalAwal, tanggalAkhir string) (
		jadopOls []antrian.JadopOl, err error)
	GetKodeBookingOperasiByNoPeserta(ctx context.Context, noPeserta string) (
		jadopOls []antrian.JadopOl, err error)
	CheckAntrean(ctx context.Context, nomorKartu, tglPeriksa, kodePoli string) (
		jumlah int64, err error)
	CheckDokterLibur(ctx context.Context, tglPeriksa string, kodeDokter string) (
		dokterLiburs antrian.LiburOl, err error)
	CheckJadwalPraktek(ctx context.Context, tglPeriksa string, idDokter string) (
		jadwal int64, err error)
	CheckKuota(ctx context.Context, tglPeriksa string, idDokter string, kuotaToday int) (isAvailable bool)
	CheckMedrek(ctx context.Context, noRM string) (dprofilpasien antrian.Dprofilpasien, err error)
	InsertAntreanMjkn(ctx context.Context, req dto.GetAntrianRequest, detailKTaripDokter antrian.KtaripDokter, kotaHariIni int, detailPoli antrian.Kpoli, detaiProfilPasien antrian.Dprofilpasien) (response dto.InsertPasienDTO, err error)
}
