package handler

import (
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type ImageVariantHandler struct {
	// TODO: Add application service when application layer is created
}

func NewImageVariantHandler(/* appSvc *imageVariantApp.Service */) *ImageVariantHandler {
	return &ImageVariantHandler{}
}

func (h *ImageVariantHandler) Create(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "ImageVariant application layer not implemented yet")
}

func (h *ImageVariantHandler) GetByID(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "ImageVariant application layer not implemented yet")
}

func (h *ImageVariantHandler) Update(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "ImageVariant application layer not implemented yet")
}

func (h *ImageVariantHandler) Delete(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "ImageVariant application layer not implemented yet")
}
