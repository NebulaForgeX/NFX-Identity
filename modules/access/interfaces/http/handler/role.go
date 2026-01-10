package handler

import (
	roleApp "nfxid/modules/access/application/roles"
	roleAppCommands "nfxid/modules/access/application/roles/commands"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	appSvc *roleApp.Service
}

func NewRoleHandler(appSvc *roleApp.Service) *RoleHandler {
	return &RoleHandler{
		appSvc: appSvc,
	}
}

// Create 创建角色
func (h *RoleHandler) Create(c *fiber.Ctx) error {
	var req reqdto.RoleCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	roleID, err := h.appSvc.CreateRole(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create role: "+err.Error())
	}

	// Get the created role
	roleView, err := h.appSvc.GetRole(c.Context(), roleID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created role: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Role created successfully", httpresp.SuccessOptions{Data: respdto.RoleROToDTO(&roleView)})
}

// GetByID 根据 ID 获取角色
func (h *RoleHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.RoleByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetRole(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Role not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Role retrieved successfully", httpresp.SuccessOptions{Data: respdto.RoleROToDTO(&result)})
}

// GetByKey 根据 Key 获取角色
func (h *RoleHandler) GetByKey(c *fiber.Ctx) error {
	var req reqdto.RoleByKeyRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetRoleByKey(c.Context(), req.Key)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Role not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Role retrieved successfully", httpresp.SuccessOptions{Data: respdto.RoleROToDTO(&result)})
}

// Update 更新角色
func (h *RoleHandler) Update(c *fiber.Ctx) error {
	var req reqdto.RoleUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateRole(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update role: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Role updated successfully")
}

// Delete 删除角色（软删除）
func (h *RoleHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.RoleByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := roleAppCommands.DeleteRoleCmd{RoleID: req.ID}
	if err := h.appSvc.DeleteRole(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete role: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Role deleted successfully")
}
