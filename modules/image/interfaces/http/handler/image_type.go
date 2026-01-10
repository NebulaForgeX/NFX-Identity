package handler

import (
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type ImageTypeHandler struct {
	// TODO: Add application service when application layer is created
}

func NewImageTypeHandler(/* appSvc *imageTypeApp.Service */) *ImageTypeHandler {
	return &ImageTypeHandler{}
}

func (h *ImageTypeHandler) Create(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "ImageType application layer not implemented yet")
}

func (h *ImageTypeHandler) GetByID(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "ImageType application layer not implemented yet")
}

func (h *ImageTypeHandler) Update(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "ImageType application layer not implemented yet")
}

func (h *ImageTypeHandler) Delete(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "ImageType application layer not implemented yet")
}
