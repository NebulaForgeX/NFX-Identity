package mapper

import (
	"encoding/json"
	"nfxid/enums"
	"nfxid/modules/tenants/domain/tenant_apps"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"gorm.io/datatypes"
)

// TenantAppDomainToModel 将 Domain TenantApp 转换为 Model TenantApp
func TenantAppDomainToModel(ta *tenant_apps.TenantApp) *models.TenantApp {
	if ta == nil {
		return nil
	}

	var settings *datatypes.JSON
	if ta.Settings() != nil && len(ta.Settings()) > 0 {
		settingsBytes, _ := json.Marshal(ta.Settings())
		jsonData := datatypes.JSON(settingsBytes)
		settings = &jsonData
	}

	return &models.TenantApp{
		ID:        ta.ID(),
		TenantID:  ta.TenantID(),
		AppID:     ta.AppID(),
		Status:    tenantAppStatusDomainToEnum(ta.Status()),
		CreatedAt: ta.CreatedAt(),
		CreatedBy: ta.CreatedBy(),
		UpdatedAt: ta.UpdatedAt(),
		Settings:  settings,
	}
}

// TenantAppModelToDomain 将 Model TenantApp 转换为 Domain TenantApp
func TenantAppModelToDomain(m *models.TenantApp) *tenant_apps.TenantApp {
	if m == nil {
		return nil
	}

	var settings map[string]interface{}
	if m.Settings != nil {
		json.Unmarshal(*m.Settings, &settings)
	}

	state := tenant_apps.TenantAppState{
		ID:        m.ID,
		TenantID:  m.TenantID,
		AppID:     m.AppID,
		Status:    tenantAppStatusEnumToDomain(m.Status),
		CreatedAt: m.CreatedAt,
		CreatedBy: m.CreatedBy,
		UpdatedAt: m.UpdatedAt,
		Settings:  settings,
	}

	return tenant_apps.NewTenantAppFromState(state)
}

// TenantAppModelToUpdates 将 Model TenantApp 转换为更新字段映射
func TenantAppModelToUpdates(m *models.TenantApp) map[string]any {
	var settings any
	if m.Settings != nil {
		settings = m.Settings
	}

	return map[string]any{
		models.TenantAppCols.TenantID:  m.TenantID,
		models.TenantAppCols.AppID:     m.AppID,
		models.TenantAppCols.Status:    m.Status,
		models.TenantAppCols.CreatedBy: m.CreatedBy,
		models.TenantAppCols.UpdatedAt: m.UpdatedAt,
		models.TenantAppCols.Settings:  settings,
	}
}

// 枚举转换辅助函数

// TenantAppStatusDomainToEnum 将 Domain TenantAppStatus 转换为 Enum TenantAppStatus
func TenantAppStatusDomainToEnum(tas tenant_apps.TenantAppStatus) enums.TenantsTenantAppStatus {
	return tenantAppStatusDomainToEnum(tas)
}

func tenantAppStatusDomainToEnum(tas tenant_apps.TenantAppStatus) enums.TenantsTenantAppStatus {
	switch tas {
	case tenant_apps.TenantAppStatusActive:
		return enums.TenantsTenantAppStatusActive
	case tenant_apps.TenantAppStatusDisabled:
		return enums.TenantsTenantAppStatusDisabled
	case tenant_apps.TenantAppStatusSuspended:
		return enums.TenantsTenantAppStatusSuspended
	default:
		return enums.TenantsTenantAppStatusActive
	}
}

func tenantAppStatusEnumToDomain(tas enums.TenantsTenantAppStatus) tenant_apps.TenantAppStatus {
	switch tas {
	case enums.TenantsTenantAppStatusActive:
		return tenant_apps.TenantAppStatusActive
	case enums.TenantsTenantAppStatusDisabled:
		return tenant_apps.TenantAppStatusDisabled
	case enums.TenantsTenantAppStatusSuspended:
		return tenant_apps.TenantAppStatusSuspended
	default:
		return tenant_apps.TenantAppStatusActive
	}
}
