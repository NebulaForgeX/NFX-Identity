package http

import (
	"encoding/json"
	"time"

	authconn "nfxidentity/connections/auth"
	audiosApp "nfxidentity/modules/asset/application/audios"
	filesApp "nfxidentity/modules/asset/application/files"
	imagesApp "nfxidentity/modules/asset/application/images"
	videosApp "nfxidentity/modules/asset/application/videos"
	"nfxidentity/modules/asset/config"
	"nfxidentity/modules/asset/interface/http/handler"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/fiberx/middleware"
	"nfxidentity/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

type Registry struct {
	Locales *handler.LocalesHandler
	Image   *handler.ImageHandler
	Video   *handler.VideoHandler
	File    *handler.FileHandler
	Audio   *handler.AudioHandler
}

type httpDeps interface {
	ImagesApp() *imagesApp.Service
	VideosApp() *videosApp.Service
	FilesApp() *filesApp.Service
	AudiosApp() *audiosApp.Service
	AuthClient() *authconn.Client
	UserTokenVerifier() token.Verifier
}

func NewHTTPServer(d httpDeps, cfg *config.Config) *fiber.App {
	accessLog := cfg.Server.AccessLog
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: fiberx.ErrorHandler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
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

	auth := d.AuthClient()
	reg := &Registry{
		Locales: handler.NewLocalesHandler(),
		Image:   handler.NewImageHandler(d.ImagesApp(), auth),
		Video:   handler.NewVideoHandler(d.VideosApp(), auth),
		File:    handler.NewFileHandler(d.FilesApp(), auth),
		Audio:   handler.NewAudioHandler(d.AudiosApp(), auth),
	}
	NewRouter(app, reg, d.UserTokenVerifier()).RegisterRoutes()
	return app
}
