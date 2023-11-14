package constant

import "errors"

var (
	SomethingWentWrong            = errors.New("Something Went Wrong")
	DataGagalDiProses             = errors.New("Data gagal diproses")
	EmailNotFound                 = errors.New("Email tidak ditemukan")
	UserHasBeenRegister           = errors.New("User Sudah Terdaftar")
	PhoneNotMatch                 = errors.New("Phone Not Match")
	PasswordNotValid              = errors.New("Password Not Valid")
	PasswordNotMatch              = errors.New("Sandi anda tidak sesuai")
	PasswordOrEmailNotMatch       = errors.New("Email atau Sandi anda tidak sesuai")
	UserNotFound                  = errors.New("User tidak ditemukan")
	FailedValidationTokenNotFound = errors.New("Validasi token tidak di temukan, silahkan verifikasi ulang")
	FailedValidationToken         = errors.New("Maaf Gagal Validasi Token")
	ValidationEmail               = errors.New("Validasi email anda terlebih dahulu")
	EmailHashBeenVerify           = errors.New("Email anda sudah di verifikasi")
	UserRoleInvalid               = errors.New("Maaf role user tidak valid")
	DataAlreadyExists             = errors.New("Data sudah ada")
	DataNotFound                  = errors.New("Data tidak ditemukan")
	StatusNotMatch                = errors.New("Status tidak ditemukan")
	ProductIDNotFound             = errors.New("Produk ID wajib di isi")
	NotChangeThemes               = errors.New("Maaf tema tidak bisa di ganti")
	DocumentHasBeenPaid           = errors.New("Dokumen sudah di bayar")
	CannotPublishDocument         = errors.New("Dokumen tidak bisa di publish, mohon lakukan pembayaran terlebih dahulu")
	CannotIsExampleDocument       = errors.New("Dokumen tidak bisa di jadikan preview, mohon publish dokumen terlebih dahulu")
)
