package fiberx

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func SendLangJSON(c fiber.Ctx, dir, lang string) error {
	lang = strings.TrimSpace(lang)
	if lang == "" {
		lang = "en"
	}
	data, err := os.ReadFile(filepath.Join(dir, lang+".json"))
	if err != nil && lang != "en" {
		data, err = os.ReadFile(filepath.Join(dir, "en.json"))
	}
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(map[string]string{"message": "lang not found"})
	}
	c.Set(fiber.HeaderContentType, "application/json; charset=utf-8")
	return c.Send(data)
}
