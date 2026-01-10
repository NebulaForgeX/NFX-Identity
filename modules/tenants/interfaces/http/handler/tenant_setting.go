package handler

import (
	tenantSettingApp "nfxid/modules/tenants/application/tenant_settings"
	tenantSettingAppCommands "nfxid/modules/tenants/application/tenant_settings/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type TenantSettingHandler struct {
	appSvc *tenantSettingApp.Service
}

func NewTenantSettingHandler(appSvc *tenantSettingApp.Service) *TenantSettingHandler {
	return &TenantSettingHandler{appSvc: appSvc}
}

func (h *TenantSettingHandler) Create(c *fiber.Ctx) error {
	var req reqdto.TenantSettingCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	tenantSettingID, err := h.appSvc.CreateTenantSetting(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create tenant setting: "+err.Error())
	}

	// Get the created tenant setting
	tenantSettingView, err := h.appSvc.GetTenantSetting(c.Context(), tenantSettingID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created tenant setting: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Tenant setting created successfully", httpresp.SuccessOptions{Data: respdto.TenantSettingROToDTO(&tenantSettingView)})
}

func (h *TenantSettingHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetTenantSetting(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Tenant setting not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Tenant setting retrieved successfully", httpresp.SuccessOptions{Data: respdto.TenantSettingROToDTO(&result)})
}

func (h *TenantSettingHandler) Update(c *fiber.Ctx) error {
	var req reqdto.TenantSettingUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateTenantSetting(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update tenant setting: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Tenant setting updated successfully")
}

func (h *TenantSettingHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := tenantSettingAppCommands.DeleteTenantSettingCmd{TenantSettingID: req.ID}
	if err := h.appSvc.DeleteTenantSetting(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete tenant setting: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Tenant setting deleted successfully")
}
