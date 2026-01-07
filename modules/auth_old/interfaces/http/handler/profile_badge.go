package handler

import (
	profileBadgeApp "nfxid/modules/auth/application/profile_badge"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type ProfileBadgeHandler struct {
	appSvc *profileBadgeApp.Service
}

func NewProfileBadgeHandler(appSvc *profileBadgeApp.Service) *ProfileBadgeHandler {
	return &ProfileBadgeHandler{
		appSvc: appSvc,
	}
}

// Create 创建用户徽章关联
func (h *ProfileBadgeHandler) Create(c *fiber.Ctx) error {
	var req reqdto.ProfileBadgeCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	result, err := h.appSvc.CreateProfileBadge(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create profile badge: "+err.Error())
	}

	// Get the created profile badge view
	profileBadgeView, err := h.appSvc.GetProfileBadge(c.Context(), result.ID())
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created profile badge: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Profile badge created successfully", httpresp.SuccessOptions{Data: respdto.ProfileBadgeViewToDTO(&profileBadgeView, nil)})
}

// GetByID 根据 ID 获取用户徽章关联
func (h *ProfileBadgeHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ProfileBadgeByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetProfileBadge(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Profile badge not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Profile badge retrieved successfully", httpresp.SuccessOptions{Data: respdto.ProfileBadgeViewToDTO(&result, nil)})
}

// GetByProfileID 根据 ProfileID 获取用户徽章关联列表
func (h *ProfileBadgeHandler) GetByProfileID(c *fiber.Ctx) error {
	var req reqdto.ProfileBadgeByProfileIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	results, err := h.appSvc.GetProfileBadgesByProfileID(c.Context(), req.ProfileID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get profile badges: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Profile badges retrieved successfully", httpresp.SuccessOptions{Data: respdto.ProfileBadgeListViewToDTO(results)})
}

// GetByBadgeID 根据 BadgeID 获取用户徽章关联列表
func (h *ProfileBadgeHandler) GetByBadgeID(c *fiber.Ctx) error {
	var req reqdto.ProfileBadgeByBadgeIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	results, err := h.appSvc.GetProfileBadgesByBadgeID(c.Context(), req.BadgeID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get profile badges: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Profile badges retrieved successfully", httpresp.SuccessOptions{Data: respdto.ProfileBadgeListViewToDTO(results)})
}

// Update 更新用户徽章关联
func (h *ProfileBadgeHandler) Update(c *fiber.Ctx) error {
	var req reqdto.ProfileBadgeUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateProfileBadge(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update profile badge: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Profile badge updated successfully")
}

// Delete 删除用户徽章关联
func (h *ProfileBadgeHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ProfileBadgeByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := profileBadgeApp.DeleteProfileBadgeCmd{ProfileBadgeID: req.ID}
	if err := h.appSvc.DeleteProfileBadge(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete profile badge: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Profile badge deleted successfully")
}

// DeleteByProfileAndBadge 根据 ProfileID 和 BadgeID 删除关联
func (h *ProfileBadgeHandler) DeleteByProfileAndBadge(c *fiber.Ctx) error {
	var req reqdto.ProfileBadgeDeleteByProfileAndBadgeRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := profileBadgeApp.DeleteProfileBadgeByProfileAndBadgeCmd{
		ProfileID: req.ProfileID,
		BadgeID:   req.BadgeID,
	}
	if err := h.appSvc.DeleteProfileBadgeByProfileAndBadge(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete profile badge: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Profile badge deleted successfully")
}
