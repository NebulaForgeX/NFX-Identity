package grpc

import (
	domainVerificationApp "nfxid/modules/tenants/application/domain_verifications"
	groupApp "nfxid/modules/tenants/application/groups"
	invitationApp "nfxid/modules/tenants/application/invitations"
	memberApp "nfxid/modules/tenants/application/member_app_roles"
	memberGroupApp "nfxid/modules/tenants/application/member_groups"
	memberRoleApp "nfxid/modules/tenants/application/member_roles"
	membersApp "nfxid/modules/tenants/application/members"
	tenantApp "nfxid/modules/tenants/application/tenants"
	tenantAppApp "nfxid/modules/tenants/application/tenant_apps"
	tenantSettingApp "nfxid/modules/tenants/application/tenant_settings"
	resourceApp "nfxid/modules/tenants/application/resource"
	grpcHandler "nfxid/modules/tenants/interfaces/grpc/handler"
	"nfxid/pkgs/postgresqlx"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	healthpb "nfxid/protos/gen/common/health"
	schemapb "nfxid/protos/gen/common/schema"
	domainverificationpb "nfxid/protos/gen/tenants/domain_verification"
	grouppb "nfxid/protos/gen/tenants/group"
	invitationpb "nfxid/protos/gen/tenants/invitation"
	membergrouppb "nfxid/protos/gen/tenants/member_group"
	memberpb "nfxid/protos/gen/tenants/member"
	tenantapplicationpb "nfxid/protos/gen/tenants/tenant_application"
	tenantsettingpb "nfxid/protos/gen/tenants/tenant_setting"
	tenantpb "nfxid/protos/gen/tenants/tenant"

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
		grpc.UnaryInterceptor(servertoken.UnaryAuthInterceptor(d.ServerTokenVerifier())),
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
