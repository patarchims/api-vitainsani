package handler

import (
	"errors"
	"strconv"
	"time"
	"vincentcoreapi/modules/antrian/dto"
)

const (
	emptyNomorKartu    = "Nomor kartu belum diisi"
	emptyNik           = "NIK belum diisi"
	emptyKK            = "Nomor KK belum diisi"
	emptyName          = "Nama belum diisi"
	emptyJK            = "Jenis Kelamin belum dipilih"
	emptyTglLahir      = "Tanggal lahir belum diisi"
	emptyAlamat        = "Alamat belum diisi"
	emptyProvinsi      = "Kode propinsi belum diisi"
	emptyNamaProvinsi  = "Nama propinsi belum diisi"
	emptyKodeDati2     = "Kode dati 2 belum diisi"
	emptyDati2         = "Dati 2 belum diisi"
	emptyKodeKecamatan = "Kode kecamatan belum diisi"
	emptyNamaKecamatan = "Nama kecamatan belum diisi"
	emptyKodeKelurahan = "Kode kelurahan belum diisi"
	emptyNamaKelurahan = "Nama kelurahan belum diisi"
	emptyRW            = "RW belum diisi"
	emptyRT            = "RT belum diisi"

	errorNomorKartu         = "Nomor Kartu tidak sesuai"
	errorNik                = "Format NIK tidak sesuai"
	errorFormatTanggalLahir = "Format Tanggal lahir tidak sesuai"
)

func validationPayloadPasienBaru(payload dto.RegisterPasienBaruRequest) (err error) {

	if payload.Nomorkartu == "" {
		return errors.New(emptyNomorKartu)
	}
	if len(payload.Nomorkartu) != 13 {
		return errors.New(errorNomorKartu)
	}

	if _, err := strconv.Atoi(payload.Nomorkartu); err != nil {
		return errors.New(errorNomorKartu)
	}

	if payload.Nik == "" {
		return errors.New(emptyNik)
	}

	if len(payload.Nik) != 16 {
		return errors.New(errorNik)
	}

	if _, err := strconv.Atoi(payload.Nik); err != nil {
		return errors.New(errorNik)
	}

	if payload.Nomorkk == "" || len(payload.Nomorkk) == 0 {
		return errors.New(emptyKK)
	}

	if payload.Nama == "" || len(payload.Nama) == 0 {
		return errors.New(emptyName)
	}

	if payload.Jeniskelamin == "" || len(payload.Jeniskelamin) == 0 {
		return errors.New(emptyJK)
	}

	if payload.Tanggallahir == "" || len(payload.Tanggallahir) == 0 {
		return errors.New(emptyTglLahir)
	}

	if payload.Alamat == "" || len(payload.Alamat) == 0 {
		return errors.New(emptyAlamat)
	}

	if payload.Kodeprop == "" || len(payload.Kodeprop) == 0 {
		return errors.New(emptyProvinsi)
	}

	if payload.Namaprop == "" || len(payload.Namaprop) == 0 {
		return errors.New(emptyNamaProvinsi)
	}

	if payload.Kodedati2 == "" || len(payload.Kodedati2) == 0 {
		return errors.New(emptyKodeDati2)
	}

	if payload.Namadati2 == "" || len(payload.Namadati2) == 0 {
		return errors.New(emptyDati2)
	}

	if payload.Kodekec == "" || len(payload.Kodekec) == 0 {
		return errors.New(emptyKodeKecamatan)
	}

	if payload.Namakec == "" || len(payload.Namakec) == 0 {
		return errors.New(emptyNamaKecamatan)
	}

	if payload.Kodekel == "" || len(payload.Kodekel) == 0 {
		return errors.New(emptyKodeKelurahan)
	}

	if payload.Namakel == "" || len(payload.Namakel) == 0 {
		return errors.New(emptyNamaKelurahan)
	}

	if payload.Rw == "" || len(payload.Rw) == 0 {
		return errors.New(emptyRW)
	}

	if payload.Rt == "" || len(payload.Rt) == 0 {
		return errors.New(emptyRT)
	}

	ok := isDateValue(payload.Tanggallahir)
	if !ok {
		return errors.New(errorFormatTanggalLahir)
	}

	ok = compareDate(payload.Tanggallahir)
	if !ok {
		return errors.New(errorFormatTanggalLahir)
	}

	return nil
}

func isDateValue(stringDate string) bool {
	_, err := time.Parse("2006-01-02", stringDate)
	return err == nil
}

func compareDate(stringDate string) bool {
	now := time.Now()
	t, _ := time.Parse("2006-01-02", stringDate)

	if t.Before(now) {
		return true
	}

	return false
}
