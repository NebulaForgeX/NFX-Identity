package handler

import (
	groupApp "nfxid/modules/tenants/application/groups"
	groupAppCommands "nfxid/modules/tenants/application/groups/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type GroupHandler struct {
	appSvc *groupApp.Service
}

func NewGroupHandler(appSvc *groupApp.Service) *GroupHandler {
	return &GroupHandler{appSvc: appSvc}
}

func (h *GroupHandler) Create(c fiber.Ctx) error {
	var req reqdto.GroupCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	groupID, err := h.appSvc.CreateGroup(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created group
	groupView, err := h.appSvc.GetGroup(c.Context(), groupID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Group created successfully", httpx.SuccessOptions{Data: respdto.GroupROToDTO(&groupView)})
}

func (h *GroupHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.GroupByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetGroup(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Group retrieved successfully", httpx.SuccessOptions{Data: respdto.GroupROToDTO(&result)})
}

func (h *GroupHandler) Update(c fiber.Ctx) error {
	var req reqdto.GroupUpdateRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateGroup(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Group updated successfully")
}

func (h *GroupHandler) Delete(c fiber.Ctx) error {
	var req reqdto.GroupByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := groupAppCommands.DeleteGroupCmd{GroupID: req.ID}
	if err := h.appSvc.DeleteGroup(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Group deleted successfully")
}

// fiber:context-methods migrated
