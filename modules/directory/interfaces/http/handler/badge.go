package handler

import (
	badgeApp "nfxid/modules/directory/application/badges"
	badgeAppCommands "nfxid/modules/directory/application/badges/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type BadgeHandler struct {
	appSvc *badgeApp.Service
}

func NewBadgeHandler(appSvc *badgeApp.Service) *BadgeHandler {
	return &BadgeHandler{
		appSvc: appSvc,
	}
}

// Create 创建徽章
func (h *BadgeHandler) Create(c *fiber.Ctx) error {
	var req reqdto.BadgeCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	badgeID, err := h.appSvc.CreateBadge(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create badge: "+err.Error())
	}

	// Get the created badge
	badgeView, err := h.appSvc.GetBadge(c.Context(), badgeID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created badge: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Badge created successfully", httpresp.SuccessOptions{Data: respdto.BadgeROToDTO(&badgeView)})
}

// GetByID 根据 ID 获取徽章
func (h *BadgeHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.BadgeByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetBadge(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Badge not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Badge retrieved successfully", httpresp.SuccessOptions{Data: respdto.BadgeROToDTO(&result)})
}

// GetByName 根据 Name 获取徽章
func (h *BadgeHandler) GetByName(c *fiber.Ctx) error {
	var req reqdto.BadgeByNameRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetBadgeByName(c.Context(), req.Name)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Badge not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Badge retrieved successfully", httpresp.SuccessOptions{Data: respdto.BadgeROToDTO(&result)})
}

// Update 更新徽章
func (h *BadgeHandler) Update(c *fiber.Ctx) error {
	var req reqdto.BadgeUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateBadge(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update badge: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Badge updated successfully")
}

// Delete 删除徽章（软删除）
func (h *BadgeHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.BadgeByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := badgeAppCommands.DeleteBadgeCmd{BadgeID: req.ID}
	if err := h.appSvc.DeleteBadge(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete badge: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Badge deleted successfully")
}
