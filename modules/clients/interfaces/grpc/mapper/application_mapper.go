package mapper

import (
	appAppResult "nfxid/modules/clients/application/apps/results"
	appDomain "nfxid/modules/clients/domain/apps"
	applicationpb "nfxid/protos/gen/clients/application"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ApplicationROToProto(v *appAppResult.AppRO) *applicationpb.Application {
	if v == nil {
		return nil
	}
	app := &applicationpb.Application{
		Id:          v.ID.String(),
		ApplicationId: v.AppID,
		TenantId:    v.TenantID.String(),
		Name:        v.Name,
		Type:        applicationTypeToProto(v.Type),
		Status:      applicationStatusToProto(v.Status),
		Environment: environmentToProto(v.Environment),
		CreatedAt:   timestamppb.New(v.CreatedAt),
		UpdatedAt:   timestamppb.New(v.UpdatedAt),
	}
	if v.Description != nil {
		app.Description = v.Description
	}
	if v.CreatedBy != nil {
		s := v.CreatedBy.String()
		app.CreatedBy = &s
	}
	if v.UpdatedBy != nil {
		s := v.UpdatedBy.String()
		app.UpdatedBy = &s
	}
	if v.Metadata != nil {
		if m, err := structpb.NewStruct(v.Metadata); err == nil {
			app.Metadata = m
		}
	}
	if v.DeletedAt != nil {
		app.DeletedAt = timestamppb.New(*v.DeletedAt)
	}
	return app
}

func ApplicationListROToProto(results []appAppResult.AppRO) []*applicationpb.Application {
	out := make([]*applicationpb.Application, len(results))
	for i := range results {
		out[i] = ApplicationROToProto(&results[i])
	}
	return out
}

func applicationTypeToProto(t appDomain.AppType) applicationpb.ClientsAppType {
	switch t {
	case appDomain.AppTypeServer:
		return applicationpb.ClientsAppType_CLIENTS_APP_TYPE_SERVER
	case appDomain.AppTypeService:
		return applicationpb.ClientsAppType_CLIENTS_APP_TYPE_SERVICE
	case appDomain.AppTypeInternal:
		return applicationpb.ClientsAppType_CLIENTS_APP_TYPE_INTERNAL
	case appDomain.AppTypePartner:
		return applicationpb.ClientsAppType_CLIENTS_APP_TYPE_PARTNER
	case appDomain.AppTypeThirdParty:
		return applicationpb.ClientsAppType_CLIENTS_APP_TYPE_THIRD_PARTY
	default:
		return applicationpb.ClientsAppType_CLIENTS_APP_TYPE_UNSPECIFIED
	}
}

func applicationStatusToProto(s appDomain.AppStatus) applicationpb.ClientsAppStatus {
	switch s {
	case appDomain.AppStatusActive:
		return applicationpb.ClientsAppStatus_CLIENTS_APP_STATUS_ACTIVE
	case appDomain.AppStatusDisabled:
		return applicationpb.ClientsAppStatus_CLIENTS_APP_STATUS_DISABLED
	case appDomain.AppStatusSuspended:
		return applicationpb.ClientsAppStatus_CLIENTS_APP_STATUS_SUSPENDED
	case appDomain.AppStatusPending:
		return applicationpb.ClientsAppStatus_CLIENTS_APP_STATUS_PENDING
	default:
		return applicationpb.ClientsAppStatus_CLIENTS_APP_STATUS_UNSPECIFIED
	}
}

func environmentToProto(e appDomain.Environment) applicationpb.ClientsEnvironment {
	switch e {
	case appDomain.EnvironmentProduction:
		return applicationpb.ClientsEnvironment_CLIENTS_ENVIRONMENT_PRODUCTION
	case appDomain.EnvironmentStaging:
		return applicationpb.ClientsEnvironment_CLIENTS_ENVIRONMENT_STAGING
	case appDomain.EnvironmentDevelopment:
		return applicationpb.ClientsEnvironment_CLIENTS_ENVIRONMENT_DEVELOPMENT
	case appDomain.EnvironmentTest:
		return applicationpb.ClientsEnvironment_CLIENTS_ENVIRONMENT_TEST
	default:
		return applicationpb.ClientsEnvironment_CLIENTS_ENVIRONMENT_UNSPECIFIED
	}
}
