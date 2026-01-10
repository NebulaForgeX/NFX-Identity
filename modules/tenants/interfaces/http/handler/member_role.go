package handler

import (
	memberRoleApp "nfxid/modules/tenants/application/member_roles"
	memberRoleAppCommands "nfxid/modules/tenants/application/member_roles/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type MemberRoleHandler struct {
	appSvc *memberRoleApp.Service
}

func NewMemberRoleHandler(appSvc *memberRoleApp.Service) *MemberRoleHandler {
	return &MemberRoleHandler{appSvc: appSvc}
}

func (h *MemberRoleHandler) Create(c *fiber.Ctx) error {
	var req reqdto.MemberRoleCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	memberRoleID, err := h.appSvc.CreateMemberRole(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create member role: "+err.Error())
	}

	// Get the created member role
	memberRoleView, err := h.appSvc.GetMemberRole(c.Context(), memberRoleID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created member role: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Member role created successfully", httpresp.SuccessOptions{Data: respdto.MemberRoleROToDTO(&memberRoleView)})
}

func (h *MemberRoleHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetMemberRole(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Member role not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Member role retrieved successfully", httpresp.SuccessOptions{Data: respdto.MemberRoleROToDTO(&result)})
}

func (h *MemberRoleHandler) Update(c *fiber.Ctx) error {
	// Update can be Revoke
	var req reqdto.MemberRoleRevokeRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToRevokeCmd()
	if err := h.appSvc.RevokeMemberRole(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to revoke member role: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Member role revoked successfully")
}

func (h *MemberRoleHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := memberRoleAppCommands.DeleteMemberRoleCmd{MemberRoleID: req.ID}
	if err := h.appSvc.DeleteMemberRole(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete member role: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Member role deleted successfully")
}
