package http

import (
	"nfxid/pkgs/security/token"

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
	// 暂无页面使用 tenants API，不注册路由
	_ = r.app.Group("/tenants")
}
