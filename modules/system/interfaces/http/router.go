package http

import (
	"nfxid/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
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

	// 公开路由（system-base 等服务会调用）
	system.Get("/system-state/latest", r.handlers.SystemState.GetLatest)
	system.Post("/system-state/initialize", r.handlers.SystemState.Initialize)
}
