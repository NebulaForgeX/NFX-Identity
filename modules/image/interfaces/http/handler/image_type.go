package handler

import (
	imageTypeApp "nfxid/modules/image/application/image_types"
	imageTypeAppCommands "nfxid/modules/image/application/image_types/commands"
	"nfxid/modules/image/interfaces/http/dto/reqdto"
	"nfxid/modules/image/interfaces/http/dto/respdto"
	"nfxid/pkgs/errx"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/httpx"

	"github.com/gofiber/fiber/v3"
)

type ImageTypeHandler struct {
	appSvc *imageTypeApp.Service
}

func NewImageTypeHandler(appSvc *imageTypeApp.Service) *ImageTypeHandler {
	return &ImageTypeHandler{
		appSvc: appSvc,
	}
}

// Create 创建图片类型
func (h *ImageTypeHandler) Create(c fiber.Ctx) error {
	var req reqdto.ImageTypeCreateRequestDTO
	if err := c.Bind().Body(&req); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := req.ToCreateCmd()
	imageTypeID, err := h.appSvc.CreateImageType(c.Context(), cmd)
	if err != nil {
		return err
	}

	// Get the created image type
	imageTypeView, err := h.appSvc.GetImageType(c.Context(), imageTypeID)
	if err != nil {
		return err
	}

	return fiberx.Created(c, "Image type created successfully", httpx.SuccessOptions{Data: respdto.ImageTypeROToDTO(&imageTypeView)})
}

// GetByID 根据 ID 获取图片类型
func (h *ImageTypeHandler) GetByID(c fiber.Ctx) error {
	var req reqdto.ImageTypeByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	result, err := h.appSvc.GetImageType(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Image type retrieved successfully", httpx.SuccessOptions{Data: respdto.ImageTypeROToDTO(&result)})
}

// Update 更新图片类型
func (h *ImageTypeHandler) Update(c fiber.Ctx) error {
	var req reqdto.ImageTypeByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	var updateReq reqdto.ImageTypeUpdateRequestDTO
	if err := c.Bind().Body(&updateReq); err != nil {
		return errx.ErrInvalidBody.WithCause(err)
	}

	cmd := updateReq.ToUpdateCmd(req.ID)
	if err := h.appSvc.UpdateImageType(c.Context(), cmd); err != nil {
		return err
	}

	// Get the updated image type
	imageTypeView, err := h.appSvc.GetImageType(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return fiberx.OK(c, "Image type updated successfully", httpx.SuccessOptions{Data: respdto.ImageTypeROToDTO(&imageTypeView)})
}

// Delete 删除图片类型
func (h *ImageTypeHandler) Delete(c fiber.Ctx) error {
	var req reqdto.ImageTypeByIDRequestDTO
	if err := c.Bind().URI(&req); err != nil {
		return errx.ErrInvalidParams.WithCause(err)
	}

	cmd := imageTypeAppCommands.DeleteImageTypeCmd{ImageTypeID: req.ID}
	if err := h.appSvc.DeleteImageType(c.Context(), cmd); err != nil {
		return err
	}

	return fiberx.OK(c, "Image type deleted successfully")
}

// fiber:context-methods migrated
