package tenants

import (
	domainverificationpb "nfxidentity/protos/gen/tenants/domain_verification"
	grouppb "nfxidentity/protos/gen/tenants/group"
	invitationpb "nfxidentity/protos/gen/tenants/invitation"
	memberpb "nfxidentity/protos/gen/tenants/member"
	membergrouppb "nfxidentity/protos/gen/tenants/member_group"
	tenantpb "nfxidentity/protos/gen/tenants/tenant"
	tenantapplicationpb "nfxidentity/protos/gen/tenants/tenant_application"
	tenantsettingpb "nfxidentity/protos/gen/tenants/tenant_setting"
)

// Client Tenants 服务 gRPC 客户端聚合
type Client struct {
	Tenant             *TenantClient
	TenantApplication  *TenantApplicationClient
	TenantSetting      *TenantSettingClient
	Member             *MemberClient
	Group              *GroupClient
	Invitation         *InvitationClient
	MemberGroup        *MemberGroupClient
	DomainVerification *DomainVerificationClient
}

// NewClient 创建 Tenants 客户端
func NewClient(
	tenantClient tenantpb.TenantServiceClient,
	tenantApplicationClient tenantapplicationpb.TenantApplicationServiceClient,
	tenantSettingClient tenantsettingpb.TenantSettingServiceClient,
	memberClient memberpb.MemberServiceClient,
	groupClient grouppb.GroupServiceClient,
	invitationClient invitationpb.InvitationServiceClient,
	memberGroupClient membergrouppb.MemberGroupServiceClient,
	domainVerificationClient domainverificationpb.DomainVerificationServiceClient,
) *Client {
	return &Client{
		Tenant:             NewTenantClient(tenantClient),
		TenantApplication:  NewTenantApplicationClient(tenantApplicationClient),
		TenantSetting:      NewTenantSettingClient(tenantSettingClient),
		Member:             NewMemberClient(memberClient),
		Group:              NewGroupClient(groupClient),
		Invitation:         NewInvitationClient(invitationClient),
		MemberGroup:        NewMemberGroupClient(memberGroupClient),
		DomainVerification: NewDomainVerificationClient(domainVerificationClient),
	}
}
