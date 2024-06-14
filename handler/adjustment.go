package handler

import (
	"gts-dry/model"
	"gts-dry/service"

	"github.com/gofiber/fiber/v2"
)

type adjustmentHandler struct {
	adjustmentService service.AdjustmentService
}

func NewAdjustmentHandler(adjustmentService service.AdjustmentService) *adjustmentHandler {
	return &adjustmentHandler{adjustmentService}
}

func (h *adjustmentHandler) GetAdjusmentAll(c *fiber.Ctx) error {
	data, err := h.adjustmentService.GetAdjustmentAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data adjustment gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data adjustment berhasil",
		Data:    data,
	})
}

func (h *adjustmentHandler) GetAdjusmentByKode(c *fiber.Ctx) error {
	kode := c.Params("kode")
	if kode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "kode parameter belum dimasukkan",
			Error:   "kode parameter belum dimasukkan",
		})
	}

	data, err := h.adjustmentService.GetAdjustmentByKode(kode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data adjustment gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data adjustment berhasil",
		Data:    data,
	})
}

func (h *adjustmentHandler) AddAdjustment(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	var req model.AdjustmentRequestModel
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	data, err := h.adjustmentService.AddAdjustment(req, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "add adjustment gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "add adjustment berhasil",
		Data:    data,
	})
}

func (h *adjustmentHandler) UpdateAdjustment(c *fiber.Ctx) error {
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

	var req model.AdjustmentRequestModel
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	data, err := h.adjustmentService.UpdateAdjustment(kode, req, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "update adjustment gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "update adjustment berhasil",
		Data:    data,
	})
}

func (h *adjustmentHandler) DeleteAdjustment(c *fiber.Ctx) error {
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

	err := h.adjustmentService.DeleteAdjustment(kode, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "delete adjustment gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "delete adjustment berhasil",
		Data:    true,
	})
}
