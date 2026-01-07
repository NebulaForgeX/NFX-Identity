package handler

import (
	badgeApp "nfxid/modules/auth/application/badge"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
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
	result, err := h.appSvc.CreateBadge(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create badge: "+err.Error())
	}

	// Get the created badge view
	badgeView, err := h.appSvc.GetBadge(c.Context(), result.ID())
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created badge: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Badge created successfully", httpresp.SuccessOptions{Data: respdto.BadgeViewToDTO(&badgeView)})
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

	return httpresp.Success(c, fiber.StatusOK, "Badge retrieved successfully", httpresp.SuccessOptions{Data: respdto.BadgeViewToDTO(&result)})
}

// GetByName 根据名称获取徽章
func (h *BadgeHandler) GetByName(c *fiber.Ctx) error {
	var req reqdto.BadgeByNameRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetBadgeByName(c.Context(), req.Name)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Badge not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Badge retrieved successfully", httpresp.SuccessOptions{Data: respdto.BadgeViewToDTO(&result)})
}

// GetAll 获取徽章列表
func (h *BadgeHandler) GetAll(c *fiber.Ctx) error {
	var query reqdto.BadgeQueryParamsDTO
	if err := c.QueryParser(&query); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid query parameters: "+err.Error())
	}

	listQuery := query.ToListQuery()
	result, err := h.appSvc.GetBadgeList(c.Context(), listQuery)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get badges: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Badges retrieved successfully", httpresp.SuccessOptions{
		Data: httpresp.ToList(respdto.BadgeListViewToDTO(result.Items), int(result.Total)),
	})
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

// Delete 删除徽章
func (h *BadgeHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.BadgeByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := badgeApp.DeleteBadgeCmd{BadgeID: req.ID}
	if err := h.appSvc.DeleteBadge(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete badge: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Badge deleted successfully")
}
