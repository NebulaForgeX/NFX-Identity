package handler

import (
	"nfxid/pkgs/netx/httpresp"

	"github.com/gofiber/fiber/v2"
)

type ImageTagHandler struct {
	// TODO: Add application service when application layer is created
}

func NewImageTagHandler(/* appSvc *imageTagApp.Service */) *ImageTagHandler {
	return &ImageTagHandler{}
}

func (h *ImageTagHandler) Create(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "ImageTag application layer not implemented yet")
}

func (h *ImageTagHandler) GetByID(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "ImageTag application layer not implemented yet")
}

func (h *ImageTagHandler) Update(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "ImageTag application layer not implemented yet")
}

func (h *ImageTagHandler) Delete(c *fiber.Ctx) error {
	return httpresp.Error(c, fiber.StatusNotImplemented, "ImageTag application layer not implemented yet")
}
