package handler

import (
	rateLimitApp "nfxid/modules/clients/application/rate_limits"
	rateLimitAppCommands "nfxid/modules/clients/application/rate_limits/commands"
	"nfxid/modules/clients/interfaces/http/dto/reqdto"
	"nfxid/modules/clients/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type RateLimitHandler struct {
	appSvc *rateLimitApp.Service
}

func NewRateLimitHandler(appSvc *rateLimitApp.Service) *RateLimitHandler {
	return &RateLimitHandler{
		appSvc: appSvc,
	}
}

// Create 创建 Rate Limit
func (h *RateLimitHandler) Create(c fiber.Ctx) error {
	var req reqdto.RateLimitCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	rateLimitID, err := h.appSvc.CreateRateLimit(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created rate limit
	rateLimitView, err := h.appSvc.GetRateLimit(c.Context(), rateLimitID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Rate limit created successfully", httpx.SuccessOptions{Data: respdto.RateLimitROToDTO(&rateLimitView)})
}

// GetByID 根据 ID 获取 Rate Limit
func (h *RateLimitHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.RateLimitByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetRateLimit(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Rate limit retrieved successfully", httpx.SuccessOptions{Data: respdto.RateLimitROToDTO(&result)})
}

// Delete 删除 Rate Limit
func (h *RateLimitHandler) Delete(c fiber.Ctx) error {
	var req reqdto.RateLimitByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := rateLimitAppCommands.DeleteRateLimitCmd{RateLimitID: req.ID}
	if err := h.appSvc.DeleteRateLimit(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Rate limit deleted successfully")
}

// fiber:context-methods migrated
