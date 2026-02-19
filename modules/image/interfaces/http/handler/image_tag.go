package handler

import (
	imageTagApp "nfxid/modules/image/application/image_tags"
	imageTagAppCommands "nfxid/modules/image/application/image_tags/commands"
	"nfxid/modules/image/interfaces/http/dto/reqdto"
	"nfxid/modules/image/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type ImageTagHandler struct {
	appSvc *imageTagApp.Service
}

func NewImageTagHandler(appSvc *imageTagApp.Service) *ImageTagHandler {
	return &ImageTagHandler{
		appSvc: appSvc,
	}
}

// Create 创建图片标签
func (h *ImageTagHandler) Create(c fiber.Ctx) error {
	var req reqdto.ImageTagCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	imageTagID, err := h.appSvc.CreateImageTag(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created image tag
	imageTagView, err := h.appSvc.GetImageTag(c.Context(), imageTagID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Image tag created successfully", httpx.SuccessOptions{Data: respdto.ImageTagROToDTO(&imageTagView)})
}

// GetByID 根据 ID 获取图片标签
func (h *ImageTagHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ImageTagByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetImageTag(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Image tag retrieved successfully", httpx.SuccessOptions{Data: respdto.ImageTagROToDTO(&result)})
}

// Update 更新图片标签
func (h *ImageTagHandler) Update(c fiber.Ctx) error {
	var req reqdto.ImageTagByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	var updateReq reqdto.ImageTagUpdateRequestDTO
	if err := c.Bind().Body(&updateReq); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := updateReq.ToUpdateCmd(req.ID)
	if err := h.appSvc.UpdateImageTag(c.Context(), cmd); err != nil {
		return err
	}

	// Get the updated image tag
	imageTagView, err := h.appSvc.GetImageTag(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Image tag updated successfully", httpx.SuccessOptions{Data: respdto.ImageTagROToDTO(&imageTagView)})
}

// Delete 删除图片标签
func (h *ImageTagHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ImageTagByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := imageTagAppCommands.DeleteImageTagCmd{ImageTagID: req.ID}
	if err := h.appSvc.DeleteImageTag(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Image tag deleted successfully")
}

// fiber:context-methods migrated
