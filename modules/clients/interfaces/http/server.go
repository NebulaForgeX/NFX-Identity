package http

import (
	"encoding/json"
	"time"

	apiKeyApp "nfxid/modules/clients/application/api_keys"
	appApp "nfxid/modules/clients/application/apps"
	clientCredentialApp "nfxid/modules/clients/application/client_credentials"
	clientScopeApp "nfxid/modules/clients/application/client_scopes"
	ipAllowlistApp "nfxid/modules/clients/application/ip_allowlist"
	rateLimitApp "nfxid/modules/clients/application/rate_limits"
	"nfxid/modules/clients/interfaces/http/handler"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/fiberx/middleware"
	"nfxid/pkgs/security/token"

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

func NewHTTPServer(d httpDeps) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		ErrorHandler:  fiberx.ErrorHandler,
		ReadTimeout:   30 * time.Second,
		WriteTimeout:  30 * time.Second,
		IdleTimeout:   120 * time.Second,
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
