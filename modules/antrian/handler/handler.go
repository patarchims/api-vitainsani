package handler

import (
	"fmt"
	"net/http"
	"time"
	"vincentcoreapi/helper"
	"vincentcoreapi/modules/antrian/dto"
	"vincentcoreapi/modules/antrian/entity"
	"vincentcoreapi/modules/antrian/mapper"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AntrianHandler struct {
	AntrianUseCase    entity.AntrianUseCase
	AntrianRepository entity.AntrianRepository
	IAntrianMapper    mapper.IAntrianMapper
	Logging           *logrus.Logger
}

// SERVICES POST STATUS ANTREAN

// @Summary			Status Antrean
// @Description		Get Status Antrean Pasien
// @Tags			Antrean
// @Accept			json
// @Produce			json
// @Security BasicAuth
// @Param			status_antrean		body		dto.StatusAntrianRequest		true	"Get Status Antrean"
// @Success			200			{object}  	helper.Response
// @Failure      	201  		{array}  	helper.FailureResponse
// @Router			/status-antrean 	[post]
func (ah *AntrianHandler) GetStatusAntrian(c *gin.Context) {
	payload := new(dto.StatusAntrianRequest)
	err := c.ShouldBindJSON(&payload)
	// data, _ := json.Marshal(payload)

	// CEK APAKAH DATA NULL
	if err != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("GET STATUS ANTREAN", response, c, data)
		return
	}

	validasi := ah.AntrianUseCase.ValidasiDateUsecase(payload.TanggalPeriksa)
	if !validasi {
		response := helper.APIResponseFailure("Format Tanggal Tidak Sesuai, format yang benar adalah yyyy-mm-dd", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("GET STATUS ANTREAN", response, c, data)
		return
	}

	// CEK BACKDATE
	now := time.Now().Format("2006-01-02")
	date, _ := time.Parse("2006-01-02", now)
	tglPeriksa, _ := time.Parse("2006-01-02", payload.TanggalPeriksa)

	if date.Unix() > tglPeriksa.Unix() {
		response := helper.APIResponseFailure("Tanggal periksa tidak berlaku", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("GET STATUS ANTREAN", response, c, data)
		return
	}

	detailPoli, err := ah.AntrianRepository.CariPoliRepository(payload.KodePoli)
	if err != nil || detailPoli.Kodepoli == "" {
		response := helper.APIResponseFailure("Poli tidak ditemukan", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("GET STATUS ANTREAN", response, c, data)
		return
	}

	m, err := ah.AntrianUseCase.GetStatusAntreanUsecase(payload, detailPoli)
	if err != nil {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("GET STATUS ANTREAN", response, c, data)
		return
	}

	response := helper.APIResponse("Ok", http.StatusOK, m)
	// telegram.RunSuccessMessage("GET STATUS ANTREAN", response, c, data)
	c.JSON(http.StatusOK, response)
}

func (ah *AntrianHandler) ListAntrianToday(c *gin.Context) {
	data, errs := ah.AntrianRepository.ListAntrianTodayRepository()

	if errs != nil {
		response := helper.APIResponseFailure(errs.Error(), http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		return
	}

	if len(data) == 0 {
		response := helper.APIResponseFailure("Data kosong", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		return
	}

	response := helper.APIResponse("Ok", http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}

// @Summary			Status Antrean
// @Description		Get Status Antrean Pasien
// @Tags			Antrean
// @Accept			json
// @Produce			json
// @Security BasicAuth
// @Param			status_antrean		body		dto.StatusAntrianRequest		true	"Get Status Antrean"
// @Success			200			{object}  	helper.Response
// @Failure      	201  		{array}  	helper.FailureResponse
// @Router			/sisa-antrean 	[post]
func (ah *AntrianHandler) GetSisaAntrian(c *gin.Context) {
	payload := new(dto.GetSisaAntrianRequest)
	// data, _ := json.Marshal(payload)

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST SISA ANTREAN", response, c, data)
		return
	}

	// CHEK APAKAH ANTRIAN TERSEBUT SUDAH BATAL ATAU TIDAK
	datas, errs := ah.AntrianRepository.GetSisaAntreanRepository(*payload)

	if errs != nil || datas.Nomorantrean == "" {
		response := helper.APIResponseFailure(errs.Error(), http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST SISA ANTREAN", response, c, data)
		return
	}

	response := helper.APIResponse("Ok", http.StatusOK, datas)
	c.JSON(http.StatusOK, response)
	// telegram.RunSuccessMessage("POST SISA ANTREAN", response, c, data)

}

// SERVICES POST BATAL ANTREAN
// @Summary			Batal Antrean
// @Description		Batal Antrean
// @Tags			Antrean
// @Accept			json
// @Produce			json
// @Security BasicAuth
// @Param			batal_antrean		body		dto.BatalAntreanRequest		true	"Batal Antrean"
// @Success			200			{object}  	helper.Response
// @Failure      	201  		{array}  	helper.FailureResponse
// @Router			/batal-antrean 	[post]
func (ah *AntrianHandler) BatalAntrean(c *gin.Context) {

	payload := new(dto.BatalAntreanRequest)
	err := c.ShouldBindJSON(&payload)
	// data, _ := json.Marshal(payload)

	if err != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST BATAL ANTREAN", response, c, data)
		return
	}

	isSuccessBatal, err := ah.AntrianUseCase.BatalAntreanUsecase(*payload)
	if err != nil || !isSuccessBatal {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST BATAL ANTREAN", response, c, data)
		return
	}

	response := helper.APIResponseFailure("Ok", http.StatusOK)
	c.JSON(http.StatusOK, response)
	// telegram.RunFailureMessage("POST BATAL ANTREAN", response, c, data)
}

// SERVICES CHECK IN
// @Summary			Check In Antrean
// @Description		Check In Antrean
// @Tags			Antrean
// @Accept			json
// @Produce			json
// @Security BasicAuth
// @Param			check-in	body		dto.CheckInRequest	true	"Batal Antrean"
// @Success			200			{object}  	helper.Response
// @Failure      	201  		{array}  	helper.FailureResponse
// @Router			/check-in 	[post]
func (ah *AntrianHandler) CheckIn(c *gin.Context) {

	payload := new(dto.CheckInRequest)
	err := c.ShouldBindJSON(&payload)
	// data, _ := json.Marshal(payload)

	if err != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST CHECK IN", response, c, data)
		return
	}

	isSuccess := ah.AntrianRepository.CheckInRepository(payload.Kodebooking, payload.Waktu)
	if !isSuccess {
		response := helper.APIResponseFailure("Gagal update", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST CHECK IN", response, c, data)
		return
	}

	// REPONSE HANYA META
	response := helper.APIResponseFailure("Ok", http.StatusOK)
	c.JSON(http.StatusOK, response)
	// telegram.RunFailureMessage("POST CHECK IN", response, c, data)
}

// SERVICES REGISTER PASIEN BARU
// @Summary			Check In Antrean
// @Description		Check In Antrean
// @Tags			Antrean
// @Accept			json
// @Produce			json
// @Security BasicAuth
// @Param			check-in	body		dto.RegisterPasienBaruRequest	true	"Pasien Baru"
// @Success			200			{object}  	helper.Response
// @Failure      	201  		{array}  	helper.FailureResponse
// @Router			/pasien-baru 	[post]
func (ah *AntrianHandler) RegisterPasienBaru(c *gin.Context) {
	payload := new(dto.RegisterPasienBaruRequest)
	err := c.ShouldBindJSON(&payload)
	// data, _ := json.Marshal(payload)
	if err != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST PASIEN BARU", response, c, data)
		return
	}

	// NOTE: VALIDASI PASIEN BARU
	err = validationPayloadPasienBaru(*payload)
	if err != nil {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST PASIEN BARU", response, c, data)
		return
	}

	// REGISTRASI PASIEN BARU
	result, err := ah.AntrianUseCase.RegisterPasienBaruUsecase(*payload)

	if err != nil || result.Norm == "" {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST PASIEN BARU", response, c, data)
		return
	}

	response := helper.APIResponse("Harap datang ke admisi untuk melengkapi data Rekam Medis", http.StatusOK, result)
	c.JSON(http.StatusOK, response)
	// telegram.RunSuccessMessage("POST PASIEN BARU", response, c, data)
}

// GetJadwalOperasi
// @Summary			Get Jadwal Operasi Pasien
// @Description		Get Jadwal Operasi Pasien
// @Tags			Antrean
// @Accept			json
// @Produce			json
// @Security BasicAuth
// @Param			jadwal-operasi	body		dto.JadwalOperasiRequest	true	"Jadwal Operasi Pasien"
// @Success			200			{object}  	helper.Response
// @Failure      	201  		{array}  	helper.FailureResponse
// @Router			/get-jadwal-operasi 	[post]
func (ah *AntrianHandler) GetJadwalOperasi(c *gin.Context) {
	payload := new(dto.JadwalOperasiRequest)
	err := c.ShouldBindJSON(&payload)
	// data, _ := json.Marshal(payload)

	if err != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada  yang null!", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST GET JADWAL OPERASI", response, c, data)
		return
	}

	// CEK FROMAT TANGGAL
	validasi := ah.AntrianUseCase.ValidasiDateUsecase(payload.Tanggalakhir)
	validasi1 := ah.AntrianUseCase.ValidasiDateUsecase(payload.Tanggalawal)
	if !validasi || !validasi1 {
		response := helper.APIResponseFailure("Format Tanggal (YYYY-MM-DD)", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST GET JADWAL OPERASI", response, c, data)
		return
	}

	// CEK "Tanggal Akhir tidak Boleh lebih Kecil Dari tanggal awal"
	date1, _ := time.Parse("2006-01-02", payload.Tanggalawal)
	date2, _ := time.Parse("2006-01-02", payload.Tanggalakhir)

	// VALIDASI TANGGAL AKHIR TIDAK BOLEH LEBIH KECIL DARI TANGAL AWAL
	if date2.Unix() < date1.Unix() {
		response := helper.APIResponseFailure("Tangal Akhir Tidak Boleh Lebih Kecil dari Tanggal Awal", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST GET JADWAL OPERASI", response, c, data)
		return
	}

	m := map[string]any{}
	jadwalOperasi, err := ah.AntrianRepository.GetJadwalOperasiRepository(payload.Tanggalawal, payload.Tanggalakhir)
	if err != nil || len(jadwalOperasi) == 0 {
		message := fmt.Sprintf("Tidak ada jadwal operasi pada tanggal %s sampai %s", payload.Tanggalawal, payload.Tanggalakhir)
		response := helper.APIResponseFailure(message, http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST GET JADWAL OPERASI", response, c, data)
		return
	}

	jadwalOperasiMapper := ah.IAntrianMapper.ToJadwalOperasiDTOMapper(jadwalOperasi, false)
	m["list"] = jadwalOperasiMapper

	response := helper.APIResponse("Ok", http.StatusOK, m)
	// telegram.RunSuccessMessage("POST GET JADWAL OPERASI", response, c, data)
	c.JSON(http.StatusOK, response)

}

// GetKodeBookingOperasi
// @Summary			Get Jadwal Operasi Pasien Pasien
// @Description		Get Jadwal Operasi Pasien Pasien
// @Tags			Antrean
// @Accept			json
// @Produce			json
// @Security BasicAuth
// @Param			jadwal-operasi-pasien	body		dto.JadwalOperasiPasienRequest	true	"Get Booking Operasi"
// @Success			200			{object}  	helper.Response
// @Failure      	201  		{array}  	helper.FailureResponse
// @Router			/list-jadwal-operasi	[post]
func (ah *AntrianHandler) GetKodeBookingOperasi(c *gin.Context) {

	payload := new(dto.JadwalOperasiPasienRequest)
	err := c.ShouldBindJSON(&payload)

	// data, _ := json.Marshal(payload)
	if err != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST JADWAL OPERASI", response, c, data)
		return
	}

	jadwalOperasi, err := ah.AntrianUseCase.GetKodeBookingOperasiByNoPesertaUsecase(*payload)

	if err != nil {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST JADWAL OPERASI", response, c, data)
		return
	}

	response := helper.APIResponse("Ok", http.StatusOK, jadwalOperasi)
	c.JSON(http.StatusOK, response)
	// telegram.RunSuccessMessage("POST JADWAL OPERASI", response, c, data)
}

// AMBIL ANTRIAN
// @Summary			Get Antrean Pasien
// @Description		Get Antrean Pasien
// @Tags			Antrean
// @Accept			json
// @Produce			json
// @Security BasicAuth
// @Param			antrian-pasien	body		dto.GetAntrianRequest	true	"Get Antrean Pasien Request"
// @Success			200			{object}  	dto.InsertPasienDTO
// @Failure      	201  		{array}  	helper.FailureResponse
// @Router			/list-jadwal-operasi	[post]
func (ah *AntrianHandler) AmbilAntrean(c *gin.Context) {
	payload := new(dto.GetAntrianRequest)
	err := c.ShouldBindJSON(&payload)
	// data, _ := json.Marshal(payload)

	if err != nil {
		response := helper.APIResponseFailure("data tidak boleh ada yang null!", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST AMBIL ANTREAN", response, c, data)
		return
	}

	detaiProfilPasien, err := ah.AntrianRepository.CheckMedrekRepository(payload.Nik)

	if err != nil || detaiProfilPasien.Id == "" {
		message := fmt.Sprintf("%s belum terdaftar rekam medis, silahkan daftar terlebih dahulu", payload.Nomorkartu)
		response := helper.APIResponseFailure(message, http.StatusAccepted)
		c.JSON(http.StatusAccepted, response)
		// telegram.RunFailureMessage("POST AMBIL ANTREAN", response, c, data)
		return
	}

	detailPoli, err := ah.AntrianRepository.CariPoliRepository(payload.Kodepoli)

	if err != nil || detailPoli.Kodepoli == "" {
		message := fmt.Sprintf("%s kode poli tersebut tidak ditemukan", payload.Kodepoli)
		response := helper.APIResponseFailure(message, http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST AMBIL ANTREAN", response, c, data)
		return
	}

	result, err := ah.AntrianUseCase.AmbilAntreanUsecase(*payload, detailPoli, detaiProfilPasien)

	if err != nil {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		// telegram.RunFailureMessage("POST AMBIL ANTREAN", response, c, data)
		return
	}

	response := helper.APIResponse("Ok", http.StatusOK, result)
	c.JSON(http.StatusOK, response)
	// telegram.RunSuccessMessage("POST AMBIL ANTREAN", response, c, data)
}
