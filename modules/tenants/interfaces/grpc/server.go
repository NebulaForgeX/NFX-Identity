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
	grpcHandler "nfxid/modules/tenants/interfaces/grpc/handler"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"
	domainverificationpb "nfxid/protos/gen/tenants/domain_verification"
	grouppb "nfxid/protos/gen/tenants/group"
	invitationpb "nfxid/protos/gen/tenants/invitation"
	memberapprolepb "nfxid/protos/gen/tenants/member_app_role"
	membergrouppb "nfxid/protos/gen/tenants/member_group"
	memberrolepb "nfxid/protos/gen/tenants/member_role"
	memberpb "nfxid/protos/gen/tenants/member"
	tenantapppb "nfxid/protos/gen/tenants/tenant_app"
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
	ServerTokenVerifier() token.Verifier
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
	tenantapppb.RegisterTenantAppServiceServer(s, grpcHandler.NewTenantAppHandler(d.TenantAppAppSvc()))
	tenantsettingpb.RegisterTenantSettingServiceServer(s, grpcHandler.NewTenantSettingHandler(d.TenantSettingAppSvc()))
	domainverificationpb.RegisterDomainVerificationServiceServer(s, grpcHandler.NewDomainVerificationHandler(d.DomainVerificationAppSvc()))
	memberrolepb.RegisterMemberRoleServiceServer(s, grpcHandler.NewMemberRoleHandler(d.MemberRoleAppSvc()))
	membergrouppb.RegisterMemberGroupServiceServer(s, grpcHandler.NewMemberGroupHandler(d.MemberGroupAppSvc()))
	memberapprolepb.RegisterMemberAppRoleServiceServer(s, grpcHandler.NewMemberAppRoleHandler(d.MemberAppRoleAppSvc()))

	return s
}
