package farmasi

import "time"

type (
	TimeElapsed struct {
		TimeElapsed    time.Time `json:"timeElapsed"`
		NoAntreanAngka int       `json:"noAntreanAngka"`
	}

	AntreanResep struct {
		KodeBookingRef string    `json:"kodeBookingRef"`
		Mrn            string    `json:"norm"`
		No             int       `json:"no"`
		Cdttm          time.Time `json:"cdttm,omitempty"`
		JenisAntrean   string    `json:"jenisAntrean"`
		JenisPasien    string    `json:"jenisPasien"`
		Tanggal        time.Time `json:"tanggal,omitempty"`
		Jam            string    `json:"jam"`
		TimeElapsed    time.Time `json:"timeElapsed"`
		NoAntrean      string    `json:"noAntrean"`
		NoAntreanAngka int       `json:"noAntreanAngka"`
		KodeBooking    string    `json:"kodeBooking"`
		Racikan        string    `json:"racikan"`
		Dilayani       string    `json:"dilayani"`
		UpdDttm        time.Time `json:"updDttm,omitempty"`
		User           string    `json:"user"`
	}

	AntreanOL struct {
		ID     string `json:"norm"`
		Nik    string `json:"nik"`
		NoBook string `json:"kodeBooking"`
		Nama   string `json:"nama"`
	}

	StatusAntrean struct {
		Tanggal        string
		KodeBooking    string
		Sisaantrean    int
		Antreanpanggil int
		Totalantrean   int
		NoAntrean      string
		NoAntreanAngka int
	}
)

func (AntreanResep) TableName() string {
	return "posfar.antrean_resep"
}

func (AntreanOL) TableName() string {
	return "rekam.antrian_ol"
}
