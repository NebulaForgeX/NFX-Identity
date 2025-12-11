package http

import (
	imageApp "nfxid/modules/image/application/image"
	imageTypeApp "nfxid/modules/image/application/image_type"
	"nfxid/modules/image/interfaces/http/handler"
	"nfxid/pkgs/recover"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type httpDeps interface {
	ImageAppSvc() *imageApp.Service
	ImageTypeAppSvc() *imageTypeApp.Service
}

func NewHTTPServer(d httpDeps) *fiber.App {
	app := fiber.New()

	// CORS 中间件
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false,
		ExposeHeaders:    "Content-Length",
	}))

	app.Use(recover.RecoverMiddleware(), logger.New())

	// 创建 handlers
	reg := &Registry{
		Image:     handler.NewImageHandler(d.ImageAppSvc()),
		ImageType: handler.NewImageTypeHandler(d.ImageTypeAppSvc()),
	}

	// 注册路由
	router := NewRouter(app, reg)
	router.RegisterRoutes()

	return app
}
