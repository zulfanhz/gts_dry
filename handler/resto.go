package handler

import (
	"gts-dry/model"
	"gts-dry/service"

	"github.com/gofiber/fiber/v2"
)

type restoHandler struct {
	restoService service.RestoService
}

func NewRestoHandler(restoService service.RestoService) *restoHandler {
	return &restoHandler{restoService}
}

func (h *restoHandler) GetRestoAll(c *fiber.Ctx) error {
	data, err := h.restoService.GetRestoAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data resto gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data resto berhasil",
		Data:    data,
	})
}

func (h *restoHandler) GetRestoByKategori(c *fiber.Ctx) error {
	kategori := c.Params("kategori")
	if kategori == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "kategori parameter belum dimasukkan",
			Error:   "kategori parameter belum dimasukkan",
		})
	}

	data, err := h.restoService.GetRestoByKategori(kategori)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data resto gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data resto berhasil",
		Data:    data,
	})
}

func (h *restoHandler) GetRestoByKode(c *fiber.Ctx) error {
	kode := c.Params("kode")
	if kode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "kode parameter belum dimasukkan",
			Error:   "kode parameter belum dimasukkan",
		})
	}

	data, err := h.restoService.GetRestoByKode(kode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data resto gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data resto berhasil",
		Data:    data,
	})
}
