package handler

import (
	imageTypeApp "nebulaid/modules/image/application/image_type"
	"nebulaid/modules/image/interfaces/http/dto/reqdto"
	"nebulaid/modules/image/interfaces/http/dto/respdto"
	"nebulaid/pkgs/netx/httpresp"

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
	result, err := h.appSvc.CreateImageType(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create image type: "+err.Error())
	}

	// Get the created image type view
	imageTypeView, err := h.appSvc.GetImageType(c.Context(), result.ID())
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created image type: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Image type created successfully", httpresp.SuccessOptions{Data: respdto.ImageTypeViewToDTO(&imageTypeView)})
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

	return httpresp.Success(c, fiber.StatusOK, "Image type retrieved successfully", httpresp.SuccessOptions{Data: respdto.ImageTypeViewToDTO(&result)})
}

// GetByKey 根据 Key 获取图片类型
func (h *ImageTypeHandler) GetByKey(c *fiber.Ctx) error {
	var req reqdto.ImageTypeByKeyRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	result, err := h.appSvc.GetImageTypeByKey(c.Context(), req.Key)
	if err != nil {
		return httpresp.Error(c, fiber.StatusNotFound, "Image type not found: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image type retrieved successfully", httpresp.SuccessOptions{Data: respdto.ImageTypeViewToDTO(&result)})
}

// GetAll 获取图片类型列表
func (h *ImageTypeHandler) GetAll(c *fiber.Ctx) error {
	var query reqdto.ImageTypeQueryParamsDTO
	if err := c.QueryParser(&query); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid query parameters: "+err.Error())
	}

	listQuery := query.ToListQuery()
	result, err := h.appSvc.GetImageTypeList(c.Context(), listQuery)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get image types: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image types retrieved successfully", httpresp.SuccessOptions{
		Data: httpresp.ToList(respdto.ImageTypeListViewToDTO(result.Items), int(result.Total)),
	})
}

// Update 更新图片类型
func (h *ImageTypeHandler) Update(c *fiber.Ctx) error {
	var req reqdto.ImageTypeUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateImageType(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update image type: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image type updated successfully")
}

// Delete 删除图片类型
func (h *ImageTypeHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ImageTypeByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := imageTypeApp.DeleteImageTypeCmd{ID: req.ID}
	if err := h.appSvc.DeleteImageType(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete image type: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image type deleted successfully")
}
