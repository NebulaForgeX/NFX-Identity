package http

import (
	"encoding/json"
	"time"

	domainVerificationApp "nfxidentity/modules/tenants/application/domain_verifications"
	groupApp "nfxidentity/modules/tenants/application/groups"
	invitationApp "nfxidentity/modules/tenants/application/invitations"
	memberAppRoleApp "nfxidentity/modules/tenants/application/member_app_roles"
	memberGroupApp "nfxidentity/modules/tenants/application/member_groups"
	memberRoleApp "nfxidentity/modules/tenants/application/member_roles"
	memberApp "nfxidentity/modules/tenants/application/members"
	tenantAppApp "nfxidentity/modules/tenants/application/tenant_apps"
	tenantSettingApp "nfxidentity/modules/tenants/application/tenant_settings"
	tenantApp "nfxidentity/modules/tenants/application/tenants"
	"nfxidentity/modules/tenants/interfaces/http/handler"
	"nfxidentity/pkgs/fiberx"
	"nfxidentity/pkgs/fiberx/middleware"
	"nfxidentity/pkgs/httpx"
	"nfxidentity/pkgs/security/token"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

type httpDeps interface {
	TenantAppSvc() *tenantApp.Service
	GroupAppSvc() *groupApp.Service
	MemberAppSvc() *memberApp.Service
	InvitationAppSvc() *invitationApp.Service
	TenantAppAppSvc() *tenantAppApp.Service
	TenantSettingAppSvc() *tenantSettingApp.Service
	DomainVerificationAppSvc() *domainVerificationApp.Service
	MemberRoleAppSvc() *memberRoleApp.Service
	MemberGroupAppSvc() *memberGroupApp.Service
	MemberAppRoleAppSvc() *memberAppRoleApp.Service
	UserTokenVerifier() token.Verifier
}

func NewHTTPServer(d httpDeps, accessLog httpx.AccessLogConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: fiberx.ErrorHandler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "X-Api-Key", "X-Request-ID"},
		AllowCredentials: false,
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		MaxAge:           3600,
	}))

	app.Use(middleware.Logger(), middleware.AccessLog(accessLog), middleware.Recover())

	reg := &Registry{
		Tenant:             handler.NewTenantHandler(d.TenantAppSvc()),
		Group:              handler.NewGroupHandler(d.GroupAppSvc()),
		Member:             handler.NewMemberHandler(d.MemberAppSvc()),
		Invitation:         handler.NewInvitationHandler(d.InvitationAppSvc()),
		TenantApp:          handler.NewTenantAppHandler(d.TenantAppAppSvc()),
		TenantSetting:      handler.NewTenantSettingHandler(d.TenantSettingAppSvc()),
		DomainVerification: handler.NewDomainVerificationHandler(d.DomainVerificationAppSvc()),
		MemberRole:         handler.NewMemberRoleHandler(d.MemberRoleAppSvc()),
		MemberGroup:        handler.NewMemberGroupHandler(d.MemberGroupAppSvc()),
		MemberAppRole:      handler.NewMemberAppRoleHandler(d.MemberAppRoleAppSvc()),
	}

	router := NewRouter(app, d.UserTokenVerifier(), reg)
	router.RegisterRoutes()

	return app
}
