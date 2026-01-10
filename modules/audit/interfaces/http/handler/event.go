package handler

import (
	eventApp "nfxid/modules/audit/application/events"
	eventAppCommands "nfxid/modules/audit/application/events/commands"
	"nfxid/modules/audit/interfaces/http/dto/reqdto"
	"nfxid/modules/audit/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	appSvc *eventApp.Service
}

func NewEventHandler(appSvc *eventApp.Service) *EventHandler {
	return &EventHandler{appSvc: appSvc}
}

// Create 创建事件
func (h *EventHandler) Create(c *fiber.Ctx) error {
	var req reqdto.EventCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	eventID, err := h.appSvc.CreateEvent(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create event: "+err.Error())
	}

	// Get the created event
	eventView, err := h.appSvc.GetEvent(c.Context(), eventID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created event: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Event created successfully", httpresp.SuccessOptions{Data: respdto.EventROToDTO(&eventView)})
}

// GetByID 根据 ID 获取事件
func (h *EventHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.EventByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetEvent(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Event not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Event retrieved successfully", httpresp.SuccessOptions{Data: respdto.EventROToDTO(&result)})
}

// Delete 删除事件
func (h *EventHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.EventByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := eventAppCommands.DeleteEventCmd{EventID: req.ID}
	if err := h.appSvc.DeleteEvent(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete event: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Event deleted successfully")
}
