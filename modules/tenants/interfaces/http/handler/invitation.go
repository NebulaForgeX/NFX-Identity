package handler

import (
	invitationApp "nfxid/modules/tenants/application/invitations"
	invitationAppCommands "nfxid/modules/tenants/application/invitations/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type InvitationHandler struct {
	appSvc *invitationApp.Service
}

func NewInvitationHandler(appSvc *invitationApp.Service) *InvitationHandler {
	return &InvitationHandler{appSvc: appSvc}
}

func (h *InvitationHandler) Create(c *fiber.Ctx) error {
	var req reqdto.InvitationCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	invitationID, err := h.appSvc.CreateInvitation(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create invitation: "+err.Error())
	}

	// Get the created invitation
	invitationView, err := h.appSvc.GetInvitation(c.Context(), invitationID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created invitation: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Invitation created successfully", httpresp.SuccessOptions{Data: respdto.InvitationROToDTO(&invitationView)})
}

func (h *InvitationHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.InvitationByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetInvitation(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Invitation not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Invitation retrieved successfully", httpresp.SuccessOptions{Data: respdto.InvitationROToDTO(&result)})
}

func (h *InvitationHandler) GetByInviteID(c *fiber.Ctx) error {
	var req reqdto.InvitationByInviteIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetInvitationByInviteID(c.Context(), req.InviteID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Invitation not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Invitation retrieved successfully", httpresp.SuccessOptions{Data: respdto.InvitationROToDTO(&result)})
}


// Accept 接受邀请
func (h *InvitationHandler) Accept(c *fiber.Ctx) error {
	var req reqdto.InvitationAcceptRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToAcceptCmd()
	if err := h.appSvc.AcceptInvitation(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to accept invitation: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Invitation accepted successfully")
}

// Revoke 撤销邀请
func (h *InvitationHandler) Revoke(c *fiber.Ctx) error {
	var req reqdto.InvitationRevokeRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToRevokeCmd()
	if err := h.appSvc.RevokeInvitation(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to revoke invitation: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Invitation revoked successfully")
}

func (h *InvitationHandler) Delete(c *fiber.Ctx) error {
	// DeleteInvitation uses InviteID, but route uses :id
	// Get invitation by ID first to get InviteID
	var idReq reqdto.InvitationByIDRequestDTO
	if err := c.ParamsParser(&idReq); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	// Get invitation by ID first to get InviteID
	invitation, err := h.appSvc.GetInvitation(c.Context(), idReq.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Invitation not found: "+err.Error())
	}

	cmd := invitationAppCommands.DeleteInvitationCmd{InviteID: invitation.InviteID}
	if err := h.appSvc.DeleteInvitation(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete invitation: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Invitation deleted successfully")
}
