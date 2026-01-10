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
	system := r.app.Group("/system")

	// 需要认证的路由（需要token）
	auth := system.Group("/auth", usertoken.AccessTokenMiddleware(r.tokenVerifier))
	{
		// 系统状态相关
		auth.Get("/system-state/latest", r.handlers.SystemState.GetLatest)
		auth.Get("/system-state/:id", r.handlers.SystemState.GetByID)
		auth.Post("/system-state/initialize", r.handlers.SystemState.Initialize)
		auth.Post("/system-state/reset", r.handlers.SystemState.Reset)
		auth.Delete("/system-state/:id", r.handlers.SystemState.Delete)
	}
}
