package handler

import (
	passwordHistoryApp "nfxid/modules/auth/application/password_history"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type PasswordHistoryHandler struct {
	appSvc *passwordHistoryApp.Service
}

func NewPasswordHistoryHandler(appSvc *passwordHistoryApp.Service) *PasswordHistoryHandler {
	return &PasswordHistoryHandler{
		appSvc: appSvc,
	}
}

// Create 创建密码历史
func (h *PasswordHistoryHandler) Create(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// GetByID 根据 ID 获取密码历史
func (h *PasswordHistoryHandler) GetByID(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}
