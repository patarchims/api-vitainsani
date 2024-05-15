package dto

type (
	PasienResponse struct {
		Nik         string `json:"nik"`
		NoKK        string `json:"noka"`
		Nama        string `json:"nama"`
		NoRm        string `json:"no_rm"`
		Agama       string `json:"agama"`
		TempatLahir string `json:"tempat_lahir"`
		Umur        int    `json:"umur"`
		Suku        string `json:"suku"`
		Kelurahan   string `json:"kelurahan"`
		Kecamatan   string `json:"kecamatan"`
		KodePos     string `json:"kode_pos"`
		Pekerjaan   string `json:"pekerjaan"`
		NamaAyah    string `json:"nama_ayah"`
		NamaIbu     string `json:"nama_ibu"`
		Telp        string `json:"telp"`
	}
)
