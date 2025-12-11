package handler

import (
	userPermissionApp "nfxid/modules/permission/application/user_permission"
	userPermissionAppCommands "nfxid/modules/permission/application/user_permission/commands"
	"nfxid/modules/permission/interfaces/http/dto/reqdto"
	"nfxid/modules/permission/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type UserPermissionHandler struct {
	appSvc *userPermissionApp.Service
}

func NewUserPermissionHandler(appSvc *userPermissionApp.Service) *UserPermissionHandler {
	return &UserPermissionHandler{
		appSvc: appSvc,
	}
}

// Assign 分配权限给用户
func (h *UserPermissionHandler) Assign(c *fiber.Ctx) error {
	var req reqdto.UserPermissionAssignRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToAssignCmd()
	err := h.appSvc.AssignPermission(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to assign permission: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Permission assigned successfully", httpresp.SuccessOptions{})
}

// Revoke 撤销用户权限
func (h *UserPermissionHandler) Revoke(c *fiber.Ctx) error {
	var req reqdto.UserPermissionRevokeRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToRevokeCmd()
	err := h.appSvc.RevokePermission(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to revoke permission: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Permission revoked successfully", httpresp.SuccessOptions{})
}

// GetByUserID 根据用户ID获取权限列表
func (h *UserPermissionHandler) GetByUserID(c *fiber.Ctx) error {
	var req reqdto.UserPermissionByUserIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid path parameter: "+err.Error())
	}

	permissions, err := h.appSvc.GetUserPermissions(c.Context(), userPermissionAppCommands.GetUserPermissionsCmd{
		UserID: req.UserID,
	})
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get user permissions: "+err.Error())
	}

	dtos := make([]*respdto.UserPermissionDTO, len(permissions))
	for i, p := range permissions {
		dtos[i] = respdto.UserPermissionViewToDTO(p)
	}

	return httpresp.Success(c, fiber.StatusOK, "User permissions retrieved successfully", httpresp.SuccessOptions{Data: dtos})
}

// GetTagsByUserID 根据用户ID获取权限标签列表
func (h *UserPermissionHandler) GetTagsByUserID(c *fiber.Ctx) error {
	var req reqdto.UserPermissionByUserIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid path parameter: "+err.Error())
	}

	tags, err := h.appSvc.GetUserPermissionTags(c.Context(), userPermissionAppCommands.GetUserPermissionsCmd{
		UserID: req.UserID,
	})
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get user permission tags: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User permission tags retrieved successfully", httpresp.SuccessOptions{Data: tags})
}

// Check 检查用户是否拥有指定权限
func (h *UserPermissionHandler) Check(c *fiber.Ctx) error {
	var req reqdto.UserPermissionCheckRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCheckCmd()
	hasPermission, err := h.appSvc.CheckPermission(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to check permission: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Permission check completed", httpresp.SuccessOptions{Data: map[string]bool{"has_permission": hasPermission}})
}

