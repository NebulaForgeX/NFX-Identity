package http

import (
	"nfxid/pkgs/fiberx/middleware"
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
	access := r.app.Group("/access")

	// 需要认证的路由（UserSecurityPage 展示 + 权限管理 UI：角色/权限/角色-权限 CRUD）
	auth := access.Group("/auth", middleware.TokenAuth(r.tokenVerifier))
	{
		auth.Get("/grants", r.handlers.Grant.GetBySubject)
		auth.Get("/grants/:id", r.handlers.Grant.GetByID)
		// 角色（key 路由在前，避免被 :id 匹配）
		auth.Post("/roles", r.handlers.Role.Create)
		auth.Get("/roles/key/:key", r.handlers.Role.GetByKey)
		auth.Get("/roles/:id", r.handlers.Role.GetByID)
		auth.Put("/roles/:id", r.handlers.Role.Update)
		auth.Delete("/roles/:id", r.handlers.Role.Delete)
		// 权限
		auth.Post("/permissions", r.handlers.Permission.Create)
		auth.Get("/permissions/key/:key", r.handlers.Permission.GetByKey)
		auth.Get("/permissions/:id", r.handlers.Permission.GetByID)
		auth.Put("/permissions/:id", r.handlers.Permission.Update)
		auth.Delete("/permissions/:id", r.handlers.Permission.Delete)
		// 角色-权限关联
		auth.Post("/role-permissions", r.handlers.RolePermission.Create)
		auth.Get("/role-permissions/role/:role_id", r.handlers.RolePermission.GetByRoleID)
		auth.Get("/role-permissions/:id", r.handlers.RolePermission.GetByID)
		auth.Delete("/role-permissions/:id", r.handlers.RolePermission.Delete)
		// Action
		auth.Post("/actions", r.handlers.Action.Create)
		auth.Get("/actions/key/:key", r.handlers.Action.GetByKey)
		auth.Get("/actions/:id", r.handlers.Action.GetByID)
		// ActionRequirement（Permission 关联的 Action，用于配置 action permission）
		auth.Post("/action-requirements", r.handlers.ActionRequirement.Create)
		auth.Get("/action-requirements/permission/:permission_id", r.handlers.ActionRequirement.GetByPermissionID)
		auth.Get("/action-requirements/:id", r.handlers.ActionRequirement.GetByID)
		auth.Delete("/action-requirements/:id", r.handlers.ActionRequirement.Delete)
	}
}
