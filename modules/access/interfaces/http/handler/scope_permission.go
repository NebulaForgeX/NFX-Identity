package handler

import (
	scopePermissionApp "nfxid/modules/access/application/scope_permissions"
	scopePermissionAppCommands "nfxid/modules/access/application/scope_permissions/commands"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
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
func (h *ScopePermissionHandler) Create(c *fiber.Ctx) error {
	var req reqdto.ScopePermissionCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	scopePermissionID, err := h.appSvc.CreateScopePermission(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create scope permission: "+err.Error())
	}

	// Get the created scope permission
	scopePermissionView, err := h.appSvc.GetScopePermission(c.Context(), scopePermissionID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created scope permission: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Scope permission created successfully", httpresp.SuccessOptions{Data: respdto.ScopePermissionROToDTO(&scopePermissionView)})
}

// GetByID 根据 ID 获取作用域权限关联
func (h *ScopePermissionHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ScopePermissionByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetScopePermission(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Scope permission not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Scope permission retrieved successfully", httpresp.SuccessOptions{Data: respdto.ScopePermissionROToDTO(&result)})
}

// Delete 删除作用域权限关联
func (h *ScopePermissionHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ScopePermissionByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := scopePermissionAppCommands.DeleteScopePermissionCmd{ScopePermissionID: req.ID}
	if err := h.appSvc.DeleteScopePermission(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete scope permission: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Scope permission deleted successfully")
}
