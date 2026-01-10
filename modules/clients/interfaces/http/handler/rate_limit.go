package handler

import (
	rateLimitApp "nfxid/modules/clients/application/rate_limits"
	rateLimitAppCommands "nfxid/modules/clients/application/rate_limits/commands"
	"nfxid/modules/clients/interfaces/http/dto/reqdto"
	"nfxid/modules/clients/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
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
func (h *RateLimitHandler) Create(c *fiber.Ctx) error {
	var req reqdto.RateLimitCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	rateLimitID, err := h.appSvc.CreateRateLimit(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create rate limit: "+err.Error())
	}

	// Get the created rate limit
	rateLimitView, err := h.appSvc.GetRateLimit(c.Context(), rateLimitID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created rate limit: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Rate limit created successfully", httpresp.SuccessOptions{Data: respdto.RateLimitROToDTO(&rateLimitView)})
}

// GetByID 根据 ID 获取 Rate Limit
func (h *RateLimitHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.RateLimitByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetRateLimit(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Rate limit not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Rate limit retrieved successfully", httpresp.SuccessOptions{Data: respdto.RateLimitROToDTO(&result)})
}

// Delete 删除 Rate Limit
func (h *RateLimitHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.RateLimitByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := rateLimitAppCommands.DeleteRateLimitCmd{RateLimitID: req.ID}
	if err := h.appSvc.DeleteRateLimit(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete rate limit: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Rate limit deleted successfully")
}
