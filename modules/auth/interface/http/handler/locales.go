package handler

import (
	"path/filepath"

	"nfxidentity/pkgs/fiberx"

	"github.com/gofiber/fiber/v3"
)

type LocalesHandler struct{}

func NewLocalesHandler() *LocalesHandler {
	return &LocalesHandler{}
}

func (h *LocalesHandler) GetErrorsByLang(c fiber.Ctx) error {
	return fiberx.SendLangJSON(c, filepath.Join("errors", "langs"), c.Params("lang"))
}

func (h *LocalesHandler) GetMessagesByLang(c fiber.Ctx) error {
	return fiberx.SendLangJSON(c, filepath.Join("messages", "langs"), c.Params("lang"))
}
