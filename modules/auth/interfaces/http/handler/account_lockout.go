package handler

import (
	accountLockoutApp "nfxid/modules/auth/application/account_lockouts"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	var req reqdto.AccountLockoutCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd, err := req.ToCreateCmd()
	if err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid user_id: "+err.Error())
	}

	if err := h.appSvc.CreateAccountLockout(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create account lockout: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Account lockout created successfully")
}

// GetByUserID 根据 UserID 获取账户锁定
func (h *AccountLockoutHandler) GetByUserID(c *fiber.Ctx) error {
	userIDStr := c.Params("user_id")
	if userIDStr == "" {
		return httpresp.Error(c, fiber.StatusBadRequest, "user_id is required")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid user_id: "+err.Error())
	}

	result, err := h.appSvc.GetAccountLockout(c.Context(), userID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Account lockout not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Account lockout retrieved successfully", httpresp.SuccessOptions{Data: respdto.AccountLockoutROToDTO(&result)})
}

// Unlock 解锁账户
func (h *AccountLockoutHandler) Unlock(c *fiber.Ctx) error {
	var req reqdto.AccountLockoutUnlockRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd, err := req.ToUnlockCmd()
	if err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid user_id: "+err.Error())
	}

	if err := h.appSvc.UnlockAccount(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to unlock account: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Account unlocked successfully")
}

// Delete 删除账户锁定
func (h *AccountLockoutHandler) Delete(c *fiber.Ctx) error {
	userIDStr := c.Params("user_id")
	if userIDStr == "" {
		return httpresp.Error(c, fiber.StatusBadRequest, "user_id is required")
	}

	req := reqdto.AccountLockoutDeleteRequestDTO{UserID: userIDStr}
	cmd, err := req.ToDeleteCmd()
	if err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid user_id: "+err.Error())
	}

	if err := h.appSvc.DeleteAccountLockout(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete account lockout: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Account lockout deleted successfully")
}
