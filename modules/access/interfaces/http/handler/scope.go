package handler

import (
	scopeApp "nfxid/modules/access/application/scopes"
	scopeAppCommands "nfxid/modules/access/application/scopes/commands"
	"nfxid/modules/access/interfaces/http/dto/reqdto"
	"nfxid/modules/access/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
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
func (h *ScopeHandler) Create(c fiber.Ctx) error {
	var req reqdto.ScopeCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	err := h.appSvc.CreateScope(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created scope
	scopeView, err := h.appSvc.GetScope(c.Context(), req.Scope)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Scope created successfully", httpx.SuccessOptions{Data: respdto.ScopeROToDTO(&scopeView)})
}

// GetByScope 根据 Scope 获取作用域
func (h *ScopeHandler) GetByScope(c fiber.Ctx) error {
	var req reqdto.ScopeByScopeRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetScope(c.Context(), req.Scope)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Scope retrieved successfully", httpx.SuccessOptions{Data: respdto.ScopeROToDTO(&result)})
}

// Update 更新作用域
func (h *ScopeHandler) Update(c fiber.Ctx) error {
	var req reqdto.ScopeUpdateRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateScope(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Scope updated successfully")
}

// Delete 删除作用域
func (h *ScopeHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ScopeByScopeRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := scopeAppCommands.DeleteScopeCmd{Scope: req.Scope}
	if err := h.appSvc.DeleteScope(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Scope deleted successfully")
}

// fiber:context-methods migrated
