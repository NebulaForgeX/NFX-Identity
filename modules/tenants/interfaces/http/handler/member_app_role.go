package handler

import (
	memberAppRoleApp "nfxid/modules/tenants/application/member_app_roles"
	memberAppRoleAppCommands "nfxid/modules/tenants/application/member_app_roles/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type MemberAppRoleHandler struct {
	appSvc *memberAppRoleApp.Service
}

func NewMemberAppRoleHandler(appSvc *memberAppRoleApp.Service) *MemberAppRoleHandler {
	return &MemberAppRoleHandler{appSvc: appSvc}
}

func (h *MemberAppRoleHandler) Create(c fiber.Ctx) error {
	var req reqdto.MemberAppRoleCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	memberAppRoleID, err := h.appSvc.CreateMemberAppRole(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created member app role
	memberAppRoleView, err := h.appSvc.GetMemberAppRole(c.Context(), memberAppRoleID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Member app role created successfully", httpx.SuccessOptions{Data: respdto.MemberAppRoleROToDTO(&memberAppRoleView)})
}

func (h *MemberAppRoleHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetMemberAppRole(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Member app role retrieved successfully", httpx.SuccessOptions{Data: respdto.MemberAppRoleROToDTO(&result)})
}

func (h *MemberAppRoleHandler) Update(c fiber.Ctx) error {
	// Update can be Revoke
	var req reqdto.MemberAppRoleRevokeRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToRevokeCmd()
	if err := h.appSvc.RevokeMemberAppRole(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Member app role revoked successfully")
}

func (h *MemberAppRoleHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := memberAppRoleAppCommands.DeleteMemberAppRoleCmd{MemberAppRoleID: req.ID}
	if err := h.appSvc.DeleteMemberAppRole(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Member app role deleted successfully")
}

// fiber:context-methods migrated
