package dto

type (
	StatusAntrianRequestV2 struct {
		KodePoli       string `json:"kodepoli" validate:"required"`
		TanggalPeriksa string `json:"tanggalperiksa" validate:"required"`
		KodeDokter     int    `json:"kodedokter" validate:"required"`
		JamPraktek     string `json:"jampraktek" validate:"required"`
	}
)
