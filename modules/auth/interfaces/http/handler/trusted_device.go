package handler

import (
	trustedDeviceApp "nfxid/modules/auth/application/trusted_devices"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type TrustedDeviceHandler struct {
	appSvc *trustedDeviceApp.Service
}

func NewTrustedDeviceHandler(appSvc *trustedDeviceApp.Service) *TrustedDeviceHandler {
	return &TrustedDeviceHandler{
		appSvc: appSvc,
	}
}

// Create 创建受信任设备
func (h *TrustedDeviceHandler) Create(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// GetByID 根据 ID 获取受信任设备
func (h *TrustedDeviceHandler) GetByID(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// Delete 删除受信任设备
func (h *TrustedDeviceHandler) Delete(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}
