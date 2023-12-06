package handler

import (
	"fmt"
	"net/http"
	"time"
	"vincentcoreapi/helper"
	"vincentcoreapi/modules/antrian/dto"

	"github.com/gofiber/fiber/v2"
)

func (ah *AntrianHandler) GetStatusAntrianFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.StatusAntrianRequestV2)
	errs := c.BodyParser(&payload)

	// data, _ := json.Marshal(payload)

	if errs != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		// telegram.RunFailureMessageFiber("GET STATUS ANTREAN", response, c, data)
		ah.Logging.Info("Data tidak boleh ada yang null!")
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	validasi := ah.AntrianUseCase.ValidasiDateUsecase(payload.TanggalPeriksa)

	if !validasi {
		response := helper.APIResponseFailure("Format Tanggal Tidak Sesuai, format yang benar adalah yyyy-mm-dd", http.StatusCreated)
		// telegram.RunFailureMessageFiber("GET STATUS ANTREAN", response, c, data)
		ah.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	// CEK BACKDATE
	now := time.Now().Format("2006-01-02")
	date, _ := time.Parse("2006-01-02", now)
	tglPeriksa, _ := time.Parse("2006-01-02", payload.TanggalPeriksa)

	if date.Unix() > tglPeriksa.Unix() {
		response := helper.APIResponseFailure("Tanggal periksa tidak berlaku", http.StatusCreated)
		ah.Logging.Info(response)
		// telegram.RunFailureMessageFiber("GET STATUS ANTREAN", response, c, data)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	detailPoli, err := ah.AntrianRepository.CariPoliRepository(payload.KodePoli)

	if err != nil || detailPoli.Kodepoli == "" {
		response := helper.APIResponseFailure("Poli tidak ditemukan", http.StatusCreated)
		ah.Logging.Info(response)
		// telegram.RunFailureMessageFiber("GET STATUS ANTREAN", response, c, data)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	m, err := ah.AntrianUseCase.GetStatusAntreanUsecaseV2(payload, detailPoli)

	if err != nil {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		ah.Logging.Info(response)
		// telegram.RunFailureMessageFiber("GET STATUS ANTREAN", response, c, data)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	response := helper.APIResponse("Ok", http.StatusOK, m)
	ah.Logging.Info(response)
	// telegram.RunSuccessMessageFiber("GET STATUS ANTREAN", response, c, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (ah *AntrianHandler) GetSisaAntrianFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.GetSisaAntrianRequestV2)
	errs := c.BodyParser(&payload)

	// data, _ := json.Marshal(payload)

	if errs != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		ah.Logging.Info(response)
		// telegram.RunFailureMessageFiber("POST SISA ANTREAN", response, c, data)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	datas, errs := ah.AntrianRepository.GetSisaAntreanRepositoryV2(*payload)

	if errs != nil || datas.Nomorantrean == "" {
		response := helper.APIResponseFailure(errs.Error(), http.StatusCreated)
		ah.Logging.Info(response)
		// telegram.RunFailureMessageFiber("POST SISA ANTREAN", response, c, data)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	response := helper.APIResponse("Ok", http.StatusOK, datas)
	ah.Logging.Info(response)
	// telegram.RunSuccessMessageFiber("POST SISA ANTREAN", response, c, data)
	return c.Status(fiber.StatusOK).JSON(response)

}

func (ah *AntrianHandler) BatalAntreanFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.BatalAntreanRequestV2)
	errs := c.BodyParser(&payload)

	// data, _ := json.Marshal(payload)

	if errs != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		ah.Logging.Info(response)
		// telegram.RunFailureMessageFiber("POST BATAL ANTREAN", response, c, data)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	isSuccessBatal, err := ah.AntrianUseCase.BatalAntreanUsecaseV2(*payload)

	if err != nil || !isSuccessBatal {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		// telegram.RunFailureMessageFiber("POST BATAL ANTREAN", response, c, data)
		ah.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	response := helper.APIResponseFailure("Ok", http.StatusOK)
	ah.Logging.Info(response)
	// telegram.RunFailureMessageFiber("POST BATAL ANTREAN", response, c, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (ah *AntrianHandler) CheckInFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.CheckInRequestV2)
	errs := c.BodyParser(&payload)

	// data, _ := json.Marshal(payload)

	if errs != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		// telegram.RunFailureMessageFiber("POST CHECK IN", response, c, data)
		ah.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	isSuccess := ah.AntrianRepository.CheckInRepository(payload.Kodebooking, payload.Waktu)

	if !isSuccess {
		response := helper.APIResponseFailure("Gagal update", http.StatusCreated)
		// telegram.RunFailureMessageFiber("POST CHECK IN", response, c, data)
		ah.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	// REPONSE HANYA
	// META
	response := helper.APIResponseFailure("Ok", http.StatusOK)
	ah.Logging.Info(response)
	// telegram.RunFailureMessageFiber("POST CHECK IN", response, c, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (ah *AntrianHandler) RegisterPasienBaruFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.RegisterPasienBaruRequest)
	errs := c.BodyParser(&payload)

	// data, _ := json.Marshal(payload)

	if errs != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		// telegram.RunFailureMessageFiber("POST PASIEN BARU", response, c, data)
		ah.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	// NOTE: VALIDASI PASIEN BARU
	errs = validationPayloadPasienBaru(*payload)
	if errs != nil {
		response := helper.APIResponseFailure(errs.Error(), http.StatusCreated)
		// telegram.RunFailureMessageFiber("POST PASIEN BARU", response, c, data)
		ah.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	// REGISTRASI PASIEN BARU
	result, err := ah.AntrianUseCase.RegisterPasienBaruUsecase(*payload)

	if err != nil || result.Norm == "" {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		ah.Logging.Info(response)
		// telegram.RunFailureMessageFiber("POST PASIEN BARU", response, c, data)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	response := helper.APIResponse("Harap datang ke admisi untuk melengkapi data Rekam Medis", http.StatusOK, result)
	ah.Logging.Info(response)
	// telegram.RunSuccessMessageFiber("POST PASIEN BARU", response, c, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (ah *AntrianHandler) GetJadwalOperasiFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.JadwalOperasiRequestV2)
	errs := c.BodyParser(&payload)

	// data, _ := json.Marshal(payload)

	if errs != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada  yang null!", http.StatusCreated)
		ah.Logging.Info(response)
		// telegram.RunFailureMessageFiber("POST GET JADWAL OPERASI", response, c, data)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	// CEK FROMAT TANGGAL
	validasi := ah.AntrianUseCase.ValidasiDateUsecase(payload.Tanggalakhir)
	validasi1 := ah.AntrianUseCase.ValidasiDateUsecase(payload.Tanggalawal)
	if !validasi || !validasi1 {
		response := helper.APIResponseFailure("Format Tanggal (YYYY-MM-DD)", http.StatusCreated)
		ah.Logging.Info(response)
		// telegram.RunFailureMessageFiber("POST GET JADWAL OPERASI", response, c, data)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	// CEK "Tanggal Akhir tidak Boleh lebih Kecil Dari tanggal awal"
	date1, _ := time.Parse("2006-01-02", payload.Tanggalawal)
	date2, _ := time.Parse("2006-01-02", payload.Tanggalakhir)

	// VALIDASI TANGGAL AKHIR TIDAK BOLEH LEBIH KECIL DARI TANGAL AWAL
	if date2.Unix() < date1.Unix() {
		response := helper.APIResponseFailure("Tangal Akhir Tidak Boleh Lebih Kecil dari Tanggal Awal", http.StatusCreated)
		ah.Logging.Info(response)
		// telegram.RunFailureMessageFiber("POST GET JADWAL OPERASI", response, c, data)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	m := map[string]any{}
	jadwalOperasi, err := ah.AntrianRepository.GetJadwalOperasiRepository(payload.Tanggalawal, payload.Tanggalakhir)
	if err != nil || len(jadwalOperasi) == 0 {
		message := fmt.Sprintf("Tidak ada jadwal operasi pada tanggal %s sampai %s", payload.Tanggalawal, payload.Tanggalakhir)
		response := helper.APIResponseFailure(message, http.StatusCreated)
		ah.Logging.Info(response)
		// telegram.RunFailureMessageFiber("POST GET JADWAL OPERASI", response, c, data)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	jadwalOperasiMapper := ah.IAntrianMapper.ToJadwalOperasiDTOMapper(jadwalOperasi, false)
	m["list"] = jadwalOperasiMapper

	response := helper.APIResponse("Ok", http.StatusOK, m)
	ah.Logging.Info(response)
	// telegram.RunSuccessMessageFiber("POST GET JADWAL OPERASI", response, c, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (ah *AntrianHandler) GetKodeBookingOperasiFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.JadwalOperasiPasienRequestV2)
	errs := c.BodyParser(&payload)

	// data, _ := json.Marshal(payload)

	if errs != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		ah.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
		// telegram.RunFailureMessageFiber("POST JADWAL OPERASI", response, c, data)
	}

	jadwalOperasi, err := ah.AntrianUseCase.GetKodeBookingOperasiByNoPesertaUsecaseV2(*payload)

	if err != nil {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		ah.Logging.Info(response)
		// telegram.RunFailureMessageFiber("POST JADWAL OPERASI", response, c, data)
		return c.Status(fiber.StatusOK).JSON(response)
	}

	response := helper.APIResponse("Ok", http.StatusOK, jadwalOperasi)
	ah.Logging.Info(response)
	// telegram.RunSuccessMessageFiber("POST JADWAL OPERASI", response, c, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (ah *AntrianHandler) AmbilAntreanFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.GetAntrianRequestV2)
	errs := c.BodyParser(&payload)

	// data, _ := json.Marshal(payload)

	if errs != nil {
		ah.Logging.Info("Data tidak boleh ada yang null!")
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		ah.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
		// telegram.RunFailureMessageFiber("POST JADWAL OPERASI", response, c, data)
	}

	detaiProfilPasien, err := ah.AntrianRepository.CheckMedrekRepository(payload.Nik)

	if err != nil || detaiProfilPasien.Id == "" {
		message := fmt.Sprintf("%s belum terdaftar rekam medis, silahkan daftar terlebih dahulu", payload.Nomorkartu)
		response := helper.APIResponseFailure(message, http.StatusAccepted)
		ah.Logging.Info(message)
		// telegram.RunFailureMessageFiber("POST AMBIL ANTREAN", response, c, data)
		return c.Status(fiber.StatusAccepted).JSON(response)
	}

	detailPoli, err := ah.AntrianRepository.CariPoliRepository(payload.Kodepoli)

	if err != nil || detailPoli.Kodepoli == "" {
		message := fmt.Sprintf("%s kode poli tersebut tidak ditemukan", payload.Kodepoli)
		response := helper.APIResponseFailure(message, http.StatusCreated)
		ah.Logging.Info(message)
		// telegram.RunFailureMessageFiber("POST AMBIL ANTREAN", response, c, data)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	result, err := ah.AntrianUseCase.AmbilAntreanUsecaseV2(*payload, detailPoli, detaiProfilPasien)

	if err != nil {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		// telegram.RunFailureMessageFiber("POST AMBIL ANTREAN", response, c, data)
		ah.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	response := helper.APIResponse("Ok", http.StatusOK, result)
	ah.Logging.Info(response)
	// telegram.RunSuccessMessageFiber("POST AMBIL ANTREAN", response, c, data)
	return c.Status(fiber.StatusOK).JSON(response)
}
