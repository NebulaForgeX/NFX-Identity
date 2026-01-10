package http

import (
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/usertoken"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app           fiber.Router
	tokenVerifier token.Verifier
	handlers      *Registry
}

func NewRouter(app fiber.Router, tokenVerifier token.Verifier, handlers *Registry) *Router {
	return &Router{
		app:           app,
		tokenVerifier: tokenVerifier,
		handlers:      handlers,
	}
}

func (r *Router) RegisterRoutes() {
	clients := r.app.Group("/clients")

	// 需要认证的路由（需要token）
	auth := clients.Group("/auth", usertoken.AccessTokenMiddleware(r.tokenVerifier))
	{
		// 应用相关
		auth.Post("/apps", r.handlers.App.Create)
		auth.Get("/apps/:id", r.handlers.App.GetByID)
		auth.Get("/apps/app-id/:app_id", r.handlers.App.GetByAppID)
		auth.Put("/apps/:id", r.handlers.App.Update)
		auth.Delete("/apps/:id", r.handlers.App.Delete)

		// API Key 相关
		auth.Post("/api-keys", r.handlers.APIKey.Create)
		auth.Get("/api-keys/:id", r.handlers.APIKey.GetByID)
		auth.Delete("/api-keys/key-id/:key_id", r.handlers.APIKey.Delete)

		// Client Credential 相关
		auth.Post("/client-credentials", r.handlers.ClientCredential.Create)
		auth.Get("/client-credentials/:id", r.handlers.ClientCredential.GetByID)
		auth.Delete("/client-credentials/client-id/:client_id", r.handlers.ClientCredential.Delete)

		// Client Scope 相关
		auth.Post("/client-scopes", r.handlers.ClientScope.Create)
		auth.Get("/client-scopes/:id", r.handlers.ClientScope.GetByID)
		auth.Delete("/client-scopes/:id", r.handlers.ClientScope.Delete)

		// IP Allowlist 相关
		auth.Post("/ip-allowlist", r.handlers.IPAllowlist.Create)
		auth.Get("/ip-allowlist/:id", r.handlers.IPAllowlist.GetByID)
		auth.Delete("/ip-allowlist/rule-id/:rule_id", r.handlers.IPAllowlist.Delete)

		// Rate Limit 相关
		auth.Post("/rate-limits", r.handlers.RateLimit.Create)
		auth.Get("/rate-limits/:id", r.handlers.RateLimit.GetByID)
		auth.Delete("/rate-limits/:id", r.handlers.RateLimit.Delete)
	}
}
