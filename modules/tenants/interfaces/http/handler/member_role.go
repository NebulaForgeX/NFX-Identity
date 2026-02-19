package handler

import (
	memberRoleApp "nfxid/modules/tenants/application/member_roles"
	memberRoleAppCommands "nfxid/modules/tenants/application/member_roles/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type MemberRoleHandler struct {
	appSvc *memberRoleApp.Service
}

func NewMemberRoleHandler(appSvc *memberRoleApp.Service) *MemberRoleHandler {
	return &MemberRoleHandler{appSvc: appSvc}
}

func (h *MemberRoleHandler) Create(c fiber.Ctx) error {
	var req reqdto.MemberRoleCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	memberRoleID, err := h.appSvc.CreateMemberRole(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created member role
	memberRoleView, err := h.appSvc.GetMemberRole(c.Context(), memberRoleID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Member role created successfully", httpx.SuccessOptions{Data: respdto.MemberRoleROToDTO(&memberRoleView)})
}

func (h *MemberRoleHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetMemberRole(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Member role retrieved successfully", httpx.SuccessOptions{Data: respdto.MemberRoleROToDTO(&result)})
}

func (h *MemberRoleHandler) Update(c fiber.Ctx) error {
	// Update can be Revoke
	var req reqdto.MemberRoleRevokeRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToRevokeCmd()
	if err := h.appSvc.RevokeMemberRole(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Member role revoked successfully")
}

func (h *MemberRoleHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := memberRoleAppCommands.DeleteMemberRoleCmd{MemberRoleID: req.ID}
	if err := h.appSvc.DeleteMemberRole(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Member role deleted successfully")
}

// fiber:context-methods migrated
