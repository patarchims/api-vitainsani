package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
	"vincentcoreapi/modules/antrian"
	"vincentcoreapi/modules/antrian/dto"
	"vincentcoreapi/modules/antrian/entity"
	"vincentcoreapi/modules/antrian/mapper"
)

type antrianUseCase struct {
	antrianRepository entity.AntrianRepository
	IAntrianMapper    mapper.IAntrianMapper
}

func NewAntrianUseCase(ar entity.AntrianRepository, IAntrianMapper mapper.IAntrianMapper) entity.AntrianUseCase {
	return &antrianUseCase{
		antrianRepository: ar,
		IAntrianMapper:    IAntrianMapper,
	}
}

func (au *antrianUseCase) GetStatusAntrean(ctx context.Context, payload *dto.StatusAntrianRequest, detailPoli antrian.Kpoli) (res dto.StatusAntreanDTO, err error) {

	detailKTaripDokter, err := au.antrianRepository.DetailTaripDokterByMapingAntrol(ctx, payload.KodeDokter)
	if err != nil || detailKTaripDokter.Iddokter == "" {
		return res, errors.New("Dokter tidak ditemukan")
	}
	lastCalled, _ := au.antrianRepository.LastCalled(ctx, payload)

	jmlAntrean, _ := au.antrianRepository.JmlAntrean(ctx, payload, detailKTaripDokter.Iddokter)

	antreanPanggil := "-"
	if lastCalled.Nomorantrean != "" {
		antreanPanggil = lastCalled.Nomorantrean
	}

	tanggal2, _ := time.Parse("2006-01-02", payload.TanggalPeriksa)
	hari := strings.ToLower(tanggal2.Format("Mon"))

	var kuotaHari int

	switch hari {
	case "mon":
		kuotaHari = detailKTaripDokter.QuotaPasienMon
	case "tue":
		kuotaHari = detailKTaripDokter.QuotaPasienTue
	case "wed":
		kuotaHari = detailKTaripDokter.QuotaPasienWed
	case "thu":
		kuotaHari = detailKTaripDokter.QuotaPasienThu
	case "fri":
		kuotaHari = detailKTaripDokter.QuotaPasienFri
	case "sat":
		kuotaHari = detailKTaripDokter.QuotaPasienSat

	}

	sisaKuota := kuotaHari - jmlAntrean
	res.Namapoli = detailPoli.Namapoli
	res.Namadokter = detailKTaripDokter.Namadokter
	res.Totalantrean = jmlAntrean
	res.Sisaantrean = jmlAntrean
	res.Antreanpanggil = antreanPanggil
	res.Sisakuotajkn = sisaKuota
	res.Kuotajkn = kuotaHari
	res.Sisakuotanonjkn = sisaKuota
	res.Kuotanonjkn = kuotaHari
	res.Keterangan = ""

	return res, nil
}

func (au *antrianUseCase) GetMobileJknByKodebooking(ctx context.Context, req dto.GetSisaAntrianRequest) (res map[string]any, err error) {

	m := map[string]any{}

	result, err := au.antrianRepository.GetMobileJknByKodebooking(ctx, req.Kodebooking)
	if err != nil || result.NoAntrian == "" {
		return res, err
	}

	return m, nil

}

func (au *antrianUseCase) BatalAntrean(ctx context.Context, req dto.BatalAntreanRequest) (isSuccessBatal bool, err error) {

	antrean, err := au.antrianRepository.GetAntreanByKodeBooking(ctx, req.Kodebooking)
	if err != nil {
		return false, err
	}

	isSuccess := au.antrianRepository.BatalAntrean(ctx, antrean.NoBook, req.Keterangan)
	if !isSuccess {
		return false, errors.New("Data gagal diupdate")
	}

	return true, nil
}

func (au *antrianUseCase) CheckedIn(ctx context.Context, req dto.CheckInRequest) (isSuccess bool, err error) {

	isSuccess = au.antrianRepository.CheckIn(ctx, req.Kodebooking, req.Waktu)
	if !isSuccess {
		return false, errors.New("Gagal update")
	}
	return true, nil
}

func (au *antrianUseCase) RegisterPasienBaru(ctx context.Context, req dto.RegisterPasienBaruRequest) (res dto.ResPasienBaru, err error) {

	exists := au.antrianRepository.CheckPasienDuplikat(ctx, req.Nomorkartu)
	if exists {
		return res, errors.New("Data peserta sudah pernah dientrikan")
	}

	mapperPayload := au.IAntrianMapper.ToAntrianOlModel(req)

	isSuccess, norm := au.antrianRepository.InsertPasienBaru(ctx, mapperPayload)
	if !isSuccess {
		return res, errors.New("Gagal insert")
	}

	res.Norm = norm

	return res, nil
}

