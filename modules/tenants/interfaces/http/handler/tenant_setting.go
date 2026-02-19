package handler

import (
	tenantSettingApp "nfxid/modules/tenants/application/tenant_settings"
	tenantSettingAppCommands "nfxid/modules/tenants/application/tenant_settings/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type TenantSettingHandler struct {
	appSvc *tenantSettingApp.Service
}

func NewTenantSettingHandler(appSvc *tenantSettingApp.Service) *TenantSettingHandler {
	return &TenantSettingHandler{appSvc: appSvc}
}

func (h *TenantSettingHandler) Create(c fiber.Ctx) error {
	var req reqdto.TenantSettingCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	ctx := c.Context()
	tenantSettingID, err := h.appSvc.CreateTenantSetting(ctx, cmd)
	if err != nil {
		return err
	}

	tenantSettingView, err := h.appSvc.GetTenantSetting(ctx, tenantSettingID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Tenant setting created successfully", httpx.SuccessOptions{Data: respdto.TenantSettingROToDTO(&tenantSettingView)})
}

func (h *TenantSettingHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetTenantSetting(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Tenant setting retrieved successfully", httpx.SuccessOptions{Data: respdto.TenantSettingROToDTO(&result)})
}

func (h *TenantSettingHandler) Update(c fiber.Ctx) error {
	var req reqdto.TenantSettingUpdateRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateTenantSetting(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Tenant setting updated successfully")
}

func (h *TenantSettingHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := tenantSettingAppCommands.DeleteTenantSettingCmd{TenantSettingID: req.ID}
	if err := h.appSvc.DeleteTenantSetting(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Tenant setting deleted successfully")
}

// fiber:context-methods migrated
