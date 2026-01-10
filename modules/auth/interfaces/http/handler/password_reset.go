package handler

import (
	passwordResetApp "nfxid/modules/auth/application/password_resets"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type PasswordResetHandler struct {
	appSvc *passwordResetApp.Service
}

func NewPasswordResetHandler(appSvc *passwordResetApp.Service) *PasswordResetHandler {
	return &PasswordResetHandler{
		appSvc: appSvc,
	}
}

// Create 创建密码重置
func (h *PasswordResetHandler) Create(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// GetByID 根据 ID 获取密码重置
func (h *PasswordResetHandler) GetByID(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// Update 更新密码重置
func (h *PasswordResetHandler) Update(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// Delete 删除密码重置
func (h *PasswordResetHandler) Delete(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}
