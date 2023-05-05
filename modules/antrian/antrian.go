package antrian

import (
	"time"

	"gorm.io/gorm"
)

type (
	AntrianOl struct {
		BookingOnsite      string `json:"booking_onsite"`
		BookingOnsiteDttm  string `json:"booking_onsite_dttm"`
		BookingOnsiteUser  string `json:"booking_onsite_user"`
		JenisAntreanPasien string `json:"jenis_antrean_pasien"`
		// omitempty -> Bidang harus dihilangkan
		TimeElapsed     string `json:"time_elapsed,omitempty"`
		Notif           string `json:"notif"`
		Batal           string `json:"batal"`
		RegType         string `json:"reg_type"`
		KetUpdate       string `json:"ket_update"`
		Kunci           string `json:"kunci"`
		Status          string `json:"status"`
		Keterangan      string `json:"keterangan"`
		Proses          string `json:"proses"`
		User            string `json:"user"`
		TglProses       string `json:"tgl_proses"`
		Pc              string `json:"pc"`
		Id              string `json:"id"`
		Noreg           string `json:"noreg"`
		Nik             string `json:"nik"`
		Noka            string `json:"noka"`
		Jeniskunjungan  string `json:"jeniskunjungan"`
		NoRujukan       string `json:"no_rujukan"`
		NoInhealth      string `json:"no_inhealth"`
		NoHp            string `json:"no_hp"`
		Nama            string `json:"nama"`
		Gender          string `json:"gender"`
		Dob             string `json:"dob,omitempty"`
		Alamat          string `json:"alamat"`
		KodeTujuan      string `json:"kode_tujuan"`
		Tujuan          string `json:"tujuan"`
		TglPeriksa      string `json:"tgl_periksa"`
		KdDokter        string `json:"kd_dokter"`
		Dokter          string `json:"dokter"`
		UnikId          string `json:"unik_id"`
		NoAntrian       string `json:"no_antrian"`
		KodeDebitur     string `json:"kode_debitur"`
		Debitur         string `json:"debitur"`
		BookDate        string `json:"book_date"`
		NoBook          string `json:"no_book"`
		JknNomorkk      string `json:"jkn_nomorkk"`
		JknTanggallahir string `json:"jkn_tanggallahir"`
		JknKodeprop     string `json:"jkn_kodeprop"`
		JknNamaprop     string `json:"jkn_namaprop"`
		JknKodedati2    string `json:"jkn_kodedati2"`
		JknNamadati2    string `json:"jkn_namadati2"`
		JknKodekec      string `json:"jkn_kodekec"`
		JknNamakec      string `json:"jkn_namakec"`
		JknKodekel      string `json:"jkn_kodekel"`
		JknNamakel      string `json:"jkn_namakel"`
		JknRw           string `json:"jkn_rw"`
		JknRt           string `json:"jkn_rt"`
		CheckedIn       string `json:"checked_in"`
	}

	AntrianOl2 struct {
		BookingOnsite      string `json:"booking_onsite"`
		BookingOnsiteDttm  string `json:"booking_onsite_dttm"`
		BookingOnsiteUser  string `json:"booking_onsite_user"`
		JenisAntreanPasien string `json:"jenis_antrean_pasien"`
		TimeElapsed        string `json:"time_elapsed"`
		Notif              string `json:"notif"`
		Batal              string `json:"batal"`
		RegType            string `json:"reg_type"`
		KetUpdate          string `json:"ket_update"`
		Kunci              string `json:"kunci"`
		Status             string `json:"status"`
		Keterangan         string `json:"keterangan"`
		Proses             string `json:"proses"`
		User               string `json:"user"`
		TglProses          string `json:"tgl_proses"`
		Pc                 string `json:"pc"`
		Id                 string `json:"id"`
		Noreg              string `json:"noreg"`
		Nik                string `json:"nik"`
		Noka               string `json:"noka"`
		Jeniskunjungan     string `json:"jeniskunjungan"`
		NoRujukan          string `json:"no_rujukan"`
		NoInhealth         string `json:"no_inhealth"`
		NoHp               string `json:"no_hp"`
		Nama               string `json:"nama"`
		Gender             string `json:"gender"`
		Dob                string `json:"dob"`
		Alamat             string `json:"alamat"`
		KodeTujuan         string `json:"kode_tujuan"`
		Tujuan             string `json:"tujuan"`
		TglPeriksa         string `json:"tgl_periksa"`
		KdDokter           string `json:"kd_dokter"`
		Dokter             string `json:"dokter"`
		UnikId             string `json:"unik_id"`
		NoAntrian          string `json:"no_antrian"`
		KodeDebitur        string `json:"kode_debitur"`
		Debitur            string `json:"debitur"`
		BookDate           string `json:"book_date"`
		NoBook             string `json:"no_book"`
		JknNomorkk         string `json:"jkn_nomorkk"`
		JknTanggallahir    string `json:"jkn_tanggallahir"`
		JknKodeprop        string `json:"jkn_kodeprop"`
		JknNamaprop        string `json:"jkn_namaprop"`
		JknKodedati2       string `json:"jkn_kodedati2"`
		JknNamadati2       string `json:"jkn_namadati2"`
		JknKodekec         string `json:"jkn_kodekec"`
		JknNamakec         string `json:"jkn_namakec"`
		JknKodekel         string `json:"jkn_kodekel"`
		JknNamakel         string `json:"jkn_namakel"`
		JknRw              string `json:"jkn_rw"`
		JknRt              string `json:"jkn_rt"`
		CheckedIn          string `json:"checked_in"`
	}

	JadopOl struct {
		User           string    `json:"user,omitempty"`
		Pc             string    `json:"pc,omitempty"`
		Id             string    `json:"id,omitempty"`
		Noreg          string    `json:"noreg,omitempty"`
		Nama           string    `json:"nama,omitempty"`
		Gender         string    `json:"gender,omitempty"`
		Debitur        string    `json:"debitur,omitempty"`
		InsertDttm     time.Time `json:"insert_dttm,omitempty"`
		UpdDttm        time.Time `json:"upd_dttm,omitempty"`
		KodeDokter     string    `json:"kode_dokter,omitempty"`
		Dokter         string    `json:"dokter,omitempty"`
		Keter          string    `json:"keter,omitempty"`
		NoBook         string    `json:"no_book,omitempty"`
		TglOperasi     time.Time `json:"tgl_operasi,omitempty"`
		KodeTindakan   string    `json:"kode_tindakan,omitempty"`
		JenisTindakan  string    `json:"jenis_tindakan,omitempty"`
		KdTujuan       string    `json:"kd_tujuan,omitempty"`
		Tujuan         string    `json:"tujuan,omitempty"`
		Status         string    `json:"status,omitempty"`
		StatusChangeBy string    `json:"status_change_by,omitempty"`
		Noka           string    `json:"noka,omitempty"`
	}

	LiburOl struct {
		InsertDttm   time.Time `json:"insert_dttm"`
		User         string    `json:"user"`
		Keterangan   string    `json:"keterangan"`
		Deskripsi    string    `json:"deskripsi"`
		Tanggal      time.Time `json:"tanggal"`
		CatatanLibur string    `json:"catatan_libur"`
	}

	KtaripDokter struct {
		Namadokter        string `json:"namadokter"`
		Alamat            string `json:"alamat"`
		EstimasiPerPasien int    `json:"estimasi_per_pasien"`
		QuotaPasienMon    int    `json:"quota_pasien_mon"`
		QuotaPasienTue    int    `json:"quota_pasien_tue"`
		QuotaPasienWed    int    `json:"quota_pasien_wed"`
		QuotaPasienThu    int    `json:"quota_pasien_thu"`
		QuotaPasienFri    int    `json:"quota_pasien_fri"`
		QuotaPasienSat    int    `json:"quota_pasien_sat"`
		Iddokter          string `json:"iddokter"`
		MapingAntrol      string `json:"maping_antrol"`
		Mon               int    `json:"mon"`
		Tue               int    `json:"tue"`
		Wed               int    `json:"wed"`
		Thu               int    `json:"thu"`
		Fri               int    `json:"fri"`
		Sat               int    `json:"sat"`
		Sun               int    `json:"sun"`
		Senin             string `json:"senin"`
		Selasa            string `json:"selasa"`
		Rabu              string `json:"rabu"`
		Kamis             string `json:"kamis"`
		Jumat             string `json:"jumat"`
		Sabtu             string `json:"sabtu"`
		Jampraktek        string `gorm:"-"`
	}

	AntreanOlPol struct {
		Nomorantrean   string `json:"nomorantrean"`
		Namapoli       string `json:"namapoli"`
		Namadokter     string `json:"namadokter"`
		SisaAntrean    int    `json:"sisaantrean"`
		Antreanpanggil string `json:"antreanpanggil"`
		Waktutunggu    int    `json:"waktutunggu"`
		Keterangan     string `json:"keterangan"`
	}

	LastCalled struct {
		Nomorantrean string `json:"nomorantrean"`
		Angkaantrean string `json:"angkaantrean"`
	}

	Dprofilpasien struct {
		gorm.Model
		Id           string `json:"id"`
		Firstname    string `json:"firstname"`
		Jeniskelamin string `json:"jeniskelamin"`
		Tgllahir     string `json:"tgllahir"`
		Alamat       string `json:"alamat"`
	}

	Kpoli struct {
		LibTarip          string  `json:"lib_tarip"`
		Dokter            string  `json:"dokter"`
		Bpjs              string  `json:"bpjs"`
		Kodepoli          string  `json:"kodepoli"`
		Namapoli          string  `json:"namapoli"`
		Tabel             string  `json:"tabel"`
		Hargapoli         float64 `json:"hargapoli"`
		Antrian           string  `json:"antrian"`
		Hargaadministrasi string  `json:"hargaadministrasi"`
	}
)

func (AntrianOl) TableName() string {
	return "antrian_ol"
}
func (AntreanOlPol) TableName() string {
	return "antrian_ol_pool"
}

func (AntrianOl2) TableName() string {
	return "antrian_ol2"
}

func (JadopOl) TableName() string {
	return "jadop_ol"
}

func (LiburOl) TableName() string {
	return "libur_ol"
}

func (KtaripDokter) TableName() string {
	return "his.ktaripdokter"
}

func (Dprofilpasien) TableName() string {
	return "his.dprofilpasien"
}

func (Kpoli) TableName() string {
	return "his.kpoli"
}
