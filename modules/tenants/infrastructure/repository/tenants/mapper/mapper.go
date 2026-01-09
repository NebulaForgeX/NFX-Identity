package mapper

import (
	"encoding/json"
	"nfxid/enums"
	"nfxid/modules/tenants/domain/tenants"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/pkgs/utils/timex"

	"gorm.io/datatypes"
)

// TenantDomainToModel 将 Domain Tenant 转换为 Model Tenant
func TenantDomainToModel(t *tenants.Tenant) *models.Tenant {
	if t == nil {
		return nil
	}

	var metadata *datatypes.JSON
	if t.Metadata() != nil && len(t.Metadata()) > 0 {
		metadataBytes, _ := json.Marshal(t.Metadata())
		jsonData := datatypes.JSON(metadataBytes)
		metadata = &jsonData
	}

	return &models.Tenant{
		ID:            t.ID(),
		TenantID:      t.TenantID(),
		Name:          t.Name(),
		DisplayName:   t.DisplayName(),
		Status:        tenantStatusDomainToEnum(t.Status()),
		PrimaryDomain: t.PrimaryDomain(),
		CreatedAt:     t.CreatedAt(),
		UpdatedAt:     t.UpdatedAt(),
		DeletedAt:     timex.TimeToGormDeletedAt(t.DeletedAt()),
		Metadata:      metadata,
	}
}

// TenantModelToDomain 将 Model Tenant 转换为 Domain Tenant
func TenantModelToDomain(m *models.Tenant) *tenants.Tenant {
	if m == nil {
		return nil
	}

	var metadata map[string]interface{}
	if m.Metadata != nil {
		json.Unmarshal(*m.Metadata, &metadata)
	}

	state := tenants.TenantState{
		ID:            m.ID,
		TenantID:      m.TenantID,
		Name:          m.Name,
		DisplayName:   m.DisplayName,
		Status:        tenantStatusEnumToDomain(m.Status),
		PrimaryDomain: m.PrimaryDomain,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
		DeletedAt:     timex.GormDeletedAtToTime(m.DeletedAt),
		Metadata:      metadata,
	}

	return tenants.NewTenantFromState(state)
}

// TenantModelToUpdates 将 Model Tenant 转换为更新字段映射
func TenantModelToUpdates(m *models.Tenant) map[string]any {
	var metadata any
	if m.Metadata != nil {
		metadata = m.Metadata
	}

	return map[string]any{
		models.TenantCols.TenantID:      m.TenantID,
		models.TenantCols.Name:          m.Name,
		models.TenantCols.DisplayName:   m.DisplayName,
		models.TenantCols.Status:        m.Status,
		models.TenantCols.PrimaryDomain: m.PrimaryDomain,
		models.TenantCols.UpdatedAt:     m.UpdatedAt,
		models.TenantCols.DeletedAt:    m.DeletedAt,
		models.TenantCols.Metadata:      metadata,
	}
}

// 枚举转换辅助函数

// TenantStatusDomainToEnum 将 Domain TenantStatus 转换为 Enum TenantStatus
func TenantStatusDomainToEnum(ts tenants.TenantStatus) enums.TenantsTenantStatus {
	return tenantStatusDomainToEnum(ts)
}

func tenantStatusDomainToEnum(ts tenants.TenantStatus) enums.TenantsTenantStatus {
	switch ts {
	case tenants.TenantStatusActive:
		return enums.TenantsTenantStatusActive
	case tenants.TenantStatusSuspended:
		return enums.TenantsTenantStatusSuspended
	case tenants.TenantStatusClosed:
		return enums.TenantsTenantStatusClosed
	case tenants.TenantStatusPending:
		return enums.TenantsTenantStatusPending
	default:
		return enums.TenantsTenantStatusPending
	}
}

func tenantStatusEnumToDomain(ts enums.TenantsTenantStatus) tenants.TenantStatus {
	switch ts {
	case enums.TenantsTenantStatusActive:
		return tenants.TenantStatusActive
	case enums.TenantsTenantStatusSuspended:
		return tenants.TenantStatusSuspended
	case enums.TenantsTenantStatusClosed:
		return tenants.TenantStatusClosed
	case enums.TenantsTenantStatusPending:
		return tenants.TenantStatusPending
	default:
		return tenants.TenantStatusPending
	}
}
