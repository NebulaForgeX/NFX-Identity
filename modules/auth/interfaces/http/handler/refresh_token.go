package handler

import (
	refreshTokenApp "nfxid/modules/auth/application/refresh_tokens"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type RefreshTokenHandler struct {
	appSvc *refreshTokenApp.Service
}

func NewRefreshTokenHandler(appSvc *refreshTokenApp.Service) *RefreshTokenHandler {
	return &RefreshTokenHandler{
		appSvc: appSvc,
	}
}

// Create 创建刷新令牌
func (h *RefreshTokenHandler) Create(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// GetByID 根据 ID 获取刷新令牌
func (h *RefreshTokenHandler) GetByID(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// Update 更新刷新令牌
func (h *RefreshTokenHandler) Update(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// Delete 删除刷新令牌
func (h *RefreshTokenHandler) Delete(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}
