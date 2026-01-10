package handler

import (
	actorSnapshotApp "nfxid/modules/audit/application/actor_snapshots"
	actorSnapshotAppCommands "nfxid/modules/audit/application/actor_snapshots/commands"
	"nfxid/modules/audit/interfaces/http/dto/reqdto"
	"nfxid/modules/audit/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type ActorSnapshotHandler struct {
	appSvc *actorSnapshotApp.Service
}

func NewActorSnapshotHandler(appSvc *actorSnapshotApp.Service) *ActorSnapshotHandler {
	return &ActorSnapshotHandler{appSvc: appSvc}
}

func (h *ActorSnapshotHandler) Create(c *fiber.Ctx) error {
	var req reqdto.ActorSnapshotCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	actorSnapshotID, err := h.appSvc.CreateActorSnapshot(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create actor snapshot: "+err.Error())
	}

	// Get the created actor snapshot
	actorSnapshotView, err := h.appSvc.GetActorSnapshot(c.Context(), actorSnapshotID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created actor snapshot: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Actor snapshot created successfully", httpresp.SuccessOptions{Data: respdto.ActorSnapshotROToDTO(&actorSnapshotView)})
}

func (h *ActorSnapshotHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ActorSnapshotByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetActorSnapshot(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Actor snapshot not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Actor snapshot retrieved successfully", httpresp.SuccessOptions{Data: respdto.ActorSnapshotROToDTO(&result)})
}

func (h *ActorSnapshotHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ActorSnapshotByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := actorSnapshotAppCommands.DeleteActorSnapshotCmd{ActorSnapshotID: req.ID}
	if err := h.appSvc.DeleteActorSnapshot(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete actor snapshot: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Actor snapshot deleted successfully")
}
