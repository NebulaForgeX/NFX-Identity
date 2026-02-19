package handler

import (
	tenantAppApp "nfxid/modules/tenants/application/tenant_apps"
	tenantAppAppCommands "nfxid/modules/tenants/application/tenant_apps/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type TenantAppHandler struct {
	appSvc *tenantAppApp.Service
}

func NewTenantAppHandler(appSvc *tenantAppApp.Service) *TenantAppHandler {
	return &TenantAppHandler{appSvc: appSvc}
}

func (h *TenantAppHandler) Create(c fiber.Ctx) error {
	var req reqdto.TenantAppCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	tenantAppID, err := h.appSvc.CreateTenantApp(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created tenant app
	tenantAppView, err := h.appSvc.GetTenantApp(c.Context(), tenantAppID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Tenant app created successfully", httpx.SuccessOptions{Data: respdto.TenantAppROToDTO(&tenantAppView)})
}

func (h *TenantAppHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetTenantApp(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Tenant app retrieved successfully", httpx.SuccessOptions{Data: respdto.TenantAppROToDTO(&result)})
}

func (h *TenantAppHandler) Update(c fiber.Ctx) error {
	// Update can be status or settings
	// Try to parse as status update first
	var statusReq reqdto.TenantAppUpdateStatusRequestDTO
	if err := c.Bind().URI(&statusReq); err == nil {
		if err := c.Bind().Body(&statusReq); err == nil {
			cmd := statusReq.ToUpdateStatusCmd()
			if err := h.appSvc.UpdateTenantAppStatus(c.Context(), cmd); err != nil {
				return err
			}
			return fiberx.OK(c, "Tenant app status updated successfully")
		}
	}

	// Try as settings update
	var settingsReq reqdto.TenantAppUpdateSettingsRequestDTO
	if err := c.Bind().URI(&settingsReq); err == nil {
		if err := c.Bind().Body(&settingsReq); err == nil {
			cmd := settingsReq.ToUpdateSettingsCmd()
			if err := h.appSvc.UpdateTenantAppSettings(c.Context(), cmd); err != nil {
				return err
			}
			return fiberx.OK(c, "Tenant app settings updated successfully")
		}
	}

	return errx.ErrInvalidParams
}

func (h *TenantAppHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := tenantAppAppCommands.DeleteTenantAppCmd{TenantAppID: req.ID}
	if err := h.appSvc.DeleteTenantApp(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Tenant app deleted successfully")
}

// fiber:context-methods migrated
