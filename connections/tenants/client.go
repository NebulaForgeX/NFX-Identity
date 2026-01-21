package tenants

import (
	domainverificationpb "nfxid/protos/gen/tenants/domain_verification"
	grouppb "nfxid/protos/gen/tenants/group"
	invitationpb "nfxid/protos/gen/tenants/invitation"
	memberapprolepb "nfxid/protos/gen/tenants/member_app_role"
	membergrouppb "nfxid/protos/gen/tenants/member_group"
	memberpb "nfxid/protos/gen/tenants/member"
	memberrolepb "nfxid/protos/gen/tenants/member_role"
	tenantapppb "nfxid/protos/gen/tenants/tenant_app"
	tenantpb "nfxid/protos/gen/tenants/tenant"
	tenantsettingpb "nfxid/protos/gen/tenants/tenant_setting"
)

// Client Tenants 服务客户端
type Client struct {
	Tenant            *TenantClient
	TenantApp         *TenantAppClient
	TenantSetting     *TenantSettingClient
	Member            *MemberClient
	Group             *GroupClient
	Invitation        *InvitationClient
	MemberGroup       *MemberGroupClient
	MemberRole        *MemberRoleClient
	MemberAppRole     *MemberAppRoleClient
	DomainVerification *DomainVerificationClient
}

// NewClient 创建 Tenants 客户端
func NewClient(
	tenantClient tenantpb.TenantServiceClient,
	tenantAppClient tenantapppb.TenantAppServiceClient,
	tenantSettingClient tenantsettingpb.TenantSettingServiceClient,
	memberClient memberpb.MemberServiceClient,
	groupClient grouppb.GroupServiceClient,
	invitationClient invitationpb.InvitationServiceClient,
	memberGroupClient membergrouppb.MemberGroupServiceClient,
	memberRoleClient memberrolepb.MemberRoleServiceClient,
	memberAppRoleClient memberapprolepb.MemberAppRoleServiceClient,
	domainVerificationClient domainverificationpb.DomainVerificationServiceClient,
) *Client {
	return &Client{
		Tenant:            NewTenantClient(tenantClient),
		TenantApp:         NewTenantAppClient(tenantAppClient),
		TenantSetting:     NewTenantSettingClient(tenantSettingClient),
		Member:            NewMemberClient(memberClient),
		Group:             NewGroupClient(groupClient),
		Invitation:        NewInvitationClient(invitationClient),
		MemberGroup:       NewMemberGroupClient(memberGroupClient),
		MemberRole:        NewMemberRoleClient(memberRoleClient),
		MemberAppRole:     NewMemberAppRoleClient(memberAppRoleClient),
		DomainVerification: NewDomainVerificationClient(domainVerificationClient),
	}
}
