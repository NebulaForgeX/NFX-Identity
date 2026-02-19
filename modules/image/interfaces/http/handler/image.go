package handler

import (
	imageApp "nfxid/modules/image/application/images"
	imageAppCommands "nfxid/modules/image/application/images/commands"
	"nfxid/modules/image/interfaces/http/dto/reqdto"
	"nfxid/modules/image/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type ImageHandler struct {
	appSvc *imageApp.Service
}

func NewImageHandler(appSvc *imageApp.Service) *ImageHandler {
	return &ImageHandler{
		appSvc: appSvc,
	}
}

// Create 创建图片
func (h *ImageHandler) Create(c fiber.Ctx) error {
	var req reqdto.ImageCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	imageID, err := h.appSvc.CreateImage(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created image
	imageView, err := h.appSvc.GetImage(c.Context(), imageID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Image created successfully", httpx.SuccessOptions{Data: respdto.ImageROToDTO(&imageView)})
}

// GetByID 根据 ID 获取图片
func (h *ImageHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ImageByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetImage(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Image retrieved successfully", httpx.SuccessOptions{Data: respdto.ImageROToDTO(&result)})
}

// Update 更新图片
func (h *ImageHandler) Update(c fiber.Ctx) error {
	var req reqdto.ImageByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	var updateReq reqdto.ImageUpdateRequestDTO
	if err := c.Bind().Body(&updateReq); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := updateReq.ToUpdateCmd(req.ID)
	if err := h.appSvc.UpdateImage(c.Context(), cmd); err != nil {
		return err
	}

	// Get the updated image
	imageView, err := h.appSvc.GetImage(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Image updated successfully", httpx.SuccessOptions{Data: respdto.ImageROToDTO(&imageView)})
}

// Delete 删除图片
func (h *ImageHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ImageByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := imageAppCommands.DeleteImageCmd{ImageID: req.ID}
	if err := h.appSvc.DeleteImage(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Image deleted successfully")
}

// fiber:context-methods migrated
