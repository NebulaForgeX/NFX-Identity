package handler

import (
	eventRetentionPolicyApp "nfxid/modules/audit/application/event_retention_policies"
	eventRetentionPolicyAppCommands "nfxid/modules/audit/application/event_retention_policies/commands"
	"nfxid/modules/audit/interfaces/http/dto/reqdto"
	"nfxid/modules/audit/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type EventRetentionPolicyHandler struct {
	appSvc *eventRetentionPolicyApp.Service
}

func NewEventRetentionPolicyHandler(appSvc *eventRetentionPolicyApp.Service) *EventRetentionPolicyHandler {
	return &EventRetentionPolicyHandler{appSvc: appSvc}
}

func (h *EventRetentionPolicyHandler) Create(c fiber.Ctx) error {
	var req reqdto.EventRetentionPolicyCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	eventRetentionPolicyID, err := h.appSvc.CreateEventRetentionPolicy(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created event retention policy
	eventRetentionPolicyView, err := h.appSvc.GetEventRetentionPolicy(c.Context(), eventRetentionPolicyID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Event retention policy created successfully", httpx.SuccessOptions{Data: respdto.EventRetentionPolicyROToDTO(&eventRetentionPolicyView)})
}

func (h *EventRetentionPolicyHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.EventRetentionPolicyByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetEventRetentionPolicy(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Event retention policy retrieved successfully", httpx.SuccessOptions{Data: respdto.EventRetentionPolicyROToDTO(&result)})
}

func (h *EventRetentionPolicyHandler) Update(c fiber.Ctx) error {
	var req reqdto.EventRetentionPolicyUpdateRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateEventRetentionPolicy(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Event retention policy updated successfully")
}

func (h *EventRetentionPolicyHandler) Delete(c fiber.Ctx) error {
	var req reqdto.EventRetentionPolicyByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := eventRetentionPolicyAppCommands.DeleteEventRetentionPolicyCmd{EventRetentionPolicyID: req.ID}
	if err := h.appSvc.DeleteEventRetentionPolicy(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Event retention policy deleted successfully")
}

// fiber:context-methods migrated