func (au *antrianUseCase) GetKodeBookingOperasiByNoPeserta(ctx context.Context, req dto.JadwalOperasiPasienRequest) (res map[string]any, err error) {

	m := map[string]any{}

	jadwalOperasiPeserta, err := au.antrianRepository.GetKodeBookingOperasiByNoPeserta(ctx, req.Nopeserta)
	if err != nil {
		return nil, err
	}

	if len(jadwalOperasiPeserta) == 0 {
		message := fmt.Sprintf("nopeserta %s belum/tidak mempunyai kodebooking!", req.Nopeserta)
		return nil, errors.New(message)
	}

	jadwalOperasiMapper := au.IAntrianMapper.ToJadwalOperasiDTO(jadwalOperasiPeserta, true)

	m["list"] = jadwalOperasiMapper

	return m, nil
}

func (au *antrianUseCase) AmbilAntrean(ctx context.Context, req dto.GetAntrianRequest, detailPoli antrian.Kpoli, detaiProfilPasien antrian.Dprofilpasien) (response dto.InsertPasienDTO, err error) {

	// validasi nomor antrean hanya boleh diambil satu kali pada tanggal dan poli yang sama
	alreadyGetAntrean, err := au.antrianRepository.CheckAntrean(ctx, req.Nomorkartu, req.Tanggalperiksa, req.Kodepoli)
	if err != nil || alreadyGetAntrean > 0 {
		return response, err
	}

	// DETAIL KTARIPDOKTER
	detailKTaripDokter, err := au.antrianRepository.DetailTaripDokterByMapingAntrol(ctx, req.Kodedokter)
	if err != nil || detailKTaripDokter.Iddokter == "" {
		str := strconv.Itoa(req.Kodedokter)
		message := fmt.Sprintf("Kode dokter %s tidak ditemukan", str)
		return response, errors.New(message)
	}

	var jumlahJadwal = detailKTaripDokter.Mon + detailKTaripDokter.Tue + detailKTaripDokter.Wed + detailKTaripDokter.Thu + detailKTaripDokter.Fri + detailKTaripDokter.Sat

	date, _ := time.Parse("2006-01-02", req.Tanggalperiksa)
	hari := strings.ToLower(date.Format("Mon"))

	var kuota int
	var jadwalToday int

	switch hari {
	case "mon":
		kuota = detailKTaripDokter.QuotaPasienMon
		jadwalToday = detailKTaripDokter.Mon
	case "tue":
		kuota = detailKTaripDokter.QuotaPasienTue
		jadwalToday = detailKTaripDokter.Tue
	case "wed":
		kuota = detailKTaripDokter.QuotaPasienWed
		jadwalToday = detailKTaripDokter.Wed
	case "thu":
		kuota = detailKTaripDokter.QuotaPasienThu
		jadwalToday = detailKTaripDokter.Thu
	case "fri":
		kuota = detailKTaripDokter.QuotaPasienFri
		jadwalToday = detailKTaripDokter.Fri
	case "sat":
		kuota = detailKTaripDokter.QuotaPasienSat
		jadwalToday = detailKTaripDokter.Sat
	case "sun":
		kuota = 0
		jadwalToday = detailKTaripDokter.Sun
	}

	if jumlahJadwal == 0 {
		message := fmt.Sprintf("Jadwal dokter %s belum tersedia, silahkan rescheduile tanggal dan jam praktek lainnya", detailKTaripDokter.Namadokter)
		return response, errors.New(message)
	}

	if jadwalToday == 0 {
		message := fmt.Sprintf("Pendaftaran ke %s Sedang Tutup", detailPoli.Namapoli)
		return response, errors.New(message)

	}

	dokterCuti, _ := au.antrianRepository.CheckDokterLibur(ctx, req.Tanggalperiksa, detailKTaripDokter.Iddokter)
	if dokterCuti.Keterangan != "" {
		message := fmt.Sprintf("Hari Libur %s ,  dalam rangka %s", dokterCuti.Deskripsi, dokterCuti.CatatanLibur)
		return response, errors.New(message)
	}

	// validasi poli tutup atau kuota habis
	checkKuota := au.antrianRepository.CheckKuota(ctx, req.Tanggalperiksa, detailKTaripDokter.Iddokter, kuota)

	fmt.Println("Check Kuota")
	fmt.Println(checkKuota)
	if !checkKuota {
		message := fmt.Sprintf("%s sudah tutup, kuota habis", detailPoli.Namapoli)
		return response, errors.New(message)
	}

	result, err := au.antrianRepository.InsertAntreanMjkn(ctx, req, detailKTaripDokter, kuota, detailPoli, detaiProfilPasien)
	if err != nil {
		log.Fatal(err.Error())
		return response, err
	}

	return result, nil
}

// DATA VALIDATION
func (au *antrianUseCase) ValidasiDate(ctx context.Context, req string) (isTrue bool) {
	value := "((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])"
	re := regexp.MustCompile(value)
	return re.MatchString(req)
}
