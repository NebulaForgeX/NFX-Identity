package handler

import (
	clientScopeApp "nfxid/modules/clients/application/client_scopes"
	clientScopeAppCommands "nfxid/modules/clients/application/client_scopes/commands"
	"nfxid/modules/clients/interfaces/http/dto/reqdto"
	"nfxid/modules/clients/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
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
func (h *ClientScopeHandler) Create(c fiber.Ctx) error {
	var req reqdto.ClientScopeCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	clientScopeID, err := h.appSvc.CreateClientScope(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created client scope
	clientScopeView, err := h.appSvc.GetClientScope(c.Context(), clientScopeID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Client scope created successfully", httpx.SuccessOptions{Data: respdto.ClientScopeROToDTO(&clientScopeView)})
}

// GetByID 根据 ID 获取 Client Scope
func (h *ClientScopeHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ClientScopeByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetClientScope(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Client scope retrieved successfully", httpx.SuccessOptions{Data: respdto.ClientScopeROToDTO(&result)})
}

// Delete 删除 Client Scope
func (h *ClientScopeHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ClientScopeByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := clientScopeAppCommands.DeleteClientScopeCmd{ClientScopeID: req.ID}
	if err := h.appSvc.DeleteClientScope(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Client scope deleted successfully")
}

// fiber:context-methods migrated
