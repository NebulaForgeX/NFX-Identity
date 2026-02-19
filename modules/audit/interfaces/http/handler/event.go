package handler

import (
	eventApp "nfxid/modules/audit/application/events"
	eventAppCommands "nfxid/modules/audit/application/events/commands"
	"nfxid/modules/audit/interfaces/http/dto/reqdto"
	"nfxid/modules/audit/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type EventHandler struct {
	appSvc *eventApp.Service
}

func NewEventHandler(appSvc *eventApp.Service) *EventHandler {
	return &EventHandler{appSvc: appSvc}
}

// Create 创建事件
func (h *EventHandler) Create(c fiber.Ctx) error {
	var req reqdto.EventCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	eventID, err := h.appSvc.CreateEvent(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created event
	eventView, err := h.appSvc.GetEvent(c.Context(), eventID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Event created successfully", httpx.SuccessOptions{Data: respdto.EventROToDTO(&eventView)})
}

// GetByID 根据 ID 获取事件
func (h *EventHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.EventByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetEvent(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Event retrieved successfully", httpx.SuccessOptions{Data: respdto.EventROToDTO(&result)})
}

// Delete 删除事件
func (h *EventHandler) Delete(c fiber.Ctx) error {
	var req reqdto.EventByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := eventAppCommands.DeleteEventCmd{EventID: req.ID}
	if err := h.appSvc.DeleteEvent(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Event deleted successfully")
}

// fiber:context-methods migrated
