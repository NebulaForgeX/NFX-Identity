package handler

import (
	hashChainCheckpointApp "nfxid/modules/audit/application/hash_chain_checkpoints"
	hashChainCheckpointAppCommands "nfxid/modules/audit/application/hash_chain_checkpoints/commands"
	"nfxid/modules/audit/interfaces/http/dto/reqdto"
	"nfxid/modules/audit/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type HashChainCheckpointHandler struct {
	appSvc *hashChainCheckpointApp.Service
}

func NewHashChainCheckpointHandler(appSvc *hashChainCheckpointApp.Service) *HashChainCheckpointHandler {
	return &HashChainCheckpointHandler{appSvc: appSvc}
}

func (h *HashChainCheckpointHandler) Create(c *fiber.Ctx) error {
	var req reqdto.HashChainCheckpointCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	hashChainCheckpointID, err := h.appSvc.CreateHashChainCheckpoint(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create hash chain checkpoint: "+err.Error())
	}

	// Get the created hash chain checkpoint
	hashChainCheckpointView, err := h.appSvc.GetHashChainCheckpoint(c.Context(), hashChainCheckpointID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created hash chain checkpoint: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Hash chain checkpoint created successfully", httpresp.SuccessOptions{Data: respdto.HashChainCheckpointROToDTO(&hashChainCheckpointView)})
}

func (h *HashChainCheckpointHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.HashChainCheckpointByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetHashChainCheckpoint(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Hash chain checkpoint not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Hash chain checkpoint retrieved successfully", httpresp.SuccessOptions{Data: respdto.HashChainCheckpointROToDTO(&result)})
}

func (h *HashChainCheckpointHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.HashChainCheckpointByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := hashChainCheckpointAppCommands.DeleteHashChainCheckpointCmd{HashChainCheckpointID: req.ID}
	if err := h.appSvc.DeleteHashChainCheckpoint(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete hash chain checkpoint: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Hash chain checkpoint deleted successfully")
}
