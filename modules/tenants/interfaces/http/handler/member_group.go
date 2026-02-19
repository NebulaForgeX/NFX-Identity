package handler

import (
	memberGroupApp "nfxid/modules/tenants/application/member_groups"
	memberGroupAppCommands "nfxid/modules/tenants/application/member_groups/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type MemberGroupHandler struct {
	appSvc *memberGroupApp.Service
}

func NewMemberGroupHandler(appSvc *memberGroupApp.Service) *MemberGroupHandler {
	return &MemberGroupHandler{appSvc: appSvc}
}

func (h *MemberGroupHandler) Create(c fiber.Ctx) error {
	var req reqdto.MemberGroupCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	memberGroupID, err := h.appSvc.CreateMemberGroup(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created member group
	memberGroupView, err := h.appSvc.GetMemberGroup(c.Context(), memberGroupID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Member group created successfully", httpx.SuccessOptions{Data: respdto.MemberGroupROToDTO(&memberGroupView)})
}

func (h *MemberGroupHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetMemberGroup(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Member group retrieved successfully", httpx.SuccessOptions{Data: respdto.MemberGroupROToDTO(&result)})
}

func (h *MemberGroupHandler) Update(c fiber.Ctx) error {
	// Update can be Revoke
	var req reqdto.MemberGroupRevokeRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToRevokeCmd()
	if err := h.appSvc.RevokeMemberGroup(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Member group revoked successfully")
}

func (h *MemberGroupHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := memberGroupAppCommands.DeleteMemberGroupCmd{MemberGroupID: req.ID}
	if err := h.appSvc.DeleteMemberGroup(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Member group deleted successfully")
}

// fiber:context-methods migrated
