package handler

import (
	ipAllowlistApp "nfxid/modules/clients/application/ip_allowlist"
	ipAllowlistAppCommands "nfxid/modules/clients/application/ip_allowlist/commands"
	"nfxid/modules/clients/interfaces/http/dto/reqdto"
	"nfxid/modules/clients/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type IPAllowlistHandler struct {
	appSvc *ipAllowlistApp.Service
}

func NewIPAllowlistHandler(appSvc *ipAllowlistApp.Service) *IPAllowlistHandler {
	return &IPAllowlistHandler{
		appSvc: appSvc,
	}
}

// Create 创建 IP Allowlist
func (h *IPAllowlistHandler) Create(c *fiber.Ctx) error {
	var req reqdto.IPAllowlistCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	ipAllowlistID, err := h.appSvc.CreateIPAllowlist(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create IP allowlist: "+err.Error())
	}

	// Get the created IP allowlist
	ipAllowlistView, err := h.appSvc.GetIPAllowlist(c.Context(), ipAllowlistID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created IP allowlist: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "IP allowlist created successfully", httpresp.SuccessOptions{Data: respdto.IPAllowlistROToDTO(&ipAllowlistView)})
}

// GetByID 根据 ID 获取 IP Allowlist
func (h *IPAllowlistHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.IPAllowlistByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetIPAllowlist(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "IP allowlist not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "IP allowlist retrieved successfully", httpresp.SuccessOptions{Data: respdto.IPAllowlistROToDTO(&result)})
}

// Delete 删除 IP Allowlist
func (h *IPAllowlistHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.IPAllowlistDeleteRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := ipAllowlistAppCommands.DeleteIPAllowlistCmd{RuleID: req.RuleID}
	if err := h.appSvc.DeleteIPAllowlist(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete IP allowlist: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "IP allowlist deleted successfully")
}
