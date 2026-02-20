package handler

import (
	permissionApp "nfxid/modules/access/application/permissions"
	permissionAppCommands "nfxid/modules/access/application/permissions/commands"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
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
func (h *PermissionHandler) Create(c fiber.Ctx) error {
	var req reqdto.PermissionCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	permissionID, err := h.appSvc.CreatePermission(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created permission
	permissionView, err := h.appSvc.GetPermission(c.Context(), permissionID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Permission created successfully", httpx.SuccessOptions{Data: respdto.PermissionROToDTO(&permissionView)})
}

// GetByID 根据 ID 获取权限
func (h *PermissionHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.PermissionByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetPermission(c.Context(), req.PermissionID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Permission retrieved successfully", httpx.SuccessOptions{Data: respdto.PermissionROToDTO(&result)})
}

// GetByKey 根据 Key 获取权限
func (h *PermissionHandler) GetByKey(c fiber.Ctx) error {
	var req reqdto.PermissionByKeyRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetPermissionByKey(c.Context(), req.Key)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Permission retrieved successfully", httpx.SuccessOptions{Data: respdto.PermissionROToDTO(&result)})
}

// Update 更新权限
func (h *PermissionHandler) Update(c fiber.Ctx) error {
	var req reqdto.PermissionUpdateRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdatePermission(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Permission updated successfully")
}

// Delete 删除权限（软删除）
func (h *PermissionHandler) Delete(c fiber.Ctx) error {
	var req reqdto.PermissionByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := permissionAppCommands.DeletePermissionCmd{PermissionID: req.PermissionID}
	if err := h.appSvc.DeletePermission(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Permission deleted successfully")
}

// fiber:context-methods migrated
