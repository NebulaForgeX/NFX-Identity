package handler

import (
	userImageApp "nfxid/modules/directory/application/user_images"
	userImageAppCommands "nfxid/modules/directory/application/user_images/commands"
	"nfxid/modules/directory/interfaces/http/dto/reqdto"
	"nfxid/modules/directory/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type UserImageHandler struct {
	appSvc *userImageApp.Service
}

func NewUserImageHandler(appSvc *userImageApp.Service) *UserImageHandler {
	return &UserImageHandler{appSvc: appSvc}
}

func (h *UserImageHandler) Create(c fiber.Ctx) error {
	var req reqdto.UserImageCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	userImageID, err := h.appSvc.CreateUserImage(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created user image
	userImageView, err := h.appSvc.GetUserImage(c.Context(), userImageID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "User image created successfully", httpx.SuccessOptions{Data: respdto.UserImageROToDTO(&userImageView)})
}

func (h *UserImageHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetUserImage(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "User image retrieved successfully", httpx.SuccessOptions{Data: respdto.UserImageROToDTO(&result)})
}

func (h *UserImageHandler) GetByUserID(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	results, err := h.appSvc.GetUserImagesByUserID(c.Context(), req.ID)
	if err != nil {
		return err
	}

	dtos := respdto.UserImageListROToDTO(results)
	return fiberx.OK(c, "User images retrieved successfully", httpx.SuccessOptions{Data: dtos})
}

func (h *UserImageHandler) GetCurrent(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetCurrentUserImageByUserID(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Current user image retrieved successfully", httpx.SuccessOptions{Data: respdto.UserImageROToDTO(&result)})
}

func (h *UserImageHandler) SetPrimary(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := userImageAppCommands.SetPrimaryUserImageCmd{UserImageID: req.ID}
	if err := h.appSvc.SetPrimaryUserImage(c.Context(), cmd); err != nil {
		return err
	}

	userImageView, err := h.appSvc.GetUserImage(c.Context(), req.ID)
	if err != nil {
		return err
	}
	return fiberx.OK(c, "Primary user image set successfully", httpx.SuccessOptions{Data: respdto.UserImageROToDTO(&userImageView)})
}

func (h *UserImageHandler) UpdateDisplayOrder(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	var updateReq reqdto.UserImageUpdateDisplayOrderRequestDTO
	if err := c.Bind().Body(&updateReq); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := updateReq.ToUpdateDisplayOrderCmd(req.ID)
	if err := h.appSvc.UpdateUserImageDisplayOrder(c.Context(), cmd); err != nil {
		return err
	}

	// Get the updated user image
	userImageView, err := h.appSvc.GetUserImage(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "User image display order updated successfully", httpx.SuccessOptions{Data: respdto.UserImageROToDTO(&userImageView)})
}

func (h *UserImageHandler) UpdateDisplayOrderBatch(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}
	userID := req.ID

	var body reqdto.UserImagesDisplayOrderBatchRequestDTO
	if err := c.Bind().Body(&body); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := body.ToBatchUpdateDisplayOrderCmd(userID)
	if err := h.appSvc.UpdateUserImagesDisplayOrderBatch(c.Context(), cmd); err != nil {
		return err
	}

	results, err := h.appSvc.GetUserImagesByUserID(c.Context(), userID)
	if err != nil {
		return err
	}
	return fiberx.OK(c, "User images display order updated successfully", httpx.SuccessOptions{Data: respdto.UserImageListROToDTO(results)})
}

func (h *UserImageHandler) Update(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	var updateReq reqdto.UserImageUpdateImageIDRequestDTO
	if err := c.Bind().Body(&updateReq); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := updateReq.ToUpdateImageIDCmd(req.ID)
	if err := h.appSvc.UpdateUserImageImageID(c.Context(), cmd); err != nil {
		return err
	}

	// Get the updated user image
	userImageView, err := h.appSvc.GetUserImage(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "User image updated successfully", httpx.SuccessOptions{Data: respdto.UserImageROToDTO(&userImageView)})
}

func (h *UserImageHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := userImageAppCommands.DeleteUserImageCmd{UserImageID: req.ID}
	if err := h.appSvc.DeleteUserImage(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "User image deleted successfully")
}

// fiber:context-methods migrated
