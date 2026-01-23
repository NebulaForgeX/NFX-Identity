package handler

import (
	loginAttemptApp "nfxid/modules/auth/application/login_attempts"
	loginAttemptAppCommands "nfxid/modules/auth/application/login_attempts/commands"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type LoginAttemptHandler struct {
	appSvc *loginAttemptApp.Service
}

func NewLoginAttemptHandler(appSvc *loginAttemptApp.Service) *LoginAttemptHandler {
	return &LoginAttemptHandler{
		appSvc: appSvc,
	}
}

// Create 创建登录尝试
func (h *LoginAttemptHandler) Create(c *fiber.Ctx) error {
	var req reqdto.LoginAttemptCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	loginAttemptID, err := h.appSvc.CreateLoginAttempt(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create login attempt: "+err.Error())
	}

	// Get the created login attempt
	loginAttemptView, err := h.appSvc.GetLoginAttempt(c.Context(), loginAttemptID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created login attempt: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Login attempt created successfully", httpresp.SuccessOptions{Data: respdto.LoginAttemptROToDTO(&loginAttemptView)})
}

// GetByID 根据 ID 获取登录尝试
func (h *LoginAttemptHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.LoginAttemptByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetLoginAttempt(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Login attempt not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Login attempt retrieved successfully", httpresp.SuccessOptions{Data: respdto.LoginAttemptROToDTO(&result)})
}

// Delete 删除登录尝试
func (h *LoginAttemptHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.LoginAttemptByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := loginAttemptAppCommands.DeleteLoginAttemptCmd{LoginAttemptID: req.ID}
	if err := h.appSvc.DeleteLoginAttempt(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete login attempt: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Login attempt deleted successfully")
}
