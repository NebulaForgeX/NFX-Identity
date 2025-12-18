package http

import (
	"nfxid/modules/permission/interfaces/http/middleware"
	"nfxid/pkgs/tokenx"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app      fiber.Router
	tokenx   *tokenx.Tokenx
	handlers *Registry
}

func NewRouter(app fiber.Router, tokenx *tokenx.Tokenx, handlers *Registry) *Router {
	return &Router{
		app:      app,
		tokenx:   tokenx,
		handlers: handlers,
	}
}

func (r *Router) RegisterRoutes() {
	api := r.app.Group("/api/v1")

	// ========== 公开路由（不需要 token） ==========
	public := api.Group("/permission")
	{
		// 认证相关
		public.Post("/login", r.handlers.Auth.Login)
		public.Post("/register", r.handlers.Auth.Register) // 注册
	}

	// ========== 需要认证的路由（需要 token） ==========
	auth := api.Group("/permission", middleware.AccessTokenMiddleware(r.tokenx))
	{
		// 权限管理
		auth.Post("/permissions", r.handlers.Permission.Create)
		auth.Put("/permissions/:id", r.handlers.Permission.Update)
		auth.Delete("/permissions/:id", r.handlers.Permission.Delete)
		auth.Get("/permissions/:id", r.handlers.Permission.GetByID)
		auth.Get("/permissions/tag/:tag", r.handlers.Permission.GetByTag)
		auth.Get("/permissions", r.handlers.Permission.List)

		// 用户权限管理
		auth.Post("/user-permissions", r.handlers.UserPermission.Assign)
		auth.Delete("/user-permissions", r.handlers.UserPermission.Revoke)
		auth.Get("/users/:user_id/permissions", r.handlers.UserPermission.GetByUserID)
		auth.Get("/users/:user_id/permission-tags", r.handlers.UserPermission.GetTagsByUserID)
		auth.Post("/user-permissions/check", r.handlers.UserPermission.Check)

		// 授权码管理
		auth.Post("/authorization-codes", r.handlers.AuthorizationCode.Create)
		auth.Get("/authorization-codes/:id", r.handlers.AuthorizationCode.GetByID)
		auth.Get("/authorization-codes/code/:code", r.handlers.AuthorizationCode.GetByCode)
		auth.Post("/authorization-codes/use", r.handlers.AuthorizationCode.Use)
		auth.Delete("/authorization-codes/:id", r.handlers.AuthorizationCode.Delete)
		auth.Post("/authorization-codes/:id/activate", r.handlers.AuthorizationCode.Activate)
		auth.Post("/authorization-codes/:id/deactivate", r.handlers.AuthorizationCode.Deactivate)
	}
}
