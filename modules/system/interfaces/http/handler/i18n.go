package handler

import (
	"os"
	"path/filepath"

	"nfxid/pkgs/errx"

	"github.com/gofiber/fiber/v3"
)

// Supported error translation languages (filenames: en.json, zh.json, fr.json).
var supportedLangs = map[string]bool{"en": true, "zh": true, "fr": true}

// I18nHandler serves error code translation JSON from a mounted directory (e.g. ./data/errors/langs).
type I18nHandler struct {
	errorsLangsPath string
}

func NewI18nHandler(errorsLangsPath string) *I18nHandler {
	return &I18nHandler{errorsLangsPath: errorsLangsPath}
}

// GetErrorTranslations returns the error code translations JSON for the given lang (en, zh, fr).
// Reads from disk each time so external updates are visible immediately.
// GET /system/i18n/errors/:lang
func (h *I18nHandler) GetErrorTranslations(c fiber.Ctx) error {
	lang := c.Params("lang")
	if lang == "" || !supportedLangs[lang] {
		return errx.ErrInvalidParams.WithMsg("lang must be one of: en, zh, fr")
	}
	name := lang + ".json"
	fpath := filepath.Join(h.errorsLangsPath, name)
	data, err := os.ReadFile(fpath)
	if err != nil {
		if os.IsNotExist(err) {
			return errx.NotFound("NOT_FOUND", "translation file not found: "+name)
		}
		return errx.ErrInternal.WithCause(err)
	}
	c.Set("Content-Type", "application/json; charset=utf-8")
	return c.Send(data)
}
