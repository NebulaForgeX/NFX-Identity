package handler

import (
	groupApp "nfxid/modules/tenants/application/groups"
	groupAppCommands "nfxid/modules/tenants/application/groups/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type GroupHandler struct {
	appSvc *groupApp.Service
}

func NewGroupHandler(appSvc *groupApp.Service) *GroupHandler {
	return &GroupHandler{appSvc: appSvc}
}

func (h *GroupHandler) Create(c *fiber.Ctx) error {
	var req reqdto.GroupCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	groupID, err := h.appSvc.CreateGroup(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create group: "+err.Error())
	}

	// Get the created group
	groupView, err := h.appSvc.GetGroup(c.Context(), groupID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created group: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Group created successfully", httpresp.SuccessOptions{Data: respdto.GroupROToDTO(&groupView)})
}

func (h *GroupHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.GroupByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetGroup(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Group not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Group retrieved successfully", httpresp.SuccessOptions{Data: respdto.GroupROToDTO(&result)})
}

func (h *GroupHandler) Update(c *fiber.Ctx) error {
	var req reqdto.GroupUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateGroup(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update group: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Group updated successfully")
}

func (h *GroupHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.GroupByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := groupAppCommands.DeleteGroupCmd{GroupID: req.ID}
	if err := h.appSvc.DeleteGroup(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete group: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Group deleted successfully")
}
