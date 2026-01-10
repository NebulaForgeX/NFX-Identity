package handler

import (
	memberApp "nfxid/modules/tenants/application/members"
	memberAppCommands "nfxid/modules/tenants/application/members/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type MemberHandler struct {
	appSvc *memberApp.Service
}

func NewMemberHandler(appSvc *memberApp.Service) *MemberHandler {
	return &MemberHandler{appSvc: appSvc}
}

func (h *MemberHandler) Create(c *fiber.Ctx) error {
	var req reqdto.MemberCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	memberID, err := h.appSvc.CreateMember(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create member: "+err.Error())
	}

	// Get the created member
	memberView, err := h.appSvc.GetMember(c.Context(), memberID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created member: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Member created successfully", httpresp.SuccessOptions{Data: respdto.MemberROToDTO(&memberView)})
}

func (h *MemberHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.MemberByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetMember(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Member not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Member retrieved successfully", httpresp.SuccessOptions{Data: respdto.MemberROToDTO(&result)})
}

func (h *MemberHandler) Update(c *fiber.Ctx) error {
	var req reqdto.MemberUpdateStatusRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateStatusCmd()
	if err := h.appSvc.UpdateMemberStatus(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update member status: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Member status updated successfully")
}

func (h *MemberHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.MemberByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := memberAppCommands.DeleteMemberCmd{MemberID: req.ID}
	if err := h.appSvc.DeleteMember(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete member: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Member deleted successfully")
}
