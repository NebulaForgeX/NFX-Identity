package handler

import (
	passwordResetApp "nfxid/modules/auth/application/password_resets"
	passwordResetAppCommands "nfxid/modules/auth/application/password_resets/commands"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
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
func (h *PasswordResetHandler) Create(c *fiber.Ctx) error {
	var req reqdto.PasswordResetCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	passwordResetID, err := h.appSvc.CreatePasswordReset(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create password reset: "+err.Error())
	}

	// Get the created password reset
	passwordResetView, err := h.appSvc.GetPasswordReset(c.Context(), passwordResetID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created password reset: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Password reset created successfully", httpresp.SuccessOptions{Data: respdto.PasswordResetROToDTO(&passwordResetView)})
}

// GetByID 根据 ID 获取密码重置
func (h *PasswordResetHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.PasswordResetByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetPasswordReset(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Password reset not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Password reset retrieved successfully", httpresp.SuccessOptions{Data: respdto.PasswordResetROToDTO(&result)})
}

// Update 更新密码重置（更新状态）
func (h *PasswordResetHandler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return httpresp.Error(c, fiber.StatusBadRequest, "id is required")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid id: "+err.Error())
	}

	// Get the password reset to get its resetID
	passwordReset, err := h.appSvc.GetPasswordReset(c.Context(), id)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Password reset not found: "+err.Error())
	}

	var req reqdto.PasswordResetUpdateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateStatusCmd(passwordReset.ResetID)
	if err := h.appSvc.UpdateStatus(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update password reset: "+err.Error())
	}

	// Get the updated password reset
	updatedPasswordReset, err := h.appSvc.GetPasswordReset(c.Context(), id)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get updated password reset: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Password reset updated successfully", httpresp.SuccessOptions{Data: respdto.PasswordResetROToDTO(&updatedPasswordReset)})
}

// Delete 删除密码重置
func (h *PasswordResetHandler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return httpresp.Error(c, fiber.StatusBadRequest, "id is required")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid id: "+err.Error())
	}

	// Get the password reset to get its resetID
	passwordReset, err := h.appSvc.GetPasswordReset(c.Context(), id)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Password reset not found: "+err.Error())
	}

	cmd := passwordResetAppCommands.DeletePasswordResetCmd{
		ResetID: passwordReset.ResetID,
	}
	if err := h.appSvc.DeletePasswordReset(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete password reset: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Password reset deleted successfully")
}
