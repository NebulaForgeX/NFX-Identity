package handler

import (
	authorizationCodeApp "nfxid/modules/permission/application/authorization_code"
	authorizationCodeAppCommands "nfxid/modules/permission/application/authorization_code/commands"
	"nfxid/modules/permission/interfaces/http/dto/reqdto"
	"nfxid/modules/permission/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type AuthorizationCodeHandler struct {
	appSvc *authorizationCodeApp.Service
}

func NewAuthorizationCodeHandler(appSvc *authorizationCodeApp.Service) *AuthorizationCodeHandler {
	return &AuthorizationCodeHandler{
		appSvc: appSvc,
	}
}

// Create 创建授权码
func (h *AuthorizationCodeHandler) Create(c *fiber.Ctx) error {
	var req reqdto.AuthorizationCodeCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	result, err := h.appSvc.CreateAuthorizationCode(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create authorization code: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Authorization code created successfully", httpresp.SuccessOptions{
		Data: respdto.AuthorizationCodeDomainToDTO(result),
	})
}

// Use 使用授权码
func (h *AuthorizationCodeHandler) Use(c *fiber.Ctx) error {
	var req reqdto.AuthorizationCodeUseRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUseCmd()
	err := h.appSvc.UseAuthorizationCode(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Failed to use authorization code: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Authorization code used successfully", httpresp.SuccessOptions{})
}

// GetByID 根据ID获取授权码
func (h *AuthorizationCodeHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.AuthorizationCodeByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid path parameter: "+err.Error())
	}

	result, err := h.appSvc.GetAuthorizationCode(c.Context(), authorizationCodeAppCommands.GetAuthorizationCodeCmd{
		ID: req.ID,
	})
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Authorization code not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Authorization code retrieved successfully", httpresp.SuccessOptions{
		Data: respdto.AuthorizationCodeDomainToDTO(result),
	})
}

// GetByCode 根据Code获取授权码
func (h *AuthorizationCodeHandler) GetByCode(c *fiber.Ctx) error {
	var req reqdto.AuthorizationCodeByCodeRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid path parameter: "+err.Error())
	}

	result, err := h.appSvc.GetAuthorizationCodeByCode(c.Context(), authorizationCodeAppCommands.GetAuthorizationCodeByCodeCmd{
		Code: req.Code,
	})
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Authorization code not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Authorization code retrieved successfully", httpresp.SuccessOptions{
		Data: respdto.AuthorizationCodeDomainToDTO(result),
	})
}

// Delete 删除授权码
func (h *AuthorizationCodeHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.AuthorizationCodeByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid path parameter: "+err.Error())
	}

	err := h.appSvc.DeleteAuthorizationCode(c.Context(), authorizationCodeAppCommands.DeleteAuthorizationCodeCmd{
		ID: req.ID,
	})
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete authorization code: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Authorization code deleted successfully", httpresp.SuccessOptions{})
}

// Activate 激活授权码
func (h *AuthorizationCodeHandler) Activate(c *fiber.Ctx) error {
	var req reqdto.AuthorizationCodeByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid path parameter: "+err.Error())
	}

	err := h.appSvc.ActivateAuthorizationCode(c.Context(), authorizationCodeAppCommands.ActivateAuthorizationCodeCmd{
		ID: req.ID,
	})
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to activate authorization code: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Authorization code activated successfully", httpresp.SuccessOptions{})
}

// Deactivate 停用授权码
func (h *AuthorizationCodeHandler) Deactivate(c *fiber.Ctx) error {
	var req reqdto.AuthorizationCodeByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid path parameter: "+err.Error())
	}

	err := h.appSvc.DeactivateAuthorizationCode(c.Context(), authorizationCodeAppCommands.DeactivateAuthorizationCodeCmd{
		ID: req.ID,
	})
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to deactivate authorization code: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Authorization code deactivated successfully", httpresp.SuccessOptions{})
}
