package handler

import (
	permissionApp "nfxid/modules/access/application/permissions"
	permissionAppCommands "nfxid/modules/access/application/permissions/commands"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
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
	permissionID, err := h.appSvc.CreatePermission(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create permission: "+err.Error())
	}

	// Get the created permission
	permissionView, err := h.appSvc.GetPermission(c.Context(), permissionID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created permission: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Permission created successfully", httpresp.SuccessOptions{Data: respdto.PermissionROToDTO(&permissionView)})
}

// GetByID 根据 ID 获取权限
func (h *PermissionHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.PermissionByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetPermission(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Permission not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Permission retrieved successfully", httpresp.SuccessOptions{Data: respdto.PermissionROToDTO(&result)})
}

// GetByKey 根据 Key 获取权限
func (h *PermissionHandler) GetByKey(c *fiber.Ctx) error {
	var req reqdto.PermissionByKeyRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetPermissionByKey(c.Context(), req.Key)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Permission not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Permission retrieved successfully", httpresp.SuccessOptions{Data: respdto.PermissionROToDTO(&result)})
}

// Update 更新权限
func (h *PermissionHandler) Update(c *fiber.Ctx) error {
	var req reqdto.PermissionUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdatePermission(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update permission: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Permission updated successfully")
}

// Delete 删除权限（软删除）
func (h *PermissionHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.PermissionByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := permissionAppCommands.DeletePermissionCmd{PermissionID: req.ID}
	if err := h.appSvc.DeletePermission(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete permission: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Permission deleted successfully")
}
