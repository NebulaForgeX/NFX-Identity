package http

import (
	"encoding/json"
	"time"

	imageTagApp "nfxidentity/modules/image/application/image_tags"
	imageTypeApp "nfxidentity/modules/image/application/image_types"
	imageVariantApp "nfxidentity/modules/image/application/image_variants"
	imageApp "nfxidentity/modules/image/application/images"
	"nfxidentity/modules/image/interfaces/http/handler"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/fiberx/middleware"
	"nfxidentity/pkgs/httpx"
	"nfxidentity/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

type httpDeps interface {
	ImageAppSvc() *imageApp.Service
	ImageTypeAppSvc() *imageTypeApp.Service
	ImageVariantAppSvc() *imageVariantApp.Service
	ImageTagAppSvc() *imageTagApp.Service
	UserTokenVerifier() token.Verifier
	StoragePath() string
}

func NewHTTPServer(d httpDeps, accessLog httpx.AccessLogConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: fiberx.ErrorHandler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
		BodyLimit:    10 * 1024 * 1024, // 10MB for file uploads
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "X-Api-Key", "X-Request-ID"},
		AllowCredentials: false,
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		MaxAge:           3600,
	}))

	app.Use(middleware.Logger(), middleware.AccessLog(accessLog), middleware.Recover())

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
