package handler

import (
	memberAppRoleApp "nfxid/modules/tenants/application/member_app_roles"
	memberAppRoleAppCommands "nfxid/modules/tenants/application/member_app_roles/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type MemberAppRoleHandler struct {
	appSvc *memberAppRoleApp.Service
}

func NewMemberAppRoleHandler(appSvc *memberAppRoleApp.Service) *MemberAppRoleHandler {
	return &MemberAppRoleHandler{appSvc: appSvc}
}

func (h *MemberAppRoleHandler) Create(c *fiber.Ctx) error {
	var req reqdto.MemberAppRoleCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	memberAppRoleID, err := h.appSvc.CreateMemberAppRole(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create member app role: "+err.Error())
	}

	// Get the created member app role
	memberAppRoleView, err := h.appSvc.GetMemberAppRole(c.Context(), memberAppRoleID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created member app role: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Member app role created successfully", httpresp.SuccessOptions{Data: respdto.MemberAppRoleROToDTO(&memberAppRoleView)})
}

func (h *MemberAppRoleHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetMemberAppRole(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Member app role not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Member app role retrieved successfully", httpresp.SuccessOptions{Data: respdto.MemberAppRoleROToDTO(&result)})
}

func (h *MemberAppRoleHandler) Update(c *fiber.Ctx) error {
	// Update can be Revoke
	var req reqdto.MemberAppRoleRevokeRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToRevokeCmd()
	if err := h.appSvc.RevokeMemberAppRole(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to revoke member app role: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Member app role revoked successfully")
}

func (h *MemberAppRoleHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := memberAppRoleAppCommands.DeleteMemberAppRoleCmd{MemberAppRoleID: req.ID}
	if err := h.appSvc.DeleteMemberAppRole(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete member app role: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Member app role deleted successfully")
}
