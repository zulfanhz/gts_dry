package handler

import (
	"gts-dry/model"
	"gts-dry/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type barangHandler struct {
	barangService service.BarangService
}

func NewBarangHandler(barangService service.BarangService) *barangHandler {
	return &barangHandler{barangService}
}

func (h *barangHandler) GetBarangAll(c *fiber.Ctx) error {
	data, err := h.barangService.GetBarangAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data barang gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data barang berhasil",
		Data:    data,
	})
}

func (h *barangHandler) GetBarangByKategori(c *fiber.Ctx) error {
	kode := c.Params("kategori")
	if kode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "kategori parameter belum dimasukkan",
			Error:   "kategori parameter belum dimasukkan",
		})
	}

	data, err := h.barangService.GetBarangByKategori(kode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data barang gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data barang berhasil",
		Data:    data,
	})
}

func (h *barangHandler) GetBarangByKode(c *fiber.Ctx) error {
	kode := c.Params("kode")
	if kode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "kode parameter belum dimasukkan",
			Error:   "kode parameter belum dimasukkan",
		})
	}

	data, err := h.barangService.GetBarangByKode(kode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data barang gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data barang berhasil",
		Data:    data,
	})
}

func (h *barangHandler) AddBarang(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	var req model.BarangWithoutUser
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	data, err := h.barangService.SaveBarang(req, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "add barang gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "add barang berhasil",
		Data:    data,
	})
}

func (h *barangHandler) UpdateBarang(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	kode := c.Params("kode")
	if kode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "kode parameter belum dimasukkan",
			Error:   "kode parameter belum dimasukkan",
		})
	}

	var req model.BarangWithoutUser
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	data, err := h.barangService.UpdateBarang(kode, req, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "update barang gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "update barang berhasil",
		Data:    data,
	})
}

func (h *barangHandler) DeleteBarang(c *fiber.Ctx) error {
	kode := c.Params("kode")
	if kode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "kode parameter belum dimasukkan",
			Error:   "kode parameter belum dimasukkan",
		})
	}

	err := h.barangService.DeleteBarang(kode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "delete barang gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "delete barang berhasil",
		Data:    true,
	})
}

func (h *barangHandler) AddSatuanBarang(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	var req model.BarangSatuanModel
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	data, err := h.barangService.SaveSatuan(req, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "add satuan barang gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "add satuan barang berhasil",
		Data:    data,
	})
}

func (h *barangHandler) UpdateSatuanBarang(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	kode := c.Params("kode")
	if kode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "kode parameter belum dimasukkan",
			Error:   "kode parameter belum dimasukkan",
		})
	}

	satuan := c.Params("satuan")
	if satuan == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "satuan parameter belum dimasukkan",
			Error:   "satuan parameter belum dimasukkan",
		})
	}

	var req model.BarangSatuanModel
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	data, err := h.barangService.UpdateSatuan(req, kode, satuan, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "update satuan barang gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "update satuan barang berhasil",
		Data:    data,
	})
}

func (h *barangHandler) DeleteSatuanBarang(c *fiber.Ctx) error {
	kode := c.Params("kode")
	if kode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "kode parameter belum dimasukkan",
			Error:   "kode parameter belum dimasukkan",
		})
	}

	satuan := c.Params("satuan")
	if satuan == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "satuan parameter belum dimasukkan",
			Error:   "satuan parameter belum dimasukkan",
		})
	}

	isHitung := c.Params("ishitung")
	if isHitung == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "isHitung parameter belum dimasukkan",
			Error:   "isHitung parameter belum dimasukkan",
		})
	}

	htg, err := strconv.Atoi(isHitung)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "isHitung parameter hanya 0 atau 1",
			Error:   "isHitung parameter hanya 0 atau 1",
		})
	}

	err = h.barangService.DeleteSatuan(kode, satuan, htg)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "delete satuan barang gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "delete satuan barang berhasil",
		Data:    true,
	})
}
