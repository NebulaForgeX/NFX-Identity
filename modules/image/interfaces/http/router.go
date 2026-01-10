package http

import (
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/usertoken"

	"github.com/gofiber/fiber/v2"
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

	// 需要认证的路由（需要token）
	auth := image.Group("/auth", usertoken.AccessTokenMiddleware(r.tokenVerifier))
	{
		// 图片相关
		auth.Post("/images", r.handlers.Image.Create)
		auth.Get("/images/:id", r.handlers.Image.GetByID)
		auth.Put("/images/:id", r.handlers.Image.Update)
		auth.Delete("/images/:id", r.handlers.Image.Delete)

		// 图片类型相关
		auth.Post("/image-types", r.handlers.ImageType.Create)
		auth.Get("/image-types/:id", r.handlers.ImageType.GetByID)
		auth.Put("/image-types/:id", r.handlers.ImageType.Update)
		auth.Delete("/image-types/:id", r.handlers.ImageType.Delete)

		// 图片变体相关
		auth.Post("/image-variants", r.handlers.ImageVariant.Create)
		auth.Get("/image-variants/:id", r.handlers.ImageVariant.GetByID)
		auth.Put("/image-variants/:id", r.handlers.ImageVariant.Update)
		auth.Delete("/image-variants/:id", r.handlers.ImageVariant.Delete)

		// 图片标签相关
		auth.Post("/image-tags", r.handlers.ImageTag.Create)
		auth.Get("/image-tags/:id", r.handlers.ImageTag.GetByID)
		auth.Put("/image-tags/:id", r.handlers.ImageTag.Update)
		auth.Delete("/image-tags/:id", r.handlers.ImageTag.Delete)
	}
}
