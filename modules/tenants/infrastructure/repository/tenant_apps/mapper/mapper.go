package mapper

import (
	"encoding/json"
	"nfxid/enums"
	"nfxid/modules/tenants/domain/tenant_apps"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"gorm.io/datatypes"
)

// TenantAppDomainToModel 将 Domain TenantApp 转换为 Model TenantApplication
func TenantAppDomainToModel(ta *tenant_apps.TenantApp) *models.TenantApplication {
	if ta == nil {
		return nil
	}

	var settings *datatypes.JSON
	if ta.Settings() != nil && len(ta.Settings()) > 0 {
		settingsBytes, _ := json.Marshal(ta.Settings())
		jsonData := datatypes.JSON(settingsBytes)
		settings = &jsonData
	}

	return &models.TenantApplication{
		ID:            ta.ID(),
		TenantID:      ta.TenantID(),
		ApplicationID: ta.AppID(),
		Status:        tenantAppStatusDomainToEnum(ta.Status()),
		CreatedAt:     ta.CreatedAt(),
		CreatedBy:     ta.CreatedBy(),
		UpdatedAt:     ta.UpdatedAt(),
		Settings:      settings,
	}
}

// TenantAppModelToDomain 将 Model TenantApplication 转换为 Domain TenantApp
func TenantAppModelToDomain(m *models.TenantApplication) *tenant_apps.TenantApp {
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
		AppID:     m.ApplicationID,
		Status:    tenantAppStatusEnumToDomain(m.Status),
		CreatedAt: m.CreatedAt,
		CreatedBy: m.CreatedBy,
		UpdatedAt: m.UpdatedAt,
		Settings:  settings,
	}

	return tenant_apps.NewTenantAppFromState(state)
}

// TenantAppModelToUpdates 将 Model TenantApplication 转换为更新字段映射
func TenantAppModelToUpdates(m *models.TenantApplication) map[string]any {
	var settings any
	if m.Settings != nil {
		settings = m.Settings
	}

	return map[string]any{
		models.TenantApplicationCols.TenantID:      m.TenantID,
		models.TenantApplicationCols.ApplicationID: m.ApplicationID,
		models.TenantApplicationCols.Status:        m.Status,
		models.TenantApplicationCols.CreatedBy:     m.CreatedBy,
		models.TenantApplicationCols.UpdatedAt:     m.UpdatedAt,
		models.TenantApplicationCols.Settings:      settings,
	}
}

// 枚举转换辅助函数

// TenantAppStatusDomainToEnum 将 Domain TenantAppStatus 转换为 Enum TenantsTenantApplicationStatus
func TenantAppStatusDomainToEnum(tas tenant_apps.TenantAppStatus) enums.TenantsTenantApplicationStatus {
	return tenantAppStatusDomainToEnum(tas)
}

func tenantAppStatusDomainToEnum(tas tenant_apps.TenantAppStatus) enums.TenantsTenantApplicationStatus {
	switch tas {
	case tenant_apps.TenantAppStatusActive:
		return enums.TenantsTenantApplicationStatusActive
	case tenant_apps.TenantAppStatusDisabled:
		return enums.TenantsTenantApplicationStatusDisabled
	case tenant_apps.TenantAppStatusSuspended:
		return enums.TenantsTenantApplicationStatusSuspended
	default:
		return enums.TenantsTenantApplicationStatusActive
	}
}

func tenantAppStatusEnumToDomain(tas enums.TenantsTenantApplicationStatus) tenant_apps.TenantAppStatus {
	switch tas {
	case enums.TenantsTenantApplicationStatusActive:
		return tenant_apps.TenantAppStatusActive
	case enums.TenantsTenantApplicationStatusDisabled:
		return tenant_apps.TenantAppStatusDisabled
	case enums.TenantsTenantApplicationStatusSuspended:
		return tenant_apps.TenantAppStatusSuspended
	default:
		return tenant_apps.TenantAppStatusActive
	}
}
