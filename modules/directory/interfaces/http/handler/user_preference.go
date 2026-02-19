package handler

import (
	userPreferenceApp "nfxid/modules/directory/application/user_preferences"
	userPreferenceAppCommands "nfxid/modules/directory/application/user_preferences/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type UserPreferenceHandler struct {
	appSvc *userPreferenceApp.Service
}

func NewUserPreferenceHandler(appSvc *userPreferenceApp.Service) *UserPreferenceHandler {
	return &UserPreferenceHandler{appSvc: appSvc}
}

func (h *UserPreferenceHandler) Create(c fiber.Ctx) error {
	var req reqdto.UserPreferenceCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	userPreferenceID, err := h.appSvc.CreateUserPreference(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created user preference
	userPreferenceView, err := h.appSvc.GetUserPreference(c.Context(), userPreferenceID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "User preference created successfully", httpx.SuccessOptions{Data: respdto.UserPreferenceROToDTO(&userPreferenceView)})
}

func (h *UserPreferenceHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetUserPreference(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "User preference retrieved successfully", httpx.SuccessOptions{Data: respdto.UserPreferenceROToDTO(&result)})
}

func (h *UserPreferenceHandler) Update(c fiber.Ctx) error {
	var req reqdto.UserPreferenceUpdateRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateUserPreference(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User preference updated successfully")
}

func (h *UserPreferenceHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := userPreferenceAppCommands.DeleteUserPreferenceCmd{UserPreferenceID: req.ID}
	if err := h.appSvc.DeleteUserPreference(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User preference deleted successfully")
}

// fiber:context-methods migrated
