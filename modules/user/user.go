package user

type (
	ApiUser struct {
		ID       string `json:"id"`
		UserName string `json:"username"`
		Password string `json:"password"`
		Ket      string `json:"ket"`
	}

	DProfilePasien struct {
		NoregTerakhir        string
		UserKeluarTerakhir   string
		KeluarTerakhirBagian string
		Nik                  string
		Nokapst              string
		KembaliTgl           string
		KembaliOleh          string
		KembaliTerima        string
		KeluarKe             string
		Tujuan               string
		TglkeluarRm          string
		Oleh                 string
		Strm                 string
		KTerakhir            string
		Gelar                string
		Id                   string
		Firstname            string
		Lastname             string
		Agama                string
		Jeniskelamin         string
		Tempatlahir          string
		Tgllahir             string
		Umurth               int
		Umurbln              int
		Suku                 string
		Alamat               string
		Alamat2              string
		Rtrw                 string
		Kelurahan            string
		Kecamatan            string
		Kotamadya            string
		Kabupaten            string
		Propinsi             string
		Negara               string
		Kodepos              string
		Telp                 string
		Hp                   string
		Namaayah             string
		Namaibu              string
		Pendidikan           string
		Pekerjaan            string
		Status               string
		Peksuami             string
		Umsuami              string
		Kunjungan            string
		CpName               string
		CpNumber             string
		CpRelasi             string
	}

	// AMBIL DATA DARI VICORE_USR
	VicoreUserModel struct {
		ReqStatus     string
		NotifStatus   string
		DelStatus     string
		CreatePc      string
		CreateUser    string
		CreateDttm    string
		UpdDttm       string
		Email         string
		Nik           string
		NoHp          string
		KodeModul     string
		VerifDttm     string
		VerifUser     string
		VerifPc       string
		UserStatus    string
		ActiveCode    string
		NotifActiated bool
		UserId        string
		Password      string
	}

	// TABLE KLINIK.USERS
	KlinikUsers struct {
		App               string
		Terminated        string
		Id                string
		User              string
		MappingKpelayanan string
		Password          string
		View              string
		Print             string
		Edit              string
		Del               string
	}
)

func (ApiUser) TableName() string {
	return "rekam.api_user"
}

// === TABLE KLINIK.USERS
func (KlinikUsers) TableName() string {
	return "klinik.users"
}

func (DProfilePasien) TableName() string {
	return "his.dprofilpasien"
}
