package handler

import (
	imageTagApp "nfxid/modules/image/application/image_tags"
	imageTagAppCommands "nfxid/modules/image/application/image_tags/commands"
	"nfxid/modules/image/interfaces/http/dto/reqdto"
	"nfxid/modules/image/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
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
func (h *ImageTagHandler) Create(c *fiber.Ctx) error {
	var req reqdto.ImageTagCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	imageTagID, err := h.appSvc.CreateImageTag(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create image tag: "+err.Error())
	}

	// Get the created image tag
	imageTagView, err := h.appSvc.GetImageTag(c.Context(), imageTagID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created image tag: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Image tag created successfully", httpresp.SuccessOptions{Data: respdto.ImageTagROToDTO(&imageTagView)})
}

// GetByID 根据 ID 获取图片标签
func (h *ImageTagHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ImageTagByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetImageTag(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Image tag not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image tag retrieved successfully", httpresp.SuccessOptions{Data: respdto.ImageTagROToDTO(&result)})
}

// Update 更新图片标签
func (h *ImageTagHandler) Update(c *fiber.Ctx) error {
	var req reqdto.ImageTagByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	var updateReq reqdto.ImageTagUpdateRequestDTO
	if err := c.BodyParser(&updateReq); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := updateReq.ToUpdateCmd(req.ID)
	if err := h.appSvc.UpdateImageTag(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update image tag: "+err.Error())
	}

	// Get the updated image tag
	imageTagView, err := h.appSvc.GetImageTag(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get updated image tag: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image tag updated successfully", httpresp.SuccessOptions{Data: respdto.ImageTagROToDTO(&imageTagView)})
}

// Delete 删除图片标签
func (h *ImageTagHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ImageTagByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := imageTagAppCommands.DeleteImageTagCmd{ImageTagID: req.ID}
	if err := h.appSvc.DeleteImageTag(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete image tag: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image tag deleted successfully")
}
