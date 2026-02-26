package grpc

import (
	domainVerificationApp "nfxidentity/modules/tenants/application/domain_verifications"
	groupApp "nfxidentity/modules/tenants/application/groups"
	invitationApp "nfxidentity/modules/tenants/application/invitations"
	memberApp "nfxidentity/modules/tenants/application/member_app_roles"
	memberGroupApp "nfxidentity/modules/tenants/application/member_groups"
	memberRoleApp "nfxidentity/modules/tenants/application/member_roles"
	membersApp "nfxidentity/modules/tenants/application/members"
	resourceApp "nfxidentity/modules/tenants/application/resource"
	tenantAppApp "nfxidentity/modules/tenants/application/tenant_apps"
	tenantSettingApp "nfxidentity/modules/tenants/application/tenant_settings"
	tenantApp "nfxidentity/modules/tenants/application/tenants"
	grpcHandler "nfxidentity/modules/tenants/interfaces/grpc/handler"
	"nfxidentity/pkgs/grpcx/interceptor"
	"nfxidentity/pkgs/postgresqlx"
	"nfxidentity/pkgs/security/token"
	"nfxidentity/pkgs/security/token/servertoken"
	healthpb "nfxidentity/protos/gen/common/health"
	schemapb "nfxidentity/protos/gen/common/schema"
	domainverificationpb "nfxidentity/protos/gen/tenants/domain_verification"
	grouppb "nfxidentity/protos/gen/tenants/group"
	invitationpb "nfxidentity/protos/gen/tenants/invitation"
	memberpb "nfxidentity/protos/gen/tenants/member"
	membergrouppb "nfxidentity/protos/gen/tenants/member_group"
	tenantpb "nfxidentity/protos/gen/tenants/tenant"
	tenantapplicationpb "nfxidentity/protos/gen/tenants/tenant_application"
	tenantsettingpb "nfxidentity/protos/gen/tenants/tenant_setting"

	"google.golang.org/grpc"
)

type Deps interface {
	TenantAppSvc() *tenantApp.Service
	GroupAppSvc() *groupApp.Service
	MemberAppSvc() *membersApp.Service
	InvitationAppSvc() *invitationApp.Service
	TenantAppAppSvc() *tenantAppApp.Service
	TenantSettingAppSvc() *tenantSettingApp.Service
	DomainVerificationAppSvc() *domainVerificationApp.Service
	MemberRoleAppSvc() *memberRoleApp.Service
	MemberGroupAppSvc() *memberGroupApp.Service
	MemberAppRoleAppSvc() *memberApp.Service
	ResourceSvc() *resourceApp.Service
	ServerTokenVerifier() token.Verifier
	Postgres() *postgresqlx.Connection
}

func NewServer(d Deps) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryErrorHandler(),
			servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier()),
		),
	}

	s := grpc.NewServer(opts...)

	// Register gRPC services
	tenantpb.RegisterTenantServiceServer(s, grpcHandler.NewTenantHandler(d.TenantAppSvc()))
	grouppb.RegisterGroupServiceServer(s, grpcHandler.NewGroupHandler(d.GroupAppSvc()))
	memberpb.RegisterMemberServiceServer(s, grpcHandler.NewMemberHandler(d.MemberAppSvc()))
	invitationpb.RegisterInvitationServiceServer(s, grpcHandler.NewInvitationHandler(d.InvitationAppSvc()))
	tenantapplicationpb.RegisterTenantApplicationServiceServer(s, grpcHandler.NewTenantApplicationHandler(d.TenantAppAppSvc()))
	tenantsettingpb.RegisterTenantSettingServiceServer(s, grpcHandler.NewTenantSettingHandler(d.TenantSettingAppSvc()))
	domainverificationpb.RegisterDomainVerificationServiceServer(s, grpcHandler.NewDomainVerificationHandler(d.DomainVerificationAppSvc()))
	membergrouppb.RegisterMemberGroupServiceServer(s, grpcHandler.NewMemberGroupHandler(d.MemberGroupAppSvc()))

	// Register health check service
	healthpb.RegisterHealthServiceServer(s, grpcHandler.NewHealthHandler(d.ResourceSvc(), "tenants"))

	// Register schema service
	schemapb.RegisterSchemaServiceServer(s, grpcHandler.NewSchemaHandler(d.Postgres().DB(), "tenants"))

	return s
}
