package handler

import (
	memberGroupApp "nfxid/modules/tenants/application/member_groups"
	memberGroupAppCommands "nfxid/modules/tenants/application/member_groups/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type MemberGroupHandler struct {
	appSvc *memberGroupApp.Service
}

func NewMemberGroupHandler(appSvc *memberGroupApp.Service) *MemberGroupHandler {
	return &MemberGroupHandler{appSvc: appSvc}
}

func (h *MemberGroupHandler) Create(c *fiber.Ctx) error {
	var req reqdto.MemberGroupCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	memberGroupID, err := h.appSvc.CreateMemberGroup(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create member group: "+err.Error())
	}

	// Get the created member group
	memberGroupView, err := h.appSvc.GetMemberGroup(c.Context(), memberGroupID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created member group: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Member group created successfully", httpresp.SuccessOptions{Data: respdto.MemberGroupROToDTO(&memberGroupView)})
}

func (h *MemberGroupHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetMemberGroup(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Member group not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Member group retrieved successfully", httpresp.SuccessOptions{Data: respdto.MemberGroupROToDTO(&result)})
}

func (h *MemberGroupHandler) Update(c *fiber.Ctx) error {
	// Update can be Revoke
	var req reqdto.MemberGroupRevokeRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToRevokeCmd()
	if err := h.appSvc.RevokeMemberGroup(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to revoke member group: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Member group revoked successfully")
}

func (h *MemberGroupHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := memberGroupAppCommands.DeleteMemberGroupCmd{MemberGroupID: req.ID}
	if err := h.appSvc.DeleteMemberGroup(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete member group: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Member group deleted successfully")
}
