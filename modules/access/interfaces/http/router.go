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
	auth := access.Group("/auth", middleware.TokenAuth(r.tokenVerifier))
	_ = auth
	if r.handlers.TenantRole != nil {
		auth.Get("/tenant-roles/tenant/:tenant_id", r.handlers.TenantRole.ListByTenantID)
		auth.Get("/tenant-roles/tenant/:tenant_id/role-key/:role_key", r.handlers.TenantRole.GetByTenantIDAndRoleKey)
		auth.Get("/tenant-roles/:id", r.handlers.TenantRole.GetByID)
		auth.Post("/tenant-roles", r.handlers.TenantRole.Create)
		auth.Put("/tenant-roles/:id", r.handlers.TenantRole.Update)
		auth.Delete("/tenant-roles/:id", r.handlers.TenantRole.DeleteByID)
	}
}
