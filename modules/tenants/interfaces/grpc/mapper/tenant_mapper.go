package mapper

import (
	tenantAppResult "nfxid/modules/tenants/application/tenants/results"
	tenantDomain "nfxid/modules/tenants/domain/tenants"
	tenantpb "nfxid/protos/gen/tenants/tenant"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TenantROToProto 将 TenantRO 转换为 proto Tenant 消息
func TenantROToProto(v *tenantAppResult.TenantRO) *tenantpb.Tenant {
	if v == nil {
		return nil
	}

	tenant := &tenantpb.Tenant{
		Id:        v.ID.String(),
		TenantId:  v.TenantID,
		Name:      v.Name,
		Status:    tenantStatusToProto(v.Status),
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.DisplayName != nil {
		tenant.DisplayName = v.DisplayName
	}

	if v.PrimaryDomain != nil {
		tenant.PrimaryDomain = v.PrimaryDomain
	}

	if v.DeletedAt != nil {
		tenant.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	if v.Metadata != nil && len(v.Metadata) > 0 {
		if metadataStruct, err := structpb.NewStruct(v.Metadata); err == nil {
			tenant.Metadata = metadataStruct
		}
	}

	return tenant
}

// TenantListROToProto 批量转换 TenantRO 到 proto Tenant
func TenantListROToProto(results []tenantAppResult.TenantRO) []*tenantpb.Tenant {
	tenants := make([]*tenantpb.Tenant, len(results))
	for i, v := range results {
		tenants[i] = TenantROToProto(&v)
	}
	return tenants
}

// tenantStatusToProto 将 domain TenantStatus 转换为 proto TenantsTenantStatus
func tenantStatusToProto(status tenantDomain.TenantStatus) tenantpb.TenantsTenantStatus {
	switch status {
	case tenantDomain.TenantStatusActive:
		return tenantpb.TenantsTenantStatus_TENANTS_TENANT_STATUS_ACTIVE
	case tenantDomain.TenantStatusSuspended:
		return tenantpb.TenantsTenantStatus_TENANTS_TENANT_STATUS_SUSPENDED
	case tenantDomain.TenantStatusClosed:
		return tenantpb.TenantsTenantStatus_TENANTS_TENANT_STATUS_CLOSED
	case tenantDomain.TenantStatusPending:
		return tenantpb.TenantsTenantStatus_TENANTS_TENANT_STATUS_PENDING
	default:
		return tenantpb.TenantsTenantStatus_TENANTS_TENANT_STATUS_UNSPECIFIED
	}
}
