package handler

import (
	"fmt"
	"net/http"
	"strings"
	"vincentcoreapi/helper"
	"vincentcoreapi/modules/transfer/dto"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type TranferHandler struct{}

// @Summary			Transfer FIle
// @Description		Transfer FIle
// @Tags			File
// @Accept			multipart/form-data
// @Produce			json
// @Security 		BasicAuth
// @Param			file_transfer		formData		file		true	"Transfer File"
// @Success			200					{object}  		helper.Response
// @Failure      	201  				{array}  		helper.FailureResponse
// @Router			/upload-file 		[post]
func (ah *TranferHandler) UploadFile(c *gin.Context) {
	payload := new(dto.DTOTransferFile)
	err := c.ShouldBind(&payload)

	if err != nil {
		log.Error().Msg(err.Error())
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Data gagal diproses", http.StatusAccepted, errorMessage)
		c.JSON(http.StatusAccepted, response)
		return
	}

	file := payload.File
	tipe := strings.Split(file.Filename, ".")

	path := fmt.Sprintf("files/%s.%s", tipe[0], tipe[1])

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		log.Info().Msg(err.Error())
		response := helper.APIResponseFailure("File gagal diupload", http.StatusCreated)
		c.JSON(http.StatusCreated, response)
		return
	}

	data := dto.DTOTransferFileResponse{FileName: path}

	response := helper.APIResponse("Data berhasil disimpan", http.StatusOK, data)
	log.Info().Msg("Data berhasil disimpan")
	c.JSON(http.StatusOK, response)

}
