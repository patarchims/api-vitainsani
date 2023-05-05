package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
	"vincentcoreapi/modules/farmasi"
	"vincentcoreapi/modules/farmasi/dto"
	"vincentcoreapi/modules/farmasi/entity"

	"gorm.io/gorm"
)

type farmasiRepository struct {
	DB *gorm.DB
}

func NewFarmasiRepository(db *gorm.DB) entity.FarmasiRepository {
	return &farmasiRepository{
		DB: db,
	}
}

func (fr *farmasiRepository) CekKodeBooking(ctx context.Context, req dto.GetAntreanFarmasiRequest) (res farmasi.AntreanOL, err error) {

	query := "SELECT * FROM rekam.antrian_ol WHERE no_book=? LIMIT 1"
	rs := fr.DB.WithContext(ctx).Raw(query, req.Kodebooking).Scan(&res)

	if rs.Error != nil {
		return res, err
	}

	if rs.RowsAffected > 0 {
		return res, err
	}

	return res, err
}

func (fr *farmasiRepository) InsertAntreanFarmasi(ctx context.Context, cekKodeBooking farmasi.AntreanOL) (res farmasi.AntreanResep, err error) {

	// GET LAST  TIME ELAPSED
	lastTime, _ := fr.LasTimeElapsed(ctx)
	var toTime time.Time
	// fmt.Println("LAST TIME")
	fmt.Println(lastTime)
	if lastTime.NoAntreanAngka > 0 {
		toTime = lastTime.TimeElapsed.Local().Add(time.Minute * 10)

	}

	if lastTime.NoAntreanAngka == 0 {
		toTime = time.Now()
	}

	// TIME ELAPSED,
	// WAKTU DI LAYANI

	// INSERT ANTREAN FARMASI
	waktu := time.Now().Format("2006-01-02")
	antre, err := fr.StatusAntreanFarmasi(ctx)
	fmt.Println(antre)
	var angka int
	var huruf string
	if err != nil || antre.Totalantrean == 0 {
		angka = 1
		huruf = "001"
	}

	count := 0

	if antre.Totalantrean > 0 {
		angka = antre.Totalantrean + 1
		for antre.Totalantrean > 0 {
			antre.Totalantrean = antre.Totalantrean / 10
			count++
		}
	}

	strAngka := strconv.Itoa(angka)

	switch count {
	case 0:
		huruf = "0001"
	case 1:
		huruf = "000" + strAngka
	case 2:
		huruf = "00" + strAngka
	case 3:
		huruf = "0" + strAngka
	default:
		huruf = strAngka
	}

	fmt.Println(angka)
	fmt.Println(huruf)

	// times := time.Now().Unix()
	// str := strconv.Itoa(int(times))
	waktuSekarang := strings.Split(waktu, "-")
	data := farmasi.AntreanResep{
		Mrn:            cekKodeBooking.ID,
		Cdttm:          time.Now(),
		JenisAntrean:   "online",
		JenisPasien:    "NRC",
		Tanggal:        time.Now(),
		Jam:            time.Now().Format("15:04:05"),
		TimeElapsed:    toTime,
		NoAntreanAngka: angka,
		NoAntrean:      huruf,
		KodeBookingRef: cekKodeBooking.NoBook,
		KodeBooking:    "NRC" + waktuSekarang[0] + waktuSekarang[1] + waktuSekarang[2] + huruf,
		Racikan:        "false",
		Dilayani:       "false",
		UpdDttm:        time.Now(),
	}

	fr.DB.Create(&data)

	return data, nil
}

func (fr *farmasiRepository) LasTimeElapsed(ctx context.Context) (res farmasi.TimeElapsed, err error) {
	now := time.Now().Format("2006-01-02")
	query := "SELECT time_elapsed , no_antrean_angka from posfar.antrean_resep WHERE    tanggal=? ORDER BY no_antrean_angka DESC  LIMIT 1"

	rs := fr.DB.WithContext(ctx).Raw(query, now).Scan(&res)

	fmt.Println(rs)

	if rs.Error != nil {
		return res, rs.Error
	}

	if rs.RowsAffected > 0 {
		return res, nil
	}

	return res, rs.Error
}

func (fr *farmasiRepository) CekKodeBookingAntreanResep(ctx context.Context, req dto.GetAntreanFarmasiRequest) (res farmasi.AntreanResep, err error) {

	value := farmasi.AntreanResep{}

	query := "SELECT * FROM posfar.antrean_resep WHERE  kode_booking_ref = ? LIMIT 1"

	rs := fr.DB.WithContext(ctx).Raw(query, req.Kodebooking).Scan(&value)

	fmt.Println(rs)

	if rs.Error != nil {
		return res, rs.Error
	}

	if rs.RowsAffected > 0 {
		return value, nil
	}

	return res, rs.Error
}

func (fr *farmasiRepository) StatusAntreanFarmasi(ctx context.Context) (res farmasi.StatusAntrean, err error) {
	// TANGGAL SEKARANG
	now := time.Now().Format("2006-01-02")

	fmt.Println(now)

	query := "SELECT no_antrean, no_antrean_angka, tanggal, kode_booking, SUM(dilayani=?) AS sisaantrean, SUM(dilayani=?) AS antreanpanggil, COUNT(dilayani) AS totalantrean FROM posfar.antrean_resep WHERE  tanggal=?"

	rs := fr.DB.WithContext(ctx).Raw(query, "false", "true", now).Scan(&res)

	fmt.Println(rs)

	if rs.Error != nil {
		return res, rs.Error
	}

	if rs.RowsAffected > 0 {
		return res, nil
	}

	return res, rs.Error
}
