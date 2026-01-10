package handler

import (
	tenantAppApp "nfxid/modules/tenants/application/tenant_apps"
	tenantAppAppCommands "nfxid/modules/tenants/application/tenant_apps/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type TenantAppHandler struct {
	appSvc *tenantAppApp.Service
}

func NewTenantAppHandler(appSvc *tenantAppApp.Service) *TenantAppHandler {
	return &TenantAppHandler{appSvc: appSvc}
}

func (h *TenantAppHandler) Create(c *fiber.Ctx) error {
	var req reqdto.TenantAppCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	tenantAppID, err := h.appSvc.CreateTenantApp(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create tenant app: "+err.Error())
	}

	// Get the created tenant app
	tenantAppView, err := h.appSvc.GetTenantApp(c.Context(), tenantAppID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created tenant app: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Tenant app created successfully", httpresp.SuccessOptions{Data: respdto.TenantAppROToDTO(&tenantAppView)})
}

func (h *TenantAppHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetTenantApp(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Tenant app not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Tenant app retrieved successfully", httpresp.SuccessOptions{Data: respdto.TenantAppROToDTO(&result)})
}

func (h *TenantAppHandler) Update(c *fiber.Ctx) error {
	// Update can be status or settings
	// Try to parse as status update first
	var statusReq reqdto.TenantAppUpdateStatusRequestDTO
	if err := c.ParamsParser(&statusReq); err == nil {
		if err := c.BodyParser(&statusReq); err == nil {
			cmd := statusReq.ToUpdateStatusCmd()
			if err := h.appSvc.UpdateTenantAppStatus(c.Context(), cmd); err != nil {
				return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update tenant app status: "+err.Error())
			}
			return httpresp.Success(c, fiber.StatusOK, "Tenant app status updated successfully")
		}
	}

	// Try as settings update
	var settingsReq reqdto.TenantAppUpdateSettingsRequestDTO
	if err := c.ParamsParser(&settingsReq); err == nil {
		if err := c.BodyParser(&settingsReq); err == nil {
			cmd := settingsReq.ToUpdateSettingsCmd()
			if err := h.appSvc.UpdateTenantAppSettings(c.Context(), cmd); err != nil {
				return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update tenant app settings: "+err.Error())
			}
			return httpresp.Success(c, fiber.StatusOK, "Tenant app settings updated successfully")
		}
	}

	return httpresp.Error(c, fiber.StatusBadRequest, "Invalid update request")
}

func (h *TenantAppHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := tenantAppAppCommands.DeleteTenantAppCmd{TenantAppID: req.ID}
	if err := h.appSvc.DeleteTenantApp(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete tenant app: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Tenant app deleted successfully")
}
