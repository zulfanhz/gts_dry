package handler

import (
	"gts-dry/model"
	"gts-dry/service"

	"github.com/gofiber/fiber/v2"
)

type outgoingHandler struct {
	outgoingService service.OutgoingService
}

func NewOutgoingHandler(outgoingService service.OutgoingService) *outgoingHandler {
	return &outgoingHandler{outgoingService}
}

func (h *outgoingHandler) GetOutgoingAll(c *fiber.Ctx) error {
	data, err := h.outgoingService.GetOutgoingAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data outgoing gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data outgoing berhasil",
		Data:    data,
	})
}

func (h *outgoingHandler) GetOutgoingByKode(c *fiber.Ctx) error {
	kode := c.Params("kode")
	if kode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "kode parameter belum dimasukkan",
			Error:   "kode parameter belum dimasukkan",
		})
	}

	data, err := h.outgoingService.GetOutgoingByKode(kode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data outgoing gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data outgoing berhasil",
		Data:    data,
	})
}

func (h *outgoingHandler) GetOutgoingBySJ(c *fiber.Ctx) error {
	kode := c.Params("sj")
	if kode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "sj parameter belum dimasukkan",
			Error:   "sj parameter belum dimasukkan",
		})
	}

	data, err := h.outgoingService.GetOutgoingBySJ(kode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data outgoing gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data outgoing berhasil",
		Data:    data,
	})
}

func (h *outgoingHandler) AddOutgoing(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	var req model.OutgoingModel
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	data, err := h.outgoingService.AddOutgoing(req, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "add outgoing gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "add outgoing berhasil",
		Data:    data,
	})
}

func (h *outgoingHandler) UpdateOutgoing(c *fiber.Ctx) error {
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

	var req model.OutgoingModel
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	data, err := h.outgoingService.UpdateOutgoing(kode, req, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "update outgoing gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "update outgoing berhasil",
		Data:    data,
	})
}

func (h *outgoingHandler) DeleteOutgoing(c *fiber.Ctx) error {
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

	err := h.outgoingService.DeleteOutgoing(kode, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "delete outgoing gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "delete outgoing berhasil",
		Data:    true,
	})
}
