package handler

import (
	imageVariantApp "nfxid/modules/image/application/image_variants"
	imageVariantAppCommands "nfxid/modules/image/application/image_variants/commands"
	"nfxid/modules/image/interfaces/http/dto/reqdto"
	"nfxid/modules/image/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
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
func (h *ImageVariantHandler) Create(c *fiber.Ctx) error {
	var req reqdto.ImageVariantCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	imageVariantID, err := h.appSvc.CreateImageVariant(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create image variant: "+err.Error())
	}

	// Get the created image variant
	imageVariantView, err := h.appSvc.GetImageVariant(c.Context(), imageVariantID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created image variant: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Image variant created successfully", httpresp.SuccessOptions{Data: respdto.ImageVariantROToDTO(&imageVariantView)})
}

// GetByID 根据 ID 获取图片变体
func (h *ImageVariantHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ImageVariantByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetImageVariant(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Image variant not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image variant retrieved successfully", httpresp.SuccessOptions{Data: respdto.ImageVariantROToDTO(&result)})
}

// Update 更新图片变体
func (h *ImageVariantHandler) Update(c *fiber.Ctx) error {
	var req reqdto.ImageVariantByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	var updateReq reqdto.ImageVariantUpdateRequestDTO
	if err := c.BodyParser(&updateReq); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := updateReq.ToUpdateCmd(req.ID)
	if err := h.appSvc.UpdateImageVariant(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update image variant: "+err.Error())
	}

	// Get the updated image variant
	imageVariantView, err := h.appSvc.GetImageVariant(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get updated image variant: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image variant updated successfully", httpresp.SuccessOptions{Data: respdto.ImageVariantROToDTO(&imageVariantView)})
}

// Delete 删除图片变体
func (h *ImageVariantHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ImageVariantByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := imageVariantAppCommands.DeleteImageVariantCmd{ImageVariantID: req.ID}
	if err := h.appSvc.DeleteImageVariant(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete image variant: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image variant deleted successfully")
}
