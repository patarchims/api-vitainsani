package handler

import (
	"net/http"
	"vincentcoreapi/helper"
	"vincentcoreapi/modules/farmasi/dto"

	"github.com/gofiber/fiber/v2"
)

func (ah *FarmasiHandler) AmbilAntreanFarmasiFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.GetAntreanFarmasiRequestV2)

	errs := c.BodyParser(&payload)

	if errs != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	// AMBIL ANTREAN USECASE
	farmasi, err11 := ah.FarmasiUseCase.AmbilAntreanFarmasiUsecaseV2(*payload)

	if err11 != nil || farmasi.JenisResep == "" {
		response := helper.APIResponseFailure(err11.Error(), http.StatusCreated)
		ah.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	response := helper.APIResponse("Ok", http.StatusOK, farmasi)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (ah *FarmasiHandler) StatusAntreanFarmasiFiberHandler(c *fiber.Ctx) error {
	payload := new(dto.GetAntreanFarmasiRequestV2)

	errs := c.BodyParser(&payload)

	if errs != nil {
		response := helper.APIResponseFailure("Data tidak boleh ada yang null!", http.StatusCreated)
		ah.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	status, err := ah.FarmasiUseCase.StatusAntreanFarmasiUsecaseV2(*payload)

	if err != nil || status.JenisResep == "" {
		response := helper.APIResponseFailure(err.Error(), http.StatusCreated)
		ah.Logging.Info(response)
		return c.Status(fiber.StatusCreated).JSON(response)
	}

	// MAPPER
	response := helper.APIResponse("Ok", http.StatusOK, status)
	ah.Logging.Info(response)
	return c.Status(fiber.StatusOK).JSON(response)
}
