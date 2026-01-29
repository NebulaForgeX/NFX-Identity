package handler

import (
	arApp "nfxid/modules/access/application/action_requirements"
	arAppCommands "nfxid/modules/access/application/action_requirements/commands"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type ActionRequirementHandler struct {
	appSvc *arApp.Service
}

func NewActionRequirementHandler(appSvc *arApp.Service) *ActionRequirementHandler {
	return &ActionRequirementHandler{appSvc: appSvc}
}

// GetByID 根据 ID 获取 ActionRequirement
func (h *ActionRequirementHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ActionRequirementByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	result, err := h.appSvc.GetActionRequirement(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Action requirement not found: "+err.Error())
	}
	return httpresp.Success(c, fiber.StatusOK, "Action requirement retrieved successfully", httpresp.SuccessOptions{Data: respdto.ActionRequirementROToDTO(&result)})
}

// GetByPermissionID 根据 PermissionID 获取 ActionRequirement 列表
func (h *ActionRequirementHandler) GetByPermissionID(c *fiber.Ctx) error {
	var req reqdto.ActionRequirementByPermissionIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	list, err := h.appSvc.GetByPermissionID(c.Context(), req.PermissionID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get action requirements: "+err.Error())
	}
	return httpresp.Success(c, fiber.StatusOK, "Action requirements retrieved successfully", httpresp.SuccessOptions{Data: respdto.ActionRequirementListROToDTO(list)})
}

// Create 创建 ActionRequirement
func (h *ActionRequirementHandler) Create(c *fiber.Ctx) error {
	var req reqdto.ActionRequirementCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}
	cmd, err := req.ToCreateCmd()
	if err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid action_id or permission_id: "+err.Error())
	}
	arID, err := h.appSvc.CreateActionRequirement(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create action requirement: "+err.Error())
	}
	result, err := h.appSvc.GetActionRequirement(c.Context(), arID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created action requirement: "+err.Error())
	}
	return httpresp.Success(c, fiber.StatusCreated, "Action requirement created successfully", httpresp.SuccessOptions{Data: respdto.ActionRequirementROToDTO(&result)})
}

// Delete 删除 ActionRequirement
func (h *ActionRequirementHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ActionRequirementDeleteRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	err := h.appSvc.DeleteActionRequirement(c.Context(), arAppCommands.DeleteActionRequirementCmd{ActionRequirementID: req.ID})
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete action requirement: "+err.Error())
	}
	return httpresp.Success(c, fiber.StatusOK, "Action requirement deleted successfully", httpresp.SuccessOptions{})
}
