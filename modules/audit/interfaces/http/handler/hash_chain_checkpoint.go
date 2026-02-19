package handler

import (
	hashChainCheckpointApp "nfxid/modules/audit/application/hash_chain_checkpoints"
	hashChainCheckpointAppCommands "nfxid/modules/audit/application/hash_chain_checkpoints/commands"
	"nfxid/modules/audit/interfaces/http/dto/reqdto"
	"nfxid/modules/audit/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type HashChainCheckpointHandler struct {
	appSvc *hashChainCheckpointApp.Service
}

func NewHashChainCheckpointHandler(appSvc *hashChainCheckpointApp.Service) *HashChainCheckpointHandler {
	return &HashChainCheckpointHandler{appSvc: appSvc}
}

func (h *HashChainCheckpointHandler) Create(c fiber.Ctx) error {
	var req reqdto.HashChainCheckpointCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	hashChainCheckpointID, err := h.appSvc.CreateHashChainCheckpoint(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created hash chain checkpoint
	hashChainCheckpointView, err := h.appSvc.GetHashChainCheckpoint(c.Context(), hashChainCheckpointID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Hash chain checkpoint created successfully", httpx.SuccessOptions{Data: respdto.HashChainCheckpointROToDTO(&hashChainCheckpointView)})
}

func (h *HashChainCheckpointHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.HashChainCheckpointByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetHashChainCheckpoint(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Hash chain checkpoint retrieved successfully", httpx.SuccessOptions{Data: respdto.HashChainCheckpointROToDTO(&result)})
}

func (h *HashChainCheckpointHandler) Delete(c fiber.Ctx) error {
	var req reqdto.HashChainCheckpointByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := hashChainCheckpointAppCommands.DeleteHashChainCheckpointCmd{HashChainCheckpointID: req.ID}
	if err := h.appSvc.DeleteHashChainCheckpoint(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Hash chain checkpoint deleted successfully")
}

// fiber:context-methods migrated
