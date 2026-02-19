package handler

import (
	accountLockoutApp "nfxid/modules/auth/application/account_lockouts"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
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
func (h *AccountLockoutHandler) Create(c fiber.Ctx) error {
	var req reqdto.AccountLockoutCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd, err := req.ToCreateCmd()
	if err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	if err := h.appSvc.CreateAccountLockout(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.Created(c, "Account lockout created successfully")
}

// GetByUserID 根据 UserID 获取账户锁定
func (h *AccountLockoutHandler) GetByUserID(c fiber.Ctx) error {
	userIDStr := c.Params("user_id")
	if userIDStr == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "user_id is required"))
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetAccountLockout(c.Context(), userID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Account lockout retrieved successfully", httpx.SuccessOptions{Data: respdto.AccountLockoutROToDTO(&result)})
}

// Unlock 解锁账户
func (h *AccountLockoutHandler) Unlock(c fiber.Ctx) error {
	var req reqdto.AccountLockoutUnlockRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd, err := req.ToUnlockCmd()
	if err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	if err := h.appSvc.UnlockAccount(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Account unlocked successfully")
}

// Delete 删除账户锁定
func (h *AccountLockoutHandler) Delete(c fiber.Ctx) error {
	userIDStr := c.Params("user_id")
	if userIDStr == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "user_id is required"))
	}

	req := reqdto.AccountLockoutDeleteRequestDTO{UserID: userIDStr}
	cmd, err := req.ToDeleteCmd()
	if err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	if err := h.appSvc.DeleteAccountLockout(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Account lockout deleted successfully")
}

// fiber:context-methods migrated
