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
	tenants := r.app.Group("/tenants")

	// 需要认证的路由（需要token）
	auth := tenants.Group("/auth", usertoken.AccessTokenMiddleware(r.tokenVerifier))
	{
		// 租户相关
		auth.Post("/", r.handlers.Tenant.Create)
		auth.Get("/:id", r.handlers.Tenant.GetByID)
		auth.Get("/tenant-id/:tenant_id", r.handlers.Tenant.GetByTenantID)
		auth.Put("/:id", r.handlers.Tenant.Update)
		auth.Patch("/:id/status", r.handlers.Tenant.UpdateStatus)
		auth.Delete("/:id", r.handlers.Tenant.Delete)

		// 组相关
		auth.Post("/groups", r.handlers.Group.Create)
		auth.Get("/groups/:id", r.handlers.Group.GetByID)
		auth.Put("/groups/:id", r.handlers.Group.Update)
		auth.Delete("/groups/:id", r.handlers.Group.Delete)

		// 成员相关
		auth.Post("/members", r.handlers.Member.Create)
		auth.Get("/members/:id", r.handlers.Member.GetByID)
		auth.Put("/members/:id", r.handlers.Member.Update)
		auth.Delete("/members/:id", r.handlers.Member.Delete)

		// 邀请相关
		auth.Post("/invitations", r.handlers.Invitation.Create)
		auth.Get("/invitations/:id", r.handlers.Invitation.GetByID)
		auth.Get("/invitations/invite-id/:invite_id", r.handlers.Invitation.GetByInviteID)
		auth.Patch("/invitations/invite-id/:invite_id/accept", r.handlers.Invitation.Accept)
		auth.Patch("/invitations/invite-id/:invite_id/revoke", r.handlers.Invitation.Revoke)
		auth.Delete("/invitations/:id", r.handlers.Invitation.Delete)

		// 租户应用相关
		auth.Post("/tenant-apps", r.handlers.TenantApp.Create)
		auth.Get("/tenant-apps/:id", r.handlers.TenantApp.GetByID)
		auth.Put("/tenant-apps/:id", r.handlers.TenantApp.Update)
		auth.Delete("/tenant-apps/:id", r.handlers.TenantApp.Delete)

		// 租户设置相关
		auth.Post("/tenant-settings", r.handlers.TenantSetting.Create)
		auth.Get("/tenant-settings/:id", r.handlers.TenantSetting.GetByID)
		auth.Put("/tenant-settings/:id", r.handlers.TenantSetting.Update)
		auth.Delete("/tenant-settings/:id", r.handlers.TenantSetting.Delete)

		// 域名验证相关
		auth.Post("/domain-verifications", r.handlers.DomainVerification.Create)
		auth.Get("/domain-verifications/:id", r.handlers.DomainVerification.GetByID)
		auth.Put("/domain-verifications/:id", r.handlers.DomainVerification.Update)
		auth.Delete("/domain-verifications/:id", r.handlers.DomainVerification.Delete)

		// 成员角色相关
		auth.Post("/member-roles", r.handlers.MemberRole.Create)
		auth.Get("/member-roles/:id", r.handlers.MemberRole.GetByID)
		auth.Patch("/member-roles/:id/revoke", r.handlers.MemberRole.Update)
		auth.Delete("/member-roles/:id", r.handlers.MemberRole.Delete)

		// 成员组相关
		auth.Post("/member-groups", r.handlers.MemberGroup.Create)
		auth.Get("/member-groups/:id", r.handlers.MemberGroup.GetByID)
		auth.Patch("/member-groups/:id/revoke", r.handlers.MemberGroup.Update)
		auth.Delete("/member-groups/:id", r.handlers.MemberGroup.Delete)

		// 成员应用角色相关
		auth.Post("/member-app-roles", r.handlers.MemberAppRole.Create)
		auth.Get("/member-app-roles/:id", r.handlers.MemberAppRole.GetByID)
		auth.Patch("/member-app-roles/:id/revoke", r.handlers.MemberAppRole.Update)
		auth.Delete("/member-app-roles/:id", r.handlers.MemberAppRole.Delete)
	}
}
