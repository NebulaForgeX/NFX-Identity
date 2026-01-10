package handler

import (
	scopeApp "nfxid/modules/access/application/scopes"
	scopeAppCommands "nfxid/modules/access/application/scopes/commands"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type ScopeHandler struct {
	appSvc *scopeApp.Service
}

func NewScopeHandler(appSvc *scopeApp.Service) *ScopeHandler {
	return &ScopeHandler{
		appSvc: appSvc,
	}
}

// Create 创建作用域
func (h *ScopeHandler) Create(c *fiber.Ctx) error {
	var req reqdto.ScopeCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	err := h.appSvc.CreateScope(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create scope: "+err.Error())
	}

	// Get the created scope
	scopeView, err := h.appSvc.GetScope(c.Context(), req.Scope)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created scope: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Scope created successfully", httpresp.SuccessOptions{Data: respdto.ScopeROToDTO(&scopeView)})
}

// GetByScope 根据 Scope 获取作用域
func (h *ScopeHandler) GetByScope(c *fiber.Ctx) error {
	var req reqdto.ScopeByScopeRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetScope(c.Context(), req.Scope)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Scope not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Scope retrieved successfully", httpresp.SuccessOptions{Data: respdto.ScopeROToDTO(&result)})
}

// Update 更新作用域
func (h *ScopeHandler) Update(c *fiber.Ctx) error {
	var req reqdto.ScopeUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateScope(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update scope: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Scope updated successfully")
}

// Delete 删除作用域
func (h *ScopeHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ScopeByScopeRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := scopeAppCommands.DeleteScopeCmd{Scope: req.Scope}
	if err := h.appSvc.DeleteScope(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete scope: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Scope deleted successfully")
}
