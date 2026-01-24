package handler

import (
	rolePermissionApp "nfxid/modules/access/application/role_permissions"
	rolePermissionAppCommands "nfxid/modules/access/application/role_permissions/commands"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type RolePermissionHandler struct {
	appSvc *rolePermissionApp.Service
}

func NewRolePermissionHandler(appSvc *rolePermissionApp.Service) *RolePermissionHandler {
	return &RolePermissionHandler{
		appSvc: appSvc,
	}
}

// Create 创建角色权限关联
func (h *RolePermissionHandler) Create(c *fiber.Ctx) error {
	var req reqdto.RolePermissionCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	rolePermissionID, err := h.appSvc.CreateRolePermission(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create role permission: "+err.Error())
	}

	// Get the created role permission
	rolePermissionView, err := h.appSvc.GetRolePermission(c.Context(), rolePermissionID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created role permission: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Role permission created successfully", httpresp.SuccessOptions{Data: respdto.RolePermissionROToDTO(&rolePermissionView)})
}

// GetByID 根据 ID 获取角色权限关联
func (h *RolePermissionHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.RolePermissionByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetRolePermission(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Role permission not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Role permission retrieved successfully", httpresp.SuccessOptions{Data: respdto.RolePermissionROToDTO(&result)})
}

// GetByRoleID 根据角色ID获取角色权限列表
func (h *RolePermissionHandler) GetByRoleID(c *fiber.Ctx) error {
	var req reqdto.RolePermissionByRoleIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	results, err := h.appSvc.GetRolePermissionsByRoleID(c.Context(), req.RoleID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Role permissions not found: "+err.Error())
	}

	dtos := make([]*respdto.RolePermissionDTO, len(results))
	for i, r := range results {
		dtos[i] = respdto.RolePermissionROToDTO(&r)
	}

	return httpresp.Success(c, fiber.StatusOK, "Role permissions retrieved successfully", httpresp.SuccessOptions{Data: dtos})
}

// Delete 删除角色权限关联
func (h *RolePermissionHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.RolePermissionByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := rolePermissionAppCommands.DeleteRolePermissionCmd{RolePermissionID: req.ID}
	if err := h.appSvc.DeleteRolePermission(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete role permission: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Role permission deleted successfully")
}
