package handler

import (
	accountLockoutApp "nfxid/modules/auth/application/account_lockouts"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type AccountLockoutHandler struct {
	appSvc *accountLockoutApp.Service
}

func NewAccountLockoutHandler(appSvc *accountLockoutApp.Service) *AccountLockoutHandler {
	return &AccountLockoutHandler{
		appSvc: appSvc,
	}
}

// Create 创建账户锁定
func (h *AccountLockoutHandler) Create(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// GetByID 根据 ID 获取账户锁定
func (h *AccountLockoutHandler) GetByID(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// Update 更新账户锁定
func (h *AccountLockoutHandler) Update(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}

// Delete 删除账户锁定
func (h *AccountLockoutHandler) Delete(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Not implemented yet")
}
