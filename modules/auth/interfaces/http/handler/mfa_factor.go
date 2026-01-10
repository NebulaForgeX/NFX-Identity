package handler

import (
	mfaFactorApp "nfxid/modules/auth/application/mfa_factors"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type MFAFactorHandler struct {
	appSvc *mfaFactorApp.Service
}

func NewMFAFactorHandler(appSvc *mfaFactorApp.Service) *MFAFactorHandler {
	return &MFAFactorHandler{
		appSvc: appSvc,
	}
}

// Create 创建 MFA 因子
func (h *MFAFactorHandler) Create(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// GetByID 根据 ID 获取 MFA 因子
func (h *MFAFactorHandler) GetByID(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// Update 更新 MFA 因子
func (h *MFAFactorHandler) Update(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// Delete 删除 MFA 因子
func (h *MFAFactorHandler) Delete(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}
