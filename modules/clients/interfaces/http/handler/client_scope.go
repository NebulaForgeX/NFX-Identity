package handler

import (
	clientScopeApp "nfxid/modules/clients/application/client_scopes"
	clientScopeAppCommands "nfxid/modules/clients/application/client_scopes/commands"
	"nfxid/modules/clients/interfaces/http/dto/reqdto"
	"nfxid/modules/clients/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type ClientScopeHandler struct {
	appSvc *clientScopeApp.Service
}

func NewClientScopeHandler(appSvc *clientScopeApp.Service) *ClientScopeHandler {
	return &ClientScopeHandler{
		appSvc: appSvc,
	}
}

// Create 创建 Client Scope
func (h *ClientScopeHandler) Create(c *fiber.Ctx) error {
	var req reqdto.ClientScopeCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	clientScopeID, err := h.appSvc.CreateClientScope(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create client scope: "+err.Error())
	}

	// Get the created client scope
	clientScopeView, err := h.appSvc.GetClientScope(c.Context(), clientScopeID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created client scope: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Client scope created successfully", httpresp.SuccessOptions{Data: respdto.ClientScopeROToDTO(&clientScopeView)})
}

// GetByID 根据 ID 获取 Client Scope
func (h *ClientScopeHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ClientScopeByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetClientScope(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Client scope not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Client scope retrieved successfully", httpresp.SuccessOptions{Data: respdto.ClientScopeROToDTO(&result)})
}

// Delete 删除 Client Scope
func (h *ClientScopeHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ClientScopeByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := clientScopeAppCommands.DeleteClientScopeCmd{ClientScopeID: req.ID}
	if err := h.appSvc.DeleteClientScope(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete client scope: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Client scope deleted successfully")
}
