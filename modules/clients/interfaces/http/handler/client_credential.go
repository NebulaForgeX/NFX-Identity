package handler

import (
	clientCredentialApp "nfxid/modules/clients/application/client_credentials"
	clientCredentialAppCommands "nfxid/modules/clients/application/client_credentials/commands"
	"nfxid/modules/clients/interfaces/http/dto/reqdto"
	"nfxid/modules/clients/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type ClientCredentialHandler struct {
	appSvc *clientCredentialApp.Service
}

func NewClientCredentialHandler(appSvc *clientCredentialApp.Service) *ClientCredentialHandler {
	return &ClientCredentialHandler{
		appSvc: appSvc,
	}
}

// Create 创建 Client Credential
func (h *ClientCredentialHandler) Create(c *fiber.Ctx) error {
	var req reqdto.ClientCredentialCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	clientCredentialID, err := h.appSvc.CreateClientCredential(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create client credential: "+err.Error())
	}

	// Get the created client credential
	clientCredentialView, err := h.appSvc.GetClientCredential(c.Context(), clientCredentialID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created client credential: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Client credential created successfully", httpresp.SuccessOptions{Data: respdto.ClientCredentialROToDTO(&clientCredentialView)})
}

// GetByID 根据 ID 获取 Client Credential
func (h *ClientCredentialHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ClientCredentialByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetClientCredential(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Client credential not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Client credential retrieved successfully", httpresp.SuccessOptions{Data: respdto.ClientCredentialROToDTO(&result)})
}

// Delete 删除 Client Credential
func (h *ClientCredentialHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ClientCredentialDeleteRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := clientCredentialAppCommands.DeleteClientCredentialCmd{ClientID: req.ClientID}
	if err := h.appSvc.DeleteClientCredential(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete client credential: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Client credential deleted successfully")
}
