package handler

import (
	"net/http"
	"vincentcoreapi/helper"
	"vincentcoreapi/modules/farmasi/dto"
	"vincentcoreapi/modules/farmasi/entity"
	"vincentcoreapi/modules/farmasi/mapper"
	"vincentcoreapi/modules/telegram"

	"github.com/goccy/go-json"

	"github.com/gin-gonic/gin"
)

type FarmasiHandler struct {
	FarmasiUseCase    entity.FarmasiUseCase
	FarmasiRepository entity.FarmasiRepository
	IFarmasiMapper    mapper.IFarmasiMapper
}

// SERVICES POST STATUS ANTREAN
// @Summary			Ambil Antrean Farmasi
// @Description		Ambil Antrean Farmasi
// @Tags			Farmasi
// @Accept			json
// @Produce			json
// @Security BasicAuth
// @Param			jadwal-operasi-pasien	body		dto.GetAntreanFarmasiRequest	true	"Get Antrean Farmasi Request"
// @Success			200			{object}  	dto.AmbilAntreanFarmasiResponse
// @Failure      	201  		{array}  	helper.FailureResponse
// @Router			/ambil-antrean-farmasi	[post]
func (ah *FarmasiHandler) AmbilAntreanFarmasi(c *gin.Context) {

	payload := new(dto.GetAntreanFarmasiRequest)
	err := c.ShouldBindJSON(&payload)
	data, _ := json.Marshal(payload)

	// CEK APAKAH DATA NULL
	if err != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		telegram.RunFailureMessage("AMBIL ANTREAN FARMASI", response, c, data)
		return
	}

	// AMBIL ANTREAN USECASE
	farmasi, err := ah.FarmasiUseCase.AmbilAntreanFarmasiUsecase(*payload)
	if err != nil || farmasi.JenisResep == "" {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		telegram.RunFailureMessage("AMBIL ANTREAN FARMASI", response, c, data)
		return
	}

	response := helper.APIResponse("Ok", http.StatusOK, farmasi)
	telegram.RunSuccessMessage("AMBIL ANTREAN FARMASI", response, c, data)
	c.JSON(http.StatusOK, response)

}

// @Summary			Status Antrean Farmasi
// @Description		Status Antrean Farmasi
// @Tags			Farmasi
// @Accept			json
// @Produce			json
// @Security BasicAuth
// @Param			jadwal-operasi-pasien	body		dto.GetAntreanFarmasiRequest	true	"Get Antrean Farmasi Request"
// @Success			200			{object}  	dto.StatusAntreanFarmasiResponse
// @Failure      	201  		{array}  	helper.FailureResponse
// @Router			/status-antrean-farmasi	[post]
func (ah *FarmasiHandler) StatusAntreanFarmasi(c *gin.Context) {
	payload := new(dto.GetAntreanFarmasiRequest)
	err := c.ShouldBindJSON(&payload)
	data, _ := json.Marshal(payload)

	// CEK APAKAH DATA NULL
	if err != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		telegram.RunFailureMessage("STATUS ANTREAN FARMASI", response, c, data)
		return
	}

	status, err := ah.FarmasiUseCase.StatusAntreanFarmasiUsecase(*payload)
	if err != nil || status.JenisResep == "" {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		telegram.RunFailureMessage("STATUS ANTREAN FARMASI", response, c, data)
		return
	}

	// MAPPER
	response := helper.APIResponse("Ok", http.StatusOK, status)
	telegram.RunSuccessMessage("STATUS ANTREAN FARMASI", response, c, data)
	c.JSON(http.StatusOK, response)

}
