package handler

import (
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type ImageHandler struct {
	// TODO: Add application service when application layer is created
	// appSvc *imageApp.Service
}

func NewImageHandler(/* appSvc *imageApp.Service */) *ImageHandler {
	return &ImageHandler{
		// appSvc: appSvc,
	}
}

func (h *ImageHandler) Create(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Image application layer not implemented yet")
}

func (h *ImageHandler) GetByID(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Image application layer not implemented yet")
}

func (h *ImageHandler) Update(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Image application layer not implemented yet")
}

func (h *ImageHandler) Delete(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "Image application layer not implemented yet")
}
