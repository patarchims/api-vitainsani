package usecase

import (
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

	"github.com/sirupsen/logrus"
)

type antrianUseCase struct {
	antrianRepository entity.AntrianRepository
	IAntrianMapper    mapper.IAntrianMapper
	logging           *logrus.Logger
}

func NewAntrianUseCase(ar entity.AntrianRepository, IAntrianMapper mapper.IAntrianMapper, logging *logrus.Logger) entity.AntrianUseCase {
	return &antrianUseCase{
		antrianRepository: ar,
		IAntrianMapper:    IAntrianMapper,
		logging:           logging,
	}
}

func (au *antrianUseCase) GetStatusAntreanUsecase(payload *dto.StatusAntrianRequest, detailPoli antrian.Kpoli) (res dto.StatusAntreanDTO, err error) {

	detailKTaripDokter, err := au.antrianRepository.DetailTaripDokterByMapingAntrolRepository(payload.KodeDokter)

	if err != nil || detailKTaripDokter.Iddokter == "" {
		return res, errors.New("dokter tidak ditemukan")
	}

	lastCalled, _ := au.antrianRepository.LastCalledRepository(payload)

	jmlAntrean, _ := au.antrianRepository.JmlAntreanRepository(payload, detailKTaripDokter.Iddokter)

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

func (au *antrianUseCase) GetMobileJknByKodebookingUsecase(req dto.GetSisaAntrianRequest) (res map[string]any, err error) {
	m := map[string]any{}

	result, err := au.antrianRepository.GetMobileJknByKodebookingRepository(req.Kodebooking)
	if err != nil || result.NoAntrian == "" {
		return res, err
	}

	return m, nil

}

func (au *antrianUseCase) BatalAntreanUsecase(req dto.BatalAntreanRequest) (isSuccessBatal bool, err error) {

	antrean, err := au.antrianRepository.GetAntreanByKodeBookingRepository(req.Kodebooking)
	if err != nil {
		return false, err
	}

	isSuccess := au.antrianRepository.BatalAntreanRepository(antrean.NoBook, req.Keterangan)
	if !isSuccess {
		return false, errors.New("data gagal diupdate")
	}

	return true, nil
}

func (au *antrianUseCase) CheckedInUsecase(req dto.CheckInRequest) (isSuccess bool, err error) {

	isSuccess = au.antrianRepository.CheckInRepository(req.Kodebooking, req.Waktu)
	if !isSuccess {
		return false, errors.New("gagal update")
	}
	return true, nil
}

func (au *antrianUseCase) RegisterPasienBaruUsecase(req dto.RegisterPasienBaruRequest) (res dto.ResPasienBaru, err error) {

	// exists := au.antrianRepository.CheckPasienDuplikatRepository(req.Nomorkartu)
	exists := au.antrianRepository.CheckPasienDProfilePasienRepository(req.Nomorkartu)

	if exists {
		message := "data peserta sudah pernah dientrikan"
		au.logging.Info(message)
		return res, errors.New(message)
	}

	// GET ID BARU
	no, errs := au.antrianRepository.GetNormPasienRepository()

	if errs != nil {
		message := "gagal mendapatkan nomor rekam medis"
		au.logging.Info(message)
		return res, errors.New(message)
	}

	var jenisKelamin = ""

	switch req.Jeniskelamin {
	case "L":
		jenisKelamin = "Laki-Laki"
	case "P":
		jenisKelamin = "Perempuan"
	case "l":
		jenisKelamin = "Laki-Laki"
	case "p":
		jenisKelamin = "Perempuan"
	default:
		jenisKelamin = ""
	}

	// tgl, _ := time.Parse("2006-01-02", (req.Tanggallahir))
	// birthDate := time.Date(tgl.Year(), tgl.Month(), tgl.Day(), 0, 0, 0, 0, time.UTC)
	// currentDate := time.Now()
	// ageDuration := currentDate.Sub(birthDate)
	// Umurth:       int(ageInYears),
	// ageInYears := ageDuration.Hours() / 24 / 365

	var pasien = antrian.Dprofilpasien{
		Id:           no.Norm,
		Nik:          req.Nik,
		Nokapst:      req.Nomorkartu,
		Firstname:    req.Nama,
		Jeniskelamin: jenisKelamin,
		Alamat:       req.Alamat,
		Tgllahir:     req.Tanggallahir,
		Rtrw:         req.Rt + "/" + req.Rw,
		Kelurahan:    req.Namakel,
		Kecamatan:    req.Namakec,
		Propinsi:     req.Namaprop,
		Kabupaten:    req.Namadati2,
		Negara:       "Indonesia",
		Hp:           req.Nohp,
	}

	_, err2 := au.antrianRepository.InsertPasienBaruDprofilePasien(pasien)

	if err2 != nil {
		au.logging.Info("data Gagal disimpan")
		return res, errors.New("data Gagal disimpan")
	}

	res.Norm = no.Norm

	return res, nil
}

func (au *antrianUseCase) GetKodeBookingOperasiByNoPesertaUsecase(req dto.JadwalOperasiPasienRequest) (res map[string]any, err error) {

	m := map[string]any{}

	jadwalOperasiPeserta, err := au.antrianRepository.GetKodeBookingOperasiByNoPesertaRepository(req.Nopeserta)
	if err != nil {
		return nil, err
	}

	if len(jadwalOperasiPeserta) == 0 {
		message := fmt.Sprintf("nopeserta %s belum/tidak mempunyai kodebooking!", req.Nopeserta)
		return nil, errors.New(message)
	}

	jadwalOperasiMapper := au.IAntrianMapper.ToJadwalOperasiDTOMapper(jadwalOperasiPeserta, true)

	m["list"] = jadwalOperasiMapper

	return m, nil
}

func (au *antrianUseCase) AmbilAntreanUsecase(req dto.GetAntrianRequest, detailPoli antrian.Kpoli, detaiProfilPasien antrian.Dprofilpasien) (response dto.InsertPasienDTO, err error) {

	// validasi nomor antrean hanya boleh diambil satu kali pada tanggal dan poli yang sama
	alreadyGetAntrean, err := au.antrianRepository.CheckAntreanRepository(req.Nomorkartu, req.Tanggalperiksa, req.Kodepoli)
	if err != nil || alreadyGetAntrean > 0 {
		return response, err
	}

	// DETAIL KTARIPDOKTER
	detailKTaripDokter, err := au.antrianRepository.DetailTaripDokterByMapingAntrolRepository(req.Kodedokter)

	if err != nil || detailKTaripDokter.Iddokter == "" {
		str := strconv.Itoa(req.Kodedokter)
		message := fmt.Sprintf("Kode dokter %s tidak ditemukan", str)
		return response, errors.New(message)
	}

	fmt.Println("DETAIL DOKTER")
	fmt.Println(detailKTaripDokter)

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

	dokterCuti, _ := au.antrianRepository.CheckDokterLiburRepository(req.Tanggalperiksa, detailKTaripDokter.Iddokter)
	if dokterCuti.Keterangan != "" {
		message := fmt.Sprintf("Hari Libur %s ,  dalam rangka %s", dokterCuti.Deskripsi, dokterCuti.CatatanLibur)
		return response, errors.New(message)
	}

	// validasi poli tutup atau kuota habis
	checkKuota := au.antrianRepository.CheckKuotaRepository(req.Tanggalperiksa, detailKTaripDokter.Iddokter, kuota)

	if !checkKuota {
		message := fmt.Sprintf("%s sudah tutup, kuota habis", detailPoli.Namapoli)
		return response, errors.New(message)
	}

	result, err := au.antrianRepository.InsertAntreanMjknRepository(req, detailKTaripDokter, kuota, detailPoli, detaiProfilPasien)
	if err != nil {
		log.Fatal(err.Error())
		return response, err
	}

	return result, nil
}

// DATA VALIDATION
func (au *antrianUseCase) ValidasiDateUsecase(req string) (isTrue bool) {
	value := "((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])"
	re := regexp.MustCompile(value)
	return re.MatchString(req)
}
