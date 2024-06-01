package repository

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"vincentcoreapi/modules/antrian"
	"vincentcoreapi/modules/antrian/dto"
	"vincentcoreapi/modules/antrian/entity"

	"gorm.io/gorm"
)

type antrianRepository struct {
	DB *gorm.DB
}

func NewAntrianRepository(db *gorm.DB) entity.AntrianRepository {
	return &antrianRepository{
		DB: db,
	}
}

func (ar *antrianRepository) ListAntrianTodayRepository() (res []antrian.AntrianOl, err error) {

	if err := ar.DB.Where("status = ? ", "tunggu").Limit(100).Take(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (ar *antrianRepository) CekPoliRepository(value string) (isTrue bool, err error) {
	kp := antrian.Kpoli{}

	query := "SELECT bpjs FROM his.kpoli WHERE bpjs = ? LIMIT 1"
	rs := ar.DB.Raw(query, value).Scan(&kp)
	if rs.Error != nil {
		log.Println("[Error Query]", rs.Error)
	}
	if rs.RowsAffected > 0 {
		return true, err
	}

	return false, err
}

func (ar *antrianRepository) DetailPoliRepository(params map[string]interface{}) (res map[string]interface{}, err error) {
	var result map[string]interface{}

	kp := antrian.Kpoli{}

	query := "SELECT namapoli,kodepoli FROM his.kpoli WHERE 1=1 AND bpjs = ? LIMIT 1"
	rs := ar.DB.Raw(query, params["kodepoli"]).Scan(&kp)

	if rs.Error != nil {
		log.Println("[Error Query]", rs.Error)
	}

	if rs.RowsAffected > 0 {
		m := map[string]interface{}{
			"namapoli": kp.Namapoli,
			"kodepoli": kp.Kodepoli,
		}
		result = m
	}

	return result, err
}

func (ar *antrianRepository) LastCalledRepository(payload *dto.StatusAntrianRequest) (res antrian.LastCalled, err error) {

	query := "SELECT nomorantrean, angkaantrean FROM antrean_ol_pool WHERE 1=1 AND tanggalperiksa = ? AND kodedokter = ? " +
		"ORDER BY angkaantrean DESC " +
		"LIMIT 1"

	rs := ar.DB.Raw(query, payload.TanggalPeriksa, payload.KodeDokter).Scan(&res)
	if rs.Error != nil {
		log.Println("[Error Query]", err)

	}

	return res, err
}

// CEK KTARIPDOKTER DUA
func (ar *antrianRepository) DetailKtaripDokter2AntrolRepository(idDokter string) (res antrian.KtaripDokter2, err error) {
	query := "SELECT iddokter,namadokter, quota_pasien, quota_pasien_mon, quota_pasien_tue, quota_pasien_wed, quota_pasien_thu, quota_pasien_fri,quota_pasien_sat FROM his.ktaripdokter2 WHERE 1=1 AND iddokter = ? "
	rs := ar.DB.Raw(query, idDokter).Scan(&res)
	if rs.Error != nil {
		return res, err
	}
	if rs.RowsAffected > 0 {
		return res, nil
	}

	return res, err
}

func (ar *antrianRepository) DetailTaripDokterByMapingAntrolRepository(mapingAntrol int) (res antrian.KtaripDokter, err error) {
	query := "SELECT * FROM his.ktaripdokter WHERE 1=1 AND maping_antrol = ? "
	rs := ar.DB.Raw(query, mapingAntrol).Scan(&res)
	if rs.Error != nil {
		return res, err
	}
	if rs.RowsAffected > 0 {
		return res, nil
	}

	return res, err
}

// GET KTARIP DOKTER DUA

func (ar *antrianRepository) GetKodeDokterRsRepository(params map[string]interface{}) (res map[string]interface{}, err error) {
	var result map[string]interface{}

	type ktaripdokter struct {
		IdDokter string `json:"iddokter"`
	}
	var kd ktaripdokter

	query := "SELECT iddokter FROM his.ktaripdokter WHERE 1=1 AND maping_antrol = ? "

	rs := ar.DB.Raw(query, params["kodedokter"]).Scan(&kd.IdDokter)
	if rs.Error != nil {
		log.Println("[Error Query]", err)
	}
	if rs.RowsAffected > 0 {
		m := map[string]interface{}{
			"iddokter": kd.IdDokter,
		}
		result = m
	}

	return result, err
}

func (ar *antrianRepository) JmlAntreanRepository(payload *dto.StatusAntrianRequest, kodeDokter string) (res int, err error) {

	type jmlAntrean struct {
		Jumlah int `json:"jumlah"`
	}
	var ja jmlAntrean

	query := "SELECT COUNT(*) AS Jumlah FROM antrian_ol WHERE 1=1 AND tgl_periksa = ? AND kd_dokter = ? AND status = 'tunggu' "
	rs := ar.DB.Raw(query, payload.TanggalPeriksa, kodeDokter).Scan(&ja)
	if rs.Error != nil {
		log.Println("[Error Query]", err)
	}

	return ja.Jumlah, err
}

func (ar *antrianRepository) GetKodeDokterJadwalRsRepository(day string, params map[string]interface{}) (res bool, err error) {
	type jmlAntrean struct {
		Jumlah int `json:"jumlah"`
	}
	var ja jmlAntrean

	query := "SELECT COUNT(*) AS Jumlah " +
		"FROM his.ktaripdokter " +
		"WHERE " + day + " = 1 AND maping_antrol = ? "

	rs := ar.DB.Raw(query, params["kodedokter"]).Scan(&ja)

	if rs.Error != nil {
		log.Println("[Error Query]", err)
	}

	if ja.Jumlah > 0 {
		return true, err
	} else {
		return false, err
	}
}

func (ar *antrianRepository) GetMobileJknByKodebookingRepository(kodebooking string) (
	res dto.GetMobileJknByKodebookingDTO, err error) {

	query := `
		SELECT 
			a.no_antrian,
			b.spesialisasi,
			a.tujuan,
			b.namadokter,
			b.maping_antrol,
			a.tgl_periksa
		FROM rekam.antrian_ol a
		JOIN his.ktaripdokter b ON b.iddokter = a.kd_dokter
		WHERE no_book = ?
	`

	result := ar.DB.Raw(query, kodebooking).Scan(&res)
	if result.Error != nil || res.NoAntrian == "" {
		log.Println("[Error Query]", err)
		return res, errors.New("antrean dengan kode booking tersebut tidak ditemukan")
	}

	return res, nil
}

func (ar *antrianRepository) GetSisaAntreanRepository(req dto.GetSisaAntrianRequest) (
	res dto.SisaAntreanResnonse, err error) {

	kodeBook := dto.GetMobileJknByKodebookingDTO{}

	query := `
				SELECT   *  FROM rekam.antrian_ol a
				JOIN his.ktaripdokter b ON b.iddokter = a.kd_dokter
				WHERE no_book = ?
			`

	result := ar.DB.Raw(query, req.Kodebooking).Scan(&kodeBook)

	if result.Error != nil || kodeBook.NoAntrian == "" {
		return res, errors.New("antrean dengan kode booking tersebut tidak ditemukan")
	}

	// PENAMBAHAN STATUS BATAL
	if kodeBook.Status == "batal" {
		return res, errors.New("antrean dengan kode booking tersebut sudah dibatalkan")
	}

	type SisaAntrean struct {
		Sisaantrean    int    `json:"sisa_antrean"`
		Antrianpanggil string `json:"antrian_panggil"`
		Waktutunggu    int    `json:"waktu_tunggu"`
	}

	var sisa SisaAntrean

	quers := `
	SELECT COUNT(*) AS sisaantrean ,(COUNT(*)*? )*60 AS waktutunggu,no_antrian AS antrianpanggil FROM rekam.antrian_ol
	WHERE kd_dokter=SUBSTRING_INDEX((SUBSTRING_INDEX(?,'-',2) ),'-',-1) AND CAST(no_antrian AS SIGNED INTEGER)<CAST(RIGHT(?,3) AS SIGNED INTEGER) AND STATUS='tunggu' AND tgl_periksa=CONCAT(SUBSTRING(SUBSTRING_INDEX(?,'-',-1) ,1,4),'-',SUBSTRING(SUBSTRING_INDEX(?,'-',-1) ,5,2),'-', SUBSTRING(SUBSTRING_INDEX(?,'-',-1) ,7,2))
	`
	value := ar.DB.Raw(quers, kodeBook.EstimasiPerPasien, req.Kodebooking, req.Kodebooking, req.Kodebooking, req.Kodebooking, req.Kodebooking).Scan(&sisa)

	if value.Error != nil {
		log.Println("[Error Query]", err)
		return res, errors.New("error query")
	}

	res.Nomorantrean = kodeBook.NoAntrian
	res.NamaPoli = kodeBook.Tujuan
	res.NamaDokter = kodeBook.Namadokter
	res.SisaAntrean = sisa.Sisaantrean
	res.AntreanPanggil = sisa.Antrianpanggil
	res.WaktuTunggu = sisa.Waktutunggu
	res.Keterangan = ""

	return res, nil
}

func (ar *antrianRepository) GetAntreanByKodeBookingRepository(kodeBooking string) (
	antrianOL antrian.AntrianOl, err error) {

	if err = ar.DB.Where("no_book = ?", kodeBooking).Take(&antrianOL).Error; err != nil {
		return antrianOL, errors.New("antrean tidak ditemukan")
	}

	if antrianOL.Status == "batal" {
		return antrianOL, errors.New("antrean tidak dapat ditemukan atau sudah Dibatalkan")
	}

	if antrianOL.Status == "proses" {
		return antrianOL, errors.New("pasien sudah dilayani antrean tidak dapat dibatalkan")
	}

	return antrianOL, nil
}

func (ar *antrianRepository) BatalAntreanRepository(kodeBooking, keterangan string) (isSuccessBatal bool) {

	if err := ar.DB.Model(antrian.AntrianOl{}).
		Where("no_book = ? ", kodeBooking).
		Updates(antrian.AntrianOl{
			Kunci:     "lock",
			Batal:     "true",
			Status:    "batal",
			KetUpdate: keterangan,
		}).Error; err != nil {

		return false
	}

	return true
}

func (ar *antrianRepository) CheckInRepository(kodeBooking string, waktu int64) (isSuccess bool) {

	t := time.UnixMilli(waktu)
	now := t.Format("2006-01-02 15:04:05")

	result := ar.DB.Model(antrian.AntrianOl{}).
		Where("no_book = ? ", kodeBooking).
		Updates(antrian.AntrianOl{CheckedIn: now})

	if result.Error != nil || result.RowsAffected == 0 {
		return false
	}

	return true
}

func (ar *antrianRepository) CheckPasienDuplikatRepository(noka string) (isDuplicate bool) {

	antrianOl := antrian.AntrianOl{}
	if err := ar.DB.Where("noka = ? ", noka).Take(&antrianOl).Error; err != nil {
		return false
	}

	return true
}

// CHECK PASIEN SUDAH ADA ATAU BELUM
func (ar *antrianRepository) CheckPasienDProfilePasienRepository(noka string) (isDuplicate bool) {
	// Dprofilpasien
	pasien := antrian.Dprofilpasien{}
	if err := ar.DB.Where("nokapst = ? ", noka).Take(&pasien).Error; err != nil {
		return false
	}

	return true
}

func (ar *antrianRepository) InsertPasienBaruRepository(pasienBaru antrian.AntrianOl) (isSuccess bool, norm string) {
	// pasienBaru.TimeElapsed =  Time.now
	result := ar.DB.Create(&pasienBaru)
	if result.Error != nil || result.RowsAffected == 0 {
		return false, ""
	}

	return true, "000000"
}

func (ar *antrianRepository) GetJadwalOperasiRepository(tanggalAwal, tanggalAkhir string) (
	jadopOls []antrian.JadopOl, err error) {

	if err = ar.DB.Where("tgl_operasi BETWEEN ? AND ? ", tanggalAwal, tanggalAkhir).
		Find(&jadopOls).Error; err != nil {

		return jadopOls, err
	}

	return jadopOls, nil
}

func (ar *antrianRepository) GetKodeBookingOperasiByNoPesertaRepository(noPeserta string) (
	jadopOls []antrian.JadopOl, err error) {

	if err = ar.DB.Where("noka = ? ", noPeserta).
		Find(&jadopOls).Error; err != nil {

		return jadopOls, err
	}

	return jadopOls, nil
}

func (ar *antrianRepository) CheckAntreanRepository(nomorKartu, tglPeriksa, kodePoli string) (
	jumlah int64, err error) {

	result := ar.DB.Table("antrian_ol a").
		Select("*").
		Joins("JOIN his.kpoli b ON b.kodepoli = a.kode_tujuan").
		Where("noka = ?  AND b.bpjs = ? AND  a.status =? AND a.reg_type=?", nomorKartu, kodePoli, "tunggu", "online").
		Count(&jumlah)
	if result.Error != nil {
		return jumlah, errors.New("data tidak di temukan")
	}

	if jumlah > 0 {
		return jumlah, errors.New("nomor antrean hanya dapat diambil 1 kali pada tanggal yang sama")
	}

	return jumlah, nil

}

func (ar *antrianRepository) CheckDokterLiburRepository(tglPeriksa string, kodeDokter string) (
	dokterLiburs antrian.LiburOl, err error) {

	if err = ar.DB.Where("keterangan = ? AND tanggal = ?", kodeDokter, tglPeriksa).Find(&dokterLiburs).Error; err != nil {
		return dokterLiburs, err
	}

	return dokterLiburs, nil
}

func (ar *antrianRepository) CheckJadwalPraktekRepository(tglPeriksa string, idDokter string) (
	jadwal int64, err error) {

	date, _ := time.Parse("2006-01-02", tglPeriksa)
	// Ubah Format hari menjadi loweercase
	hari := strings.ToLower(date.Format("Mon"))

	fmt.Println("Hari")
	fmt.Println(hari)

	query := `
		SELECT count(*)
		FROM his.ktaripdokter
		WHERE ` + hari + ` = 1 AND iddokter = ?
	`
	// JUMLAH
	result := ar.DB.Raw(query, idDokter).Scan(&jadwal)
	if result.Error != nil {
		log.Println("[Error Query]", err)
		return jadwal, err
	}

	return jadwal, nil
}

func (ar *antrianRepository) GetDokterNameRepository(kodeDokter int) (
	dokter antrian.KtaripDokter, err error) {

	if err := ar.DB.Where("maping_antrol = ? ", kodeDokter).Take(&dokter).Error; err != nil {
		return dokter, err
	}

	return dokter, nil
}

func (ar *antrianRepository) CheckKuotaRepository(tglPeriksa string, idDokter string, kuotaToday int) (isAvailable bool) {
	var jmlhDaftar int64
	var antrianOl antrian.AntrianOl

	resJmlhDaftar := ar.DB.Where("kd_dokter = ? AND tgl_periksa = ? AND kode_debitur = ?", idDokter, tglPeriksa, "BPJS").Find(&antrianOl).Count(&jmlhDaftar)

	if resJmlhDaftar.Error != nil {
		return false
	}

	fmt.Println("JUMLAH DAFTAR")
	fmt.Println(jmlhDaftar)

	// kuota pasien
	var ktaripdokter antrian.KtaripDokter

	resKuota := ar.DB.Where("iddokter = ?", idDokter).Find(&ktaripdokter)

	if resKuota.Error != nil {
		return false
	}

	if int64(kuotaToday) <= jmlhDaftar {
		return false
	}

	return true
}

func (ar *antrianRepository) CheckJamStartRepository(tglPeriksa string, kodeDokter int, jamPrakterk string) (jamTutup string) {

	date, _ := time.Parse("2006-01-02", tglPeriksa)
	number := strings.ToLower(date.Format("Mon"))

	var hari string
	switch number {
	case "mon":
		hari = "senin"
	case "tue":
		hari = "selasa"
	case "wed":
		hari = "rabu"
	case "thu":
		hari = "kamis"
	case "fri":
		hari = "jumat"
	case "sat":
		hari = "sabtu"
	}

	params := map[string]any{}
	params["kodedokter"] = kodeDokter

	query := `
		SELECT ` + hari + ` AS jampraktek
		FROM his.ktaripdokter
		WHERE maping_antrol = ? AND ` + hari + ` LIKE ?
	`

	type Result struct {
		Jampraktek string
	}
	ktaripdokter := Result{}
	result := ar.DB.Raw(query, kodeDokter, jamPrakterk+"%").Scan(&ktaripdokter)

	if result.Error != nil || ktaripdokter.Jampraktek == "" {
		return ""
	}

	jam := strings.Split(ktaripdokter.Jampraktek, "|")
	jamTutup = jam[1]

	return jamTutup
}

func (ar *antrianRepository) CheckMedrekRepository(nik string) (dprofilpasien antrian.Dprofilpasien, err error) {
	query1 := `SELECT * FROM his.dprofilpasien WHERE nik= ?`
	err = ar.DB.Raw(query1, nik).Scan(&dprofilpasien).Error

	if err != nil {
		return dprofilpasien, err
	}

	return dprofilpasien, nil
}

func (ar *antrianRepository) InsertAntreanMjknRepository(req dto.GetAntrianRequest, detailKTaripDokter antrian.KtaripDokter, kotaHariIni int, detailPoli antrian.Kpoli, detaiProfilPasien antrian.Dprofilpasien) (response dto.InsertPasienDTO, err error) {

	// TODO : RUMAH SAKIT VITA INSANI

	query1 := `
		SELECT COALESCE(LPAD(CONVERT(@last_no_antrian :=MAX(no_antrian),SIGNED INTEGER)+1,3,0),'001') AS no_antre,
			CONCAT('RSVI-', ? ,'-',REPLACE(?,'-',''),COALESCE(LPAD(CONVERT(@last_no_antrian :=MAX(no_antrian),SIGNED INTEGER)+1,3,0),'001')) AS kobook,
			? AS date
		FROM rekam.antrian_ol
		WHERE tgl_periksa = ?
		AND kd_dokter = ?

	`
	type Result1 struct {
		NoAntre string
		Kobook  string
		Date    string
	}
	result1 := Result1{}
	err = ar.DB.Raw(query1, detailKTaripDokter.Iddokter, req.Tanggalperiksa, req.Tanggalperiksa, req.Tanggalperiksa, detailKTaripDokter.Iddokter).
		Scan(&result1).Error

	if err != nil {
		fmt.Println("ERROR QUERY RESULT")
		log.Fatal(err.Error())
		return
	}

	// Jumlah Antrian
	antrians := []antrian.AntrianOl{}
	if err = ar.DB.Where("status = ? AND tgl_periksa = ? AND kd_dokter = ? AND kode_debitur=? ", "tunggu", req.Tanggalperiksa, detailKTaripDokter.Iddokter, "BPJS").Find(&antrians).Error; err != nil {
		return response, err
	}
	fmt.Println("Jumlah Antrean")
	fmt.Println(len(antrians))
	jumlahAntrian := len(antrians) + 1

	// GET TIME ELAPSED HERE
	timeElapsed := ar.getTimeElapsed(req.Tanggalperiksa, strconv.Itoa(req.Kodedokter), int64(detailKTaripDokter.EstimasiPerPasien), int64(jumlahAntrian), detailKTaripDokter)

	data2 := antrian.AntrianOl{
		// SIMPAN DARI DPROFILE PASIEN
		Dob:                detaiProfilPasien.Tgllahir[0:10],
		Nama:               detaiProfilPasien.Firstname,
		Alamat:             detaiProfilPasien.Alamat,
		Id:                 detaiProfilPasien.Id,
		Gender:             detaiProfilPasien.Jeniskelamin,
		TimeElapsed:        timeElapsed,
		BookingOnsite:      "false",
		Notif:              "false",
		JenisAntreanPasien: "0",
		Batal:              "false",
		RegType:            "online",
		Kunci:              "unlock",
		Status:             "tunggu",
		Proses:             "false",
		Nik:                req.Nik,
		Noka:               req.Nomorkartu,
		Jeniskunjungan:     strconv.Itoa(req.Jeniskunjungan),
		NoRujukan:          req.Nomorreferensi,
		// SIMPAN NO HP KE ANTRIAN OL
		NoHp:            req.Nohp,
		KodeTujuan:      detailPoli.Kodepoli,
		Tujuan:          detailPoli.Namapoli,
		TglPeriksa:      req.Tanggalperiksa,
		KdDokter:        detailKTaripDokter.Iddokter,
		Dokter:          detailKTaripDokter.Namadokter,
		NoAntrian:       result1.NoAntre,
		KodeDebitur:     "BPJS",
		Debitur:         "JKN-BPJS",
		BookDate:        time.Now().Format("2006-01-02 15:04:05"),
		NoBook:          result1.Kobook,
		JknNomorkk:      "0000000000000000",
		JknTanggallahir: "000-00-00",
		JknKodeprop:     "00",
		JknKodedati2:    "0000",
		JknNamadati2:    "kosong",
		JknKodekec:      "0000",
		JknNamakec:      "kosong",
		JknKodekel:      "0000",
		JknNamakel:      "kosong",
		JknRw:           "00",
		JknRt:           "00",
		CheckedIn:       "",
	}

	ar.DB.Create(&data2)
	angkaAntrean, _ := strconv.Atoi(result1.NoAntre)

	response.Nomorantrean = result1.NoAntre
	response.Angkaantrean = angkaAntrean

	date, _ := time.Parse("2006-01-02 15:04:05", timeElapsed)

	// LAKUKAN PERUBAHAN ETIMASI DILAYANI
	theTime := time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), time.Local)

	sisaKuota := kotaHariIni - jumlahAntrian
	tanggal := theTime.Unix() * 1000
	response.Estimasidilayani = int(tanggal)
	response.Kodebooking = result1.Kobook
	response.Namapoli = detailPoli.Namapoli
	response.Namadokter = detailKTaripDokter.Namadokter
	response.Sisakuotajkn = sisaKuota
	response.Kuotajkn = kotaHariIni
	response.Sisakuotanonjkn = sisaKuota
	response.Kuotanonjkn = kotaHariIni
	response.Norm = detaiProfilPasien.Id
	response.Keterangan = "Peserta harap 60 menit lebih awal guna pencatatan administrasi."

	return response, nil
}

