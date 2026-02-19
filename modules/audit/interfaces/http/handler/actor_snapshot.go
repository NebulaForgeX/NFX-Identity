package handler

import (
	actorSnapshotApp "nfxid/modules/audit/application/actor_snapshots"
	actorSnapshotAppCommands "nfxid/modules/audit/application/actor_snapshots/commands"
	"nfxid/modules/audit/interfaces/http/dto/reqdto"
	"nfxid/modules/audit/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type ActorSnapshotHandler struct {
	appSvc *actorSnapshotApp.Service
}

func NewActorSnapshotHandler(appSvc *actorSnapshotApp.Service) *ActorSnapshotHandler {
	return &ActorSnapshotHandler{appSvc: appSvc}
}

func (h *ActorSnapshotHandler) Create(c fiber.Ctx) error {
	var req reqdto.ActorSnapshotCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	actorSnapshotID, err := h.appSvc.CreateActorSnapshot(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created actor snapshot
	actorSnapshotView, err := h.appSvc.GetActorSnapshot(c.Context(), actorSnapshotID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Actor snapshot created successfully", httpx.SuccessOptions{Data: respdto.ActorSnapshotROToDTO(&actorSnapshotView)})
}

func (h *ActorSnapshotHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ActorSnapshotByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetActorSnapshot(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Actor snapshot retrieved successfully", httpx.SuccessOptions{Data: respdto.ActorSnapshotROToDTO(&result)})
}

func (h *ActorSnapshotHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ActorSnapshotByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := actorSnapshotAppCommands.DeleteActorSnapshotCmd{ActorSnapshotID: req.ID}
	if err := h.appSvc.DeleteActorSnapshot(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Actor snapshot deleted successfully")
}

// fiber:context-methods migrated
