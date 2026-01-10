package http

import (
	"encoding/json"

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
	"nfxid/pkgs/recover"
	"nfxid/pkgs/security/token"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false,
		ExposeHeaders:    "Content-Length",
	}))

	app.Use(recover.RecoverMiddleware(), logger.New())

	reg := &Registry{
		Tenant:            handler.NewTenantHandler(d.TenantAppSvc()),
		Group:             handler.NewGroupHandler(d.GroupAppSvc()),
		Member:            handler.NewMemberHandler(d.MemberAppSvc()),
		Invitation:        handler.NewInvitationHandler(d.InvitationAppSvc()),
		TenantApp:         handler.NewTenantAppHandler(d.TenantAppAppSvc()),
		TenantSetting:     handler.NewTenantSettingHandler(d.TenantSettingAppSvc()),
		DomainVerification: handler.NewDomainVerificationHandler(d.DomainVerificationAppSvc()),
		MemberRole:        handler.NewMemberRoleHandler(d.MemberRoleAppSvc()),
		MemberGroup:       handler.NewMemberGroupHandler(d.MemberGroupAppSvc()),
		MemberAppRole:     handler.NewMemberAppRoleHandler(d.MemberAppRoleAppSvc()),
	}

	router := NewRouter(app, d.UserTokenVerifier(), reg)
	router.RegisterRoutes()

	return app
}
