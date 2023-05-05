package handler

import (
	"encoding/json"
	"net/http"
	"vincentcoreapi/helper"
	"vincentcoreapi/modules/farmasi/dto"
	"vincentcoreapi/modules/farmasi/entity"
	"vincentcoreapi/modules/farmasi/mapper"
	"vincentcoreapi/modules/telegram"

	"github.com/gin-gonic/gin"
)

type FarmasiHandler struct {
	FarmasiUseCase    entity.FarmasiUseCase
	FarmasiRepository entity.FarmasiRepository
	IFarmasiMapper    mapper.IFarmasiMapper
}

// SERVICES POST STATUS ANTREAN
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
	farmasi, err := ah.FarmasiUseCase.AmbilAntreanFarmasi(c, *payload)
	if err != nil || farmasi.JenisResep == "" {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		telegram.RunFailureMessage("AMBIL ANTREAN FARMASI", response, c, data)
		return
	}

	response := helper.APIResponse("Ok", http.StatusOK, "Ok", farmasi)
	telegram.RunSuccessMessage("AMBIL ANTREAN FARMASI", response, c, data)
	c.JSON(http.StatusOK, response)

}

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

	status, err := ah.FarmasiUseCase.StatusAntreanFarmasi(c, *payload)
	if err != nil || status.JenisResep == "" {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		telegram.RunFailureMessage("STATUS ANTREAN FARMASI", response, c, data)
		return
	}

	// MAPPER
	response := helper.APIResponse("Ok", http.StatusOK, "Ok", status)
	telegram.RunSuccessMessage("STATUS ANTREAN FARMASI", response, c, data)
	c.JSON(http.StatusOK, response)

}
