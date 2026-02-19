package handler

import (
	refreshTokenApp "nfxid/modules/auth/application/refresh_tokens"
	refreshTokenAppCommands "nfxid/modules/auth/application/refresh_tokens/commands"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
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
func (h *RefreshTokenHandler) Create(c fiber.Ctx) error {
	var req reqdto.RefreshTokenCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	refreshTokenID, err := h.appSvc.CreateRefreshToken(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created refresh token
	refreshTokenView, err := h.appSvc.GetRefreshToken(c.Context(), refreshTokenID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Refresh token created successfully", httpx.SuccessOptions{Data: respdto.RefreshTokenROToDTO(&refreshTokenView)})
}

// GetByID 根据 ID 获取刷新令牌
func (h *RefreshTokenHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.RefreshTokenByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetRefreshToken(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Refresh token retrieved successfully", httpx.SuccessOptions{Data: respdto.RefreshTokenROToDTO(&result)})
}

// Update 更新刷新令牌（撤销）
func (h *RefreshTokenHandler) Update(c fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_PARAMS", "id is required"))
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	// Get the refresh token to get its tokenID
	refreshToken, err := h.appSvc.GetRefreshToken(c.Context(), id)
	if err != nil {
		return err
	}

	var req reqdto.RefreshTokenRevokeRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToRevokeCmd(refreshToken.TokenID)
	if err := h.appSvc.RevokeRefreshToken(c.Context(), cmd); err != nil {
		return err
	}

	// Get the updated refresh token
	updatedRefreshToken, err := h.appSvc.GetRefreshToken(c.Context(), id)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Refresh token revoked successfully", httpx.SuccessOptions{Data: respdto.RefreshTokenROToDTO(&updatedRefreshToken)})
}

// Delete 删除刷新令牌
func (h *RefreshTokenHandler) Delete(c fiber.Ctx) error {
	var req reqdto.RefreshTokenByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := refreshTokenAppCommands.DeleteRefreshTokenCmd{
		RefreshTokenID: req.ID,
	}
	if err := h.appSvc.DeleteRefreshToken(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Refresh token deleted successfully")
}

// fiber:context-methods migrated
