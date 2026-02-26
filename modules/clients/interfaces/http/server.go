package http

import (
	"encoding/json"
	"time"

	apiKeyApp "nfxidentity/modules/clients/application/api_keys"
	appApp "nfxidentity/modules/clients/application/apps"
	clientCredentialApp "nfxidentity/modules/clients/application/client_credentials"
	clientScopeApp "nfxidentity/modules/clients/application/client_scopes"
	ipAllowlistApp "nfxidentity/modules/clients/application/ip_allowlist"
	rateLimitApp "nfxidentity/modules/clients/application/rate_limits"
	"nfxidentity/modules/clients/interfaces/http/handler"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/fiberx/middleware"
	"nfxidentity/pkgs/httpx"
	"nfxidentity/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
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

func NewHTTPServer(d httpDeps, accessLog httpx.AccessLogConfig) *fiber.App {
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

	reg := &Registry{
		App:              handler.NewAppHandler(d.AppAppSvc()),
		APIKey:           handler.NewAPIKeyHandler(d.APIKeyAppSvc()),
		ClientCredential: handler.NewClientCredentialHandler(d.ClientCredentialAppSvc()),
		ClientScope:      handler.NewClientScopeHandler(d.ClientScopeAppSvc()),
		IPAllowlist:      handler.NewIPAllowlistHandler(d.IPAllowlistAppSvc()),
		RateLimit:        handler.NewRateLimitHandler(d.RateLimitAppSvc()),
	}

	router := NewRouter(app, d.UserTokenVerifier(), reg)
	router.RegisterRoutes()

	return app
}
