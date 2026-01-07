package handler

import (
	imageApp "nfxid/modules/image/application/image"
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
	result, err := h.appSvc.CreateImage(c.Context(), cmd)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to create image: "+err.Error())
	}

	// Get the created image view
	imageView, err := h.appSvc.GetImage(c.Context(), result.ID())
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get created image: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusCreated, "Image created successfully", httpresp.SuccessOptions{Data: respdto.ImageViewToDTO(&imageView)})
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

	return httpresp.Success(c, fiber.StatusOK, "Image retrieved successfully", httpresp.SuccessOptions{Data: respdto.ImageViewToDTO(&result)})
}

// GetAll 获取图片列表
func (h *ImageHandler) GetAll(c *fiber.Ctx) error {
	var query reqdto.ImageQueryParamsDTO
	if err := c.QueryParser(&query); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid query parameters: "+err.Error())
	}

	appQuery := query.ToListQuery()
	domainQuery := appQuery.ToDomainListQuery()
	result, err := h.appSvc.GetImageList(c.Context(), domainQuery)
	if err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to get images: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Images retrieved successfully", httpresp.SuccessOptions{
		Data: httpresp.ToList(respdto.ImageListViewToDTO(result.Items), int(result.Total)),
	})
}

// Update 更新图片
func (h *ImageHandler) Update(c *fiber.Ctx) error {
	var req reqdto.ImageUpdateRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}
	if err := c.BodyParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	cmd := req.ToUpdateCmd()
	if err := h.appSvc.UpdateImage(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to update image: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image updated successfully")
}

// Delete 删除图片
func (h *ImageHandler) Delete(c *fiber.Ctx) error {
	var req reqdto.ImageByIDRequestDTO
	if err := c.ParamsParser(&req); err != nil {
		return httpresp.Error(c, fiber.StatusBadRequest, "Invalid request params: "+err.Error())
	}

	cmd := imageApp.DeleteImageCmd{ID: req.ID}
	if err := h.appSvc.DeleteImage(c.Context(), cmd); err != nil {
		return httpresp.Error(c, fiber.StatusInternalServerError, "Failed to delete image: "+err.Error())
	}

	return httpresp.Success(c, fiber.StatusOK, "Image deleted successfully")
}
