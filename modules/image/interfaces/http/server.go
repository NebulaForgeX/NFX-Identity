package http

import (
	"encoding/json"

	imageApp "nfxid/modules/image/application/images"
	imageTagApp "nfxid/modules/image/application/image_tags"
	imageTypeApp "nfxid/modules/image/application/image_types"
	imageVariantApp "nfxid/modules/image/application/image_variants"
	"nfxid/modules/image/interfaces/http/handler"
	"nfxid/pkgs/recover"
	"nfxid/pkgs/security/token"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type httpDeps interface {
	ImageAppSvc() *imageApp.Service
	ImageTypeAppSvc() *imageTypeApp.Service
	ImageVariantAppSvc() *imageVariantApp.Service
	ImageTagAppSvc() *imageTagApp.Service
	UserTokenVerifier() token.Verifier
	StoragePath() string
}

func NewHTTPServer(d httpDeps) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		BodyLimit:   10 * 1024 * 1024, // 10MB for file uploads
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
	reg := &Registry{
		Image:        handler.NewImageHandler(d.ImageAppSvc()),
		ImageType:    handler.NewImageTypeHandler(d.ImageTypeAppSvc()),
		ImageVariant: handler.NewImageVariantHandler(d.ImageVariantAppSvc()),
		ImageTag:     handler.NewImageTagHandler(d.ImageTagAppSvc()),
		Upload:       handler.NewUploadHandler(d.ImageAppSvc(), d.StoragePath()),
	}

	// 注册路由
	router := NewRouter(app, d.UserTokenVerifier(), reg)
	router.RegisterRoutes()

	return app
}
