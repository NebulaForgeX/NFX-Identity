package handler

import (
	rolePermissionApp "nfxid/modules/access/application/role_permissions"
	rolePermissionAppCommands "nfxid/modules/access/application/role_permissions/commands"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
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
func (h *RolePermissionHandler) Create(c fiber.Ctx) error {
	var req reqdto.RolePermissionCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	rolePermissionID, err := h.appSvc.CreateRolePermission(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created role permission
	rolePermissionView, err := h.appSvc.GetRolePermission(c.Context(), rolePermissionID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Role permission created successfully", httpx.SuccessOptions{Data: respdto.RolePermissionROToDTO(&rolePermissionView)})
}

// GetByID 根据 ID 获取角色权限关联
func (h *RolePermissionHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.RolePermissionByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetRolePermission(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Role permission retrieved successfully", httpx.SuccessOptions{Data: respdto.RolePermissionROToDTO(&result)})
}

// GetByRoleID 根据角色ID获取角色权限列表
func (h *RolePermissionHandler) GetByRoleID(c fiber.Ctx) error {
	var req reqdto.RolePermissionByRoleIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	results, err := h.appSvc.GetRolePermissionsByRoleID(c.Context(), req.RoleID)
	if err != nil {
		return err
	}

	dtos := make([]*respdto.RolePermissionDTO, len(results))
	for i, r := range results {
		dtos[i] = respdto.RolePermissionROToDTO(&r)
	}

	return fiberx.OK(c, "Role permissions retrieved successfully", httpx.SuccessOptions{Data: dtos})
}

// Delete 删除角色权限关联
func (h *RolePermissionHandler) Delete(c fiber.Ctx) error {
	var req reqdto.RolePermissionByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := rolePermissionAppCommands.DeleteRolePermissionCmd{RolePermissionID: req.ID}
	if err := h.appSvc.DeleteRolePermission(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Role permission deleted successfully")
}

// fiber:context-methods migrated
