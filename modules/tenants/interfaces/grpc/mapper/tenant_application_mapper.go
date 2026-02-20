package mapper

import (
	tenantAppAppResult "nfxid/modules/tenants/application/tenant_apps/results"
	tenantAppDomain "nfxid/modules/tenants/domain/tenant_apps"
	tenantapplicationpb "nfxid/protos/gen/tenants/tenant_application"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TenantApplicationROToProto(v *tenantAppAppResult.TenantAppRO) *tenantapplicationpb.TenantApplication {
	if v == nil {
		return nil
	}
	ta := &tenantapplicationpb.TenantApplication{
		Id:            v.ID.String(),
		TenantId:      v.TenantID.String(),
		ApplicationId: v.AppID.String(),
		Status:        tenantApplicationStatusToProto(v.Status),
		CreatedAt:     timestamppb.New(v.CreatedAt),
		UpdatedAt:     timestamppb.New(v.UpdatedAt),
	}
	if v.CreatedBy != nil {
		createdBy := v.CreatedBy.String()
		ta.CreatedBy = &createdBy
	}
	if v.Settings != nil && len(v.Settings) > 0 {
		if settingsStruct, err := structpb.NewStruct(v.Settings); err == nil {
			ta.Settings = settingsStruct
		}
	}
	return ta
}

func TenantApplicationListROToProto(results []tenantAppAppResult.TenantAppRO) []*tenantapplicationpb.TenantApplication {
	out := make([]*tenantapplicationpb.TenantApplication, len(results))
	for i := range results {
		out[i] = TenantApplicationROToProto(&results[i])
	}
	return out
}

func tenantApplicationStatusToProto(status tenantAppDomain.TenantAppStatus) tenantapplicationpb.TenantsTenantApplicationStatus {
	switch status {
	case tenantAppDomain.TenantAppStatusActive:
		return tenantapplicationpb.TenantsTenantApplicationStatus_TENANTS_TENANT_APPLICATION_STATUS_ACTIVE
	case tenantAppDomain.TenantAppStatusDisabled:
		return tenantapplicationpb.TenantsTenantApplicationStatus_TENANTS_TENANT_APPLICATION_STATUS_DISABLED
	case tenantAppDomain.TenantAppStatusSuspended:
		return tenantapplicationpb.TenantsTenantApplicationStatus_TENANTS_TENANT_APPLICATION_STATUS_SUSPENDED
	default:
		return tenantapplicationpb.TenantsTenantApplicationStatus_TENANTS_TENANT_APPLICATION_STATUS_UNSPECIFIED
	}
}
