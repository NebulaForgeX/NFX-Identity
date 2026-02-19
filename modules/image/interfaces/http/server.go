package http

import (
	"encoding/json"
	"time"

	imageTagApp "nfxid/modules/image/application/image_tags"
	imageTypeApp "nfxid/modules/image/application/image_types"
	imageVariantApp "nfxid/modules/image/application/image_variants"
	imageApp "nfxid/modules/image/application/images"
	"nfxid/modules/image/interfaces/http/handler"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/fiberx/middleware"
	"nfxid/pkgs/security/token"

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

func NewHTTPServer(d httpDeps) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		ErrorHandler:  fiberx.ErrorHandler,
		ReadTimeout:   30 * time.Second,
		WriteTimeout:  30 * time.Second,
		IdleTimeout:   120 * time.Second,
		BodyLimit:     10 * 1024 * 1024, // 10MB for file uploads
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "X-Api-Key", "X-Request-ID"},
		AllowCredentials: false,
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		MaxAge:           3600,
	}))

	app.Use(middleware.Logger(), middleware.Recover())

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
