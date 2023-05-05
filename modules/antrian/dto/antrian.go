package dto

import "time"

type (
	GetSisaAntrianRequest struct {
		Kodebooking string `json:"kodebooking" binding:"required"`
	}

	ResPasienBaru struct {
		Norm string `json:"norm"`
	}

	SisaAntreanResnonse struct {
		Nomorantrean   string `json:"nomorantrean"`
		NamaPoli       string `json:"namapoli"`
		NamaDokter     string `json:"namadokter"`
		SisaAntrean    int    `json:"sisaantrean"`
		AntreanPanggil string `json:"antreanpanggil"`
		WaktuTunggu    int    `json:"waktutunggu"`
		Keterangan     string `json:"keterangan"`
	}

	StatusAntrianRequest struct {
		KodePoli       string `json:"kodepoli" binding:"required"`
		TanggalPeriksa string `json:"tanggalperiksa" binding:"required"`
		KodeDokter     int    `json:"kodedokter" binding:"required"`
		JamPraktek     string `json:"jampraktek" binding:"required"`
	}

	BatalAntreanRequest struct {
		Kodebooking string `json:"kodebooking" binding:"required"`
		Keterangan  string `json:"keterangan" binding:"required"`
	}

	CheckInRequest struct {
		Kodebooking string `json:"kodebooking" binding:"required"`
		Waktu       int64  `json:"waktu" binding:"required"`
	}

	RegisterPasienBaruRequest struct {
		Nomorkartu   string `json:"nomorkartu"`
		Nik          string `json:"nik"`
		Nomorkk      string `json:"nomorkk"`
		Nama         string `json:"nama"`
		Jeniskelamin string `json:"jeniskelamin"`
		Tanggallahir string `json:"tanggallahir"`
		Nohp         string `json:"nohp"`
		Alamat       string `json:"alamat"`
		Kodeprop     string `json:"kodeprop"`
		Namaprop     string `json:"namaprop"`
		Kodedati2    string `json:"kodedati2"`
		Namadati2    string `json:"namadati2"`
		Kodekec      string `json:"kodekec"`
		Namakec      string `json:"namakec"`
		Kodekel      string `json:"kodekel"`
		Namakel      string `json:"namakel"`
		Rw           string `json:"rw"`
		Rt           string `json:"rt"`
	}

	JadwalOperasiRequest struct {
		Tanggalawal  string `json:"tanggalawal" binding:"required"`
		Tanggalakhir string `json:"tanggalakhir" binding:"required"`
	}

	JadwalOperasiPasienRequest struct {
		Nopeserta string `json:"nopeserta" binding:"required"`
	}

	GetAntrianRequest struct {
		Nomorkartu     string `json:"nomorkartu"`
		Nik            string `json:"nik" binding:"required"`
		Nohp           string `json:"nohp" binding:"required"`
		Kodepoli       string `json:"kodepoli" binding:"required"`
		Norm           string `json:"norm,omitempty"`
		Tanggalperiksa string `json:"tanggalperiksa" binding:"required"`
		Kodedokter     int    `json:"kodedokter" binding:"required"`
		Jampraktek     string `json:"jampraktek" binding:"required"`
		Jeniskunjungan int    `json:"jeniskunjungan" binding:"required"`
		Nomorreferensi string `json:"nomorreferensi" binding:"required"`
	}

	GetMobileJknByKodebookingDTO struct {
		EstimasiPerPasien int       `json:"estimasi_per_pasien"`
		NoAntrian         string    `json:"no_antrian"`
		TglPeriksa        time.Time `json:"tgl_periksa"`
		Spesialisasi      string    `json:"spesialisasi"`
		Tujuan            string    `json:"tujuan"`
		Namadokter        string    `json:"namadokter"`
		MapingAntrol      string    `json:"maping_antrol"`
	}

	StatusAntreanDTO struct {
		Namapoli        string `json:"namapoli"`
		Namadokter      string `json:"namadokter"`
		Totalantrean    int    `json:"totalantrean"`
		Sisaantrean     int    `json:"sisaantrean"`
		Antreanpanggil  string `json:"antreanpanggil"`
		Sisakuotajkn    int    `json:"sisakuotajkn"`
		Kuotajkn        int    `json:"kuotajkn"`
		Sisakuotanonjkn int    `json:"sisakuotanonjkn"`
		Kuotanonjkn     int    `json:"kuotanonjkn"`
		Keterangan      string `json:"keterangan"`
	}

	SisaANtreanDTO struct {
		Nomorantrean   string `json:"nomorantrean"`
		Namapoli       string `json:"namapoli"`
		Namadokter     string `json:"namadokter"`
		SisaAntrean    int    `json:"sisaantrean"`
		Antreanpanggil string `json:"antreanpanggil"`
		Waktutunggu    int    `json:"waktutunggu"`
		Keterangan     string `json:"keterangan"`
	}

	JadwalOperasiDTO struct {
		Kodebooking    string `json:"kodebooking"`
		Tanggaloperasi string `json:"tanggaloperasi"`
		Jenistindakan  string `json:"jenistindakan"`
		Kodepoli       string `json:"kodepoli"`
		Namapoli       string `json:"namapoli"`
		Terlaksana     int    `json:"terlaksana"`
		Nopeserta      string `json:"nopeserta,omitempty"`
		Lastupdate     int64  `json:"lastupdate,omitempty"`
	}

	InsertPasienDTO struct {
		Nomorantrean     string `json:"nomorantrean"`
		Angkaantrean     int    `json:"angkaantrean"`
		Kodebooking      string `json:"kodebooking"`
		Norm             string `json:"norm"`
		Namapoli         string `json:"namapoli"`
		Namadokter       string `json:"namadokter"`
		Estimasidilayani int    `json:"estimasidilayani"`
		Sisakuotajkn     int    `json:"sisakuotajkn"`
		Kuotajkn         int    `json:"kuotajkn"`
		Sisakuotanonjkn  int    `json:"sisakuotanonjkn"`
		Kuotanonjkn      int    `json:"kuotanonjkn"`
		Keterangan       string `json:"keterangan"`
	}
)