// TIMEELAPSED
func (ar *antrianRepository) getTimeElapsed(tglPeriksa, _ string, estimasiPerPasien, jumlahAntrean int64, detailKTaripDokter antrian.KtaripDokter) string {
	date, _ := time.Parse("2006-01-02", tglPeriksa)
	number := strings.ToLower(date.Format("Mon"))

	var hari string
	switch number {
	case "mon":
		hari = detailKTaripDokter.Senin
	case "tue":
		hari = detailKTaripDokter.Selasa
	case "wed":
		hari = detailKTaripDokter.Rabu
	case "thu":
		hari = detailKTaripDokter.Kamis
	case "fri":
		hari = detailKTaripDokter.Jumat
	case "sat":
		hari = detailKTaripDokter.Sabtu
	}

	jam := strings.Split(hari, "|")
	jambuka := jam[0]

	minutesToAdd := jumlahAntrean * estimasiPerPasien

	timestampPelayanan := fmt.Sprintf("%s %s", tglPeriksa, jambuka)

	/**
	 * jika daftar setelah jam buka maka $estimasi = $minutes_to_add + $jam_daftar
	 * jika sebelumnya maka $estimasi = $minutes_to_add + $jam_buka
	 */

	jamPraktek, _ := time.Parse("2006-01-02 15:04:05", timestampPelayanan)
	fmt.Println("Jam Praktek")
	fmt.Println(jamPraktek)

	if minutesToAdd == 0 {
		return timestampPelayanan
	} else {
		var timein time.Time
		timein = jamPraktek.Add(time.Minute * time.Duration(minutesToAdd))
		waktu := timein.Format("2006-01-02 15:04:05")
		return waktu
	}

}

