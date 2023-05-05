package mutiara

import "time"

type (
	DKaryawan struct {
		Kelas          string    `json:"kelas"`
		Ktp            string    `json:"ktp"`
		AlamatSekarang string    `json:"alamatSekarang"`
		Bagian         string    `json:"bagian"`
		Tglmasuk       time.Time `json:"tglMasuk"`
		Id             string    `gorm:"primaryKey:Id json:id"`
		Nama           string    `json:"nama"`
		Jeniskelamin   string    `json:"jenisKelamin"`
		Tgllahir       time.Time `json:"tglLahir"`
		Usia           int32     `json:"usia"`
		Agama          string    `json:"agama"`
		Alamat         string    `json:"alamat"`
		Dinasmalam     string    `json:"dinasMalam"`
		Pendidikan     string    `json:"pendidikan"`
		Status         string    `json:"status"`
		Anak           string    `json:"anak"`
		Posisi         string    `json:"posisi"`
		Kota           string    `json:"kota"`
		Bank           string    `json:"bank"`
		Account        string    `json:"account"`
		Atasnama       string    `json:"atasNama"`
		Cabang         string    `json:"cabang"`
	}

	DGaji struct {
		Id        string    `json:"id"`
		Periode   string    `json:"periode"`
		Jumlah    float64   `json:"jumlah"`
		Terbilang string    `json:"terbilang"`
		Tglbayar  time.Time `json:"tglBayar"`
		Nogaji    string    `json:"noGaji"`
		Bulan     string    `json:"bulan"`
		Golongan  string    `json:"golongan"`
		Gajipokok float64   `json:"gajiPokok"`
		Jabatan   float64   `json:"jabatan"`
		Makan     float64   `json:"makan"`
		Tansport  float64   `json:"transport"`
		Jamlem    float64   `json:"jamLembur"`
		Bruto     float64   `json:"bruto"`
		Potot     float64   `json:"potongan"`
	}
)

func (DGaji) TableName() string {
	return "mutiara.dgaji"
}

func (DKaryawan) TableName() string {
	return "mutiara.pengajar"
}
