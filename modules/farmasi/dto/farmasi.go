package dto

type (
	GetAntreanFarmasiRequest struct {
		Kodebooking string `json:"kodebooking" binding:"required"`
	}

	AmbilAntreanFarmasiResponse struct {
		JenisResep   string `json:"jenisresep"`
		NomorAntrean int    `json:"nomorantrean"`
		Keterangan   string `json:"keterangan"`
	}

	StatusAntreanFarmasiResponse struct {
		JenisResep     string `json:"jenisresep"`
		TotalAntrean   int    `json:"totalantrean"`
		SisaAntrean    int    `json:"sisaantrean"`
		AntreanPanggil int    `json:"antreanpanggil"`
		Keterangan     string `json:"keterangan"`
	}
)
