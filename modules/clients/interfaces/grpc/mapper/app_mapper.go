package mapper

import (
	appAppResult "nfxid/modules/clients/application/apps/results"
	appDomain "nfxid/modules/clients/domain/apps"
	apppb "nfxid/protos/gen/clients/app"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AppROToProto 将 AppRO 转换为 proto App 消息
func AppROToProto(v *appAppResult.AppRO) *apppb.App {
	if v == nil {
		return nil
	}

	app := &apppb.App{
		Id:        v.ID.String(),
		AppId:     v.AppID,
		TenantId:  v.TenantID.String(),
		Name:      v.Name,
		Type:      appTypeToProto(v.Type),
		Status:    appStatusToProto(v.Status),
		Environment: environmentToProto(v.Environment),
		CreatedAt: timestamppb.New(v.CreatedAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.Description != nil {
		app.Description = v.Description
	}
	if v.CreatedBy != nil {
		createdByStr := v.CreatedBy.String()
		app.CreatedBy = &createdByStr
	}
	if v.UpdatedBy != nil {
		updatedByStr := v.UpdatedBy.String()
		app.UpdatedBy = &updatedByStr
	}
	if v.Metadata != nil {
		if metadata, err := structpb.NewStruct(v.Metadata); err == nil {
			app.Metadata = metadata
		}
	}
	if v.DeletedAt != nil {
		app.DeletedAt = timestamppb.New(*v.DeletedAt)
	}

	return app
}

// AppListROToProto 批量转换 AppRO 到 proto App
func AppListROToProto(results []appAppResult.AppRO) []*apppb.App {
	apps := make([]*apppb.App, len(results))
	for i, v := range results {
		apps[i] = AppROToProto(&v)
	}
	return apps
}

func appTypeToProto(t appDomain.AppType) apppb.ClientsAppType {
	switch t {
	case appDomain.AppTypeServer:
		return apppb.ClientsAppType_CLIENTS_APP_TYPE_SERVER
	case appDomain.AppTypeService:
		return apppb.ClientsAppType_CLIENTS_APP_TYPE_SERVICE
	case appDomain.AppTypeInternal:
		return apppb.ClientsAppType_CLIENTS_APP_TYPE_INTERNAL
	case appDomain.AppTypePartner:
		return apppb.ClientsAppType_CLIENTS_APP_TYPE_PARTNER
	case appDomain.AppTypeThirdParty:
		return apppb.ClientsAppType_CLIENTS_APP_TYPE_THIRD_PARTY
	default:
		return apppb.ClientsAppType_CLIENTS_APP_TYPE_UNSPECIFIED
	}
}

func appStatusToProto(s appDomain.AppStatus) apppb.ClientsAppStatus {
	switch s {
	case appDomain.AppStatusActive:
		return apppb.ClientsAppStatus_CLIENTS_APP_STATUS_ACTIVE
	case appDomain.AppStatusDisabled:
		return apppb.ClientsAppStatus_CLIENTS_APP_STATUS_DISABLED
	case appDomain.AppStatusSuspended:
		return apppb.ClientsAppStatus_CLIENTS_APP_STATUS_SUSPENDED
	case appDomain.AppStatusPending:
		return apppb.ClientsAppStatus_CLIENTS_APP_STATUS_PENDING
	default:
		return apppb.ClientsAppStatus_CLIENTS_APP_STATUS_UNSPECIFIED
	}
}

func environmentToProto(e appDomain.Environment) apppb.ClientsEnvironment {
	switch e {
	case appDomain.EnvironmentProduction:
		return apppb.ClientsEnvironment_CLIENTS_ENVIRONMENT_PRODUCTION
	case appDomain.EnvironmentStaging:
		return apppb.ClientsEnvironment_CLIENTS_ENVIRONMENT_STAGING
	case appDomain.EnvironmentDevelopment:
		return apppb.ClientsEnvironment_CLIENTS_ENVIRONMENT_DEVELOPMENT
	case appDomain.EnvironmentTest:
		return apppb.ClientsEnvironment_CLIENTS_ENVIRONMENT_TEST
	default:
		return apppb.ClientsEnvironment_CLIENTS_ENVIRONMENT_UNSPECIFIED
	}
}
