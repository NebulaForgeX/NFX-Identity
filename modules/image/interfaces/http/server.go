package http

import (
	"encoding/json"

	"nfxid/modules/image/interfaces/http/handler"
	"nfxid/pkgs/recover"
	"nfxid/pkgs/security/token"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type httpDeps interface {
	// TODO: Add application services when application layer is created
	// ImageAppSvc() *imageApp.Service
	// ImageTypeAppSvc() *imageTypeApp.Service
	// ImageVariantAppSvc() *imageVariantApp.Service
	// ImageTagAppSvc() *imageTagApp.Service
	UserTokenVerifier() token.Verifier
}

func NewHTTPServer(d httpDeps) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// CORS 中间件 - 必须在其他中间件之前
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false,
		ExposeHeaders:    "Content-Length",
	}))

	app.Use(recover.RecoverMiddleware(), logger.New())

	// 创建handlers
	// TODO: Pass application services when available
	reg := &Registry{
		Image:        handler.NewImageHandler(),
		ImageType:    handler.NewImageTypeHandler(),
		ImageVariant: handler.NewImageVariantHandler(),
		ImageTag:     handler.NewImageTagHandler(),
	}

	// 注册路由
	router := NewRouter(app, d.UserTokenVerifier(), reg)
	router.RegisterRoutes()

	return app
}
