package handler

import (
	userAvatarApp "nfxid/modules/directory/application/user_avatars"
	userAvatarAppCommands "nfxid/modules/directory/application/user_avatars/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type UserAvatarHandler struct {
	appSvc *userAvatarApp.Service
}

func NewUserAvatarHandler(appSvc *userAvatarApp.Service) *UserAvatarHandler {
	return &UserAvatarHandler{appSvc: appSvc}
}

func (h *UserAvatarHandler) CreateOrUpdate(c *fiber.Ctx) error {
	var req reqdto.UserAvatarCreateOrUpdateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateOrUpdateCmd()
	if err := h.appSvc.CreateOrUpdateUserAvatar(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create or update user avatar: "+err.Error())
	}

	// Get the created/updated user avatar
	userAvatarView, err := h.appSvc.GetUserAvatarByUserID(c.Context(), req.UserID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get user avatar: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User avatar created or updated successfully", httpresp.SuccessOptions{Data: respdto.UserAvatarROToDTO(&userAvatarView)})
}

func (h *UserAvatarHandler) GetByUserID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetUserAvatarByUserID(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "User avatar not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User avatar retrieved successfully", httpresp.SuccessOptions{Data: respdto.UserAvatarROToDTO(&result)})
}

func (h *UserAvatarHandler) Update(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	var updateReq reqdto.UserAvatarUpdateImageIDRequestDTO
	if err := c.BodyParser(&updateReq); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := updateReq.ToUpdateImageIDCmd(req.ID)
	if err := h.appSvc.UpdateUserAvatarImageID(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update user avatar: "+err.Error())
	}

	// Get the updated user avatar
	userAvatarView, err := h.appSvc.GetUserAvatarByUserID(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get updated user avatar: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User avatar updated successfully", httpresp.SuccessOptions{Data: respdto.UserAvatarROToDTO(&userAvatarView)})
}

func (h *UserAvatarHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := userAvatarAppCommands.DeleteUserAvatarCmd{UserID: req.ID}
	if err := h.appSvc.DeleteUserAvatar(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete user avatar: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User avatar deleted successfully")
}
