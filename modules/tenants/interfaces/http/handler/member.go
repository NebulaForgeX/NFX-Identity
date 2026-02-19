package handler

import (
	memberApp "nfxid/modules/tenants/application/members"
	memberAppCommands "nfxid/modules/tenants/application/members/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type MemberHandler struct {
	appSvc *memberApp.Service
}

func NewMemberHandler(appSvc *memberApp.Service) *MemberHandler {
	return &MemberHandler{appSvc: appSvc}
}

func (h *MemberHandler) Create(c fiber.Ctx) error {
	var req reqdto.MemberCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	memberID, err := h.appSvc.CreateMember(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created member
	memberView, err := h.appSvc.GetMember(c.Context(), memberID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Member created successfully", httpx.SuccessOptions{Data: respdto.MemberROToDTO(&memberView)})
}

func (h *MemberHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.MemberByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetMember(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Member retrieved successfully", httpx.SuccessOptions{Data: respdto.MemberROToDTO(&result)})
}

func (h *MemberHandler) Update(c fiber.Ctx) error {
	var req reqdto.MemberUpdateStatusRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateStatusCmd()
	if err := h.appSvc.UpdateMemberStatus(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Member status updated successfully")
}

func (h *MemberHandler) Delete(c fiber.Ctx) error {
	var req reqdto.MemberByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := memberAppCommands.DeleteMemberCmd{MemberID: req.ID}
	if err := h.appSvc.DeleteMember(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Member deleted successfully")
}

// fiber:context-methods migrated
