package usecase

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"vincentcoreapi/modules/antrian"
	"vincentcoreapi/modules/antrian/dto"
)

func (au *antrianUseCase) GetStatusAntreanUsecaseV2(payload *dto.StatusAntrianRequestV2, detailPoli antrian.Kpoli) (res dto.StatusAntreanDTO, err error) {

	detailKTaripDokter, err := au.antrianRepository.DetailTaripDokterByMapingAntrolRepository(payload.KodeDokter)

	if err != nil || detailKTaripDokter.Iddokter == "" {
		return res, errors.New("dokter tidak ditemukan")
	}

	lastCalled, _ := au.antrianRepository.LastCalledRepositoryV2(payload)

	jmlAntrean, _ := au.antrianRepository.JmlAntreanRepositoryV2(payload, detailKTaripDokter.Iddokter)

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

func (au *antrianUseCase) BatalAntreanUsecaseV2(req dto.BatalAntreanRequestV2) (isSuccessBatal bool, err error) {
	antrean, err := au.antrianRepository.GetAntreanByKodeBookingRepository(req.Kodebooking)

	if err != nil {
		return false, err
	}

	isSuccess := au.antrianRepository.BatalAntreanRepository(antrean.NoBook, req.Keterangan)

	if !isSuccess {
		return false, errors.New("Data gagal diupdate")
	}

	return true, nil
}

func (au *antrianUseCase) GetKodeBookingOperasiByNoPesertaUsecaseV2(req dto.JadwalOperasiPasienRequestV2) (res map[string]any, err error) {

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

func (au *antrianUseCase) AmbilAntreanUsecaseV2(req dto.GetAntrianRequestV2, detailPoli antrian.Kpoli, detaiProfilPasien antrian.Dprofilpasien) (response dto.InsertPasienDTO, err error) {

	alreadyGetAntrean, err := au.antrianRepository.CheckAntreanRepository(req.Nomorkartu, req.Tanggalperiksa, req.Kodepoli)

	if err != nil || alreadyGetAntrean > 0 {
		return response, err
	}

	detailKTaripDokter, err := au.antrianRepository.DetailTaripDokterByMapingAntrolRepository(req.Kodedokter)

	if err != nil || detailKTaripDokter.Iddokter == "" {
		str := strconv.Itoa(req.Kodedokter)
		message := fmt.Sprintf("Kode dokter %s tidak ditemukan", str)
		return response, errors.New(message)
	}

	var jumlahJadwal = detailKTaripDokter.Mon + detailKTaripDokter.Tue + detailKTaripDokter.Wed + detailKTaripDokter.Thu + detailKTaripDokter.Fri + detailKTaripDokter.Sat

	date, _ := time.Parse("2006-01-02", req.Tanggalperiksa)
	au.logging.Info(date)
	hari := strings.ToLower(date.Format("Mon"))
	au.logging.Info(hari)
	ktaripDokter2, _ := au.antrianRepository.DetailKtaripDokter2AntrolRepository(detailKTaripDokter.Iddokter)

	var kuota int
	var jadwalToday int
	var kuotaUmum int

	switch hari {
	case "mon":
		kuota = detailKTaripDokter.QuotaPasienMon
		jadwalToday = detailKTaripDokter.Mon
		kuotaUmum = ktaripDokter2.QuotaPasienMon
	case "tue":
		kuota = detailKTaripDokter.QuotaPasienTue
		jadwalToday = detailKTaripDokter.Tue
		kuotaUmum = ktaripDokter2.QuotaPasienTue
	case "wed":
		kuota = detailKTaripDokter.QuotaPasienWed
		jadwalToday = detailKTaripDokter.Wed
		kuotaUmum = ktaripDokter2.QuotaPasienWed
	case "thu":
		kuota = detailKTaripDokter.QuotaPasienThu
		jadwalToday = detailKTaripDokter.Thu
		kuotaUmum = ktaripDokter2.QuotaPasienThu
	case "fri":
		kuota = detailKTaripDokter.QuotaPasienFri
		jadwalToday = detailKTaripDokter.Fri
		kuotaUmum = ktaripDokter2.QuotaPasienFri
	case "sat":
		kuota = detailKTaripDokter.QuotaPasienSat
		jadwalToday = detailKTaripDokter.Sat
		kuotaUmum = ktaripDokter2.QuotaPasienSat
	case "sun":
		kuota = 0
		jadwalToday = detailKTaripDokter.Sun
		kuotaUmum = 0
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

	checkKuota := au.antrianRepository.CheckKuotaRepository(req.Tanggalperiksa, detailKTaripDokter.Iddokter, kuota)

	if !checkKuota {
		message := fmt.Sprintf("%s sudah tutup, kuota habis", detailPoli.Namapoli)
		return response, errors.New(message)
	}

	// KOUTA SAAT INI
	var saatIniKuoata = ""

	if lenLoop(kuotaUmum) == 3 {
		saatIniKuoata = strconv.Itoa(kuotaUmum + 1)
	}

	if lenLoop(kuotaUmum) == 2 {
		saatIniKuoata = "0" + strconv.Itoa(kuotaUmum+1)
	}

	if lenLoop(kuotaUmum) == 1 {
		saatIniKuoata = "00" + strconv.Itoa(kuotaUmum+1)
	}

	result, err := au.antrianRepository.InsertAntreanMjknRepositoryV2(req, detailKTaripDokter, kuota, detailPoli, detaiProfilPasien, saatIniKuoata)

	if err != nil {
		log.Fatal(err.Error())
		return response, err
	}

	return result, nil
}

func lenLoop(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}
