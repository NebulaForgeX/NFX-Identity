package handler

import (
	refreshTokenApp "nfxid/modules/auth/application/refresh_tokens"
	refreshTokenAppCommands "nfxid/modules/auth/application/refresh_tokens/commands"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	var req reqdto.RefreshTokenCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	refreshTokenID, err := h.appSvc.CreateRefreshToken(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create refresh token: "+err.Error())
	}

	// Get the created refresh token
	refreshTokenView, err := h.appSvc.GetRefreshToken(c.Context(), refreshTokenID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created refresh token: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Refresh token created successfully", httpresp.SuccessOptions{Data: respdto.RefreshTokenROToDTO(&refreshTokenView)})
}

// GetByID 根据 ID 获取刷新令牌
func (h *RefreshTokenHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.RefreshTokenByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetRefreshToken(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Refresh token not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Refresh token retrieved successfully", httpresp.SuccessOptions{Data: respdto.RefreshTokenROToDTO(&result)})
}

// Update 更新刷新令牌（撤销）
func (h *RefreshTokenHandler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return httpresp.Error(c, fiber.StatusBadRequest, "id is required")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid id: "+err.Error())
	}

	// Get the refresh token to get its tokenID
	refreshToken, err := h.appSvc.GetRefreshToken(c.Context(), id)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Refresh token not found: "+err.Error())
	}

	var req reqdto.RefreshTokenRevokeRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToRevokeCmd(refreshToken.TokenID)
	if err := h.appSvc.RevokeRefreshToken(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to revoke refresh token: "+err.Error())
	}

	// Get the updated refresh token
	updatedRefreshToken, err := h.appSvc.GetRefreshToken(c.Context(), id)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get updated refresh token: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Refresh token revoked successfully", httpresp.SuccessOptions{Data: respdto.RefreshTokenROToDTO(&updatedRefreshToken)})
}

// Delete 删除刷新令牌
func (h *RefreshTokenHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.RefreshTokenByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := refreshTokenAppCommands.DeleteRefreshTokenCmd{
		RefreshTokenID: req.ID,
	}
	if err := h.appSvc.DeleteRefreshToken(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete refresh token: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Refresh token deleted successfully")
}
