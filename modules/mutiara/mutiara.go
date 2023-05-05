package mutiara

import "time"

type DGaji struct {
	Bagian    string
	Periode   string
	Jumlah    float64
	Terbilang string
	Tglbayar  time.Time
	Nogaji    string
	Bulan     string
	Id        string
	Nama      string
	Tglmasuk  time.Time
	Golongan  string
	Gajipokok float64
	Jabatan   float64
	Makan     float64
	Tansport  float64
	Jamlem    float64
	Bruto     float64
	Potot     float64
}

func (DGaji) TableName() string {
	return "mutiara.dgaji"
}
