package http

import (
	"encoding/json"
	"time"

	domainVerificationApp "nfxid/modules/tenants/application/domain_verifications"
	groupApp "nfxid/modules/tenants/application/groups"
	invitationApp "nfxid/modules/tenants/application/invitations"
	memberAppRoleApp "nfxid/modules/tenants/application/member_app_roles"
	memberGroupApp "nfxid/modules/tenants/application/member_groups"
	memberRoleApp "nfxid/modules/tenants/application/member_roles"
	memberApp "nfxid/modules/tenants/application/members"
	tenantAppApp "nfxid/modules/tenants/application/tenant_apps"
	tenantSettingApp "nfxid/modules/tenants/application/tenant_settings"
	tenantApp "nfxid/modules/tenants/application/tenants"
	"nfxid/modules/tenants/interfaces/http/handler"
	"nfxid/pkgs/fiberx"
	"nfxid/pkgs/fiberx/middleware"
	"nfxid/pkgs/security/token"

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

func NewHTTPServer(d httpDeps) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		ErrorHandler:  fiberx.ErrorHandler,
		ReadTimeout:   30 * time.Second,
		WriteTimeout:  30 * time.Second,
		IdleTimeout:   120 * time.Second,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "X-Api-Key", "X-Request-ID"},
		AllowCredentials: false,
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		MaxAge:           3600,
	}))

	app.Use(middleware.Logger(), middleware.Recover())

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
