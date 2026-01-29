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
	access := r.app.Group("/access")

	// 需要认证的路由（仅注册页面用到的：UserSecurityPage 展示授权/角色/权限）
	auth := access.Group("/auth", usertoken.AccessTokenMiddleware(r.tokenVerifier))
	{
		auth.Get("/grants", r.handlers.Grant.GetBySubject)
		auth.Get("/grants/:id", r.handlers.Grant.GetByID)
		auth.Get("/roles/:id", r.handlers.Role.GetByID)
		auth.Get("/permissions/:id", r.handlers.Permission.GetByID)
		auth.Get("/role-permissions/role/:role_id", r.handlers.RolePermission.GetByRoleID)
	}
}
