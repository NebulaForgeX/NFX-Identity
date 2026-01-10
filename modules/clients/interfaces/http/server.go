package http

import (
	"encoding/json"

	apiKeyApp "nfxid/modules/clients/application/api_keys"
	appApp "nfxid/modules/clients/application/apps"
	clientCredentialApp "nfxid/modules/clients/application/client_credentials"
	clientScopeApp "nfxid/modules/clients/application/client_scopes"
	ipAllowlistApp "nfxid/modules/clients/application/ip_allowlist"
	rateLimitApp "nfxid/modules/clients/application/rate_limits"
	"nfxid/modules/clients/interfaces/http/handler"
	"nfxid/pkgs/recover"
	"nfxid/pkgs/security/token"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type httpDeps interface {
	AppAppSvc() *appApp.Service
	APIKeyAppSvc() *apiKeyApp.Service
	ClientCredentialAppSvc() *clientCredentialApp.Service
	ClientScopeAppSvc() *clientScopeApp.Service
	IPAllowlistAppSvc() *ipAllowlistApp.Service
	RateLimitAppSvc() *rateLimitApp.Service
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
	reg := &Registry{
		App:              handler.NewAppHandler(d.AppAppSvc()),
		APIKey:           handler.NewAPIKeyHandler(d.APIKeyAppSvc()),
		ClientCredential: handler.NewClientCredentialHandler(d.ClientCredentialAppSvc()),
		ClientScope:      handler.NewClientScopeHandler(d.ClientScopeAppSvc()),
		IPAllowlist:      handler.NewIPAllowlistHandler(d.IPAllowlistAppSvc()),
		RateLimit:        handler.NewRateLimitHandler(d.RateLimitAppSvc()),
	}

	// 注册路由
	router := NewRouter(app, d.UserTokenVerifier(), reg)
	router.RegisterRoutes()

	return app
}
