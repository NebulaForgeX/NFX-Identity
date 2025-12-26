package http

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app      fiber.Router
	handlers *Registry
}

func NewRouter(app fiber.Router, handlers *Registry) *Router {
	return &Router{
		app:      app,
		handlers: handlers,
	}
}

func (r *Router) RegisterRoutes() {
	image := r.app.Group("/image")

	// ========== 图片相关路由 ==========
	images := image.Group("/images")
	{
		images.Post("", r.handlers.Image.Create)
		images.Get("", r.handlers.Image.GetAll)
		images.Get("/:id", r.handlers.Image.GetByID)
		images.Put("/:id", r.handlers.Image.Update)
		images.Delete("/:id", r.handlers.Image.Delete)
	}

	// ========== 图片类型相关路由 ==========
	imageTypes := image.Group("/image-types")
	{
		imageTypes.Post("", r.handlers.ImageType.Create)
		imageTypes.Get("", r.handlers.ImageType.GetAll)
		imageTypes.Get("/:id", r.handlers.ImageType.GetByID)
		imageTypes.Get("/key/:key", r.handlers.ImageType.GetByKey)
		imageTypes.Put("/:id", r.handlers.ImageType.Update)
		imageTypes.Delete("/:id", r.handlers.ImageType.Delete)
	}
}

