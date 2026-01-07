package handler

import (
	profileApp "nfxid/modules/auth/application/profile"
	"nfxid/modules/auth/interfaces/http/dto/reqdto"
	"nfxid/modules/auth/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type ProfileHandler struct {
	appSvc *profileApp.Service
}

func NewProfileHandler(appSvc *profileApp.Service) *ProfileHandler {
	return &ProfileHandler{
		appSvc: appSvc,
	}
}

// Create 创建资料
func (h *ProfileHandler) Create(c *fiber.Ctx) error {
	var req reqdto.ProfileCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	result, err := h.appSvc.CreateProfile(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create profile: "+err.Error())
	}

	// Get the created profile view
	profileView, err := h.appSvc.GetProfile(c.Context(), result.ID())
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created profile: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Profile created successfully", httpresp.SuccessOptions{Data: respdto.ProfileViewToDTO(&profileView)})
}

// GetByID 根据 ID 获取资料
func (h *ProfileHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ProfileByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetProfile(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Profile not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Profile retrieved successfully", httpresp.SuccessOptions{Data: respdto.ProfileViewToDTO(&result)})
}

// GetByUserID 根据用户ID获取资料
func (h *ProfileHandler) GetByUserID(c *fiber.Ctx) error {
	var req reqdto.ProfileByUserIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetProfileByUserID(c.Context(), req.UserID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Profile not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Profile retrieved successfully", httpresp.SuccessOptions{Data: respdto.ProfileViewToDTO(&result)})
}

// GetAll 获取资料列表
func (h *ProfileHandler) GetAll(c *fiber.Ctx) error {
	var query reqdto.ProfileQueryParamsDTO
	if err := c.QueryParser(&query); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid query parameters: "+err.Error())
	}

	listQuery := query.ToListQuery()
	result, err := h.appSvc.GetProfileList(c.Context(), listQuery)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get profiles: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Profiles retrieved successfully", httpresp.SuccessOptions{
		Data: httpresp.ToList(respdto.ProfileListViewToDTO(result.Items), int(result.Total)),
	})
}

// Update 更新资料
func (h *ProfileHandler) Update(c *fiber.Ctx) error {
	var req reqdto.ProfileUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateProfile(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update profile: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Profile updated successfully")
}

// Delete 删除资料
func (h *ProfileHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ProfileByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := profileApp.DeleteProfileCmd{ProfileID: req.ID}
	if err := h.appSvc.DeleteProfile(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete profile: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Profile deleted successfully")
}
