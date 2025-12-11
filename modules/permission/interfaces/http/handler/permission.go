package handler

import (
	permissionApp "nfxid/modules/permission/application/permission"
	permissionAppCommands "nfxid/modules/permission/application/permission/commands"
	"nfxid/modules/permission/interfaces/http/dto/reqdto"
	"nfxid/modules/permission/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type PermissionHandler struct {
	appSvc *permissionApp.Service
}

func NewPermissionHandler(appSvc *permissionApp.Service) *PermissionHandler {
	return &PermissionHandler{
		appSvc: appSvc,
	}
}

// Create 创建权限
func (h *PermissionHandler) Create(c *fiber.Ctx) error {
	var req reqdto.PermissionCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	result, err := h.appSvc.CreatePermission(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create permission: "+err.Error())
	}

	// Get the created permission view
	permissionView, err := h.appSvc.GetPermission(c.Context(), permissionAppCommands.GetPermissionCmd{
		ID: result.ID(),
	})
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created permission: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Permission created successfully", httpresp.SuccessOptions{Data: respdto.PermissionViewToDTO(permissionView)})
}

// Update 更新权限
func (h *PermissionHandler) Update(c *fiber.Ctx) error {
	var req reqdto.PermissionUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid path parameter: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	err := h.appSvc.UpdatePermission(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update permission: "+err.Error())
	}

	// Get the updated permission view
	permissionView, err := h.appSvc.GetPermission(c.Context(), permissionAppCommands.GetPermissionCmd{
		ID: req.ID,
	})
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get updated permission: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Permission updated successfully", httpresp.SuccessOptions{Data: respdto.PermissionViewToDTO(permissionView)})
}

// Delete 删除权限
func (h *PermissionHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.PermissionByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid path parameter: "+err.Error())
	}

	err := h.appSvc.DeletePermission(c.Context(), permissionAppCommands.DeletePermissionCmd{
		ID: req.ID,
	})
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete permission: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Permission deleted successfully", httpresp.SuccessOptions{})
}

// GetByID 根据ID获取权限
func (h *PermissionHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.PermissionByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid path parameter: "+err.Error())
	}

	permissionView, err := h.appSvc.GetPermission(c.Context(), permissionAppCommands.GetPermissionCmd{
		ID: req.ID,
	})
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Permission not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Permission retrieved successfully", httpresp.SuccessOptions{Data: respdto.PermissionViewToDTO(permissionView)})
}

// GetByTag 根据Tag获取权限
func (h *PermissionHandler) GetByTag(c *fiber.Ctx) error {
	var req reqdto.PermissionByTagRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid path parameter: "+err.Error())
	}

	permissionView, err := h.appSvc.GetPermissionByTag(c.Context(), permissionAppCommands.GetPermissionByTagCmd{
		Tag: req.Tag,
	})
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Permission not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Permission retrieved successfully", httpresp.SuccessOptions{Data: respdto.PermissionViewToDTO(permissionView)})
}

// List 获取权限列表
func (h *PermissionHandler) List(c *fiber.Ctx) error {
	var query reqdto.PermissionQueryParamsDTO
	if err := c.QueryParser(&query); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid query parameters: "+err.Error())
	}

	cmd := permissionAppCommands.ListPermissionsCmd{}
	if query.Category != nil {
		cmd.Category = *query.Category
	}

	permissions, err := h.appSvc.ListPermissions(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to list permissions: "+err.Error())
	}

	dtos := make([]*respdto.PermissionDTO, len(permissions))
	for i, p := range permissions {
		dtos[i] = respdto.PermissionViewToDTO(p)
	}

	return httpresp.Success(c, fiber.StatusOK, "Permissions retrieved successfully", httpresp.SuccessOptions{Data: dtos})
}

