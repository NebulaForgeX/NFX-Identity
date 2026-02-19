package handler

import (
	apiKeyApp "nfxid/modules/clients/application/api_keys"
	apiKeyAppCommands "nfxid/modules/clients/application/api_keys/commands"
	"nfxid/modules/clients/interfaces/http/dto/reqdto"
	"nfxid/modules/clients/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type APIKeyHandler struct {
	appSvc *apiKeyApp.Service
}

func NewAPIKeyHandler(appSvc *apiKeyApp.Service) *APIKeyHandler {
	return &APIKeyHandler{
		appSvc: appSvc,
	}
}

// Create 创建 API Key
func (h *APIKeyHandler) Create(c fiber.Ctx) error {
	var req reqdto.APIKeyCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	apiKeyID, err := h.appSvc.CreateAPIKey(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created API key
	apiKeyView, err := h.appSvc.GetAPIKey(c.Context(), apiKeyID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "API key created successfully", httpx.SuccessOptions{Data: respdto.APIKeyROToDTO(&apiKeyView)})
}

// GetByID 根据 ID 获取 API Key
func (h *APIKeyHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.APIKeyByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetAPIKey(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "API key retrieved successfully", httpx.SuccessOptions{Data: respdto.APIKeyROToDTO(&result)})
}

// Delete 删除 API Key
func (h *APIKeyHandler) Delete(c fiber.Ctx) error {
	var req reqdto.APIKeyDeleteRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := apiKeyAppCommands.DeleteAPIKeyCmd{KeyID: req.KeyID}
	if err := h.appSvc.DeleteAPIKey(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "API key deleted successfully")
}

// fiber:context-methods migrated
