package handler

import (
	"gts-dry/model"
	"gts-dry/service"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) LoginUser(c *fiber.Ctx) error {
	var user model.LoginRequest
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.ErrBadRequest.Code,
			Message: "Login Gagal",
			Error:   err.Error(),
		})
	}

	data, token, err := h.userService.LoginUser(user.Email, user.Password)
	if err != nil {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.ErrUnauthorized.Code,
			Message: "Login Gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "Login Berhasil",
		Data: struct {
			User  model.UserResponse `json:"user"`
			Token *string            `json:"token"`
		}{
			User:  data,
			Token: token,
		},
	})
}

func (h *userHandler) GetUserByToken(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   "user id not found",
		})
	}

	data, err := h.userService.GetUserByEmail(id)
	if err != nil {
		return c.Status(401).JSON(model.ErrorResponse{
			Success: false,
			Code:    401,
			Message: "your user login not found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "get data user berhasil",
		Data:    data,
	})
}

func (h *userHandler) ChangePassword(c *fiber.Ctx) error {
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

	var req model.ChangePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	_, err = h.userService.GetUserByEmail(req.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "akun tidak ditemukan",
			Error:   err.Error(),
		})
	}

	data, err := h.userService.ChangePassword(req.Email, req.PasswordSekarang, req.PasswordBaru, req.PasswordBaruRepeat, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Message: "change password gagal",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.SuccessResponse{
		Success: true,
		Code:    fiber.StatusOK,
		Message: "change password berhasil",
		Data:    data,
	})
}
