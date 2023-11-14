package dto

import "time"

type (
	GetSisaAntrianRequest struct {
		Kodebooking string `json:"kodebooking" binding:"required"  bson:"kodebooking"`
	}

	ResPasienBaru struct {
		Norm string `json:"norm" bson:"norm"`
	}

	SisaAntreanResnonse struct {
		Nomorantrean   string `json:"nomorantrean" bson:"nomorantrean"`
		NamaPoli       string `json:"namapoli" bson:"namapoli"`
		NamaDokter     string `json:"namadokter" bson:"namadokter"`
		SisaAntrean    int    `json:"sisaantrean" bson:"sisaantrean"`
		AntreanPanggil string `json:"antreanpanggil" bson:"antreanpanggil"`
		WaktuTunggu    int    `json:"waktutunggu" bson:"waktutunggu"`
		Keterangan     string `json:"keterangan" bson:"keterangan"`
	}

	StatusAntrianRequest struct {
		KodePoli       string `json:"kodepoli" binding:"required" bson:"kodepoli"`
		TanggalPeriksa string `json:"tanggalperiksa" binding:"required" bson:"tanggalperiksa"`
		KodeDokter     int    `json:"kodedokter" binding:"required" bson:"kodedokter"`
		JamPraktek     string `json:"jampraktek" binding:"required" bson:"jampraktek"`
	}

	BatalAntreanRequest struct {
		Kodebooking string `json:"kodebooking" binding:"required" bson:"kodebooking"`
		Keterangan  string `json:"keterangan" binding:"required" bson:"keterangan"`
	}

	CheckInRequest struct {
		Kodebooking string `json:"kodebooking" binding:"required" bson:"kodebooking"`
		Waktu       int64  `json:"waktu" binding:"required" bson:"waktu"`
	}

	RegisterPasienBaruRequest struct {
		Nomorkartu   string `json:"nomorkartu" bson:"nomorkartu"`
		Nik          string `json:"nik" bson:"nik"`
		Nomorkk      string `json:"nomorkk" bson:"nomorkk"`
		Nama         string `json:"nama" bson:"nama"`
		Jeniskelamin string `json:"jeniskelamin" bson:"jeniskelamin"`
		Tanggallahir string `json:"tanggallahir" bson:"tanggallahir"`
		Nohp         string `json:"nohp" bson:"nohp"`
		Alamat       string `json:"alamat" bson:"alamat"`
		Kodeprop     string `json:"kodeprop" bson:"kodeprop"`
		Namaprop     string `json:"namaprop" bson:"namaprop"`
		Kodedati2    string `json:"kodedati2" bson:"kodedati2"`
		Namadati2    string `json:"namadati2" bson:"namadati2"`
		Kodekec      string `json:"kodekec" bson:"kodekec"`
		Namakec      string `json:"namakec" bson:"namakec"`
		Kodekel      string `json:"kodekel" bson:"kodekel"`
		Namakel      string `json:"namakel" bson:"namakel"`
		Rw           string `json:"rw" bson:"rw"`
		Rt           string `json:"rt" bson:"rt"`
	}

	JadwalOperasiRequest struct {
		Tanggalawal  string `json:"tanggalawal" binding:"required" bson:"tanggalawal"`
		Tanggalakhir string `json:"tanggalakhir" binding:"required" bson:"tanggalakhir"`
	}

	JadwalOperasiPasienRequest struct {
		Nopeserta string `json:"nopeserta" binding:"required" bson:"nopeserta"`
	}

	GetAntrianRequest struct {
		Nomorkartu     string `json:"nomorkartu"  bson:"nomorkartu"`
		Nik            string `json:"nik" binding:"required"  bson:"nik"`
		Nohp           string `json:"nohp" binding:"required"  bson:"nohp"`
		Kodepoli       string `json:"kodepoli" binding:"required"  bson:"kodepoli"`
		Norm           string `json:"norm,omitempty"  bson:"norm,omitempty"`
		Tanggalperiksa string `json:"tanggalperiksa" binding:"required"  bson:"tanggalperiksa"`
		Kodedokter     int    `json:"kodedokter" binding:"required"  bson:"kodedokter"`
		Jampraktek     string `json:"jampraktek" binding:"required"  bson:"jampraktek"`
		Jeniskunjungan int    `json:"jeniskunjungan" binding:"required"  bson:"jeniskunjungan"`
		Nomorreferensi string `json:"nomorreferensi" binding:"required"  bson:"nomorreferensi"`
	}

	GetMobileJknByKodebookingDTO struct {
		EstimasiPerPasien int       `json:"estimasi_per_pasien" bson:"estimasi_per_pasien"`
		NoAntrian         string    `json:"no_antrian" bson:"no_antrian"`
		TglPeriksa        time.Time `json:"tgl_periksa" bson:"tgl_periksa"`
		Spesialisasi      string    `json:"spesialisasi" bson:"spesialisasi"`
		Tujuan            string    `json:"tujuan" bson:"tujuan"`
		Namadokter        string    `json:"namadokter" bson:"namadokter"`
		MapingAntrol      string    `json:"maping_antrol" bson:"maping_antrol"`
	}

	StatusAntreanDTO struct {
		Namapoli        string `json:"namapoli" bson:"namapoli"`
		Namadokter      string `json:"namadokter" bson:"namadokter"`
		Totalantrean    int    `json:"totalantrean" bson:"totalantrean"`
		Sisaantrean     int    `json:"sisaantrean" bson:"sisaantrean"`
		Antreanpanggil  string `json:"antreanpanggil" bson:"antreanpanggil"`
		Sisakuotajkn    int    `json:"sisakuotajkn" bson:"sisakuotajkn"`
		Kuotajkn        int    `json:"kuotajkn" bson:"kuotajkn"`
		Sisakuotanonjkn int    `json:"sisakuotanonjkn" bson:"sisakuotanonjkn"`
		Kuotanonjkn     int    `json:"kuotanonjkn" bson:"kuotanonjkn"`
		Keterangan      string `json:"keterangan" bson:"keterangan"`
	}

	SisaANtreanDTO struct {
		Nomorantrean   string `json:"nomorantrean" bson:"nomorantrean"`
		Namapoli       string `json:"namapoli" bson:"namapoli"`
		Namadokter     string `json:"namadokter" bson:"namadokter"`
		SisaAntrean    int    `json:"sisaantrean" bson:"sisaantrean"`
		Antreanpanggil string `json:"antreanpanggil" bson:"antreanpanggil"`
		Waktutunggu    int    `json:"waktutunggu" bson:"waktutunggu"`
		Keterangan     string `json:"keterangan" bson:"keterangan"`
	}

	JadwalOperasiDTO struct {
		Kodebooking    string `json:"kodebooking" bson:"kodebooking"`
		Tanggaloperasi string `json:"tanggaloperasi" bson:"tanggaloperasi"`
		Jenistindakan  string `json:"jenistindakan" bson:"jenistindakan"`
		Kodepoli       string `json:"kodepoli" bson:"kodepoli"`
		Namapoli       string `json:"namapoli" bson:"namapoli"`
		Terlaksana     int    `json:"terlaksana" bson:"terlaksana"`
		Nopeserta      string `json:"nopeserta,omitempty" bson:"nopeserta,omitempty"`
		Lastupdate     int64  `json:"lastupdate,omitempty" bson:"lastupdate,omitempty"`
	}

	InsertPasienDTO struct {
		Nomorantrean     string `json:"nomorantrean" bson:"nomorantrean"`
		Angkaantrean     int    `json:"angkaantrean" bson:"angkaantrean"`
		Kodebooking      string `json:"kodebooking" bson:"kodebooking"`
		Norm             string `json:"norm" bson:"norm"`
		Namapoli         string `json:"namapoli" bson:"namapoli"`
		Namadokter       string `json:"namadokter" bson:"namadokter"`
		Estimasidilayani int    `json:"estimasidilayani" bson:"estimasidilayani"`
		Sisakuotajkn     int    `json:"sisakuotajkn" bson:"sisakuotajkn"`
		Kuotajkn         int    `json:"kuotajkn" bson:"kuotajkn"`
		Sisakuotanonjkn  int    `json:"sisakuotanonjkn" bson:"sisakuotanonjkn"`
		Kuotanonjkn      int    `json:"kuotanonjkn" bson:"kuotanonjkn"`
		Keterangan       string `json:"keterangan" bson:"keterangan"`
	}

	GetSisaAntrianRequestV2 struct {
		Kodebooking string `json:"kodebooking" validate:"required"`
	}

	BatalAntreanRequestV2 struct {
		Kodebooking string `json:"kodebooking" validate:"required"`
		Keterangan  string `json:"keterangan" validate:"required"`
	}

	//

	CheckInRequestV2 struct {
		Kodebooking string `json:"kodebooking" validate:"required"`
		Waktu       int64  `json:"waktu" validate:"required"`
	}

	JadwalOperasiRequestV2 struct {
		Tanggalawal  string `json:"tanggalawal" validate:"required"`
		Tanggalakhir string `json:"tanggalakhir" validate:"required"`
	}

	JadwalOperasiPasienRequestV2 struct {
		Nopeserta string `json:"nopeserta" validate:"required"`
	}

	GetAntrianRequestV2 struct {
		Nomorkartu     string `json:"nomorkartu"`
		Nik            string `json:"nik" validate:"required"`
		Nohp           string `json:"nohp" validate:"required"`
		Kodepoli       string `json:"kodepoli" validate:"required"`
		Norm           string `json:"norm,omitempty"`
		Tanggalperiksa string `json:"tanggalperiksa" validate:"required"`
		Kodedokter     int    `json:"kodedokter" validate:"required"`
		Jampraktek     string `json:"jampraktek" validate:"required"`
		Jeniskunjungan int    `json:"jeniskunjungan" validate:"required"`
		Nomorreferensi string `json:"nomorreferensi" validate:"required"`
	}
)
