package handler

import (
	userProfileApp "nfxid/modules/directory/application/user_profiles"
	userProfileAppCommands "nfxid/modules/directory/application/user_profiles/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type UserProfileHandler struct {
	appSvc *userProfileApp.Service
}

func NewUserProfileHandler(appSvc *userProfileApp.Service) *UserProfileHandler {
	return &UserProfileHandler{appSvc: appSvc}
}

func (h *UserProfileHandler) Create(c fiber.Ctx) error {
	var req reqdto.UserProfileCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	userProfileID, err := h.appSvc.CreateUserProfile(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created user profile
	userProfileView, err := h.appSvc.GetUserProfile(c.Context(), userProfileID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "User profile created successfully", httpx.SuccessOptions{Data: respdto.UserProfileROToDTO(&userProfileView)})
}

func (h *UserProfileHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.UserProfileByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetUserProfile(c.Context(), req.UserProfileID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "User profile retrieved successfully", httpx.SuccessOptions{Data: respdto.UserProfileROToDTO(&result)})
}

func (h *UserProfileHandler) Update(c fiber.Ctx) error {
	var req reqdto.UserProfileUpdateRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateUserProfile(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User profile updated successfully")
}

func (h *UserProfileHandler) Delete(c fiber.Ctx) error {
	var req reqdto.UserProfileByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := userProfileAppCommands.DeleteUserProfileCmd{UserProfileID: req.UserProfileID}
	if err := h.appSvc.DeleteUserProfile(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User profile deleted successfully")
}

// fiber:context-methods migrated
