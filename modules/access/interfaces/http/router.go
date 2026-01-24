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

	// 需要认证的路由（需要token）
	auth := access.Group("/auth", usertoken.AccessTokenMiddleware(r.tokenVerifier))
	{
		// 角色相关
		auth.Post("/roles", r.handlers.Role.Create)
		auth.Get("/roles/:id", r.handlers.Role.GetByID)
		auth.Get("/roles/key/:key", r.handlers.Role.GetByKey)
		auth.Put("/roles/:id", r.handlers.Role.Update)
		auth.Delete("/roles/:id", r.handlers.Role.Delete)

		// 权限相关
		auth.Post("/permissions", r.handlers.Permission.Create)
		auth.Get("/permissions/:id", r.handlers.Permission.GetByID)
		auth.Get("/permissions/key/:key", r.handlers.Permission.GetByKey)
		auth.Put("/permissions/:id", r.handlers.Permission.Update)
		auth.Delete("/permissions/:id", r.handlers.Permission.Delete)

		// 作用域相关
		auth.Post("/scopes", r.handlers.Scope.Create)
		auth.Get("/scopes/:scope", r.handlers.Scope.GetByScope)
		auth.Put("/scopes/:scope", r.handlers.Scope.Update)
		auth.Delete("/scopes/:scope", r.handlers.Scope.Delete)

		// 授权相关
		auth.Post("/grants", r.handlers.Grant.Create)
		auth.Get("/grants", r.handlers.Grant.GetBySubject)
		auth.Get("/grants/:id", r.handlers.Grant.GetByID)
		auth.Put("/grants/:id", r.handlers.Grant.Update)
		auth.Delete("/grants/:id", r.handlers.Grant.Delete)

		// 角色权限关联相关
		auth.Post("/role-permissions", r.handlers.RolePermission.Create)
		auth.Get("/role-permissions/:id", r.handlers.RolePermission.GetByID)
		auth.Delete("/role-permissions/:id", r.handlers.RolePermission.Delete)

		// 作用域权限关联相关
		auth.Post("/scope-permissions", r.handlers.ScopePermission.Create)
		auth.Get("/scope-permissions/:id", r.handlers.ScopePermission.GetByID)
		auth.Delete("/scope-permissions/:id", r.handlers.ScopePermission.Delete)
	}
}