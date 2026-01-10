package handler

import (
	loginAttemptApp "nfxid/modules/auth/application/login_attempts"
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
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// GetByID 根据 ID 获取登录尝试
func (h *LoginAttemptHandler) GetByID(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// Delete 删除登录尝试
func (h *LoginAttemptHandler) Delete(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}
