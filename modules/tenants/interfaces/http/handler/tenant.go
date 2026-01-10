package handler

import (
	tenantApp "nfxid/modules/tenants/application/tenants"
	tenantAppCommands "nfxid/modules/tenants/application/tenants/commands"
	"nfxid/modules/tenants/interfaces/http/dto/reqdto"
	"nfxid/modules/tenants/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
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
func (h *TenantHandler) Create(c *fiber.Ctx) error {
	var req reqdto.TenantCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	tenantID, err := h.appSvc.CreateTenant(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create tenant: "+err.Error())
	}

	// Get the created tenant
	tenantView, err := h.appSvc.GetTenant(c.Context(), tenantID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created tenant: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Tenant created successfully", httpresp.SuccessOptions{Data: respdto.TenantROToDTO(&tenantView)})
}

// GetByID 根据 ID 获取租户
func (h *TenantHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.TenantByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetTenant(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Tenant not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Tenant retrieved successfully", httpresp.SuccessOptions{Data: respdto.TenantROToDTO(&result)})
}

// GetByTenantID 根据 TenantID 获取租户
func (h *TenantHandler) GetByTenantID(c *fiber.Ctx) error {
	var req reqdto.TenantByTenantIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetTenantByTenantID(c.Context(), req.TenantID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Tenant not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Tenant retrieved successfully", httpresp.SuccessOptions{Data: respdto.TenantROToDTO(&result)})
}

// Update 更新租户
func (h *TenantHandler) Update(c *fiber.Ctx) error {
	var req reqdto.TenantUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateTenant(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update tenant: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Tenant updated successfully")
}

// UpdateStatus 更新租户状态
func (h *TenantHandler) UpdateStatus(c *fiber.Ctx) error {
	var req reqdto.TenantUpdateStatusRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateStatusCmd()
	if err := h.appSvc.UpdateTenantStatus(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update tenant status: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Tenant status updated successfully")
}

// Delete 删除租户（软删除）
func (h *TenantHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.TenantByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := tenantAppCommands.DeleteTenantCmd{TenantID: req.ID}
	if err := h.appSvc.DeleteTenant(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete tenant: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Tenant deleted successfully")
}
