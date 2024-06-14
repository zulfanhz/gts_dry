package handler

import (
	"gts-dry/model"
	"gts-dry/service"

	"github.com/gofiber/fiber/v2"
)

type restoTerimaHandler struct {
	restoTerimaService service.RestoTerimaService
}

func NewRestoTerimaHandler(restoTerimaService service.RestoTerimaService) *restoTerimaHandler {
	return &restoTerimaHandler{restoTerimaService}
}

func (h *restoTerimaHandler) GetRestoTerimaAll(c *fiber.Ctx) error {
	data, err := h.restoTerimaService.GetTerimaRestoAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data resto terima gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data resto terima berhasil",
		Data:    data,
	})
}

func (h *restoTerimaHandler) GetTerimaRestoByCodeResto(c *fiber.Ctx) error {
	kode := c.Params("kode")
	if kode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "kode parameter belum dimasukkan",
			Error:   "kode parameter belum dimasukkan",
		})
	}

	data, err := h.restoTerimaService.GetTerimaRestoByCodeResto(kode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data resto terima gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data resto terima berhasil",
		Data:    data,
	})
}

func (h *restoTerimaHandler) AddRestoTerima(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	var req model.RestoTerimaModel
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	data, err := h.restoTerimaService.AddRestoTerima(req, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "update resto terima gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "update resto terima berhasil",
		Data:    data,
	})
}
