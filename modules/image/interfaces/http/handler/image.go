package handler

import (
	imageApp "nfxid/modules/image/application/images"
	imageAppCommands "nfxid/modules/image/application/images/commands"
	"nfxid/modules/image/interfaces/http/dto/reqdto"
	"nfxid/modules/image/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
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
func (h *ImageHandler) Create(c *fiber.Ctx) error {
	var req reqdto.ImageCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	imageID, err := h.appSvc.CreateImage(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create image: "+err.Error())
	}

	// Get the created image
	imageView, err := h.appSvc.GetImage(c.Context(), imageID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created image: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Image created successfully", httpresp.SuccessOptions{Data: respdto.ImageROToDTO(&imageView)})
}

// GetByID 根据 ID 获取图片
func (h *ImageHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ImageByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetImage(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Image not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image retrieved successfully", httpresp.SuccessOptions{Data: respdto.ImageROToDTO(&result)})
}

// Update 更新图片
func (h *ImageHandler) Update(c *fiber.Ctx) error {
	var req reqdto.ImageByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	var updateReq reqdto.ImageUpdateRequestDTO
	if err := c.BodyParser(&updateReq); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := updateReq.ToUpdateCmd(req.ID)
	if err := h.appSvc.UpdateImage(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update image: "+err.Error())
	}

	// Get the updated image
	imageView, err := h.appSvc.GetImage(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get updated image: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image updated successfully", httpresp.SuccessOptions{Data: respdto.ImageROToDTO(&imageView)})
}

// Delete 删除图片
func (h *ImageHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ImageByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := imageAppCommands.DeleteImageCmd{ImageID: req.ID}
	if err := h.appSvc.DeleteImage(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete image: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image deleted successfully")
}
