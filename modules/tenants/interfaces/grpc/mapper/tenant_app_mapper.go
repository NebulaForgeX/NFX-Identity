package mapper

import (
	tenantAppAppResult "nfxid/modules/tenants/application/tenant_apps/results"
	tenantAppDomain "nfxid/modules/tenants/domain/tenant_apps"
	tenantapppb "nfxid/protos/gen/tenants/tenant_app"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TenantAppROToProto 将 TenantAppRO 转换为 proto TenantApp 消息
func TenantAppROToProto(v *tenantAppAppResult.TenantAppRO) *tenantapppb.TenantApp {
	if v == nil {
		return nil
	}

	tenantApp := &tenantapppb.TenantApp{
		Id:        v.ID.String(),
		TenantId:  v.TenantID.String(),
		AppId:     v.AppID.String(),
		Status:    tenantAppStatusToProto(v.Status),
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.CreatedBy != nil {
		createdBy := v.CreatedBy.String()
		tenantApp.CreatedBy = &createdBy
	}

	if v.Settings != nil && len(v.Settings) > 0 {
		if settingsStruct, err := structpb.NewStruct(v.Settings); err == nil {
			tenantApp.Settings = settingsStruct
		}
	}

	return tenantApp
}

// TenantAppListROToProto 批量转换 TenantAppRO 到 proto TenantApp
func TenantAppListROToProto(results []tenantAppAppResult.TenantAppRO) []*tenantapppb.TenantApp {
	tenantApps := make([]*tenantapppb.TenantApp, len(results))
	for i, v := range results {
		tenantApps[i] = TenantAppROToProto(&v)
	}
	return tenantApps
}

// tenantAppStatusToProto 将 domain TenantAppStatus 转换为 proto TenantsTenantAppStatus
func tenantAppStatusToProto(status tenantAppDomain.TenantAppStatus) tenantapppb.TenantsTenantAppStatus {
	switch status {
	case tenantAppDomain.TenantAppStatusActive:
		return tenantapppb.TenantsTenantAppStatus_TENANTS_TENANT_APP_STATUS_ACTIVE
	case tenantAppDomain.TenantAppStatusDisabled:
		return tenantapppb.TenantsTenantAppStatus_TENANTS_TENANT_APP_STATUS_DISABLED
	case tenantAppDomain.TenantAppStatusSuspended:
		return tenantapppb.TenantsTenantAppStatus_TENANTS_TENANT_APP_STATUS_SUSPENDED
	default:
		return tenantapppb.TenantsTenantAppStatus_TENANTS_TENANT_APP_STATUS_UNSPECIFIED
	}
}
