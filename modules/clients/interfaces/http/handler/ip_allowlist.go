package handler

import (
	ipAllowlistApp "nfxid/modules/clients/application/ip_allowlist"
	ipAllowlistAppCommands "nfxid/modules/clients/application/ip_allowlist/commands"
	"nfxid/modules/clients/interfaces/http/dto/reqdto"
	"nfxid/modules/clients/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
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
func (h *IPAllowlistHandler) Create(c fiber.Ctx) error {
	var req reqdto.IPAllowlistCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	ipAllowlistID, err := h.appSvc.CreateIPAllowlist(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created IP allowlist
	ipAllowlistView, err := h.appSvc.GetIPAllowlist(c.Context(), ipAllowlistID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "IP allowlist created successfully", httpx.SuccessOptions{Data: respdto.IPAllowlistROToDTO(&ipAllowlistView)})
}

// GetByID 根据 ID 获取 IP Allowlist
func (h *IPAllowlistHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.IPAllowlistByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetIPAllowlist(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "IP allowlist retrieved successfully", httpx.SuccessOptions{Data: respdto.IPAllowlistROToDTO(&result)})
}

// Delete 删除 IP Allowlist
func (h *IPAllowlistHandler) Delete(c fiber.Ctx) error {
	var req reqdto.IPAllowlistDeleteRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := ipAllowlistAppCommands.DeleteIPAllowlistCmd{RuleID: req.RuleID}
	if err := h.appSvc.DeleteIPAllowlist(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "IP allowlist deleted successfully")
}

// fiber:context-methods migrated
