package handler

import (
	eventSearchIndexApp "nfxid/modules/audit/application/event_search_index"
	eventSearchIndexAppCommands "nfxid/modules/audit/application/event_search_index/commands"
	"nfxid/modules/audit/interfaces/http/dto/reqdto"
	"nfxid/modules/audit/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type EventSearchIndexHandler struct {
	appSvc *eventSearchIndexApp.Service
}

func NewEventSearchIndexHandler(appSvc *eventSearchIndexApp.Service) *EventSearchIndexHandler {
	return &EventSearchIndexHandler{appSvc: appSvc}
}

func (h *EventSearchIndexHandler) Create(c *fiber.Ctx) error {
	var req reqdto.EventSearchIndexCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	eventSearchIndexID, err := h.appSvc.CreateEventSearchIndex(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create event search index: "+err.Error())
	}

	// Get the created event search index
	eventSearchIndexView, err := h.appSvc.GetEventSearchIndex(c.Context(), eventSearchIndexID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created event search index: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Event search index created successfully", httpresp.SuccessOptions{Data: respdto.EventSearchIndexROToDTO(&eventSearchIndexView)})
}

func (h *EventSearchIndexHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.EventSearchIndexByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetEventSearchIndex(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Event search index not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Event search index retrieved successfully", httpresp.SuccessOptions{Data: respdto.EventSearchIndexROToDTO(&result)})
}

func (h *EventSearchIndexHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.EventSearchIndexByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := eventSearchIndexAppCommands.DeleteEventSearchIndexCmd{EventSearchIndexID: req.ID}
	if err := h.appSvc.DeleteEventSearchIndex(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete event search index: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Event search index deleted successfully")
}
