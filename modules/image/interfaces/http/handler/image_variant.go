package handler

import (
	imageVariantApp "nfxid/modules/image/application/image_variants"
	imageVariantAppCommands "nfxid/modules/image/application/image_variants/commands"
	"nfxid/modules/image/interfaces/http/dto/reqdto"
	"nfxid/modules/image/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type ImageVariantHandler struct {
	appSvc *imageVariantApp.Service
}

func NewImageVariantHandler(appSvc *imageVariantApp.Service) *ImageVariantHandler {
	return &ImageVariantHandler{
		appSvc: appSvc,
	}
}

// Create 创建图片变体
func (h *ImageVariantHandler) Create(c fiber.Ctx) error {
	var req reqdto.ImageVariantCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	imageVariantID, err := h.appSvc.CreateImageVariant(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created image variant
	imageVariantView, err := h.appSvc.GetImageVariant(c.Context(), imageVariantID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Image variant created successfully", httpx.SuccessOptions{Data: respdto.ImageVariantROToDTO(&imageVariantView)})
}

// GetByID 根据 ID 获取图片变体
func (h *ImageVariantHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ImageVariantByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetImageVariant(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Image variant retrieved successfully", httpx.SuccessOptions{Data: respdto.ImageVariantROToDTO(&result)})
}

// Update 更新图片变体
func (h *ImageVariantHandler) Update(c fiber.Ctx) error {
	var req reqdto.ImageVariantByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	var updateReq reqdto.ImageVariantUpdateRequestDTO
	if err := c.Bind().Body(&updateReq); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := updateReq.ToUpdateCmd(req.ID)
	if err := h.appSvc.UpdateImageVariant(c.Context(), cmd); err != nil {
		return err
	}

	// Get the updated image variant
	imageVariantView, err := h.appSvc.GetImageVariant(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Image variant updated successfully", httpx.SuccessOptions{Data: respdto.ImageVariantROToDTO(&imageVariantView)})
}

// Delete 删除图片变体
func (h *ImageVariantHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ImageVariantByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := imageVariantAppCommands.DeleteImageVariantCmd{ImageVariantID: req.ID}
	if err := h.appSvc.DeleteImageVariant(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Image variant deleted successfully")
}

// fiber:context-methods migrated
