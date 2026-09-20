package handler

import (
	"nfxidentity/errors/src/sys"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v3"
)

// Supported error translation languages (filenames: en.json, zh.json, fr.json).
var supportedLangs = map[string]bool{"en": true, "zh": true, "fr": true}

// I18nHandler serves error code translation JSON from a mounted directory (e.g. ./data/errors/langs).
type I18nHandler struct {
	errorsLangsPath   string
	messagesLangsPath string
}

func NewI18nHandler(errorsLangsPath string) *I18nHandler {
	messagesLangsPath := strings.Replace(errorsLangsPath, "errors/langs", "messages/langs", 1)
	if messagesLangsPath == errorsLangsPath {
		messagesLangsPath = "./messages/langs"
	}
	return &I18nHandler{errorsLangsPath: errorsLangsPath, messagesLangsPath: messagesLangsPath}
}

// GetErrorTranslations GET /system/locales/:lang
func (h *I18nHandler) GetErrorTranslations(c fiber.Ctx) error {
	return h.sendLangJSON(c, h.errorsLangsPath)
}

// GetMessageTranslations GET /system/messages/:lang
func (h *I18nHandler) GetMessageTranslations(c fiber.Ctx) error {
	return h.sendLangJSON(c, h.messagesLangsPath)
}

func (h *I18nHandler) sendLangJSON(c fiber.Ctx, dir string) error {
	lang := c.Params("lang")
	if lang == "" || !supportedLangs[lang] {
		return sys.ErrInvalidParams.WithMsg("lang must be one of: en, zh, fr")
	}
	name := lang + ".json"
	fpath := filepath.Join(dir, name)
	data, err := os.ReadFile(fpath)
	if err != nil {
		if os.IsNotExist(err) {
			return sys.ErrNotFound.WithMsg("translation file not found: " + name)
		}
		return sys.ErrInternal.WithCause(err)
	}
	c.Set("Content-Type", "application/json; charset=utf-8")
	return c.Send(data)
}
