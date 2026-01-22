package mapper

import (
	"encoding/json"
	"strings"
	"nfxid/modules/tenants/domain/tenant_settings"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"gorm.io/datatypes"
)

// TenantSettingDomainToModel 将 Domain TenantSetting 转换为 Model TenantSetting
func TenantSettingDomainToModel(ts *tenant_settings.TenantSetting) *models.TenantSetting {
	if ts == nil {
		return nil
	}

	var allowedEmailDomains *string
	if len(ts.AllowedEmailDomains()) > 0 {
		domainsStr := strings.Join(ts.AllowedEmailDomains(), ",")
		allowedEmailDomains = &domainsStr
	}

	var passwordPolicy *datatypes.JSON
	if ts.PasswordPolicy() != nil && len(ts.PasswordPolicy()) > 0 {
		policyBytes, _ := json.Marshal(ts.PasswordPolicy())
		jsonData := datatypes.JSON(policyBytes)
		passwordPolicy = &jsonData
	}

	var loginPolicy *datatypes.JSON
	if ts.LoginPolicy() != nil && len(ts.LoginPolicy()) > 0 {
		policyBytes, _ := json.Marshal(ts.LoginPolicy())
		jsonData := datatypes.JSON(policyBytes)
		loginPolicy = &jsonData
	}

	var mfaPolicy *datatypes.JSON
	if ts.MFAPolicy() != nil && len(ts.MFAPolicy()) > 0 {
		policyBytes, _ := json.Marshal(ts.MFAPolicy())
		jsonData := datatypes.JSON(policyBytes)
		mfaPolicy = &jsonData
	}

	return &models.TenantSetting{
		ID:                  ts.ID(), // id 直接引用 tenants.id
		EnforceMfa:          ts.EnforceMFA(), // Model 使用 EnforceMfa，Domain 使用 EnforceMFA
		AllowedEmailDomains: allowedEmailDomains,
		SessionTtlMinutes:   ts.SessionTTLMinutes(), // Model 使用 SessionTtlMinutes，Domain 使用 SessionTTLMinutes
		PasswordPolicy:      passwordPolicy,
		LoginPolicy:         loginPolicy,
		MfaPolicy:           mfaPolicy, // Model 使用 MfaPolicy，Domain 使用 MFAPolicy
		CreatedAt:           ts.CreatedAt(),
		UpdatedAt:           ts.UpdatedAt(),
		UpdatedBy:           ts.UpdatedBy(),
	}
}

// TenantSettingModelToDomain 将 Model TenantSetting 转换为 Domain TenantSetting
func TenantSettingModelToDomain(m *models.TenantSetting) *tenant_settings.TenantSetting {
	if m == nil {
		return nil
	}

	var allowedEmailDomains []string
	if m.AllowedEmailDomains != nil && *m.AllowedEmailDomains != "" {
		allowedEmailDomains = strings.Split(*m.AllowedEmailDomains, ",")
	}

	var passwordPolicy map[string]interface{}
	if m.PasswordPolicy != nil {
		json.Unmarshal(*m.PasswordPolicy, &passwordPolicy)
	}

	var loginPolicy map[string]interface{}
	if m.LoginPolicy != nil {
		json.Unmarshal(*m.LoginPolicy, &loginPolicy)
	}

	var mfaPolicy map[string]interface{}
	if m.MfaPolicy != nil {
		json.Unmarshal(*m.MfaPolicy, &mfaPolicy)
	}

	state := tenant_settings.TenantSettingState{
		ID:                 m.ID, // id 直接引用 tenants.id
		TenantID:           m.ID, // TenantID 从 ID 获取（一对一关系）
		EnforceMFA:         m.EnforceMfa, // Model 使用 EnforceMfa，Domain 使用 EnforceMFA
		AllowedEmailDomains: allowedEmailDomains,
		SessionTTLMinutes:  m.SessionTtlMinutes, // Model 使用 SessionTtlMinutes，Domain 使用 SessionTTLMinutes
		PasswordPolicy:     passwordPolicy,
		LoginPolicy:        loginPolicy,
		MFAPolicy:          mfaPolicy, // Model 使用 MfaPolicy，Domain 使用 MFAPolicy
		CreatedAt:          m.CreatedAt,
		UpdatedAt:          m.UpdatedAt,
		UpdatedBy:          m.UpdatedBy,
	}

	return tenant_settings.NewTenantSettingFromState(state)
}

// TenantSettingModelToUpdates 将 Model TenantSetting 转换为更新字段映射
func TenantSettingModelToUpdates(m *models.TenantSetting) map[string]any {
	var passwordPolicy any
	if m.PasswordPolicy != nil {
		passwordPolicy = m.PasswordPolicy
	}

	var loginPolicy any
	if m.LoginPolicy != nil {
		loginPolicy = m.LoginPolicy
	}

	var mfaPolicy any
	if m.MfaPolicy != nil {
		mfaPolicy = m.MfaPolicy
	}

	return map[string]any{
		// 注意：TenantID 不再存在，id 直接引用 tenants.id
		models.TenantSettingCols.EnforceMfa:         m.EnforceMfa,
		models.TenantSettingCols.AllowedEmailDomains: m.AllowedEmailDomains,
		models.TenantSettingCols.SessionTtlMinutes:   m.SessionTtlMinutes,
		models.TenantSettingCols.PasswordPolicy:     passwordPolicy,
		models.TenantSettingCols.LoginPolicy:        loginPolicy,
		models.TenantSettingCols.MfaPolicy:           mfaPolicy,
		models.TenantSettingCols.UpdatedAt:           m.UpdatedAt,
		models.TenantSettingCols.UpdatedBy:           m.UpdatedBy,
	}
}
