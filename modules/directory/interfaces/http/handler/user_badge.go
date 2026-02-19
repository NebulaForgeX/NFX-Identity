package handler

import (
	userBadgeApp "nfxid/modules/directory/application/user_badges"
	userBadgeAppCommands "nfxid/modules/directory/application/user_badges/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type UserBadgeHandler struct {
	appSvc *userBadgeApp.Service
}

func NewUserBadgeHandler(appSvc *userBadgeApp.Service) *UserBadgeHandler {
	return &UserBadgeHandler{appSvc: appSvc}
}

func (h *UserBadgeHandler) Create(c fiber.Ctx) error {
	var req reqdto.UserBadgeCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	userBadgeID, err := h.appSvc.CreateUserBadge(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created user badge
	userBadgeView, err := h.appSvc.GetUserBadge(c.Context(), userBadgeID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "User badge created successfully", httpx.SuccessOptions{Data: respdto.UserBadgeROToDTO(&userBadgeView)})
}

func (h *UserBadgeHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetUserBadge(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "User badge retrieved successfully", httpx.SuccessOptions{Data: respdto.UserBadgeROToDTO(&result)})
}

func (h *UserBadgeHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := userBadgeAppCommands.DeleteUserBadgeCmd{UserBadgeID: req.ID}
	if err := h.appSvc.DeleteUserBadge(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User badge deleted successfully")
}

// fiber:context-methods migrated
