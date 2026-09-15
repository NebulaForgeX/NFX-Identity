package http

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"nfxidentity/modules/asset/application/media"
	"nfxidentity/pkgs/errx"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/fiberx/middleware"
	"nfxidentity/pkgs/httpx"
	"nfxidentity/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/google/uuid"
)

type httpDeps interface {
	MediaSvc() *media.Service
	UserTokenVerifier() token.Verifier
}

func NewHTTPServer(d httpDeps, accessLog httpx.AccessLogConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: fiberx.ErrorHandler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "X-Api-Key", "X-Request-ID"},
		MaxAge:       3600,
	}))
	app.Use(middleware.Logger(), middleware.AccessLog(accessLog), middleware.Recover())
	h := &Handler{svc: d.MediaSvc()}
	asset := app.Group("/asset")
	asset.Get("/locales/:lang", h.Locales)
	asset.Get("/messages/:lang", h.Locales)
	app.Get("/health", func(c fiber.Ctx) error {
		return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]string{"service": "asset"}})
	})
	for _, kind := range []string{media.KindImages, media.KindFiles, media.KindVideos, media.KindAudios} {
		k := kind
		g := asset.Group("/"+k, middleware.TokenAuth(d.UserTokenVerifier()))
		g.Get("/", h.list(k))
		g.Post("/upload-url", h.prepare(k))
		g.Post("/upload-urls", h.prepareMany(k))
		g.Post("/confirm", h.confirm(k))
		g.Post("/confirm-"+k, h.confirmMany(k))
		g.Delete("/:id", h.remove(k))
		asset.Get("/"+k+"/:id/file", h.file(k))
	}
	return app
}

type Handler struct {
	svc *media.Service
}

func wrap(c fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*errx.Error); ok {
		return fiberx.ErrorFromErrx(c, e)
	}
	return fiberx.ErrorFromErrx(c, errx.Internal("INTERNAL", err.Error()))
}

func (h *Handler) Locales(c fiber.Ctx) error {
	lang := c.Params("lang")
	path := filepath.Join("errors", "langs", lang+".json")
	data, err := os.ReadFile(path)
	if err != nil {
		data, err = os.ReadFile(filepath.Join("errors", "langs", "en.json"))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(map[string]string{"message": "lang not found"})
		}
	}
	c.Set(fiber.HeaderContentType, "application/json; charset=utf-8")
	return c.Send(data)
}

func accountID(c fiber.Ctx) (uuid.UUID, error) {
	aid, ok := fiberx.AccountIDFromContext(c.Context())
	if !ok {
		return uuid.Nil, errx.Unauthorized("INVALID_TOKEN", "missing account")
	}
	return aid, nil
}

func (h *Handler) list(kind string) fiber.Handler {
	return func(c fiber.Ctx) error {
		aid, err := accountID(c)
		if err != nil {
			return wrap(c, err)
		}
		rows, err := h.svc.List(c.Context(), aid, kind)
		if err != nil {
			return wrap(c, err)
		}
		return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: rows})
	}
}

func (h *Handler) prepare(kind string) fiber.Handler {
	return func(c fiber.Ctx) error {
		aid, err := accountID(c)
		if err != nil {
			return wrap(c, err)
		}
		var req struct {
			FileName string `json:"file_name"`
			MimeType string `json:"mime_type"`
		}
		if err := c.Bind().Body(&req); err != nil {
			return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_BODY", "invalid body"))
		}
		out, err := h.svc.Prepare(c.Context(), aid, kind, req.FileName, req.MimeType)
		if err != nil {
			return wrap(c, err)
		}
		return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: out})
	}
}

func (h *Handler) prepareMany(kind string) fiber.Handler {
	return func(c fiber.Ctx) error {
		aid, err := accountID(c)
		if err != nil {
			return wrap(c, err)
		}
		var req struct {
			Items []struct {
				FileName string `json:"file_name"`
				MimeType string `json:"mime_type"`
			} `json:"items"`
		}
		if err := c.Bind().Body(&req); err != nil || len(req.Items) == 0 {
			return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_BODY", "invalid body"))
		}
		out := make([]*media.PrepareResult, 0, len(req.Items))
		for _, item := range req.Items {
			one, err := h.svc.Prepare(c.Context(), aid, kind, item.FileName, item.MimeType)
			if err != nil {
				return wrap(c, err)
			}
			out = append(out, one)
		}
		return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]any{"results": out}})
	}
}

func (h *Handler) confirm(kind string) fiber.Handler {
	return func(c fiber.Ctx) error {
		aid, err := accountID(c)
		if err != nil {
			return wrap(c, err)
		}
		var req struct {
			ID string `json:"id"`
		}
		if err := c.Bind().Body(&req); err != nil {
			return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_BODY", "invalid body"))
		}
		if err := h.svc.Confirm(c.Context(), aid, kind, req.ID); err != nil {
			return wrap(c, err)
		}
		return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
	}
}

func (h *Handler) confirmMany(kind string) fiber.Handler {
	return func(c fiber.Ctx) error {
		aid, err := accountID(c)
		if err != nil {
			return wrap(c, err)
		}
		var req struct {
			IDs []string `json:"ids"`
		}
		if err := c.Bind().Body(&req); err != nil || len(req.IDs) == 0 {
			return fiberx.ErrorFromErrx(c, errx.InvalidArg("INVALID_BODY", "invalid body"))
		}
		for _, id := range req.IDs {
			if err := h.svc.Confirm(c.Context(), aid, kind, id); err != nil {
				return wrap(c, err)
			}
		}
		return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
	}
}

func (h *Handler) remove(kind string) fiber.Handler {
	return func(c fiber.Ctx) error {
		aid, err := accountID(c)
		if err != nil {
			return wrap(c, err)
		}
		if err := h.svc.Delete(c.Context(), aid, kind, c.Params("id")); err != nil {
			return wrap(c, err)
		}
		return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: nil})
	}
}

func (h *Handler) file(kind string) fiber.Handler {
	return func(c fiber.Ctx) error {
		rc, mime, name, err := h.svc.Open(c.Context(), kind, c.Params("id"))
		if err != nil {
			return wrap(c, err)
		}
		defer rc.Close()
		c.Set("Content-Type", mime)
		c.Set("Content-Disposition", "inline; filename=\""+name+"\"")
		_, copyErr := io.Copy(c.Response().BodyWriter(), rc)
		return copyErr
	}
}
