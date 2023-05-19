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
		response := helper.APIResponse("Data gagal diproses", http.StatusAccepted, "error", errorMessage)
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

	response := helper.APIResponse("Data berhasil disimpan", http.StatusOK, "OK", data)
	log.Info().Msg("Data berhasil disimpan")
	c.JSON(http.StatusOK, response)

}

// @Summary			File Directory
// @Description		File Directory
// @Tags			File
// @Accept			json
// @Produce			json
// @Security 		BasicAuth
// @Param			file_transfer		formData		file		true	"Transfer File"
// @Success			200					{object}  		helper.Response
// @Failure      	201  				{array}  		helper.FailureResponse
// @Router			/file-directories		[get]
// func (ah *TranferHandler) FileDirectory(c *gin.Context) {
// 	root := "./files"

// 	// Slice to store directory paths
// 	var directories []string

// 	// Define the WalkFunc callback function
// 	walkFunc := func(path string, info os.DirEntry, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		if info.IsDir() {
// 			directories = append(directories, path)
// 		}
// 		return nil
// 	}

// 	// Walk the directory tree and collect directories
// 	err := filepath.Walk(root, walkFunc)
// 	if err != nil {
// 		log.Info().Msg(err.Error())
// 		response := helper.APIResponseFailure("Error", http.StatusCreated)
// 		c.JSON(http.StatusCreated, response)
// 		return
// 	}

// 	response := helper.APIResponse("Data berhasil disimpan", http.StatusOK, "OK", gin.H{"directories": directories})
// 	log.Info().Msg("Data berhasil disimpan")
// 	c.JSON(http.StatusOK, response)
// }
