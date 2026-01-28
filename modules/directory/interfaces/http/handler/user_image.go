package handler

import (
	userImageApp "nfxid/modules/directory/application/user_images"
	userImageAppCommands "nfxid/modules/directory/application/user_images/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type UserImageHandler struct {
	appSvc *userImageApp.Service
}

func NewUserImageHandler(appSvc *userImageApp.Service) *UserImageHandler {
	return &UserImageHandler{appSvc: appSvc}
}

func (h *UserImageHandler) Create(c *fiber.Ctx) error {
	var req reqdto.UserImageCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	userImageID, err := h.appSvc.CreateUserImage(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create user image: "+err.Error())
	}

	// Get the created user image
	userImageView, err := h.appSvc.GetUserImage(c.Context(), userImageID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created user image: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "User image created successfully", httpresp.SuccessOptions{Data: respdto.UserImageROToDTO(&userImageView)})
}

func (h *UserImageHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetUserImage(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "User image not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User image retrieved successfully", httpresp.SuccessOptions{Data: respdto.UserImageROToDTO(&result)})
}

func (h *UserImageHandler) GetByUserID(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	results, err := h.appSvc.GetUserImagesByUserID(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get user images: "+err.Error())
	}

	dtos := respdto.UserImageListROToDTO(results)
	return httpresp.Success(c, fiber.StatusOK, "User images retrieved successfully", httpresp.SuccessOptions{Data: dtos})
}

func (h *UserImageHandler) GetCurrent(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetCurrentUserImageByUserID(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Current user image not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Current user image retrieved successfully", httpresp.SuccessOptions{Data: respdto.UserImageROToDTO(&result)})
}

func (h *UserImageHandler) UpdateDisplayOrder(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	var updateReq reqdto.UserImageUpdateDisplayOrderRequestDTO
	if err := c.BodyParser(&updateReq); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := updateReq.ToUpdateDisplayOrderCmd(req.ID)
	if err := h.appSvc.UpdateUserImageDisplayOrder(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update user image display order: "+err.Error())
	}

	// Get the updated user image
	userImageView, err := h.appSvc.GetUserImage(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get updated user image: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User image display order updated successfully", httpresp.SuccessOptions{Data: respdto.UserImageROToDTO(&userImageView)})
}

func (h *UserImageHandler) Update(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	var updateReq reqdto.UserImageUpdateImageIDRequestDTO
	if err := c.BodyParser(&updateReq); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := updateReq.ToUpdateImageIDCmd(req.ID)
	if err := h.appSvc.UpdateUserImageImageID(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update user image: "+err.Error())
	}

	// Get the updated user image
	userImageView, err := h.appSvc.GetUserImage(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get updated user image: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User image updated successfully", httpresp.SuccessOptions{Data: respdto.UserImageROToDTO(&userImageView)})
}

func (h *UserImageHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := userImageAppCommands.DeleteUserImageCmd{UserImageID: req.ID}
	if err := h.appSvc.DeleteUserImage(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete user image: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "User image deleted successfully")
}
