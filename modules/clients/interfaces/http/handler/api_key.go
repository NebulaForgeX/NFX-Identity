package handler

import (
	apiKeyApp "nfxid/modules/clients/application/api_keys"
	apiKeyAppCommands "nfxid/modules/clients/application/api_keys/commands"
	"nfxid/modules/clients/interfaces/http/dto/reqdto"
	"nfxid/modules/clients/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
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
func (h *APIKeyHandler) Create(c *fiber.Ctx) error {
	var req reqdto.APIKeyCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	apiKeyID, err := h.appSvc.CreateAPIKey(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create API key: "+err.Error())
	}

	// Get the created API key
	apiKeyView, err := h.appSvc.GetAPIKey(c.Context(), apiKeyID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created API key: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "API key created successfully", httpresp.SuccessOptions{Data: respdto.APIKeyROToDTO(&apiKeyView)})
}

// GetByID 根据 ID 获取 API Key
func (h *APIKeyHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.APIKeyByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetAPIKey(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "API key not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "API key retrieved successfully", httpresp.SuccessOptions{Data: respdto.APIKeyROToDTO(&result)})
}

// Delete 删除 API Key
func (h *APIKeyHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.APIKeyDeleteRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := apiKeyAppCommands.DeleteAPIKeyCmd{KeyID: req.KeyID}
	if err := h.appSvc.DeleteAPIKey(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete API key: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "API key deleted successfully")
}
