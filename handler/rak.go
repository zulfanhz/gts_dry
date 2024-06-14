package handler

import (
	"gts-dry/model"
	"gts-dry/service"

	"github.com/gofiber/fiber/v2"
)

type rakHandler struct {
	rakService  service.RakService
	userService service.UserService
}

func NewRakHandler(rakService service.RakService, userService service.UserService) *rakHandler {
	return &rakHandler{rakService, userService}
}

func (h *rakHandler) GetRakAll(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	_, err := h.userService.GetUserByEmail(id)
	if err != nil {
		return c.Status(401).JSON(model.ErrorResponse{
			Success: false,
			Code:    401,
			Message: "your user login not found",
			Error:   err.Error(),
		})
	}

	data, err := h.rakService.GetRakAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data rak gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data rak berhasil",
		Data:    data,
	})
}

func (h *rakHandler) GetRakByType(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	types := c.Params("type")
	if types == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "type parameter belum dimasukkan",
			Error:   "type parameter belum dimasukkan",
		})
	}

	_, err := h.userService.GetUserByEmail(id)
	if err != nil {
		return c.Status(401).JSON(model.ErrorResponse{
			Success: false,
			Code:    401,
			Message: "your user login not found",
			Error:   err.Error(),
		})
	}

	data, err := h.rakService.GetRakByType(types)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data rak gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data rak berhasil",
		Data:    data,
	})
}

func (h *rakHandler) GetRakByJenis(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	jenis := c.Params("jenis")
	if jenis == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "jenis parameter belum dimasukkan",
			Error:   "jenis parameter belum dimasukkan",
		})
	}

	_, err := h.userService.GetUserByEmail(id)
	if err != nil {
		return c.Status(401).JSON(model.ErrorResponse{
			Success: false,
			Code:    401,
			Message: "your user login not found",
			Error:   err.Error(),
		})
	}

	data, err := h.rakService.GetRakByJenis(jenis)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data rak gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data rak berhasil",
		Data:    data,
	})
}

func (h *rakHandler) GetRakByKodeRak(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	codeRak := c.Params("kode")
	if codeRak == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "kode parameter belum dimasukkan",
			Error:   "kode parameter belum dimasukkan",
		})
	}

	_, err := h.userService.GetUserByEmail(id)
	if err != nil {
		return c.Status(401).JSON(model.ErrorResponse{
			Success: false,
			Code:    401,
			Message: "your user login not found",
			Error:   err.Error(),
		})
	}

	data, err := h.rakService.GetRakByKodeRak(codeRak)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data rak gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data rak berhasil",
		Data:    data,
	})
}

func (h *rakHandler) GetRakIsiByProductCode(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	productCode := c.Params("kode")
	if productCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "kode parameter belum dimasukkan",
			Error:   "kode parameter belum dimasukkan",
		})
	}

	_, err := h.userService.GetUserByEmail(id)
	if err != nil {
		return c.Status(401).JSON(model.ErrorResponse{
			Success: false,
			Code:    401,
			Message: "your user login not found",
			Error:   err.Error(),
		})
	}

	data, err := h.rakService.GetRakIsiByProductCode(productCode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data rak gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data rak berhasil",
		Data:    data,
	})
}

func (h *rakHandler) GetRakIsiByProductCodedanRak(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	productCode := c.Params("productCode")
	if productCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "code parameter belum dimasukkan",
			Error:   "code parameter belum dimasukkan",
		})
	}

	rakCode := c.Params("rakCode")
	if rakCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "rak parameter belum dimasukkan",
			Error:   "rak parameter belum dimasukkan",
		})
	}

	exp := c.Params("exp")
	if exp == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "exp parameter belum dimasukkan",
			Error:   "exp parameter belum dimasukkan",
		})
	}

	_, err := h.userService.GetUserByEmail(id)
	if err != nil {
		return c.Status(401).JSON(model.ErrorResponse{
			Success: false,
			Code:    401,
			Message: "your user login not found",
			Error:   err.Error(),
		})
	}

	data, err := h.rakService.GetRakIsiByProductRakExp(productCode, rakCode, exp)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data rak gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data rak berhasil",
		Data:    data,
	})
}

func (h *rakHandler) CekRakListAvailableIncoming(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	productCode := c.Params("kode")
	if productCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "type parameter belum dimasukkan",
			Error:   "type parameter belum dimasukkan",
		})
	}

	productCategory := c.Params("productCategory")

	if productCategory == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "kategori parameter belum dimasukkan",
			Error:   "kategori parameter belum dimasukkan",
		})
	}

	exp := c.Params("exp")
	if exp == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "exp parameter belum dimasukkan",
			Error:   "exp parameter belum dimasukkan",
		})
	}

	_, err := h.userService.GetUserByEmail(id)
	if err != nil {
		return c.Status(401).JSON(model.ErrorResponse{
			Success: false,
			Code:    401,
			Message: "your user login not found",
			Error:   err.Error(),
		})
	}

	data, err := h.rakService.CekRakListAvailableIncoming(productCode, productCategory, exp)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "get data rak gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data rak berhasil",
		Data:    data,
	})
}

func (h *rakHandler) AddRak(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	var req model.RakModelWithoutUser
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	data, err := h.rakService.AddRak(req, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "add rak gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "add rak berhasil",
		Data:    data,
	})
}

func (h *rakHandler) UpdateRak(c *fiber.Ctx) error {
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

	var req model.RakModelWithoutUser
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	data, err := h.rakService.UpdateRak(kode, req, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "update rak gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "update rak berhasil",
		Data:    data,
	})
}

func (h *rakHandler) DeleteRak(c *fiber.Ctx) error {
	kode := c.Params("kode")
	if kode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "kode parameter belum dimasukkan",
			Error:   "kode parameter belum dimasukkan",
		})
	}

	err := h.rakService.DeleteRak(kode)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "delete rak gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "delete rak berhasil",
		Data:    true,
	})
}
