package repository

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
	"vincentcoreapi/modules/antrian"
	"vincentcoreapi/modules/antrian/dto"
)

func (ar *antrianRepository) LastCalledRepositoryV2(payload *dto.StatusAntrianRequestV2) (res antrian.LastCalled, err error) {

	query := "SELECT nomorantrean, angkaantrean " +
		"FROM antrean_ol_pool " +
		"WHERE 1=1 AND tanggalperiksa = ? AND kodedokter = ? " +
		"ORDER BY angkaantrean DESC " +
		"LIMIT 1"
	rs := ar.DB.Raw(query, payload.TanggalPeriksa, payload.KodeDokter).Scan(&res)
	if rs.Error != nil {
		log.Println("[Error Query]", rs.Error)
	}

	return res, nil
}

func (ar *antrianRepository) JmlAntreanRepositoryV2(payload *dto.StatusAntrianRequestV2, kodeDokter string) (res int, err error) {

	type jmlAntrean struct {
		Jumlah int `json:"jumlah"`
	}
	var ja jmlAntrean

	query := "SELECT COUNT(*) AS Jumlah " +
		"FROM antrian_ol " +
		"WHERE 1=1 AND tgl_periksa = ? AND kd_dokter = ? AND status = 'tunggu' "
	rs := ar.DB.Raw(query, payload.TanggalPeriksa, kodeDokter).Scan(&ja)
	if rs.Error != nil {
		log.Println("[Error Query]", rs.Error)
	}

	return ja.Jumlah, nil
}

func (ar *antrianRepository) GetSisaAntreanRepositoryV2(req dto.GetSisaAntrianRequestV2) (
	res dto.SisaAntreanResnonse, err error) {

	kodeBook := dto.GetMobileJknByKodebookingDTO{}

	query := `
				SELECT   *  FROM rekam.antrian_ol a
				JOIN his.ktaripdokter b ON b.iddokter = a.kd_dokter
				WHERE no_book = ?
			`

	result := ar.DB.Raw(query, req.Kodebooking).Scan(&kodeBook)

	if result.Error != nil || kodeBook.NoAntrian == "" {
		log.Println("[Error Query]", err)
		return res, errors.New("antrean dengan kode booking tersebut tidak ditemukan")
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
		log.Println("[Error Query]", value.Error)
		return res, value.Error
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

func (ar *antrianRepository) InsertAntreanMjknRepositoryV2(req dto.GetAntrianRequestV2, detailKTaripDokter antrian.KtaripDokter, kotaHariIni int, detailPoli antrian.Kpoli, detaiProfilPasien antrian.Dprofilpasien, umum int) (response dto.InsertPasienDTO, err error) {

	// TODO : RUMAH SAKIT VITA INSANI

	// query1 := `
	// 		SELECT COALESCE(LPAD(CONVERT(@last_no_antrian :=MAX(no_antrian),SIGNED INTEGER)+1,3,0),'001') AS no_antre,CONCAT('VIAJ-',?,'-',REPLACE(CURDATE(),'-',''), COALESCE(LPAD(CONVERT(@last_no_antrian :=MAX(no_antrian),SIGNED INTEGER)+1,3,0),'001')) AS kobook, CURDATE() AS date,CURTIME() AS time FROM rekam.antrian_ol  WHERE tgl_periksa=? AND kd_dokter=? AND kode_debitur='BPJS'
	// `

	query1 := `
		SELECT COALESCE(LPAD(CONVERT(@last_no_antrian :=MAX(no_antrian),SIGNED INTEGER)+1,3,0), ?) AS no_antre,
			CONCAT('VIAJ-', ? ,'-',REPLACE(?,'-',''),COALESCE(LPAD(CONVERT(@last_no_antrian :=MAX(no_antrian),SIGNED INTEGER)+1,3,0), ?)) AS kobook,
			? AS date
		FROM rekam.antrian_ol
		WHERE tgl_periksa = ? AND kd_dokter = ? AND kode_debitur='BPJS'
	`

	type Result1 struct {
		NoAntre string
		Kobook  string
		Date    string
		// Time    string
	}

	result1 := Result1{}

	// err = ar.DB.Raw(query1, detailKTaripDokter.Iddokter, detailKTaripDokter.Iddokter).
	// 	Scan(&result1).Error

	err = ar.DB.Raw(query1, umum, detailKTaripDokter.Iddokter, req.Tanggalperiksa, umum, req.Tanggalperiksa, req.Tanggalperiksa, detailKTaripDokter.Iddokter).
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
		Dob:                detaiProfilPasien.Tgllahir,
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
		NoHp:               req.Nohp,
		KodeTujuan:         detailPoli.Kodepoli,
		Tujuan:             detailPoli.Namapoli,
		TglPeriksa:         req.Tanggalperiksa,
		KdDokter:           detailKTaripDokter.Iddokter,
		Dokter:             detailKTaripDokter.Namadokter,
		NoAntrian:          result1.NoAntre,
		KodeDebitur:        "BPJS",
		Debitur:            "JKN-BPJS",
		BookDate:           time.Now().Format("2006-01-02 15:04:05"),
		NoBook:             result1.Kobook,
		JknNomorkk:         "0000000000000000",
		JknTanggallahir:    "000-00-00",
		JknKodeprop:        "00",
		JknKodedati2:       "0000",
		JknNamadati2:       "kosong",
		JknKodekec:         "0000",
		JknNamakec:         "kosong",
		JknKodekel:         "0000",
		JknNamakel:         "kosong",
		JknRw:              "00",
		JknRt:              "00",
		CheckedIn:          "",
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
