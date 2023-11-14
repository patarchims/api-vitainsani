package handler

import (
	"net/http"
	"vincentcoreapi/helper"
	"vincentcoreapi/modules/mutiara/dto"
	"vincentcoreapi/modules/mutiara/entity"

	"github.com/gin-gonic/gin"
)

type MutiaraHandler struct {
	MutiaraUseCase    entity.MutiaraUseCase
	MutiaraRepository entity.MutiaraRepository
}

func (ah *MutiaraHandler) GetDataGaji(c *gin.Context) {

	var input dto.GetDataKaryawan

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponseFailure("Tidak dapat menemukan", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		return
	}

	data, err := ah.MutiaraUseCase.GetDataKaryawanUsecase(input.ID)
	if err != nil {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		return
	}

	// MAPPER
	response := helper.APIResponse("Ok", http.StatusOK, data)
	c.JSON(http.StatusOK, response)

}

func (ah *MutiaraHandler) Pengajar(c *gin.Context) {

	data, err := ah.MutiaraRepository.GetPengajarRepository()

	if err != nil {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		return
	}

	if len(data) == 0 {
		response := helper.APIResponseFailure("Data kosong", http.StatusAccepted)
		c.JSON(http.StatusAccepted, response)
		return
	}

	// MAPPER
	response := helper.APIResponse("Ok", http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}
