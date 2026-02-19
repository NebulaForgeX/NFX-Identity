package handler

import (
	invitationApp "nfxid/modules/tenants/application/invitations"
	invitationAppCommands "nfxid/modules/tenants/application/invitations/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type InvitationHandler struct {
	appSvc *invitationApp.Service
}

func NewInvitationHandler(appSvc *invitationApp.Service) *InvitationHandler {
	return &InvitationHandler{appSvc: appSvc}
}

func (h *InvitationHandler) Create(c fiber.Ctx) error {
	var req reqdto.InvitationCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	invitationID, err := h.appSvc.CreateInvitation(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created invitation
	invitationView, err := h.appSvc.GetInvitation(c.Context(), invitationID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Invitation created successfully", httpx.SuccessOptions{Data: respdto.InvitationROToDTO(&invitationView)})
}

func (h *InvitationHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.InvitationByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetInvitation(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Invitation retrieved successfully", httpx.SuccessOptions{Data: respdto.InvitationROToDTO(&result)})
}

func (h *InvitationHandler) GetByInviteID(c fiber.Ctx) error {
	var req reqdto.InvitationByInviteIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetInvitationByInviteID(c.Context(), req.InviteID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Invitation retrieved successfully", httpx.SuccessOptions{Data: respdto.InvitationROToDTO(&result)})
}

// Accept 接受邀请
func (h *InvitationHandler) Accept(c fiber.Ctx) error {
	var req reqdto.InvitationAcceptRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToAcceptCmd()
	if err := h.appSvc.AcceptInvitation(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Invitation accepted successfully")
}

// Revoke 撤销邀请
func (h *InvitationHandler) Revoke(c fiber.Ctx) error {
	var req reqdto.InvitationRevokeRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToRevokeCmd()
	if err := h.appSvc.RevokeInvitation(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Invitation revoked successfully")
}

func (h *InvitationHandler) Delete(c fiber.Ctx) error {
	// DeleteInvitation uses InviteID, but route uses :id
	// Get invitation by ID first to get InviteID
	var idReq reqdto.InvitationByIDRequestDTO
	if err := c.Bind().URI(&idReq); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	// Get invitation by ID first to get InviteID
	invitation, err := h.appSvc.GetInvitation(c.Context(), idReq.ID)
	if err != nil {
		return err
	}

	cmd := invitationAppCommands.DeleteInvitationCmd{InviteID: invitation.InviteID}
	if err := h.appSvc.DeleteInvitation(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Invitation deleted successfully")
}

// fiber:context-methods migrated