func (ar *antrianRepository) CariPoliRepository(kdPoli string) (res antrian.Kpoli, err error) {
	query := "SELECT * FROM his.kpoli WHERE bpjs = ? LIMIT 1"
	rs := ar.DB.Raw(query, kdPoli).Scan(&res)
	if rs.Error != nil {
		return res, err
	}
	if rs.RowsAffected > 0 {
		return res, err
	}
	return res, err
}

func (ar *antrianRepository) GetNormPasienRepository() (res antrian.IDPasien, err error) {

	query := `SELECT COALESCE(LPAD(CONVERT(@last_id :=MAX(id),SIGNED INTEGER)+1,6,0),'000001') AS norm FROM his.dprofilpasien WHERE length(id)=6;`

	result := ar.DB.Raw(query).Scan(&res)

	if result.Error != nil {
		message := fmt.Sprintf("Error %s, Data tidak dapat dieksekusi", result.Error.Error())
		return res, errors.New(message)
	}

	return res, nil
}

// ============================== PENAMBAHAN FITUR
// ======================== INSERT PASIEN BARU, KE DPROFILPASIEN
func (ar *antrianRepository) InsertPasienBaruDprofilePasien(pasienBaru antrian.Dprofilpasien) (res antrian.Dprofilpasien, err error) {
	result := ar.DB.Create(&pasienBaru).Scan(&res)

	if result.Error != nil || result.RowsAffected == 0 {
		return res, result.Error
	}

	return res, nil
}
