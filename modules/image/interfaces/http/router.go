package http

import (
	"nfxid/pkgs/fiberx/middleware"
	"nfxid/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
)

type Router struct {
	app           fiber.Router
	tokenVerifier token.Verifier
	handlers      *Registry
}

func NewRouter(app fiber.Router, tokenVerifier token.Verifier, handlers *Registry) *Router {
	return &Router{
		app:           app,
		tokenVerifier: tokenVerifier,
		handlers:      handlers,
	}
}

func (r *Router) RegisterRoutes() {
	image := r.app.Group("/image")

	// 公开：按 ID 返回图片文件（头像/背景等 <img src> 用）
	image.Get("/public/images/:image_id", r.handlers.Upload.ServeImage)

	// 需要认证：上传
	auth := image.Group("/auth", middleware.TokenAuth(r.tokenVerifier))
	{
		auth.Post("/upload", r.handlers.Upload.UploadImage)
	}
}
