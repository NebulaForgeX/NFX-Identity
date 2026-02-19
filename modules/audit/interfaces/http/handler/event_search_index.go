package handler

import (
	eventSearchIndexApp "nfxid/modules/audit/application/event_search_index"
	eventSearchIndexAppCommands "nfxid/modules/audit/application/event_search_index/commands"
	"nfxid/modules/audit/interfaces/http/dto/reqdto"
	"nfxid/modules/audit/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type EventSearchIndexHandler struct {
	appSvc *eventSearchIndexApp.Service
}

func NewEventSearchIndexHandler(appSvc *eventSearchIndexApp.Service) *EventSearchIndexHandler {
	return &EventSearchIndexHandler{appSvc: appSvc}
}

func (h *EventSearchIndexHandler) Create(c fiber.Ctx) error {
	var req reqdto.EventSearchIndexCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	eventSearchIndexID, err := h.appSvc.CreateEventSearchIndex(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created event search index
	eventSearchIndexView, err := h.appSvc.GetEventSearchIndex(c.Context(), eventSearchIndexID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Event search index created successfully", httpx.SuccessOptions{Data: respdto.EventSearchIndexROToDTO(&eventSearchIndexView)})
}

func (h *EventSearchIndexHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.EventSearchIndexByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetEventSearchIndex(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Event search index retrieved successfully", httpx.SuccessOptions{Data: respdto.EventSearchIndexROToDTO(&result)})
}

func (h *EventSearchIndexHandler) Delete(c fiber.Ctx) error {
	var req reqdto.EventSearchIndexByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := eventSearchIndexAppCommands.DeleteEventSearchIndexCmd{EventSearchIndexID: req.ID}
	if err := h.appSvc.DeleteEventSearchIndex(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Event search index deleted successfully")
}

// fiber:context-methods migrated
