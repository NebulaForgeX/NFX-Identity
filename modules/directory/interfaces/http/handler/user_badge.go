package handler

import (
	userBadgeApp "nfxid/modules/directory/application/user_badges"
	userBadgeAppCommands "nfxid/modules/directory/application/user_badges/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type UserBadgeHandler struct {
	appSvc *userBadgeApp.Service
}

func NewUserBadgeHandler(appSvc *userBadgeApp.Service) *UserBadgeHandler {
	return &UserBadgeHandler{appSvc: appSvc}
}

func (h *UserBadgeHandler) Create(c *fiber.Ctx) error {
	var req reqdto.UserBadgeCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	userBadgeID, err := h.appSvc.CreateUserBadge(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create user badge: "+err.Error())
	}

	// Get the created user badge
	userBadgeView, err := h.appSvc.GetUserBadge(c.Context(), userBadgeID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created user badge: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "User badge created successfully", httpresp.SuccessOptions{Data: respdto.UserBadgeROToDTO(&userBadgeView)})
}

func (h *UserBadgeHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetUserBadge(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "User badge not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User badge retrieved successfully", httpresp.SuccessOptions{Data: respdto.UserBadgeROToDTO(&result)})
}

func (h *UserBadgeHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := userBadgeAppCommands.DeleteUserBadgeCmd{UserBadgeID: req.ID}
	if err := h.appSvc.DeleteUserBadge(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete user badge: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User badge deleted successfully")
}
