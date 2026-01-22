package handler

import (
	imageTypeApp "nfxid/modules/image/application/image_types"
	imageTypeAppCommands "nfxid/modules/image/application/image_types/commands"
	"nfxid/modules/image/interfaces/http/dto/reqdto"
	"nfxid/modules/image/interfaces/http/dto/respdto"
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
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
func (h *ImageTypeHandler) Create(c *fiber.Ctx) error {
	var req reqdto.ImageTypeCreateRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToCreateCmd()
	imageTypeID, err := h.appSvc.CreateImageType(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create image type: "+err.Error())
	}

	// Get the created image type
	imageTypeView, err := h.appSvc.GetImageType(c.Context(), imageTypeID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created image type: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Image type created successfully", httpresp.SuccessOptions{Data: respdto.ImageTypeROToDTO(&imageTypeView)})
}

// GetByID 根据 ID 获取图片类型
func (h *ImageTypeHandler) GetByID(c *fiber.Ctx) error {
	var req reqdto.ImageTypeByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetImageType(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Image type not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image type retrieved successfully", httpresp.SuccessOptions{Data: respdto.ImageTypeROToDTO(&result)})
}

// Update 更新图片类型
func (h *ImageTypeHandler) Update(c *fiber.Ctx) error {
	var req reqdto.ImageTypeByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	var updateReq reqdto.ImageTypeUpdateRequestDTO
	if err := c.BodyParser(&updateReq); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := updateReq.ToUpdateCmd(req.ID)
	if err := h.appSvc.UpdateImageType(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update image type: "+err.Error())
	}

	// Get the updated image type
	imageTypeView, err := h.appSvc.GetImageType(c.Context(), req.ID)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get updated image type: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image type updated successfully", httpresp.SuccessOptions{Data: respdto.ImageTypeROToDTO(&imageTypeView)})
}

// Delete 删除图片类型
func (h *ImageTypeHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ImageTypeByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := imageTypeAppCommands.DeleteImageTypeCmd{ImageTypeID: req.ID}
	if err := h.appSvc.DeleteImageType(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete image type: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image type deleted successfully")
}
