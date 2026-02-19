package handler

import (
	clientCredentialApp "nfxid/modules/clients/application/client_credentials"
	clientCredentialAppCommands "nfxid/modules/clients/application/client_credentials/commands"
	"nfxid/modules/clients/interfaces/http/dto/reqdto"
	"nfxid/modules/clients/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
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
func (h *ClientCredentialHandler) Create(c fiber.Ctx) error {
	var req reqdto.ClientCredentialCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	clientCredentialID, err := h.appSvc.CreateClientCredential(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created client credential
	clientCredentialView, err := h.appSvc.GetClientCredential(c.Context(), clientCredentialID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Client credential created successfully", httpx.SuccessOptions{Data: respdto.ClientCredentialROToDTO(&clientCredentialView)})
}

// GetByID 根据 ID 获取 Client Credential
func (h *ClientCredentialHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ClientCredentialByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetClientCredential(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Client credential retrieved successfully", httpx.SuccessOptions{Data: respdto.ClientCredentialROToDTO(&result)})
}

// Delete 删除 Client Credential
func (h *ClientCredentialHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ClientCredentialDeleteRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := clientCredentialAppCommands.DeleteClientCredentialCmd{ClientID: req.ClientID}
	if err := h.appSvc.DeleteClientCredential(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Client credential deleted successfully")
}

// fiber:context-methods migrated
