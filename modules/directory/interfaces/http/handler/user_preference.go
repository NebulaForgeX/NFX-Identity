package handler

import (
	userPreferenceApp "nfxid/modules/directory/application/user_preferences"
	userPreferenceAppCommands "nfxid/modules/directory/application/user_preferences/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type UserPreferenceHandler struct {
	appSvc *userPreferenceApp.Service
}

func NewUserPreferenceHandler(appSvc *userPreferenceApp.Service) *UserPreferenceHandler {
	return &UserPreferenceHandler{appSvc: appSvc}
}

func (h *UserPreferenceHandler) Create(c *fiber.Ctx) error {
	var req reqdto.UserPreferenceCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	userPreferenceID, err := h.appSvc.CreateUserPreference(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create user preference: "+err.Error())
	}

	// Get the created user preference
	userPreferenceView, err := h.appSvc.GetUserPreference(c.Context(), userPreferenceID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created user preference: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "User preference created successfully", httpresp.SuccessOptions{Data: respdto.UserPreferenceROToDTO(&userPreferenceView)})
}

func (h *UserPreferenceHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetUserPreference(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "User preference not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User preference retrieved successfully", httpresp.SuccessOptions{Data: respdto.UserPreferenceROToDTO(&result)})
}

func (h *UserPreferenceHandler) Update(c *fiber.Ctx) error {
	var req reqdto.UserPreferenceUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateUserPreference(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update user preference: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User preference updated successfully")
}

func (h *UserPreferenceHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := userPreferenceAppCommands.DeleteUserPreferenceCmd{UserPreferenceID: req.ID}
	if err := h.appSvc.DeleteUserPreference(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete user preference: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User preference deleted successfully")
}
