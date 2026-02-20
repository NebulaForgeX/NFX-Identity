package mapper

import (
	"encoding/json"
	"nfxid/enums"
	"nfxid/modules/clients/domain/apps"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"

	"gorm.io/datatypes"
)

// AppDomainToModel 将 Domain App 转换为 Model Application
func AppDomainToModel(a *apps.App) *models.Application {
	if a == nil {
		return nil
	}

	var metadata *datatypes.JSON
	if a.Metadata() != nil {
		metadataBytes, _ := json.Marshal(a.Metadata())
		metadata = (*datatypes.JSON)(&metadataBytes)
	}

	return &models.Application{
		ID:            a.ID(),
		ApplicationID: a.AppID(),
		TenantID:      a.TenantID(),
		Name:          a.Name(),
		Description:   a.Description(),
		Type:          appTypeDomainToEnum(a.Type()),
		Status:        appStatusDomainToEnum(a.Status()),
		Environment:   environmentDomainToEnum(a.Environment()),
		CreatedAt:     a.CreatedAt(),
		UpdatedAt:     a.UpdatedAt(),
		CreatedBy:     a.CreatedBy(),
		UpdatedBy:     a.UpdatedBy(),
		Metadata:      metadata,
		DeletedAt:     timex.TimeToGormDeletedAt(a.DeletedAt()),
	}
}

// AppModelToDomain 将 Model Application 转换为 Domain App
func AppModelToDomain(m *models.Application) *apps.App {
	if m == nil {
		return nil
	}

	var metadata map[string]interface{}
	if m.Metadata != nil {
		json.Unmarshal(*m.Metadata, &metadata)
	}

	state := apps.AppState{
		ID:          m.ID,
		AppID:       m.ApplicationID,
		TenantID:    m.TenantID,
		Name:        m.Name,
		Description: m.Description,
		Type:        appTypeEnumToDomain(m.Type),
		Status:      appStatusEnumToDomain(m.Status),
		Environment: environmentEnumToDomain(m.Environment),
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		CreatedBy:   m.CreatedBy,
		UpdatedBy:   m.UpdatedBy,
		Metadata:    metadata,
		DeletedAt:   timex.GormDeletedAtToTime(m.DeletedAt),
	}

	return apps.NewAppFromState(state)
}

// AppModelToUpdates 将 Model Application 转换为更新字段映射
func AppModelToUpdates(m *models.Application) map[string]any {
	updates := map[string]any{
		models.ApplicationCols.Name:        m.Name,
		models.ApplicationCols.Description: m.Description,
		models.ApplicationCols.Type:        m.Type,
		models.ApplicationCols.Status:      m.Status,
		models.ApplicationCols.Environment: m.Environment,
		models.ApplicationCols.UpdatedAt:   m.UpdatedAt,
		models.ApplicationCols.UpdatedBy:   m.UpdatedBy,
		models.ApplicationCols.Metadata:    m.Metadata,
		models.ApplicationCols.DeletedAt:   m.DeletedAt,
	}
	return updates
}

// 枚举转换辅助函数

func appTypeDomainToEnum(at apps.AppType) enums.ClientsAppType {
	switch at {
	case apps.AppTypeServer:
		return enums.ClientsAppTypeServer
	case apps.AppTypeService:
		return enums.ClientsAppTypeService
	case apps.AppTypeInternal:
		return enums.ClientsAppTypeInternal
	case apps.AppTypePartner:
		return enums.ClientsAppTypePartner
	case apps.AppTypeThirdParty:
		return enums.ClientsAppTypeThirdParty
	default:
		return enums.ClientsAppTypeServer
	}
}

func appTypeEnumToDomain(at enums.ClientsAppType) apps.AppType {
	switch at {
	case enums.ClientsAppTypeServer:
		return apps.AppTypeServer
	case enums.ClientsAppTypeService:
		return apps.AppTypeService
	case enums.ClientsAppTypeInternal:
		return apps.AppTypeInternal
	case enums.ClientsAppTypePartner:
		return apps.AppTypePartner
	case enums.ClientsAppTypeThirdParty:
		return apps.AppTypeThirdParty
	default:
		return apps.AppTypeServer
	}
}

// AppStatusDomainToEnum 将 Domain AppStatus 转换为 Enum AppStatus（导出供其他包使用）
func AppStatusDomainToEnum(as apps.AppStatus) enums.ClientsAppStatus {
	return appStatusDomainToEnum(as)
}

func appStatusDomainToEnum(as apps.AppStatus) enums.ClientsAppStatus {
	switch as {
	case apps.AppStatusActive:
		return enums.ClientsAppStatusActive
	case apps.AppStatusDisabled:
		return enums.ClientsAppStatusDisabled
	case apps.AppStatusSuspended:
		return enums.ClientsAppStatusSuspended
	case apps.AppStatusPending:
		return enums.ClientsAppStatusPending
	default:
		return enums.ClientsAppStatusPending
	}
}

func appStatusEnumToDomain(as enums.ClientsAppStatus) apps.AppStatus {
	switch as {
	case enums.ClientsAppStatusActive:
		return apps.AppStatusActive
	case enums.ClientsAppStatusDisabled:
		return apps.AppStatusDisabled
	case enums.ClientsAppStatusSuspended:
		return apps.AppStatusSuspended
	case enums.ClientsAppStatusPending:
		return apps.AppStatusPending
	default:
		return apps.AppStatusPending
	}
}

// EnvironmentDomainToEnum 将 Domain Environment 转换为 Enum Environment（导出供其他包使用）
func EnvironmentDomainToEnum(e apps.Environment) enums.ClientsEnvironment {
	return environmentDomainToEnum(e)
}

func environmentDomainToEnum(e apps.Environment) enums.ClientsEnvironment {
	switch e {
	case apps.EnvironmentProduction:
		return enums.ClientsEnvironmentProduction
	case apps.EnvironmentStaging:
		return enums.ClientsEnvironmentStaging
	case apps.EnvironmentDevelopment:
		return enums.ClientsEnvironmentDevelopment
	case apps.EnvironmentTest:
		return enums.ClientsEnvironmentTest
	default:
		return enums.ClientsEnvironmentDevelopment
	}
}

func environmentEnumToDomain(e enums.ClientsEnvironment) apps.Environment {
	switch e {
	case enums.ClientsEnvironmentProduction:
		return apps.EnvironmentProduction
	case enums.ClientsEnvironmentStaging:
		return apps.EnvironmentStaging
	case enums.ClientsEnvironmentDevelopment:
		return apps.EnvironmentDevelopment
	case enums.ClientsEnvironmentTest:
		return apps.EnvironmentTest
	default:
		return apps.EnvironmentDevelopment
	}
}
