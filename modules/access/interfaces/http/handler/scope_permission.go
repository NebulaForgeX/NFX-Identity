package handler

import (
	scopePermissionApp "nfxid/modules/access/application/scope_permissions"
	scopePermissionAppCommands "nfxid/modules/access/application/scope_permissions/commands"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type ScopePermissionHandler struct {
	appSvc *scopePermissionApp.Service
}

func NewScopePermissionHandler(appSvc *scopePermissionApp.Service) *ScopePermissionHandler {
	return &ScopePermissionHandler{
		appSvc: appSvc,
	}
}

// Create 创建作用域权限关联
func (h *ScopePermissionHandler) Create(c fiber.Ctx) error {
	var req reqdto.ScopePermissionCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	scopePermissionID, err := h.appSvc.CreateScopePermission(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created scope permission
	scopePermissionView, err := h.appSvc.GetScopePermission(c.Context(), scopePermissionID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Scope permission created successfully", httpx.SuccessOptions{Data: respdto.ScopePermissionROToDTO(&scopePermissionView)})
}

// GetByID 根据 ID 获取作用域权限关联
func (h *ScopePermissionHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ScopePermissionByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetScopePermission(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Scope permission retrieved successfully", httpx.SuccessOptions{Data: respdto.ScopePermissionROToDTO(&result)})
}

// Delete 删除作用域权限关联
func (h *ScopePermissionHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ScopePermissionByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := scopePermissionAppCommands.DeleteScopePermissionCmd{ScopePermissionID: req.ID}
	if err := h.appSvc.DeleteScopePermission(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Scope permission deleted successfully")
}

// fiber:context-methods migrated
