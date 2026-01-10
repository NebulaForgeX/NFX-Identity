package handler

import (
	eventRetentionPolicyApp "nfxid/modules/audit/application/event_retention_policies"
	eventRetentionPolicyAppCommands "nfxid/modules/audit/application/event_retention_policies/commands"
	"nfxid/modules/audit/interfaces/http/dto/reqdto"
	"nfxid/modules/audit/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type EventRetentionPolicyHandler struct {
	appSvc *eventRetentionPolicyApp.Service
}

func NewEventRetentionPolicyHandler(appSvc *eventRetentionPolicyApp.Service) *EventRetentionPolicyHandler {
	return &EventRetentionPolicyHandler{appSvc: appSvc}
}

func (h *EventRetentionPolicyHandler) Create(c *fiber.Ctx) error {
	var req reqdto.EventRetentionPolicyCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	eventRetentionPolicyID, err := h.appSvc.CreateEventRetentionPolicy(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create event retention policy: "+err.Error())
	}

	// Get the created event retention policy
	eventRetentionPolicyView, err := h.appSvc.GetEventRetentionPolicy(c.Context(), eventRetentionPolicyID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created event retention policy: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Event retention policy created successfully", httpresp.SuccessOptions{Data: respdto.EventRetentionPolicyROToDTO(&eventRetentionPolicyView)})
}

func (h *EventRetentionPolicyHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.EventRetentionPolicyByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetEventRetentionPolicy(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Event retention policy not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Event retention policy retrieved successfully", httpresp.SuccessOptions{Data: respdto.EventRetentionPolicyROToDTO(&result)})
}

func (h *EventRetentionPolicyHandler) Update(c *fiber.Ctx) error {
	var req reqdto.EventRetentionPolicyUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateEventRetentionPolicy(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update event retention policy: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Event retention policy updated successfully")
}

func (h *EventRetentionPolicyHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.EventRetentionPolicyByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := eventRetentionPolicyAppCommands.DeleteEventRetentionPolicyCmd{EventRetentionPolicyID: req.ID}
	if err := h.appSvc.DeleteEventRetentionPolicy(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete event retention policy: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Event retention policy deleted successfully")
}
