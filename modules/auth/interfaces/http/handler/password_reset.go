package handler

import (
	passwordResetApp "nfxid/modules/auth/application/password_resets"
	passwordResetAppCommands "nfxid/modules/auth/application/password_resets/commands"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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
func (h *PasswordResetHandler) Create(c fiber.Ctx) error {
	var req reqdto.PasswordResetCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	passwordResetID, err := h.appSvc.CreatePasswordReset(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created password reset
	passwordResetView, err := h.appSvc.GetPasswordReset(c.Context(), passwordResetID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Password reset created successfully", httpx.SuccessOptions{Data: respdto.PasswordResetROToDTO(&passwordResetView)})
}

// GetByID 根据 ID 获取密码重置
func (h *PasswordResetHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.PasswordResetByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetPasswordReset(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Password reset retrieved successfully", httpx.SuccessOptions{Data: respdto.PasswordResetROToDTO(&result)})
}

// Update 更新密码重置（更新状态）
func (h *PasswordResetHandler) Update(c fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "id is required"))
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	// Get the password reset to get its resetID
	passwordReset, err := h.appSvc.GetPasswordReset(c.Context(), id)
	if err != nil {
		return err
	}

	var req reqdto.PasswordResetUpdateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateStatusCmd(passwordReset.ResetID)
	if err := h.appSvc.UpdateStatus(c.Context(), cmd); err != nil {
		return err
	}

	// Get the updated password reset
	updatedPasswordReset, err := h.appSvc.GetPasswordReset(c.Context(), id)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Password reset updated successfully", httpx.SuccessOptions{Data: respdto.PasswordResetROToDTO(&updatedPasswordReset)})
}

// Delete 删除密码重置
func (h *PasswordResetHandler) Delete(c fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "id is required"))
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	// Get the password reset to get its resetID
	passwordReset, err := h.appSvc.GetPasswordReset(c.Context(), id)
	if err != nil {
		return err
	}

	cmd := passwordResetAppCommands.DeletePasswordResetCmd{
		ResetID: passwordReset.ResetID,
	}
	if err := h.appSvc.DeletePasswordReset(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Password reset deleted successfully")
}

// fiber:context-methods migrated
