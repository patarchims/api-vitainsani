package dto

type (
	GetDataKaryawan struct {
		ID string `uri:"id" binding:"required"`
	}

	ResDataGaji struct {
		Jumlah    string `json:"jumlah"`
		Tglbayar  string `json:"tglBayar"`
		Bulan     string `json:"bulan"`
		Gajipokok string `json:"gajiPokok"`
		Makan     string `json:"makan"`
		Tansport  string `json:"transport"`
		Jamlem    string `json:"jamLembur"`
		Bruto     string `json:"bruto"`
		Potot     string `json:"potongan"`
	}
)
