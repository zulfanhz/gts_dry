package handler

import (
	"gts-dry/model"
	"gts-dry/service"

	"github.com/gofiber/fiber/v2"
)

type mutasiHandler struct {
	mutasiService service.MutasiService
}

func NewMutasiHandler(mutasiService service.MutasiService) *mutasiHandler {
	return &mutasiHandler{mutasiService}
}

func (h *mutasiHandler) GetMutasiAll(c *fiber.Ctx) error {
	data, err := h.mutasiService.GetMutasiAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data mutasi gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data mutasi berhasil",
		Data:    data,
	})
}

func (h *mutasiHandler) AddMutasi(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	var req model.MutasiRakRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	data, err := h.mutasiService.AddMutasi(req, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "add mutasi gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "add mutasi berhasil",
		Data:    data,
	})
}
