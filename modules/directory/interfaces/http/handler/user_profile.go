package handler

import (
	userProfileApp "nfxid/modules/directory/application/user_profiles"
	userProfileAppCommands "nfxid/modules/directory/application/user_profiles/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type UserProfileHandler struct {
	appSvc *userProfileApp.Service
}

func NewUserProfileHandler(appSvc *userProfileApp.Service) *UserProfileHandler {
	return &UserProfileHandler{appSvc: appSvc}
}

func (h *UserProfileHandler) Create(c *fiber.Ctx) error {
	var req reqdto.UserProfileCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	userProfileID, err := h.appSvc.CreateUserProfile(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create user profile: "+err.Error())
	}

	// Get the created user profile
	userProfileView, err := h.appSvc.GetUserProfile(c.Context(), userProfileID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created user profile: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "User profile created successfully", httpresp.SuccessOptions{Data: respdto.UserProfileROToDTO(&userProfileView)})
}

func (h *UserProfileHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetUserProfile(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "User profile not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User profile retrieved successfully", httpresp.SuccessOptions{Data: respdto.UserProfileROToDTO(&result)})
}

func (h *UserProfileHandler) Update(c *fiber.Ctx) error {
	var req reqdto.UserProfileUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateUserProfile(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update user profile: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User profile updated successfully")
}

func (h *UserProfileHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := userProfileAppCommands.DeleteUserProfileCmd{UserProfileID: req.ID}
	if err := h.appSvc.DeleteUserProfile(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete user profile: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User profile deleted successfully")
}
