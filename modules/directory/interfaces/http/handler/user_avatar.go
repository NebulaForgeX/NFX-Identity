package handler

import (
	userAvatarApp "nfxid/modules/directory/application/user_avatars"
	userAvatarAppCommands "nfxid/modules/directory/application/user_avatars/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type UserAvatarHandler struct {
	appSvc *userAvatarApp.Service
}

func NewUserAvatarHandler(appSvc *userAvatarApp.Service) *UserAvatarHandler {
	return &UserAvatarHandler{appSvc: appSvc}
}

func (h *UserAvatarHandler) CreateOrUpdate(c fiber.Ctx) error {
	var req reqdto.UserAvatarCreateOrUpdateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateOrUpdateCmd()
	if err := h.appSvc.CreateOrUpdateUserAvatar(c.Context(), cmd); err != nil {
		return err
	}

	// Get the created/updated user avatar
	userAvatarView, err := h.appSvc.GetUserAvatarByUserID(c.Context(), req.UserID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "User avatar created or updated successfully", httpx.SuccessOptions{Data: respdto.UserAvatarROToDTO(&userAvatarView)})
}

func (h *UserAvatarHandler) GetByUserID(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetUserAvatarByUserID(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "User avatar retrieved successfully", httpx.SuccessOptions{Data: respdto.UserAvatarROToDTO(&result)})
}

func (h *UserAvatarHandler) Update(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	var updateReq reqdto.UserAvatarUpdateImageIDRequestDTO
	if err := c.Bind().Body(&updateReq); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := updateReq.ToUpdateImageIDCmd(req.ID)
	if err := h.appSvc.UpdateUserAvatarImageID(c.Context(), cmd); err != nil {
		return err
	}

	// Get the updated user avatar
	userAvatarView, err := h.appSvc.GetUserAvatarByUserID(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "User avatar updated successfully", httpx.SuccessOptions{Data: respdto.UserAvatarROToDTO(&userAvatarView)})
}

func (h *UserAvatarHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := userAvatarAppCommands.DeleteUserAvatarCmd{UserID: req.ID}
	if err := h.appSvc.DeleteUserAvatar(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User avatar deleted successfully")
}

// fiber:context-methods migrated
