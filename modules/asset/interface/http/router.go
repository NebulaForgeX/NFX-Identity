package http

import (
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/fiberx/middleware"
	"nfxidentity/pkgs/httpx"
	"nfxidentity/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
)

type Router struct {
	app      fiber.Router
	handlers *Registry
	verifier token.Verifier
}

func NewRouter(app fiber.Router, handlers *Registry, verifier token.Verifier) *Router {
	return &Router{app: app, handlers: handlers, verifier: verifier}
}

func (r *Router) RegisterRoutes() {
	asset := r.app.Group("/asset")
	r.RegisterLocalesGroup(asset)
	r.RegisterImagesGroup(asset)
	r.RegisterVideosGroup(asset)
	r.RegisterFilesGroup(asset)
	r.RegisterAudiosGroup(asset)
	r.app.Get("/health", func(c fiber.Ctx) error {
		return fiberx.OK(c, "ok", httpx.SuccessOptions{Data: map[string]string{"service": "asset"}})
	})
}

func (r *Router) RegisterLocalesGroup(asset fiber.Router) {
	locales := asset.Group("/locales")
	locales.Get("/:lang", r.handlers.Locales.GetErrorsByLang)
	messages := asset.Group("/messages")
	messages.Get("/:lang", r.handlers.Locales.GetMessagesByLang)
}

func (r *Router) RegisterImagesGroup(asset fiber.Router) {
	images := asset.Group("/images")
	images.Get("/:id/file", r.handlers.Image.ServeImageFile)

	me := images.Group("", middleware.TokenAuth(r.verifier))
	me.Get("/", r.handlers.Image.ListImages)
	me.Post("/upload-url", r.handlers.Image.PrepareImageUpload)
	me.Post("/upload-urls", r.handlers.Image.PrepareImagesUpload)
	me.Post("/confirm", r.handlers.Image.ConfirmImageUpload)
	me.Post("/confirm-images", r.handlers.Image.ConfirmImagesUpload)
	me.Delete("/:id", r.handlers.Image.DeleteImage)
}

func (r *Router) RegisterVideosGroup(asset fiber.Router) {
	videos := asset.Group("/videos")
	videos.Get("/:id/file", r.handlers.Video.ServeVideoFile)

	me := videos.Group("", middleware.TokenAuth(r.verifier))
	me.Get("/", r.handlers.Video.ListVideos)
	me.Post("/upload-url", r.handlers.Video.PrepareVideoUpload)
	me.Post("/upload-urls", r.handlers.Video.PrepareVideosUpload)
	me.Post("/confirm", r.handlers.Video.ConfirmVideoUpload)
	me.Post("/confirm-videos", r.handlers.Video.ConfirmVideosUpload)
	me.Delete("/:id", r.handlers.Video.DeleteVideo)
}

func (r *Router) RegisterFilesGroup(asset fiber.Router) {
	files := asset.Group("/files")
	files.Get("/:id/file", r.handlers.File.ServeFileFile)

	me := files.Group("", middleware.TokenAuth(r.verifier))
	me.Get("/", r.handlers.File.ListFiles)
	me.Post("/upload-url", r.handlers.File.PrepareFileUpload)
	me.Post("/upload-urls", r.handlers.File.PrepareFilesUpload)
	me.Post("/confirm", r.handlers.File.ConfirmFileUpload)
	me.Post("/confirm-files", r.handlers.File.ConfirmFilesUpload)
	me.Delete("/:id", r.handlers.File.DeleteFile)
}

func (r *Router) RegisterAudiosGroup(asset fiber.Router) {
	audios := asset.Group("/audios")
	audios.Get("/:id/file", r.handlers.Audio.ServeAudioFile)

	me := audios.Group("", middleware.TokenAuth(r.verifier))
	me.Get("/", r.handlers.Audio.ListAudios)
	me.Post("/upload-url", r.handlers.Audio.PrepareAudioUpload)
	me.Post("/upload-urls", r.handlers.Audio.PrepareAudiosUpload)
	me.Post("/confirm", r.handlers.Audio.ConfirmAudioUpload)
	me.Post("/confirm-audios", r.handlers.Audio.ConfirmAudiosUpload)
	me.Delete("/:id", r.handlers.Audio.DeleteAudio)
}
