package handler

import (
	tenantApp "nfxid/modules/tenants/application/tenants"
	tenantAppCommands "nfxid/modules/tenants/application/tenants/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type TenantHandler struct {
	appSvc *tenantApp.Service
}

func NewTenantHandler(appSvc *tenantApp.Service) *TenantHandler {
	return &TenantHandler{
		appSvc: appSvc,
	}
}

// Create 创建租户
func (h *TenantHandler) Create(c fiber.Ctx) error {
	var req reqdto.TenantCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	tenantID, err := h.appSvc.CreateTenant(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created tenant
	tenantView, err := h.appSvc.GetTenant(c.Context(), tenantID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Tenant created successfully", httpx.SuccessOptions{Data: respdto.TenantROToDTO(&tenantView)})
}

// GetByID 根据 ID 获取租户
func (h *TenantHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.TenantByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetTenant(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Tenant retrieved successfully", httpx.SuccessOptions{Data: respdto.TenantROToDTO(&result)})
}

// GetByTenantID 根据 TenantID 获取租户
func (h *TenantHandler) GetByTenantID(c fiber.Ctx) error {
	var req reqdto.TenantByTenantIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetTenantByTenantID(c.Context(), req.TenantID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Tenant retrieved successfully", httpx.SuccessOptions{Data: respdto.TenantROToDTO(&result)})
}

// Update 更新租户
func (h *TenantHandler) Update(c fiber.Ctx) error {
	var req reqdto.TenantUpdateRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateTenant(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Tenant updated successfully")
}

// UpdateStatus 更新租户状态
func (h *TenantHandler) UpdateStatus(c fiber.Ctx) error {
	var req reqdto.TenantUpdateStatusRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateStatusCmd()
	if err := h.appSvc.UpdateTenantStatus(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Tenant status updated successfully")
}

// Delete 删除租户（软删除）
func (h *TenantHandler) Delete(c fiber.Ctx) error {
	var req reqdto.TenantByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := tenantAppCommands.DeleteTenantCmd{TenantID: req.ID}
	if err := h.appSvc.DeleteTenant(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Tenant deleted successfully")
}

// fiber:context-methods migrated
